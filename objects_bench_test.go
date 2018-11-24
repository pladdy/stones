package stones

import (
	"testing"
)

// not sure if this benchmark is useful; it's a small json file/object...
func BenchmarkNewObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObject("malware")
	}
}
