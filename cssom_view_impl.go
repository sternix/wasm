// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type screenImpl struct {
	js.Value
}

func wrapScreen(v js.Value) Screen {
	if isNil(v) {
		return nil
	}

	return &screenImpl{
		Value: v,
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

func wrapMediaQueryList(v js.Value) MediaQueryList {
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

type mediaQueryListEventImpl struct {
	*eventImpl
}

func wrapMediaQueryListEvent(v js.Value) MediaQueryListEvent {
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

type caretPositionImpl struct {
	js.Value
}

func wrapCaretPosition(v js.Value) CaretPosition {
	if isNil(v) {
		return nil
	}

	return &caretPositionImpl{
		Value: v,
	}
}

func (p *caretPositionImpl) OffsetNode() Node {
	return wrapNode(p.Get("offsetNode"))
}

func (p *caretPositionImpl) Offset() int {
	return p.Get("offset").Int()
}

func (p *caretPositionImpl) ClientRect() DOMRect {
	return wrapDOMRect(p.Call("getClientRect"))
}

// -------------8<---------------------------------------

type geometryUtilsImpl struct {
	js.Value
}

func wrapGeometryUtils(v js.Value) GeometryUtils {
	if p := newGeometryUtilsImpl(v); p != nil {
		return p
	}
	return nil
}

func newGeometryUtilsImpl(v js.Value) *geometryUtilsImpl {
	if isNil(v) {
		return nil
	}
	return &geometryUtilsImpl{
		Value: v,
	}
}

func (p *geometryUtilsImpl) BoxQuads(options ...BoxQuadOptions) []DOMQuad {
	switch len(options) {
	case 0:
		return domQuadArrayToSlice(p.Call("getBoxQuads"))
	default:
		return domQuadArrayToSlice(p.Call("getBoxQuads", options[0].toDict()))
	}
}

func (p *geometryUtilsImpl) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.Call("convertQuadFromNode", quad.toDict(), from.JSValue()))
	default:
		return wrapDOMQuad(p.Call("convertQuadFromNode", quad.toDict(), from.JSValue(), options[0].toDict()))
	}
}

func (p *geometryUtilsImpl) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.Call("convertRectFromNode", rect.JSValue(), from.JSValue()))
	default:
		return wrapDOMQuad(p.Call("convertRectFromNode", rect.JSValue(), from.JSValue(), options[0].toDict()))
	}
}

func (p *geometryUtilsImpl) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMPoint {
	switch len(options) {
	case 0:
		return wrapDOMPoint(p.Call("convertPointFromNode", point.toDict(), from.JSValue()))
	default:
		return wrapDOMPoint(p.Call("convertPointFromNode", point.toDict(), from.JSValue(), options[0].toDict()))
	}
}

// -------------8<---------------------------------------

type cssPseudoElementImpl struct {
	*eventTargetImpl
}

func wrapCSSPseudoElement(v js.Value) CSSPseudoElement {
	if isNil(v) {
		return nil
	}
	return &cssPseudoElementImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *cssPseudoElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *cssPseudoElementImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssPseudoElementListImpl struct {
	js.Value
}

func wrapCSSPseudoElementList(v js.Value) CSSPseudoElementList {
	if isNil(v) {
		return nil
	}
	return &cssPseudoElementListImpl{
		Value: v,
	}
}

func (p *cssPseudoElementListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *cssPseudoElementListImpl) Item(index int) CSSPseudoElement {
	return wrapCSSPseudoElement(p.Call("item", index))
}

func (p *cssPseudoElementListImpl) ByType(typ string) CSSPseudoElement {
	return wrapCSSPseudoElement(p.Call("getByType", typ))
}

// -------------8<---------------------------------------
func NewMediaQueryListEvent(typ string, eventInitDict ...MediaQueryListEventInit) MediaQueryListEvent {
	jsMQLE := js.Global().Get("MediaQueryListEvent")

	switch len(eventInitDict) {
	case 0:
		return wrapMediaQueryListEvent(jsMQLE.New(typ))
	default:
		return wrapMediaQueryListEvent(jsMQLE.New(typ, eventInitDict[0].toDict()))
	}
}
