// +build js,wasm

package wasm

import (
	"reflect"
	"syscall/js"
	"unsafe"
)

type (
	jsValue      = js.Value
	jsWrapper    = js.Wrapper
	jsTypedArray = js.TypedArray
)

var (
	jsObject    = js.Global().Get("Object")
	jsArray     = js.Global().Get("Array")
	jsTypeFunc  = jsObject.Get("prototype").Get("toString")
	jsGlobal    = Value{js.Global()}
	jsNull      = js.Null()
	jsUndefined = js.Undefined()
)

const (
	TypeUndefined = js.TypeUndefined
	TypeNull      = js.TypeNull
	TypeBoolean   = js.TypeBoolean
	TypeNumber    = js.TypeNumber
	TypeString    = js.TypeString
	TypeSymbol    = js.TypeSymbol
	TypeObject    = js.TypeObject
	TypeFunction  = js.TypeFunction
)

func jsTypedArrayOf(slice interface{}) jsTypedArray {
	return js.TypedArrayOf(slice)
}

type Value struct {
	jsValue jsValue
}

func (p Value) valid() bool {
	if p.jsValue == jsNull || p.jsValue == jsUndefined {
		return false
	}
	return true
}

func (p Value) toBool() bool {
	if p.valid() {
		return p.jsValue.Bool()
	}
	return false
}

func (p Value) toString() string {
	if p.valid() {
		return p.jsValue.String()
	}
	return ""
}

func (p Value) toFloat64() float64 {
	if p.valid() {
		return p.jsValue.Float()
	}
	return 0.0
}

func (p Value) toInt64() int64 {
	return int64(p.toFloat64())
}

func (p Value) toUint64() uint64 {
	return uint64(p.toFloat64())
}

func (p Value) toInt() int {
	return int(p.toFloat64())
}

func (p Value) toUint16() uint16 {
	return uint16(p.toFloat64())
}

func (p Value) toInt16() int16 {
	return int16(p.toFloat64())
}

func (p Value) toUint8() uint8 {
	return uint8(p.toFloat64())
}

func (p Value) toUint() uint {
	return uint(p.toFloat64())
}

func (p Value) toFloat32() float32 {
	return float32(p.toFloat64())
}

func (v Value) instanceOf(t Value) bool {
	return v.jsValue.InstanceOf(t.jsValue)
}

func (p Value) JSValue() js.Value {
	return p.jsValue
}

func (p Value) jsType() string {
	if p.jsValue.Type() == js.TypeObject {
		str := jsTypeFunc.Call("call", p.jsValue).String()
		return str[8 : len(str)-1]
	}

	return p.jsValue.Type().String()
}

func (p Value) get(property string) Value {
	return Value{
		p.jsValue.Get(property),
	}
}

func (p Value) set(property string, x interface{}) {
	p.jsValue.Set(property, x)
}

func (p Value) setIndex(i int, x interface{}) {
	p.jsValue.SetIndex(i, x)
}

func (p Value) length() int {
	return p.jsValue.Length()
}

func (p Value) call(m string, args ...interface{}) Value {
	return Value{
		p.jsValue.Call(m, args...),
	}
}

func (p Value) invoke(args ...interface{}) Value {
	return Value{
		p.jsValue.Invoke(args...),
	}
}

func (p Value) jsNew(args ...interface{}) Value {
	return Value{
		p.jsValue.New(args...),
	}
}

func (p Value) index(i int) Value {
	return Value{
		p.jsValue.Index(i),
	}
}

func (p Value) toSlice() []Value {
	if p.valid() {
		slc := make([]Value, p.length())
		for i := range slc {
			slc[i] = p.index(i)
		}
		return slc
	}
	return nil
}

/*
func nodeListToSlice(v Value) []Node {
	if v.valid() && v.length() > 0 {
		ret := make([]Node, v.length())
		for i := range ret {
			ret[i] = wrapAsNode(v.index(i))
		}
		return ret
	}
	return nil
}


valuePtr := reflect.ValueOf(arrPtr)
 value := valuePtr.Elem()

 value.Set(reflect.Append(value, reflect.ValueOf(55)))
 value.Set(reflect.Append(value, reflect.ValueOf(56)))
 value.Set(reflect.Append(value, reflect.ValueOf(57)))

 fmt.Println(value.Len())


*/

// Append's Values to given slice according to wrapfn function
// expects slice pointer
func (p Value) AppendToSlice(t interface{}, wrapfn func(v Value) interface{}) {
	if p.valid() && p.length() > 0 {
		if rv := reflect.ValueOf(t); rv.Kind() == reflect.Ptr {
			if slice := rv.Elem(); slice.Kind() == reflect.Slice {
				for i := 0; i < p.length(); i++ {
					if item := wrapfn(p.index(i)); item != nil {
						slice.Set(reflect.Append(slice, reflect.ValueOf(item)))
					}
				}
			}
		}
	}
}

// -------------8<---------------------------------------

// taken from https://go-review.googlesource.com/c/go/+/150917/
// modified as standalone func
func await(v Value) (result Value, ok bool) {
	if v.jsValue.Type() != js.TypeObject || v.get("then").jsValue.Type() != js.TypeFunction {
		return v, true

	}

	done := make(chan struct{})

	onResolve := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result = Value{args[0]}
		ok = true
		close(done)
		return nil
	})
	defer onResolve.Release()

	onReject := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result = Value{args[0]}
		ok = false
		close(done)
		return nil
	})
	defer onReject.Release()

	v.call("then", onResolve, onReject)
	<-done
	return
}

// -------------8<---------------------------------------

func JSValueOf(x interface{}) jsValue {
	switch x := x.(type) {
	case jsValue: // should precede Wrapper to avoid a loop
		return x
	case jsWrapper:
		return x.JSValue()
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, unsafe.Pointer, float32, float64, string, []interface{}, map[string]interface{}:
		return js.ValueOf(x)
	default: // if type has embedded Value field
		if v, ok := reflect.ValueOf(x).Elem().FieldByName("Value").Interface().(Value); ok {
			return v.JSValue()
		}
		return jsNull
	}
}

// -------------8<---------------------------------------

func Equal(x interface{}, y interface{}) bool {
	return JSValueOf(x) == JSValueOf(y)
}

// -------------8<---------------------------------------

type Func struct {
	js.Func
}

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	fx := func(xthis js.Value, xargs []js.Value) interface{} {
		var (
			fxThis = Value{xthis}
			fxArgs []Value
		)

		if xargs != nil && len(xargs) > 0 {
			fxArgs = make([]Value, len(xargs))
			for i, v := range xargs {
				fxArgs[i] = Value{v}
			}
		}
		return fn(fxThis, fxArgs)
	}
	return Func{js.FuncOf(fx)}
}

// -------------8<---------------------------------------

func uint8ArrayToByteSlice(v Value) []byte {
	jsa := jsUint8Array.jsNew(v)
	ret := make([]byte, jsa.get("byteLength").toInt())
	ta := js.TypedArrayOf(ret)
	ta.Call("set", jsa)
	ta.Release()
	return ret
}

// -------------8<---------------------------------------

// expects a go slice and returns JavaScript Array with the slice values
func ToJSArray(t interface{}) jsValue {
	if reflect.TypeOf(t).Kind() == reflect.Slice {
		slc := reflect.ValueOf(t)
		l := slc.Len()
		jsArr := jsArray.New(l)
		if l > 0 {
			for i := 0; i < l; i++ {
				jsArr.SetIndex(i, JSValueOf(slc.Index(i).Interface()))
			}
		}
		return jsArr
	}
	return js.Null()
}

// -------------8<---------------------------------------

var (
	ifaceSliceValue = reflect.ValueOf([]interface{}{})
	ifaceSliceType  = ifaceSliceValue.Type()
	ifaceSliceZero  = reflect.Zero(ifaceSliceType).Interface().([]interface{})
)

// expects a go slice with any type, returns new []interface{} slice with given slice's values
func ToIfaceSlice(t interface{}) []interface{} {
	if v := reflect.ValueOf(t); v.Kind() == reflect.Slice {
		if l := v.Len(); l > 0 {
			ifaceSlice := reflect.MakeSlice(ifaceSliceType, l, l)
			for i := 0; i < l; i++ {
				ifaceSlice.Index(i).Set(v.Index(i))
			}
			return ifaceSlice.Interface().([]interface{})
		}
	}
	return ifaceSliceZero
}

// -------------8<---------------------------------------
