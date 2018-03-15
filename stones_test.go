package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

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

func TestMarshalBundle(t *testing.T) {
	bundle := Bundle{Type: "bundle", ID: "bundle--5d0092c5-5f74-4287-9642-33f4c354e56d", SpecVersion: "2.0"}
	bundle.AddObject(`{"type": "malware", "id": "malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b"}`)

	if len(bundle.Objects) != 1 {
		t.Error("Got:", len(bundle.Objects), "Expected:", 1)
	}

	b, err := json.Marshal(bundle)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"type":"bundle","id":"bundle--5d0092c5-5f74-4287-9642-33f4c354e56d","spec_version":"2.0",` +
		`"objects":[{"type":"malware","id":"malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b"}]}`

	if string(b) != expected {
		t.Error("Got:", string(b), "Expected:", expected)
	}
}
