//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsUpdateReader is a Reader for the ObjectsUpdate structure.
type ObjectsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObjectsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObjectsUpdateOK creates a ObjectsUpdateOK with default headers values
func NewObjectsUpdateOK() *ObjectsUpdateOK {
	return &ObjectsUpdateOK{}
}

/*
ObjectsUpdateOK describes a response with status code 200, with default header values.

Successfully received.
*/
type ObjectsUpdateOK struct {
	Payload *models.Object
}

// IsSuccess returns true when this objects update o k response has a 2xx status code
func (o *ObjectsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this objects update o k response has a 3xx status code
func (o *ObjectsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update o k response has a 4xx status code
func (o *ObjectsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects update o k response has a 5xx status code
func (o *ObjectsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this objects update o k response a status code equal to that given
func (o *ObjectsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the objects update o k response
func (o *ObjectsUpdateOK) Code() int {
	return 200
}

func (o *ObjectsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateOK  %+v", 200, o.Payload)
}

func (o *ObjectsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateOK  %+v", 200, o.Payload)
}

func (o *ObjectsUpdateOK) GetPayload() *models.Object {
	return o.Payload
}

func (o *ObjectsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Object)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsUpdateUnauthorized creates a ObjectsUpdateUnauthorized with default headers values
func NewObjectsUpdateUnauthorized() *ObjectsUpdateUnauthorized {
	return &ObjectsUpdateUnauthorized{}
}

/*
ObjectsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsUpdateUnauthorized struct {
}

// IsSuccess returns true when this objects update unauthorized response has a 2xx status code
func (o *ObjectsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects update unauthorized response has a 3xx status code
func (o *ObjectsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update unauthorized response has a 4xx status code
func (o *ObjectsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects update unauthorized response has a 5xx status code
func (o *ObjectsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this objects update unauthorized response a status code equal to that given
func (o *ObjectsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the objects update unauthorized response
func (o *ObjectsUpdateUnauthorized) Code() int {
	return 401
}

func (o *ObjectsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateUnauthorized ", 401)
}

func (o *ObjectsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateUnauthorized ", 401)
}

func (o *ObjectsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsUpdateForbidden creates a ObjectsUpdateForbidden with default headers values
func NewObjectsUpdateForbidden() *ObjectsUpdateForbidden {
	return &ObjectsUpdateForbidden{}
}

/*
ObjectsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ObjectsUpdateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects update forbidden response has a 2xx status code
func (o *ObjectsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects update forbidden response has a 3xx status code
func (o *ObjectsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update forbidden response has a 4xx status code
func (o *ObjectsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects update forbidden response has a 5xx status code
func (o *ObjectsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this objects update forbidden response a status code equal to that given
func (o *ObjectsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the objects update forbidden response
func (o *ObjectsUpdateForbidden) Code() int {
	return 403
}

func (o *ObjectsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsUpdateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsUpdateNotFound creates a ObjectsUpdateNotFound with default headers values
func NewObjectsUpdateNotFound() *ObjectsUpdateNotFound {
	return &ObjectsUpdateNotFound{}
}

/*
ObjectsUpdateNotFound describes a response with status code 404, with default header values.

Successful query result but no resource was found.
*/
type ObjectsUpdateNotFound struct {
}

// IsSuccess returns true when this objects update not found response has a 2xx status code
func (o *ObjectsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects update not found response has a 3xx status code
func (o *ObjectsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update not found response has a 4xx status code
func (o *ObjectsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects update not found response has a 5xx status code
func (o *ObjectsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this objects update not found response a status code equal to that given
func (o *ObjectsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the objects update not found response
func (o *ObjectsUpdateNotFound) Code() int {
	return 404
}

func (o *ObjectsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateNotFound ", 404)
}

func (o *ObjectsUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateNotFound ", 404)
}

func (o *ObjectsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsUpdateUnprocessableEntity creates a ObjectsUpdateUnprocessableEntity with default headers values
func NewObjectsUpdateUnprocessableEntity() *ObjectsUpdateUnprocessableEntity {
	return &ObjectsUpdateUnprocessableEntity{}
}

/*
ObjectsUpdateUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ObjectsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects update unprocessable entity response has a 2xx status code
func (o *ObjectsUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects update unprocessable entity response has a 3xx status code
func (o *ObjectsUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update unprocessable entity response has a 4xx status code
func (o *ObjectsUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects update unprocessable entity response has a 5xx status code
func (o *ObjectsUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this objects update unprocessable entity response a status code equal to that given
func (o *ObjectsUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the objects update unprocessable entity response
func (o *ObjectsUpdateUnprocessableEntity) Code() int {
	return 422
}

func (o *ObjectsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsUpdateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsUpdateInternalServerError creates a ObjectsUpdateInternalServerError with default headers values
func NewObjectsUpdateInternalServerError() *ObjectsUpdateInternalServerError {
	return &ObjectsUpdateInternalServerError{}
}

/*
ObjectsUpdateInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects update internal server error response has a 2xx status code
func (o *ObjectsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects update internal server error response has a 3xx status code
func (o *ObjectsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects update internal server error response has a 4xx status code
func (o *ObjectsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects update internal server error response has a 5xx status code
func (o *ObjectsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this objects update internal server error response a status code equal to that given
func (o *ObjectsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the objects update internal server error response
func (o *ObjectsUpdateInternalServerError) Code() int {
	return 500
}

func (o *ObjectsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /objects/{id}][%d] objectsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
