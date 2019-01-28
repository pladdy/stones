// Package stones is for STIX object validation
package stones

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// TODO: move these doc strings to the structs when defined
// Campaign is a grouping of behaviors describing malicious activites or attacks
// CourseOfAction is an action taken to prevent or respond to an attack
// Identity represents individuals, groups, or organizations
// Indicator contains a pattern that can be used to detect suspicious or malicious activity
// IntrusionSet is a grouped set of advesarial behaviors and resources with common behavior
// Malware is a type of TTP also known as malcious code or software
// ObservedData conveys information observed on networks or systems using the Cyber Observable specification
// Report are collections of threat intelligence focused on one or more topics
// ThreatActor are actual individuals, groups, organizations, believed to operate with malicious intent
// Tool is a piece of software that can be used by threat actors to conduct attacks
// Vulnerability is a "a mistake in software that can be directly used by a hacker to gain access to a system or network"
// http://docs.oasis-open.org/cti/stix/v2.0/cs01/part2-stix-objects/stix-v2.0-cs01-part2-stix-objects.html#omgb5053jgfy

// Relationship is used to link two SDOs in order to descibre how they're related
// Sighting denotes that something was believed to be seen; they track who and what are being targeted

// list of types
const (
	// SDO, SRO, and STO types
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

	// object types
	domainObject       = "domain object"
	relationshipObject = "relationship object"
	transportObject    = "transoport object"

	// version
	specVersion = "2.0"
)

// Validatable defines the interface objects need to satisfy to be able to validate them selves
type Validatable interface {
	Valid() (bool, []error)
}

// ErrorsToString converts a slice of errors to one error
// TODO: this is gross, how about a ValidationErrors type with a method to pretty print errors?
func ErrorsToString(errs []error) error {
	s := []string{}
	for _, err := range errs {
		s = append(s, fmt.Sprintf("%v", err))
	}
	return fmt.Errorf(strings.Join(s, "; "))
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

func invalidSpecVersion() error {
	return fmt.Errorf(`SpecVersion should be` + specVersion)
}

func invalidType() error {
	return fmt.Errorf(`STIX Type is invalid`)
}

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

// peek into a JSON object and return its type
func objectType(o []byte) string {
	re := regexp.MustCompile(`"type":\s"(.+?)"`)
	matches := re.FindStringSubmatch(string(o))

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

type objectValidator func(b []byte) (bool, []error)

var objectValidators = map[string]objectValidator{
	attackPatternType: validAttackPattern,
}

// Validator specfies what methods each object should implement for validation
type Validator interface {
	Valid() (bool, []error)
}

// Validate will take a raw JSON object and run validation against it.
func Validate(b []byte) (bool, []error) {
	t := objectType(b)
	if validStixType(t) && objectValidators[t] != nil {
		return objectValidators[t](b)
	}
	return false, []error{fmt.Errorf("Invalid STIX type '%v' for object", t)}
}

func validStixType(s string) bool {
	return validStixTypes[s]
}

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
