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

func TestBundleValid(t *testing.T) {
	id, _ := NewIdentifier("bundle")

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

func TestMarshalBundle(t *testing.T) {
	id := Identifier{Type: "bundle"}
	id.ID, _ = uuid.FromString("5d0092c5-5f74-4287-9642-33f4c354e56d")

	bundle := Bundle{Type: "bundle", ID: id, SpecVersion: "2.0"}
	bundle.AddObject(`{"type": "malware", "id": "malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b"}`)

	if len(bundle.Objects) != 1 {
		t.Error("Got:", len(bundle.Objects), "Expected:", 1)
	}

	b, err := json.Marshal(bundle)
	if err != nil {
		t.Fatal(err)
	}

	// no whitespace for exact matching; seems brittle...
	expected := `{"type":"bundle","id":"bundle--5d0092c5-5f74-4287-9642-33f4c354e56d","spec_version":"2.0",` +
		`"objects":[{"type":"malware","id":"malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b"}]}`

	if string(b) != expected {
		t.Error("Got:", string(b), "Expected:", expected)
	}
}

func TestNewBundle(t *testing.T) {
	b, _ := NewBundle()
	empty := uuid.UUID{}

	if b.ID.String() == empty.String() {
		t.Error("Got:", b.ID.String(), "Expected: Not an empty UUID", empty.String())
	}
}

func TestUnmarshalBundle(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/malware_bundle.json")
	if err != nil {
		t.Fatal(err)
	}

	var bundle Bundle
	err = json.Unmarshal(b, &bundle)
	if err != nil {
		t.Fatal(err)
	}

	if len(bundle.Objects) == 0 {
		t.Error("Got:", len(bundle.Objects), "Expected: > 0")
	}

	expected := `{
      "type": "indicator",
      "id": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
      "created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
      "created": "2016-04-06T20:03:48.000Z",
      "modified": "2016-04-06T20:03:48.000Z",
      "labels": ["malicious-activity"],
      "name": "Poison Ivy Malware",
      "description": "This file is part of Poison Ivy",
      "pattern": "[ file:hashes.'SHA-256' = '4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877' ]",
      "valid_from": "2016-01-01T00:00:00Z"
    }`

	if string(bundle.Objects[0]) != expected {
		t.Error("Got:", string(b), "Expected:", expected)
	}
}
