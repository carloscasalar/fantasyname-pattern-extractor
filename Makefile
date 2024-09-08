# Define the binary name
BINARY_NAME=./out/extract-pattern

# Define the main package
MAIN_PACKAGE=./cmd/extract-pattern

# Go Bin
GO_BIN=$(shell which go)

# Ensure the out directory exists
out:
	mkdir -p out

# Install the required tools for go generators
install-tools:
	@echo "Parsing tools.go and installing dependencies..."
	@go list -e -f '{{join .Imports " "}}' tools.go | xargs -t -n 1 $(GO_BIN) install
	@echo "all tools installed"

# Build the Go application
build: out
	@go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the Go application
run: build
	$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

# Run tests
test:
	@go test -v ./...

# Run the linter
lint:
	@revive -config .revive.toml -formatter friendly ./...

# Run all checks
check: lint test

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)

# Default target
all: build

# Hack to make run proxy the arguments to the binary
%:
	@true

.PHONY: out install-tools build run test lint check clean all
