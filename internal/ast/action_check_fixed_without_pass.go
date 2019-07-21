package ast

import (
	"fmt"
)

var _ Action = &CheckFixedWithoutPass{}

// CheckFixedWithoutPass ...
type CheckFixedWithoutPass struct {
	access
	Limit *Target
}

func (pu *CheckFixedWithoutPass) Accept(d ActionDispatcher) error {
	if err := d.DispatchCheckFixedWithoutPass(pu); err != nil {
		return err
	}
	return nil
}

func (pu *CheckFixedWithoutPass) String() string {
	if pu.Limit.Type == String {
		return fmt.Sprintf("Check if the rest after %s character starts with prefix %s", posLit(pu.Limit.Lower+1), pu.Limit.Value)
	}
	return fmt.Sprintf("Check if %s character in the rest equals to %s", posLit(pu.Limit.Lower+1), pu.Limit.Value)
}

// CheckFixedTargetWithoutPass ...
func CheckFixedTargetWithoutPass() *CheckFixedWithoutPass {
	return &CheckFixedWithoutPass{
		Limit: NewTarget(),
	}
}
