package stones

import (
	"encoding/json"
	"fmt"
)

// Bundle is a STIX bundle: a collection of STIX objects
type Bundle struct {
	Type        string            `json:"type"`
	ID          string            `json:"id"`
	SpecVersion string            `json:"spec_version"`
	Objects     []json.RawMessage `json:"objects"`
}

// NewBundle returns a STIX bundle object
func NewBundle() (Bundle, error) {
	b := Bundle{}
	id, err := NewIdentifier(bundleType)

	b.ID = id.String()
	b.Type = bundleType
	b.SpecVersion = specVersion

	return b, err
}

// AddObject adds a object to the bundle
func (b *Bundle) AddObject(o string) {
	bundle := *b
	bundle.Objects = append(bundle.Objects, json.RawMessage(o))
	*b = bundle
}

// Valid is called to validate a bundle
func (b *Bundle) Valid() (valid bool, errs []error) {
	if b.Type != bundleType {
		errs = append(errs, invalidType())
	}

	if _, err := UnmarshalIdentifier(b.ID); err != nil {
		errs = append(errs, invalidID(err))
	}

	if b.SpecVersion != specVersion {
		errs = append(errs, invalidSpecVersion())
	}

	if len(b.Objects) == 0 {
		errs = append(errs, fmt.Errorf(`Objects must be more than 0`))
	}

	if len(errs) == 0 {
		valid = true
	}
	return valid, errs
}
