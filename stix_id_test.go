package stones

import (
	"testing"
)

func TestNewStixID(t *testing.T) {
	tests := []struct {
		stixType string
		id       string
		expected string
		valid    bool
	}{
		{"bundle", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"nonType", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "nonType--00000000-0000-0000-0000-000000000000", false},
		{"bundle", "error", "bundle--00000000-0000-0000-0000-000000000000", false},
	}

	for _, test := range tests {
		result, err := NewStixID(test.stixType, test.id)
		if !test.valid && err == nil {
			t.Error("Should have generated an error", "Test:", test)
		}

		if result.String() != test.expected {
			t.Error("Got:", result, "Expected:", test.expected)
		}
	}
}

func TestValidStixID(t *testing.T) {
	tests := []struct {
		id    string
		valid bool
	}{
		{"bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"invalidType--6ba7b810-9dad-11d1-80b4-00c04fd430c8", false},
		{"bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8--extra", false},
		{"bundle--00000000-0000-0000-0000-000000000000", false},
		{"bundle", false},
		{"bundle--invalid-id", false},
	}

	for _, test := range tests {
		valid, err := validStixID(test.id)
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "ID:", test.id, "Error:", err)
		}
	}
}

func TestStixIDValidate(t *testing.T) {
	validStixID, _ := NewStixID("bundle")

	tests := []struct {
		id    StixID
		valid bool
	}{
		{validStixID, true},
		{StixID{}, false},
	}

	for _, test := range tests {
		valid, _ := test.id.validate()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test)
		}
	}
}
