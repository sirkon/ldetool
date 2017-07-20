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
	data, err := compile(`rule = Date(int8) "border";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) "border";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) "border");`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) "border"));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeStringLimited(t *testing.T) {
	data, err := compile(`rule = Date(int8) "border"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) "border"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) "border"[:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) "border"[:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

}

func TestTakeBeforeStringBounded(t *testing.T) {
	data, err := compile(`rule = Date(int8) "border"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) "border"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(string) "border"[12:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) "border"[12:40])) _"end" Name(string) " not really";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeChar(t *testing.T) {
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

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ' '));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeLimitedChar(t *testing.T) {
	data, err := compile(`rule = Date(int8) ' '[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ' '[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ' '[:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ' '[:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeBoundedChar(t *testing.T) {
	data, err := compile(`rule = Date(int8) ' '[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ' '[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ' '[12:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ' '[12:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeStringOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? "border";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? "border";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? "border");`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) ?? "border"));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeLimitedStringOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? "border"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? "border"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? "border"[:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) ?? "border"[:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeBoundedStringOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? "border"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? "border"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? "border"[12:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(int64) ?? "border"[12:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeCharOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? ' ';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? ' ';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? ' ');`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ?? ' '));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeLimitedCharOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? ' '[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? ' '[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? ' '[:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ?? ' '[:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestTakeBeforeBoundedCharOrRest(t *testing.T) {
	data, err := compile(`rule = Date(int8) ?? ' '[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! Date(int8) ?? ' '[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  Data(int64) ?? ' '[12:40]);`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt(  ?Opt2 (Data(string) ?? ' '[12:40]));`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
