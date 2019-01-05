package stones

import (
	"encoding/json"
	"fmt"
)

// Bundle is a collection of arbitrary STIX objects grouped in a container
type Bundle struct {
	// required
	Type        string            `json:"type" stones:"required"`
	ID          Identifier        `json:"id" stones:"required"`
	SpecVersion string            `json:"spec_version" stones:"required"`
	Objects     []json.RawMessage `json:"objects" stones:"optional"`
}

// NewBundle returns an empty STIX Bundle object (no objects).
//
// It creates a new UUIDv4 and sets the Type, ID, and SpecVersion properties automatically.
func NewBundle() (b Bundle, err error) {
	b.ID, err = NewIdentifier(bundleType)
	b.Type = bundleType
	b.SpecVersion = specVersion
	return b, err
}

// AddObject adds an object to the Bundle.  It expects a JSON string that will be appended as a json.RawMessage.
func (b *Bundle) AddObject(o string) {
	// skip empty objects
	if len(o) == 0 {
		return
	}

	bundle := *b
	bundle.Objects = append(bundle.Objects, json.RawMessage(o))
	*b = bundle
}

// Valid is called to check for STiX 2.0 specification conformance.
//
// If the Bundle is invalid, it returns the list of errors from validation.
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
