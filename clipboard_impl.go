// +build js,wasm

package wasm

// -------------8<---------------------------------------

type clipboardImpl struct {
	*eventTargetImpl
}

func wrapClipboard(v Value) Clipboard {
	if v.valid() {
		return &clipboardImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *clipboardImpl) Read() func() (DataTransfer, error) {
	return func() (DataTransfer, error) {
		result, ok := await(p.call("read"))
		if ok {
			return wrapDataTransfer(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

func (p *clipboardImpl) ReadText() func() (string, bool) {
	return func() (string, bool) {
		result, ok := await(p.call("readText"))
		if ok {
			return result.toString(), true
		}
		return "", false
	}
}

func (p *clipboardImpl) Write(data DataTransfer) func() bool {
	return func() bool {
		_, ok := await(p.call("write", JSValue(data)))
		return ok
	}
}

func (p *clipboardImpl) WriteText(data string) func() bool {
	return func() bool {
		_, ok := await(p.call("writeText", data))
		return ok
	}
}

// -------------8<---------------------------------------

type clipboardEventImpl struct {
	*eventImpl
}

func wrapClipboardEvent(v Value) ClipboardEvent {
	if v.valid() {
		return &clipboardEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *clipboardEventImpl) ClipboardData() DataTransfer {
	return wrapDataTransfer(p.get("clipboardData"))
}

// -------------8<---------------------------------------

func NewClipboardEvent(typ string, eventInitDict ...ClipboardEventInit) ClipboardEvent {
	if jsClipboardEvent := jsGlobal.get("ClipboardEvent"); jsClipboardEvent.valid() {
		switch len(eventInitDict) {
		case 0:
			return wrapClipboardEvent(jsClipboardEvent.jsNew(typ))
		default:
			return wrapClipboardEvent(jsClipboardEvent.jsNew(typ, eventInitDict[0].toJSObject()))
		}
	}
	return nil
}
