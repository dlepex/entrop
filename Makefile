EXE=entrop
EXE_WASM=$(EXE).wasm
EXE_LINUX=$(EXE)

all: test build build-wasm

test:
	go test -v
clean:
	go clean
	rm -f $(EXE) $(EXE_WASM)
run:
	go build -o $(EXE) -v
	./$(EXE)
build:
	go build -o $(EXE)
build-wasm:
	GOOS=js GOARCH=wasm go build -o $(EXE_WASM)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(EXE_LINUX) -v
