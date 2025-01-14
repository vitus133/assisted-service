// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UpdateHostReader is a Reader for the V2UpdateHost structure.
type V2UpdateHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UpdateHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2UpdateHostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UpdateHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UpdateHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UpdateHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UpdateHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UpdateHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2UpdateHostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UpdateHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UpdateHostCreated creates a V2UpdateHostCreated with default headers values
func NewV2UpdateHostCreated() *V2UpdateHostCreated {
	return &V2UpdateHostCreated{}
}

/*V2UpdateHostCreated handles this case with default header values.

Success.
*/
type V2UpdateHostCreated struct {
	Payload *models.Host
}

func (o *V2UpdateHostCreated) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostCreated  %+v", 201, o.Payload)
}

func (o *V2UpdateHostCreated) GetPayload() *models.Host {
	return o.Payload
}

func (o *V2UpdateHostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostBadRequest creates a V2UpdateHostBadRequest with default headers values
func NewV2UpdateHostBadRequest() *V2UpdateHostBadRequest {
	return &V2UpdateHostBadRequest{}
}

/*V2UpdateHostBadRequest handles this case with default header values.

Error.
*/
type V2UpdateHostBadRequest struct {
	Payload *models.Error
}

func (o *V2UpdateHostBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostBadRequest  %+v", 400, o.Payload)
}

func (o *V2UpdateHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostUnauthorized creates a V2UpdateHostUnauthorized with default headers values
func NewV2UpdateHostUnauthorized() *V2UpdateHostUnauthorized {
	return &V2UpdateHostUnauthorized{}
}

/*V2UpdateHostUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2UpdateHostUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UpdateHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostForbidden creates a V2UpdateHostForbidden with default headers values
func NewV2UpdateHostForbidden() *V2UpdateHostForbidden {
	return &V2UpdateHostForbidden{}
}

/*V2UpdateHostForbidden handles this case with default header values.

Forbidden.
*/
type V2UpdateHostForbidden struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostForbidden  %+v", 403, o.Payload)
}

func (o *V2UpdateHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostNotFound creates a V2UpdateHostNotFound with default headers values
func NewV2UpdateHostNotFound() *V2UpdateHostNotFound {
	return &V2UpdateHostNotFound{}
}

/*V2UpdateHostNotFound handles this case with default header values.

Error.
*/
type V2UpdateHostNotFound struct {
	Payload *models.Error
}

func (o *V2UpdateHostNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostNotFound  %+v", 404, o.Payload)
}

func (o *V2UpdateHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostMethodNotAllowed creates a V2UpdateHostMethodNotAllowed with default headers values
func NewV2UpdateHostMethodNotAllowed() *V2UpdateHostMethodNotAllowed {
	return &V2UpdateHostMethodNotAllowed{}
}

/*V2UpdateHostMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2UpdateHostMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2UpdateHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UpdateHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostConflict creates a V2UpdateHostConflict with default headers values
func NewV2UpdateHostConflict() *V2UpdateHostConflict {
	return &V2UpdateHostConflict{}
}

/*V2UpdateHostConflict handles this case with default header values.

Error.
*/
type V2UpdateHostConflict struct {
	Payload *models.Error
}

func (o *V2UpdateHostConflict) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostConflict  %+v", 409, o.Payload)
}

func (o *V2UpdateHostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInternalServerError creates a V2UpdateHostInternalServerError with default headers values
func NewV2UpdateHostInternalServerError() *V2UpdateHostInternalServerError {
	return &V2UpdateHostInternalServerError{}
}

/*V2UpdateHostInternalServerError handles this case with default header values.

Error.
*/
type V2UpdateHostInternalServerError struct {
	Payload *models.Error
}

func (o *V2UpdateHostInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2UpdateHostInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UpdateHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
