// +build js,wasm

package wasm

// https://encoding.spec.whatwg.org/#idl-index

type (
	TextDecoderCommon interface {
		Encoding() string
		Fatal() bool
		IgnoreBOM() bool
	}

	TextDecoder interface {
		TextDecoderCommon

		Decode(...interface{}) string
	}

	TextEncoderCommon interface {
		Encoding() string
	}

	TextEncoder interface {
		TextEncoderCommon
		Encode(...string) []byte
	}

	GenericTransformStream interface {
		Readable() ReadableStream
		Writable() WritableStream
	}

	TextDecoderStream interface {
		TextDecoderCommon
		GenericTransformStream
	}

	TextEncoderStream interface {
		TextEncoderCommon
		GenericTransformStream
	}
)

// -------------8<---------------------------------------

type TextDecoderOptions struct {
	Fatal     bool
	IgnoreBOM bool
}

func (p TextDecoderOptions) toJSObject() Value {
	o := jsObject.jsNew()
	o.set("fatal", p.Fatal)
	o.set("ignoreBOM", p.IgnoreBOM)
	return o
}

// -------------8<---------------------------------------

type TextDecodeOptions struct {
	Stream bool
}

func (p TextDecodeOptions) toJSObject() Value {
	o := jsObject.jsNew()
	o.set("stream", p.Stream)
	return o
}
