package ast

import (
	"fmt"
	"strconv"
)

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
func (t *Target) SetChar(text string) error {
	// TODO get rid of this after https://github.com/antlr/antlr4/pull/2642 will be merged
	if _, err := strconv.Unquote(text); err != nil {
		return fmt.Errorf("process rune %s: %w", text, err)
	}

	t.Type = Char
	t.Value = text

	return nil
}

// SetString sets target into String
func (t *Target) SetString(text string) error {
	// TODO get rid of this after https://github.com/antlr/antlr4/pull/2642 will be merged
	if _, err := strconv.Unquote(text); err != nil {
		return fmt.Errorf("process string %s: %w", text, err)
	}

	t.Type = String
	t.Value = text

	return nil
}

// SetLimit sets target limit
func (t *Target) SetLimit(upper int) {
	t.Upper = upper
}

// SetBound sets target bound
func (t *Target) SetBound(lower, upper int) {
	t.Lower = lower
	t.Upper = upper
}

// SetJump sets target offset jump
func (t *Target) SetJump(lower int) {
	t.Lower = lower
}
