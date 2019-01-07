// +build js,wasm

package wasm

type (
	// https://heycam.github.io/webidl/#BufferSource
	BufferSource interface {
	}

	// https://heycam.github.io/webidl/#ArrayBufferView
	ArrayBufferView interface {
	}

	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer
	ArrayBuffer interface {
		ByteLength() int
		IsView(arg interface{}) bool
		Slice(int, ...int) ArrayBuffer
		ToByteSlice() []byte
	}

	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/DataView
	DataView interface {
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
		// from , of
	}
)

// -------------8<---------------------------------------

func NewArrayBuffer(length int) ArrayBuffer {
	if jsArrayBuffer := jsGlobal.Get("ArrayBuffer"); jsArrayBuffer.Valid() {
		return wrapArrayBuffer(jsArrayBuffer.New(length))
	}
	return nil
}

func NewDataView(buf ArrayBuffer, args ...int) DataView {
	if jsDataView := jsGlobal.Get("DataView"); jsDataView.Valid() {
		switch len(args) {
		case 0:
			return wrapDataView(jsDataView.New(JSValue(buf)))
		case 1:
			return wrapDataView(jsDataView.New(JSValue(buf), args[0]))
		default:
			return wrapDataView(jsDataView.New(JSValue(buf), args[0], args[1]))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type arrayBufferViewImpl struct {
	Value
}

func wrapArrayBufferView(v Value) ArrayBufferView {
	if v.Valid() {
		return &arrayBufferViewImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type arrayBufferImpl struct {
	Value
}

func wrapArrayBuffer(v Value) ArrayBuffer {
	if v.Valid() {
		return &arrayBufferImpl{
			Value: v,
		}
	}
	return nil
}

func (p *arrayBufferImpl) ByteLength() int {
	return p.Get("byteLength").Int()
}

func (p *arrayBufferImpl) IsView(arg interface{}) bool {
	if v := JSValue(arg); v != jsNull {
		return p.Call("isView", v).Bool()
	}

	return false
}

func (p *arrayBufferImpl) Slice(begin int, end ...int) ArrayBuffer {
	switch len(end) {
	case 0:
		return wrapArrayBuffer(p.Call("slice", begin))
	default:
		return wrapArrayBuffer(p.Call("slice", begin, end[0]))
	}
}

func (p *arrayBufferImpl) ToByteSlice() []byte {
	return uint8ArrayToByteSlice(p.Value)
}

// -------------8<---------------------------------------

type dataViewImpl struct {
	Value
}

func wrapDataView(v Value) DataView {
	if v.Valid() {
		return &dataViewImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataViewImpl) Buffer() ArrayBuffer {
	return wrapArrayBuffer(p.Get("buffer"))
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
	Value
}

func wrapBufferSource(v Value) BufferSource {
	if v.Valid() {
		return &bufferSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------
