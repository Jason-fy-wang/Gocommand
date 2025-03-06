.PHONY: all clean fmt import build test display

# Define the Go binary name
BINARY_NAME=console

# Define the main Go file
CONSOLE_MAIN_FILE=cmd/console/main.go

dirs=trending tui utils parser

all: clean build

fmt:
	find . -path ./vendor -prune -o -name '*.go' -exec gofmt -w {} +

import:
	find . -path ./vendor -prune -o -name '*.go' -exec goimports -w {} +

# Clean up the previous builds
clean:
	rm -f $(BINARY_NAME)

display:
	@echo "display fmt difference"
	gofmt -d -e -l  ${dirs}
	@echo "display import difference"
	goimports -d -e -l ${dirs}

check_update:
	go list -m -mod=mod -u all 

# ./...  the mean of ... in here is: all the packages in the current directory and all of its subdirectories
test:
	go test -cover -coverprofile=coverage.out ./...

# Build the Go binary
build: fmt import test
	go build -o $(BINARY_NAME) ${CONSOLE_MAIN_FILE}


