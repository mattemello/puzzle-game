all: src/wasm/main server

clear:
	rm -f *.wasm

src/wasm/main:
	cd src/wasm
	GOOS=js GOARCH=wasm go build -o ./assets/main.wasm

server:
	go run src/server/main.go
