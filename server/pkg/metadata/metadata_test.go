package metadata

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMeta(t *testing.T) {
	m := All()
	require.True(t, len(m.Amps) > 1)
	require.True(t, len(m.Amps[0].Channels) > 1)
}