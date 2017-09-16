package main

import (
	"fmt"
	"strings"

	"github.com/sirkon/message"
)

//go:generate ldetool generate --package main msg_translator_rules.lde

// ErrorTranslator translates gocc error messages in human readable form
type ErrorTranslator struct {
	te *TypeError
}

// NewErrorTranslator ...
func NewErrorTranslator() *ErrorTranslator {
	return &ErrorTranslator{
		te: &TypeError{},
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
		return fmt.Sprintf(
			"%d:%d: unsupported type `\033[1m%s\033[0m`, expected one of %s",
			et.te.Line,
			et.te.Column,
			string(et.te.Name),
			strings.Join(strings.Split(string(et.te.Choices), " "), ", "),
		)
	}
	return err.Error()
}

// Translate translates an error into human readable form
func (et *ErrorTranslator) Translate(err error) error {
	return fmt.Errorf("%s", et.translate(err))
}
