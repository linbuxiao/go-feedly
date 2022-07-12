// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Collection collection
//
// swagger:model Collection
type Collection struct {

	// cover
	Cover string `json:"cover,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// feeds
	Feeds []*Feed `json:"feeds"`

	// id
	// Required: true
	ID *string `json:"id"`

	// label
	Label string `json:"label,omitempty"`
}

// Validate validates this collection
func (m *Collection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeeds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Collection) validateFeeds(formats strfmt.Registry) error {
	if swag.IsZero(m.Feeds) { // not required
		return nil
	}

	for i := 0; i < len(m.Feeds); i++ {
		if swag.IsZero(m.Feeds[i]) { // not required
			continue
		}

		if m.Feeds[i] != nil {
			if err := m.Feeds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feeds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feeds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Collection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this collection based on the context it is used
func (m *Collection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeeds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Collection) contextValidateFeeds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Feeds); i++ {

		if m.Feeds[i] != nil {
			if err := m.Feeds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feeds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feeds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Collection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Collection) UnmarshalBinary(b []byte) error {
	var res Collection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
