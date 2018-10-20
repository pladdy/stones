package stones

// Object is a generic STIX object
type Object struct {
	// required
	ID       Identifier `json:"id"`
	Type     string     `json:"type"`
	Created  string     `json:"created"`
	Modified string     `json:"modified"`
	// optional
	CratedByRef        Identifier          `json:"created_by_ref,omitempty"`
	Revoked            bool                `json:"revoked,omitempty"`
	Labels             []string            `json:"labels,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs  []Identifier        `json:"object_marking_refs,omitempty"`
	GranularMarkings   []string            `json:"granular_markings,omitempty"`
}
