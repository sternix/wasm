wasm:
	GOOS=js GOARCH=wasm go112 build -o ballgame.wasm main.go

server:
	go build server.go

serve: wasm server
	./server

clean:
	rm -f ballgame.wasm
	rm -f server
