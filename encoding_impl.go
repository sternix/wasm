// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type textDecoderCommonImpl struct {
	js.Value
}

func wrapTextDecoderCommon(v js.Value) TextDecoderCommon {
	if p := newTextDecoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextDecoderCommonImpl(v js.Value) *textDecoderCommonImpl {
	if isNil(v) {
		return nil
	}

	return &textDecoderCommonImpl{
		Value: v,
	}
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

func wrapTextDecoder(v js.Value) TextDecoder {
	if isNil(v) {
		return nil
	}

	return &textDecoderImpl{
		textDecoderCommonImpl: newTextDecoderCommonImpl(v),
	}
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
	js.Value
}

func wrapTextEncoderCommon(v js.Value) TextEncoderCommon {
	if p := newTextEncoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextEncoderCommonImpl(v js.Value) *textEncoderCommonImpl {
	if isNil(v) {
		return nil
	}

	return &textEncoderCommonImpl{
		Value: v,
	}
}

func (p *textEncoderCommonImpl) Encoding() string {
	return p.Get("encoding").String()
}

// -------------8<---------------------------------------

type textEncoderImpl struct {
	*textEncoderCommonImpl
}

func wrapTextEncoder(v js.Value) TextEncoder {
	if isNil(v) {
		return nil
	}

	return &textEncoderImpl{
		textEncoderCommonImpl: newTextEncoderCommonImpl(v),
	}
}

func (p *textEncoderImpl) Encode(input ...string) []byte {
	if len(input) > 0 {
		return uint8ArrayToByteSlice(p.Call("encode", input[0]))
	}

	return uint8ArrayToByteSlice(p.Call("encode"))
}

// -------------8<---------------------------------------

type genericTransformStreamImpl struct {
	js.Value
}

func wrapGenericTransformStream(v js.Value) GenericTransformStream {
	if p := newGenericTransformStreamImpl(v); p != nil {
		return p
	}
	return nil
}

func newGenericTransformStreamImpl(v js.Value) *genericTransformStreamImpl {
	if isNil(v) {
		return nil
	}

	return &genericTransformStreamImpl{
		Value: v,
	}
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

func wrapTextDecoderStream(v js.Value) TextDecoderStream {
	if isNil(v) {
		return nil
	}

	return &textDecoderStreamImpl{
		textDecoderCommonImpl:      newTextDecoderCommonImpl(v),
		genericTransformStreamImpl: newGenericTransformStreamImpl(v),
	}
}

// -------------8<---------------------------------------

type textEncoderStreamImpl struct {
	*textEncoderCommonImpl
	*genericTransformStreamImpl
}

func wrapTextEncoderStream(v js.Value) TextEncoderStream {
	if isNil(v) {
		return nil
	}

	return &textEncoderStreamImpl{
		textEncoderCommonImpl:      newTextEncoderCommonImpl(v),
		genericTransformStreamImpl: newGenericTransformStreamImpl(v),
	}
}

// -------------8<---------------------------------------

func NewTextDecoderStream(args ...interface{}) TextDecoderStream {
	jsDecStream := js.Global().Get("TextDecoderStream")
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
	jsTextDecoder := js.Global().Get("TextDecoder")
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
	return wrapTextEncoderStream(js.Global().Get("TextEncoderStream").New())
}

func NewTextEncoder() TextEncoder {
	return wrapTextEncoder(js.Global().Get("TextEncoder").New())
}
