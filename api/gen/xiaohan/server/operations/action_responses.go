// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ActionOKCode is the HTTP code returned for type ActionOK
const ActionOKCode int = 200

/*
ActionOK OK

swagger:response actionOK
*/
type ActionOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewActionOK creates ActionOK with default headers values
func NewActionOK() *ActionOK {

	return &ActionOK{}
}

// WithPayload adds the payload to the action o k response
func (o *ActionOK) WithPayload(payload string) *ActionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the action o k response
func (o *ActionOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ActionInternalServerErrorCode is the HTTP code returned for type ActionInternalServerError
const ActionInternalServerErrorCode int = 500

/*
ActionInternalServerError Internal Server Error

swagger:response actionInternalServerError
*/
type ActionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewActionInternalServerError creates ActionInternalServerError with default headers values
func NewActionInternalServerError() *ActionInternalServerError {

	return &ActionInternalServerError{}
}

// WithPayload adds the payload to the action internal server error response
func (o *ActionInternalServerError) WithPayload(payload string) *ActionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the action internal server error response
func (o *ActionInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}