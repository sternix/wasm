// +build js,wasm

package wasm

import (
	"errors"
	"time"
)

// -------------8<---------------------------------------

var (
	errUnsupporttedType = errors.New("Unsupported Type")
	errInvalidType      = errors.New("Invalid Type")
	jsDate              = jsGlobal.get("Date")
	jsWindowProxy       = jsGlobal
	jsMessagePort       = jsGlobal.get("MessagePort")
	jsUint8Array        = jsGlobal.get("Uint8Array")
	jsJSON              = jsGlobal.get("JSON")
)

// -------------8<---------------------------------------
// TODO
func sliceToJsArray(slc interface{}) jsValue {
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
			arr.SetIndex(i, JSValueOf(t))
		}
		return arr

	case []MessagePort:
		arr := jsArray.New(len(x))
		for i, t := range x {
			arr.SetIndex(i, JSValueOf(t))
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
	if v.valid() && v.length() > 0 {
		ret := make([]Node, v.length())
		for i := range ret {
			ret[i] = wrapAsNode(v.index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func elementArrayToSlice(v Value) []Element {
	ret := make([]Element, v.length())
	for i := range ret {
		ret[i] = wrapAsElement(v.index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func domQuadArrayToSlice(v Value) []DOMQuad {
	if v.valid() && v.length() > 0 {
		ret := make([]DOMQuad, v.length())
		for i := range ret {
			ret[i] = wrapDOMQuad(v.index(i))
		}
		return ret
	}

	return nil
}

// -------------8<---------------------------------------

func stringSequenceToSlice(s Value) []string {
	if s.valid() && s.length() > 0 {
		ret := make([]string, s.length())
		for i := range ret {
			ret[i] = s.index(i).toString()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func boolSequenceToSlice(s Value) []bool {
	if s.valid() && s.length() > 0 {
		ret := make([]bool, s.length())
		for i := range ret {
			ret[i] = s.index(i).toBool()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func floatSequenceToSlice(s Value) []float64 {
	if s.valid() && s.length() > 0 {
		ret := make([]float64, s.length())
		for i := range ret {
			ret[i] = s.index(i).toFloat64()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func mutationRecordSequenceToSlice(v Value) []MutationRecord {
	if v.valid() && v.length() > 0 {
		ret := make([]MutationRecord, v.length())
		for i := range ret {
			ret[i] = wrapMutationRecord(v.index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func fileListToSlice(v Value) []File {
	ret := make([]File, v.length())
	for i := range ret {
		ret[i] = wrapFile(v.index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func keys(v Value) []string {
	if v.valid() {
		a := Value{jsObject.Call("keys", v)}
		ret := make([]string, a.length())
		for i := range ret {
			ret[i] = a.index(i).toString()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func domStringMapToMap(v Value) map[string]string {
	m := make(map[string]string)

	for _, key := range keys(v) {
		m[key] = v.get(key).toString()
	}

	return m
}

// -------------8<---------------------------------------

// expects v is js Date object
func jsDateToTime(v Value) time.Time {
	ms := int64(v.call("getTime").toFloat64()) * int64(time.Millisecond)
	return time.Unix(0, ms)
}

// -------------8<---------------------------------------
// https://heycam.github.io/webidl/#DOMTimeStamp

func domTimeStampToTime(ts int) time.Time {
	return time.Unix(0, int64(ts)*int64(time.Millisecond))
}

// -------------8<---------------------------------------

func domStringListToSlice(dsl Value) []string {
	if dsl.valid() && dsl.length() > 0 {
		ret := make([]string, dsl.length())
		for i := range ret {
			ret[i] = dsl.call("item").toString()
		}

		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func touchListToSlice(v Value) []Touch {
	if v.valid() && v.length() > 0 {
		ret := make([]Touch, v.length())
		for i := range ret {
			ret[i] = wrapTouch(v.index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toFloat32Slice(v Value) []float32 {
	if v.valid() && v.length() > 0 {
		ret := make([]float32, v.length())
		for i := range ret {
			ret[i] = float32(v.index(i).toFloat64())
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toFloat64Slice(v Value) []float64 {
	if v.valid() && v.length() > 0 {
		ret := make([]float64, v.length())
		for i := range ret {
			ret[i] = v.index(i).toFloat64()
		}

		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToElementSlice(v Value) []Element {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		ret := make([]Element, c.Length())
		for i := uint(0); i < c.Length(); i++ {
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
		for i := uint(0); i < c.Length(); i++ {
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
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLOptionElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type DOMError uint16

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
