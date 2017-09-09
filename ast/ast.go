package ast

import (
	"fmt"

	"github.com/sirkon/ldetool/token"
)

// Attrib ...
type Attrib interface{}

// RuleSeq ...
func RuleSeq(cur Attrib, next Attrib) (attr Attrib, err error) {
	res := cur.([]RuleItem)
	item := next.(RuleItem)
	res = append(res, item)
	return res, nil
}

// LastItem ...
func LastItem(cur Attrib) (attr Attrib, err error) {
	item := cur.(RuleItem)
	res := []RuleItem{item}
	return res, nil
}

// RuleItem ...
type RuleItem struct {
	Name    string
	Actions ActionSequence

	NameToken *token.Token
}

// Rule ...
func Rule(name Attrib, act Attrib) (attr Attrib, err error) {
	n := name.(*token.Token)
	res := RuleItem{
		Name:    string(n.Lit),
		Actions: act.(ActionSequence),

		NameToken: n,
	}
	return res, nil
}

func fasten(as1, as2 ActionSequence) ActionSequence {
	if as1.Tail != nil {
		res := fasten(*as1.Tail, as2)
		return ActionSequence{Head: as1.Head, Tail: &res}
	}
	return ActionSequence{Head: as1.Head, Tail: &as2}
}

func assignAction(act Attrib) (res ActionItem, aseq bool, err error) {
	switch t := act.(type) {
	case AtEnd:
		res.End = &t
	case Optional:
		res.Option = &t
	case PassUntil:
		res.Pass = &t
	case PassUntilOrIgnore:
		res.PassOrIgnore = &t
	case StartChar:
		res.StartWithChar = &t
	case StartString:
		res.StartWithString = &t
	case MayBeStartChar:
		res.MayBeStartWithChar = &t
	case MayBeStartString:
		res.MayBeStartWithString = &t
	case Take:
		res.Take = &t
	case TakeRest:
		res.TakeRest = &t
	case TakeUntilOrRest:
		res.TakeUntilOrRest = &t
	case PassFixed:
		res.PassFirst = &t
	case ActionSequence:
		aseq = true
	default:
		err = fmt.Errorf("Unsupported action object %T", act)
	}
	return
}

// ActionSeq ...
func ActionSeq(act Attrib, next Attrib) (attr Attrib, err error) {
	tail := next.(ActionSequence)
	res, aseq, err := assignAction(act)
	if aseq {
		return fasten(act.(ActionSequence), tail), nil
	}
	if err != nil {
		return
	}
	return ActionSequence{Head: res, Tail: &tail}, err

}

// Action ...
func Action(act Attrib) (attr Attrib, err error) {
	res, aseq, err := assignAction(act)
	if aseq {
		err = fmt.Errorf("Composed actions (ActionSequence) is not expected at this step")
	}
	return ActionSequence{Head: res, Tail: nil}, err
}

// MatchRequired ...
func MatchRequired(act Attrib) (attr Attrib, err error) {
	res := act.(ActionSequence)
	res.ErrorOnMismatch = true
	return res, nil
}

// ActionItem ...
type ActionItem struct {
	End                  *AtEnd
	Option               *Optional
	Pass                 *PassUntil
	PassOrIgnore         *PassUntilOrIgnore
	StartWithChar        *StartChar
	StartWithString      *StartString
	MayBeStartWithChar   *MayBeStartChar
	MayBeStartWithString *MayBeStartString
	Take                 *Take
	TakeRest             *TakeRest
	TakeUntilOrRest      *TakeUntilOrRest
	PassFirst            *PassFixed
}

func (ai ActionItem) String() string {
	switch {
	case ai.End != nil:
		return "AtEnd"
	case ai.Option != nil:
		return fmt.Sprintf("Option %s", ai.Option.Name)
	case ai.Pass != nil:
		return fmt.Sprintf("Pass until \033[1m%s\033[0m", ai.Pass.Limit.Value)
	case ai.PassOrIgnore != nil:
		return fmt.Sprintf("Pass until \033[1m%s\033[0m or ignore", ai.PassOrIgnore.Limit.Value)
	case ai.StartWithChar != nil:
		return fmt.Sprintf("Check and pass character \033[1m%s\033[0m", ai.StartWithChar.Value)
	case ai.StartWithString != nil:
		return fmt.Sprintf("Check and pass \033[1m%s\033[0m", ai.StartWithString.Value)
	case ai.MayBeStartWithChar != nil:
		return fmt.Sprintf("Pass character \033[1m%s\033[0m if starts with", ai.MayBeStartWithChar.Value)
	case ai.MayBeStartWithString != nil:
		return fmt.Sprintf("Pass \033[1m%s\033[0m if starts with", ai.MayBeStartWithString.Value)
	case ai.Take != nil:
		return fmt.Sprintf("Take until \033[1m%s\033[0m as \033[32m%s(%s)\033[0m",
			ai.Take.Limit.Value, ai.Take.Field.Name, ai.Take.Field.Type)
	case ai.TakeRest != nil:
		return fmt.Sprintf("Take the rest as \033[32m%s(%s)\033[0m", ai.TakeRest.Field.Name, ai.TakeRest.Field.Type)
	case ai.TakeUntilOrRest != nil:
		return fmt.Sprintf(
			"Take until %s (or all the rest if not found) as \033[32m%s(%s)\033[0m",
			ai.TakeUntilOrRest.Limit.Value, ai.TakeUntilOrRest.Field.Name, ai.TakeUntilOrRest.Field.Type)
	case ai.PassFirst != nil:
		return fmt.Sprintf("Passing first %d letters", ai.PassFirst)

	default:
		panic("Must not happen!")
	}
}

// ActionSequence holds sequence of actions
type ActionSequence struct {
	ErrorOnMismatch bool
	Head            ActionItem
	Tail            *ActionSequence
}

func (as *ActionSequence) String() string {
	res := as.Head.String()
	if as.Tail != nil {
		return res + ": " + as.Tail.String()
	}
	return res
}
