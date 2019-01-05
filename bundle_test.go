package stones

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/gofrs/uuid"
)

func TestBundleAddObject(t *testing.T) {
	tests := []struct {
		objectToAdd     string
		objectsInBundle int
	}{
		{``, 0},
		{`{"type": "malware"}`, 1},
	}

	for _, test := range tests {
		b, _ := NewBundle()
		b.AddObject(test.objectToAdd)

		if len(b.Objects) != test.objectsInBundle {
			t.Error("Got:", len(b.Objects), "Expected:", test.objectsInBundle)
		}
	}
}

func TestNewBundle(t *testing.T) {
	b, _ := NewBundle()
	empty := uuid.UUID{}

	if b.ID.String() == empty.String() {
		t.Error("Got:", b.ID.String(), "Expected: Not an empty UUID", empty.String())
	}
}

func TestBundleValid(t *testing.T) {
	id, err := NewIdentifier("bundle")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile("testdata/malware_bundle.json")
	if err != nil {
		t.Fatal(err)
	}
	objects := []json.RawMessage{b}

	tests := []struct {
		bundle Bundle
		valid  bool
		errs   []error
	}{
		{bundle: Bundle{Type: "fail", ID: id, SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{errors.New("wrong type")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: "1.0", Objects: objects},
			valid: false,
			errs:  []error{errors.New("wrong spec version")}},
		{bundle: Bundle{Type: "bundle", SpecVersion: specVersion, Objects: objects},
			valid: false,
			errs:  []error{errors.New("no id"), errors.New("invalid stix type")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: specVersion, Objects: []json.RawMessage{}},
			valid: false,
			errs:  []error{errors.New("no objects")}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: specVersion, Objects: objects},
			valid: true,
			errs:  []error{}},
		{bundle: Bundle{Type: "bundle", ID: id, SpecVersion: specVersion},
			valid: false,
			errs:  []error{errors.New("no objects")}},
	}

	for _, test := range tests {
		valid, errs := test.bundle.Valid()

		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid)
		}

		if len(errs) != len(test.errs) {
			t.Error("Got:", len(errs), "Expected:", len(test.errs))
		}
	}
}
