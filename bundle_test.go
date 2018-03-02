package stones

import "testing"

func TestNewBundle(t *testing.T) {
	b, _ := NewBundle()
	empty := StixID{}

	if b.ID == empty {
		t.Error("Got:", b.ID, "Expected: NOT", empty)
	}
}
