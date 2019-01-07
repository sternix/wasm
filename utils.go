// +build js,wasm

package wasm

import (
	"errors"
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

var (
	errUnsupporttedType = errors.New("Unsupported Type")
	errInvalidType      = errors.New("Invalid Type")
	jsArray             = jsGlobal.Get("Array")
	//jsObject            = jsGlobal.Get("Object")
	//jsTypeFunc          = jsObject.Get("prototype").Get("toString")
	//jsTypeFunc = jsGlobal.Get("Object").Get("prototype").Get("toString")
	jsDate        = jsGlobal.Get("Date")
	jsWindowProxy = jsGlobal
	jsMessagePort = jsGlobal.Get("MessagePort")
	jsUint8Array  = jsGlobal.Get("Uint8Array")
	//jsUint8ClampedArray = jsGlobal.Get("Uint8ClampedArray")
	//jsBufferSource      = jsGlobal.Get("BufferSource") --> typedef
	jsJSON = jsGlobal.Get("JSON")
)

// -------------8<---------------------------------------
// TODO
func sliceToJsArray(slc interface{}) Value {
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
	case []Value:
		arr := jsArray.New(len(x))
		for i, s := range x {
			arr.SetIndex(i, s)
		}
		return arr
	case []bool:
		arr := jsArray.New(len(x))
		for i, b := range x {
			arr.SetIndex(i, b)
		}
		return arr
	case []Touch:
		arr := jsArray.New(len(x))
		for i, t := range x {
			arr.SetIndex(i, JSValue(t))
		}
		return arr

	case []MessagePort:
		arr := jsArray.New(len(x))
		for i, t := range x {
			arr.SetIndex(i, JSValue(t))
		}
		return arr
	default:
		// TODO: remove this when all types ok
		panic("sliceToJsArray: unregistered type")
	}
	//return js.Null()
}

// -------------8<---------------------------------------

func nodeListToSlice(v Value) []Node {
	if v.Valid() && v.Length() > 0 {
		ret := make([]Node, v.Length())
		for i := range ret {
			ret[i] = wrapAsNode(v.Index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func elementArrayToSlice(v Value) []Element {
	ret := make([]Element, v.Length())
	for i := range ret {
		ret[i] = wrapAsElement(v.Index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func domQuadArrayToSlice(v Value) []DOMQuad {
	if v.Valid() && v.Length() > 0 {
		ret := make([]DOMQuad, v.Length())
		for i := range ret {
			ret[i] = wrapDOMQuad(v.Index(i))
		}
		return ret
	}

	return nil
}

// -------------8<---------------------------------------

func stringSequenceToSlice(s Value) []string {
	if s.Valid() && s.Length() > 0 {
		ret := make([]string, s.Length())
		for i := range ret {
			ret[i] = s.Index(i).String()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func boolSequenceToSlice(s Value) []bool {
	if s.Valid() && s.Length() > 0 {
		ret := make([]bool, s.Length())
		for i := range ret {
			ret[i] = s.Index(i).Bool()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func floatSequenceToSlice(s Value) []float64 {
	if s.Valid() && s.Length() > 0 {
		ret := make([]float64, s.Length())
		for i := range ret {
			ret[i] = s.Index(i).Float()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func mutationRecordSequenceToSlice(v Value) []MutationRecord {
	if v.Valid() && v.Length() > 0 {
		ret := make([]MutationRecord, v.Length())
		for i := range ret {
			ret[i] = wrapMutationRecord(v.Index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func fileListToSlice(v Value) []File {
	ret := make([]File, v.Length())
	for i := range ret {
		ret[i] = wrapFile(v.Index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func keys(v Value) []string {
	if v.Valid() {
		a := jsObject.Call("keys", v)
		ret := make([]string, a.Length())
		for i := range ret {
			ret[i] = a.Index(i).String()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func domStringMapToMap(v Value) map[string]string {
	m := make(map[string]string)

	for _, key := range keys(v) {
		m[key] = v.Get(key).String()
	}

	return m
}

// -------------8<---------------------------------------

// expects v is js Date object
func jsDateToTime(v Value) time.Time {
	ms := int64(v.Call("getTime").Float()) * int64(time.Millisecond)
	return time.Unix(0, ms)
}

// -------------8<---------------------------------------
// https://heycam.github.io/webidl/#DOMTimeStamp

func domTimeStampToTime(ts int) time.Time {
	return time.Unix(0, int64(ts)*int64(time.Millisecond))
}

// -------------8<---------------------------------------

func domStringListToSlice(dsl Value) []string {
	if dsl.Valid() && dsl.Length() > 0 {
		ret := make([]string, dsl.Length())
		for i := range ret {
			ret[i] = dsl.Call("item").String()
		}

		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func touchListToSlice(v Value) []Touch {
	if v.Valid() && v.Length() > 0 {
		ret := make([]Touch, v.Length())
		for i := range ret {
			ret[i] = wrapTouch(v.Index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toFloat32Slice(v Value) []float32 {
	if v.Valid() && v.Length() > 0 {
		ret := make([]float32, v.Length())
		for i := range ret {
			ret[i] = float32(v.Index(i).Float())
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toFloat64Slice(v Value) []float64 {
	if v.Valid() && v.Length() > 0 {
		ret := make([]float64, v.Length())
		for i := range ret {
			ret[i] = v.Index(i).Float()
		}

		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToElementSlice(v Value) []Element {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		ret := make([]Element, c.Length())
		for i := 0; i < c.Length(); i++ {
			ret[i] = c.Item(i)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToHTMLElementSlice(v Value) []HTMLElement {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		var ret []HTMLElement
		for i := 0; i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToHTMLOptionElementSlice(v Value) []HTMLOptionElement {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		var ret []HTMLOptionElement
		for i := 0; i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLOptionElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func uint8ArrayToByteSlice(v Value) []byte {
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
