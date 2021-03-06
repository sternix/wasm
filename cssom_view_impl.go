// +build js,wasm

package wasm

// -------------8<---------------------------------------

type screenImpl struct {
	Value
}

func wrapScreen(v Value) Screen {
	if v.valid() {
		return &screenImpl{
			Value: v,
		}
	}
	return nil
}

func (p *screenImpl) AvailWidth() int {
	return p.get("availWidth").toInt()
}

func (p *screenImpl) AvailHeight() int {
	return p.get("availHeight").toInt()
}

func (p *screenImpl) Width() int {
	return p.get("width").toInt()
}

func (p *screenImpl) Height() int {
	return p.get("height").toInt()
}

func (p *screenImpl) ColorDepth() uint {
	return p.get("colorDepth").toUint()
}

func (p *screenImpl) PixelDepth() uint {
	return p.get("pixelDepth").toUint()
}

// -------------8<---------------------------------------

type mediaQueryListImpl struct {
	*eventTargetImpl
}

func wrapMediaQueryList(v Value) MediaQueryList {
	if v.valid() {
		return &mediaQueryListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaQueryListImpl) Media() string {
	return p.get("media").toString()
}

func (p *mediaQueryListImpl) Matches() bool {
	return p.get("matches").toBool()
}

func (p *mediaQueryListImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

// -------------8<---------------------------------------

type mediaQueryListEventImpl struct {
	*eventImpl
}

func wrapMediaQueryListEvent(v Value) MediaQueryListEvent {
	if v.valid() {
		return &mediaQueryListEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *mediaQueryListEventImpl) Media() string {
	return p.get("media").toString()
}

func (p *mediaQueryListEventImpl) Matches() bool {
	return p.get("matches").toBool()
}

// -------------8<---------------------------------------

type caretPositionImpl struct {
	Value
}

func wrapCaretPosition(v Value) CaretPosition {
	if v.valid() {
		return &caretPositionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *caretPositionImpl) OffsetNode() Node {
	return wrapAsNode(p.get("offsetNode"))
}

func (p *caretPositionImpl) Offset() uint {
	return p.get("offset").toUint()
}

func (p *caretPositionImpl) ClientRect() DOMRect {
	return wrapDOMRect(p.call("getClientRect"))
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
	if v.valid() {
		return &geometryUtilsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *geometryUtilsImpl) BoxQuads(options ...BoxQuadOptions) []DOMQuad {
	switch len(options) {
	case 0:
		return domQuadArrayToSlice(p.call("getBoxQuads"))
	default:
		return domQuadArrayToSlice(p.call("getBoxQuads", options[0].JSValue()))
	}
}

func (p *geometryUtilsImpl) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.call("convertQuadFromNode", quad.JSValue(), JSValueOf(from)))
	default:
		return wrapDOMQuad(p.call("convertQuadFromNode", quad.JSValue(), JSValueOf(from), options[0].JSValue()))
	}
}

func (p *geometryUtilsImpl) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ...ConvertCoordinateOptions) DOMQuad {
	switch len(options) {
	case 0:
		return wrapDOMQuad(p.call("convertRectFromNode", JSValueOf(rect), JSValueOf(from)))
	default:
		return wrapDOMQuad(p.call("convertRectFromNode", JSValueOf(rect), JSValueOf(from), options[0].JSValue()))
	}
}

func (p *geometryUtilsImpl) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ...ConvertCoordinateOptions) DOMPoint {
	switch len(options) {
	case 0:
		return wrapDOMPoint(p.call("convertPointFromNode", point.JSValue(), JSValueOf(from)))
	default:
		return wrapDOMPoint(p.call("convertPointFromNode", point.JSValue(), JSValueOf(from), options[0].JSValue()))
	}
}

// -------------8<---------------------------------------

type cssPseudoElementImpl struct {
	*eventTargetImpl
}

func wrapCSSPseudoElement(v Value) CSSPseudoElement {
	if v.valid() {
		return &cssPseudoElementImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *cssPseudoElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *cssPseudoElementImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.get("style"))
}

// -------------8<---------------------------------------

type cssPseudoElementListImpl struct {
	Value
}

func wrapCSSPseudoElementList(v Value) CSSPseudoElementList {
	if v.valid() {
		return &cssPseudoElementListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *cssPseudoElementListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *cssPseudoElementListImpl) Item(index uint) CSSPseudoElement {
	return wrapCSSPseudoElement(p.call("item", index))
}

func (p *cssPseudoElementListImpl) ByType(typ string) CSSPseudoElement {
	return wrapCSSPseudoElement(p.call("getByType", typ))
}

// -------------8<---------------------------------------
func NewMediaQueryListEvent(typ string, eventInitDict ...MediaQueryListEventInit) MediaQueryListEvent {
	jsMQLE := jsGlobal.get("MediaQueryListEvent")

	switch len(eventInitDict) {
	case 0:
		return wrapMediaQueryListEvent(jsMQLE.jsNew(typ))
	default:
		return wrapMediaQueryListEvent(jsMQLE.jsNew(typ, eventInitDict[0].JSValue()))
	}
}
