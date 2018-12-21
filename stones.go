// Package stones is for STIX object validation
package stones

import (
	"errors"
	"fmt"
)

const specVersion = "2.0"

// DomainObject represents a SDO (STIX Domain Object)
type DomainObject int

const (
	// AttackPattern is a TTP that describes ways advesaries attempt to compromise targets
	_AttackPattern DomainObject = 1 + iota
	// Campaign is a grouping of behaviors describing malicious activites or attacks
	_Campaign
	// CourseOfAction is an action taken to prevent or respond to an attack
	_CourseOfAction
	// Identity represents individuals, groups, or organizations
	_Identity
	// Indicator contains a pattern that can be used to detect suspicious or malicious activity
	_Indicator
	// IntrusionSet is a grouped set of advesarial behaviors and resources with common behavior
	_IntrusionSet
	// Malware is a type of TTP also known as malcious code or software
	_Malware
	// ObservedData conveys information observed on networks or systems using the Cyber Observable specification
	_ObservedData
	// Report are collections of threat intelligence focused on one or more topics
	_Report
	// ThreatActor are actual individuals, groups, organizations, believed to operate with malicious intent
	_ThreatActor
	// Tool is a piece of software that can be used by threat actors to conduct attacks
	_Tool
	// Vulnerability is a "a mistake in software that can be directly used by a hacker to gain access to a system or network"
	// http://docs.oasis-open.org/cti/stix/v2.0/cs01/part2-stix-objects/stix-v2.0-cs01-part2-stix-objects.html#omgb5053jgfy
	_Vulnerability
)

// RelationshipObject represents a SRO (STIX Relationship Object)
type RelationshipObject int

const (
	// Relationship is used to link two SDOs in order to descibre how they're related
	_Relationship RelationshipObject = 101 + iota
	// Sighting denotes that something was believed to be seen; they track who and what are being targeted
	_Sighting
)

// TransportObject represents a STIX Transport (currently a bundle)
type TransportObject int

const (
	// Bundle is a collection of arbitrary STIX objects grouped in a container
	_Bundle TransportObject = 201 + iota
)

// DomainObjects is a map of domain object to it's name
var DomainObjects = map[DomainObject]string{
	_AttackPattern:  "attack-pattern",
	_Campaign:       "campaign",
	_CourseOfAction: "course-of-action",
	_Identity:       "identity",
	_Indicator:      "indicator",
	_IntrusionSet:   "intrusion-set",
	_Malware:        "malware",
	_ObservedData:   "observed-data",
	_Report:         "report",
	_ThreatActor:    "threat-actor",
	_Tool:           "tool",
	_Vulnerability:  "vulnerability",
}

// RelationshipObjects is a map of relationship object to it's name
var RelationshipObjects = map[RelationshipObject]string{
	_Relationship: "relationship",
	_Sighting:     "sighting",
}

// TransportObjects is a map of transport objects to it's name
var TransportObjects = map[TransportObject]string{
	_Bundle: "bundle",
}

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

// ExternalReference data type for pointing to references that are external to STIX
type ExternalReference struct {
	SourceName  string   `json:"source_name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Hashes      []string `json:"hashes"`
	ExternalID  string   `json:"external_id"`
}

// Valid returns whether an ExternalReference is valid
// func (er *ExternalReference) Valid() (bool, []error) {
// 	// SourceName is required
// 	// at least one of description, url, or external id are required
// 	return true, []error{}
// }

// Hashes represents 1 or more cryptographic hashes as key/value pairs
// hasing algorithm -> hash string
type Hashes map[string]string

// Valid returns whether hash definitions are valid
// func (h *Hashes) Valid() (bool, []error) {
// 	// keys are 30 ascii chars max
// 	// values must be no longer than 256 chars
// 	// values are characters a-z (lowercase ASCII), A-Z (uppercase ASCII), numerals 0-9, hyphen (-), and underscore (_)
// 	return true, []error{}
// }

// KillChainPhase represents a phase in a kill chain
type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
}

// Valid will run validation on a KillChainPhase
func (k *KillChainPhase) Valid() (valid bool, errs []error) {
	if k.KillChainName == "" {
		errs = append(errs, errors.New("KillChainName can't be empty"))
	}

	if k.PhaseName == "" {
		errs = append(errs, errors.New("PhaseName can't be empty"))
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}

// Validator specfies what methods each object should implement for validation
type Validator interface {
	Valid() (bool, []error)
}

/* helpers */

func invalidType() error {
	return fmt.Errorf(`STIX Type is invalid`)
}

func invalidSpecVersion() error {
	return fmt.Errorf(`SpecVersion should be` + specVersion)
}

func validationErrors(errs []error) error {
	return fmt.Errorf(fmt.Sprintf("Errors: %v", errs))
}
