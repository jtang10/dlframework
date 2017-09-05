// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rai-project/dlframework/httpapi/restapi/operations/predictor"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/registry"
)

// NewDlframeworkAPI creates a new Dlframework instance
func NewDlframeworkAPI(spec *loads.Document) *DlframeworkAPI {
	return &DlframeworkAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		PredictorDatasetHandler: predictor.DatasetHandlerFunc(func(params predictor.DatasetParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictorDataset has not yet been implemented")
		}),
		RegistryFrameworkAgentsHandler: registry.FrameworkAgentsHandlerFunc(func(params registry.FrameworkAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryFrameworkAgents has not yet been implemented")
		}),
		RegistryFrameworkManifestsHandler: registry.FrameworkManifestsHandlerFunc(func(params registry.FrameworkManifestsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryFrameworkManifests has not yet been implemented")
		}),
		PredictorImagesHandler: predictor.ImagesHandlerFunc(func(params predictor.ImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictorImages has not yet been implemented")
		}),
		RegistryModelAgentsHandler: registry.ModelAgentsHandlerFunc(func(params registry.ModelAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryModelAgents has not yet been implemented")
		}),
		RegistryModelManifestsHandler: registry.ModelManifestsHandlerFunc(func(params registry.ModelManifestsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryModelManifests has not yet been implemented")
		}),
		PredictorUrlsHandler: predictor.UrlsHandlerFunc(func(params predictor.UrlsParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictorUrls has not yet been implemented")
		}),
	}
}

/*DlframeworkAPI the dlframework API */
type DlframeworkAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PredictorDatasetHandler sets the operation handler for the dataset operation
	PredictorDatasetHandler predictor.DatasetHandler
	// RegistryFrameworkAgentsHandler sets the operation handler for the framework agents operation
	RegistryFrameworkAgentsHandler registry.FrameworkAgentsHandler
	// RegistryFrameworkManifestsHandler sets the operation handler for the framework manifests operation
	RegistryFrameworkManifestsHandler registry.FrameworkManifestsHandler
	// PredictorImagesHandler sets the operation handler for the images operation
	PredictorImagesHandler predictor.ImagesHandler
	// RegistryModelAgentsHandler sets the operation handler for the model agents operation
	RegistryModelAgentsHandler registry.ModelAgentsHandler
	// RegistryModelManifestsHandler sets the operation handler for the model manifests operation
	RegistryModelManifestsHandler registry.ModelManifestsHandler
	// PredictorUrlsHandler sets the operation handler for the urls operation
	PredictorUrlsHandler predictor.UrlsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *DlframeworkAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DlframeworkAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DlframeworkAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DlframeworkAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DlframeworkAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DlframeworkAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DlframeworkAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DlframeworkAPI
func (o *DlframeworkAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PredictorDatasetHandler == nil {
		unregistered = append(unregistered, "predictor.DatasetHandler")
	}

	if o.RegistryFrameworkAgentsHandler == nil {
		unregistered = append(unregistered, "registry.FrameworkAgentsHandler")
	}

	if o.RegistryFrameworkManifestsHandler == nil {
		unregistered = append(unregistered, "registry.FrameworkManifestsHandler")
	}

	if o.PredictorImagesHandler == nil {
		unregistered = append(unregistered, "predictor.ImagesHandler")
	}

	if o.RegistryModelAgentsHandler == nil {
		unregistered = append(unregistered, "registry.ModelAgentsHandler")
	}

	if o.RegistryModelManifestsHandler == nil {
		unregistered = append(unregistered, "registry.ModelManifestsHandler")
	}

	if o.PredictorUrlsHandler == nil {
		unregistered = append(unregistered, "predictor.UrlsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DlframeworkAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DlframeworkAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *DlframeworkAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *DlframeworkAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DlframeworkAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the dlframework API
func (o *DlframeworkAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DlframeworkAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/predict/dataset"] = predictor.NewDataset(o.context, o.PredictorDatasetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/registry/frameworks/agent"] = registry.NewFrameworkAgents(o.context, o.RegistryFrameworkAgentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/registry/frameworks/manifest"] = registry.NewFrameworkManifests(o.context, o.RegistryFrameworkManifestsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/predict/images"] = predictor.NewImages(o.context, o.PredictorImagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/registry/models/agent"] = registry.NewModelAgents(o.context, o.RegistryModelAgentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/registry/models/manifest"] = registry.NewModelManifests(o.context, o.RegistryModelManifestsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/predict/urls"] = predictor.NewUrls(o.context, o.PredictorUrlsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DlframeworkAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *DlframeworkAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
