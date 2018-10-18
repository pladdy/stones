// Package stones is for STIX object validation
package stones

import (
	"fmt"
)

/* STIX types */

// list of types
const (
	attackPatternType  = "attack-pattern"
	bundleType         = "bundle"
	campaignType       = "campaign"
	courseOfActionType = "course-of-action"
	identityType       = "identity"
	indicatorType      = "indicator"
	intrusionSetType   = "intrusion-set"
	malwareType        = "malware"
	observedDataType   = "observed-data"
	relationshipType   = "relationship"
	reportType         = "report"
	sightingType       = "sighting"
	threatActorType    = "threat-actor"
	toolType           = "tool"
	vulnerabilityType  = "vulnerability"
)

// map of types used to validate
var validStixTypes = map[string]bool{
	attackPatternType:  true,
	bundleType:         true,
	campaignType:       true,
	courseOfActionType: true,
	identityType:       true,
	indicatorType:      true,
	intrusionSetType:   true,
	malwareType:        true,
	observedDataType:   true,
	relationshipType:   true,
	reportType:         true,
	sightingType:       true,
	threatActorType:    true,
	toolType:           true,
	vulnerabilityType:  true}

func validStixType(s string) bool {
	return validStixTypes[s]
}

const specVersion = "2.0"

// Validator specfies what methods each object should implement for validation
type Validator interface {
	Valid() (bool, []error)
}

/* helpers */

func invalidType() error {
	return fmt.Errorf(`STIX Type is invalid`)
}

func invalidID(err error) error {
	return fmt.Errorf(`STIX ID should have a valid STIX type and valid v4 UUID: %v`, err)
}

func invalidSpecVersion() error {
	return fmt.Errorf(`SpecVersion should be` + specVersion)
}
