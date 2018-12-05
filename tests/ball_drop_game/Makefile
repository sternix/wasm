ballgame.wasm:
	GOOS=js GOARCH=wasm go build -o ballgame.wasm main.go

serve: ballgame.wasm
	go run server.go

clean:
	rm -f ballgame.wasm
