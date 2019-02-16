package stones

import (
	"encoding/json"
	"testing"

	"github.com/gofrs/uuid"
)

func TestIdentifierMarshal(t *testing.T) {
	id, err := IdentifierFromString("bundle--5d0092c5-5f74-4287-9642-33f4c354e56d")
	if err != nil {
		t.Fatal(err)
	}

	i, err := json.Marshal(&id)
	if err != nil {
		t.Fatal(err)
	}

	expected := `"bundle--5d0092c5-5f74-4287-9642-33f4c354e56d"`

	if string(i) != expected {
		t.Error("Got:", string(i), "Expected:", expected)
	}
}

func TestIdentifierUnmarshalJSON(t *testing.T) {
	testUUID, _ := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	tests := []struct {
		rawJSON            string
		expectedIdentifier Identifier
		expectError        bool
	}{
		{`"bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8"`, Identifier{Type: "bundle", ID: testUUID}, false},
		{`bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8`, Identifier{Type: "bundle", ID: testUUID}, true},
		{`"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`, Identifier{Type: "bundle", ID: testUUID}, true},
	}

	for _, test := range tests {
		eid := test.expectedIdentifier

		var id Identifier
		err := id.UnmarshalJSON([]byte(test.rawJSON))

		if test.expectError && err == nil {
			t.Error("Expected error for test:", test)
		}

		if !test.expectError && err != nil {
			if id.Type != eid.Type {
				t.Error("Got:", id.Type, "Expected:", eid.Type)
			}

			if id.ID.String() != eid.ID.String() {
				t.Error("Got:", id.ID.String(), "Expected:", eid.ID.String())
			}
		}
	}
}

func TestIdentifierValid(t *testing.T) {
	validIdentifier, _ := NewIdentifier(bundleType)

	// test with unsupported uuid versions
	ns := uuid.Must(uuid.FromString("123e4567-e89b-12d3-a456-426655440000"))
	v1UUID := uuid.Must(uuid.NewV1())
	// v2 is not supported in the uuid library as of ~2019-02; removing from test
	v3UUID := uuid.NewV3(ns, "test")
	v4UUID := uuid.Must(uuid.NewV4())
	v5UUID := uuid.NewV5(ns, "test")

	tests := []struct {
		id    Identifier
		valid bool
	}{
		{validIdentifier, true},
		{Identifier{}, false},
		{Identifier{Type: bundleType, ID: v1UUID}, false},
		{Identifier{Type: bundleType, ID: v3UUID}, false},
		{Identifier{Type: bundleType, ID: v4UUID}, true},
		{Identifier{Type: bundleType, ID: v5UUID}, false},
	}

	for _, test := range tests {
		valid, _ := test.id.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test)
		}
	}
}

func TestNewIdentifier(t *testing.T) {
	tests := []struct {
		stixType     string
		expectedType string
	}{
		{"bundle", "bundle"},
		{"nonType", "nonType"},
	}

	for _, test := range tests {
		result, err := NewIdentifier(test.stixType)
		if result.Type != test.expectedType {
			t.Error("Got:", result, "Expected:", test.expectedType)
		}

		if err != nil {
			t.Error("Got:", err, "Expected no error")
		}
	}
}
