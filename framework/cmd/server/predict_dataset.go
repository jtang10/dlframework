package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"github.com/levigross/grequests"
	"github.com/pkg/errors"
	"github.com/rai-project/caffe/predict"
	"github.com/rai-project/config"
	mongodb "github.com/rai-project/database/mongodb"
	"github.com/rai-project/dldataset"
	dl "github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework/agent"
	"github.com/rai-project/dlframework/framework/options"
	"github.com/rai-project/dlframework/steps"
	"github.com/rai-project/evaluation"
	nvidiasmi "github.com/rai-project/nvidia-smi"
	"github.com/rai-project/pipeline"
	"github.com/rai-project/tracer"
	"github.com/rai-project/uuid"
	"github.com/schollz/progressbar"
	"github.com/spf13/cobra"
	jaeger "github.com/uber/jaeger-client-go"
	"gopkg.in/mgo.v2/bson"
)

var (
	datasetCategory      string
	datasetName          string
	modelName            string
	modelVersion         string
	batchSize            int
	partitionDatasetSize int
	publishEvaluation    bool
	traceLevel           tracer.Level = tracer.STEP_TRACE
)

var (
	DefaultChannelBuffer = 100000
)

var datasetCmd = &cobra.Command{
	Use:   "dataset",
	Short: "dataset",
	RunE: func(c *cobra.Command, args []string) error {
		defer tracer.Close()
		span, ctx := tracer.StartSpanFromContext(context.Background(), traceLevel, "dataset")
		defer span.Finish()

		db, err := mongodb.NewDatabase(config.App.Name)
		defer db.Close()

		predictorFramework, err := agent.GetPredictor(framework)
		if err != nil {
			return errors.Wrapf(err,
				"⚠️ failed to get predictor for %s. make sure you have "+
					"imported the framework's predictor package",
				framework.MustCanonicalName(),
			)
		}

		model, err := framework.FindModel(modelName + ":" + modelVersion)
		if err != nil {
			return err
		}

		var dc map[string]int32
		if nvidiasmi.HasGPU {
			dc = map[string]int32{"GPU": 0}
		} else {
			dc = map[string]int32{"CPU": 0}
		}

		execOpts := &dl.ExecutionOptions{
			TraceLevel: dl.ExecutionOptions_TraceLevel(
				dl.ExecutionOptions_TraceLevel_value[traceLevel.String()],
			),
			DeviceCount: dc,
		}

		predOpts := &dl.PredictionOptions{
			// FeatureLimit:     5,
			BatchSize:        uint32(batchSize),
			ExecutionOptions: execOpts,
		}

		predictor, err := predictorFramework.Load(ctx, *model, options.PredictorOptions(predOpts))
		if err != nil {
			return err
		}

		if datasetName == "ilsvrc2012_validation" {
			imagePredictor, ok := predictor.(predict.ImagePredictor)
			if !ok {
				return errors.Errorf("expecting an image predictor for %v", model.GetCannonicalName())
			}
			dims, err := predictor.GetImageDimensions(ctx)
			if err != nil {
				return err
			}
			if len(dims) != 3 {
				return errors.Errorf("expecting a 3 element vector for dimensions %v", dims)
			}
			width, height := dims[1], dims[2]
			if width != height {
				return errors.Errorf("expecting a square image dimensions width = %v, height = %v", width, height)
			}

			imageDimensionString := strconv.Itoa(width)
			datasetName = datasetName + "_" + imageDimensionString
		}

		log.WithField("dataset_category", datasetCategory).
			WithField("dataset_name", datasetName).
			Debug("using specified dataset")

		dataset, err := dldataset.Get(datasetCategory, datasetName)
		if err != nil {
			return err
		}
		defer dataset.Close()

		err = dataset.Download(ctx)
		if err != nil {
			return err
		}

		fileList, err := dataset.List(ctx)
		if err != nil {
			return err
		}

		err = dataset.Load(ctx)
		if err != nil {
			return err
		}

		inputPredictionIds := []bson.ObjectId{}

		evaluationEntry := evaluation.Evaluation{
			ID:              bson.NewObjectId(),
			CreatedAt:       time.Now(),
			Framework:       *model.GetFramework(),
			Model:           *model,
			DatasetCategory: dataset.Category(),
			DatasetName:     dataset.Name(),
			Public:          false,
		}

		evaluationTable, err := mongodb.NewTable(db, evaluationEntry.TableName())
		if err != nil {
			return err
		}
		evaluationTable.Create(nil)

		modelAccuracyTable, err := mongodb.NewTable(db, evaluation.ModelAccuracy{}.TableName())
		if err != nil {
			return err
		}
		modelAccuracyTable.Create(nil)

		performanceTable, err := mongodb.NewTable(db, evaluation.Performance{}.TableName())
		if err != nil {
			return err
		}
		performanceTable.Create(nil)

		inputPredictionsTable, err := mongodb.NewTable(db, evaluation.InputPrediction{}.TableName())
		if err != nil {
			return err
		}
		inputPredictionsTable.Create(nil)

		preprocessOptions, err := predictor.GetPreprocessOptions(ctx)
		if err != nil {
			return err
		}
		_ = preprocessOptions

		fileNameParts := partitionDataset(fileList, partitionDatasetSize)

		var cntTop1 = 0
		var cntTop5 = 0

		outputs := make(chan interface{}, DefaultChannelBuffer)
		partlabels := map[string]string{}

		progress := progressbar.New(len(fileList))
		for _, part := range fileNameParts[0:16] {
			input := make(chan interface{}, DefaultChannelBuffer)
			go func() {
				defer close(input)
				for _, fileName := range part {
					lda, err := dataset.Get(ctx, fileName)
					if err != nil {
						continue
					}
					id := uuid.NewV4()
					lbl := steps.NewIDWrapper(id, lda)
					partlabels[lbl.GetID()] = lda.Label()
					input <- lbl
				}
			}()

			output := pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(DefaultChannelBuffer)).
				Then(steps.NewReadImage(preprocessOptions)).
				Then(steps.NewPreprocessImage(preprocessOptions)).
				Run(input)

			var images []interface{}
			for out := range output {
				images = append(images, out)
			}

			parts := agent.Partition(images, batchSize)

			input = make(chan interface{}, DefaultChannelBuffer)
			go func() {
				defer close(input)
				for _, part := range parts {
					input <- part
				}
			}()

			output = pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(DefaultChannelBuffer)).
				Then(steps.NewPredictImage(predictor)).
				Run(input)

			for o := range output {
				progress.Add(batchSize)
				outputs <- o
			}

		}

		close(outputs)

		for out0 := range outputs {
			out, ok := out0.(steps.IDer)
			if !ok {
				return errors.Errorf("expecting steps.IDer, but got %v", out0)
			}
			_ = out
			id := out.GetID()
			label := partlabels[id]

			features := out.GetData().(dl.Features)
			if !ok {
				return errors.Errorf("expecting a dlframework.Features type, but got %v", out.GetData())
			}

			inputPrediction := evaluation.InputPrediction{
				ID:            bson.NewObjectId(),
				CreatedAt:     time.Now(),
				InputID:       id,
				ExpectedLabel: label,
				Features:      features,
			}
			if err := inputPredictionsTable.Insert(inputPrediction); err != nil {
				log.WithError(err).Error("failed to publish input prediction entry")
			}
			inputPredictionIds = append(inputPredictionIds, inputPrediction.ID)

			features.Sort()

			if strings.Fields(features[0].GetName())[0] == label {
				cntTop1++
			}
			for _, f := range features[:5] {
				if strings.Fields(f.GetName())[0] == label {
					cntTop5++
				}
			}
		}

		modelAccuracy := evaluation.ModelAccuracy{
			ID:        bson.NewObjectId(),
			CreatedAt: time.Now(),
			Top1:      float64(cntTop1) / float64(len(fileList)),
			Top5:      float64(cntTop5) / float64(len(fileList)),
		}
		if err := modelAccuracyTable.Insert(modelAccuracy); err != nil {
			log.WithError(err).Error("failed to publish model accuracy entry")
		}

		traceID := span.Context().(jaeger.SpanContext).TraceID()
		query := fmt.Sprintf("http://localhost:16686/api/traces/%v", strconv.FormatUint(traceID.Low, 16))
		resp, err := grequests.Get(query, nil)

		if err == nil {
			var trace evaluation.TraceInformation
			dec := json.NewDecoder(resp)
			if err := dec.Decode(&trace); err != nil {
				log.WithError(err).Error("failed to decode trace information")
			}
			performance := evaluation.Performance{
				ID:         bson.NewObjectId(),
				CreatedAt:  time.Now(),
				Trace:      trace,
				TraceLevel: traceLevel,
			}
			evaluationEntry.PerformanceID = performance.ID
			performanceTable.Insert(performance)
		}
		evaluationEntry.ModelAccuracyID = modelAccuracy.ID
		evaluationEntry.InputPredictionIDs = inputPredictionIds

		if err := evaluationTable.Insert(evaluationEntry); err != nil {
			log.WithError(err).Error("failed to publish evaluation entry")
		}

		pp.Println("cntTop1 = ", cntTop1, "cntTop5 = ", cntTop5)

		return nil
	},
}

func partitionDataset(in []string, partitionSize int) (out [][]string) {
	cnt := (len(in)-1)/partitionSize + 1
	for i := 0; i < cnt; i++ {
		start := i * partitionSize
		end := (i + 1) * partitionSize
		if end > len(in) {
			end = len(in)
		}
		part := in[start:end]
		out = append(out, part)
	}

	return out
}

func init() {
	datasetCmd.PersistentFlags().StringVar(&datasetCategory, "dataset_category", "vision", "dataset category (e.g. \"vision\")")
	datasetCmd.PersistentFlags().StringVar(&datasetName, "dataset_name", "ilsvrc2012_validation", "dataset name (e.g. \"ilsvrc2012_validation_folder\")")
	datasetCmd.PersistentFlags().StringVar(&modelName, "modelName", "BVLC-AlexNet", "modelName")
	datasetCmd.PersistentFlags().StringVar(&modelVersion, "modelVersion", "1.0", "modelVersion")
	datasetCmd.PersistentFlags().IntVarP(&batchSize, "batchSize", "b", 64, "batch size")
	datasetCmd.PersistentFlags().BoolVar(&publishEvaluation, "publish", true, "publish evaluation to database")
	datasetCmd.PersistentFlags().IntVarP(&partitionDatasetSize, "partitionDatasetSize", "p", 64, "partition dataset size")
}