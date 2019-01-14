// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewWebSocket(url string, protocols ...string) WebSocket {
	if jsWebSocket := jsGlobal.get("WebSocket"); jsWebSocket.valid() {
		switch len(protocols) {
		case 0:
			return wrapWebSocket(jsWebSocket.jsNew(url))
		case 1:
			return wrapWebSocket(jsWebSocket.jsNew(url, protocols[0]))
		default:
			return wrapWebSocket(jsWebSocket.jsNew(url, sliceToJsArray(protocols)))
		}
	}
	return nil
}

func NewCloseEvent(typ string, cei ...CloseEventInit) CloseEvent {
	if jsCloseEvent := jsGlobal.get("CloseEvent"); jsCloseEvent.valid() {
		switch len(cei) {
		case 0:
			return wrapCloseEvent(jsCloseEvent.jsNew(typ))
		default:
			return wrapCloseEvent(jsCloseEvent.jsNew(typ, cei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webSocketImpl struct {
	*eventTargetImpl
}

func wrapWebSocket(v Value) WebSocket {
	if v.valid() {
		return &webSocketImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *webSocketImpl) URL() string {
	return p.get("url").toString()
}

func (p *webSocketImpl) ReadyState() WebSocketReadyState {
	return WebSocketReadyState(p.get("readyState").toInt())
}

func (p *webSocketImpl) BufferedAmount() int {
	return p.get("bufferedAmount").toInt()
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
	return p.get("extensions").toString()
}

func (p *webSocketImpl) Protocol() string {
	return p.get("protocol").toString()
}

func (p *webSocketImpl) Close(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("close")
	case 1:
		if code, ok := args[0].(int); ok {
			p.call("close", code)
		}
	case 2:
		if code, ok := args[0].(int); ok {
			if reason, ok := args[1].(string); ok {
				p.call("close", code, reason)
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
	return BinaryType(p.get("binaryType").toString())
}

func (p *webSocketImpl) SetBinaryType(bt BinaryType) {
	p.set("binaryType", bt)
}

func (p *webSocketImpl) Send(typ interface{}) {
	switch x := typ.(type) {
	case string:
		p.call("send", x)
	case []byte:
		ta := js.TypedArrayOf(x)
		blob := NewBlob(ta)
		p.call("send", JSValue(blob))
		ta.Release()
	case Blob, ArrayBuffer, ArrayBufferView:
		p.call("send", JSValue(x))
	}
}

// -------------8<---------------------------------------

type closeEventImpl struct {
	*eventImpl
}

func wrapCloseEvent(v Value) CloseEvent {
	if v.valid() {
		return &closeEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *closeEventImpl) WasClean() bool {
	return p.get("wasClean").toBool()
}

func (p *closeEventImpl) Code() uint16 {
	return p.get("code").toUint16()
}

func (p *closeEventImpl) Reason() string {
	return p.get("reason").toString()
}

// -------------8<---------------------------------------
