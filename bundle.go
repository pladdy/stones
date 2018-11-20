package stones

import (
	"encoding/json"
	"fmt"
)

// Bundle is a STIX bundle: a collection of STIX objects
type Bundle struct {
	// required
	Type        string     `json:"type"`
	ID          Identifier `json:"id"`
	SpecVersion string     `json:"spec_version"`
	// optional
	Objects []json.RawMessage `json:"objects"`
}

// NewBundle returns a STIX bundle object
func NewBundle() (b Bundle, err error) {
	b.ID, err = NewIdentifier(bundleType)
	b.Type = bundleType
	b.SpecVersion = specVersion
	return b, err
}

// AddObject adds a object to the bundle
func (b *Bundle) AddObject(o string) {
	// skip empty objects
	if len(o) == 0 {
		return
	}

	bundle := *b
	bundle.Objects = append(bundle.Objects, json.RawMessage(o))
	*b = bundle
}

// Valid is called to validate a bundle
func (b *Bundle) Valid() (valid bool, errs []error) {
	if b.Type != bundleType {
		errs = append(errs, invalidType())
	}

	_, es := b.ID.Valid()
	if len(es) != 0 {
		errs = append(errs, es...)
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
