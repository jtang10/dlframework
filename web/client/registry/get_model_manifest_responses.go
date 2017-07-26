// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rai-project/dlframework/web/models"
)

// GetModelManifestReader is a Reader for the GetModelManifest structure.
type GetModelManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetModelManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetModelManifestOK creates a GetModelManifestOK with default headers values
func NewGetModelManifestOK() *GetModelManifestOK {
	return &GetModelManifestOK{}
}

/*GetModelManifestOK handles this case with default header values.

GetModelManifestOK get model manifest o k
*/
type GetModelManifestOK struct {
	Payload *models.DlframeworkModelManifest
}

func (o *GetModelManifestOK) Error() string {
	return fmt.Sprintf("[GET /v1/model/{model_name}/info][%d] getModelManifestOK  %+v", 200, o.Payload)
}

func (o *GetModelManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkModelManifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
