package stones

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

const identifierJoin = "--"

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

// IdentifierFromString takes a identifier string and returns an Identifier
func IdentifierFromString(s string) (id Identifier, err error) {
	var maxParts = 2

	parts := strings.Split(s, identifierJoin)
	if len(parts) != maxParts {
		return id, fmt.Errorf("Invalid STIX ID")
	}

	id.Type = parts[0]
	id.ID, err = uuid.FromString(parts[1])
	return
}

// UnmarshalJSON is used by encoding/json.Unmarshal to Unmarshal JSON encodings to Identifier types
func (id *Identifier) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	newID, err := IdentifierFromString(s)
	id.Type, id.ID = newID.Type, newID.ID
	return err
}

// Valid will validate an Identifier
func (id *Identifier) Valid() (valid bool, errs []error) {
	if !validStixType(id.Type) {
		errs = append(errs, fmt.Errorf("Invalid STIX type: %v", id.Type))
	}

	if !id.validID() {
		errs = append(errs, fmt.Errorf("Invalid identifier: %v", id.Type))
	}

	if len(errs) == 0 {
		valid = true
	}
	return
}

func (id *Identifier) validID() bool {
	empty := &uuid.UUID{}

	if id.ID.String() == empty.String() {
		return false
	}

	if id.ID.Version() != uuid.V4 {
		return false
	}
	return true
}
