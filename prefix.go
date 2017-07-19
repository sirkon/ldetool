package main

import "strings"

// Prefix for gravity names
type Prefix []string

// NewPrefix constructor
func NewPrefix() Prefix {
	return nil
}

// Add item to the chain
func (p Prefix) Add(name string) Prefix {
	return Prefix(append(p, name))
}

// String ...
func (p Prefix) String() string {
	return strings.Join(p, ".")
}
