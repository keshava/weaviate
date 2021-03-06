//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsPropertiesAddOKCode is the HTTP code returned for type SchemaActionsPropertiesAddOK
const SchemaActionsPropertiesAddOKCode int = 200

/*SchemaActionsPropertiesAddOK Added the property.

swagger:response schemaActionsPropertiesAddOK
*/
type SchemaActionsPropertiesAddOK struct {

	/*
	  In: Body
	*/
	Payload *models.Property `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesAddOK creates SchemaActionsPropertiesAddOK with default headers values
func NewSchemaActionsPropertiesAddOK() *SchemaActionsPropertiesAddOK {

	return &SchemaActionsPropertiesAddOK{}
}

// WithPayload adds the payload to the schema actions properties add o k response
func (o *SchemaActionsPropertiesAddOK) WithPayload(payload *models.Property) *SchemaActionsPropertiesAddOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties add o k response
func (o *SchemaActionsPropertiesAddOK) SetPayload(payload *models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsPropertiesAddUnauthorizedCode is the HTTP code returned for type SchemaActionsPropertiesAddUnauthorized
const SchemaActionsPropertiesAddUnauthorizedCode int = 401

/*SchemaActionsPropertiesAddUnauthorized Unauthorized or invalid credentials.

swagger:response schemaActionsPropertiesAddUnauthorized
*/
type SchemaActionsPropertiesAddUnauthorized struct {
}

// NewSchemaActionsPropertiesAddUnauthorized creates SchemaActionsPropertiesAddUnauthorized with default headers values
func NewSchemaActionsPropertiesAddUnauthorized() *SchemaActionsPropertiesAddUnauthorized {

	return &SchemaActionsPropertiesAddUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesAddUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaActionsPropertiesAddForbiddenCode is the HTTP code returned for type SchemaActionsPropertiesAddForbidden
const SchemaActionsPropertiesAddForbiddenCode int = 403

/*SchemaActionsPropertiesAddForbidden Forbidden

swagger:response schemaActionsPropertiesAddForbidden
*/
type SchemaActionsPropertiesAddForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesAddForbidden creates SchemaActionsPropertiesAddForbidden with default headers values
func NewSchemaActionsPropertiesAddForbidden() *SchemaActionsPropertiesAddForbidden {

	return &SchemaActionsPropertiesAddForbidden{}
}

// WithPayload adds the payload to the schema actions properties add forbidden response
func (o *SchemaActionsPropertiesAddForbidden) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesAddForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties add forbidden response
func (o *SchemaActionsPropertiesAddForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesAddForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsPropertiesAddUnprocessableEntityCode is the HTTP code returned for type SchemaActionsPropertiesAddUnprocessableEntity
const SchemaActionsPropertiesAddUnprocessableEntityCode int = 422

/*SchemaActionsPropertiesAddUnprocessableEntity Invalid property.

swagger:response schemaActionsPropertiesAddUnprocessableEntity
*/
type SchemaActionsPropertiesAddUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesAddUnprocessableEntity creates SchemaActionsPropertiesAddUnprocessableEntity with default headers values
func NewSchemaActionsPropertiesAddUnprocessableEntity() *SchemaActionsPropertiesAddUnprocessableEntity {

	return &SchemaActionsPropertiesAddUnprocessableEntity{}
}

// WithPayload adds the payload to the schema actions properties add unprocessable entity response
func (o *SchemaActionsPropertiesAddUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesAddUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties add unprocessable entity response
func (o *SchemaActionsPropertiesAddUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesAddUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsPropertiesAddInternalServerErrorCode is the HTTP code returned for type SchemaActionsPropertiesAddInternalServerError
const SchemaActionsPropertiesAddInternalServerErrorCode int = 500

/*SchemaActionsPropertiesAddInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaActionsPropertiesAddInternalServerError
*/
type SchemaActionsPropertiesAddInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsPropertiesAddInternalServerError creates SchemaActionsPropertiesAddInternalServerError with default headers values
func NewSchemaActionsPropertiesAddInternalServerError() *SchemaActionsPropertiesAddInternalServerError {

	return &SchemaActionsPropertiesAddInternalServerError{}
}

// WithPayload adds the payload to the schema actions properties add internal server error response
func (o *SchemaActionsPropertiesAddInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaActionsPropertiesAddInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions properties add internal server error response
func (o *SchemaActionsPropertiesAddInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsPropertiesAddInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
