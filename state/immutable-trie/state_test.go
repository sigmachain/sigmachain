package itrie

import (
	"testing"

	"github.com/sigmachain/sigmachain/state"
)

func TestState(t *testing.T) {
	state.TestState(t, buildPreState)
}

func buildPreState(pre state.PreStates) state.Snapshot {
	storage := NewMemoryStorage()
	st := NewState(storage)
	snap := st.NewSnapshot()

	return snap
}
