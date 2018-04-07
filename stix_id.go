package stones

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const stixIDJoin = "--"

// StixID represents a STIX type concatenated with a version 4 UUID
type StixID struct {
	Type string
	ID   uuid.UUID
}

// NewStixID takes a STIX type and returns a StixID struct
// The ID field is a v4 UUID
//   v4 UUID:
//     128 bit; 16 octets of 32 hexadecimal numbers
//     String representation: 32 bit - 16 bit - 16 bit - 16 bit - 48 bit
//     Example:               6ba7b810-9dad-11d1-80b4-00c04fd430c8
func NewStixID(t string) (StixID, error) {
	s := StixID{Type: t}

	_, err := s.validType()
	if err != nil {
		return s, err
	}

	s.ID, err = uuid.NewV4()
	return s, err
}

func (s *StixID) String() string {
	return s.Type + stixIDJoin + s.ID.String()
}

func (s *StixID) validate() (bool, error) {
	valid, err := s.validType()
	if err != nil {
		return valid, err
	}

	valid, err = s.validID()
	return valid, err
}

func (s *StixID) validType() (bool, error) {
	if !validStixType(s.Type) {
		return false, fmt.Errorf("Invalid STIX type: %v", s.Type)
	}
	return true, nil
}

func (s *StixID) validID() (bool, error) {
	empty := &uuid.UUID{}
	if s.ID.String() == empty.String() {
		return false, fmt.Errorf("Empty uuid")
	}
	return true, nil
}

/* helpers */

// MarshalStixID takes a raw stix id string and converts it to a StixID type
func MarshalStixID(id string) (StixID, error) {
	var maxParts = 2

	parts := strings.Split(id, stixIDJoin)
	if len(parts) != maxParts {
		return StixID{}, fmt.Errorf("Invalid STIX ID")
	}
	s := StixID{Type: parts[0]}

	var err error
	s.ID, err = uuid.FromString(parts[1])
	return s, err
}

func validStixID(s string) (bool, error) {
	id, err := MarshalStixID(s)
	if err != nil {
		return false, err
	}

	valid, err := id.validate()
	return valid, err
}
