package steps

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorgonia.org/tensor"

	"github.com/rai-project/dlframework/framework/options"
	"github.com/rai-project/dlframework/framework/predictor"
	"github.com/rai-project/pipeline"
	"github.com/rai-project/tracer"
)

type predict struct {
	base
	predictor predictor.Predictor
}

func NewPredict(predictor predictor.Predictor) pipeline.Step {
	res := predict{
		base: base{
			info: "Predict",
		},
	}
	res.predictor = predictor
	res.doer = res.do

	return res
}

func (p predict) do(ctx context.Context, in0 interface{}, pipelineOpts *pipeline.Options) interface{} {
	iData, ok := in0.([]interface{})
	if !ok {
		return errors.Errorf("expecting []interface{} for predict image step, but got %v", in0)
	}

	data, err := p.castToTensorType(iData)
	if err != nil {
		return err
	}

	if p.predictor == nil {
		return errors.New("the predict image step was created with a nil predictor")
	}

	opts, err := p.predictor.GetPredictionOptions(ctx)
	if err != nil {
		return err
	}

	framework, model, err := p.predictor.Info()
	if err != nil {
		return err
	}

	span, ctx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
		"model_name":        model.GetName(),
		"model_version":     model.GetVersion(),
		"framework_name":    framework.GetName(),
		"framework_version": framework.GetVersion(),
		"batch_size":        opts.BatchSize(),
		"feature_limit":     opts.FeatureLimit(),
		"device":            opts.Devices().String(),
		"trace_level":       opts.TraceLevel().String(),
		"uses_gpu":          opts.UsesGPU(),
	})
	defer span.Finish()

	err = p.predictor.Predict(ctx, data, options.WithOptions(opts), options.Context(ctx))
	if err != nil {
		return err
	}

	features, err := p.predictor.ReadPredictedFeatures(ctx)
	if err != nil {
		return err
	}

	lst := make([]interface{}, len(iData))
	for ii := 0; ii < len(iData); ii++ {
		lst[ii] = features[ii]
	}

	return lst
}

func (p predict) castToTensorType(inputs []interface{}) (interface{}, error) {
	data := make([]*tensor.Dense, len(inputs))
	for ii, input := range inputs {
		v, ok := input.(*tensor.Dense)
		if !ok {
			return nil, errors.Errorf("unable to cast to dense tensor in %v step", p.info)
		}
		data[ii] = v
	}
	return data, nil
}

func (p predict) castToElementType(inputs []interface{}) (interface{}, error) {
	_, model, _ := p.predictor.Info()
	switch t := model.GetElementType(); t {
	case "raw_image":
		data := make([][]byte, len(inputs))
		for ii, input := range inputs {
			r, err := toByteSlice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to uint8 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	case "uint8":
		data := make([][]uint8, len(inputs))
		for ii, input := range inputs {
			r, err := toUInt8Slice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to uint8 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	case "float32":
		data := make([][]float32, len(inputs))
		for ii, input := range inputs {
			r, err := toFloat32Slice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to float32 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	default:
		return nil, errors.Errorf("unsupported element type %v", t)
	}
}

func (p predict) Close() error {
	return nil
}
