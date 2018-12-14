// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	Callback interface {
		js.Wrapper

		Release()
		jsCallback() js.Func
		jsFunc(js.Value, []js.Value) interface{}
	}

	// for SetTimeout and SetInterval
	TimerCallback interface {
		Callback
	}

	// https://www.w3.org/TR/html52/browsers.html#callbackdef-framerequestcallback
	FrameRequestCallback interface {
		Callback
	}

	// https://dom.spec.whatwg.org/#callbackdef-mutationcallback
	MutationCallback interface {
		Callback
	}

	// https://www.w3.org/TR/html52/semantics-scripting.html#callbackdef-blobcallback
	BlobCallback interface {
		Callback
	}

	// https://w3c.github.io/geolocation-api/#dom-positioncallback
	PositionCallback interface {
		Callback
	}

	// https://w3c.github.io/geolocation-api/#dom-positionerrorcallback
	PositionErrorCallback interface {
		Callback
	}

	// https://www.w3.org/TR/html52/editing.html#callbackdef-functionstringcallback
	FunctionStringCallback interface {
		Callback
	}

	// https://heycam.github.io/webidl/#VoidFunction
	VoidFunction interface {
		Callback
	}
)

// -------------8<---------------------------------------

type callbackImpl struct {
	js.Func
}

func newCallbackImpl() *callbackImpl {
	return &callbackImpl{}
}

func (p *callbackImpl) Release() {
	if !isNil(p.Func.Value) {
		p.Func.Release()
	}
}

func (p *callbackImpl) jsCallback() js.Func {
	return p.Func
}

// -------------8<---------------------------------------

func NewTimerCallback(fn func(...interface{}), args ...interface{}) TimerCallback {
	h := &timerCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
		args:         args,
	}

	h.Func = js.FuncOf(h.jsFunc)
	return h
}

type timerCallbackImpl struct {
	*callbackImpl
	fn   func(...interface{})
	args []interface{}
}

func (p *timerCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	p.fn(p.args...)
	return nil
}

// -------------8<---------------------------------------

func NewFrameRequestCallback(fn func(FrameRequestCallback, float64)) FrameRequestCallback {
	h := &frameRequestCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	h.Func = js.FuncOf(h.jsFunc)
	return h
}

type frameRequestCallbackImpl struct {
	*callbackImpl
	fn func(FrameRequestCallback, float64)
}

func (p *frameRequestCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	p.fn(p, args[0].Float())
	return nil
}

// -------------8<---------------------------------------

func NewBlobCallback(fn func(Blob)) BlobCallback {
	cb := &blobCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type blobCallbackImpl struct {
	*callbackImpl
	fn func(Blob)
}

func (p *blobCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		p.fn(newBlob(args[0]))
	} else {
		p.fn(nil)
	}
	return nil
}

// -------------8<---------------------------------------

func NewMutationCallback(fn func([]MutationRecord, MutationObserver)) MutationCallback {
	cb := &mutationCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type mutationCallbackImpl struct {
	*callbackImpl
	fn func([]MutationRecord, MutationObserver)
}

func (p *mutationCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) == 2 {
		p.fn(mutationRecordSequenceToSlice(args[0]), newMutationObserver(args[1]))
	}
	return nil
}

// -------------8<---------------------------------------

func NewPositionCallback(fn func(Position)) PositionCallback {
	cb := &positionCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type positionCallbackImpl struct {
	*callbackImpl
	fn func(Position)
}

func (p *positionCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		p.fn(newPosition(args[0]))
	}
	return nil
}

// -------------8<---------------------------------------

func NewPositionErrorCallback(fn func(PositionError)) PositionErrorCallback {
	cb := &positionErrorCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type positionErrorCallbackImpl struct {
	*callbackImpl
	fn func(PositionError)
}

func (p *positionErrorCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		p.fn(newPositionError(args[0]))
	}
	return nil
}

// -------------8<---------------------------------------

func NewFunctionStringCallback(fn func(string)) FunctionStringCallback {
	cb := &functionStringCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type functionStringCallbackImpl struct {
	*callbackImpl
	fn func(string)
}

func (p *functionStringCallbackImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		p.fn(args[0].String())
	}
	return nil
}

// -------------8<---------------------------------------

func NewVoidFunction(fn func()) VoidFunction {
	cb := &voidFunctionImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.Func = js.FuncOf(cb.jsFunc)
	return cb
}

type voidFunctionImpl struct {
	*callbackImpl
	fn func()
}

func (p *voidFunctionImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	p.fn()
	return nil
}

// -------------8<---------------------------------------
