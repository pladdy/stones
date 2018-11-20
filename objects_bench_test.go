package stones

import (
	"io/ioutil"
	"testing"
)

// not sure if this benchmark is useful; it's a small json file/object...
func BenchmarkNewObject(b *testing.B) {
	content, err := ioutil.ReadFile("testdata/malware.json")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		NewObject(content)
	}
}
