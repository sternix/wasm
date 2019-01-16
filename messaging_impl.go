// +build js,wasm

package wasm

// -------------8<---------------------------------------

type messageChannelImpl struct {
	Value
}

func wrapMessageChannel(v Value) MessageChannel {
	if v.valid() {
		return &messageChannelImpl{
			Value: v,
		}
	}
	return nil
}

func (p *messageChannelImpl) Port1() MessagePort {
	return wrapMessagePort(p.get("port1"))
}

func (p *messageChannelImpl) Port2() MessagePort {
	return wrapMessagePort(p.get("port2"))
}

// -------------8<---------------------------------------

type messagePortImpl struct {
	*eventTargetImpl
}

func wrapMessagePort(v Value) MessagePort {
	if v.valid() {
		return &messagePortImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

// TODO optional sequence<object> transfer = [] omitted
func (p *messagePortImpl) PostMessage(message interface{}) {
	//XXX: panicable
	p.call("postMessage", message)
}

func (p *messagePortImpl) Start() {
	p.call("start")
}

func (p *messagePortImpl) Close() {
	p.call("close")
}

func (p *messagePortImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return p.On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *messagePortImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

// -------------8<---------------------------------------

type broadcastChannelImpl struct {
	*eventTargetImpl
}

func wrapBroadcastChannel(v Value) BroadcastChannel {
	if v.valid() {
		return &broadcastChannelImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *broadcastChannelImpl) Name() string {
	return p.get("name").toString()
}

func (p *broadcastChannelImpl) PostMessage(message interface{}) {
	p.call("postMessage", message)
}

func (p *broadcastChannelImpl) Close() {
	p.call("close")
}

func (p *broadcastChannelImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return p.On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *broadcastChannelImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

// -------------8<---------------------------------------

type messageEventImpl struct {
	*eventImpl
}

func wrapMessageEvent(v Value) MessageEvent {
	if v.valid() {
		return &messageEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *messageEventImpl) Data() interface{} {
	return Wrap(p.get("data"))
}

func (p *messageEventImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *messageEventImpl) LastEventId() string {
	return p.get("lastEventId").toString()
}

func (p *messageEventImpl) Source() MessageEventSource {
	if v := p.get("source"); v.valid() {
		if v.instanceOf(jsWindowProxy) {
			return wrapWindowProxy(v)
		} else if v.instanceOf(jsMessagePort) { // TODO remove from util
			return wrapMessagePort(v)
		} /* TODO: ServiceWorker  else if v.InstanceOf(jsServiceWorker) {
			return wrapServiceWorker(v)
		}*/
	}
	return nil
}

func (p *messageEventImpl) Ports() []MessagePort {
	if ports := p.get("ports").toSlice(); ports != nil {
		var ret []MessagePort
		for _, port := range ports {
			ret = append(ret, wrapMessagePort(port))
		}
		return ret
	}
	return nil
}

func (p *messageEventImpl) InitMessageEvent(typ string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("initMessageEvent", typ)
	case 1:
		if bubbles, ok := args[0].(bool); ok {
			p.call("initMessageEvent", typ, bubbles)
		}
	case 2:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.call("initMessageEvent", typ, bubbles, cancelable)
			}
		}
	case 3:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.call("initMessageEvent", typ, bubbles, cancelable, args[2])
			}
		}
	case 4:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				if origin, ok := args[3].(string); ok {
					p.call("initMessageEvent", typ, bubbles, cancelable, args[2], origin)
				}
			}
		}
	case 5:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				if origin, ok := args[3].(string); ok {
					if lastEventId, ok := args[4].(string); ok {
						p.call("initMessageEvent", typ, bubbles, cancelable, args[2], origin, lastEventId)
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
							p.call("initMessageEvent", typ, bubbles, cancelable, args[2], origin, lastEventId, JSValueOf(source))
						}
					}
				}
			}
		}
	}
}

// -------------8<---------------------------------------

type messageEventSourceImpl struct {
	Value
}

func wrapMessageEventSource(v Value) MessageEventSource {
	if v.valid() {
		return &messageEventSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

func NewBroadcastChannel(name string) BroadcastChannel {
	if jsBroadcastChannel := jsGlobal.get("BroadcastChannel"); jsBroadcastChannel.valid() {
		return wrapBroadcastChannel(jsBroadcastChannel.jsNew(name))
	}
	return nil
}

func NewMessageChannel() MessageChannel {
	if jsMessageChannel := jsGlobal.get("MessageChannel"); jsMessageChannel.valid() {
		return wrapMessageChannel(jsMessageChannel.jsNew())
	}
	return nil
}
