BINARY_NAME=crasher
BIN_PATH=bin/
LINTER=golangci-lint

.ONESHELL: test build lint all run 
.PHONY: test build lint all run 

all: build test lint

build:
	@go build -o $(BIN_PATH)$(BINARY_NAME) cmd/crasher/main.go

test:
	@go test -v -race -cover ./...

lint:
	@$(LINTER) run ./...

run:
	@export CRASHER_SERVER_ADDRESS=:8080
	@go run cmd/crasher/main.go
