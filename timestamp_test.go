package stones

import (
	"encoding/json"
	"testing"
)

func TestTimestampFromString(t *testing.T) {
	tests := []struct {
		timestamp string
		hasError  bool
	}{
		{"2016-04-06T20:07:09.000Z", false},
		{"malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b", true},
	}

	for _, test := range tests {
		ts, err := TimestampFromString(test.timestamp)
		if test.hasError && err == nil {
			t.Error("Expected an error", "Test:", ts)
		}
	}
}

func TestTimestampMarshal(t *testing.T) {
	ts, err := TimestampFromString("2016-04-06T20:07:09.000Z")
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(ts)
	expected := `"2016-04-06T20:07:09Z"`

	if string(b) != expected {
		t.Error("Got:", string(b), "Expected:", expected)
	}
}

func TestTimestampString(t *testing.T) {
	ts := Timestamp{}
	result := ts.String()
	expected := "0001-01-01T00:00:00Z"

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}
