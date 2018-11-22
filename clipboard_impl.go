// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type clipboardImpl struct {
	*eventTargetImpl
}

func newClipboard(v js.Value) Clipboard {
	if isNil(v) {
		return nil
	}

	return &clipboardImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *clipboardImpl) Read() Promise {
	return newPromiseImpl(p.Call("read"))
}

func (p *clipboardImpl) ReadText() Promise {
	return newPromiseImpl(p.Call("readText"))
}

func (p *clipboardImpl) Write(data DataTransfer) Promise {
	return newPromiseImpl(p.Call("write", data.JSValue()))
}

func (p *clipboardImpl) WriteText(data string) Promise {
	return newPromiseImpl(p.Call("writeText", data))
}

// -------------8<---------------------------------------

type clipboardEventImpl struct {
	*eventImpl
}

func newClipboardEvent(v js.Value) ClipboardEvent {
	if isNil(v) {
		return nil
	}

	return &clipboardEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *clipboardEventImpl) ClipboardData() DataTransfer {
	return newDataTransfer(p.Get("clipboardData"))
}

// -------------8<---------------------------------------

func NewClipboardEvent(typ string, eventInitDict ...ClipboardEventInit) ClipboardEvent {
	jsClipboardEvent := js.Global().Get("ClipboardEvent")
	if isNil(jsClipboardEvent) {
		return nil
	}

	switch len(eventInitDict) {
	case 0:
		return newClipboardEvent(jsClipboardEvent.New(typ))
	default:
		return newClipboardEvent(jsClipboardEvent.New(typ, toJSONObject(eventInitDict[0])))
	}
}
