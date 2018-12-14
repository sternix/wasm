// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type eventHandlerImpl struct {
	jsCb js.Func
	fn   func(Event)
	typ  string
}

func (p *eventHandlerImpl) Type() string {
	return p.typ
}

func (p *eventHandlerImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		p.Handle(wrapEvent(args[0]))
	}
	return nil
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

func EventHandlerFunc(typ string, fn func(Event)) EventHandler {
	eh := &eventHandlerImpl{
		fn:  fn,
		typ: typ,
	}

	eh.jsCb = js.FuncOf(eh.jsFunc)

	js.Global().Get("addEventListener").Invoke(typ, eh.jsCb)
	return eh
}

// -------------8<---------------------------------------

type elementEventHandlerImpl struct {
	js.Value
	jsCb js.Func
	fn   func(Event)
	typ  string
}

func (p *elementEventHandlerImpl) Type() string {
	return p.typ
}

func (p *elementEventHandlerImpl) jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		p.Handle(wrapEvent(args[0]))
	}
	return nil
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
