// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2UploadLogsNoContentCode is the HTTP code returned for type V2UploadLogsNoContent
const V2UploadLogsNoContentCode int = 204

/*V2UploadLogsNoContent Success.

swagger:response v2UploadLogsNoContent
*/
type V2UploadLogsNoContent struct {
}

// NewV2UploadLogsNoContent creates V2UploadLogsNoContent with default headers values
func NewV2UploadLogsNoContent() *V2UploadLogsNoContent {

	return &V2UploadLogsNoContent{}
}

// WriteResponse to the client
func (o *V2UploadLogsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// V2UploadLogsUnauthorizedCode is the HTTP code returned for type V2UploadLogsUnauthorized
const V2UploadLogsUnauthorizedCode int = 401

/*V2UploadLogsUnauthorized Unauthorized.

swagger:response v2UploadLogsUnauthorized
*/
type V2UploadLogsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2UploadLogsUnauthorized creates V2UploadLogsUnauthorized with default headers values
func NewV2UploadLogsUnauthorized() *V2UploadLogsUnauthorized {

	return &V2UploadLogsUnauthorized{}
}

// WithPayload adds the payload to the v2 upload logs unauthorized response
func (o *V2UploadLogsUnauthorized) WithPayload(payload *models.InfraError) *V2UploadLogsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs unauthorized response
func (o *V2UploadLogsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsForbiddenCode is the HTTP code returned for type V2UploadLogsForbidden
const V2UploadLogsForbiddenCode int = 403

/*V2UploadLogsForbidden Forbidden.

swagger:response v2UploadLogsForbidden
*/
type V2UploadLogsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2UploadLogsForbidden creates V2UploadLogsForbidden with default headers values
func NewV2UploadLogsForbidden() *V2UploadLogsForbidden {

	return &V2UploadLogsForbidden{}
}

// WithPayload adds the payload to the v2 upload logs forbidden response
func (o *V2UploadLogsForbidden) WithPayload(payload *models.InfraError) *V2UploadLogsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs forbidden response
func (o *V2UploadLogsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsNotFoundCode is the HTTP code returned for type V2UploadLogsNotFound
const V2UploadLogsNotFoundCode int = 404

/*V2UploadLogsNotFound Error.

swagger:response v2UploadLogsNotFound
*/
type V2UploadLogsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UploadLogsNotFound creates V2UploadLogsNotFound with default headers values
func NewV2UploadLogsNotFound() *V2UploadLogsNotFound {

	return &V2UploadLogsNotFound{}
}

// WithPayload adds the payload to the v2 upload logs not found response
func (o *V2UploadLogsNotFound) WithPayload(payload *models.Error) *V2UploadLogsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs not found response
func (o *V2UploadLogsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsMethodNotAllowedCode is the HTTP code returned for type V2UploadLogsMethodNotAllowed
const V2UploadLogsMethodNotAllowedCode int = 405

/*V2UploadLogsMethodNotAllowed Method Not Allowed.

swagger:response v2UploadLogsMethodNotAllowed
*/
type V2UploadLogsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UploadLogsMethodNotAllowed creates V2UploadLogsMethodNotAllowed with default headers values
func NewV2UploadLogsMethodNotAllowed() *V2UploadLogsMethodNotAllowed {

	return &V2UploadLogsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 upload logs method not allowed response
func (o *V2UploadLogsMethodNotAllowed) WithPayload(payload *models.Error) *V2UploadLogsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs method not allowed response
func (o *V2UploadLogsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsConflictCode is the HTTP code returned for type V2UploadLogsConflict
const V2UploadLogsConflictCode int = 409

/*V2UploadLogsConflict Error.

swagger:response v2UploadLogsConflict
*/
type V2UploadLogsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UploadLogsConflict creates V2UploadLogsConflict with default headers values
func NewV2UploadLogsConflict() *V2UploadLogsConflict {

	return &V2UploadLogsConflict{}
}

// WithPayload adds the payload to the v2 upload logs conflict response
func (o *V2UploadLogsConflict) WithPayload(payload *models.Error) *V2UploadLogsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs conflict response
func (o *V2UploadLogsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsInternalServerErrorCode is the HTTP code returned for type V2UploadLogsInternalServerError
const V2UploadLogsInternalServerErrorCode int = 500

/*V2UploadLogsInternalServerError Error.

swagger:response v2UploadLogsInternalServerError
*/
type V2UploadLogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UploadLogsInternalServerError creates V2UploadLogsInternalServerError with default headers values
func NewV2UploadLogsInternalServerError() *V2UploadLogsInternalServerError {

	return &V2UploadLogsInternalServerError{}
}

// WithPayload adds the payload to the v2 upload logs internal server error response
func (o *V2UploadLogsInternalServerError) WithPayload(payload *models.Error) *V2UploadLogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs internal server error response
func (o *V2UploadLogsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UploadLogsServiceUnavailableCode is the HTTP code returned for type V2UploadLogsServiceUnavailable
const V2UploadLogsServiceUnavailableCode int = 503

/*V2UploadLogsServiceUnavailable Unavailable.

swagger:response v2UploadLogsServiceUnavailable
*/
type V2UploadLogsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UploadLogsServiceUnavailable creates V2UploadLogsServiceUnavailable with default headers values
func NewV2UploadLogsServiceUnavailable() *V2UploadLogsServiceUnavailable {

	return &V2UploadLogsServiceUnavailable{}
}

// WithPayload adds the payload to the v2 upload logs service unavailable response
func (o *V2UploadLogsServiceUnavailable) WithPayload(payload *models.Error) *V2UploadLogsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 upload logs service unavailable response
func (o *V2UploadLogsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UploadLogsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
