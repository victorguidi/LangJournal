all: build

build:
	@go build -v -o langjournal ./src/

test:
	@go test -v ./src/...

run: all
	@./langjournal

.PHONY: build run all test
