package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewObject(t *testing.T) {
	_, err := NewObject("malware")
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
		{"testdata/malware_invalid_created.json", true},
		{"testdata/malware_invalid_modified.json", true},
	}

	for _, test := range tests {
		b, err := ioutil.ReadFile(test.testFile)
		if err != nil {
			t.Fatal(err)
		}

		// ignore the error from Unmarshal (which will include any Validation errors
		var o Object
		if err := json.Unmarshal(b, &o); err == nil && test.expectError {
			t.Error("Expected an error for object:", string(b))
		}

		// test Valid itself
		_, errs := o.Valid()
		if test.expectError && len(errs) == 0 {
			t.Error("Expected error, Test:", test, "Object:", o)
		}

		if !test.expectError && len(errs) > 0 {
			t.Error("Error unexpected, Test:", test, "Object:", o, "Errors:", errs)
		}
	}
}
