package ast

var _ Action = &PassBeforeOrIgnore{}

// PassBeforeOrIgnore ...
type PassBeforeOrIgnore struct {
	access
	Limit *Target
}

func (p *PassBeforeOrIgnore) Accept(d ActionDispatcher) error {
	return d.DispatchPassBeforeOrIgnore(p)
}

func (p *PassBeforeOrIgnore) String() string {
	pu := &PassBefore{
		Limit: p.Limit,
	}
	if p.Limit.Lower == p.Limit.Upper && p.Limit.Lower > 0 {
		return pu.String() + " or ignore otherwise"
	} else {
		return pu.String() + " or ignore if not found"
	}
}

// PassBeforeTargetOrIgnore ...
func PassBeforeTargetOrIgnore() *PassBeforeOrIgnore {
	return &PassBeforeOrIgnore{
		Limit: NewTarget(),
	}
}
