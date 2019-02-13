module github.com/sirkon/ldetool/v3

require (
	github.com/antlr/antlr4 v0.0.0-20190207013812-1c6c62afc7cb
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sanity-io/litter v1.1.0
	github.com/sirkon/decconv v1.0.0
	github.com/sirkon/gosrcfmt v1.5.0
	github.com/sirkon/gotify v0.5.0
	github.com/sirkon/ldetool/internal v0.0.0-00010101000000-000000000000
	github.com/sirkon/message v1.5.1
	github.com/stretchr/testify v1.2.2
	github.com/urfave/cli v1.20.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

replace github.com/sirkon/ldetool/internal => ./internal
