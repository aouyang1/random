usage:
	@echo "make all       : Runs all tests and benchmarks"
	@echo "make test      : Runs test suite"
	@echo "make bench     : Runs benchmarks"
	@echo "make travis-ci : Travis CI specific testing"

all: test bench

test:
	go test -race -cover ./...

bench:
	go test ./... -run=XX -bench=. -test.benchmem

travis-ci:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic
