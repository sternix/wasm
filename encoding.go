// +build js,wasm

package wasm

import (
	"syscall/js"
)

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
	Fatal     bool `json:"fatal"`
	IgnoreBOM bool `json:"ignoreBOM"`
}

func (p TextDecoderOptions) toDict() js.Value {
	o := jsObject.New()
	o.Set("fatal", p.Fatal)
	o.Set("ignoreBOM", p.IgnoreBOM)
	return o
}

// -------------8<---------------------------------------

type TextDecodeOptions struct {
	Stream bool `json:"stream"`
}

func (p TextDecodeOptions) toDict() js.Value {
	o := jsObject.New()
	o.Set("stream", p.Stream)
	return o
}
