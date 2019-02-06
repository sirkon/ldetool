package ast

import (
	"fmt"
)

var _ Action = PassHeadingCharacters('1')

type PassHeadingCharacters string

func (a PassHeadingCharacters) String() string {
	return fmt.Sprintf("pass all '%s' characters in the head of rest", string(a))
}

func (a PassHeadingCharacters) Accept(d ActionDispatcher) error {
	return d.DispatchPassHeadingCharacters(a)
}

func (a PassHeadingCharacters) neverCallMe() {}
