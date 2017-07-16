package generator

// Generator describes methods needed of data lookup and extraction
type Generator interface {
	// Lookups
	LookupString(anchor string)
	LookupLimitedString(anchor string)
	LookupBoundedString(anchor string)
	LookupChar(anchor string)
	LookupLimitedChar(anchor string)
	LookupBoundedChar(anchor string)

	// Passes
	PassFoundString(anchor string)
	PassFoundChar()
	PassBoundedFoundString(anchor string)
	PassBoundedFoundChar(anchor string)
	PassN(n int)

	// Head
	HeadString(anchor string)
	HeadChar(anchor string)

	// Data handlers
	AddField(name string, fieldType string)

	// Consumers
	ConsumeField(name string)

	// Optionals
	OpenOptionalScope(name string)
	ExitOptionalScope() // We always know what scope we are in

	// Stress trigger mismatch treatment as serious error
	Stress()

	// Generate code
	Generate() error
}
