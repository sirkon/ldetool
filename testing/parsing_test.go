package ldetesting

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicBeforeCharDecoder(t *testing.T) {
	d := &Decoders{}
	ok, err := d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(d.Rest))
	ok, err = d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdefrest`))
	require.False(t, ok)

	dl := &DecodersLimited{}
	ok, err = dl.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest               `))
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
	require.Equal(t, "rest               ", string(dl.Rest))

	db := &DecodersBounded{}
	ok, err = db.Extract([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int64(11122321312313), db.Int64)
	require.Equal(t, "        ", string(db.Rest))
}

func TestBasicBeforeStringDecoder(t *testing.T) {
	d := &DecodersString{}
	ok, err := d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(d.Rest))

	dl := &DecodersLimitedString{}
	ok, err = dl.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(dl.Rest))

	db := &DecodersBoundedString{}
	ok, err = db.Extract([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, float64(11122321312313), db.Float64)
	require.Equal(t, "        ", string(db.Rest))
}

func TestBasicBeforeCharStressDecoder(t *testing.T) {
	d := &DecodersStress{}
	ok, err := d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(d.Rest))
	ok, err = d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7abcdef rest`))
	require.NotNil(t, err)

	dl := &DecodersLimitedStress{}
	ok, err = dl.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(dl.Rest))

	db := &DecodersBoundedStress{}
	ok, err = db.Extract([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int64(11122321312313), db.Int64)
	require.Equal(t, "        ", string(db.Rest))
}

func TestBasicBeforeStringStressDecoder(t *testing.T) {
	d := &DecodersStringStress{}
	ok, err := d.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(d.Rest))

	dl := &DecodersLimitedStringStress{}
	ok, err = dl.Extract([]byte(`1 2 3 4 5 6 7 8 9 11e7 abcdef rest`))
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
	require.Equal(t, "rest", string(dl.Rest))

	db := &DecodersBoundedStringStress{}
	ok, err = db.Extract([]byte(`11122321312313         `))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, float64(11122321312313), db.Float64)
	require.Equal(t, "        ", string(db.Rest))
}

func TestDecoderOptionals(t *testing.T) {
	d := &DecoderOptionals{}

	ok, err := d.Extract([]byte(`12 head=13 end`))
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, ok)
	require.Equal(t, int8(12), d.Int8)
	require.Equal(t, "13", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`12 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), d.Int8)
	require.Equal(t, "", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`12 head=13`))
	require.False(t, ok)
	require.Nil(t, err)

	ds := &DecoderOptionalsStress{}

	ok, err = ds.Extract([]byte(`12 head=13 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), ds.Int8)
	require.Equal(t, "13", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`12 end`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int8(12), ds.Int8)
	require.Equal(t, "", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`12 head=13`))
	require.False(t, ok)
	require.Error(t, err)
}

func TestDecoderBranching(t *testing.T) {
	d := &DecoderBranching{}

	ok, err := d.Extract([]byte(`start head=data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`start head=data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`start data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`start data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))

	ok, err = d.Extract([]byte(`data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(d.GetHeadData()))
	ds := &DecoderBranchingStress{}
	ok, err = ds.Extract([]byte(`start head=data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`start head=data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`start data `))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`start data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))

	ok, err = ds.Extract([]byte(`data`))
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "data", string(ds.GetHeadData()))
}

func TestPrefixes(t *testing.T) {
	p := &Prefixes{}
	if ok, err := p.Extract([]byte("1234567890123 1234554321 ")); !ok {
		if err != nil {
			t.Fatal(err)
		} else {
			t.Fatalf("Should not be here")
		}
	}
	require.Equal(t, p.Data, int32(123))
	require.Equal(t, p.Rest1, int32(54321))
}

func TestFixedLook(t *testing.T) {
	p := &FixedLook{}
	if ok, err := p.Extract([]byte("9012345678901234")); !ok {
		if err != nil {
			t.Fatal(err)
		} else {
			t.Fatalf("Should not be here")
		}
	}
	require.Equal(t, p.Data, int32(90))
	require.Equal(t, p.Rest1, int32(12))
}

func TestAnonymousAreas(t *testing.T) {
	p := &AnonymousAreas{}
	src := "  data=1234end  "
	ok, err := p.Extract([]byte(src))
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "end", string(p.Data))

	src = "1234end"
	ok, err = p.Extract([]byte(src))
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "end", string(p.Data))

	src = "  data=end"
	ok, err = p.Extract([]byte(src))
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "end", string(p.Data))

	src = "end  "
	ok, err = p.Extract([]byte(src))
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "end", string(p.Data))
}

func TestSplit(t *testing.T) {
	p := &Split{}
	src := []byte("Name|1|2|3|4|5")
	if ok, err := p.Extract(src); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.NotNil(t, err)
	}

	require.Equal(t, "Name", string(p.Name))
	require.Equal(t, "4", string(p.Count))
}

func TestShift(t *testing.T) {
	p := &Shift1{}
	src := []byte("1233330000000000000000000000000000000000000000000000003")
	if ok, _ := p.Extract(src); ok {
		t.Errorf("Rule Shift1 must give a error on \033[1m%s\033[0m", string(src))
	}

	p2 := &Shift2{}
	if ok, _ := p2.Extract(src); ok {
		t.Errorf("Rule Shift2 must give a error on \033[1m%s\033[0m", string(src))
	}

	src = []byte("ba12ba              ")
	p3 := &Shift3{}
	if ok, err := p3.Extract(src); !ok {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("Rule Shift2 must give a error on \033[1m%s\033[0m", string(src))
	}
	require.Equal(t, "ba12", string(p3.B))

	p4 := &Shift4{}
	if ok, err := p4.Extract(src); !ok {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("Rule Shift2 must give a error on \033[1m%s\033[0m", string(src))
	}
	require.Equal(t, "ba12", string(p4.B))
}

func TestJump(t *testing.T) {
	p := &Jump{}
	src := []byte("1  34 15@@1@@@")
	if ok, err := p.Extract(src); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.NotNil(t, err)
	}
	require.Equal(t, "1 ", string(p.First))
	require.Equal(t, "34", string(p.Second))
	require.Equal(t, "15", string(p.Third))
	require.Equal(t, "1@", string(p.Fourth))
	require.Equal(t, "", string(p.Rest))
}

func TestLookupJump(t *testing.T) {
	p := &LookupJump{}
	src := []byte("1  3445@@123     ll123     ee123     e      f123")
	if ok, err := p.Extract(src); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.NotNil(t, err)
	}
	require.Equal(t, "123", string(p.Rest))
}

func TestTargetConstraintsCheck(t *testing.T) {
	p := &TargetConstraintsCheck{}
	src := []byte(" 1")

	ok, err := p.Extract(src)
	require.False(t, ok)
	if err != nil {
		t.Fatal(err)
	}

	src = []byte("1 1")
	ok, err = p.Extract(src)
	require.False(t, ok)
	require.NotNil(t, err)

	src = []byte("1 1 abcdef")
	ok, err = p.Extract(src)
	require.True(t, ok)
	require.Equal(t, "abcdef", string(p.Rest))
}

func TestIncludeChar(t *testing.T) {
	p := &IncludeChar{}
	src := []byte("abcd@12@")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "abcd@", string(p.Data))
	require.Equal(t, 12, p.Field2)
}

func TestIncludeString(t *testing.T) {
	p := &IncludeString{}
	src := []byte("adcdab12ab")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "adcdab", string(p.Data))
	require.Equal(t, 12, p.Field2)
}

func TestHex(t *testing.T) {
	p := &Hex{}
	src := []byte("fffe ff ffff ffffffff ffffffffffffffff")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, &Hex{
		Rest: []byte{},
		F1:   0xfffe,
		F2:   0xff,
		F3:   0xffff,
		F4:   0xffffffff,
		F5:   0xffffffffffffffff,
	}, p)
}

func TestOct(t *testing.T) {
	p := &Oct{}
	src := []byte("77 77 77 77 77")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, &Oct{
		Rest: []byte{},
		F1:   077,
		F2:   077,
		F3:   077,
		F4:   077,
		F5:   077,
	}, p)
}

func TestDec(t *testing.T) {
	p := &Dec{}
	src := []byte("5445.333 121212.22 512.22")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, &Dec{
		Rest: []byte{},
		F1:   5445333,
		F2:   12121222,
		F3: struct {
			Lo uint64
			Hi uint64
		}{Lo: 51222000000, Hi: 0},
	}, p)
}

func TestRestLegth(t *testing.T) {
	p := &RestLength{}
	src := bytes.Repeat([]byte{'#'}, 15)
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, strings.Repeat("#", 15), string(p.Rest))
}

func TestStr(t *testing.T) {
	p := &Str{}
	src := []byte("abc abcdef")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, &Str{
		Rest: []byte{},
		F1:   "abc",
		F2:   []byte("abcdef"),
	}, p)
}

func TestStar(t *testing.T) {
	var p Star
	src := []byte("aaaaaaaaaaaaaaaa123")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, Star{
		Rest: []byte{},
		F:    123,
	}, p)
}

func TestJustToCompile(t *testing.T) {
	var p JustToCompile
	src := []byte("aaaa-9999")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, JustToCompile{
		Rest: []byte(""),
		Head: 0xAAAA,
		Tail: 0x9999,
	}, p)

	ok, err = p.Extract([]byte("gggg-123"))
	require.False(t, ok)
	require.Error(t, err)
}

func TestJustToString(t *testing.T) {
	var p JustToCompileString
	src := []byte("aaaaabcd9999")
	ok, err := p.Extract(src)
	require.True(t, ok)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, JustToCompileString{
		Rest: []byte(""),
		Head: 0xAAAA,
		Tail: 0x9999,
	}, p)

	ok, err = p.Extract([]byte("gggg-123"))
	require.False(t, ok)
	if err != nil {
		t.Fatal(err)
	}
}
