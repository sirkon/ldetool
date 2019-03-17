test:
		PATH=${GOPATH}/bin:${PATH}
		go install
		go get -u github.com/stretchr/testify
		go get -u github.com/sirkon/decconv
		go generate github.com/sirkon/ldetool/testing
		which ldetool
		go test -test.v github.com/sirkon/ldetool/testing
