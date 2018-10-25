package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// VolumeConfig volume config
// swagger:model VolumeConfig
type VolumeConfig struct {

	// flags
	Flags map[string]string `json:"flags,omitempty"`

	// mount point
	MountPoint string `json:"mountPoint,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// read write
	ReadWrite bool `json:"readWrite,omitempty"`
}

// Validate validates this volume config
func (m *VolumeConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
