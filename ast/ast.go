package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// RuleItem ...
type RuleItem struct {
	Name    string
	Actions []*ActionItem

	NameToken antlr.Token
}

// NewRule constructor
func NewRule(name antlr.Token) *RuleItem {
	n := name.(antlr.Token)
	res := &RuleItem{
		Name:      string(n.GetText()),
		NameToken: n,
	}
	return res
}

func (r *RuleItem) Append(ai *ActionItem) {
	r.Actions = append(r.Actions, ai)
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
	ErrorOnMismatch      bool
}

func (ai *ActionItem) String() string {
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
	case ai.ErrorOnMismatch:
		return fmt.Sprintf("Treating all remaining mismatches in the rule as critical errors")

	default:
		panic("Must not happen!")
	}
}
