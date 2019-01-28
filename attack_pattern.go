package stones

import (
	"encoding/json"
	"fmt"
)

// AttackPatternProperites defines the fields only used by this SDO
type AttackPatternProperites struct {
	Name            string           `json:"name" stones:"required"`
	Description     string           `json:"description,omitempty" stones:"optional"`
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty" stones:"optional"`
}

// AttackPattern is a TTP (Tactic, technique, or procedure) that describes how advesaries attempt to compromise targets.
type AttackPattern struct {
	Object
	AttackPatternProperites
}

// NewAttackPattern returns an AttackPattern object
func NewAttackPattern(name string) (ap AttackPattern, err error) {
	ap.ID, err = NewIdentifier(attackPatternType)
	ap.Name = name
	ap.Type = attackPatternType
	ap.Created = NewTimestamp()
	ap.Modified = NewTimestamp()
	return
}

// UnmarshalJSON implements the encoding/json Unmarshaler interface (https://golang.org/pkg/encoding/json/#Unmarshaler).
//
// It will take JSON and deserialize to an Object.  This should not be called directly, but instead
// json.Unmarshal(b []byte, v interface{}) should be used.
func (ap *AttackPattern) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &ap.Object); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &ap.AttackPatternProperites); err != nil {
		return err
	}

	return nil
}

// Valid is called to check for STiX 2.0 specification conformance.
//
// If the AttackPattern is invalid, it returns the list of errors from validation.
func (ap *AttackPattern) Valid() (valid bool, errs []error) {
	ov, err := ap.Object.Valid()
	if !ov {
		errs = append(errs, fmt.Errorf("Invalid Object: %v", err))
	}

	if ap.Type != attackPatternType {
		errs = append(errs, fmt.Errorf("Field 'type' is %s, should be %s", ap.Type, attackPatternType))
	}

	if ap.Name == "" {
		errs = append(errs, fmt.Errorf("Field 'name' required for %s", attackPatternType))
	}

	for _, kcp := range ap.KillChainPhases {
		_, newErrs := kcp.Valid()
		if len(newErrs) != 0 {
			errs = append(errs, newErrs...)
		}
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}

/* helpers */

func validAttackPattern(b []byte) (result bool, errs []error) {
	var ap AttackPattern

	if err := json.Unmarshal(b, &ap); err != nil {
		errs = append(errs, err)
	}

	result, newErrs := ap.Valid()
	errs = append(errs, newErrs...)
	return
}
