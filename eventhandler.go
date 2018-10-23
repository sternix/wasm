// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type eventHandlerImpl struct {
	jsCb js.Callback
	fn   func(Event)
	typ  string
}

func (p *eventHandlerImpl) Type() string {
	return p.typ
}

func (p *eventHandlerImpl) jsFunc(v js.Value) {
	p.Handle(wrapEvent(v))
}

func (p *eventHandlerImpl) Handle(event Event) {
	p.fn(event)
}

func (p *eventHandlerImpl) Release() {
	p.jsCb.Release()
}

func (p *eventHandlerImpl) Remove() {
	js.Global().Get("removeEventListener").Invoke(p.typ, p.jsCb)
	p.Release()
}

func (p *eventHandlerImpl) Dispatch() bool {
	return js.Global().Get("dispatchEvent").Invoke(p.jsCb).Bool()
}

func EventHandlerFunc(typ string, fn func(Event), flag ...js.EventCallbackFlag) EventHandler {
	eh := &eventHandlerImpl{
		fn:  fn,
		typ: typ,
	}

	if len(flag) > 0 {
		eh.jsCb = js.NewEventCallback(flag[0], eh.jsFunc)
	} else {
		eh.jsCb = js.NewEventCallback(0, eh.jsFunc)
	}

	js.Global().Get("addEventListener").Invoke(typ, eh.jsCb)
	return eh
}

// -------------8<---------------------------------------

type elementEventHandlerImpl struct {
	js.Value
	jsCb js.Callback
	fn   func(Event)
	typ  string
}

func (p *elementEventHandlerImpl) Type() string {
	return p.typ
}

func (p *elementEventHandlerImpl) jsFunc(v js.Value) {
	p.Handle(wrapEvent(v))
}

func (p *elementEventHandlerImpl) Handle(event Event) {
	p.fn(event)
}

func (p *elementEventHandlerImpl) Release() {
	p.jsCb.Release()
}

func (p *elementEventHandlerImpl) Remove() {
	p.Call("removeEventListener", p.typ, p.jsCb)
	p.Release()
}

func (p *elementEventHandlerImpl) Dispatch() bool {
	return p.Call("dispatchEvent", p.jsCb).Bool()
}

// -------------8<---------------------------------------
