package stones

import (
	"testing"
)

func BenchmarkIDValid(b *testing.B) {
	id := ID("malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b")

	for i := 0; i < b.N; i++ {
		id.Valid()
	}
}

func BenchIdentifierValid(b *testing.B) {
	id, _ := NewIdentifier("bundle")

	for i := 0; i < b.N; i++ {
		id.Valid()
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

func BenchmarkUnmarshalIdentifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshalIdentifier("bundle--6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	}
}
