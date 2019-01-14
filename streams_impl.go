// +build js,wasm

package wasm

// -------------8<---------------------------------------

type readableStreamImpl struct {
	Value
}

func wrapReadableStream(v Value) ReadableStream {
	if v.valid() {
		return &readableStreamImpl{
			Value: v,
		}
	}
	return nil
}

func (p *readableStreamImpl) Locked() bool {
	return p.get("locked").toBool()
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
	Value
}

func wrapWritableStream(v Value) WritableStream {
	if v.valid() {
		return &writableStreamImpl{
			Value: v,
		}
	}
	return nil
}

func (p *writableStreamImpl) Locked() bool {
	return p.get("locked").toBool()
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
