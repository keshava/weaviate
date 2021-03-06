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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Schema Definitions of semantic schemas (also see: https://github.com/semi-technologies/weaviate-semantic-schemas).
// swagger:model Schema
type Schema struct {

	// Semantic classes that are available.
	Classes []*Class `json:"classes"`

	// Email of the maintainer.
	// Format: email
	Maintainer strfmt.Email `json:"maintainer,omitempty"`

	// Name of the schema.
	Name string `json:"name,omitempty"`

	// Type of schema, should be "thing" or "action".
	// Enum: [thing action]
	Type string `json:"type,omitempty"`
}

// Validate validates this schema
func (m *Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClasses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema) validateClasses(formats strfmt.Registry) error {

	if swag.IsZero(m.Classes) { // not required
		return nil
	}

	for i := 0; i < len(m.Classes); i++ {
		if swag.IsZero(m.Classes[i]) { // not required
			continue
		}

		if m.Classes[i] != nil {
			if err := m.Classes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Schema) validateMaintainer(formats strfmt.Registry) error {

	if swag.IsZero(m.Maintainer) { // not required
		return nil
	}

	if err := validate.FormatOf("maintainer", "body", "email", m.Maintainer.String(), formats); err != nil {
		return err
	}

	return nil
}

var schemaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["thing","action"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemaTypeTypePropEnum = append(schemaTypeTypePropEnum, v)
	}
}

const (

	// SchemaTypeThing captures enum value "thing"
	SchemaTypeThing string = "thing"

	// SchemaTypeAction captures enum value "action"
	SchemaTypeAction string = "action"
)

// prop value enum
func (m *Schema) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemaTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Schema) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema) UnmarshalBinary(b []byte) error {
	var res Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
