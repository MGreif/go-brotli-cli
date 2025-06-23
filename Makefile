prepare:
	mkdir -p build

build: prepare
	go build -o ./build/brotli-cli ./cmd/*

test:
	go test ./... -v