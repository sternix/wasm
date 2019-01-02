// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://www.w3.org/TR/html52/semantics-scripting.html#htmlscriptelement
	HTMLScriptElement interface {
		HTMLElement

		Src() string
		SetSrc(string)
		Type() string
		SetType(string)
		Charset() string
		SetCharset(string)
		Async() bool
		SetAsync(bool)
		Defer() bool
		SetDefer(bool)
		CrossOrigin() string
		SetCrossOrigin(string)
		Text() string
		SetText(string)
		Nonce() string
		SetNonce(string)
	}

	// https://www.w3.org/TR/html52/semantics-scripting.html#htmltemplateelement
	HTMLTemplateElement interface {
		HTMLElement

		Content() DocumentFragment
	}

	//(CanvasRenderingContext2D or WebGLRenderingContext or WebGL2RenderingContext or ImageBitmapRenderingContext)
	RenderingContext interface{}

	// https://www.w3.org/TR/html52/semantics-scripting.html#htmlcanvaselement
	HTMLCanvasElement interface {
		HTMLElement
		Width() int
		SetWidth(int)
		Height() int
		SetHeight(int)
		Context2D(...CanvasRenderingContext2DSettings) CanvasRenderingContext2D
		ContextWebGL(attrs ...WebGLContextAttributes) WebGLRenderingContext
		ProbablySupportsContext(string, ...interface{}) bool
		ToDataURL(...interface{}) string
		ToBlob(BlobCallback, ...interface{})
	}
)

type CanvasRenderingContext2DSettings struct {
	Alpha bool // default true
}

func (p CanvasRenderingContext2DSettings) toJSObject() js.Value {
	o := jsObject.New()
	o.Set("alpha", p.Alpha)
	return o
}
