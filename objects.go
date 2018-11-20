package stones

import "encoding/json"

// Object is a generic STIX object with properties common to all SIIX objects
type Object struct {
	// required
	ID       Identifier `json:"id"`
	Type     string     `json:"type"`
	Created  Timestamp  `json:"created"`
	Modified Timestamp  `json:"modified"`
	// optional
	CreatedByRef       Identifier          `json:"created_by_ref,omitempty"`
	Revoked            bool                `json:"revoked,omitempty"`
	Labels             []string            `json:"labels,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs  []Identifier        `json:"object_marking_refs,omitempty"`
	GranularMarkings   []string            `json:"granular_markings,omitempty"`
	// track the original json
	Source []byte
}

// NewObject returns an object based on JSON input in bytes
func NewObject(b []byte) (o Object, err error) {
	err = json.Unmarshal(b, &o)
	o.Source = b
	return
}
