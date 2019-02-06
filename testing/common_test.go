package ldetesting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommon(t *testing.T) {
	data := `[bugaga] -123  234 abcdef`
	e := &Rule{}
	if ok, err := e.Extract(data); !ok || err != nil {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "bugaga", e.Data)
	require.Equal(t, -123, e.Signed)
	require.Equal(t, uint(234), e.Unsigned)
	require.Equal(t, "abcdef", e.Str)
}
