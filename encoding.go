// +build js,wasm

package wasm

// https://encoding.spec.whatwg.org/#idl-index

type (
	TextDecoderCommon interface {
		Encoding() string
		Fatal() bool
		IgnoreBOM() bool
	}

	TextDecoderOptions struct {
		Fatal     bool `json:"fatal"`
		IgnoreBOM bool `json:"ignoreBOM"`
	}

	TextDecodeOptions struct {
		Stream bool `json:"stream"`
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
