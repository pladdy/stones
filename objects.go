package stones

import (
	"encoding/json"
	"fmt"
)

// Object is a generic STIX object with properties common to all SIIX objects
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
	Source             []byte              `stones:"optional"`
}

// NewObject returns an object based on JSON input in bytes
func NewObject(b []byte) (o Object, err error) {
	err = json.Unmarshal(b, &o)
	o.Source = b
	return
}

// Valid returns whether the objecti s valid or not
func (o *Object) Valid() (valid bool, errs []error) {
	if !validStixType(o.Type) {
		errs = append(errs, fmt.Errorf("Invalid STIX type: %v", o.Type))
	}

	if valid, err := o.ID.Valid(); !valid {
		errs = append(errs, fmt.Errorf("Invalid ID: %v, Error: %v", o.ID, err))
	}

	if o.Type != o.ID.Type {
		errs = append(errs, fmt.Errorf("Object 'Type' (%v) and 'Identfier' Type (%v) don't match", o.Type, o.ID.Type))
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}
