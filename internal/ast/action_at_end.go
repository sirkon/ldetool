package ast

var _ Action = AtEnd{}

// AtEnd ...
type AtEnd struct {
	access
}

func (a AtEnd) Accept(d ActionDispatcher) error {
	return d.DispatchAtEnd(a)
}

func (AtEnd) String() string {
	return "Must be the end here"
}
