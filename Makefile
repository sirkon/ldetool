test:
		PATH=${GOPATH}/bin:${PATH}
		cd ldetool && go install
		cd ldetool && go generate ./testing
		which ldetool
		cd ldetool && go test -test.v ./testing

grammar:
		antlr4 -no-visitor -listener -o internal/parser -Dlanguage=Go LDE.g4
