all: wasm server

wasm:
	$(MAKE) -C ball_drop_game
	$(MAKE) -C dom
	$(MAKE) -C webgl

server:
	go build -o server/server server/server.go

serve: server wasm
	./server/server

clean:
	rm -f server/assets/*.wasm
	rm -f server/server
