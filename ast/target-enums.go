package ast

// TargetEnum ...
type TargetEnum int

//go:generate stringer -type TargetEnum target-enums.go

const (
	String TargetEnum = iota
	Char
)
