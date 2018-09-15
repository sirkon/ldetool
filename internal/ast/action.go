package ast

import (
	"fmt"
)

// Action абстракция действия
type Action interface {
	fmt.Stringer
	Accept(d ActionDispatcher) error
	neverCallMe()
}

type access struct{}

func (access) neverCallMe() {}
