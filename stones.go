package stones

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const specVersion = "2.0"

// Validator specfies what methods each object should implement for validation
type Validator interface {
	Validate() (bool error)
}

func invalidType(t string) error {
	return fmt.Errorf(`Type should be "` + t + `"`)
}

func invalidID() error {
	return fmt.Errorf(`ID should be a non-empty, valid StixID`)
}

func invalidSpecVersion() error {
	return fmt.Errorf(`SpecVersion should be` + specVersion)
}

// StixID is a v4 UUID
type StixID struct {
	uuid.UUID
}

// NewStixID returns a STIX ID: a Version 4 UUID
//   128 bit; 16 octets of 32 hexadecimal numbers
//   String representation: 32 bit - 16 bit - 16 bit - 16 bit - 48 bit
//   Example:               6ba7b810-9dad-11d1-80b4-00c04fd430c8
//
// An optional uuid can be passed in and if valid will be converted to StixID type
func NewStixID(u ...string) (StixID, error) {
	if len(u) > 0 && len(u[0]) > 0 {
		id, err := uuid.FromString(u[0])
		return StixID{id}, err
	}

	id, err := uuid.NewV4()
	return StixID{id}, err
}

func (s *StixID) isEmpty() bool {
	empty := &StixID{}
	if s.String() == empty.String() {
		return true
	}
	return false
}
