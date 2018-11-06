// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

type (
	Callback interface {
		Release()
		jsCallback() js.Callback
		jsFunc([]js.Value)
		JSValue() js.Value
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
	cb js.Callback
}

func newCallbackImpl() *callbackImpl {
	return &callbackImpl{}
}

func (p *callbackImpl) Release() {
	if !isNil(p.cb.Value) {
		p.cb.Release()
	}
}

func (p *callbackImpl) jsCallback() js.Callback {
	return p.cb
}

func (p *callbackImpl) JSValue() js.Value {
	return p.cb.Value
}

// -------------8<---------------------------------------

func NewTimerCallback(fn func(...interface{}), args ...interface{}) TimerCallback {
	h := &timerCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
		args:         args,
	}

	h.cb = js.NewCallback(h.jsFunc)
	return h
}

type timerCallbackImpl struct {
	*callbackImpl
	fn   func(...interface{})
	args []interface{}
}

func (p *timerCallbackImpl) jsFunc(args []js.Value) {
	p.fn(p.args...)
}

// -------------8<---------------------------------------

func NewFrameRequestCallback(fn func(time.Time)) FrameRequestCallback {
	h := &frameRequestCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	h.cb = js.NewCallback(h.jsFunc)
	return h
}

type frameRequestCallbackImpl struct {
	*callbackImpl
	fn func(time.Time)
}

func (p *frameRequestCallbackImpl) jsFunc(args []js.Value) {
	p.fn(highResTimeStampToTime(args[0]))
}

// -------------8<---------------------------------------

func NewBlobCallback(fn func(Blob)) BlobCallback {
	cb := &blobCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type blobCallbackImpl struct {
	*callbackImpl
	fn func(Blob)
}

func (p *blobCallbackImpl) jsFunc(args []js.Value) {
	if len(args) == 1 {
		p.fn(newBlob(args[0]))
	} else {
		p.fn(nil)
	}
}

// -------------8<---------------------------------------

func NewMutationCallback(fn func([]MutationRecord, MutationObserver)) MutationCallback {
	cb := &mutationCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type mutationCallbackImpl struct {
	*callbackImpl
	fn func([]MutationRecord, MutationObserver)
}

func (p *mutationCallbackImpl) jsFunc(args []js.Value) {
	if len(args) == 2 {
		p.fn(mutationRecordSequenceToSlice(args[0]), newMutationObserver(args[1]))
	}
}

// -------------8<---------------------------------------

func NewPositionCallback(fn func(Position)) PositionCallback {
	cb := &positionCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type positionCallbackImpl struct {
	*callbackImpl
	fn func(Position)
}

func (p *positionCallbackImpl) jsFunc(args []js.Value) {
	if len(args) == 1 {
		p.fn(newPosition(args[0]))
	}
}

// -------------8<---------------------------------------

func NewPositionErrorCallback(fn func(PositionError)) PositionErrorCallback {
	cb := &positionErrorCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type positionErrorCallbackImpl struct {
	*callbackImpl
	fn func(PositionError)
}

func (p *positionErrorCallbackImpl) jsFunc(args []js.Value) {
	if len(args) == 1 {
		p.fn(newPositionError(args[0]))
	}
}

// -------------8<---------------------------------------

func NewFunctionStringCallback(fn func(string)) FunctionStringCallback {
	cb := &functionStringCallbackImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type functionStringCallbackImpl struct {
	*callbackImpl
	fn func(string)
}

func (p *functionStringCallbackImpl) jsFunc(args []js.Value) {
	if len(args) == 1 {
		p.fn(args[0].String())
	}
}

// -------------8<---------------------------------------

func NewVoidFunction(fn func()) VoidFunction {
	cb := &voidFunctionImpl{
		callbackImpl: newCallbackImpl(),
		fn:           fn,
	}

	cb.cb = js.NewCallback(cb.jsFunc)
	return cb
}

type voidFunctionImpl struct {
	*callbackImpl
	fn func()
}

func (p *voidFunctionImpl) jsFunc([]js.Value) {
	p.fn()
}

// -------------8<---------------------------------------
