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
		if id.Valid() != test.valid {
			t.Error("Got:", id.Valid(), "Expected:", test.valid, "ID:", id)
		}
	}
}

func TestMarshalStixID(t *testing.T) {
	tests := []struct {
		rawID        string
		expectedType string
		expectedID   string
	}{
		{"bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8", "bundle", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
		{"bundle--bad", "bundle", "00000000-0000-0000-0000-000000000000"},
	}

	for _, test := range tests {
		s, _ := MarshalStixID(test.rawID)
		if s.Type != test.expectedType {
			t.Error("Got:", s.Type, "Expected:", test.expectedType)
		}

		if s.ID.String() != test.expectedID {
			t.Error("Got:", s.ID.String(), "Expected:", test.expectedID)
		}
	}
}

func TestNewStixID(t *testing.T) {
	tests := []struct {
		stixType     string
		expectedType string
		valid        bool
	}{
		{"bundle", "bundle", true},
		{"nonType", "nonType", false},
	}

	for _, test := range tests {
		result, err := NewStixID(test.stixType)
		if !test.valid && err == nil {
			t.Error("Should have generated an error", "Test:", test)
		}

		if result.Type != test.expectedType {
			t.Error("Got:", result, "Expected:", test.expectedType)
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
