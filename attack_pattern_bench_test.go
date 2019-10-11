package stones

import (
	"encoding/json"
	"io/ioutil"
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
		ap, _ := NewAttackPattern("test")
		ap.Valid()
	}
}

func BenchmarkUnmarshalAttackPattern(b *testing.B) {
	d, err := ioutil.ReadFile("testdata/attack_pattern.json")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var o AttackPattern
		json.Unmarshal(d, &o)
	}
}

func BenchmarkValidateAttackPattern(b *testing.B) {
	d, err := ioutil.ReadFile("testdata/attack_pattern.json")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		validAttackPattern(d)
	}
}
