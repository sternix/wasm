// +build js,wasm

package wasm

// -------------8<---------------------------------------

type textDecoderCommonImpl struct {
	Value
}

func wrapTextDecoderCommon(v Value) TextDecoderCommon {
	if p := newTextDecoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextDecoderCommonImpl(v Value) *textDecoderCommonImpl {
	if v.valid() {
		return &textDecoderCommonImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textDecoderCommonImpl) Encoding() string {
	return p.get("encoding").toString()
}

func (p *textDecoderCommonImpl) Fatal() bool {
	return p.get("fatal").toBool()
}

func (p *textDecoderCommonImpl) IgnoreBOM() bool {
	return p.get("ignoreBOM").toBool()
}

// -------------8<---------------------------------------

type textDecoderImpl struct {
	*textDecoderCommonImpl
}

func wrapTextDecoder(v Value) TextDecoder {
	if v.valid() {
		return &textDecoderImpl{
			textDecoderCommonImpl: newTextDecoderCommonImpl(v),
		}
	}
	return nil
}

func (p *textDecoderImpl) Decode(args ...interface{}) string {
	switch len(args) {
	case 1:
		if input, ok := args[0].(BufferSource); ok {
			return p.call("decode", JSValue(input)).toString()
		}
	case 2:
		if input, ok := args[0].(BufferSource); ok {
			if options, ok := args[1].(TextDecodeOptions); ok {
				return p.call("decode", JSValue(input), options.toJSObject()).toString()
			}
		}
	}

	return p.call("decode").toString()
}

// -------------8<---------------------------------------

type textEncoderCommonImpl struct {
	Value
}

func wrapTextEncoderCommon(v Value) TextEncoderCommon {
	if p := newTextEncoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextEncoderCommonImpl(v Value) *textEncoderCommonImpl {
	if v.valid() {
		return &textEncoderCommonImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textEncoderCommonImpl) Encoding() string {
	return p.get("encoding").toString()
}

// -------------8<---------------------------------------

type textEncoderImpl struct {
	*textEncoderCommonImpl
}

func wrapTextEncoder(v Value) TextEncoder {
	if v.valid() {
		return &textEncoderImpl{
			textEncoderCommonImpl: newTextEncoderCommonImpl(v),
		}
	}
	return nil
}

func (p *textEncoderImpl) Encode(input ...string) []byte {
	if len(input) > 0 {
		return uint8ArrayToByteSlice(p.call("encode", input[0]))
	}

	return uint8ArrayToByteSlice(p.call("encode"))
}

// -------------8<---------------------------------------

type genericTransformStreamImpl struct {
	Value
}

func wrapGenericTransformStream(v Value) GenericTransformStream {
	if p := newGenericTransformStreamImpl(v); p != nil {
		return p
	}
	return nil
}

func newGenericTransformStreamImpl(v Value) *genericTransformStreamImpl {
	if v.valid() {
		return &genericTransformStreamImpl{
			Value: v,
		}
	}
	return nil
}

func (p *genericTransformStreamImpl) Readable() ReadableStream {
	return wrapReadableStream(p.get("readable"))
}

func (p *genericTransformStreamImpl) Writable() WritableStream {
	return wrapWritableStream(p.get("writable"))
}

// -------------8<---------------------------------------

type textDecoderStreamImpl struct {
	*textDecoderCommonImpl
	*genericTransformStreamImpl
}

func wrapTextDecoderStream(v Value) TextDecoderStream {
	if v.valid() {
		return &textDecoderStreamImpl{
			textDecoderCommonImpl:      newTextDecoderCommonImpl(v),
			genericTransformStreamImpl: newGenericTransformStreamImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type textEncoderStreamImpl struct {
	*textEncoderCommonImpl
	*genericTransformStreamImpl
}

func wrapTextEncoderStream(v Value) TextEncoderStream {
	if v.valid() {
		return &textEncoderStreamImpl{
			textEncoderCommonImpl:      newTextEncoderCommonImpl(v),
			genericTransformStreamImpl: newGenericTransformStreamImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

func NewTextDecoderStream(args ...interface{}) TextDecoderStream {
	jsDecStream := jsGlobal.get("TextDecoderStream")
	switch len(args) {
	case 1:
		if label, ok := args[0].(string); ok {
			return wrapTextDecoderStream(jsDecStream.jsNew(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return wrapTextDecoderStream(jsDecStream.jsNew(label, options.toJSObject()))
			}
		}
	}

	return wrapTextDecoderStream(jsDecStream.jsNew())
}

func NewTextDecoder(args ...interface{}) TextDecoder {
	jsTextDecoder := jsGlobal.get("TextDecoder")
	switch len(args) {
	case 1:
		if label, ok := args[0].(string); ok {
			return wrapTextDecoder(jsTextDecoder.jsNew(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return wrapTextDecoder(jsTextDecoder.jsNew(label, options.toJSObject()))
			}
		}
	}

	return wrapTextDecoder(jsTextDecoder.jsNew())
}

func NewTextEncoderStream() TextEncoderStream {
	return wrapTextEncoderStream(jsGlobal.get("TextEncoderStream").jsNew())
}

func NewTextEncoder() TextEncoder {
	return wrapTextEncoder(jsGlobal.get("TextEncoder").jsNew())
}
