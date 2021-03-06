package stones

import (
	"encoding/json"
	"testing"
)

func BenchIdentifierValid(b *testing.B) {
	id, _ := NewIdentifier("bundle")

	for i := 0; i < b.N; i++ {
		id.Valid()
	}
}

func BenchmarkIdentifierFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IdentifierFromString("bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	}
}

func BenchmarkNewIdentifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewIdentifier("malware")
	}
}

func BenchNewIdentifierAndValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id, _ := NewIdentifier("bundle")
		id.Valid()
	}
}

func BenchmarkMarshalIdentifier(b *testing.B) {
	id, err := NewIdentifier("bundle")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(id)
	}
}

func BenchmarkUnmarshalIdentifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var id Identifier
		json.Unmarshal([]byte(`{"id": "bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8"}`), &id)
	}
}
