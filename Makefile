build:
	@go build -o build/wasmlife cmd/wasmlife/*.go 

run:
	@go run cmd/wasmlife/*.go

wasm:
	@GOOS=js GOARCH=wasm go build \
		-o docs/lib.wasm \
		cmd/wasmlife/*.go 
