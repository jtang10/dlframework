// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/rai-project/dlframework/httpapi/models"
)

// ResetReader is a Reader for the Reset structure.
type ResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResetOK creates a ResetOK with default headers values
func NewResetOK() *ResetOK {
	return &ResetOK{}
}

/*ResetOK handles this case with default header values.

A successful response.
*/
type ResetOK struct {
	Payload *models.DlframeworkResetResponse
}

func (o *ResetOK) Error() string {
	return fmt.Sprintf("[POST /predict/reset][%d] resetOK  %+v", 200, o.Payload)
}

func (o *ResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkResetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
