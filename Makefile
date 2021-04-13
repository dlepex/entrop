
		# thanks to: https://sohlich.github.io/post/go_makefile/
		# Go parameters
BINARY_NAME=entrop
BINARY_WASM=$(BINARY_NAME).wasm
BINARY_UNIX=entrop

all: test build build-wasm

test:
	go test -v
clean:
	go clean
	rm -f $(BINARY_NAME) $(BINARY_WASM)
run:
	go build -o $(BINARY_NAME) -v
	./$(BINARY_NAME)


# build:
build:
	go build -o $(BINARY_NAME)
build-wasm:
	GOOS=js GOARCH=wasm go build -o $(BINARY_WASM)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v
