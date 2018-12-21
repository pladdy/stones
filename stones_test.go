package stones

import (
	"testing"
)

func TestKillChainPhaseValid(t *testing.T) {
	tests := []struct {
		killChainName string
		phaseName     string
		valid         bool
	}{
		{"", "", false},
		{"", "test", false},
		{"test", "", false},
		{"test", "test", true},
	}

	for _, test := range tests {
		kcp := KillChainPhase{}
		kcp.KillChainName = test.killChainName
		kcp.PhaseName = test.phaseName

		valid, _ := kcp.Valid()
		if valid != test.valid {
			t.Error("Got:", valid, "Expected:", test.valid, "Test:", test, "Kill Chain Phase:", kcp)
		}
	}
}
