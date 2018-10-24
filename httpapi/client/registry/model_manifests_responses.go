// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/rai-project/dlframework/httpapi/models"
)

// ModelManifestsReader is a Reader for the ModelManifests structure.
type ModelManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModelManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModelManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModelManifestsOK creates a ModelManifestsOK with default headers values
func NewModelManifestsOK() *ModelManifestsOK {
	return &ModelManifestsOK{}
}

/*ModelManifestsOK handles this case with default header values.

A successful response.
*/
type ModelManifestsOK struct {
	Payload *models.DlframeworkModelManifestsResponse
}

func (o *ModelManifestsOK) Error() string {
	return fmt.Sprintf("[GET /registry/models/manifest][%d] modelManifestsOK  %+v", 200, o.Payload)
}

func (o *ModelManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkModelManifestsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
