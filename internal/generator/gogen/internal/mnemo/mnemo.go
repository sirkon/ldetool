package mnemo

import (
	"bytes"
	"fmt"
	"strings"
)

// Maker takes care of correct rule (re)usage
type Maker struct {
	firstWrite    bool
	characters    map[string]rule
	prevCharacter string
	mnemonic      string
	buf           *bytes.Buffer
}

type rule struct {
	Singular string
	Plural   string
}

func newRule(sing, pl string) rule {
	return rule{
		Singular: sing,
		Plural:   pl,
	}
}

// New constructor
func New() *Maker {
	var specialCharacters = map[string]rule{
		" ":  newRule("space", "spaces"),
		";":  newRule("scolon", "scolons"),
		":":  newRule("colon", "colons"),
		".":  newRule("dot", "dots"),
		",":  newRule("comma", "commas"),
		"_":  newRule("uscore", "uscores"),
		"-":  newRule("minus", "minuses"),
		"+":  newRule("plus", "pluses"),
		"*":  newRule("star", "stars"),
		"/":  newRule("slash", "slashes"),
		"(":  newRule("lbrack", "lbracks"),
		")":  newRule("rbrack", "rbracks"),
		"[":  newRule("lsbrck", "lsbrcks"),
		"]":  newRule("rsbrck", "rsbrcks"),
		"{":  newRule("lbrace", "lbraces"),
		"}":  newRule("rbrace", "rbraces"),
		"&":  newRule("amp", "amps"),
		"%":  newRule("prcnt", "prcnts"),
		"$":  newRule("buck", "bucks"),
		"#":  newRule("hash", "hashes"),
		"@":  newRule("dog", "dogs"),
		"!":  newRule("xclam", "xclams"),
		"<":  newRule("less", "lesses"),
		">":  newRule("more", "mores"),
		"\"": newRule("qoute", "quotes"),
		"'":  newRule("apo", "apos"),
		"|":  newRule("bar", "bars"),
		"\\": newRule("bslash", "bslashes"),
		"^":  newRule("pow", "pows"),
		"=":  newRule("eq", "eqs"),
	}

	values := make(map[string]string, len(specialCharacters))
	for key, value := range specialCharacters {
		if value.Singular == value.Plural {
			panic(fmt.Errorf("Same singular and plural for `\033[1m%s\033[0m`", key))
		}
		if thatKey, ok := values[value.Singular]; ok {
			panic(fmt.Errorf("Same rule `\033[1m%s\033[0m` for characters \033[1m%s\033[0m and \033[1m%s\033[0m",
				value.Singular, key, thatKey))
		}
		values[value.Singular] = key
		if thatKey, ok := values[value.Plural]; ok {
			if thatKey == key {
				continue
			}
			panic(fmt.Errorf("Same rule `\033[1m%s\033[0m` for characters \033[1m%s\033[0m and \033[1m%s\033[0m",
				value.Plural, key, thatKey))
		}
		values[value.Plural] = key
	}
	return &Maker{
		firstWrite: true,
		characters: specialCharacters,
		buf:        &bytes.Buffer{},
	}
}

// Flush flushes collected characters
func (m *Maker) Flush() error {
	if len(m.mnemonic) > 0 {
		if m.firstWrite {
			_ = m.buf.WriteByte('_')
		}
		_, _ = m.buf.WriteString(m.mnemonic)
		_ = m.buf.WriteByte('_')
		m.mnemonic = ""
	}
	return nil
}

// WriteRune writes a rune
func (m *Maker) WriteRune(r rune) (n int, err error) {
	rr := string(r)
	if m.firstWrite && r == '"' || r == '\'' {
		return 1, nil
	}
	if mnemonic, ok := m.characters[rr]; ok {
		if rr == m.prevCharacter {
			m.mnemonic = mnemonic.Plural
		} else {
			_ = m.Flush()
			m.prevCharacter = rr
			m.mnemonic = mnemonic.Singular
		}
	} else {
		_ = m.Flush()
		m.prevCharacter = ""
		m.buf.WriteRune(r)
	}
	return len(rr), nil
}

// String ...
func (m *Maker) String() string {
	_ = m.Flush()
	res := m.buf.String()
	if strings.HasSuffix(res, "Quote") {
		return res[:len(res)-5]
	} else if strings.HasSuffix(res, "Apo") {
		return res[:len(res)-3]
	}
	if len(res) == 0 {
		return "constValue"
	}
	return res
}
