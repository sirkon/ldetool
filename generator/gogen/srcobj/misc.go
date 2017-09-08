package srcobj

import (
	"bytes"
	"fmt"
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
	fmt.Println(buf.String())
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
