// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/rai-project/dlframework/httpapi/models"
	"golang.org/x/net/context"
)

// NewImagesParams creates a new ImagesParams object
// with the default values initialized.
func NewImagesParams() *ImagesParams {
	var ()
	return &ImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImagesParamsWithTimeout creates a new ImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImagesParamsWithTimeout(timeout time.Duration) *ImagesParams {
	var ()
	return &ImagesParams{

		timeout: timeout,
	}
}

// NewImagesParamsWithContext creates a new ImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewImagesParamsWithContext(ctx context.Context) *ImagesParams {
	var ()
	return &ImagesParams{

		Context: ctx,
	}
}

// NewImagesParamsWithHTTPClient creates a new ImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImagesParamsWithHTTPClient(client *http.Client) *ImagesParams {
	var ()
	return &ImagesParams{
		HTTPClient: client,
	}
}

/*ImagesParams contains all the parameters to send to the API endpoint
for the images operation typically these are written to a http.Request
*/
type ImagesParams struct {

	/*Body*/
	Body *models.DlframeworkImagesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the images params
func (o *ImagesParams) WithTimeout(timeout time.Duration) *ImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images params
func (o *ImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images params
func (o *ImagesParams) WithContext(ctx context.Context) *ImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images params
func (o *ImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the images params
func (o *ImagesParams) WithHTTPClient(client *http.Client) *ImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the images params
func (o *ImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the images params
func (o *ImagesParams) WithBody(body *models.DlframeworkImagesRequest) *ImagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the images params
func (o *ImagesParams) SetBody(body *models.DlframeworkImagesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
