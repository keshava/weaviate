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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReferenceMeta Additional Meta information about this particular reference.
// swagger:model ReferenceMeta
type ReferenceMeta struct {

	// If a property was set through a classification, this meta field contains additional info
	Classification *ReferenceMetaClassification `json:"classification,omitempty"`
}

// Validate validates this reference meta
func (m *ReferenceMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReferenceMeta) validateClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.Classification) { // not required
		return nil
	}

	if m.Classification != nil {
		if err := m.Classification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReferenceMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferenceMeta) UnmarshalBinary(b []byte) error {
	var res ReferenceMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
