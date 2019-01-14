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
	if jsArrayBuffer := jsGlobal.get("ArrayBuffer"); jsArrayBuffer.valid() {
		return wrapArrayBuffer(jsArrayBuffer.jsNew(length))
	}
	return nil
}

func NewDataView(buf ArrayBuffer, args ...int) DataView {
	if jsDataView := jsGlobal.get("DataView"); jsDataView.valid() {
		switch len(args) {
		case 0:
			return wrapDataView(jsDataView.jsNew(JSValue(buf)))
		case 1:
			return wrapDataView(jsDataView.jsNew(JSValue(buf), args[0]))
		default:
			return wrapDataView(jsDataView.jsNew(JSValue(buf), args[0], args[1]))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type arrayBufferViewImpl struct {
	Value
}

func wrapArrayBufferView(v Value) ArrayBufferView {
	if v.valid() {
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
	if v.valid() {
		return &arrayBufferImpl{
			Value: v,
		}
	}
	return nil
}

func (p *arrayBufferImpl) ByteLength() int {
	return p.get("byteLength").toInt()
}

func (p *arrayBufferImpl) IsView(arg interface{}) bool {
	if v := JSValue(arg); v != jsNull {
		return p.call("isView", v).toBool()
	}

	return false
}

func (p *arrayBufferImpl) Slice(begin int, end ...int) ArrayBuffer {
	switch len(end) {
	case 0:
		return wrapArrayBuffer(p.call("slice", begin))
	default:
		return wrapArrayBuffer(p.call("slice", begin, end[0]))
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
	if v.valid() {
		return &dataViewImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataViewImpl) Buffer() ArrayBuffer {
	return wrapArrayBuffer(p.get("buffer"))
}

func (p *dataViewImpl) ByteLength() int {
	return p.get("byteLength").toInt()
}

func (p *dataViewImpl) ByteOffset() int {
	return p.get("byteOffset").toInt()
}

func (p *dataViewImpl) Int8() int8 {
	return int8(p.call("getInt8").toInt())
}

func (p *dataViewImpl) Uint8() uint8 {
	return uint8(p.call("getUint8").toInt())
}

func (p *dataViewImpl) Int16() int16 {
	return int16(p.call("getInt16").toInt())
}

func (p *dataViewImpl) Uint16() uint16 {
	return uint16(p.call("getUint16").toInt())
}

func (p *dataViewImpl) Int32() int32 {
	return int32(p.call("getInt32").toInt())
}

func (p *dataViewImpl) Uint32() uint32 {
	return uint32(p.call("getUint32").toInt())
}

func (p *dataViewImpl) Float32() float32 {
	return float32(p.call("getFloat32").toFloat64())
}

func (p *dataViewImpl) Float64() float64 {
	return p.call("getFloat64").toFloat64()
}

func (p *dataViewImpl) SetInt8(v int8) {
	p.call("setInt8", v)
}

func (p *dataViewImpl) SetUint8(v uint8) {
	p.call("setUint8", v)
}

func (p *dataViewImpl) SetInt16(v int16) {
	p.call("setInt16", v)
}

func (p *dataViewImpl) SetUint16(v uint16) {
	p.call("setUint16", v)
}

func (p *dataViewImpl) SetInt32(v int32) {
	p.call("setInt32", v)
}

func (p *dataViewImpl) SetUint32(v uint32) {
	p.call("setUint32", v)
}

func (p *dataViewImpl) SetFloat32(v float32) {
	p.call("setFloat32", v)
}

func (p *dataViewImpl) SetFloat64(v float64) {
	p.call("setFloat64", v)
}

// -------------8<---------------------------------------

type bufferSourceImpl struct {
	Value
}

func wrapBufferSource(v Value) BufferSource {
	if v.valid() {
		return &bufferSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------
