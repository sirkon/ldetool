package ast

import (
	"fmt"
)

type RestLengthCheck struct {
	access
	Operator string
	Length   int
}

func (a RestLengthCheck) String() string {
	switch a.Operator {
	case "<":
		return fmt.Sprintf("Check if less than %d symbols left in the rest", a.Length)
	case "==":
		return fmt.Sprintf("Check if exactly %d symbols left in the rest", a.Length)
	case ">":
		return fmt.Sprintf("Check if more than %d symbols left in the rest", a.Length)
	default:
		panic(fmt.Errorf("unsupported operator %s in rest length check", a.Operator))
	}
}

func (a RestLengthCheck) Accept(d ActionDispatcher) error {
	return d.DispatchRestLengthCheck(a)
}

func RestCheck(operator string, length int) RestLengthCheck {
	if len(operator) == 0 {
		operator = "=="
	}
	return RestLengthCheck{
		Operator: operator,
		Length:   length,
	}
}
