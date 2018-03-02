package stones

import (
	uuid "github.com/satori/go.uuid"
)

// StixID is a v4 UUID
type StixID struct {
	uuid.UUID
}

// NewStixID returns a ID for a STIX object
func NewStixID() (StixID, error) {
	id, err := uuid.NewV4()
	return StixID{id}, err
}
