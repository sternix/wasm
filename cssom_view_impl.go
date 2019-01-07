// +build js,wasm

package wasm

// -------------8<---------------------------------------

type screenImpl struct {
	Value
}

func wrapScreen(v Value) Screen {
	if v.Valid() {
		return &screenImpl{
			Value: v,
		}
	}
	return nil
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

func wrapMediaQueryList(v Value) MediaQueryList {
	if v.Valid() {
		return &mediaQueryListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
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

func wrapMediaQueryListEvent(v Value) MediaQueryListEvent {
	if v.Valid() {
		return &mediaQueryListEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *mediaQueryListEventImpl) Media() string {
	return p.Get("media").String()
}

func (p *mediaQueryListEventImpl) Matches() bool {
	return p.Get("matches").Bool()
}

// -------------8<---------------------------------------

type caretPositionImpl struct {
	Value
}

func wrapCaretPosition(v Value) CaretPosition {
	if v.Valid() {
		return &caretPositionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *caretPositionImpl) OffsetNode() Node {
	return wrapAsNode(p.Get("offsetNode"))
}

func (p *caretPositionImpl) Offset() int {
	return p.Get("offset").Int()
}

func (p *caretPositionImpl) ClientRect() DOMRect {
	return wrapDOMRect(p.Call("getClientRect"))
}

// -------------8<---------------------------------------

type geometryUtilsImpl struct {
	Value
}

func wrapGeometryUtils(v Value) GeometryUtils {
	if p := newGeometryUtilsImpl(v); p != nil {
		return p
	}
	return nil
}

func newGeometryUtilsImpl(v Value) *geometryUtilsImpl {
	if v.Valid() {
		return &geometryUtilsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *geometryUtilsImpl) BoxQuads(options ...BoxQuadOptions) []DOMQuad {
	switch len(options) {
	case 0:
		return domQuadArrayToSlice(p.Call("getBoxQuads"))
	default:
		return domQuadArrayToSlice(p.Call("getBoxQuads", options[0].toJSObject()))
	}
}

func (p *geometryUtilsImpl) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.Call("convertQuadFromNode", quad.toJSObject(), JSValue(from)))
	default:
		return wrapDOMQuad(p.Call("convertQuadFromNode", quad.toJSObject(), JSValue(from), options[0].toJSObject()))
	}
}

func (p *geometryUtilsImpl) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.Call("convertRectFromNode", JSValue(rect), JSValue(from)))
	default:
		return wrapDOMQuad(p.Call("convertRectFromNode", JSValue(rect), JSValue(from), options[0].toJSObject()))
	}
}

func (p *geometryUtilsImpl) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMPoint {
	switch len(options) {
	case 0:
		return wrapDOMPoint(p.Call("convertPointFromNode", point.toJSObject(), JSValue(from)))
	default:
		return wrapDOMPoint(p.Call("convertPointFromNode", point.toJSObject(), JSValue(from), options[0].toJSObject()))
	}
}

// -------------8<---------------------------------------

type cssPseudoElementImpl struct {
	*eventTargetImpl
}

func wrapCSSPseudoElement(v Value) CSSPseudoElement {
	if v.Valid() {
		return &cssPseudoElementImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *cssPseudoElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *cssPseudoElementImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssPseudoElementListImpl struct {
	Value
}

func wrapCSSPseudoElementList(v Value) CSSPseudoElementList {
	if v.Valid() {
		return &cssPseudoElementListImpl{
			Value: v,
		}
	}
	return nil
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
	jsMQLE := jsGlobal.Get("MediaQueryListEvent")

	switch len(eventInitDict) {
	case 0:
		return wrapMediaQueryListEvent(jsMQLE.New(typ))
	default:
		return wrapMediaQueryListEvent(jsMQLE.New(typ, eventInitDict[0].toJSObject()))
	}
}
