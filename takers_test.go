package main

import "testing"

func TestTakeRest(t *testing.T) {
	data, err := compile(`rule = Data(int8);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = Data(uint16);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = Data(float32);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = Buf(string);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeString(t *testing.T) {
	data, err := compile(`rule = Date(int8) ' ';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ' ';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ' ');`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) ' '));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	panic(0)
}
