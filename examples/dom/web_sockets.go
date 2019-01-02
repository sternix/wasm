// +build js,wasm

package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func TestWebSocket(doc wasm.Document) {
	wsclock := wasm.NewHTMLDivElement()
	doc.Body().AppendChild(wsclock)

	sock := wasm.NewWebSocket("ws://" + doc.Location().Host() + "/ws_time")
	if sock == nil {
		errx("Browser does not support WebSockets")
	}

	sock.OnOpen(func(wasm.Event) {
		fmt.Println("WebSocket is Opened")
	})

	sock.OnClose(func(wasm.CloseEvent) {
		fmt.Println("WebSocket is Closed")
	})

	sock.OnMessage(func(e wasm.MessageEvent) {
		if data, ok := e.Data().(string); ok {
			wsclock.SetInnerHTML("Server Time : " + data)
		}
	})
}
