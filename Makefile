test:
		PATH=${GOPATH}/bin:${PATH}
		go install github.com/sirkon/ldetool
		go get -u github.com/stretchr/testify
		go generate github.com/sirkon/ldetool/testing
		which ldetool
		go test -test.v github.com/sirkon/ldetool/testing
