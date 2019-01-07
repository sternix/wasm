// +build js,wasm

package wasm

import (
	"syscall/js"
	"reflect"
)

var (
	jsObject    = Value{js.Global().Get("Object")}
	jsTypeFunc  = jsObject.Get("prototype").Get("toString")
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
	js.Value
}

func (p Value) Valid() bool {
	if p == jsNull || p == jsUndefined {
		return false
	}
	return true
}

func (p Value) Bool() bool {
	if p.Valid() {
		return p.Value.Bool()
	}
	return false
}

func (p Value) String() string {
	if p.Valid() {
		return p.Value.String()
	}
	return ""
}

func (p Value) Float() float64 {
	if p.Valid() {
		return p.Value.Float()
	}
	return 0.0
}

func (p Value) Int() int {
	return int(p.Float())
}

func (p Value) Uint16() uint16 {
	return uint16(p.Float())
}

func (p Value) Uint() uint {
	return uint(p.Float())
}

func (p Value) Float32() float32 {
	return float32(p.Value.Float())
}

func (v Value) InstanceOf(t Value) bool {
	return v.Value.InstanceOf(t.Value)
}

func (p Value) JSValue() js.Value {
	return p.Value
}

func (p Value) JSType() string {
	if p.Value.Type() == js.TypeObject {
		str := jsTypeFunc.Call("call", p.Value).String()
		return str[8 : len(str)-1]
	}

	return p.Value.Type().String()
}

func (p Value) Get(property string) Value {
	return Value{
		p.Value.Get(property),
	}
}

func (p Value) Call(m string, args ...interface{}) Value {
	return Value{
		p.Value.Call(m, args...),
	}
}

func (p Value) Invoke(args ...interface{}) Value {
	return Value{
		p.Value.Invoke(args...),
	}
}

func (p Value) New(args ...interface{}) Value {
	return Value{
		p.Value.New(args...),
	}
}

func (p Value) Index(i int) Value {
	return Value{
		p.Value.Index(i),
	}
}

// -------------8<---------------------------------------

func (p Value) ToSlice() []Value {
	if p.Valid() {
		slc := make([]Value, p.Length())
		for i := range slc {
			slc[i] = p.Index(i)
		}
		return slc
	}
	return nil
}

// -------------8<---------------------------------------

// taken from https://go-review.googlesource.com/c/go/+/150917/
// modified as standalone func
func await(v Value) (result Value, ok bool) {
	if v.Value.Type() != TypeObject || v.Get("then").Type() != TypeFunction {
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

	v.Call("then", onResolve, onReject)
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
		// is typed array
		if ta, ok := o.(js.Wrapper); ok {
			return Value{ta.JSValue()}
		}

		// embedded Value
		if v, ok := reflect.ValueOf(o).Elem().FieldByName("Value").Interface().(Value); ok {
			return v
		}
	}
	return jsNull
}

// -------------8<---------------------------------------
