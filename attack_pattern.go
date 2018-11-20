package stones

import "fmt"

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
