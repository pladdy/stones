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

func TestObjectUnmarshal(t *testing.T) {
	tests := []struct {
		file        string
		expectError bool
	}{
		{"testdata/malware.json", false},
		{"testdata/malware_invalid_object.json", true},
	}

	for _, test := range tests {
		in, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Fatal(err)
		}

		var o Object
		err = json.Unmarshal(in, &o)

		if err != nil && !test.expectError {
			t.Error("Expected no error:", err)
		}
		if err == nil && test.expectError {
			t.Error("Expected error from file:", test.file)
		}
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

		o, err := ObjectFromBytes(b)
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
