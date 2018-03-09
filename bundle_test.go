package stones

import (
	"fmt"
	"testing"
)

func TestNewBundle(t *testing.T) {
	b, _ := NewBundle()
	empty := StixID{}

	if b.ID == empty {
		t.Error("Got:", b.ID, "Expected: NOT", empty)
	}
}

func TestBundleValidate(t *testing.T) {
	id, _ := NewStixID()
	objects := []string{"object 1", "object 2"}

	tests := []struct {
		bundle Bundle
		valid  bool
		errs   []error
	}{
		{bundle: Bundle{Type: "fail", ID: id, SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{invalidType("invalid type")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: "1.0", Objects: objects},
			valid: false,
			errs:  []error{invalidType("invalid spec version")}},
		{bundle: Bundle{Type: "bundle", SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{invalidType("invalid id")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: specVersion, Objects: []string{}},
			valid: false,
			errs:  []error{fmt.Errorf("no objects")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: specVersion, Objects: objects},
			valid: true,
			errs:  []error{}},
	}

	for _, test := range tests {
		valid, errs := test.bundle.Validate()

		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid)
		}

		if len(test.errs) != len(errs) {
			t.Error("Got:", len(test.errs), "Expected:", len(errs))
		}
	}
}

func BenchmarkBundleValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bundle, _ := NewBundle()
		bundle.Validate()
	}
}
