package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassFirstCharacters(t *testing.T) {
	data, err := compile(`rule = _[51:];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _[51:];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt( _[51:]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestEnd(t *testing.T) {
	data, err := compile(`rule = _"bugaga" $;`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _"bugaga" $;`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestRegression(t *testing.T) {
	data, err := compile(`
decoders =
	Int8(int8) ' '
	Int16(int16) ' '
	Int32(int32) ' '
	Int64(int64) ' '
	Uint8(uint8) ' '
	Uint16(uint16) ' '
	Uint32(uint32) ' '
	Uint64(uint64) ' '
	Float32(float32) ' '
	Float64(float64) ' '
	String(string);`)
	if err != nil {
		t.Fatal(err)
	}
	require.NotEqual(t, "", string(data))

	data, err = compile(`
decodersString =
	Int8(int8) " "
	Int16(int16) " "
	Int32(int32) " "
	Int64(int64) " "
	Uint8(uint8) " "
	Uint16(uint16) " "
	Uint32(uint32) " "
	Uint64(uint64) " "
	Float32(float32) " "
	Float64(float64) " "
	String(string) " ";`)
	if err != nil {
		t.Fatal(err)
	}
	require.NotEqual(t, "", string(data))
}
