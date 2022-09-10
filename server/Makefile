BINARY_NAME=crasher
BIN_PATH=bin/
LINTER=golangci-lint

.ONESHELL: test build lint all
.PHONY: test build lint all

all: test build lint

test:
	@go test ./...

build:
	@go build -o $(BIN_PATH)$(BINARY_NAME) cmd/crasher/main.go

lint:
	@$(LINTER) run ./...
