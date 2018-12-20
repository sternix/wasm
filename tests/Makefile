all: wasm server

wasm:
	$(MAKE) -C ball_drop_game
	$(MAKE) -C dom
	$(MAKE) -C webgl

server:
	$(MAKE) -C server

serve: server wasm
	./server/server

clean:
	rm -f server/assets/*.wasm
	rm -f server/server
