test:
		go generate github.com/sirkon/ldetool/testing
		go test -test.v github.com/sirkon/ldetool/testing
