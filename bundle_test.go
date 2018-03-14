package stones

import (
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestNewBundle(t *testing.T) {
	b, _ := NewBundle()
	empty := uuid.UUID{}

	if b.ID == empty.String() {
		t.Error("Got:", b.ID, "Expected: NOT", empty.String())
	}
}

func TestBundleValidate(t *testing.T) {
	id, _ := newStixID("bundle")
	ids := id.String()
	objects := []string{"object 1", "object 2"}

	tests := []struct {
		bundle Bundle
		valid  bool
		errs   []error
	}{
		{bundle: Bundle{Type: "fail", ID: ids, SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{errors.New("wrong type")}},
		{bundle: Bundle{Type: "bundle", ID: ids, SpecVersion: "1.0", Objects: objects},
			valid: false,
			errs:  []error{errors.New("wrong spec version")}},
		{bundle: Bundle{Type: "bundle", SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{errors.New("no id")}},
		{bundle: Bundle{Type: "bundle", ID: ids, SpecVersion: specVersion, Objects: []string{}},
			valid: false,
			errs:  []error{errors.New("no objects")}},
		{bundle: Bundle{Type: "bundle", ID: ids, SpecVersion: specVersion, Objects: objects},
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
