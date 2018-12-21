package stones

import (
	"testing"
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

func TestAttackPatternValidNoKillChain(t *testing.T) {
	tests := []struct {
		name  string
		obj   Object
		valid bool
	}{
		{"", Object{}, false},
		{"", Object{Type: attackPatternType}, false},
		{"test", Object{}, false},
		{"test", validTestObject(attackPatternType), true},
	}

	for _, test := range tests {
		ap := AttackPattern{Object: test.obj, Name: test.name}

		valid, errs := ap.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Attack Pattern:", ap, "Errs:", errs)
		}
	}
}

func TestAttackPatternValidWithKillChain(t *testing.T) {
	tests := []struct {
		name            string
		obj             Object
		killChainPhases []KillChainPhase
		valid           bool
	}{
		{"test", validTestObject(attackPatternType), []KillChainPhase{}, true},
		{"test", validTestObject(attackPatternType), []KillChainPhase{KillChainPhase{PhaseName: "foo"}}, false},
	}

	for _, test := range tests {
		ap := AttackPattern{Object: test.obj, Name: test.name, KillChainPhases: test.killChainPhases}

		valid, errs := ap.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Attack Pattern:", ap, "Errs:", errs)
		}
	}
}
