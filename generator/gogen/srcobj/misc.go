package srcobj

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Found is `pos >= 0`
var Found = OperatorGE(Raw("pos"), Raw("0"))

var OK = Raw("ok")

// Goto generate goto <label>
func Goto(label string) Source {
	return Raw(fmt.Sprintf("goto %s", label))
}

func Error(format string, params ...Source) Source {
	replacements := map[byte]string{
		'"':  `\"`,
		'\r': `\r`,
		'\n': `\n`,
		'\t': `\t`,
		'\a': `\a`,
		'\\': `\`,
		0x1b: `\033`,
	}
	buf := &bytes.Buffer{}
	buf.Grow(len(format) * 7 / 6)
	for _, char := range []byte(format) {
		if replacement, ok := replacements[char]; ok {
			buf.WriteString(replacement)
		} else {
			buf.WriteByte(char)
		}
	}
	p := make([]Source, 0, len(params)+1)
	p = append(p, Raw(fmt.Sprintf(`"%s"`, buf.String())))
	p = append(p, params...)
	return NewCall("fmt.Errorf", p...)
}

// Comment generate comment
func Comment(comment string) Source {
	return Raw("// " + comment + "\n")
}

// Break represents break operator
const Break = Raw("break")

type hardToAccessDecodingStuff struct {
	Dest     string
	Decoding Call
	Failure  Source
}

func (d hardToAccessDecodingStuff) Dump(w io.Writer) error {
	expr := Assign(fmt.Sprintf("%s, err", d.Dest), d.Decoding)
	return If{
		Expr: OperatorSemicolon(expr, OperatorNEq(Raw("err"), Raw("nil"))),
		Then: d.Failure,
	}.Dump(w)
}

// Decode decoding generator
func Decode(dest string, decoding Call, fail Source) (res hardToAccessDecodingStuff) {
	res.Dest = dest
	res.Decoding = decoding
	res.Failure = fail
	return
}

// Trim trims generated code from new line characters and semicolons
func Trim(s Source) Source {
	data := String(s)
	data = strings.TrimRight(data, "\n;")
	return Raw(data)
}

// False ...
const False = Raw("false")

// Semicolon ...
const Semicolon = Raw(";")
