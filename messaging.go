// +build js,wasm

package wasm

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

		OnMessage(func(MessageEvent)) EventHandler
		OnMessageError(func(Event)) EventHandler
	}

	// https://html.spec.whatwg.org/multipage/web-messaging.html#broadcastchannel
	BroadcastChannel interface {
		EventTarget

		Name() string
		PostMessage(interface{})
		Close()
		OnMessage(func(MessageEvent)) EventHandler
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
	MessageEventSource interface{}
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

func (p MessageEventInit) toJSObject() Value {
	o := p.EventInit.toJSObject()
	o.set("data", p.Data)
	o.set("origin", p.Origin)
	o.set("lastEventId", p.LastEventId)
	o.set("source", JSValue(p.Source))
	o.set("ports", sliceToJsArray(p.Ports))
	return o
}
