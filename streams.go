// +build js,wasm

package wasm

import (
	"syscall/js"
)

// EXPERIMENTAL

// https://streams.spec.whatwg.org/

// for compiled successfully
// TODO: remove when types completed
type Promise interface{}

type (
	ReadableStream interface {
		Locked() bool
		Cancel(string) Promise // Promise<reason>
		Reader() GenericReader
		PipeThrough(TransformStream, ...PipeToOptions) ReadableStream
		PipeTo(WritableStream, ...PipeToOptions) Promise // Promise<void>
		Tee()
	}

	GenericReader interface {
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

	WritableStream interface {
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
	// TODO
	return nil
}

// -------------8<---------------------------------------

type PipeToOptions struct {
	PreventClose  bool
	PreventAbort  bool
	PreventCancel bool
}

func (p PipeToOptions) toDict() js.Value {
	o := jsObject.New()
	o.Set("preventClose", p.PreventClose)
	o.Set("preventAbort", p.PreventAbort)
	o.Set("preventCancel", p.PreventCancel)
	return o
}

// -------------8<---------------------------------------

type TransformStream struct {
	Readable ReadableStream
	Writable WritableStream
}

func (p TransformStream) toDict() js.Value {
	o := jsObject.New()
	o.Set("writable", JSValue(p.Readable))
	o.Set("readable", JSValue(p.Writable))
	return o
}
