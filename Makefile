EXE=entrop
EXE_WASM=$(EXE).wasm
EXE_LINUX=$(EXE)
SALT_FILE=salt_v5

all: test build

release: test mk-version build build-wasm

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
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(EXE_LINUX) -v
mk-salt:
	dd if=/dev/random of=embed/$(SALT_FILE) bs=64 count=1
mk-version:
	git rev-parse HEAD > embed/gitcommit
 	git describe --abbrev=0 --tags > embed/gittag
