package agent

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rai-project/dldataset"
	dl "github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework/options"
	"github.com/rai-project/dlframework/framework/predictor"
	"github.com/rai-project/dlframework/steps"
	rgrpc "github.com/rai-project/grpc"
	"github.com/rai-project/pipeline"
	"github.com/rai-project/registry"
	"github.com/rai-project/tracer"
	"github.com/rai-project/utils"
	"github.com/rai-project/uuid"
	jaeger "github.com/uber/jaeger-client-go"
	"golang.org/x/sync/syncmap"
	"google.golang.org/grpc"

	_ "github.com/rai-project/dldataset/vision"
)

type Agent struct {
	base
	loadedPredictors syncmap.Map
	predictors       []predictor.Predictor
	options          *Options
	channelBuffer    int
}

var (
	DefaultChannelBuffer = 1000
)

func New(predictors []predictor.Predictor, opts ...Option) (*Agent, error) {
	options, err := NewOptions(opts...)
	if err != nil {
		return nil, err
	}
	framework, _, err := predictors[0].Info()
	if err != nil {
		return nil, err
	}

	return &Agent{
		base: base{
			Framework: framework,
		},
		predictors:    predictors,
		options:       options,
		channelBuffer: DefaultChannelBuffer,
	}, nil
}

// Opens a predictor and returns an id where the predictor
// is accessible. The id can be used to perform inference
// requests.
func (p *Agent) Open(ctx context.Context, req *dl.PredictorOpenRequest) (*dl.Predictor, error) {
	_, model, err := p.FindFrameworkModel(ctx, req)
	if err != nil {
		return nil, err
	}

	opts := req.GetOptions()
	if opts == nil {
		opts = &dl.PredictionOptions{}
	}

	level := tracer.LevelFromName(opts.GetExecutionOptions().GetTraceLevel().String())
	tracer.SetLevel(level)

	var predictorHandle predictor.Predictor
	for _, pred := range p.predictors {
		predModality, err := pred.Modality()
		if err != nil {
			continue
		}
		modelModality, err := model.Modality()
		if err != nil {
			continue
		}
		if predModality == modelModality {
			predictorHandle = pred
			break
		}
	}
	if predictorHandle == nil {
		return nil, errors.New("unable to find predictor for requested modality")
	}

	predictor, err := predictorHandle.Load(
		ctx,
		*model,
		options.PredictorOptions(opts),
		options.DisableFrameworkAutoTuning(true),
		options.InferencePrecision("fp32"),
	)
	if err != nil {
		return nil, err
	}
	if predictor == nil {
		return nil, errors.New("predictor in Open is nil")
	}

	id := uuid.NewV4()
	p.loadedPredictors.Store(id, predictor)

	return &dl.Predictor{ID: id}, nil
}

func (p *Agent) getLoadedPredictor(id string) (predictor.Predictor, error) {
	val, ok := p.loadedPredictors.Load(id)
	if !ok {
		return nil, errors.Errorf("predictor %v was not found", id)
	}

	predictor, ok := val.(predictor.Predictor)

	if !ok {
		return nil, errors.Errorf("predictor %v is not a valid image predictor", predictor)
	}

	return predictor, nil
}

// Close a predictor and clear it's memory.
func (p *Agent) Close(ctx context.Context, req *dl.PredictorCloseRequest) (*dl.PredictorCloseResponse, error) {
	id := req.GetPredictor().GetID()
	predictor, err := p.getLoadedPredictor(id)
	if err != nil {
		return nil, err
	}
	predictor.Close()
	p.loadedPredictors.Delete(id)

	return &dl.PredictorCloseResponse{}, nil
}

func getTraceID(ctx context.Context) (string, error) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return "", errors.New("unable to find span")
	}
	jaegerSpan, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		return "", errors.New("invalid cast")
	}

	return strconv.FormatUint(jaegerSpan.TraceID().Low, 16), nil
}

func (p *Agent) toFeaturesResponse(ctx context.Context, output <-chan interface{}, options *dl.PredictionOptions) (*dl.FeaturesResponse, error) {
	res := &dl.FeaturesResponse{}

	if traceId, err := getTraceID(ctx); err == nil {
		res.TraceId = &dl.TraceID{
			Id: traceId,
		}
	}

	if options == nil {
		options = &dl.PredictionOptions{
			RequestID: "request-id-not-found",
		}
	}

	for out := range output {
		if err, ok := out.(error); ok {
			return nil, err
		}

		var features []*dl.Feature

		inputId := "undefined-input-id"

		switch out := out.(type) {
		case steps.IDer:
			inputId = out.GetID()
			data := out.GetData()
			switch data := data.(type) {
			case []*dl.Feature:
				features = data
			case dl.Features:
				features = []*dl.Feature(data)
			default:
				return nil, errors.Errorf("expecting a []*Feature type, but got %v", data)
			}
		case []*dl.Feature:
			features = out
		case dl.Features:
			features = []*dl.Feature(out)
		default:
			return nil, errors.Errorf("expecting an ider or []*Feature type, but got %v", reflect.TypeOf(out))
		}

		res.Responses = append(res.Responses, &dl.FeatureResponse{
			ID:        uuid.NewV4(),
			InputID:   inputId,
			RequestID: options.GetRequestID(),
			Features:  features,
		})
	}

	return res, nil
}

func (p *Agent) urls(ctx context.Context, req *dl.URLsRequest) (<-chan interface{}, error) {
	if req.GetPredictor() == nil {
		return nil, errors.New("request does not have a valid predictor set")
	}

	predictorId := req.GetPredictor().GetID()
	if predictorId == "" {
		return nil, errors.New("predictor id cannot be an empty string")
	}

	predictor, err := p.getLoadedPredictor(predictorId)
	if err != nil {
		return nil, err
	}

	preprocessOptions, err := predictor.GetPreprocessOptions()
	if err != nil {
		return nil, err
	}

	input := make(chan interface{}, p.channelBuffer)
	go func() {
		defer close(input)
		for _, url := range req.GetUrls() {
			input <- url
		}
	}()

	output := pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(p.channelBuffer)).
		Then(steps.NewReadURL()).
		Then(steps.NewReadImage(preprocessOptions)).
		Then(steps.NewPreprocessImage(preprocessOptions)).
		Run(input)

	var outputs []interface{}
	for out := range output {
		outputs = append(outputs, out)
	}

	predictionOptions, err := predictor.GetPredictionOptions()
	if err != nil {
		return nil, err
	}

	batchSize := predictionOptions.BatchSize()

	parts := dl.Partition(outputs, batchSize)

	input = make(chan interface{}, p.channelBuffer)
	go func() {
		defer close(input)
		for _, part := range parts {
			input <- part
		}
	}()

	output = pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(p.channelBuffer)).
		Then(steps.NewPredict(predictor)).
		Run(input)

	return output, nil
}

// URLs method receives a list of urls and runs
// the predictor on all the urls.
// The result is a list of predicted features for all the urls.
func (p *Agent) URLs(ctx context.Context, req *dl.URLsRequest) (*dl.FeaturesResponse, error) {
	output, err := p.urls(ctx, req)
	if err != nil {
		return nil, err
	}

	return p.toFeaturesResponse(ctx, output, req.GetOptions())
}

// URLsStream method receives a stream of urls and runs
// the predictor on all the urls.
// The result is a prediction feature stream for each url.
// func (p *Agent) URLsStream(req *dl.URLsRequest, svr dl.Predict_URLsStreamServer) error {
// 	ctx := svr.Context()
// 	output, err := p.urls(ctx, req)
// 	if err != nil {
// 		return err
// 	}

// 	for o := range output {
// 		svr.Send(o.(*dl.FeatureResponse))
// 	}

// 	return nil
// }

func (p *Agent) images(ctx context.Context, req *dl.ImagesRequest) (<-chan interface{}, error) {
	if req.GetPredictor() == nil {
		return nil, errors.New("request does not have a valid predictor set")
	}
	predictorId := req.GetPredictor().GetID()
	if predictorId == "" {
		return nil, errors.New("predictor id cannot be an empty string")
	}

	predictor, err := p.getLoadedPredictor(predictorId)
	if err != nil {
		return nil, err
	}

	preprocessOptions, err := predictor.GetPreprocessOptions()
	if err != nil {
		return nil, err
	}

	input := make(chan interface{}, p.channelBuffer)
	go func() {
		defer close(input)
		for _, img := range req.GetImages() {
			input <- img
		}
	}()

	output := pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(p.channelBuffer)).
		Then(steps.NewReadImage(preprocessOptions)).
		Then(steps.NewPreprocessImage(preprocessOptions)).
		Run(input)

	var outputs []interface{}
	for out := range output {
		outputs = append(outputs, out)
	}

	opts, err := predictor.GetPredictionOptions()
	if err != nil {
		return nil, err
	}

	batchSize := int(opts.BatchSize())

	parts := dl.Partition(outputs, batchSize)

	input = make(chan interface{}, p.channelBuffer)
	go func() {
		defer close(input)
		for _, part := range parts {
			input <- part
		}
	}()

	output = pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(p.channelBuffer)).
		Then(steps.NewPredict(predictor)).
		Run(input)

	return output, nil
}

// Image method receives a list base64 encoded images and runs
// the predictor on all the images.
// The result is a prediction feature list for each image.
func (p *Agent) Images(ctx context.Context, req *dl.ImagesRequest) (*dl.FeaturesResponse, error) {
	output, err := p.images(ctx, req)
	if err != nil {
		return nil, err
	}

	return p.toFeaturesResponse(ctx, output, req.GetOptions())
}

// Image method receives a list base64 encoded images and runs
// the predictor on all the images.
// The result is a prediction feature stream for each image.
// func (p *Agent) ImagesStream(req *dl.ImagesRequest, svr dl.Predict_ImagesStreamServer) error {
// 	ctx := svr.Context()
// 	output, err := p.images(ctx, req)
// 	if err != nil {
// 		return err
// 	}

// 	for o := range output {
// 		svr.Send(o.(*dl.FeatureResponse))
// 	}

// 	return nil
// }

// Dataset method receives a single dataset and runs
// the predictor on all elements of the dataset.
// The result is a prediction feature list.
func (p *Agent) dataset(ctx context.Context, req *dl.DatasetRequest) (<-chan interface{}, error) {
	if req.GetPredictor() == nil {
		return nil, errors.New("request does not have a valid predictor set")
	}
	predictorId := req.GetPredictor().GetID()
	if predictorId == "" {
		return nil, errors.New("predictor id cannot be an empty string")
	}

	predictor, err := p.getLoadedPredictor(predictorId)
	if err != nil {
		return nil, err
	}

	preprocessOptions, err := predictor.GetPreprocessOptions()
	if err != nil {
		return nil, err
	}

	if req.GetDataset() == nil {
		return nil, errors.New("invalid empty dataset parameter in request")
	}

	dataset, err := dldataset.Get(req.Dataset.GetCategory(), req.Dataset.GetName())
	if err != nil {
		return nil, err
	}

	dataset, err = dataset.New(ctx)
	if err != nil {
		return nil, err
	}

	defer dataset.Close()

	if err := dataset.Download(ctx); err != nil {
		return nil, err
	}

	elems, err := dataset.List(ctx)
	if err != nil {
		return nil, err
	}

	input := make(chan interface{}, p.channelBuffer)
	go func() {
		defer close(input)
		for _, e := range elems {
			input <- e
		}
	}()

	output := pipeline.New(pipeline.Context(ctx), pipeline.ChannelBuffer(p.channelBuffer)).
		Then(steps.NewGetDataset(dataset)).
		Then(steps.NewReadImage(preprocessOptions)).
		Then(steps.NewPreprocessImage(preprocessOptions)).
		Then(steps.NewPredict(predictor)).
		Run(input)

	return output, nil
}

// Dataset method receives a single dataset and runs
// the predictor on all elements of the dataset.
// The result is a prediction feature list.
func (p *Agent) Dataset(ctx context.Context, req *dl.DatasetRequest) (*dl.FeaturesResponse, error) {
	output, err := p.dataset(ctx, req)
	if err != nil {
		return nil, err
	}

	return p.toFeaturesResponse(ctx, output, req.GetOptions())
}

// Dataset method receives a single dataset and runs
// the predictor on all elements of the dataset.
// The result is a prediction feature stream.
// func (p *Agent) DatasetStream(req *dl.DatasetRequest, svr dl.Predict_DatasetStreamServer) error {
// 	ctx := svr.Context()
// 	output, err := p.dataset(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	for o := range output {
// 		svr.Send(o.(*dl.FeatureResponse))
// 	}

// 	return nil
// }

// Clear method clears the internal cache of the predictors
func (p *Agent) Reset(ctx context.Context, req *dl.ResetRequest) (*dl.ResetResponse, error) {
	return nil, nil
}

func (p *Agent) FindFrameworkModel(ctx context.Context, req *dl.PredictorOpenRequest) (*dl.FrameworkManifest, *dl.ModelManifest, error) {
	framework, err := dl.FindFramework(req.GetFrameworkName() + ":" + req.GetFrameworkVersion())
	if err != nil {
		return nil, nil, err
	}
	model, err := framework.FindModel(req.GetModelName() + ":" + req.GetModelVersion())
	if err != nil {
		return nil, nil, err
	}

	return framework, model, nil
}

func (p *Agent) RegisterManifests() (*grpc.Server, error) {
	log.Info("populating registry")
	var grpcServer *grpc.Server
	grpcServer = rgrpc.NewServer(dl.RegistryServiceDescription)
	svr := &Registry{
		base: base{
			Framework: p.base.Framework,
		},
	}
	go func() {
		utils.Every(
			registry.Config.Timeout/2,
			func() {
				svr.PublishInRegistery()
			},
		)
	}()
	dl.RegisterRegistryServer(grpcServer, svr)

	return grpcServer, nil
}

func (p *Agent) RegisterPredictor() (*grpc.Server, error) {
	grpcServer := rgrpc.NewServer(dl.PredictServiceDescription)

	host := fmt.Sprintf("%s:%d", p.options.host, p.options.port)
	log.Info("registering predictor service at ", host)

	go func() {
		utils.Every(
			registry.Config.Timeout/2,
			func() {
				p.base.PublishInPredictor(host, "predictor")
			},
		)
	}()
	dl.RegisterPredictServer(grpcServer, p)

	return grpcServer, nil
}
