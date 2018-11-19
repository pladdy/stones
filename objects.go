package stones

import "fmt"

// Object is a generic STIX object with properties common to all SIIX objects
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
	// track the original json object
	Source []byte
}

// AttackPattern is a TTP (Tactic, technique, or procedure) that describes how advesaries attempt to compromise targets
type AttackPattern struct {
	Object
	Name            string
	Description     string
	KillChainPhases []KillChainPhase
}

// NewAttackPattern returns an AttackPattern object
func NewAttackPattern(name string) (ap AttackPattern, err error) {
	ap.ID, err = NewIdentifier(attackPatternType)
	ap.Name = name
	ap.Type = attackPatternType
	return
}

// Valid checks if object is valid STIX
func (ap *AttackPattern) Valid() (valid bool, errs []error) {
	if ap.Type != attackPatternType {
		errs = append(errs, invalidType())
	}

	_, es := ap.ID.Valid()
	if len(es) != 0 {
		errs = append(errs, es...)
	}

	if ap.Name == "" {
		errs = append(errs, fmt.Errorf("Property 'Name' required for %s", attackPatternType))
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}
