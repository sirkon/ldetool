package ast

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
func StartsWithString(target antlr.Token) (*StartString, error) {
	return &StartString{Value: target.GetText()}, nil
}

// StartsWithChar ...
func StartsWithChar(target antlr.Token) (*StartChar, error) {
	return &StartChar{Value: target.GetText()}, nil
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
func MayBeStartsWithString(value antlr.Token) (*MayBeStartString, error) {
	return &MayBeStartString{Value: value.GetText()}, nil
}

// MayBeStartsWithChar ...
func MayBeStartsWithChar(value antlr.Token) (*MayBeStartChar, error) {
	return &MayBeStartChar{Value: value.GetText()}, nil
}

// PassUntil ...
type PassUntil struct {
	Limit *Target
}

// PassUntilTarget ...
func PassUntilTarget() (res *PassUntil, err error) {
	res = &PassUntil{
		Limit: NewTarget(),
	}
	return res, nil
}

// PassUntilOrIgnore ...
type PassUntilOrIgnore struct {
	Limit *Target
}

// PassUntilTargetOrIgnore ...
func PassUntilTargetOrIgnore() (res *PassUntilOrIgnore, err error) {
	res = &PassUntilOrIgnore{
		Limit: NewTarget(),
	}
	return res, nil
}

// Field ...
type Field struct {
	Name string
	Type string

	NameToken antlr.Token
	TypeToken antlr.Token
}

// NewField constructor
func NewField(name antlr.Token, typeToken antlr.Token) Field {
	return Field{
		Name: string(name.GetText()),
		Type: string(typeToken.GetText()),

		NameToken: name,
		TypeToken: typeToken,
	}
}

// Take ...
type Take struct {
	Field Field
	Limit *Target
}

// TakeUntilTarget ...
func TakeUntilTarget(field, fieldType antlr.Token) (res *Take, err error) {
	f := NewField(field, fieldType)
	res = &Take{
		Field: f,
		Limit: NewTarget(),
	}
	return
}

// TakeRest ...
type TakeRest struct {
	Field Field
}

// TakeTheRest ...
func TakeTheRest(field, fieldType antlr.Token) (*TakeRest, error) {
	f := NewField(field, fieldType)
	res := &TakeRest{
		Field: f,
	}
	return res, nil
}

// AtEnd ...
type AtEnd struct{}

// AtTheEnd ...
func AtTheEnd() (*AtEnd, error) {
	return &AtEnd{}, nil
}

// Optional ...
type Optional struct {
	Name    string
	Actions []*ActionItem

	NameToken antlr.Token
}

// Option ...
func Option(opt antlr.Token) (*Optional, error) {
	res := &Optional{
		Name:    opt.GetText(),
		Actions: []*ActionItem{},

		NameToken: opt,
	}
	return res, nil
}

func (o *Optional) Append(i *ActionItem) {
	o.Actions = append(o.Actions, i)
}

// AnonymousOption ...
type AnonymousOption struct {
	Name    string
	Actions []*ActionItem

	StartToken antlr.Token
}

// Anonymous ...
func Anonymous(opt antlr.Token) (*AnonymousOption, error) {
	res := &AnonymousOption{
		Actions: []*ActionItem{},

		StartToken: opt,
	}
	return res, nil
}

func (a *AnonymousOption) Append(i *ActionItem) {
	a.Actions = append(a.Actions, i)
}

// TakeUntilOrRest ...
type TakeUntilOrRest struct {
	Field Field
	Limit *Target
}

// TakeUntilTargetOrRest ...
func TakeUntilTargetOrRest(field antlr.Token, fieldType antlr.Token) (res *TakeUntilOrRest, err error) {
	f := NewField(field, fieldType)
	res = &TakeUntilOrRest{
		Field: f,
		Limit: NewTarget(),
	}
	return res, nil
}

// PassFixed ...
type PassFixed int

// PassFirst ...
func PassFirst(field antlr.Token) (res *PassFixed, err error) {
	f := field.(antlr.Token)
	value, err := strconv.ParseInt(string(f.GetText()), 10, 64)
	if err != nil {
		return
	}
	result := PassFixed(value)
	return &result, nil
}
