package ast

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
)

// TokenError printer
func TokenError(t antlr.Token) func(format string, args ...interface{}) error {
	return func(format string, args ...interface{}) error {
		prefix := fmt.Sprintf("%d:column %d: ", t.GetLine(), t.GetColumn())
		return errors.New(prefix + fmt.Sprintf(format, args...))
	}

}

func posLit(i int) string {
	switch i {
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	default:
		return fmt.Sprintf("%dth", i)
	}
}

const decimalPrefix = "dec"

// decimalExtractor ...
type decimalExtractor struct {
	Rest      string
	Precision uint
	Scale     uint
}

// Extract ...
func (p *decimalExtractor) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var tmp string
	var tmpUint uint64

	// Checks if the rest starts with `"decimalPrefix"` and pass it
	if strings.HasPrefix(p.Rest, decimalPrefix) {
		p.Rest = p.Rest[len(decimalPrefix):]
	} else {
		return false, nil
	}

	// Take until '_' as Precision(uint)
	pos = strings.IndexByte(p.Rest, '_')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpUint, err = strconv.ParseUint(tmp, 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Precision = uint(tmpUint)

	// Take the rest as Scale(uint)
	if tmpUint, err = strconv.ParseUint(p.Rest, 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(p.Rest), err)
	}
	p.Scale = uint(tmpUint)
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}
