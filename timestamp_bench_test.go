package stones

import (
	"encoding/json"
	"testing"
	"time"
)

var now = time.Now().UTC()

func BenchmarkMarshalTimestamp(b *testing.B) {
	ts, err := NewTimestamp("2016-04-06T20:07:09.000Z")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(ts)
	}
}

func BenchmarkNewTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTimestamp(now.Format(time.RFC3339Nano))
	}
}
