test:
		PATH=${GOPATH}/bin:${PATH}
		go install
		go get -u github.com/stretchr/testify
		go get -u github.com/sirkon/decconv
		go generate github.com/sirkon/ldetool/testing
		which ldetool
		go test -test.v github.com/sirkon/ldetool/testing


version:
		git tag --list | tail -1 | xargs printf 'package main\n\nconst ldetoolVersion = "%s"\n' > version.go
		git commit -a --amend

grammar:
		antlr4 -no-visitor -listener -o internal/parser -Dlanguage=Go LDE.g4
