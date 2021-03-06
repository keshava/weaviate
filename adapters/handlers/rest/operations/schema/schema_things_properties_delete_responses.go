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

// SchemaThingsPropertiesDeleteOKCode is the HTTP code returned for type SchemaThingsPropertiesDeleteOK
const SchemaThingsPropertiesDeleteOKCode int = 200

/*SchemaThingsPropertiesDeleteOK Removed the property from the ontology.

swagger:response schemaThingsPropertiesDeleteOK
*/
type SchemaThingsPropertiesDeleteOK struct {
}

// NewSchemaThingsPropertiesDeleteOK creates SchemaThingsPropertiesDeleteOK with default headers values
func NewSchemaThingsPropertiesDeleteOK() *SchemaThingsPropertiesDeleteOK {

	return &SchemaThingsPropertiesDeleteOK{}
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SchemaThingsPropertiesDeleteUnauthorizedCode is the HTTP code returned for type SchemaThingsPropertiesDeleteUnauthorized
const SchemaThingsPropertiesDeleteUnauthorizedCode int = 401

/*SchemaThingsPropertiesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response schemaThingsPropertiesDeleteUnauthorized
*/
type SchemaThingsPropertiesDeleteUnauthorized struct {
}

// NewSchemaThingsPropertiesDeleteUnauthorized creates SchemaThingsPropertiesDeleteUnauthorized with default headers values
func NewSchemaThingsPropertiesDeleteUnauthorized() *SchemaThingsPropertiesDeleteUnauthorized {

	return &SchemaThingsPropertiesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaThingsPropertiesDeleteForbiddenCode is the HTTP code returned for type SchemaThingsPropertiesDeleteForbidden
const SchemaThingsPropertiesDeleteForbiddenCode int = 403

/*SchemaThingsPropertiesDeleteForbidden Forbidden

swagger:response schemaThingsPropertiesDeleteForbidden
*/
type SchemaThingsPropertiesDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesDeleteForbidden creates SchemaThingsPropertiesDeleteForbidden with default headers values
func NewSchemaThingsPropertiesDeleteForbidden() *SchemaThingsPropertiesDeleteForbidden {

	return &SchemaThingsPropertiesDeleteForbidden{}
}

// WithPayload adds the payload to the schema things properties delete forbidden response
func (o *SchemaThingsPropertiesDeleteForbidden) WithPayload(payload *models.ErrorResponse) *SchemaThingsPropertiesDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties delete forbidden response
func (o *SchemaThingsPropertiesDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsPropertiesDeleteInternalServerErrorCode is the HTTP code returned for type SchemaThingsPropertiesDeleteInternalServerError
const SchemaThingsPropertiesDeleteInternalServerErrorCode int = 500

/*SchemaThingsPropertiesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaThingsPropertiesDeleteInternalServerError
*/
type SchemaThingsPropertiesDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesDeleteInternalServerError creates SchemaThingsPropertiesDeleteInternalServerError with default headers values
func NewSchemaThingsPropertiesDeleteInternalServerError() *SchemaThingsPropertiesDeleteInternalServerError {

	return &SchemaThingsPropertiesDeleteInternalServerError{}
}

// WithPayload adds the payload to the schema things properties delete internal server error response
func (o *SchemaThingsPropertiesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaThingsPropertiesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties delete internal server error response
func (o *SchemaThingsPropertiesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
