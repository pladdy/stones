package stones

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var ap AttackPattern
var valid bool
var err error
var errs []error

func BenchmarkAttackPatternValid(b *testing.B) {
	attackPattern, _ := NewAttackPattern("test")

	for i := 0; i < b.N; i++ {
		valid, errs = attackPattern.Valid()
	}
}

func BenchmarkNewAttackPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ap, err = NewAttackPattern("test")
	}
}

func BenchmarkNewAttackPatternAndValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ap, err = NewAttackPattern("test")
		valid, errs = ap.Valid()
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
		valid, errs = validAttackPattern(d)
	}
}
