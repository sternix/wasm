// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type textDecoderCommonImpl struct {
	*objectImpl
}

func newTextDecoderCommon(v js.Value) TextDecoderCommon {
	if p := newTextDecoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

// KEEP
func newTextDecoderCommonImpl(v js.Value) *textDecoderCommonImpl {
	if isNil(v) {
		return nil
	}

	return &textDecoderCommonImpl{
		objectImpl: newObjectImpl(v),
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

func newTextDecoder(v js.Value) TextDecoder {
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
			return p.Call("decode", input.JSValue()).String()
		}
	case 2:
		if input, ok := args[0].(BufferSource); ok {
			if options, ok := args[1].(TextDecodeOptions); ok {
				return p.Call("decode", input.JSValue(), toJSONObject(options)).String()
			}
		}
	}

	return p.Call("decode").String()
}

// -------------8<---------------------------------------

type textEncoderCommonImpl struct {
	*objectImpl
}

func newTextEncoderCommon(v js.Value) TextEncoderCommon {
	if p := newTextEncoderCommonImpl(v); p != nil {
		return p
	}
	return nil
}

// KEEP
func newTextEncoderCommonImpl(v js.Value) *textEncoderCommonImpl {
	if isNil(v) {
		return nil
	}

	return &textEncoderCommonImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *textEncoderCommonImpl) Encoding() string {
	return p.Get("encoding").String()
}

// -------------8<---------------------------------------

type textEncoderImpl struct {
	*textEncoderCommonImpl
}

func newTextEncoder(v js.Value) TextEncoder {
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
	*objectImpl
}

func newGenericTransformStream(v js.Value) GenericTransformStream {
	if p := newGenericTransformStreamImpl(v); p != nil {
		return p
	}
	return nil
}

// KEEP
func newGenericTransformStreamImpl(v js.Value) *genericTransformStreamImpl {
	if isNil(v) {
		return nil
	}

	return &genericTransformStreamImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *genericTransformStreamImpl) Readable() ReadableStream {
	return newReadableStream(p.Get("readable"))
}

func (p *genericTransformStreamImpl) Writable() WritableStream {
	return newWritableStream(p.Get("writable"))
}

// -------------8<---------------------------------------

type textDecoderStreamImpl struct {
	*textDecoderCommonImpl
	*genericTransformStreamImpl
}

func newTextDecoderStream(v js.Value) TextDecoderStream {
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

func newTextEncoderStream(v js.Value) TextEncoderStream {
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
			return newTextDecoderStream(jsDecStream.New(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return newTextDecoderStream(jsDecStream.New(label, toJSONObject(options)))
			}
		}
	}

	return newTextDecoderStream(jsDecStream.New())
}

func NewTextDecoder(args ...interface{}) TextDecoder {
	jsTextDecoder := js.Global().Get("TextDecoder")
	switch len(args) {
	case 1:
		if label, ok := args[0].(string); ok {
			return newTextDecoder(jsTextDecoder.New(label))
		}
	case 2:
		if label, ok := args[0].(string); ok {
			if options, ok := args[1].(TextDecoderOptions); ok {
				return newTextDecoder(jsTextDecoder.New(label, toJSONObject(options)))
			}
		}
	}

	return newTextDecoder(jsTextDecoder.New())
}

func NewTextEncoderStream() TextEncoderStream {
	return newTextEncoderStream(js.Global().Get("TextEncoderStream").New())
}

func NewTextEncoder() TextEncoder {
	return newTextEncoder(js.Global().Get("TextEncoder").New())
}
