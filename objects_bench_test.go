package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchNewObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObject("bundle")
	}
}

func BenchObjectValid(b *testing.B) {
	o, _ := NewObject("test")

	for i := 0; i < b.N; i++ {
		o.Valid()
	}
}

func BenchUnmarshalObject(b *testing.B) {
	d, err := ioutil.ReadFile("testdata/malware.json")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var o Object
		json.Unmarshal(d, &o)
	}
}
