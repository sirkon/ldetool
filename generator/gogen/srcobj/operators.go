package srcobj

import (
	"io"
	"strings"
)

type hardToAccessBinaryOperator struct {
	operator string
	operand1 Source
	operand2 Source
}

func (h hardToAccessBinaryOperator) Dump(w io.Writer) error {
	data := String(h.operand1)
	data = strings.TrimRight(data, "\n")
	if _, err := io.WriteString(w, data); err != nil {
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

func OperatorDot(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: ".",
	}
}

func OperatorInstruction(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: ";",
	}
}

func OperatorInc(op1, op2 Source) Source {
	return hardToAccessBinaryOperator{
		operand1: op1,
		operand2: op2,
		operator: "+=",
	}
}

type hardToAccessUnaryOperator struct {
	operator string
	operand  Source
}

func (uo hardToAccessUnaryOperator) Dump(w io.Writer) error {
	if _, err := io.WriteString(w, uo.operator); err != nil {
		return err
	}
	if err := uo.operand.Dump(w); err != nil {
		return err
	}
	return nil
}

func unaryAccess(operator string, operand Source) Source {
	return hardToAccessUnaryOperator{
		operator: operator,
		operand:  operand,
	}
}

func Reference(operand Source) Source {
	return unaryAccess("&", operand)
}
