package stones

import (
	"testing"
)

func TestTimestampString(t *testing.T) {
	ts := Timestamp{}
	result := ts.String()
	expected := "0001-01-01T00:00:00Z"

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}
