// +build js,wasm

package wasm

// EXPERIMENTAL

// https://streams.spec.whatwg.org/

type (
	ReadableStream interface {
		Object

		Locked() bool
		Cancel(string) Promise // Promise<reason>
		Reader() GenericReader
		PipeThrough(TransformStream, ...PipeToOptions) ReadableStream
		PipeTo(WritableStream, ...PipeToOptions) Promise // Promise<void>
		Tee()
	}

	PipeToOptions struct {
		PreventClose  bool `json:"preventClose"`
		PreventAbort  bool `json:"preventAbort"`
		PreventCancel bool `json:"preventCancel"`
	}

	GenericReader interface {
		Object

		Closed() Promise
		Cancel(string) Promise // Promise<reason>
		Read() Promise         // Promise<result>
		ReleaseLock()
	}

	ReadableStreamDefaultReader interface {
		GenericReader
	}

	ReadableStreamBYOBReader interface {
		GenericReader
	}

	TransformStream struct {
		Readable ReadableStream `json:"writable"`
		Writable WritableStream `json:"readable"`
	}

	WritableStream interface {
		Object

		Locked() bool
		Abort(string) Promise // Promise<reason>
		Writer() WritableStreamDefaultWriter
	}

	WritableStreamDefaultWriter interface {
		Closed() Promise
		DesiredSize() int
		Ready() Promise
		Abort() Promise // Promise<reason>
		Close() Promise // Promise<undefined>
		ReleaseLock()
		Write([]byte) Promise
	}
)

func NewReadableStream() ReadableStream {
	return nil
}
