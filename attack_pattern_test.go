package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestAttackPatternSTIXObjectType(t *testing.T) {
	o := validTestObject(attackPatternType)
	ap := AttackPattern{Object: o}
	ap.Name = "test"

	result := ap.STIXObjectType()
	expected := domainObject

	if result != domainObject {
		t.Error("Got:", result, "Expected:", expected)
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
		ap := AttackPattern{Object: test.obj}
		ap.Name = test.name

		valid, errs := ap.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Attack Pattern:", ap, "Errs:", errs)
		}
	}
}

func TestAttackPatternUnmarshal(t *testing.T) {
	tests := []struct {
		file        string
		expectError bool
	}{
		{"testdata/attack_pattern.json", false},
		{"testdata/attack_pattern_invalid_name.json", true},
		{"testdata/attack_pattern_invalid_object.json", true},
	}

	for _, test := range tests {
		in, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Fatal(err)
		}

		var o AttackPattern
		err = json.Unmarshal(in, &o)
		if err != nil && !test.expectError {
			t.Error("Expected no error:", err)
		}
		if err == nil && test.expectError {
			t.Error("Expected error from file:", test.file)
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
		{"test", validTestObject(attackPatternType), nil, true},
		{"test", validTestObject(attackPatternType), []KillChainPhase{KillChainPhase{PhaseName: "foo"}}, false},
	}

	for _, test := range tests {
		ap := AttackPattern{Object: test.obj}
		ap.Name = test.name
		ap.KillChainPhases = test.killChainPhases

		valid, errs := ap.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Attack Pattern:", ap, "Errs:", errs)
		}
	}
}

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

func TestValidAttackPattern(t *testing.T) {
	invalidAttackPattern, err := ioutil.ReadFile("testdata/attack_pattern_invalid_name.json")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		object []byte
		valid  bool
	}{
		{attackPatternJSON(), true},
		{malwareJSON(), false},
		{invalidAttackPattern, false},
	}

	for _, test := range tests {
		valid, errs := validAttackPattern(test.object)
		if valid != test.valid {
			t.Error(
				"Got:", valid,
				"Expected:", test.valid,
				"Validation Errors:", errorsToString(errs),
				"Object:", string(test.object))
		}
	}
}
