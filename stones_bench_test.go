package stones

import (
	"testing"
)

func BenchmarkKillChainPhaseValid(b *testing.B) {
	kcp := KillChainPhase{"kill chain name", "phase name"}

	for i := 0; i < b.N; i++ {
		kcp.Valid()
	}
}
