// Package stones is for STIX object validation
package stones

import (
	"fmt"
)

/* STIX types */

const (
	bundleType = "bundle"
)

var validStixTypes = map[string]bool{bundleType: true}

func validStixType(s string) bool {
	return validStixTypes[s]
}

/* Validator */

const specVersion = "2.0"

// Validator specfies what methods each object should implement for validation
type Validator interface {
	Validate() (bool error)
}

func invalidType() error {
	return fmt.Errorf(`STIX Type is invalid`)
}

func invalidID(err error) error {
	return fmt.Errorf(`STIX ID should have a valid STIX type and valid v4 UUID: %v`, err)
}

func invalidSpecVersion() error {
	return fmt.Errorf(`SpecVersion should be` + specVersion)
}
