package main

import (
	"fmt"
	"strings"

	"github.com/sirkon/message"
)

//go:generate ldetool generate --package main msg_translator_rules.lde

// ErrorTranslator translates gocc error messages in human readable form
type ErrorTranslator struct {
	te  *TypeError
	ne1 *NoEnd1Error
}

// NewErrorTranslator ...
func NewErrorTranslator() *ErrorTranslator {
	return &ErrorTranslator{
		te:  &TypeError{},
		ne1: &NoEnd1Error{},
	}
}

func (et *ErrorTranslator) translate(err error) string {
	var ok bool
	var errParsing error
	defer func() {
		if errParsing != nil {
			message.Warning(errParsing)
		}
	}()

	data := []byte(err.Error())
	if ok, errParsing = et.te.Extract(data); ok {
		var msg string
		if len(et.te.Name) == 0 {
			msg = fmt.Sprintf("no type given")
		} else {
			msg = fmt.Sprintf("unsupported type `\033[1m%s\033[0m`", string(et.te.Name))
		}
		choices := strings.Split(string(et.te.Choices), " ")
		for i, choice := range choices {
			choices[i] = fmt.Sprintf("\033[1m%s\033[0m", choice)
		}
		return fmt.Sprintf(
			"%d:%d: %s, use one of %s",
			et.te.Line,
			et.te.Column,
			msg,
			strings.Join(choices, ", "),
		)
	} else if ok, errParsing = et.ne1.Extract(data); ok {
		return fmt.Sprintf("%d:%d: it looks like `;` is missed in the previous rule", et.ne1.Line, et.ne1.Column)
	}
	return err.Error()
}

// Translate translates an error into human readable form
func (et *ErrorTranslator) Translate(err error) error {
	return fmt.Errorf("%s", et.translate(err))
}
