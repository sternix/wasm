// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://html.spec.whatwg.org/multipage/web-sockets.html#websocket
	WebSocket interface {
		EventTarget

		URL() string
		ReadyState() WebSocketReadyState
		BufferedAmount() int
		OnOpen(func(Event)) EventHandler
		OnError(func(ErrorEvent)) EventHandler
		OnClose(func(CloseEvent)) EventHandler
		Extensions() string
		Protocol() string
		Close(...interface{})
		OnMessage(func(MessageEvent)) EventHandler
		BinaryType() BinaryType
		SetBinaryType(BinaryType)
		Send(interface{})
	}

	// https://html.spec.whatwg.org/multipage/web-sockets.html#websocket
	CloseEvent interface {
		Event

		WasClean() bool
		Code() int
		Reason() string
	}
)

type BinaryType string

const (
	BinaryTypeBlob        BinaryType = "blob"
	BinaryTypeArrayBuffer BinaryType = "arraybuffer"
)

type WebSocketReadyState int

const (
	WebSocketReadyStateConnecting WebSocketReadyState = 0
	WebSocketReadyStateOpen       WebSocketReadyState = 1
	WebSocketReadyStateClosing    WebSocketReadyState = 2
	WebSocketReadyStateClosed     WebSocketReadyState = 3
)

// -------------8<---------------------------------------

// https://html.spec.whatwg.org/multipage/web-sockets.html#closeeventinit
type CloseEventInit struct {
	WasClean bool //false
	Code     int  // 0
	Reason   string
}

func (p CloseEventInit) toDict() js.Value {
	o := jsObject.New()
	o.Set("wasClean", p.WasClean)
	o.Set("code", p.Code)
	o.Set("reason", p.Reason)
	return o
}
