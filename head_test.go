package main

import "testing"

func TestHeadGeneration(t *testing.T) {
	data, err := compile(`rule = ^"head";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! ^"head";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ^"head" );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ^'[';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! ^'[';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	panic(0)
}

func TestMayBeHeadGeneration(t *testing.T) {
	data, err := compile(`rule = ^ ?? "head";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! ^ ?? "head";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ^ ?? "head" );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ^ ?? '[';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! ^ ??'[';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ^ ?? '[' );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

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

	panic(0)
}
