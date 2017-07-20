package ldetesting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicDecoder(t *testing.T) {
	d := &decoders{}
	ok, err := d.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(1), d.Int8)
	require.Equal(t, int16(2), d.Int16)
	require.Equal(t, int32(3), d.Int32)
	require.Equal(t, int64(4), d.Int64)
	require.Equal(t, uint8(5), d.Uint8)
	require.Equal(t, uint16(6), d.Uint16)
	require.Equal(t, uint32(7), d.Uint32)
	require.Equal(t, uint64(8), d.Uint64)
	require.Equal(t, float32(9), d.Float32)
	require.Equal(t, float64(11e7), d.Float64)
	require.Equal(t, "abcdef", string(d.String))
	require.Equal(t, "rest", string(d.rest))
}
