// +build js,wasm

package wasm

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
		//Context(string, ...interface{}) RenderingContext
		Context2D(alpha ...bool) CanvasRenderingContext2D
		// TODO
		//ContextWebGl()
		ProbablySupportsContext(string, ...interface{}) bool
		ToDataURL(...interface{}) string
		ToBlob(BlobCallback, ...interface{})
	}
)
