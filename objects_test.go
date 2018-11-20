package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewObject(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/malware.json")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewObject(b)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObjectSerializeDeserialize(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/malware.json")

	var o Object
	err = json.Unmarshal(in, &o)
	if err != nil {
		t.Fatal(err)
	}

	_, err = json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObjectValid(t *testing.T) {
	tests := []struct {
		testFile    string
		expectError bool
	}{
		{"testdata/malware.json", false},
		{"testdata/malware_invalid_type.json", true},
		{"testdata/malware_type_mismatch.json", true},
	}

	for _, test := range tests {
		b, err := ioutil.ReadFile(test.testFile)
		if err != nil {
			t.Fatal(err)
		}

		o, err := NewObject(b)
		if err != nil {
			t.Fatal(err)
		}

		_, errs := o.Valid()
		if test.expectError && len(errs) == 0 {
			t.Error("Expected error, Test:", test, "Object:", o)
		}

		if !test.expectError && len(errs) > 0 {
			t.Error("Error unexpected, Test:", test, "Object:", o, "Errors:", errs)
		}
	}
}
