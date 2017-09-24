package generator

// PlatformType is used to point what platform to generate code for
type PlatformType int

const (
	// Universal for universally compatible code generation
	Universal PlatformType = iota

	// LittleEndian is for little endian architecture code generation
	LittleEndian

	// BigEndian is for big endian architecture code generation
	BigEndian
)
