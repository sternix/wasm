// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type screenImpl struct {
	*objectImpl
}

func newScreen(v js.Value) Screen {
	if isNil(v) {
		return nil
	}

	return &screenImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *screenImpl) AvailWidth() int {
	return p.Get("availWidth").Int()
}

func (p *screenImpl) AvailHeight() int {
	return p.Get("availHeight").Int()
}

func (p *screenImpl) Width() int {
	return p.Get("width").Int()
}

func (p *screenImpl) Height() int {
	return p.Get("height").Int()
}

func (p *screenImpl) ColorDepth() int {
	return p.Get("colorDepth").Int()
}

func (p *screenImpl) PixelDepth() int {
	return p.Get("pixelDepth").Int()
}

// -------------8<---------------------------------------

type mediaQueryListImpl struct {
	*eventTargetImpl
}

func newMediaQueryList(v js.Value) MediaQueryList {
	if isNil(v) {
		return nil
	}

	return &mediaQueryListImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *mediaQueryListImpl) Media() string {
	return p.Get("media").String()
}

func (p *mediaQueryListImpl) Matches() bool {
	return p.Get("matches").Bool()
}

func (p *mediaQueryListImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

// -------------8<---------------------------------------

type caretPositionImpl struct {
	*objectImpl
}

func newCaretPosition(v js.Value) CaretPosition {
	if isNil(v) {
		return nil
	}

	return &caretPositionImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *caretPositionImpl) OffsetNode() Node {
	return newNode(p.Get("offsetNode"))
}

func (p *caretPositionImpl) Offset() int {
	return p.Get("offset").Int()
}

func (p *caretPositionImpl) ClientRect() DOMRect {
	return newDOMRect(p.Call("getClientRect"))
}

// -------------8<---------------------------------------

type mediaQueryListEventImpl struct {
	*eventImpl
}

func newMediaQueryListEvent(v js.Value) MediaQueryListEvent {
	if isNil(v) {
		return nil
	}

	return &mediaQueryListEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *mediaQueryListEventImpl) Media() string {
	return p.Get("media").String()
}

func (p *mediaQueryListEventImpl) Matches() bool {
	return p.Get("matches").Bool()
}

// -------------8<---------------------------------------

func NewMediaQueryListEvent(typ string, eventInitDict ...MediaQueryListEventInit) MediaQueryListEvent {
	jsMQLE := js.Global().Get("MediaQueryListEvent")

	switch len(eventInitDict) {
	case 0:
		return newMediaQueryListEvent(jsMQLE.New(typ))
	default:
		return newMediaQueryListEvent(jsMQLE.New(typ, toJSONObject(eventInitDict[0])))
	}
}
