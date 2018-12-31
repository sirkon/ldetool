package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

// ErrorListener representation of evil
type ErrorListener struct {
	Line int
	Col  int
	Msg  string
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Line = line
	el.Col = column
	el.Msg = msg

	if strings.HasPrefix(msg, "no viable alternative") {
		el.Msg = "syntax error"
	}
	panic(el)
}

func (el *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	return
}

func (el *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	return
}

func (el *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	return
}
