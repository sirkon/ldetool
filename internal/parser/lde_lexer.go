// Code generated from LDE.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 122,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 7, 14, 72, 10, 14, 12, 14, 14, 14, 75, 11, 14, 3, 15, 6,
	15, 78, 10, 15, 13, 15, 14, 15, 79, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 7, 17, 88, 10, 17, 12, 17, 14, 17, 91, 11, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 101, 10, 19, 12, 19, 14,
	19, 104, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	7, 21, 114, 10, 21, 12, 21, 14, 21, 117, 11, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 4, 89, 102, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 2, 33, 17, 35,
	2, 37, 18, 39, 19, 41, 20, 43, 21, 3, 2, 8, 5, 2, 67, 92, 97, 97, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 12, 15,
	15, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 126, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 3, 45, 3, 2, 2,
	2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 51, 3, 2, 2, 2, 11, 53, 3,
	2, 2, 2, 13, 55, 3, 2, 2, 2, 15, 57, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19,
	61, 3, 2, 2, 2, 21, 63, 3, 2, 2, 2, 23, 65, 3, 2, 2, 2, 25, 67, 3, 2, 2,
	2, 27, 69, 3, 2, 2, 2, 29, 77, 3, 2, 2, 2, 31, 81, 3, 2, 2, 2, 33, 84,
	3, 2, 2, 2, 35, 94, 3, 2, 2, 2, 37, 97, 3, 2, 2, 2, 39, 107, 3, 2, 2, 2,
	41, 111, 3, 2, 2, 2, 43, 120, 3, 2, 2, 2, 45, 46, 7, 63, 2, 2, 46, 4, 3,
	2, 2, 2, 47, 48, 7, 61, 2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7, 42, 2, 2, 50,
	8, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 10, 3, 2, 2, 2, 53, 54, 7, 96,
	2, 2, 54, 12, 3, 2, 2, 2, 55, 56, 7, 93, 2, 2, 56, 14, 3, 2, 2, 2, 57,
	58, 7, 95, 2, 2, 58, 16, 3, 2, 2, 2, 59, 60, 7, 65, 2, 2, 60, 18, 3, 2,
	2, 2, 61, 62, 7, 97, 2, 2, 62, 20, 3, 2, 2, 2, 63, 64, 7, 60, 2, 2, 64,
	22, 3, 2, 2, 2, 65, 66, 7, 38, 2, 2, 66, 24, 3, 2, 2, 2, 67, 68, 7, 128,
	2, 2, 68, 26, 3, 2, 2, 2, 69, 73, 9, 2, 2, 2, 70, 72, 9, 3, 2, 2, 71, 70,
	3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 28, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 9, 4, 2, 2, 77, 76, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80,
	30, 3, 2, 2, 2, 81, 82, 7, 94, 2, 2, 82, 83, 7, 36, 2, 2, 83, 32, 3, 2,
	2, 2, 84, 89, 7, 36, 2, 2, 85, 88, 5, 31, 16, 2, 86, 88, 10, 5, 2, 2, 87,
	85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 89, 87, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93,
	7, 36, 2, 2, 93, 34, 3, 2, 2, 2, 94, 95, 7, 94, 2, 2, 95, 96, 7, 41, 2,
	2, 96, 36, 3, 2, 2, 2, 97, 102, 7, 41, 2, 2, 98, 101, 5, 35, 18, 2, 99,
	101, 10, 5, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3,
	2, 2, 2, 102, 103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 3, 2, 2,
	2, 104, 102, 3, 2, 2, 2, 105, 106, 7, 41, 2, 2, 106, 38, 3, 2, 2, 2, 107,
	108, 9, 6, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 8, 20, 2, 2, 110, 40,
	3, 2, 2, 2, 111, 115, 7, 37, 2, 2, 112, 114, 10, 7, 2, 2, 113, 112, 3,
	2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2,
	2, 116, 118, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 8, 21, 2, 2, 119,
	42, 3, 2, 2, 2, 120, 121, 7, 35, 2, 2, 121, 44, 3, 2, 2, 2, 10, 2, 73,
	79, 87, 89, 100, 102, 115, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "';'", "'('", "')'", "'^'", "'['", "']'", "'?'", "'_'", "':'",
	"'$'", "'~'", "", "", "", "", "", "", "'!'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Identifier", "IntLit",
	"StringLit", "CharLit", "WS", "LineComment", "Stress",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "Identifier", "IntLit", "EscapedQuote", "StringLit",
	"EscapedApo", "CharLit", "WS", "LineComment", "Stress",
}

type LDELexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewLDELexer(input antlr.CharStream) *LDELexer {

	l := new(LDELexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "LDE.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LDELexer tokens.
const (
	LDELexerT__0        = 1
	LDELexerT__1        = 2
	LDELexerT__2        = 3
	LDELexerT__3        = 4
	LDELexerT__4        = 5
	LDELexerT__5        = 6
	LDELexerT__6        = 7
	LDELexerT__7        = 8
	LDELexerT__8        = 9
	LDELexerT__9        = 10
	LDELexerT__10       = 11
	LDELexerT__11       = 12
	LDELexerIdentifier  = 13
	LDELexerIntLit      = 14
	LDELexerStringLit   = 15
	LDELexerCharLit     = 16
	LDELexerWS          = 17
	LDELexerLineComment = 18
	LDELexerStress      = 19
)
