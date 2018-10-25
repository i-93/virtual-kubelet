package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BindContainerResponse bind container response
// swagger:model BindContainerResponse
type BindContainerResponse struct {

	// endpoints
	// Required: true
	Endpoints []*EndpointConfig `json:"endpoints"`

	// handle
	// Required: true
	Handle string `json:"handle"`
}

// Validate validates this bind container response
func (m *BindContainerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BindContainerResponse) validateEndpoints(formats strfmt.Registry) error {

	if err := validate.Required("endpoints", "body", m.Endpoints); err != nil {
		return err
	}

	for i := 0; i < len(m.Endpoints); i++ {

		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {

			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BindContainerResponse) validateHandle(formats strfmt.Registry) error {

	if err := validate.RequiredString("handle", "body", string(m.Handle)); err != nil {
		return err
	}

	return nil
}
