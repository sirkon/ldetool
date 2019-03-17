package ast

import (
	"fmt"
)

var _ Action = &PassAfter{}

// PassAfter ...
type PassAfter struct {
	access
	Limit *Target
}

func (pu *PassAfter) Accept(d ActionDispatcher) error {
	return d.DispatchPassAfter(pu)
}

func (pu *PassAfter) String() string {
	if pu.Limit.Lower == pu.Limit.Upper && pu.Limit.Lower > 0 {
		switch pu.Limit.Type {
		case String:
			return fmt.Sprintf("Check if the rest at %s character and further starts with prefix %s and pass the prefix then", posLit(pu.Limit.Lower+1), pu.Limit.Value)
		case Char:
			return fmt.Sprintf("Check if %s character equals to %s and pass it", posLit(pu.Limit.Lower+1), pu.Limit.Value)
		}
	}

	var area string
	if pu.Limit.Lower > 0 && pu.Limit.Upper > 0 {
		area = fmt.Sprintf("rest[%d:%d]", pu.Limit.Lower, pu.Limit.Upper)
	} else if pu.Limit.Lower > 0 {
		area = fmt.Sprintf("rest[%d:]", pu.Limit.Lower)
	} else {
		area = "the rest"
	}
	res := fmt.Sprintf("Look for \033[1m%s\033[0m in %s and pass it", pu.Limit.Value, area)
	return res
}

// PassAfterTarget ...
func PassAfterTarget() *PassAfter {
	return &PassAfter{
		Limit: NewTarget(),
	}
}
