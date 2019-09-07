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
	return "Check if we reached an end of data"
}
