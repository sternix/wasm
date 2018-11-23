// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (

	// https://html.spec.whatwg.org/multipage/web-messaging.html#messagechannel
	MessageChannel interface {
		Port1() MessagePort
		Port2() MessagePort
	}

	// https://html.spec.whatwg.org/multipage/web-messaging.html#messageport
	MessagePort interface {
		EventTarget

		PostMessage(interface{}) // optional sequence<object> transfer = []
		Start()
		Close()

		OnMessage(func(Event)) EventHandler
		OnMessageError(func(Event)) EventHandler
	}

	// https://html.spec.whatwg.org/multipage/web-messaging.html#broadcastchannel
	BroadcastChannel interface {
		EventTarget

		Name() string
		PostMessage(interface{})
		Close()
		OnMessage(func(Event)) EventHandler
		OnMessageError(func(Event)) EventHandler
	}

	// https://html.spec.whatwg.org/multipage/comms.html#messageevent
	MessageEvent interface {
		Event

		Data() interface{}
		Origin() string
		LastEventId() string
		Source() MessageEventSource
		Ports() []MessagePort
		InitMessageEvent(string, ...interface{})
	}

	// https://html.spec.whatwg.org/multipage/comms.html#messageeventsource
	// typedef (WindowProxy or MessagePort or ServiceWorker) MessageEventSource;
	MessageEventSource interface {
		js.Wrapper
	}
)

// -------------8<---------------------------------------

// https://html.spec.whatwg.org/multipage/comms.html#messageeventinit
type MessageEventInit struct {
	EventInit

	Data        interface{}
	Origin      string
	LastEventId string
	Source      MessageEventSource
	Ports       []MessagePort
}

func (p MessageEventInit) toDict() js.Value {
	o := p.EventInit.toDict()
	o.Set("data", p.Data)
	o.Set("origin", p.Origin)
	o.Set("lastEventId", p.LastEventId)
	o.Set("source", p.Source.JSValue())
	o.Set("ports", messagePortSliceToJsArray(p.Ports))
	return o
}
