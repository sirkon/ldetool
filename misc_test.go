package main

import "testing"

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
	panic(0)
}
