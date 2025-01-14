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

// V2ListClustersReader is a Reader for the V2ListClusters structure.
type V2ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ListClustersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2ListClustersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListClustersOK creates a V2ListClustersOK with default headers values
func NewV2ListClustersOK() *V2ListClustersOK {
	return &V2ListClustersOK{}
}

/*V2ListClustersOK handles this case with default header values.

Success.
*/
type V2ListClustersOK struct {
	Payload models.ClusterList
}

func (o *V2ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersOK  %+v", 200, o.Payload)
}

func (o *V2ListClustersOK) GetPayload() models.ClusterList {
	return o.Payload
}

func (o *V2ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClustersUnauthorized creates a V2ListClustersUnauthorized with default headers values
func NewV2ListClustersUnauthorized() *V2ListClustersUnauthorized {
	return &V2ListClustersUnauthorized{}
}

/*V2ListClustersUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2ListClustersUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListClustersUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClustersForbidden creates a V2ListClustersForbidden with default headers values
func NewV2ListClustersForbidden() *V2ListClustersForbidden {
	return &V2ListClustersForbidden{}
}

/*V2ListClustersForbidden handles this case with default header values.

Forbidden.
*/
type V2ListClustersForbidden struct {
	Payload *models.InfraError
}

func (o *V2ListClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersForbidden  %+v", 403, o.Payload)
}

func (o *V2ListClustersForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClustersMethodNotAllowed creates a V2ListClustersMethodNotAllowed with default headers values
func NewV2ListClustersMethodNotAllowed() *V2ListClustersMethodNotAllowed {
	return &V2ListClustersMethodNotAllowed{}
}

/*V2ListClustersMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2ListClustersMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2ListClustersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListClustersMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClustersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClustersInternalServerError creates a V2ListClustersInternalServerError with default headers values
func NewV2ListClustersInternalServerError() *V2ListClustersInternalServerError {
	return &V2ListClustersInternalServerError{}
}

/*V2ListClustersInternalServerError handles this case with default header values.

Error.
*/
type V2ListClustersInternalServerError struct {
	Payload *models.Error
}

func (o *V2ListClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListClustersServiceUnavailable creates a V2ListClustersServiceUnavailable with default headers values
func NewV2ListClustersServiceUnavailable() *V2ListClustersServiceUnavailable {
	return &V2ListClustersServiceUnavailable{}
}

/*V2ListClustersServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2ListClustersServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2ListClustersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/clusters][%d] v2ListClustersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ListClustersServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListClustersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
