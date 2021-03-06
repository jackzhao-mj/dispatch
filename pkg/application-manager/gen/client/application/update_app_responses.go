///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateAppReader is a Reader for the UpdateApp structure.
type UpdateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAppOK creates a UpdateAppOK with default headers values
func NewUpdateAppOK() *UpdateAppOK {
	return &UpdateAppOK{}
}

/*UpdateAppOK handles this case with default header values.

Successful update
*/
type UpdateAppOK struct {
	Payload *v1.Application
}

func (o *UpdateAppOK) Error() string {
	return fmt.Sprintf("[PUT /{application}][%d] updateAppOK  %+v", 200, o.Payload)
}

func (o *UpdateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppBadRequest creates a UpdateAppBadRequest with default headers values
func NewUpdateAppBadRequest() *UpdateAppBadRequest {
	return &UpdateAppBadRequest{}
}

/*UpdateAppBadRequest handles this case with default header values.

Invalid input
*/
type UpdateAppBadRequest struct {
	Payload *v1.Error
}

func (o *UpdateAppBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{application}][%d] updateAppBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppNotFound creates a UpdateAppNotFound with default headers values
func NewUpdateAppNotFound() *UpdateAppNotFound {
	return &UpdateAppNotFound{}
}

/*UpdateAppNotFound handles this case with default header values.

Application not found
*/
type UpdateAppNotFound struct {
	Payload *v1.Error
}

func (o *UpdateAppNotFound) Error() string {
	return fmt.Sprintf("[PUT /{application}][%d] updateAppNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInternalServerError creates a UpdateAppInternalServerError with default headers values
func NewUpdateAppInternalServerError() *UpdateAppInternalServerError {
	return &UpdateAppInternalServerError{}
}

/*UpdateAppInternalServerError handles this case with default header values.

Internal error
*/
type UpdateAppInternalServerError struct {
	Payload *v1.Error
}

func (o *UpdateAppInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /{application}][%d] updateAppInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
