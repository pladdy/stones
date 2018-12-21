package stones

import "fmt"

// AttackPattern is a TTP (Tactic, technique, or procedure) that describes how advesaries attempt to compromise targets.
type AttackPattern struct {
	Object          `stones:"required"`
	Name            string           `stones:"required"`
	Description     string           `json:"omitempty" stones:"optional"`
	KillChainPhases []KillChainPhase `json:"omitempty" stones:"optional"`
}

// NewAttackPattern returns an AttackPattern object
func NewAttackPattern(name string) (ap AttackPattern, err error) {
	ap.ID, err = NewIdentifier(attackPatternType)
	ap.Name = name
	ap.Type = attackPatternType
	return
}

// Valid is called to check for STiX 2.0 specification conformance.
//
// If the AttackPattern is invalid, it returns the list of errors from validation.
func (ap *AttackPattern) Valid() (valid bool, errs []error) {
	_, newErrs := ap.Object.Valid()
	if len(newErrs) != 0 {
		errs = append(errs, newErrs...)
	}

	if ap.Type != attackPatternType {
		errs = append(errs, invalidType())
	}

	_, newErrs = ap.ID.Valid()
	if len(newErrs) != 0 {
		errs = append(errs, newErrs...)
	}

	if ap.Name == "" {
		errs = append(errs, fmt.Errorf("Property 'Name' required for %s", attackPatternType))
	}

	for _, kcp := range ap.KillChainPhases {
		_, newErrs = kcp.Valid()
		if len(newErrs) != 0 {
			errs = append(errs, newErrs...)
		}
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}
