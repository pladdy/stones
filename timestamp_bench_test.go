package stones

import (
	"testing"
	"time"
)

var now = time.Now().UTC()

func BenchmarkNewTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTimestamp(now.Format(time.RFC3339Nano))
	}
}
