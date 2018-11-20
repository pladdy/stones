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
