package ast

import (
	"strconv"

	"github.com/sirkon/ldetool/token"
)

// Target ...
type Target struct {
	Type  TargetEnum
	Value string
	Lower int
	Upper int
	Close bool
}

// BoundedScopeStringTarget ...
func BoundedScopeStringTarget(needle Attrib, lower Attrib, upper Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := lower.(*token.Token)
	ll, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil || ll == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	u := upper.(*token.Token)
	uu, err := strconv.ParseUint(string(u.Lit), 10, 64)
	if err != nil || uu == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	if ll >= uu {
		err = TokenError(l)("lower bound must be lower than upper")
	}
	res := Target{
		Type:  String,
		Value: string(n.Lit),
		Lower: int(ll),
		Upper: int(uu),
	}
	return res, nil
}

// LimitedScopeStringTarget ...
func LimitedScopeStringTarget(needle Attrib, limitation Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := limitation.(*token.Token)
	limit, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil || limit == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	res := Target{
		Type:  String,
		Value: string(n.Lit),
		Upper: int(limit),
	}
	return res, nil
}

// FixedStringTarget ...
func FixedStringTarget(needle Attrib, limitation Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := limitation.(*token.Token)
	limit, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil || limit == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	res := Target{
		Type:  String,
		Value: string(n.Lit),
		Lower: int(limit),
		Upper: int(limit),
	}
	return res, nil
}

// CloseStringTarget ...
func CloseStringTarget(needle Attrib) (attr Attrib, err error) {
	res := Target{
		Type:  String,
		Value: string(needle.(*token.Token).Lit),
		Close: true,
	}
	return res, nil
}

// StringTarget ...
func StringTarget(needle Attrib) (attr Attrib, err error) {
	res := Target{
		Type:  String,
		Value: string(needle.(*token.Token).Lit),
	}
	return res, nil
}

// BoundedScopeCharTarget ...
func BoundedScopeCharTarget(needle Attrib, lower, upper Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := lower.(*token.Token)
	ll, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	u := upper.(*token.Token)
	uu, err := strconv.ParseUint(string(u.Lit), 10, 64)
	if err != nil || uu == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	if ll >= uu {
		err = TokenError(l)("lower bound must be lower than upper")
	}
	res := Target{
		Type:  Char,
		Value: string(n.Lit),
		Lower: int(ll),
		Upper: int(uu),
	}
	return res, nil
}

// LimitedScopeCharTarget ...
func LimitedScopeCharTarget(needle Attrib, limitation Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := limitation.(*token.Token)
	limit, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	res := Target{
		Type:  Char,
		Value: string(n.Lit),
		Upper: int(limit),
	}
	return res, nil
}

// FixedCharTarget ...
func FixedCharTarget(needle Attrib, limitation Attrib) (attr Attrib, err error) {
	n := needle.(*token.Token)
	l := limitation.(*token.Token)
	limit, err := strconv.ParseUint(string(l.Lit), 10, 64)
	if err != nil || limit == 0 {
		err = TokenError(l)("bad number %s", string(l.Lit))
		return
	}
	res := Target{
		Type:  Char,
		Value: string(n.Lit),
		Lower: int(limit),
		Upper: int(limit),
	}
	return res, nil
}

// CloseCharTarget ...
func CloseCharTarget(needle Attrib) (attr Attrib, err error) {
	res := Target{
		Type:  Char,
		Value: string(needle.(*token.Token).Lit),
		Close: true,
	}
	return res, nil
}

// CharTarget ...
func CharTarget(needle Attrib) (attr Attrib, err error) {
	res := Target{
		Type:  Char,
		Value: string(needle.(*token.Token).Lit),
		Upper: 0,
	}
	return res, nil
}
