package gogen

import (
	"bytes"
	"fmt"
)

// Mnemowriter takes care of correct mnemonic (re)usage
type Mnemowriter struct {
	firstWrite    bool
	characters    map[string]mnemonic
	prevCharacter string
	mnemonic      string
	buf           *bytes.Buffer
}

type mnemonic struct {
	Singular string
	Plural   string
}

func newMnemonic(sing, pl string) mnemonic {
	return mnemonic{
		Singular: sing,
		Plural:   pl,
	}
}

// NewMnemowriter constructor
func NewMnemowriter() *Mnemowriter {
	var specialCharacters = map[string]mnemonic{
		";":  newMnemonic("scolon", "scolons"),
		":":  newMnemonic("colon", "colons"),
		".":  newMnemonic("dot", "dots"),
		",":  newMnemonic("comma", "commas"),
		"_":  newMnemonic("uscore", "uscores"),
		"-":  newMnemonic("minus", "minuses"),
		"+":  newMnemonic("plus", "pluses"),
		"*":  newMnemonic("star", "stars"),
		"/":  newMnemonic("slash", "slashes"),
		"(":  newMnemonic("lbrack", "lbracks"),
		")":  newMnemonic("rbrack", "rbracks"),
		"[":  newMnemonic("lsbrck", "lsbrcks"),
		"]":  newMnemonic("rsbrck", "rsbrcks"),
		"{":  newMnemonic("lbrace", "lbraces"),
		"}":  newMnemonic("rbrace", "rbraces"),
		"&":  newMnemonic("amp", "amps"),
		"%":  newMnemonic("prcnt", "prcnts"),
		"$":  newMnemonic("buck", "bucks"),
		"#":  newMnemonic("hash", "hashes"),
		"@":  newMnemonic("dog", "dogs"),
		"!":  newMnemonic("xclam", "xclams"),
		"<":  newMnemonic("less", "lesses"),
		">":  newMnemonic("more", "mores"),
		"\"": newMnemonic("qoute", "quotes"),
		"'":  newMnemonic("apo", "apos"),
		"|":  newMnemonic("bar", "bars"),
		"\\": newMnemonic("bslash", "bslashes"),
		"^":  newMnemonic("pow", "pows"),
		"=":  newMnemonic("eq", "eqs"),
	}

	values := make(map[string]string, len(specialCharacters))
	for key, value := range specialCharacters {
		if value.Singular == value.Plural {
			panic(fmt.Errorf("Same singular and plural for `\033[1m%s\033[0m`", key))
		}
		if thatKey, ok := values[value.Singular]; ok {
			panic(fmt.Errorf("Same mnemonic `\033[1m%s\033[0m` for characters \033[1m%s\033[0m and \033[1m%s\033[0m",
				value.Singular, key, thatKey))
		}
		values[value.Singular] = key
		if thatKey, ok := values[value.Plural]; ok {
			if thatKey == key {
				continue
			}
			panic(fmt.Errorf("Same mnemonic `\033[1m%s\033[0m` for characters \033[1m%s\033[0m and \033[1m%s\033[0m",
				value.Plural, key, thatKey))
		}
		values[value.Plural] = key
	}
	return &Mnemowriter{
		firstWrite: true,
		characters: specialCharacters,
		buf:        &bytes.Buffer{},
	}
}

// Flush flushes collected characters
func (m *Mnemowriter) Flush() error {
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
func (m *Mnemowriter) WriteRune(r rune) (n int, err error) {
	rr := string(r)
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
func (m *Mnemowriter) String() string {
	_ = m.Flush()
	return m.buf.String()
}
