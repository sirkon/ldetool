package ast

// Target ...
type Target struct {
	Type  TargetEnum
	Value string
	Lower int
	Upper int
	Close bool
}

// NewTarget ...
func NewTarget() *Target {
	return &Target{}
}

// SetClose sets target type to close
func (t *Target) SetClose() {
	t.Close = true
}

// SetChar sets target into Char
func (t *Target) SetChar(text string) {
	t.Type = Char
	t.Value = text
}

// SetString sets target into String
func (t *Target) SetString(text string) {
	t.Type = String
	t.Value = text
}

// SetLimit sets target limit
func (t *Target) SetLimit(lower int) {
	t.Lower = lower
}

// SetBound sets target bound
func (t *Target) SetBound(lower, upper int) {
	t.Lower = lower
	t.Upper = upper
}
