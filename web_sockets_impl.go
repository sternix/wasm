// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewWebSocket(url string, protocols ...string) WebSocket {
	if jsWebSocket := jsGlobal.Get("WebSocket"); jsWebSocket.Valid() {
		switch len(protocols) {
		case 0:
			return wrapWebSocket(jsWebSocket.New(url))
		case 1:
			return wrapWebSocket(jsWebSocket.New(url, protocols[0]))
		default:
			return wrapWebSocket(jsWebSocket.New(url, sliceToJsArray(protocols)))
		}
	}
	return nil
}

func NewCloseEvent(typ string, cei ...CloseEventInit) CloseEvent {
	if jsCloseEvent := jsGlobal.Get("CloseEvent"); jsCloseEvent.Valid() {
		switch len(cei) {
		case 0:
			return wrapCloseEvent(jsCloseEvent.New(typ))
		default:
			return wrapCloseEvent(jsCloseEvent.New(typ, cei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webSocketImpl struct {
	*eventTargetImpl
}

func wrapWebSocket(v Value) WebSocket {
	if v.Valid() {
		return &webSocketImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *webSocketImpl) URL() string {
	return p.Get("url").String()
}

func (p *webSocketImpl) ReadyState() WebSocketReadyState {
	return WebSocketReadyState(p.Get("readyState").Int())
}

func (p *webSocketImpl) BufferedAmount() int {
	return p.Get("bufferedAmount").Int()
}

func (p *webSocketImpl) OnOpen(fn func(Event)) EventHandler {
	return p.On("open", fn)
}

func (p *webSocketImpl) OnError(fn func(ErrorEvent)) EventHandler {
	return p.On("error", func(e Event) {
		if ee, ok := e.(ErrorEvent); ok {
			fn(ee)
		}
	})
}

func (p *webSocketImpl) OnClose(fn func(CloseEvent)) EventHandler {
	return p.On("close", func(e Event) {
		if ce, ok := e.(CloseEvent); ok {
			fn(ce)
		}
	})
}

func (p *webSocketImpl) Extensions() string {
	return p.Get("extensions").String()
}

func (p *webSocketImpl) Protocol() string {
	return p.Get("protocol").String()
}

func (p *webSocketImpl) Close(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("close")
	case 1:
		if code, ok := args[0].(int); ok {
			p.Call("close", code)
		}
	case 2:
		if code, ok := args[0].(int); ok {
			if reason, ok := args[1].(string); ok {
				p.Call("close", code, reason)
			}
		}
	}
}

func (p *webSocketImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return p.On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *webSocketImpl) BinaryType() BinaryType {
	return BinaryType(p.Get("binaryType").String())
}

func (p *webSocketImpl) SetBinaryType(bt BinaryType) {
	p.Set("binaryType", bt)
}

func (p *webSocketImpl) Send(typ interface{}) {
	switch x := typ.(type) {
	case string:
		p.Call("send", x)
	case []byte:
		ta := js.TypedArrayOf(x)
		blob := NewBlob(ta)
		p.Call("send", JSValue(blob))
		ta.Release()
	case Blob, ArrayBuffer, ArrayBufferView:
		p.Call("send", JSValue(x))
	}
}

// -------------8<---------------------------------------

type closeEventImpl struct {
	*eventImpl
}

func wrapCloseEvent(v Value) CloseEvent {
	if v.Valid() {
		return &closeEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *closeEventImpl) WasClean() bool {
	return p.Get("wasClean").Bool()
}

func (p *closeEventImpl) Code() int {
	return p.Get("code").Int()
}

func (p *closeEventImpl) Reason() string {
	return p.Get("reason").String()
}

// -------------8<---------------------------------------
