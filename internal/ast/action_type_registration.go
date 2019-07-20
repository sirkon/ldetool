package ast

import (
	"github.com/sirkon/ldetool/internal/types"
)

var _ Action = TypeRegistration{}

type TypeRegistration struct {
	access

	tr map[string]types.TypeRegistration
}

func NewTypeRegistration(tr map[string]types.TypeRegistration) TypeRegistration {
	return TypeRegistration{
		tr: tr,
	}
}

func (t TypeRegistration) String() string {
	return "type registration"
}

func (t TypeRegistration) Types() map[string]types.TypeRegistration {
	return t.tr
}

func (t TypeRegistration) Accept(d ActionDispatcher) error {
	return d.DispatchTypeRegistration(t)
}
