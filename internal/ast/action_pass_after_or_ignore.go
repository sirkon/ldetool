package ast

var _ Action = &PassAfterOrIgnore{}

// PassAfterOrIgnore ...
type PassAfterOrIgnore struct {
	access
	Limit *Target
}

func (p *PassAfterOrIgnore) Accept(d ActionDispatcher) error {
	return d.DispatchPassAfterOrIgnore(p)
}

func (p *PassAfterOrIgnore) String() string {
	pu := &PassAfter{
		Limit: p.Limit,
	}
	if p.Limit.Lower == p.Limit.Upper && p.Limit.Lower > 0 {
		return pu.String() + " or ignore otherwise"
	} else {
		return pu.String() + " or ignore if not found"
	}
}

// PassAfterTargetOrIgnore ...
func PassAfterTargetOrIgnore() *PassAfterOrIgnore {
	return &PassAfterOrIgnore{
		Limit: NewTarget(),
	}
}
