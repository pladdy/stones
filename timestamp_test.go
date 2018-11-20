package stones

import "testing"

func TestNewTimestamp(t *testing.T) {
	tests := []struct {
		timestamp string
		hasError  bool
	}{
		{"2016-04-06T20:07:09.000Z", false},
		{"malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b", true},
	}

	for _, test := range tests {
		ts, err := NewTimestamp(test.timestamp)
		if test.hasError && err == nil {
			t.Error("Expected an error", "Test:", ts)
		}
	}
}
