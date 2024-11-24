all: cmd/wasm/main cmd/server/main

clear:
	rm -f *.wasm

cmd/wasm/main:
	cd cmd/wasm & GOOS=js GOARCH=wasm go build -o ./assets/main.wasm

server:
	go run cmd/server/main.go
