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

func (p *clipboardImpl) Read() func() (DataTransfer, error) {
	return func() (DataTransfer, error) {
		result, ok := Await(p.Call("read"))
		if ok {
			return newDataTransfer(result), nil
		}
		return nil, newDOMException(result)
	}
}

func (p *clipboardImpl) ReadText() func() (string, bool) {
	return func() (string, bool) {
		result, ok := Await(p.Call("readText"))
		if ok {
			return result.String(), true
		}
		return "", false
	}
}

func (p *clipboardImpl) Write(data DataTransfer) func() bool {
	return func() bool {
		_, ok := Await(p.Call("write", data.JSValue()))
		return ok
	}
}

func (p *clipboardImpl) WriteText(data string) func() bool {
	return func() bool {
		_, ok := Await(p.Call("writeText", data))
		return ok
	}
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
		return newClipboardEvent(jsClipboardEvent.New(typ, eventInitDict[0].toDict()))
	}
}
