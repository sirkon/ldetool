package ldetesting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicBeforeCharDecoder(t *testing.T) {
	d := &Decoders{}
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
	ok, err = d.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdefrest`))
	require.False(t, ok)

	dl := &DecodersLimited{}
	ok, err = dl.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(1), dl.Int8)
	require.Equal(t, int16(2), dl.Int16)
	require.Equal(t, int32(3), dl.Int32)
	require.Equal(t, int64(4), dl.Int64)
	require.Equal(t, uint8(5), dl.Uint8)
	require.Equal(t, uint16(6), dl.Uint16)
	require.Equal(t, uint32(7), dl.Uint32)
	require.Equal(t, uint64(8), dl.Uint64)
	require.Equal(t, float32(9), dl.Float32)
	require.Equal(t, float64(11e7), dl.Float64)
	require.Equal(t, "abcdef", string(dl.String))
	require.Equal(t, "rest", string(dl.rest))

	db := &DecodersBounded{}
	ok, err = db.Parse([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int64(11122321312313), db.Int64)
	require.Equal(t, "        ", string(db.rest))
}

func TestBasicBeforeStringDecoder(t *testing.T) {
	d := &DecodersString{}
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

	dl := &DecodersLimitedString{}
	ok, err = dl.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(1), dl.Int8)
	require.Equal(t, int16(2), dl.Int16)
	require.Equal(t, int32(3), dl.Int32)
	require.Equal(t, int64(4), dl.Int64)
	require.Equal(t, uint8(5), dl.Uint8)
	require.Equal(t, uint16(6), dl.Uint16)
	require.Equal(t, uint32(7), dl.Uint32)
	require.Equal(t, uint64(8), dl.Uint64)
	require.Equal(t, float32(9), dl.Float32)
	require.Equal(t, float64(11e7), dl.Float64)
	require.Equal(t, "abcdef", string(dl.String))
	require.Equal(t, "rest", string(dl.rest))

	db := &DecodersBoundedString{}
	ok, err = db.Parse([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, float64(11122321312313), db.Float64)
	require.Equal(t, "        ", string(db.rest))
}

func TestBasicBeforeCharStressDecoder(t *testing.T) {
	d := &DecodersStress{}
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
	ok, err = d.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7abcdef rest`))
	require.NotNil(t, err)

	dl := &DecodersLimitedStress{}
	ok, err = dl.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(1), dl.Int8)
	require.Equal(t, int16(2), dl.Int16)
	require.Equal(t, int32(3), dl.Int32)
	require.Equal(t, int64(4), dl.Int64)
	require.Equal(t, uint8(5), dl.Uint8)
	require.Equal(t, uint16(6), dl.Uint16)
	require.Equal(t, uint32(7), dl.Uint32)
	require.Equal(t, uint64(8), dl.Uint64)
	require.Equal(t, float32(9), dl.Float32)
	require.Equal(t, float64(11e7), dl.Float64)
	require.Equal(t, "abcdef", string(dl.String))
	require.Equal(t, "rest", string(dl.rest))

	db := &DecodersBoundedStress{}
	ok, err = db.Parse([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int64(11122321312313), db.Int64)
	require.Equal(t, "        ", string(db.rest))
}

func TestBasicBeforeStringStressDecoder(t *testing.T) {
	d := &DecodersStringStress{}
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

	dl := &DecodersLimitedStringStress{}
	ok, err = dl.Parse([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(1), dl.Int8)
	require.Equal(t, int16(2), dl.Int16)
	require.Equal(t, int32(3), dl.Int32)
	require.Equal(t, int64(4), dl.Int64)
	require.Equal(t, uint8(5), dl.Uint8)
	require.Equal(t, uint16(6), dl.Uint16)
	require.Equal(t, uint32(7), dl.Uint32)
	require.Equal(t, uint64(8), dl.Uint64)
	require.Equal(t, float32(9), dl.Float32)
	require.Equal(t, float64(11e7), dl.Float64)
	require.Equal(t, "abcdef", string(dl.String))
	require.Equal(t, "rest", string(dl.rest))

	db := &DecodersBoundedStringStress{}
	ok, err = db.Parse([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, float64(11122321312313), db.Float64)
	require.Equal(t, "        ", string(db.rest))
}

func TestDecoderOptionals(t *testing.T) {
	d := &DecoderOptionals{}

	ok, err := d.Parse([]byte(`12 head=13 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), d.Int8)
	require.Equal(t, "13", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`12 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), d.Int8)
	require.Equal(t, "", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`12 head=13`))
	require.False(t, ok)
	if err != nil {
		t.Fatal(err)
	}

	ds := &DecoderOptionalsStress{}

	ok, err = ds.Parse([]byte(`12 head=13 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), ds.Int8)
	require.Equal(t, "13", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`12 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), ds.Int8)
	require.Equal(t, "", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`12 head=13`))
	require.False(t, ok)
	require.NotNil(t, err)
}

func TestDecoderBranching(t *testing.T) {
	d := &DecoderBranching{}

	ok, err := d.Parse([]byte(`start head=data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`start head=data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`start data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`start data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Parse([]byte(`data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))
	ds := &DecoderBranchingStress{}
	ok, err = ds.Parse([]byte(`start head=data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`start head=data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`start data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`start data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Parse([]byte(`data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))
}
