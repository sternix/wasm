// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://heycam.github.io/webidl/#BufferSource
	BufferSource interface {
		js.Wrapper
	}

	// https://heycam.github.io/webidl/#ArrayBufferView
	ArrayBufferView interface {
		js.Wrapper
	}

	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer
	ArrayBuffer interface {
		js.Wrapper

		ByteLength() int
		IsView(arg interface{}) bool
		Slice(int, ...int) ArrayBuffer
		ToByteSlice() []byte
	}

	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/DataView
	DataView interface {
		js.Wrapper

		Buffer() ArrayBuffer
		ByteLength() int
		ByteOffset() int
		Int8() int8
		Uint8() uint8
		Int16() int16
		Uint16() uint16
		Int32() int32
		Uint32() uint32
		Float32() float32
		Float64() float64
		SetInt8(v int8)
		SetUint8(v uint8)
		SetInt16(v int16)
		SetUint16(v uint16)
		SetInt32(v int32)
		SetUint32(v uint32)
		SetFloat32(v float32)
		SetFloat64(v float64)
	}

	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Uint8ClampedArray
	Uint8ClampedArray interface {
		js.Wrapper
		// from , of
	}
)

// -------------8<---------------------------------------

func NewArrayBuffer(length int) ArrayBuffer {
	jsArrayBuffer := js.Global().Get("ArrayBuffer")
	if isNil(jsArrayBuffer) {
		return nil
	}

	return newArrayBuffer(jsArrayBuffer.New(length))
}

func NewDataView(buf ArrayBuffer, args ...int) DataView {
	jsDataView := js.Global().Get("DataView")
	if isNil(jsDataView) {
		return nil
	}

	switch len(args) {
	case 0:
		return newDataView(jsDataView.New(buf.JSValue()))
	case 1:
		return newDataView(jsDataView.New(buf.JSValue(), args[0]))
	default:
		return newDataView(jsDataView.New(buf.JSValue(), args[0], args[1]))
	}
}

// -------------8<---------------------------------------

type arrayBufferViewImpl struct {
	js.Value
}

func newArrayBufferView(v js.Value) ArrayBufferView {
	if isNil(v) {
		return nil
	}

	return &arrayBufferViewImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type arrayBufferImpl struct {
	js.Value
}

func newArrayBuffer(v js.Value) ArrayBuffer {
	if isNil(v) {
		return nil
	}

	return &arrayBufferImpl{
		Value: v,
	}
}

func (p *arrayBufferImpl) ByteLength() int {
	return p.Get("byteLength").Int()
}

func (p *arrayBufferImpl) IsView(arg interface{}) bool {
	if jv, ok := arg.(js.Wrapper); ok {
		return p.Call("isView", jv.JSValue()).Bool()
	}

	return false
}

func (p *arrayBufferImpl) Slice(begin int, end ...int) ArrayBuffer {
	switch len(end) {
	case 0:
		return newArrayBuffer(p.Call("slice", begin))
	default:
		return newArrayBuffer(p.Call("slice", begin, end[0]))
	}
}

func (p *arrayBufferImpl) ToByteSlice() []byte {
	return uint8ArrayToByteSlice(p.Value)
}

// -------------8<---------------------------------------

type dataViewImpl struct {
	js.Value
}

func newDataView(v js.Value) DataView {
	if isNil(v) {
		return nil
	}

	return &dataViewImpl{
		Value: v,
	}
}

func (p *dataViewImpl) Buffer() ArrayBuffer {
	return newArrayBuffer(p.Get("buffer"))
}

func (p *dataViewImpl) ByteLength() int {
	return p.Get("byteLength").Int()
}

func (p *dataViewImpl) ByteOffset() int {
	return p.Get("byteOffset").Int()
}

func (p *dataViewImpl) Int8() int8 {
	return int8(p.Call("getInt8").Int())
}

func (p *dataViewImpl) Uint8() uint8 {
	return uint8(p.Call("getUint8").Int())
}

func (p *dataViewImpl) Int16() int16 {
	return int16(p.Call("getInt16").Int())
}

func (p *dataViewImpl) Uint16() uint16 {
	return uint16(p.Call("getUint16").Int())
}

func (p *dataViewImpl) Int32() int32 {
	return int32(p.Call("getInt32").Int())
}

func (p *dataViewImpl) Uint32() uint32 {
	return uint32(p.Call("getUint32").Int())
}

func (p *dataViewImpl) Float32() float32 {
	return float32(p.Call("getFloat32").Float())
}

func (p *dataViewImpl) Float64() float64 {
	return p.Call("getFloat64").Float()
}

func (p *dataViewImpl) SetInt8(v int8) {
	p.Call("setInt8", v)
}

func (p *dataViewImpl) SetUint8(v uint8) {
	p.Call("setUint8", v)
}

func (p *dataViewImpl) SetInt16(v int16) {
	p.Call("setInt16", v)
}

func (p *dataViewImpl) SetUint16(v uint16) {
	p.Call("setUint16", v)
}

func (p *dataViewImpl) SetInt32(v int32) {
	p.Call("setInt32", v)
}

func (p *dataViewImpl) SetUint32(v uint32) {
	p.Call("setUint32", v)
}

func (p *dataViewImpl) SetFloat32(v float32) {
	p.Call("setFloat32", v)
}

func (p *dataViewImpl) SetFloat64(v float64) {
	p.Call("setFloat64", v)
}

// -------------8<---------------------------------------

type bufferSourceImpl struct {
	js.Value
}

func newBufferSource(v js.Value) BufferSource {
	if isNil(v) {
		return nil
	}

	return &bufferSourceImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------
