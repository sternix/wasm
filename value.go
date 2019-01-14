// +build js,wasm

package wasm

import (
	"reflect"
	"syscall/js"
)

var (
	jsObject    = Value{js.Global().Get("Object")}
	jsTypeFunc  = jsObject.get("prototype").get("toString")
	jsGlobal    = Value{js.Global()}
	jsNull      = Value{js.Null()}
	jsUndefined = Value{js.Undefined()}
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

type Value struct {
	jsValue js.Value
}

func (p Value) valid() bool {
	if p == jsNull || p == jsUndefined {
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

func (p Value) toInt() int {
	return int(p.toFloat64())
}

func (p Value) toUint16() uint16 {
	return uint16(p.toFloat64())
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
	if p.jsValue.Type() == TypeObject {
		str := jsTypeFunc.call("call", p.jsValue).toString()
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

// -------------8<---------------------------------------

// taken from https://go-review.googlesource.com/c/go/+/150917/
// modified as standalone func
func await(v Value) (result Value, ok bool) {
	if v.jsValue.Type() != TypeObject || v.get("then").jsValue.Type() != TypeFunction {
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

func ValueOf(x interface{}) Value {
	return Value{js.ValueOf(x)}
}

// -------------8<---------------------------------------

func Equal(x interface{}, y interface{}) bool {
	return JSValue(x) == JSValue(y)
}

// -------------8<---------------------------------------

// returns interface types underlying Value
// it excepts Value is embedded in struct
func JSValue(o interface{}) Value {
	if o != nil {
		if ta, ok := o.(js.Wrapper); ok {
			return Value{ta.JSValue()}
		}

		if v, ok := reflect.ValueOf(o).Elem().FieldByName("Value").Interface().(Value); ok {
			return v
		}
	}
	return jsNull
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
