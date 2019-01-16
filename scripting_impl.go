// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlScriptElementImpl struct {
	*htmlElementImpl
}

func NewHTMLScriptElement() HTMLScriptElement {
	if el := CurrentDocument().CreateElement("script"); el != nil {
		if script, ok := el.(HTMLScriptElement); ok {
			return script
		}
	}
	return nil
}

func wrapHTMLScriptElement(v Value) HTMLScriptElement {
	if v.valid() {
		return &htmlScriptElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlScriptElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlScriptElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlScriptElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlScriptElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlScriptElementImpl) Charset() string {
	return p.get("charset").toString()
}

func (p *htmlScriptElementImpl) SetCharset(charset string) {
	p.set("charset", charset)
}

func (p *htmlScriptElementImpl) Async() bool {
	return p.get("async").toBool()
}

func (p *htmlScriptElementImpl) SetAsync(async bool) {
	p.set("async", async)
}

func (p *htmlScriptElementImpl) Defer() bool {
	return p.get("defer").toBool()
}

func (p *htmlScriptElementImpl) SetDefer(d bool) {
	p.set("defer", d)
}

func (p *htmlScriptElementImpl) CrossOrigin() string {
	return p.get("crossOrigin").toString()
}

func (p *htmlScriptElementImpl) SetCrossOrigin(co string) {
	p.set("crossOrigin", co)
}

func (p *htmlScriptElementImpl) Text() string {
	return p.get("text").toString()
}

func (p *htmlScriptElementImpl) SetText(t string) {
	p.set("text", t)
}

func (p *htmlScriptElementImpl) Nonce() string {
	return p.get("nonce").toString()
}

func (p *htmlScriptElementImpl) SetNonce(n string) {
	p.set("nonce", n)
}

// -------------8<---------------------------------------

type htmlTemplateElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTemplateElement() HTMLTemplateElement {
	if el := CurrentDocument().CreateElement("template"); el != nil {
		if template, ok := el.(HTMLTemplateElement); ok {
			return template
		}
	}
	return nil
}

func wrapHTMLTemplateElement(v Value) HTMLTemplateElement {
	if v.valid() {
		return &htmlTemplateElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTemplateElementImpl) Content() DocumentFragment {
	return wrapDocumentFragment(p.get("content"))
}

// -------------8<---------------------------------------

type htmlCanvasElementImpl struct {
	*htmlElementImpl
}

func NewHTMLCanvasElement(size ...uint) HTMLCanvasElement {
	if el := CurrentDocument().CreateElement("canvas"); el != nil {
		if canvas, ok := el.(HTMLCanvasElement); ok {
			if len(size) == 2 {
				canvas.SetWidth(size[0])
				canvas.SetHeight(size[1])
			}
			return canvas
		}
	}
	return nil
}

func wrapHTMLCanvasElement(v Value) HTMLCanvasElement {
	if v.valid() {
		return &htmlCanvasElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlCanvasElementImpl) Width() uint {
	return p.get("width").toUint()
}

func (p *htmlCanvasElementImpl) SetWidth(w uint) {
	p.set("width", w)
}

func (p *htmlCanvasElementImpl) Height() uint {
	return p.get("height").toUint()
}

func (p *htmlCanvasElementImpl) SetHeight(h uint) {
	p.set("height", h)
}

func (p *htmlCanvasElementImpl) Context2D(settings ...CanvasRenderingContext2DSettings) CanvasRenderingContext2D {
	switch len(settings) {
	case 0:
		return wrapCanvasRenderingContext2D(p.call("getContext", "2d"))
	default:
		return wrapCanvasRenderingContext2D(p.call("getContext", "2d", settings[0].JSValue()))
	}
}

func (p *htmlCanvasElementImpl) ContextWebGL(attrs ...WebGLContextAttributes) WebGLRenderingContext {
	var v Value

	switch len(attrs) {
	case 0:
		v = p.call("getContext", "webgl")
		if !v.valid() {
			v = p.call("getContext", "experimental-webgl")
		}

	default:
		v = p.call("getContext", "webgl", attrs[0].JSValue())
		if !v.valid() {
			v = p.call("getContext", "experimental-webgl", attrs[0].JSValue())
		}
	}

	if !v.valid() {
		return nil
	}

	return wrapWebGLRenderingContext(v)
}

// TODO: removed from standart
// https://github.com/whatwg/html/commit/2cfb8e3f03d3166842d2ad0f661459d26e2a40eb
func (p *htmlCanvasElementImpl) ProbablySupportsContext(ctxId string, args ...interface{}) bool {
	// TODO
	return p.call("probablySupportsContext", ctxId).toBool()
}

func (p *htmlCanvasElementImpl) ToDataURL(args ...interface{}) string {
	switch len(args) {
	case 1:
		if typ, ok := args[0].(string); ok { // type
			return p.call("toDataURL", typ).toString()
		}
	case 2:
		if typ, ok := args[0].(string); ok { // type
			if quality, ok := args[1].(float64); ok { //quality
				return p.call("toDataURL", typ, quality).toString()
			}
		}
	}

	return p.call("toDataURL").toString()
}

func (p *htmlCanvasElementImpl) ToBlob(cb BlobCallback, args ...interface{}) {
	switch len(args) {
	case 1:
		if typ, ok := args[0].(string); ok { // type
			p.call("toBlob", cb.jsCallback(), typ)
		}
	case 2:
		if typ, ok := args[0].(string); ok { // type
			if quality, ok := args[1].(float64); ok { //quality
				p.call("toBlob", cb.jsCallback(), typ, quality)
			}
		}
	default:
		p.call("toBlob", cb.jsCallback())
	}
}

// -------------8<---------------------------------------
