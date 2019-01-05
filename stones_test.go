package stones

import (
	"fmt"
	"testing"
)

func TestErrorsToString(t *testing.T) {
	errs := []error{fmt.Errorf("a fake error"), fmt.Errorf("another fake error")}

	result := fmt.Sprintf("%v", ErrorsToString(errs))
	expected := "a fake error; another fake error"

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestKillChainPhaseValid(t *testing.T) {
	tests := []struct {
		killChainName string
		phaseName     string
		valid         bool
	}{
		{"", "", false},
		{"", "test", false},
		{"test", "", false},
		{"test", "test", true},
	}

	for _, test := range tests {
		kcp := KillChainPhase{}
		kcp.KillChainName = test.killChainName
		kcp.PhaseName = test.phaseName

		valid, _ := kcp.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Kill Chain Phase:", kcp)
		}
	}
}

func TestObjectType(t *testing.T) {
	tests := []struct {
		objectSource string
		expected     string
	}{
		{`{"type": "foo"}`, "foo"},
		{`{"id": "foo---some-idenfitier", "type": "foo"}`, "foo"},
		{`{"id": "foo---some-idenfitier", "type": "foo", "created_at": "2018-12-12 20:35:35Z"}`, "foo"},
		{`{"id": "foo---some-idenfitier", "created_at": "2018-12-12 20:35:35Z"}`, ""},
	}

	for _, test := range tests {
		result := objectType([]byte(test.objectSource))
		if result != test.expected {
			t.Error("Got:", result, "Exepcted:", test.expected)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		object string
		valid  bool
	}{
		{`{"type": "foo"}`, false},
		{`{"type": "attack-pattern"}`, false},
	}

	for _, test := range tests {
		result, err := Validate([]byte(test.object))
		if result != test.valid {
			t.Error("Got:", result, "Expected:", test.valid, "Object:", test.object, "Error:", err)
		}
	}
}
