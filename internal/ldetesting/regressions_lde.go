// Code generated by ldetool --package ldetesting --go-string regressions.lde. DO NOT EDIT.

package ldetesting

import (
	"fmt"
	"strconv"
	"strings"
)

var constFooBarBazLessBar = "<Bar"
var constFooBarBazLessBazMore = "<baz>"
var constFooBarBazLessFooMore = "<foo>"
var constFooBarBazLessSlashBazMore = "</baz>"
var constFooBarBazLessSlashFooMore = "</foo>"
var constFooBarBazRbrace = "}'"
var constFooBarBazSlashMore = "/>"
var constFooBarBazSpaceFoobarEqLbrace = " foobar='{"

// Regression1 ...
type Regression1 struct {
	Rest  string
	Pid   int32
	Comm  string
	State uint8
	Ppid  int32
}

// Extract ...
func (p *Regression1) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var tmp string
	var tmpInt int64

	// Take until ' ' as Pid(int32)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 32); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Pid(int32): %s", tmp, err)
	}
	p.Pid = int32(tmpInt)

	// Take until ' ' as Comm(string)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		p.Comm = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until 2nd character  if it is equal to ' ' character as State($uint8)
	if len(p.Rest) < 1+1 || p.Rest[1] != ' ' {
		return false, nil
	}
	tmp = p.Rest[:1]
	p.Rest = p.Rest[1+1:]
	if p.State, err = p.unmarshalState(tmp); err != nil {
		return false, fmt.Errorf("parsing `%s` into field State(uint8): %s", tmp, err)
	}

	// Take until ' ' as Ppid(int32)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 32); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Ppid(int32): %s", tmp, err)
	}
	p.Ppid = int32(tmpInt)

	return true, nil
}

// FooBarBaz relates to https://github.com/sirkon/ldetool/issues/40
type FooBarBaz struct {
	Rest  string
	Stuff string
	Bar   struct {
		Valid bool
		ID    struct {
			Valid     bool
			Foobarbaz string
		}
	}
	Baz string
}

// Extract ...
func (p *FooBarBaz) Extract(line string) (bool, error) {
	p.Rest = line
	var barIDRest string
	var barRest string
	var pos int

	// Checks if the rest starts with `"<foo>"` and pass it
	if strings.HasPrefix(p.Rest, constFooBarBazLessFooMore) {
		p.Rest = p.Rest[len(constFooBarBazLessFooMore):]
	} else {
		return false, nil
	}

	// Take until "</foo>" as Stuff(string)
	pos = strings.Index(p.Rest, constFooBarBazLessSlashFooMore)
	if pos >= 0 {
		p.Stuff = p.Rest[:pos]
		p.Rest = p.Rest[pos+len(constFooBarBazLessSlashFooMore):]
	} else {
		return false, nil
	}
	barRest = p.Rest

	// Checks if the rest starts with `"<Bar"` and pass it
	if strings.HasPrefix(barRest, constFooBarBazLessBar) {
		barRest = barRest[len(constFooBarBazLessBar):]
	} else {
		p.Bar.Valid = false
		goto foobarbazBarLabel
	}
	barIDRest = barRest

	// Checks if the rest starts with `" foobar='{"` and pass it
	if strings.HasPrefix(barIDRest, constFooBarBazSpaceFoobarEqLbrace) {
		barIDRest = barIDRest[len(constFooBarBazSpaceFoobarEqLbrace):]
	} else {
		p.Bar.ID.Valid = false
		goto foobarbazBarIDLabel
	}

	// Take until "}'" as Foobarbaz(string)
	pos = strings.Index(barIDRest, constFooBarBazRbrace)
	if pos >= 0 {
		p.Bar.ID.Foobarbaz = barIDRest[:pos]
		barIDRest = barIDRest[pos+len(constFooBarBazRbrace):]
	} else {
		p.Bar.ID.Valid = false
		goto foobarbazBarIDLabel
	}

	p.Bar.ID.Valid = true
	barRest = barIDRest
foobarbazBarIDLabel:

	// Checks if the rest starts with `"/>"` and pass it
	if strings.HasPrefix(barRest, constFooBarBazSlashMore) {
		barRest = barRest[len(constFooBarBazSlashMore):]
	} else {
		p.Bar.Valid = false
		goto foobarbazBarLabel
	}

	p.Bar.Valid = true
	p.Rest = barRest
foobarbazBarLabel:

	// Checks if the rest starts with `"<baz>"` and pass it
	if strings.HasPrefix(p.Rest, constFooBarBazLessBazMore) {
		p.Rest = p.Rest[len(constFooBarBazLessBazMore):]
	} else {
		return false, nil
	}

	// Take until "</baz>" as Baz(string)
	pos = strings.Index(p.Rest, constFooBarBazLessSlashBazMore)
	if pos >= 0 {
		p.Baz = p.Rest[:pos]
		p.Rest = p.Rest[pos+len(constFooBarBazLessSlashBazMore):]
	} else {
		return false, nil
	}

	return true, nil
}

// GetBarIDFoobarbaz ...
func (p *FooBarBaz) GetBarIDFoobarbaz() (res string) {
	if p.Bar.Valid {
		if p.Bar.ID.Valid {
			res = p.Bar.ID.Foobarbaz
		}
	}
	return
}
