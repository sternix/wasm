// +build js,wasm

package wasm

import (
	"encoding/json"
	"errors"
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

var (
	errUnsupporttedType = errors.New("Unsupported Type")
	errInvalidType      = errors.New("Invalid Type")
	jsJSON              = js.Global().Get("JSON")
	jsArray             = js.Global().Get("Array")
	jsObject            = js.Global().Get("Object")
	jsTypeFunc          = jsObject.Get("prototype").Get("toString")
	//jsTypeFunc = js.Global().Get("Object").Get("prototype").Get("toString")
	jsDate        = js.Global().Get("Date")
	jsWindowProxy = js.Global()
	jsMessagePort = js.Global().Get("MessagePort")
	jsUint8Array  = js.Global().Get("Uint8Array")

	//jsUint8ClampedArray = js.Global().Get("Uint8ClampedArray")
	//jsBufferSource      = js.Global().Get("BufferSource") --> typedef

)

// -------------8<---------------------------------------

func JSType(v js.Value) string {
	if v.Type() == js.TypeObject {
		str := jsTypeFunc.Call("call", v).String()
		return str[8 : len(str)-1]
	}

	return v.Type().String()
}

// -------------8<---------------------------------------

func valueOf(x interface{}) interface{} {
	if v, ok := x.(Wrapper); ok {
		return v.JSValue()
	}

	return x
}

// -------------8<---------------------------------------

// TODO use map[string]interface{}
// convert structs to js.Value

func toJSONObject(t interface{}) js.Value {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	return jsJSON.Call("parse", string(data))
}

// -------------8<---------------------------------------

func isNil(v js.Value) bool {
	if v == js.Null() || v == js.Undefined() {
		return true
	}
	return false
}

// -------------8<---------------------------------------

func arrayToSlice(a js.Value) []js.Value {
	if isNil(a) {
		return nil
	}

	ret := make([]js.Value, a.Length())
	for i := range ret {
		ret[i] = a.Index(i)
	}

	return ret
}

// -------------8<---------------------------------------

func stringSliceToJsArray(slc []string) js.Value {
	arr := jsArray.New(len(slc))
	for i, s := range slc {
		arr.SetIndex(i, s)
	}
	return arr
}

// -------------8<---------------------------------------

func promiseSliceToJsArray(slc []Promise) js.Value {
	arr := jsArray.New(len(slc))
	for i, p := range slc {
		arr.SetIndex(i, p)
	}
	return arr
}

// -------------8<---------------------------------------

// TODO
func sliceToJsArray(slc interface{}) js.Value {
	switch x := slc.(type) {
	case []string:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []float64:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []int:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []uint:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []byte:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []js.Value:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	default:
		// TODO: remove this when all types ok
		panic("sliceToJsArray: unregistered type")
	}
	//return js.Null()
}

// -------------8<---------------------------------------

func nodeListToSlice(nl js.Value) []Node {
	if isNil(nl) {
		return nil
	}

	ret := make([]Node, nl.Length())

	for i := range ret {
		ret[i] = newNode(nl.Index(i))
	}

	return ret
}

// -------------8<---------------------------------------

func elementArrayToSlice(v js.Value) []Element {
	ret := make([]Element, v.Length())

	for i := range ret {
		ret[i] = newElement(v.Index(i))
	}

	return ret
}

// -------------8<---------------------------------------

func floatSliceToJsArray(sl []float64) js.Value {
	obj := jsArray.New(len(sl))

	for i, f := range sl {
		obj.SetIndex(i, f)
	}

	return obj
}

// -------------8<---------------------------------------

func stringSequenceToSlice(s js.Value) []string {
	if isNil(s) {
		return nil
	}

	ret := make([]string, s.Length())
	for i := range ret {
		ret[i] = s.Index(i).String()
	}
	return ret
}

// -------------8<---------------------------------------

func boolSequenceToSlice(s js.Value) []bool {
	if isNil(s) {
		return nil
	}

	ret := make([]bool, s.Length())
	for i := range ret {
		ret[i] = s.Index(i).Bool()
	}
	return ret
}

// -------------8<---------------------------------------

func floatSequenceToSlice(s js.Value) []float64 {
	if isNil(s) {
		return nil
	}

	ret := make([]float64, s.Length())
	for i := range ret {
		ret[i] = s.Index(i).Float()
	}
	return ret
}

// -------------8<---------------------------------------

func mutationRecordSequenceToSlice(v js.Value) []MutationRecord {
	if isNil(v) {
		return nil
	}

	ret := make([]MutationRecord, v.Length())
	for i := range ret {
		ret[i] = newMutationRecord(v.Index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func fileListToSlice(v js.Value) []File {
	ret := make([]File, v.Length())
	for i := range ret {
		ret[i] = newFile(v.Index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func keys(v js.Value) []string {
	if isNil(v) {
		return nil
	}

	a := jsObject.Call("keys", v)

	ret := make([]string, a.Length())

	for i := range ret {
		ret[i] = a.Index(i).String()
	}

	return ret
}

// -------------8<---------------------------------------

func domStringMapToMap(v js.Value) map[string]string {
	m := make(map[string]string)

	for _, key := range keys(v) {
		m[key] = v.Get(key).String()
	}

	return m
}

// -------------8<---------------------------------------

// expects v is js Date object
func jsDateToTime(v js.Value) time.Time {
	ms := int64(v.Call("getTime").Int())
	return time.Unix(0, ms*int64(time.Millisecond))
}

// -------------8<---------------------------------------
// https://heycam.github.io/webidl/#DOMTimeStamp

func domTimeStampToTime(ts int) time.Time {
	return time.Unix(0, int64(ts)*int64(time.Millisecond))
}

// -------------8<---------------------------------------

func domStringListToSlice(dsl js.Value) []string {
	if isNil(dsl) {
		return nil
	}

	ret := make([]string, dsl.Length())
	for i := range ret {
		ret[i] = dsl.Call("item").String()
	}

	return ret
}

// -------------8<---------------------------------------

func highResTimeStampToTime(v js.Value) time.Time {
	return time.Unix(0, int64(v.Float())*int64(time.Millisecond))
}

// -------------8<---------------------------------------

func touchListToSlice(v js.Value) []Touch {
	if isNil(v) {
		return nil
	}

	ret := make([]Touch, v.Length())
	for i := range ret {
		ret[i] = newTouch(v.Index(i))
	}

	return ret
}

// -------------8<---------------------------------------

func toFloat32Slice(v js.Value) []float32 {
	if isNil(v) {
		return nil
	}

	ret := make([]float32, v.Length())
	for i := range ret {
		ret[i] = float32(v.Index(i).Float())
	}

	return ret
}

// -------------8<---------------------------------------

func toFloat64Slice(v js.Value) []float64 {
	if isNil(v) {
		return nil
	}

	ret := make([]float64, v.Length())
	for i := range ret {
		ret[i] = v.Index(i).Float()
	}

	return ret
}

// -------------8<---------------------------------------

func uint8ArrayToByteSlice(v js.Value) []byte {
	jsa := jsUint8Array.New(v)
	ret := make([]byte, jsa.Get("byteLength").Int())
	ta := js.TypedArrayOf(ret)
	ta.Call("set", jsa)
	ta.Release()
	return ret
}

// -------------8<---------------------------------------

type DOMError int

const (
	ErrIndexSize             DOMError = 1
	ErrDOMStringSize         DOMError = 2
	ErrHierarchyRequest      DOMError = 3
	ErrWrongDocument         DOMError = 4
	ErrInvalidCharacter      DOMError = 5
	ErrNoDataAllowed         DOMError = 6
	ErrNoModificationAllowed DOMError = 7
	ErrNotFound              DOMError = 8
	ErrNotSupported          DOMError = 9
	ErrInuseAttribute        DOMError = 10
	ErrInvalidState          DOMError = 11
	ErrSyntax                DOMError = 12
	ErrInvalidModification   DOMError = 13
	ErrNamespace             DOMError = 14
	ErrInvalidAccess         DOMError = 15
	ErrValidation            DOMError = 16
	ErrTypeMismatch          DOMError = 17
	ErrSecurity              DOMError = 18
	ErrNetwork               DOMError = 19
	ErrAbort                 DOMError = 20
	ErrURLMismatch           DOMError = 21
	ErrQuotaExceeded         DOMError = 22
	ErrTimeout               DOMError = 23
	ErrInvalidNodeType       DOMError = 24
	ErrDataClone             DOMError = 25
)

// -------------8<---------------------------------------
