package main

import (
	"fmt"

	"github.com/sirkon/message"
)

//go:generate ldetool generate --package main msg_translator_rules.lde

// ErrorTranslator translates gocc error messages in human readable form
type ErrorTranslator struct {
	o *Orig
}

// NewErrorTranslator ...
func NewErrorTranslator() *ErrorTranslator {
	return &ErrorTranslator{
		o: &Orig{},
	}
}

func (et *ErrorTranslator) translate(err error) string {
	var ok bool
	var errParsing error
	data := []byte(err.Error())
	defer func() {
		if errParsing != nil {
			message.Warning(errParsing)
		}
	}()

	if ok, errParsing = et.o.Extract(data); ok {
		return fmt.Sprintf("%d:%d: %s", et.o.Line, et.o.Column, string(et.o.Message))
	}

	return err.Error()
}

// Translate translates an error into human readable form
func (et *ErrorTranslator) Translate(err error) error {
	return fmt.Errorf("%s", et.translate(err))
}
