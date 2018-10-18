package stones

import (
	"testing"
)

func TestIDValid(t *testing.T) {
	tests := []struct {
		id    string
		valid bool
	}{
		{"malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b", true},
		{"31b940d4-6f7f-459a-80ea-9c1f17b5891b", false},
		{"malware", false},
		{"", false},
	}

	for _, test := range tests {
		id := ID(test.id)
		valid, err := id.Valid()

		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "ID:", id, "Error:", err)
		}
	}
}

func TestUnmarshalIdentifier(t *testing.T) {
	tests := []struct {
		rawID        string
		expectedType string
		expectedID   string
	}{
		{"bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8", "bundle", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
		{"bundle--bad", "bundle", "00000000-0000-0000-0000-000000000000"},
	}

	for _, test := range tests {
		s, _ := UnmarshalIdentifier(test.rawID)
		if s.Type != test.expectedType {
			t.Error("Got:", s.Type, "Expected:", test.expectedType)
		}

		if s.ID.String() != test.expectedID {
			t.Error("Got:", s.ID.String(), "Expected:", test.expectedID)
		}
	}
}

func TestNewIdentifier(t *testing.T) {
	tests := []struct {
		stixType     string
		expectedType string
	}{
		{"bundle", "bundle"},
		{"nonType", "nonType"},
	}

	for _, test := range tests {
		result, err := NewIdentifier(test.stixType)
		if result.Type != test.expectedType {
			t.Error("Got:", result, "Expected:", test.expectedType)
		}

		if err != nil {
			t.Error("Got:", err, "Expected no error")
		}
	}
}

func TestIdentifierValid(t *testing.T) {
	validIdentifier, _ := NewIdentifier("bundle")

	tests := []struct {
		id    Identifier
		valid bool
	}{
		{validIdentifier, true},
		{Identifier{}, false},
	}

	for _, test := range tests {
		valid, _ := test.id.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test)
		}
	}
}
