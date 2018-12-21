package stones

import (
	"encoding/json"
	"fmt"
	"time"
)

// Object is a generic STIX object with properties common to all SIIX objects.
type Object struct {
	ID                 Identifier          `json:"id" stones:"required"`
	Type               string              `json:"type" stones:"required"`
	Created            Timestamp           `json:"created" stones:"required"`
	Modified           Timestamp           `json:"modified" stones:"required"`
	CreatedByRef       Identifier          `json:"created_by_ref,omitempty" stones:"optional"`
	Revoked            bool                `json:"revoked,omitempty" stones:"optional"`
	Labels             []string            `json:"labels,omitempty" stones:"optional"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty" stones:"optional"`
	ObjectMarkingRefs  []Identifier        `json:"object_marking_refs,omitempty" stones:"optional"`
	GranularMarkings   []string            `json:"granular_markings,omitempty" stones:"optional"`
	Source             []byte              `stones:"required"`
}

// NewObject takes a STIX Type as a string and returns an Object with that Type and a new ID.
func NewObject(t string) (o Object, err error) {
	o.ID, err = NewIdentifier(t)
	o.Type = t
	return
}

// UnmarshalJSON implements the encoding/json Unmarshaler interface (https://golang.org/pkg/encoding/json/#Unmarshaler).
//
// It will take JSON and deserialize to an Object.  This should not be called directly, but instead
// json.Unmarshal(b []byte, v interface{}) should be used.
//
// Validation is run on the Object; if invalid, validation errors are returned as one error.
func (o *Object) UnmarshalJSON(d []byte) error {
	// use an aliase to avoid infinite loop by using Unmarshal on the object
	type ObjectAlias Object
	alias := struct{ *ObjectAlias }{
		ObjectAlias: (*ObjectAlias)(o),
	}

	if err := json.Unmarshal(d, &alias); err != nil {
		return err
	}
	o.Source = d

	if valid, errs := o.Valid(); !valid {
		return validationErrors(errs)
	}

	return nil
}

// Valid is called to check for STiX 2.0 specification conformance.
//
// If the Object is invalid, it returns the list of errors from validation.
func (o *Object) Valid() (valid bool, errs []error) {
	if !validStixType(o.Type) {
		errs = append(errs, fmt.Errorf("Invalid 'Type': %v", o.Type))
	}

	if valid, err := o.ID.Valid(); !valid {
		errs = append(errs, fmt.Errorf("Invalid 'Identifier': %v, Error: %v", o.ID, err))
	}

	if o.Type != o.ID.Type {
		errs = append(errs, fmt.Errorf("Object 'Type' (%v) and 'Identfier' Type (%v) must match", o.Type, o.ID.Type))
	}

	notSet := Timestamp{time.Time{}}

	if o.Created == notSet {
		errs = append(errs, fmt.Errorf("Created time must be set; should be > 1970-01-01 00:00:00Z"))
	}

	if o.Modified == notSet {
		errs = append(errs, fmt.Errorf("Modified time must be set; should be > 1970-01-01 00:00:00Z"))
	}

	if len(o.Source) == 0 {
		errs = append(errs, fmt.Errorf("Object 'Source' cannot be empty"))
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}
