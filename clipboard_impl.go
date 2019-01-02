// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type clipboardImpl struct {
	*eventTargetImpl
}

func wrapClipboard(v js.Value) Clipboard {
	if isNil(v) {
		return nil
	}

	return &clipboardImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *clipboardImpl) Read() func() (DataTransfer, error) {
	return func() (DataTransfer, error) {
		result, ok := await(p.Call("read"))
		if ok {
			return wrapDataTransfer(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

func (p *clipboardImpl) ReadText() func() (string, bool) {
	return func() (string, bool) {
		result, ok := await(p.Call("readText"))
		if ok {
			return result.String(), true
		}
		return "", false
	}
}

func (p *clipboardImpl) Write(data DataTransfer) func() bool {
	return func() bool {
		_, ok := await(p.Call("write", JSValue(data)))
		return ok
	}
}

func (p *clipboardImpl) WriteText(data string) func() bool {
	return func() bool {
		_, ok := await(p.Call("writeText", data))
		return ok
	}
}

// -------------8<---------------------------------------

type clipboardEventImpl struct {
	*eventImpl
}

func wrapClipboardEvent(v js.Value) ClipboardEvent {
	if isNil(v) {
		return nil
	}

	return &clipboardEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *clipboardEventImpl) ClipboardData() DataTransfer {
	return wrapDataTransfer(p.Get("clipboardData"))
}

// -------------8<---------------------------------------

func NewClipboardEvent(typ string, eventInitDict ...ClipboardEventInit) ClipboardEvent {
	jsClipboardEvent := js.Global().Get("ClipboardEvent")
	if isNil(jsClipboardEvent) {
		return nil
	}

	switch len(eventInitDict) {
	case 0:
		return wrapClipboardEvent(jsClipboardEvent.New(typ))
	default:
		return wrapClipboardEvent(jsClipboardEvent.New(typ, eventInitDict[0].toJSObject()))
	}
}
