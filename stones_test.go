package stones

import (
	"testing"
)

func TestNewStixID(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		shouldError bool
	}{
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", false},
		{"error", "00000000-0000-0000-0000-000000000000", true},
	}

	for _, test := range tests {
		result, err := NewStixID(test.input)
		if test.shouldError && err == nil {
			t.Error("Should have generated an error")
		}

		if result.String() != test.expected {
			t.Error("Got:", result.String(), "Expected:", test.expected)
		}
	}
}
