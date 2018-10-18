package stones

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

const identifierJoin = "--"

// ID is a string representation of an Identifier (example malware--6ba7b810-9dad-11d1-80b4-00c04fd430c8)
// it exists to more easily unmarshal a database value into an ID (vs unmarshalling a string into an Identifier
// struct).
type ID string

// Valid returns true if the ID is a valid Identifier
func (id *ID) Valid() (bool, error) {
	_, err := UnmarshalIdentifier(fmt.Sprint(*id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Identifier represents a STIX type concatenated with a version 4 UUID
type Identifier struct {
	Type string
	ID   uuid.UUID
}

// NewIdentifier takes a STIX type and returns a Identifier struct
// The ID field is a v4 UUID
//   v4 UUID:
//     128 bit; 16 octets of 32 hexadecimal numbers
//     String representation: 32 bit - 16 bit - 16 bit - 16 bit - 48 bit
//     Example:               6ba7b810-9dad-11d1-80b4-00c04fd430c8
func NewIdentifier(t string) (id Identifier, err error) {
	id = Identifier{Type: t}
	id.ID, err = uuid.NewV4()
	return id, err
}

func (id *Identifier) String() string {
	return strings.Join([]string{id.Type, id.ID.String()}, identifierJoin)
}

// Valid will validate the Identifier
func (id *Identifier) Valid() (result bool, errs []error) {
	result = true

	if !validStixType(id.Type) {
		errs = append(errs, fmt.Errorf("Invalid STIX type: %v", id.Type))
		result = false
	}

	if !id.validID() {
		errs = append(errs, fmt.Errorf("Invalid identifier: %v", id.Type))
		result = false
	}
	return
}

func (id *Identifier) validID() bool {
	empty := &uuid.UUID{}
	if id.ID.String() == empty.String() {
		return false
	}
	return true
}

/* helpers */

// UnmarshalIdentifier takes a STIX identifier string and converts it to a Identifier type
func UnmarshalIdentifier(raw string) (Identifier, error) {
	var maxParts = 2

	parts := strings.Split(raw, identifierJoin)
	if len(parts) != maxParts {
		return Identifier{}, fmt.Errorf("Invalid STIX ID")
	}
	id := Identifier{Type: parts[0]}

	var err error
	id.ID, err = uuid.FromString(parts[1])
	return id, err
}
