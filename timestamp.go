package stones

import "time"

// Timestamp represents a a STIX date/time format
// alias works (type Timestamp = time.Time) but can't extend it
type Timestamp struct {
	time.Time
}

// NewTimestamp returns a new Timestamp based on a timestamp string
func NewTimestamp(s string) (Timestamp, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	return Timestamp{t}, err
}

// MarshalJSON will serialize Timestamp into a JSON string
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Timestamp) String() string {
	return t.Format(time.RFC3339Nano)
}
