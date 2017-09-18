package main

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorListener representation of evil
type ErrorListener struct {
	line int
	col  int
	msg  string
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.line = line
	el.col = column
	el.msg = msg

	if strings.HasPrefix(msg, "no viable alternative") {
		el.msg = "syntax error"
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
