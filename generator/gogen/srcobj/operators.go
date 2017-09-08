package srcobj

import (
	"io"
)

type hardToAccessBinaryOperator struct {
	operator string
	operand1 Source
	operand2 Source
}

func (h hardToAccessBinaryOperator) Dump(w io.Writer) error {
	if err := h.operand1.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, h.operator); err != nil {
		return err
	}
	if err := h.operand2.Dump(w); err != nil {
		return err
	}
	return nil
}

// OperatorGE generate Greater or Equal comparison
func OperatorGE(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: ">=",
	}
}

// OperatorGT generate Greater or Equal comparison
func OperatorGT(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: ">",
	}
}

// OperatorLT generate Greater or Equal comparison
func OperatorLT(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "<",
	}
}

// OperatorEq generate equality check
func OperatorEq(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "==",
	}
}

// OperatorNEq generate equality check
func OperatorNEq(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "!=",
	}
}

// OperatorAnd generate equality check
func OperatorAnd(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "&&",
	}
}

func OperatorColon(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: ":",
	}
}

func OperatorAdd(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "+",
	}
}
