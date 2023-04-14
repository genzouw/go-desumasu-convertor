.PHONY: build test clean

.DEFAULT_GOAL := all

build: clean
	go build -o desumasu-convertor ./cmd/
test: build
	go test -v ./...
integration_test: test
	bin/test.sh
clean:
	rm -f desumasu-convertor
all: build test integration_test
