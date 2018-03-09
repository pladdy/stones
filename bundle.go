package stones

import "fmt"

const bundleType = "bundle"

// Bundle is a STIX bundle: a collection of STIX objects
type Bundle struct {
	Type        string   `json:"type"`
	ID          StixID   `json:"id"`
	SpecVersion string   `json:"spec_version"`
	Objects     []string `json:"objects"`
}

// NewBundle returns a STIX bundle object
func NewBundle() (Bundle, error) {
	b := Bundle{}
	id, err := NewStixID()

	b.ID = id
	b.Type = bundleType
	b.SpecVersion = specVersion

	return b, err
}

// Validate is called to validate a bundle
func (b *Bundle) Validate() (bool, []error) {
	var errs []error

	if b.Type != bundleType {
		errs = append(errs, invalidType(b.Type))
	}

	if b.ID.isEmpty() {
		errs = append(errs, invalidID())
	}

	if b.SpecVersion != specVersion {
		errs = append(errs, invalidSpecVersion())
	}

	if len(b.Objects) == 0 {
		errs = append(errs, fmt.Errorf(`Objects must be more than 0`))
	}

	if len(errs) == 0 {
		return true, errs
	}
	return false, errs
}
