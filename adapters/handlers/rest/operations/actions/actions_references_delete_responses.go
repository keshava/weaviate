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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsReferencesDeleteNoContentCode is the HTTP code returned for type ActionsReferencesDeleteNoContent
const ActionsReferencesDeleteNoContentCode int = 204

/*ActionsReferencesDeleteNoContent Successfully deleted.

swagger:response actionsReferencesDeleteNoContent
*/
type ActionsReferencesDeleteNoContent struct {
}

// NewActionsReferencesDeleteNoContent creates ActionsReferencesDeleteNoContent with default headers values
func NewActionsReferencesDeleteNoContent() *ActionsReferencesDeleteNoContent {

	return &ActionsReferencesDeleteNoContent{}
}

// WriteResponse to the client
func (o *ActionsReferencesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ActionsReferencesDeleteUnauthorizedCode is the HTTP code returned for type ActionsReferencesDeleteUnauthorized
const ActionsReferencesDeleteUnauthorizedCode int = 401

/*ActionsReferencesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response actionsReferencesDeleteUnauthorized
*/
type ActionsReferencesDeleteUnauthorized struct {
}

// NewActionsReferencesDeleteUnauthorized creates ActionsReferencesDeleteUnauthorized with default headers values
func NewActionsReferencesDeleteUnauthorized() *ActionsReferencesDeleteUnauthorized {

	return &ActionsReferencesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *ActionsReferencesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActionsReferencesDeleteForbiddenCode is the HTTP code returned for type ActionsReferencesDeleteForbidden
const ActionsReferencesDeleteForbiddenCode int = 403

/*ActionsReferencesDeleteForbidden Forbidden

swagger:response actionsReferencesDeleteForbidden
*/
type ActionsReferencesDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsReferencesDeleteForbidden creates ActionsReferencesDeleteForbidden with default headers values
func NewActionsReferencesDeleteForbidden() *ActionsReferencesDeleteForbidden {

	return &ActionsReferencesDeleteForbidden{}
}

// WithPayload adds the payload to the actions references delete forbidden response
func (o *ActionsReferencesDeleteForbidden) WithPayload(payload *models.ErrorResponse) *ActionsReferencesDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions references delete forbidden response
func (o *ActionsReferencesDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsReferencesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsReferencesDeleteNotFoundCode is the HTTP code returned for type ActionsReferencesDeleteNotFound
const ActionsReferencesDeleteNotFoundCode int = 404

/*ActionsReferencesDeleteNotFound Successful query result but no resource was found.

swagger:response actionsReferencesDeleteNotFound
*/
type ActionsReferencesDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsReferencesDeleteNotFound creates ActionsReferencesDeleteNotFound with default headers values
func NewActionsReferencesDeleteNotFound() *ActionsReferencesDeleteNotFound {

	return &ActionsReferencesDeleteNotFound{}
}

// WithPayload adds the payload to the actions references delete not found response
func (o *ActionsReferencesDeleteNotFound) WithPayload(payload *models.ErrorResponse) *ActionsReferencesDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions references delete not found response
func (o *ActionsReferencesDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsReferencesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsReferencesDeleteInternalServerErrorCode is the HTTP code returned for type ActionsReferencesDeleteInternalServerError
const ActionsReferencesDeleteInternalServerErrorCode int = 500

/*ActionsReferencesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response actionsReferencesDeleteInternalServerError
*/
type ActionsReferencesDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsReferencesDeleteInternalServerError creates ActionsReferencesDeleteInternalServerError with default headers values
func NewActionsReferencesDeleteInternalServerError() *ActionsReferencesDeleteInternalServerError {

	return &ActionsReferencesDeleteInternalServerError{}
}

// WithPayload adds the payload to the actions references delete internal server error response
func (o *ActionsReferencesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ActionsReferencesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions references delete internal server error response
func (o *ActionsReferencesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsReferencesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
