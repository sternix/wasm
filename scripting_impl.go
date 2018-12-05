// +build js,wasm

package wasm

import (
	"syscall/js"
)

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

func newHTMLScriptElement(v js.Value) HTMLScriptElement {
	if isNil(v) {
		return nil
	}

	return &htmlScriptElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlScriptElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlScriptElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlScriptElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlScriptElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlScriptElementImpl) Charset() string {
	return p.Get("charset").String()
}

func (p *htmlScriptElementImpl) SetCharset(charset string) {
	p.Set("charset", charset)
}

func (p *htmlScriptElementImpl) Async() bool {
	return p.Get("async").Bool()
}

func (p *htmlScriptElementImpl) SetAsync(async bool) {
	p.Set("async", async)
}

func (p *htmlScriptElementImpl) Defer() bool {
	return p.Get("defer").Bool()
}

func (p *htmlScriptElementImpl) SetDefer(d bool) {
	p.Set("defer", d)
}

func (p *htmlScriptElementImpl) CrossOrigin() string {
	return p.Get("crossOrigin").String()
}

func (p *htmlScriptElementImpl) SetCrossOrigin(co string) {
	p.Set("crossOrigin", co)
}

func (p *htmlScriptElementImpl) Text() string {
	return p.Get("text").String()
}

func (p *htmlScriptElementImpl) SetText(t string) {
	p.Set("text", t)
}

func (p *htmlScriptElementImpl) Nonce() string {
	return p.Get("nonce").String()
}

func (p *htmlScriptElementImpl) SetNonce(n string) {
	p.Set("nonce", n)
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

func newHTMLTemplateElement(v js.Value) HTMLTemplateElement {
	if isNil(v) {
		return nil
	}

	return &htmlTemplateElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTemplateElementImpl) Content() DocumentFragment {
	return newDocumentFragment(p.Get("content"))
}

// -------------8<---------------------------------------

type htmlCanvasElementImpl struct {
	*htmlElementImpl
}

func NewHTMLCanvasElement(size ...int) HTMLCanvasElement {
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

func newHTMLCanvasElement(v js.Value) HTMLCanvasElement {
	if isNil(v) {
		return nil
	}

	return &htmlCanvasElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlCanvasElementImpl) Width() int {
	return p.Get("width").Int()
}

func (p *htmlCanvasElementImpl) SetWidth(w int) {
	p.Set("width", w)
}

func (p *htmlCanvasElementImpl) Height() int {
	return p.Get("height").Int()
}

func (p *htmlCanvasElementImpl) SetHeight(h int) {
	p.Set("height", h)
}

/*
func (p *htmlCanvasElementImpl) Context(ctxId string, args ...interface{}) RenderingContext {
	// TODO
	return nil
}
*/

func (p *htmlCanvasElementImpl) Context2D(alpha ...bool) CanvasRenderingContext2D {
	// TODO: alpha param omitted
	return newCanvasRenderingContext2D(p.Call("getContext", "2d"))
}

func (p *htmlCanvasElementImpl) ContextWebGL(attrs ...WebGLContextAttributes) WebGLRenderingContext {
	var v js.Value

	switch len(attrs) {
	case 0:
		v = p.Call("getContext", "webgl")
		if v == js.Undefined() {
			v = p.Call("getContext", "experimental-webgl")
		}

		if isNil(v) {
			return nil
		}
	default:
		v = p.Call("getContext", "webgl", attrs[0].toDict())
		if v == js.Undefined() {
			v = p.Call("getContext", "experimental-webgl", attrs[0].toDict())
		}

		if isNil(v) {
			return nil
		}
	}

	return newWebGLRenderingContext(v)
}

// TODO: removed from standart
// https://github.com/whatwg/html/commit/2cfb8e3f03d3166842d2ad0f661459d26e2a40eb
func (p *htmlCanvasElementImpl) ProbablySupportsContext(ctxId string, args ...interface{}) bool {
	// TODO
	return p.Call("probablySupportsContext", ctxId).Bool()
}

func (p *htmlCanvasElementImpl) ToDataURL(args ...interface{}) string {
	switch len(args) {
	case 1:
		if typ, ok := args[0].(string); ok { // type
			return p.Call("toDataURL", typ).String()
		}
	case 2:
		if typ, ok := args[0].(string); ok { // type
			if quality, ok := args[1].(float64); ok { //quality
				return p.Call("toDataURL", typ, quality).String()
			}
		}
	}

	return p.Call("toDataURL").String()
}

func (p *htmlCanvasElementImpl) ToBlob(cb BlobCallback, args ...interface{}) {
	switch len(args) {
	case 1:
		if typ, ok := args[0].(string); ok { // type
			p.Call("toBlob", cb.jsCallback(), typ)
		}
	case 2:
		if typ, ok := args[0].(string); ok { // type
			if quality, ok := args[1].(float64); ok { //quality
				p.Call("toBlob", cb.jsCallback(), typ, quality)
			}
		}
	default:
		p.Call("toBlob", cb.jsCallback())
	}
}

// -------------8<---------------------------------------
