package ast

import (
	"strconv"

	"github.com/DenisCheremisov/ldegen/token"
)

// StartString ...
type StartString struct {
	Value string
}

// StartChar ...
type StartChar struct {
	Value string
}

// StartsWithString ...
func StartsWithString(target Attrib) (Attrib, error) {
	return StartString{Value: string(target.(*token.Token).Lit)}, nil
}

// StartsWithChar ...
func StartsWithChar(target Attrib) (Attrib, error) {
	return StartChar{Value: string(target.(*token.Token).Lit)}, nil
}

// MayBeStartString ...
type MayBeStartString struct {
	Value string
}

// MayBeStartChar ...
type MayBeStartChar struct {
	Value string
}

// MayBeStartsWithString ...
func MayBeStartsWithString(target Attrib) (Attrib, error) {
	return MayBeStartString{Value: string(target.(*token.Token).Lit)}, nil
}

// MayBeStartsWithChar ...
func MayBeStartsWithChar(target Attrib) (Attrib, error) {
	return MayBeStartChar{Value: string(target.(*token.Token).Lit)}, nil
}

// PassUntil ...
type PassUntil struct {
	Limit Target
}

// PassUntilTarget ...
func PassUntilTarget(target Attrib) (attr Attrib, err error) {
	res := PassUntil{
		Limit: target.(Target),
	}
	return res, nil
}

// Field ...
type Field struct {
	Name string
	Type string

	NameToken *token.Token
	TypeToken *token.Token
}

// NewField constructor
func NewField(name *token.Token, typeToken *token.Token) Field {
	return Field{
		Name: string(name.Lit),
		Type: string(typeToken.Lit),

		NameToken: name,
		TypeToken: typeToken,
	}
}

// Take ...
type Take struct {
	Field Field
	Limit Target
}

// TakeUntilTarget ...
func TakeUntilTarget(field Attrib, fieldType Attrib, target Attrib) (attr Attrib, err error) {
	f := NewField(field.(*token.Token), fieldType.(*token.Token))
	res := Take{
		Field: f,
		Limit: target.(Target),
	}
	return res, nil
}

// TakeRest ...
type TakeRest struct {
	Field Field
}

// TakeTheRest ...
func TakeTheRest(field Attrib, fieldType Attrib) (attr Attrib, err error) {
	f := NewField(field.(*token.Token), fieldType.(*token.Token))
	res := TakeRest{
		Field: f,
	}
	return res, nil
}

// AtEnd ...
type AtEnd struct{}

// AtTheEnd ...
func AtTheEnd() (attr Attrib, err error) {
	return AtEnd{}, nil
}

// Optional ...
type Optional struct {
	Name    string
	Actions ActionSequence

	NameToken *token.Token
}

// Option ...
func Option(opt Attrib, act Attrib) (attr Attrib, err error) {
	optName := opt.(*token.Token)
	res := Optional{
		Name:    string(optName.Lit),
		Actions: act.(ActionSequence),

		NameToken: optName,
	}
	return res, nil
}

// TakeUntilOrRest ...
type TakeUntilOrRest struct {
	Field Field
	Limit Target
}

// TakeUntilTargetOrRest ...
func TakeUntilTargetOrRest(field Attrib, fieldType Attrib, target Attrib) (attr Attrib, err error) {
	f := NewField(field.(*token.Token), fieldType.(*token.Token))
	res := Take{
		Field: f,
		Limit: target.(Target),
	}
	return res, nil
}

// PassFixed ...
type PassFixed int

// PassFirst ...
func PassFirst(field Attrib) (attr Attrib, err error) {
	f := field.(*token.Token)
	value, err := strconv.ParseInt(string(f.Lit), 10, 64)
	if err != nil {
		return field, err
	}
	return PassFixed(value), nil
}
