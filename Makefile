test:
		PATH=${GOPATH}/bin:${PATH}
		go install
		go generate ./testing
		which ldetool
		go test -test.v ./testing

grammar:
		antlr4 -no-visitor -listener -o internal/parser -Dlanguage=Go LDE.g4
