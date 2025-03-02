.PHONY: all clean fmt import build

# Define the Go binary name
BINARY_NAME=console

# Define the main Go file
MAIN_FILE=main.go

dirs=trending tui utils parser

all: clean build

fmt:
	gofmt -w ${dirs}

import:
	goimports -w ${dirs}

# Clean up the previous builds
clean:
	rm -f $(BINARY_NAME)

display:
	@echo "display fmt difference"
	gofmt -d -e -l  ${dirs}
	@echo "display import difference"
	goimports -d -e -l ${dirs}

# Build the Go binary
build: fmt import
	go build -o $(BINARY_NAME) .


