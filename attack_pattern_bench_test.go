package stones

import (
	"testing"
)

func BenchmarkAttackPatternValid(b *testing.B) {
	attackPattern, _ := NewAttackPattern("test")

	for i := 0; i < b.N; i++ {
		attackPattern.Valid()
	}
}

func BenchmarkNewAttackPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewAttackPattern("test")
	}
}

func BenchmarkNewAttackPatternAndValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		attackPattern, _ := NewAttackPattern("test")
		attackPattern.Valid()
	}
}
