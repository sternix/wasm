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
	if v.Valid() {
		return &textDecoderCommonImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textDecoderCommonImpl) Encoding() string {
	return p.Get("encoding").String()
}

func (p *textDecoderCommonImpl) Fatal() bool {
	return p.Get("fatal").Bool()
}

func (p *textDecoderCommonImpl) IgnoreBOM() bool {
	return p.Get("ignoreBOM").Bool()
}

// -------------8<---------------------------------------

type textDecoderImpl struct {
	*textDecoderCommonImpl
}

func wrapTextDecoder(v Value) TextDecoder {
	if v.Valid() {
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
			return p.Call("decode", JSValue(input)).String()
		}
	case 2:
		if input, ok := args[0].(BufferSource); ok {
			if options, ok := args[1].(TextDecodeOptions); ok {
				return p.Call("decode", JSValue(input), options.toJSObject()).String()
			}
		}
	}

	return p.Call("decode").String()
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
	if v.Valid() {
		return &textEncoderCommonImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textEncoderCommonImpl) Encoding() string {
	return p.Get("encoding").String()
}

// -------------8<---------------------------------------

type textEncoderImpl struct {
	*textEncoderCommonImpl
}

func wrapTextEncoder(v Value) TextEncoder {
	if v.Valid() {
		return &textEncoderImpl{
			textEncoderCommonImpl: newTextEncoderCommonImpl(v),
		}
	}
	return nil
}

func (p *textEncoderImpl) Encode(input ...string) []byte {
	if len(input) > 0 {
		return uint8ArrayToByteSlice(p.Call("encode", input[0]))
	}

	return uint8ArrayToByteSlice(p.Call("encode"))
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
	if v.Valid() {
		return &genericTransformStreamImpl{
			Value: v,
		}
	}
	return nil
}

func (p *genericTransformStreamImpl) Readable() ReadableStream {
	return wrapReadableStream(p.Get("readable"))
}

func (p *genericTransformStreamImpl) Writable() WritableStream {
	return wrapWritableStream(p.Get("writable"))
}

// -------------8<---------------------------------------

type textDecoderStreamImpl struct {
	*textDecoderCommonImpl
	*genericTransformStreamImpl
}

func wrapTextDecoderStream(v Value) TextDecoderStream {
	if v.Valid() {
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
	if v.Valid() {
		return &textEncoderStreamImpl{
			textEncoderCommonImpl:      newTextEncoderCommonImpl(v),
			genericTransformStreamImpl: newGenericTransformStreamImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

func NewTextDecoderStream(args ...interface{}) TextDecoderStream {
	jsDecStream := jsGlobal.Get("TextDecoderStream")
	switch len(args) {
	case 1:
		if label, ok := args[0].(string); ok {
			return wrapTextDecoderStream(jsDecStream.New(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return wrapTextDecoderStream(jsDecStream.New(label, options.toJSObject()))
			}
		}
	}

	return wrapTextDecoderStream(jsDecStream.New())
}

func NewTextDecoder(args ...interface{}) TextDecoder {
	jsTextDecoder := jsGlobal.Get("TextDecoder")
	switch len(args) {
	case 1:
		if label, ok := args[0].(string); ok {
			return wrapTextDecoder(jsTextDecoder.New(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return wrapTextDecoder(jsTextDecoder.New(label, options.toJSObject()))
			}
		}
	}

	return wrapTextDecoder(jsTextDecoder.New())
}

func NewTextEncoderStream() TextEncoderStream {
	return wrapTextEncoderStream(jsGlobal.Get("TextEncoderStream").New())
}

func NewTextEncoder() TextEncoder {
	return wrapTextEncoder(jsGlobal.Get("TextEncoder").New())
}
