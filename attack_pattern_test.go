package stones

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestNewAttackPattern(t *testing.T) {
	tests := []struct {
		name      string
		hasErrors bool
	}{
		{"", false},
		{"some attack pattern", false},
	}

	for _, test := range tests {
		result, err := NewAttackPattern(test.name)

		if test.hasErrors && err == nil {
			t.Error("Got:", result, "Expected an error, test:", test)
		}
	}
}

func TestAttackPatternValid(t *testing.T) {
	testID, err := NewIdentifier(attackPatternType)
	if err != nil {
		t.Fatal(err)
	}

	invalidID, err := NewIdentifier(attackPatternType)
	if err != nil {
		t.Fatal(err)
	}
	invalidID.ID = uuid.UUID{}

	tests := []struct {
		newType string
		id      Identifier
		name    string
		valid   bool
	}{
		{"", testID, "", false},
		{attackPatternType, testID, "", false},
		{attackPatternType, invalidID, "test", false},
		{attackPatternType, testID, "test", true},
	}

	for _, test := range tests {
		// set up object
		ap := AttackPattern{}
		ap.ID = test.id
		ap.Type = test.newType
		ap.Name = test.name

		valid, _ := ap.Valid()

		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Attack Pattern:", ap)
		}
	}
}
