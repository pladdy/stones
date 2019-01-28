package stones

import (
	"encoding/json"
	"testing"
	"time"
)

var now = time.Now().UTC()

func BenchmarkMarshalTimestamp(b *testing.B) {
	ts, err := TimestampFromString("2016-04-06T20:07:09.000Z")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(ts)
	}
}

func BenchmarkTimestampFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimestampFromString(now.Format(time.RFC3339Nano))
	}
}
