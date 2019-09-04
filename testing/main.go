package ldetesting

//go:generate ldetool --little-endian --package ldetesting parsing.lde
//go:generate ldetool --package ldetesting --go-string string.lde
//go:generate ldetool --package ldetesting --go-string regressions.lde
//go:generate ldetool --package ldetesting --go-string missing_import.lde
