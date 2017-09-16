
/*
 This file was autogenerated via
 --------------------------------------------------------
 ldetool generate --package main msg_translator_rules.lde
 --------------------------------------------------------
 do not touch it with bare hands!
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

var commaSpaceExpectedSpaceOneSpaceOfColonSpace = []byte(", expected one of: ")
var identifierLbrack = []byte("identifier(")
var lineEq = []byte("line=")
var posLbrack = []byte("Pos(")

// TypeError ...
type TypeError struct {
	rest    []byte
	Name    []byte
	Line    uint32
	Column  uint32
	Choices []byte
}

// Extract ...
func (p *TypeError) Extract(line []byte) (bool, error) {
	p.rest = line
	var err error
	var pos int
	var tmp []byte
	var tmpUint uint64

	// Checks if the rest starts with `"Error in"` and pass it
	if len(p.rest) >= 8 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&18446744073709551615 == 7955925892594823749 {
		p.rest = p.rest[8:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), "Error in")
	}

	// Looking for ':' and then pass it
	pos = bytes.IndexByte(p.rest[:8], ':')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", ':', string(p.rest[:8]))
	}

	// Looking for "identifier(" and then pass it
	pos = bytes.Index(p.rest, identifierLbrack)
	if pos >= 0 {
		p.rest = p.rest[pos+len(identifierLbrack):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", identifierLbrack, string(p.rest))
	}

	// Looking for ',' and then pass it
	pos = bytes.IndexByte(p.rest, ',')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", ',', string(p.rest))
	}

	// Take until ')' as Name(string)
	pos = bytes.IndexByte(p.rest, ')')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Name", ')', string(p.rest))
	}
	p.Name = tmp

	// Looking for "Pos(" and then pass it
	pos = bytes.Index(p.rest, posLbrack)
	if pos >= 0 {
		p.rest = p.rest[pos+len(posLbrack):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", posLbrack, string(p.rest))
	}

	// Looking for "line=" and then pass it
	pos = bytes.Index(p.rest, lineEq)
	if pos >= 0 {
		p.rest = p.rest[pos+len(lineEq):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", lineEq, string(p.rest))
	}

	// Take until ',' as Line(uint32)
	pos = bytes.IndexByte(p.rest, ',')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Line", ',', string(p.rest))
	}
	if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 32); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Line = uint32(tmpUint)

	// Checks if the rest starts with `" column="` and pass it
	if len(p.rest) >= 8 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&18446744073709551615 == 4426595834849616672 {
		p.rest = p.rest[8:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " column=")
	}

	// Take until ')' as Column(uint32)
	pos = bytes.IndexByte(p.rest, ')')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Column", ')', string(p.rest))
	}
	if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 32); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Column = uint32(tmpUint)

	// Checks if the rest starts with `", expected one of: "` and pass it
	if bytes.HasPrefix(p.rest, commaSpaceExpectedSpaceOneSpaceOfColonSpace) {
		p.rest = p.rest[len(commaSpaceExpectedSpaceOneSpaceOfColonSpace):]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), ", expected one of: ")
	}

	// Take the rest as Choices(string)
	p.Choices = p.rest
	p.rest = p.rest[len(p.rest):]
	return true, nil
}