package ast

var _ Action = &PassUntilOrIgnore{}

// PassUntilOrIgnore ...
type PassUntilOrIgnore struct {
	access
	Limit *Target
}

func (p *PassUntilOrIgnore) Accept(d ActionDispatcher) error {
	return d.DispatchPassUntilOrIgnore(p)
}

func (p *PassUntilOrIgnore) String() string {
	pu := &PassUntil{
		Limit: p.Limit,
	}
	if p.Limit.Lower == p.Limit.Upper && p.Limit.Lower > 0 {
		return pu.String() + " or ignore otherwise"
	} else {
		return pu.String() + " or ignore if not found"
	}
}

// PassUntilTargetOrIgnore ...
func PassUntilTargetOrIgnore() *PassUntilOrIgnore {
	return &PassUntilOrIgnore{
		Limit: NewTarget(),
	}
}
