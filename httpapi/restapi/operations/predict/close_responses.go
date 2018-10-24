// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	models "github.com/rai-project/dlframework/httpapi/models"
)

// CloseOKCode is the HTTP code returned for type CloseOK
const CloseOKCode int = 200

/*CloseOK A successful response.

swagger:response closeOK
*/
type CloseOK struct {

	/*
	  In: Body
	*/
	Payload models.DlframeworkPredictorCloseResponse `json:"body,omitempty"`
}

// NewCloseOK creates CloseOK with default headers values
func NewCloseOK() *CloseOK {

	return &CloseOK{}
}

// WithPayload adds the payload to the close o k response
func (o *CloseOK) WithPayload(payload models.DlframeworkPredictorCloseResponse) *CloseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the close o k response
func (o *CloseOK) SetPayload(payload models.DlframeworkPredictorCloseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
