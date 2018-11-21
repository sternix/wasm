// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type messageChannelImpl struct {
	js.Value
}

func newMessageChannel(v js.Value) MessageChannel {
	if isNil(v) {
		return nil
	}

	return &messageChannelImpl{
		Value: v,
	}
}

func (p *messageChannelImpl) Port1() MessagePort {
	return newMessagePort(p.Get("port1"))
}

func (p *messageChannelImpl) Port2() MessagePort {
	return newMessagePort(p.Get("port2"))
}

// -------------8<---------------------------------------

type messagePortImpl struct {
	*eventTargetImpl
}

func newMessagePort(v js.Value) MessagePort {
	if isNil(v) {
		return nil
	}

	return &messagePortImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

// TODO optional sequence<object> transfer = [] omitted
func (p *messagePortImpl) PostMessage(message interface{}) {
	//XXX: panicable
	p.Call("postMessage", message)
}

func (p *messagePortImpl) Start() {
	p.Call("start")
}

func (p *messagePortImpl) Close() {
	p.Call("close")
}

func (p *messagePortImpl) OnMessage(fn func(Event)) EventHandler {
	return p.On("message", fn)
}

func (p *messagePortImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

func (p *messagePortImpl) JSValue() js.Value {
	return p.Value
}

// -------------8<---------------------------------------

type broadcastChannelImpl struct {
	*eventTargetImpl
}

func newBroadcastChannel(v js.Value) BroadcastChannel {
	if isNil(v) {
		return nil
	}

	return &broadcastChannelImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *broadcastChannelImpl) Name() string {
	return p.Get("name").String()
}

func (p *broadcastChannelImpl) PostMessage(message interface{}) {
	p.Call("postMessage", message)
}

func (p *broadcastChannelImpl) Close() {
	p.Call("close")
}

func (p *broadcastChannelImpl) OnMessage(fn func(Event)) EventHandler {
	return p.On("message", fn)
}

func (p *broadcastChannelImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

// -------------8<---------------------------------------

type messageEventImpl struct {
	*eventImpl
}

func newMessageEvent(v js.Value) MessageEvent {
	if isNil(v) {
		return nil
	}

	return &messageEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *messageEventImpl) Data() interface{} {
	return Wrap(p.Get("data"))
}

func (p *messageEventImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *messageEventImpl) LastEventId() string {
	return p.Get("lastEventId").String()
}

func (p *messageEventImpl) Source() MessageEventSource {
	v := p.Get("source")
	if isNil(v) {
		return nil
	}

	if v.InstanceOf(jsWindowProxy) {
		return newWindowProxy(v)
	} else if v.InstanceOf(jsMessagePort) {
		return newMessagePort(v)
	} /* TODO: ServiceWorker  else if v.InstanceOf(jsServiceWorker) {
		return newServiceWorker(v)
	}*/

	return nil
}

func (p *messageEventImpl) Ports() []MessagePort {
	var ret []MessagePort

	ports := arrayToSlice(p.Get("ports"))
	for _, port := range ports {
		ret = append(ret, newMessagePort(port))
	}

	return ret
}

func (p *messageEventImpl) InitMessageEvent(typ string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("initMessageEvent", typ)
	case 1:
		if bubbles, ok := args[0].(bool); ok {
			p.Call("initMessageEvent", typ, bubbles)
		}
	case 2:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.Call("initMessageEvent", typ, bubbles, cancelable)
			}
		}
	case 3:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.Call("initMessageEvent", typ, bubbles, cancelable, args[2])
			}
		}
	case 4:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				if origin, ok := args[3].(string); ok {
					p.Call("initMessageEvent", typ, bubbles, cancelable, args[2], origin)
				}
			}
		}
	case 5:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				if origin, ok := args[3].(string); ok {
					if lastEventId, ok := args[4].(string); ok {
						p.Call("initMessageEvent", typ, bubbles, cancelable, args[2], origin, lastEventId)
					}
				}
			}
		}
	case 6:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				if origin, ok := args[3].(string); ok {
					if lastEventId, ok := args[4].(string); ok {
						if source, ok := args[5].(MessageEventSource); ok {
							p.Call("initMessageEvent", typ, bubbles, cancelable, args[2], origin, lastEventId, source.JSValue())
						}
					}
				}
			}
		}
	}
}

// -------------8<---------------------------------------

type messageEventSourceImpl struct {
	*objectImpl
}

func newMessageEventSource(v js.Value) MessageEventSource {
	if isNil(v) {
		return nil
	}

	return &messageEventSourceImpl{
		objectImpl: newObjectImpl(v),
	}
}

// -------------8<---------------------------------------

func NewBroadcastChannel(name string) BroadcastChannel {
	jsBroadcastChannel := js.Global().Get("BroadcastChannel")
	if isNil(jsBroadcastChannel) {
		return nil
	}

	return newBroadcastChannel(jsBroadcastChannel.New(name))
}

func NewMessageChannel() MessageChannel {
	jsMessageChannel := js.Global().Get("MessageChannel")
	if isNil(jsMessageChannel) {
		return nil
	}

	return newMessageChannel(jsMessageChannel.New())
}
