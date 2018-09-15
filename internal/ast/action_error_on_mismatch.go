package ast

var _ Action = ErrorOnMismatch{}

// ErrorOnMismatch ...
type ErrorOnMismatch struct {
	access
}

func (e ErrorOnMismatch) Accept(d ActionDispatcher) error {
	return d.DispatchErrorMismatch(e)
}

func (ErrorOnMismatch) String() string {
	return "Treating all mismatches at the root level in the rest of the rule as critical errors"
}
