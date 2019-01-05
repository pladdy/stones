package stones

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

const identifierJoin = "--"

// Identifier represents a STIX Identifier Data Type.
//
// Identifiers uniquely identify STIX Objects.  They follow the form of <object type>--<uuid V4>.
//
// The ID field is a v4 UUID
//   v4 UUID:
//     128 bit; 16 octets of 32 hexadecimal numbers
//     String representation: 32 bit - 16 bit - 16 bit - 16 bit - 48 bit
//     Example:               6ba7b810-9dad-11d1-80b4-00c04fd430c8
type Identifier struct {
	Type string
	ID   uuid.UUID
}

// NewIdentifier takes a STIX Type string and returns an Identifier struct.
func NewIdentifier(t string) (id Identifier, err error) {
	id = Identifier{Type: t}
	id.ID, err = uuid.NewV4()
	return id, err
}

// IdentifierFromString takes a Identifier string representation and returns an Identifier struct.
func IdentifierFromString(s string) (id Identifier, err error) {
	var maxParts = 2

	parts := strings.Split(s, identifierJoin)
	if len(parts) != maxParts {
		return id, fmt.Errorf("Invalid 'Identifier': %v", s)
	}

	id.Type = parts[0]
	id.ID, err = uuid.FromString(parts[1])
	return
}

// MarshalJSON implements the encoding/json Marshaler interface (https://golang.org/pkg/encoding/json/#Marshaler)
//
// It is used to serialize an Identifier into JSON format
func (id Identifier) MarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

func (id *Identifier) String() string {
	return strings.Join([]string{id.Type, id.ID.String()}, identifierJoin)
}

// UnmarshalJSON implements the encoding/json Unmarshaler interface (https://golang.org/pkg/encoding/json/#Unmarshaler).
//
// It will take JSON and deserialize to an Object.  This should not be called directly, but instead
// json.Unmarshal(b []byte, v interface{}) should be used.
func (id *Identifier) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	newID, err := IdentifierFromString(s)
	id.Type, id.ID = newID.Type, newID.ID

	return err
}

// Valid is called to check for STiX 2.0 specification conformance.
//
// If the Identifier is invalid, it returns the list of errors from validation.
func (id *Identifier) Valid() (valid bool, errs []error) {
	if !validStixType(id.Type) {
		errs = append(errs, fmt.Errorf("Invalid 'Type' in Identifier: %v", id.Type))
	}

	if !id.validID() {
		errs = append(errs, fmt.Errorf("Invalid 'UUID': %v", id.ID))
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
