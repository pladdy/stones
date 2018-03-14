package stones

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const stixIDJoin = "--"

type stixID struct {
	Type string
	ID   uuid.UUID
}

// newStixID contains a stix Type and an ID: a Version 4 UUID
//   128 bit; 16 octets of 32 hexadecimal numbers
//   String representation: 32 bit - 16 bit - 16 bit - 16 bit - 48 bit
//   Example:               6ba7b810-9dad-11d1-80b4-00c04fd430c8
//
// An optional uuid can be passed in and if valid will be converted to StixID type
func newStixID(t string, u ...string) (stixID, error) {
	s := stixID{Type: t}

	_, err := s.validType()
	if err != nil {
		return s, err
	}

	if len(u) > 0 && len(u[0]) > 0 {
		s.ID, err = uuid.FromString(u[0])
	} else {
		s.ID, err = uuid.NewV4()
	}
	return s, err
}

func (s *stixID) String() string {
	return s.Type + stixIDJoin + s.ID.String()
}

func (s *stixID) validate() (bool, error) {
	valid, err := s.validType()
	if err != nil {
		return valid, err
	}

	valid, err = s.validID()
	return valid, err
}

func (s *stixID) validType() (bool, error) {
	if !validStixType(s.Type) {
		return false, fmt.Errorf("Invalid type")
	}
	return true, nil
}

func (s *stixID) validID() (bool, error) {
	empty := &uuid.UUID{}
	if s.ID.String() == empty.String() {
		return false, fmt.Errorf("Empty uuid")
	}
	return true, nil
}

func validStixID(s string) (bool, error) {
	var err error
	var maxParts = 2

	parts := strings.Split(s, stixIDJoin)
	if len(parts) != maxParts {
		return false, fmt.Errorf("Invalid id")
	}

	id, err := newStixID(parts[0], parts[1])
	if err != nil {
		return false, err
	}

	valid, err := id.validate()
	return valid, err
}
