test:
		PATH=${GOPATH}/bin:${PATH}
		go install
		go generate ./internal/ldetesting
		which ldetool
		go test -test.v ./internal/ldetesting

grammar:
		antlr4 -no-visitor -listener -o internal/parser -Dlanguage=Go LDE.g4
