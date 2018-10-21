package stones

import (
	"testing"
)

func BenchmarkBundleValid(b *testing.B) {
	bundle, _ := NewBundle()

	for i := 0; i < b.N; i++ {
		bundle.Valid()
	}
}

func BenchmarkNewBundle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBundle()
	}
}

func BenchmarkNewBundleAndValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bundle, _ := NewBundle()
		bundle.Valid()
	}
}
