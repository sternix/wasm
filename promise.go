// +build js,wasm

package wasm

import (
	"syscall/js"
)

/*
type PromiseState int

const (
	PromiseStatePending PromiseState = iota
	PromiseStateFulfilled
	PromiseStateRejected
	PromiseStateSettled
)
*/

// taken from https://go-review.googlesource.com/c/go/+/150917/
// modified as standalone func
func Await(v js.Value) (result js.Value, ok bool) {
	if v.Type() != js.TypeObject || v.Get("then").Type() != js.TypeFunction {
		return v, true

	}

	done := make(chan struct{})

	onResolve := js.NewCallback(func(this js.Value, args []js.Value) interface{} {
		result = args[0]
		ok = true
		close(done)
		return nil
	})
	defer onResolve.Release()

	onReject := js.NewCallback(func(this js.Value, args []js.Value) interface{} {
		result = args[0]
		ok = false
		close(done)
		return nil
	})
	defer onReject.Release()

	v.Call("then", onResolve, onReject)
	<-done
	return
}

var _ Promise = &promiseImpl{}

type Promise interface {
	js.Wrapper

	All(...js.Callback) Promise
	Then(js.Callback, ...js.Callback) Promise
	Catch(js.Callback) Promise
	Finally(js.Callback) Promise
	Race([]Promise) Promise
	Reject(string) Promise
	Resolve(interface{}) Promise
}

type promiseImpl struct {
	js.Value
}

func newPromiseImpl(v js.Value) *promiseImpl {
	if isNil(v) {
		return nil
	}

	return &promiseImpl{
		Value: v,
	}
}

func (p *promiseImpl) All(cbs ...js.Callback) Promise {
	var params []interface{}

	for _, cb := range cbs {
		params = append(params, cb)
	}
	return newPromiseImpl(p.Call("all", params...))
}

func (p *promiseImpl) Then(resolve js.Callback, reject ...js.Callback) Promise {
	var v js.Value
	if len(reject) > 0 {
		v = p.Call("then", resolve, reject[0])
	} else {
		v = p.Call("then", resolve)
	}

	return newPromiseImpl(v)
}

func (p *promiseImpl) Catch(cb js.Callback) Promise {
	return newPromiseImpl(p.Call("catch", cb))
}

func (p *promiseImpl) Finally(cb js.Callback) Promise {
	return newPromiseImpl(p.Call("finally", cb))
}

func (p *promiseImpl) Race(s []Promise) Promise {
	return newPromiseImpl(p.Call("race", sliceToJsArray(s)))
}

func (p *promiseImpl) Reject(reason string) Promise {
	return newPromiseImpl(p.Call("reject", reason))
}

func (p *promiseImpl) Resolve(arg interface{}) Promise {
	var v js.Value
	switch x := arg.(type) {
	case js.Value:
		v = p.Call("resolve", x)
	case Promise:
		v = p.Call("resolve", x.JSValue())
	case Callback: // for thenable
		v = p.Call("resolve", x.JSValue())
	default:
		panic("Wrong parameter type for Promise.Resolve")
	}
	return newPromiseImpl(v)
}

// -------------8<---------------------------------------

type ArrayBufferPromise struct {
	*promiseImpl
}

func NewArrayBufferPromise(v js.Value) *ArrayBufferPromise {
	if isNil(v) {
		return nil
	}

	return &ArrayBufferPromise{
		promiseImpl: newPromiseImpl(v),
	}
}

/*
func (p *ArrayBufferPromise) Then(resolve func(ArrayBuffer), reject ...func(ArrayBuffer)) ArrayBufferPromise {
	cbResolve := js.NewCallback()
	p.promiseImpl.Then()
}
*/

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
