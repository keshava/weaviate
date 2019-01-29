/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaActionsPropertiesAddReader is a Reader for the WeaviateSchemaActionsPropertiesAdd structure.
type WeaviateSchemaActionsPropertiesAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaActionsPropertiesAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaActionsPropertiesAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaActionsPropertiesAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaActionsPropertiesAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaActionsPropertiesAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaActionsPropertiesAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaActionsPropertiesAddOK creates a WeaviateSchemaActionsPropertiesAddOK with default headers values
func NewWeaviateSchemaActionsPropertiesAddOK() *WeaviateSchemaActionsPropertiesAddOK {
	return &WeaviateSchemaActionsPropertiesAddOK{}
}

/*WeaviateSchemaActionsPropertiesAddOK handles this case with default header values.

Added the property.
*/
type WeaviateSchemaActionsPropertiesAddOK struct {
}

func (o *WeaviateSchemaActionsPropertiesAddOK) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] weaviateSchemaActionsPropertiesAddOK ", 200)
}

func (o *WeaviateSchemaActionsPropertiesAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesAddUnauthorized creates a WeaviateSchemaActionsPropertiesAddUnauthorized with default headers values
func NewWeaviateSchemaActionsPropertiesAddUnauthorized() *WeaviateSchemaActionsPropertiesAddUnauthorized {
	return &WeaviateSchemaActionsPropertiesAddUnauthorized{}
}

/*WeaviateSchemaActionsPropertiesAddUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaActionsPropertiesAddUnauthorized struct {
}

func (o *WeaviateSchemaActionsPropertiesAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] weaviateSchemaActionsPropertiesAddUnauthorized ", 401)
}

func (o *WeaviateSchemaActionsPropertiesAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesAddForbidden creates a WeaviateSchemaActionsPropertiesAddForbidden with default headers values
func NewWeaviateSchemaActionsPropertiesAddForbidden() *WeaviateSchemaActionsPropertiesAddForbidden {
	return &WeaviateSchemaActionsPropertiesAddForbidden{}
}

/*WeaviateSchemaActionsPropertiesAddForbidden handles this case with default header values.

Could not find the Action class.
*/
type WeaviateSchemaActionsPropertiesAddForbidden struct {
}

func (o *WeaviateSchemaActionsPropertiesAddForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] weaviateSchemaActionsPropertiesAddForbidden ", 403)
}

func (o *WeaviateSchemaActionsPropertiesAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesAddUnprocessableEntity creates a WeaviateSchemaActionsPropertiesAddUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsPropertiesAddUnprocessableEntity() *WeaviateSchemaActionsPropertiesAddUnprocessableEntity {
	return &WeaviateSchemaActionsPropertiesAddUnprocessableEntity{}
}

/*WeaviateSchemaActionsPropertiesAddUnprocessableEntity handles this case with default header values.

Invalid property.
*/
type WeaviateSchemaActionsPropertiesAddUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] weaviateSchemaActionsPropertiesAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsPropertiesAddInternalServerError creates a WeaviateSchemaActionsPropertiesAddInternalServerError with default headers values
func NewWeaviateSchemaActionsPropertiesAddInternalServerError() *WeaviateSchemaActionsPropertiesAddInternalServerError {
	return &WeaviateSchemaActionsPropertiesAddInternalServerError{}
}

/*WeaviateSchemaActionsPropertiesAddInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaActionsPropertiesAddInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] weaviateSchemaActionsPropertiesAddInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
