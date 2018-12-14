all: wasm server

wasm:
	GOOS=js GOARCH=wasm go112 build -o webgltriangle.wasm main.go

server:
	go build server.go

serve: server wasm
	./server

clean:
	rm -f *.wasm
	rm -f server
