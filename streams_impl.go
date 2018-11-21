// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type readableStreamImpl struct {
	js.Value
}

func newReadableStream(v js.Value) ReadableStream {
	if isNil(v) {
		return nil
	}

	return &readableStreamImpl{
		Value: v,
	}
}

func (p *readableStreamImpl) Locked() bool {
	return p.Get("locked").Bool()
}

func (p *readableStreamImpl) Cancel(reason string) Promise { // Promise<reason>
	// TODO
	return nil
}

func (p *readableStreamImpl) Reader() GenericReader {
	// TODO
	return nil
}

func (p *readableStreamImpl) PipeThrough(TransformStream, ...PipeToOptions) ReadableStream {
	// TODO
	return nil
}

func (p *readableStreamImpl) PipeTo(WritableStream, ...PipeToOptions) Promise { // Promise<void>
	// TODO
	return nil
}

func (p *readableStreamImpl) Tee() {
	// TODO
}

// -------------8<---------------------------------------

type writableStreamImpl struct {
	js.Value
}

func newWritableStream(v js.Value) WritableStream {
	if isNil(v) {
		return nil
	}

	return &writableStreamImpl{
		Value: v,
	}
}

func (p *writableStreamImpl) Locked() bool {
	return p.Get("locked").Bool()
}

func (p *writableStreamImpl) Abort(string) Promise { // Promise<reason>
	// TODO
	return nil
}

func (p *writableStreamImpl) Writer() WritableStreamDefaultWriter {
	// TODO
	return nil
}

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
