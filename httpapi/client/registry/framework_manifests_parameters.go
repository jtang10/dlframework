// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewFrameworkManifestsParams creates a new FrameworkManifestsParams object
// with the default values initialized.
func NewFrameworkManifestsParams() *FrameworkManifestsParams {
	var ()
	return &FrameworkManifestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFrameworkManifestsParamsWithTimeout creates a new FrameworkManifestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFrameworkManifestsParamsWithTimeout(timeout time.Duration) *FrameworkManifestsParams {
	var ()
	return &FrameworkManifestsParams{

		timeout: timeout,
	}
}

// NewFrameworkManifestsParamsWithContext creates a new FrameworkManifestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFrameworkManifestsParamsWithContext(ctx context.Context) *FrameworkManifestsParams {
	var ()
	return &FrameworkManifestsParams{

		Context: ctx,
	}
}

// NewFrameworkManifestsParamsWithHTTPClient creates a new FrameworkManifestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFrameworkManifestsParamsWithHTTPClient(client *http.Client) *FrameworkManifestsParams {
	var ()
	return &FrameworkManifestsParams{
		HTTPClient: client,
	}
}

/*FrameworkManifestsParams contains all the parameters to send to the API endpoint
for the framework manifests operation typically these are written to a http.Request
*/
type FrameworkManifestsParams struct {

	/*FrameworkName*/
	FrameworkName *string
	/*FrameworkVersion*/
	FrameworkVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the framework manifests params
func (o *FrameworkManifestsParams) WithTimeout(timeout time.Duration) *FrameworkManifestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the framework manifests params
func (o *FrameworkManifestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the framework manifests params
func (o *FrameworkManifestsParams) WithContext(ctx context.Context) *FrameworkManifestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the framework manifests params
func (o *FrameworkManifestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the framework manifests params
func (o *FrameworkManifestsParams) WithHTTPClient(client *http.Client) *FrameworkManifestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the framework manifests params
func (o *FrameworkManifestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrameworkName adds the frameworkName to the framework manifests params
func (o *FrameworkManifestsParams) WithFrameworkName(frameworkName *string) *FrameworkManifestsParams {
	o.SetFrameworkName(frameworkName)
	return o
}

// SetFrameworkName adds the frameworkName to the framework manifests params
func (o *FrameworkManifestsParams) SetFrameworkName(frameworkName *string) {
	o.FrameworkName = frameworkName
}

// WithFrameworkVersion adds the frameworkVersion to the framework manifests params
func (o *FrameworkManifestsParams) WithFrameworkVersion(frameworkVersion *string) *FrameworkManifestsParams {
	o.SetFrameworkVersion(frameworkVersion)
	return o
}

// SetFrameworkVersion adds the frameworkVersion to the framework manifests params
func (o *FrameworkManifestsParams) SetFrameworkVersion(frameworkVersion *string) {
	o.FrameworkVersion = frameworkVersion
}

// WriteToRequest writes these params to a swagger request
func (o *FrameworkManifestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FrameworkName != nil {

		// query param framework_name
		var qrFrameworkName string
		if o.FrameworkName != nil {
			qrFrameworkName = *o.FrameworkName
		}
		qFrameworkName := qrFrameworkName
		if qFrameworkName != "" {
			if err := r.SetQueryParam("framework_name", qFrameworkName); err != nil {
				return err
			}
		}

	}

	if o.FrameworkVersion != nil {

		// query param framework_version
		var qrFrameworkVersion string
		if o.FrameworkVersion != nil {
			qrFrameworkVersion = *o.FrameworkVersion
		}
		qFrameworkVersion := qrFrameworkVersion
		if qFrameworkVersion != "" {
			if err := r.SetQueryParam("framework_version", qFrameworkVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
