// +build js,wasm

package wasm

import (
	"fmt"
	"time"
)

// -------------8<---------------------------------------

func NewEvent(typ string, ei ...EventInit) Event {
	if jsEvent := jsGlobal.Get("Event"); jsEvent.Valid() {
		switch len(ei) {
		case 0:
			return wrapEvent(jsEvent.New(typ))
		default:
			return wrapEvent(jsEvent.New(typ, ei[0].toJSObject()))
		}
	}
	return nil
}

func NewCustomEvent(typ string, cei ...CustomEventInit) CustomEvent {
	if jsCustomEvent := jsGlobal.Get("CustomEvent"); jsCustomEvent.Valid() {
		switch len(cei) {
		case 0:
			return wrapCustomEvent(jsCustomEvent.New(typ))
		default:
			return wrapCustomEvent(jsCustomEvent.New(typ, cei[0].toJSObject()))
		}
	}
	return nil
}

func NewFocusEvent(typ string, ini ...FocusEventInit) FocusEvent {
	if jsFocusEvent := jsGlobal.Get("FocusEvent"); jsFocusEvent.Valid() {
		switch len(ini) {
		case 0:
			return wrapFocusEvent(jsFocusEvent.New(typ))
		default:
			return wrapFocusEvent(jsFocusEvent.New(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewMouseEvent(typ string, ini ...MouseEventInit) MouseEvent {
	if jsMouseEvent := jsGlobal.Get("MouseEvent"); jsMouseEvent.Valid() {
		switch len(ini) {
		case 0:
			return wrapMouseEvent(jsMouseEvent.New(typ))
		default:
			return wrapMouseEvent(jsMouseEvent.New(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewWheelEvent(typ string, ini ...WheelEventInit) WheelEvent {
	if jsWheelEvent := jsGlobal.Get("WheelEvent"); jsWheelEvent.Valid() {
		switch len(ini) {
		case 0:
			return wrapWheelEvent(jsWheelEvent.New(typ))
		default:
			return wrapWheelEvent(jsWheelEvent.New(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewInputEvent(typ string, ini ...InputEventInit) InputEvent {
	if jsInputEvent := jsGlobal.Get("InputEvent"); jsInputEvent.Valid() {
		switch len(ini) {
		case 0:
			return wrapInputEvent(jsInputEvent.New(typ))
		default:
			return wrapInputEvent(jsInputEvent.New(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewKeyboardEvent(typ string, ini ...KeyboardEventInit) KeyboardEvent {
	if jsKeyboardEvent := jsGlobal.Get("KeyboardEvent"); jsKeyboardEvent.Valid() {
		switch len(ini) {
		case 0:
			return wrapKeyboardEvent(jsKeyboardEvent.New(typ))
		default:
			return wrapKeyboardEvent(jsKeyboardEvent.New(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewErrorEvent(typ string, eei ...ErrorEventInit) ErrorEvent {
	if jsErrorEvent := jsGlobal.Get("ErrorEvent"); jsErrorEvent.Valid() {
		switch len(eei) {
		case 0:
			return wrapErrorEvent(jsErrorEvent.New(typ))
		default:
			return wrapErrorEvent(jsErrorEvent.New(typ, eei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type eventHandlerImpl struct {
	Value
	jsCb Func
	fn   func(Event)
	typ  string
}

func (p *eventHandlerImpl) Type() string {
	return p.typ
}

func (p *eventHandlerImpl) jsFunc(this Value, args []Value) interface{} {
	if len(args) > 0 {
		p.Handle(wrapAsEvent(args[0]))
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
	p.Call("removeEventListener", p.typ, p.jsCb)
	p.Release()
}

// -------------8<---------------------------------------

type eventTargetImpl struct {
	Value
}

func wrapEventTarget(v Value) EventTarget {
	if p := newEventTargetImpl(v); p != nil {
		return p
	}
	return nil
}

func newEventTargetImpl(v Value) *eventTargetImpl {
	if v.Valid() {
		return &eventTargetImpl{
			Value: v,
		}
	}
	return nil
}

func (p *eventTargetImpl) On(event string, fn func(ev Event)) EventHandler {
	eh := &eventHandlerImpl{
		Value: p.Value,
		fn:    fn,
		typ:   event,
	}

	eh.jsCb = FuncOf(eh.jsFunc)
	p.Call("addEventListener", event, eh.jsCb)

	return eh
}

func (p *eventTargetImpl) DispatchEvent(e Event) bool {
	return p.Call("dispatchEvent", JSValue(e)).Bool()
}

// -------------8<---------------------------------------

type eventImpl struct {
	Value
}

func wrapEvent(v Value) Event {
	if p := newEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newEventImpl(v Value) *eventImpl {
	if v.Valid() {
		return &eventImpl{
			Value: v,
		}
	}
	return nil
}

func (p *eventImpl) Type() string {
	return p.Get("type").String()
}

func (p *eventImpl) Target() EventTarget {
	return wrapAsEventTarget(p.Get("target"))
}

func (p *eventImpl) CurrentTarget() EventTarget {
	return wrapAsEventTarget(p.Get("currentTarget"))
}

func (p *eventImpl) ComposedPath() []EventTarget {
	if s := p.Call("composedPath").ToSlice(); s != nil {
		ret := make([]EventTarget, len(s))
		for i, v := range s {
			ret[i] = wrapAsEventTarget(v)
		}
		return ret
	}
	return nil
}

func (p *eventImpl) EventPhase() EventPhase {
	return EventPhase(p.Get("eventPhase").Int())
}

func (p *eventImpl) StopPropagation() {
	p.Call("stopPropagation")
}

func (p *eventImpl) StopImmediatePropagation() {
	p.Call("stopImmediatePropagation")
}

func (p *eventImpl) Bubbles() bool {
	return p.Get("bubbles").Bool()
}

func (p *eventImpl) Cancelable() bool {
	return p.Get("cancelable").Bool()
}

func (p *eventImpl) PreventDefault() {
	p.Call("preventDefault")
}

func (p *eventImpl) DefaultPrevented() bool {
	return p.Get("defaultPrevented").Bool()
}

func (p *eventImpl) Composed() bool {
	return p.Get("composed").Bool()
}

func (p *eventImpl) IsTrusted() bool {
	return p.Get("isTrusted").Bool()
}

func (p *eventImpl) TimeStamp() time.Time {
	ms := int64(p.Get("timeStamp").Float())
	return time.Unix(0, ms*int64(time.Millisecond))
}

// -------------8<---------------------------------------

type customEventImpl struct {
	*eventImpl
}

func wrapCustomEvent(v Value) CustomEvent {
	if v.Valid() {
		return &customEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *customEventImpl) Detail() interface{} {
	return Wrap(p.Get("detail"))
}

func (p *customEventImpl) InitCustomEvent(typ string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("initCustomEvent", typ)
	case 1:
		if bubbles, ok := args[0].(bool); ok {
			p.Call("initCustomEvent", typ, bubbles)
		}
	case 2:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.Call("initCustomEvent", typ, bubbles, cancelable)
			}
		}
	case 3:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.Call("initCustomEvent", typ, bubbles, cancelable, args[2])
			}
		}
	}
}

// -------------8<---------------------------------------

var _ GlobalEventHandlers = &globalEventHandlersImpl{}

type globalEventHandlersImpl struct {
	*eventTargetImpl
}

func newGlobalEventHandlersImpl(et *eventTargetImpl) *globalEventHandlersImpl {
	return &globalEventHandlersImpl{
		eventTargetImpl: et,
	}
}

func (p *globalEventHandlersImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

func (p *globalEventHandlersImpl) OnAuxClick(fn func(Event)) EventHandler {
	return p.On("auxclick", fn)
}

func (p *globalEventHandlersImpl) OnBlur(fn func(Event)) EventHandler {
	return p.On("blur", fn)
}

func (p *globalEventHandlersImpl) OnCancel(fn func(Event)) EventHandler {
	return p.On("cancel", fn)
}

func (p *globalEventHandlersImpl) OnCanPlay(fn func(Event)) EventHandler {
	return p.On("canplay", fn)
}

func (p *globalEventHandlersImpl) OnCanPlayThrough(fn func(Event)) EventHandler {
	return p.On("canplaythrough", fn)
}

func (p *globalEventHandlersImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

func (p *globalEventHandlersImpl) OnClick(fn func(MouseEvent)) EventHandler {
	return p.On("click", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnClose(fn func(CloseEvent)) EventHandler {
	return p.On("close", func(e Event) {
		if ce, ok := e.(CloseEvent); ok {
			fn(ce)
		}
	})
}

func (p *globalEventHandlersImpl) OnContextMenu(fn func(MouseEvent)) EventHandler {
	return p.On("contextmenu", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnCueChange(fn func(Event)) EventHandler {
	return p.On("cuechange", fn)
}

func (p *globalEventHandlersImpl) OnDblClick(fn func(MouseEvent)) EventHandler {
	return p.On("dblclick", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnDrag(fn func(DragEvent)) EventHandler {
	return p.On("drag", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragEnd(fn func(DragEvent)) EventHandler {
	return p.On("dragend", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragEnter(fn func(DragEvent)) EventHandler {
	return p.On("dragenter", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragExit(fn func(DragEvent)) EventHandler {
	return p.On("dragexit", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragLeave(fn func(DragEvent)) EventHandler {
	return p.On("dragleave", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragOver(fn func(DragEvent)) EventHandler {
	return p.On("dragover", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragStart(fn func(DragEvent)) EventHandler {
	return p.On("dragstart", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDrop(fn func(DragEvent)) EventHandler {
	return p.On("drop", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDurationChange(fn func(Event)) EventHandler {
	return p.On("durationchange", fn)
}

func (p *globalEventHandlersImpl) OnEmptied(fn func(Event)) EventHandler {
	return p.On("emptied", fn)
}

func (p *globalEventHandlersImpl) OnEnded(fn func(Event)) EventHandler {
	return p.On("ended", fn)
}

func (p *globalEventHandlersImpl) OnError(fn func(ErrorEvent)) EventHandler {
	return p.On("error", func(e Event) {
		if ee, ok := e.(ErrorEvent); ok {
			fn(ee)
		}
	})
}

func (p *globalEventHandlersImpl) OnFocus(fn func(FocusEvent)) EventHandler {
	return p.On("focus", func(e Event) {
		if fe, ok := e.(FocusEvent); ok {
			fn(fe)
		}
	})
}

func (p *globalEventHandlersImpl) OnInput(fn func(InputEvent)) EventHandler {
	return p.On("input", func(e Event) {
		if ie, ok := e.(InputEvent); ok {
			fn(ie)
		}
	})
}

func (p *globalEventHandlersImpl) OnInvalid(fn func(Event)) EventHandler {
	return p.On("invalid", fn)
}

func (p *globalEventHandlersImpl) OnKeyDown(fn func(KeyboardEvent)) EventHandler {
	return p.On("keydown", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnKeyPress(fn func(KeyboardEvent)) EventHandler {
	return p.On("keypress", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnKeyUp(fn func(KeyboardEvent)) EventHandler {
	return p.On("keyup", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoad(fn func(UIEvent)) EventHandler {
	return p.On("load", func(e Event) {
		if ue, ok := e.(UIEvent); ok {
			fn(ue)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoadedData(fn func(Event)) EventHandler {
	return p.On("loadeddata", fn)
}

func (p *globalEventHandlersImpl) OnLoadedMetadata(fn func(Event)) EventHandler {
	return p.On("loadedmetadata", fn)
}

func (p *globalEventHandlersImpl) OnLoadEnd(fn func(ProgressEvent)) EventHandler {
	return p.On("loadend", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoadStart(fn func(ProgressEvent)) EventHandler {
	return p.On("loadstart", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseDown(fn func(MouseEvent)) EventHandler {
	return p.On("mousedown", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseEnter(fn func(MouseEvent)) EventHandler {
	return p.On("mouseenter", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseLeave(fn func(MouseEvent)) EventHandler {
	return p.On("mouseleave", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseMove(fn func(MouseEvent)) EventHandler {
	return p.On("mousemove", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseOut(fn func(MouseEvent)) EventHandler {
	return p.On("mouseout", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseOver(fn func(MouseEvent)) EventHandler {
	return p.On("mouseover", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})

}

func (p *globalEventHandlersImpl) OnMouseUp(fn func(MouseEvent)) EventHandler {
	return p.On("mouseup", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnWheel(fn func(WheelEvent)) EventHandler {
	return p.On("wheel", func(e Event) {
		if we, ok := e.(WheelEvent); ok {
			fn(we)
		}
	})
}

func (p *globalEventHandlersImpl) OnPause(fn func(Event)) EventHandler {
	return p.On("pause", fn)
}

func (p *globalEventHandlersImpl) OnPlay(fn func(Event)) EventHandler {
	return p.On("play", fn)
}

func (p *globalEventHandlersImpl) OnPlaying(fn func(Event)) EventHandler {
	return p.On("playing", fn)
}

func (p *globalEventHandlersImpl) OnProgress(fn func(ProgressEvent)) EventHandler {
	return p.On("progress", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnRateChange(fn func(Event)) EventHandler {
	return p.On("ratechange", fn)
}

func (p *globalEventHandlersImpl) OnReset(fn func(Event)) EventHandler {
	return p.On("reset", fn)
}

func (p *globalEventHandlersImpl) OnResize(fn func(Event)) EventHandler {
	return p.On("resize", fn)
}

func (p *globalEventHandlersImpl) OnScroll(fn func(Event)) EventHandler {
	return p.On("scroll", fn)
}

func (p *globalEventHandlersImpl) OnSecurityPolicyViolation(fn func(Event)) EventHandler {
	return p.On("securitypolicyviolation", fn)
}

func (p *globalEventHandlersImpl) OnSeeked(fn func(Event)) EventHandler {
	return p.On("seeked", fn)
}

func (p *globalEventHandlersImpl) OnSeeking(fn func(Event)) EventHandler {
	return p.On("seeking", fn)
}

func (p *globalEventHandlersImpl) OnSelect(fn func(Event)) EventHandler {
	return p.On("select", fn)
}

func (p *globalEventHandlersImpl) OnStalled(fn func(Event)) EventHandler {
	return p.On("stalled", fn)
}

func (p *globalEventHandlersImpl) OnSubmit(fn func(Event)) EventHandler {
	return p.On("submit", fn)
}

func (p *globalEventHandlersImpl) OnSuspend(fn func(Event)) EventHandler {
	return p.On("suspend", fn)
}

func (p *globalEventHandlersImpl) OnTimeUpdate(fn func(Event)) EventHandler {
	return p.On("timeupdate", fn)
}

func (p *globalEventHandlersImpl) OnToggle(fn func(Event)) EventHandler {
	return p.On("toggle", fn)
}

func (p *globalEventHandlersImpl) OnVolumeChange(fn func(Event)) EventHandler {
	return p.On("volumechange", fn)
}

func (p *globalEventHandlersImpl) OnWaiting(fn func(Event)) EventHandler {
	return p.On("waiting", fn)
}

func (p *globalEventHandlersImpl) OnTouchStart(fn func(TouchEvent)) EventHandler {
	return p.On("touchstart", func(e Event) {
		if te, ok := e.(TouchEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTouchEnd(fn func(TouchEvent)) EventHandler {
	return p.On("touchend", func(e Event) {
		if te, ok := e.(TouchEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTouchMove(fn func(TouchEvent)) EventHandler {
	return p.On("touchmove", func(e Event) {
		if te, ok := e.(TouchEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTouchCancel(fn func(TouchEvent)) EventHandler {
	return p.On("touchcancel", func(e Event) {
		if te, ok := e.(TouchEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTransitionRun(fn func(TransitionEvent)) EventHandler {
	return p.On("transitionrun", func(e Event) {
		if te, ok := e.(TransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTransitionStart(fn func(TransitionEvent)) EventHandler {
	return p.On("transitionstart", func(e Event) {
		if te, ok := e.(TransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTransitionEnd(fn func(TransitionEvent)) EventHandler {
	return p.On("transitionend", func(e Event) {
		if te, ok := e.(TransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnTransitionCancel(fn func(TransitionEvent)) EventHandler {
	return p.On("transitioncancel", func(e Event) {
		if te, ok := e.(TransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *globalEventHandlersImpl) OnGotPointerCapture(fn func(PointerEvent)) EventHandler {
	return p.On("gotpointercapture", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnLostPointerCapture(fn func(PointerEvent)) EventHandler {
	return p.On("lostpointercapture", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerDown(fn func(PointerEvent)) EventHandler {
	return p.On("pointerdown", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerMove(fn func(PointerEvent)) EventHandler {
	return p.On("pointermove", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerUp(fn func(PointerEvent)) EventHandler {
	return p.On("pointerup", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerCancel(fn func(PointerEvent)) EventHandler {
	return p.On("pointercancel", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerOver(fn func(PointerEvent)) EventHandler {
	return p.On("pointerover", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerOut(fn func(PointerEvent)) EventHandler {
	return p.On("pointerout", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerEnter(fn func(PointerEvent)) EventHandler {
	return p.On("pointerenter", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnPointerLeave(fn func(PointerEvent)) EventHandler {
	return p.On("pointerleave", func(e Event) {
		if pe, ok := e.(PointerEvent); ok {
			fn(pe)
		}
	})
}

// -------------8<---------------------------------------

var _ DocumentAndElementEventHandlers = &documentAndElementEventHandlersImpl{}

type documentAndElementEventHandlersImpl struct {
	*eventTargetImpl
}

func newDocumentAndElementEventHandlersImpl(et *eventTargetImpl) *documentAndElementEventHandlersImpl {
	return &documentAndElementEventHandlersImpl{
		eventTargetImpl: et,
	}
}

func (p *documentAndElementEventHandlersImpl) OnCopy(fn func(Event)) EventHandler {
	return p.On("copy", fn)
}

func (p *documentAndElementEventHandlersImpl) OnCut(fn func(Event)) EventHandler {
	return p.On("cut", fn)
}

func (p *documentAndElementEventHandlersImpl) OnPaste(fn func(Event)) EventHandler {
	return p.On("paste", fn)
}

// -------------8<---------------------------------------

var _ WindowOrWorkerGlobalScope = &windowOrWorkerGlobalScopeImpl{}

type windowOrWorkerGlobalScopeImpl struct {
	Value
}

func newWindowOrWorkerGlobalScopeImpl(v Value) *windowOrWorkerGlobalScopeImpl {
	if v.Valid() {
		return &windowOrWorkerGlobalScopeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *windowOrWorkerGlobalScopeImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *windowOrWorkerGlobalScopeImpl) Btoa(btoa string) string {
	return p.Call("btoa", btoa).String()
}

func (p *windowOrWorkerGlobalScopeImpl) Atob(atob string) string {
	return p.Call("atob", atob).String()
}

func (p *windowOrWorkerGlobalScopeImpl) SetTimeout(cb TimerCallback, delay int) int {
	return p.Call("setTimeout", cb.jsCallback(), delay).Int()
}

func (p *windowOrWorkerGlobalScopeImpl) ClearTimeout(handle int) {
	p.Call("clearTimeout", handle)
}

func (p *windowOrWorkerGlobalScopeImpl) SetInterval(cb TimerCallback, delay int) int {
	return p.Call("setInterval", cb.jsCallback(), delay).Int()
}

func (p *windowOrWorkerGlobalScopeImpl) ClearInterval(handle int) {
	p.Call("clearInterval", handle)
}

func (p *windowOrWorkerGlobalScopeImpl) QueueMicrotask(vfn VoidFunction) {
	p.Call("queueMicrotask", vfn.jsCallback())
}

func (p *windowOrWorkerGlobalScopeImpl) CreateImageBitmap(image ImageBitmapSource, options ...ImageBitmapOptions) func() (ImageBitmap, error) {
	return func() (ImageBitmap, error) {
		var (
			result Value
			ok     bool
		)

		switch len(options) {
		case 0:
			result, ok = await(p.Call("createImageBitmap", JSValue(image)))
		default:
			result, ok = await(p.Call("createImageBitmap", JSValue(image), options[0].toJSObject()))
		}

		if ok {
			return wrapImageBitmap(result), nil
		}

		return nil, wrapDOMException(result)
	}
}

func (p *windowOrWorkerGlobalScopeImpl) CreateImageBitmapWithSize(image ImageBitmapSource, sx int, sy int, sw int, sh int, options ...ImageBitmapOptions) func() (ImageBitmap, error) {
	return func() (ImageBitmap, error) {
		var (
			result Value
			ok     bool
		)

		switch len(options) {
		case 0:
			result, ok = await(p.Call("createImageBitmap", JSValue(image), sx, sy, sw, sh))
		default:
			result, ok = await(p.Call("createImageBitmap", JSValue(image), sx, sy, sw, sh, options[0].toJSObject()))
		}

		if ok {
			return wrapImageBitmap(result), nil
		}

		return nil, wrapDOMException(result)
	}
}

func (p *windowOrWorkerGlobalScopeImpl) IndexedDB() IDBFactory {
	return wrapIDBFactory(p.Get("indexedDB"))
}

func (p *windowOrWorkerGlobalScopeImpl) Fetch(input RequestInfo, ri ...RequestInit) func() (Response, error) {
	return func() (Response, error) {
		var in Value
		switch x := input.(type) {
		case string:
			in = ValueOf(x)
		case Request:
			in = JSValue(x)
		default:
			return nil, fmt.Errorf("Wrong parameter type for RequestInfo")
		}

		var (
			result Value
			ok     bool
		)

		switch len(ri) {
		case 0:
			result, ok = await(p.Call("fetch", in))
		default:
			result, ok = await(p.Call("fetch", in, ri[0].toJSObject()))
		}

		if ok {
			return wrapResponse(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

// -------------8<---------------------------------------

var _ WindowEventHandlers = &windowEventHandlersImpl{}

type windowEventHandlersImpl struct {
	*eventTargetImpl
}

func newWindowEventHandlersImpl(et *eventTargetImpl) *windowEventHandlersImpl {
	return &windowEventHandlersImpl{
		eventTargetImpl: et,
	}
}

func (p *windowEventHandlersImpl) OnAfterPrint(fn func(Event)) EventHandler {
	return p.On("afterprint", fn)
}

func (p *windowEventHandlersImpl) OnBeforePrint(fn func(Event)) EventHandler {
	return p.On("beforeprint", fn)
}

func (p *windowEventHandlersImpl) OnBeforeUnload(fn func(BeforeUnloadEvent)) EventHandler {
	return p.On("beforeunload", func(e Event) {
		if be, ok := e.(BeforeUnloadEvent); ok {
			fn(be)
		}
	})
}

func (p *windowEventHandlersImpl) OnHashChange(fn func(HashChangeEvent)) EventHandler {
	return p.On("hashchange", func(e Event) {
		if hce, ok := e.(HashChangeEvent); ok {
			fn(hce)
		}
	})
}

func (p *windowEventHandlersImpl) OnLanguageChange(fn func(Event)) EventHandler {
	return p.On("languagechange", fn)
}

func (p *windowEventHandlersImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return p.On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *windowEventHandlersImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

func (p *windowEventHandlersImpl) OnOffline(fn func(Event)) EventHandler {
	return p.On("offline", fn)
}

func (p *windowEventHandlersImpl) OnOnline(fn func(Event)) EventHandler {
	return p.On("online", fn)
}

func (p *windowEventHandlersImpl) OnPageHide(fn func(PageTransitionEvent)) EventHandler {
	return p.On("pagehide", func(e Event) {
		if te, ok := e.(PageTransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *windowEventHandlersImpl) OnPageShow(fn func(PageTransitionEvent)) EventHandler {
	return p.On("pageshow", func(e Event) {
		if te, ok := e.(PageTransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *windowEventHandlersImpl) OnPopState(fn func(PopStateEvent)) EventHandler {
	return p.On("popstate", func(e Event) {
		if pse, ok := e.(PopStateEvent); ok {
			fn(pse)
		}
	})
}

func (p *windowEventHandlersImpl) OnRejectionHandled(fn func(Event)) EventHandler {
	return p.On("rejectionhandled", fn)
}

func (p *windowEventHandlersImpl) OnStorage(fn func(StorageEvent)) EventHandler {
	return p.On("storage", func(e Event) {
		if se, ok := e.(StorageEvent); ok {
			fn(se)
		}
	})
}

func (p *windowEventHandlersImpl) OnUnhandledRejection(fn func(Event)) EventHandler {
	return p.On("unhandledrejection", fn)
}

func (p *windowEventHandlersImpl) OnUnload(fn func(Event)) EventHandler {
	return p.On("unload", fn)
}

// -------------8<---------------------------------------

type uiEventImpl struct {
	*eventImpl
}

func wrapUIEvent(v Value) UIEvent {
	if p := newUIEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newUIEventImpl(v Value) *uiEventImpl {
	if v.Valid() {
		return &uiEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *uiEventImpl) View() Window {
	return wrapWindow(p.Get("view"))
}

func (p *uiEventImpl) Detail() int {
	return p.Get("detail").Int()
}

// -------------8<---------------------------------------

type mouseEventImpl struct {
	*uiEventImpl
}

func wrapMouseEvent(v Value) MouseEvent {
	if p := newMouseEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newMouseEventImpl(v Value) *mouseEventImpl {
	if v.Valid() {
		return &mouseEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *mouseEventImpl) ScreenX() float64 {
	return p.Get("screenX").Float()
}

func (p *mouseEventImpl) ScreenY() float64 {
	return p.Get("screenY").Float()
}

func (p *mouseEventImpl) ClientX() float64 {
	return p.Get("clientX").Float()
}

func (p *mouseEventImpl) ClientY() float64 {
	return p.Get("clientY").Float()
}

func (p *mouseEventImpl) CtrlKey() bool {
	return p.Get("ctrlKey").Bool()
}

func (p *mouseEventImpl) ShiftKey() bool {
	return p.Get("shiftKey").Bool()
}

func (p *mouseEventImpl) AltKey() bool {
	return p.Get("altKey").Bool()
}

func (p *mouseEventImpl) MetaKey() bool {
	return p.Get("metaKey").Bool()
}

func (p *mouseEventImpl) Button() int {
	return p.Get("button").Int()
}

func (p *mouseEventImpl) Buttons() int {
	return int(p.Get("buttons").Int())
}

func (p *mouseEventImpl) RelatedTarget() EventTarget {
	return wrapAsEventTarget(p.Get("relatedTarget"))
}

func (p *mouseEventImpl) ModifierState(keyArg string) bool {
	return p.Call("getModifierState", keyArg).Bool()
}

func (p *mouseEventImpl) PageX() float64 {
	return p.Get("pageX").Float()
}

func (p *mouseEventImpl) PageY() float64 {
	return p.Get("pageY").Float()
}

func (p *mouseEventImpl) X() float64 {
	return p.Get("x").Float()
}

func (p *mouseEventImpl) Y() float64 {
	return p.Get("y").Float()
}

func (p *mouseEventImpl) OffsetX() float64 {
	return p.Get("offsetX").Float()
}

func (p *mouseEventImpl) OffsetY() float64 {
	return p.Get("offsetY").Float()
}

// -------------8<---------------------------------------

type focusEventImpl struct {
	*uiEventImpl
}

func wrapFocusEvent(v Value) FocusEvent {
	if v.Valid() {
		return &focusEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *focusEventImpl) RelatedTarget() EventTarget {
	return newEventTargetImpl(p.Get("relatedTarget"))
}

// -------------8<---------------------------------------

type wheelEventImpl struct {
	*mouseEventImpl
}

func wrapWheelEvent(v Value) WheelEvent {
	if v.Valid() {
		return &wheelEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *wheelEventImpl) DeltaX() float64 {
	return p.Get("deltaX").Float()
}

func (p *wheelEventImpl) DeltaY() float64 {
	return p.Get("deltaY").Float()
}

func (p *wheelEventImpl) DeltaZ() float64 {
	return p.Get("deltaZ").Float()
}

func (p *wheelEventImpl) DeltaMode() WheelEventDeltaMode {
	return WheelEventDeltaMode(p.Get("deltaMode").Int())
}

// -------------8<---------------------------------------

type inputEventImpl struct {
	*uiEventImpl
}

func wrapInputEvent(v Value) InputEvent {
	if v.Valid() {
		return &inputEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *inputEventImpl) Data() string {
	return p.Get("data").String()
}

func (p *inputEventImpl) IsComposing() bool {
	return p.Get("isComposing").Bool()
}

// -------------8<---------------------------------------

type keyboardEventImpl struct {
	*uiEventImpl
}

func wrapKeyboardEvent(v Value) KeyboardEvent {
	if v.Valid() {
		return &keyboardEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *keyboardEventImpl) Key() string {
	return p.Get("key").String()
}

func (p *keyboardEventImpl) Code() string {
	return p.Get("code").String()
}

func (p *keyboardEventImpl) Location() int {
	return p.Get("location").Int()
}

func (p *keyboardEventImpl) CtrlKey() bool {
	return p.Get("ctrlKey").Bool()
}

func (p *keyboardEventImpl) ShiftKey() bool {
	return p.Get("shiftKey").Bool()
}

func (p *keyboardEventImpl) AltKey() bool {
	return p.Get("altKey").Bool()
}

func (p *keyboardEventImpl) MetaKey() bool {
	return p.Get("metaKey").Bool()
}

func (p *keyboardEventImpl) Repeat() bool {
	return p.Get("repeat").Bool()
}

func (p *keyboardEventImpl) IsComposing() bool {
	return p.Get("isComposing").Bool()
}

func (p *keyboardEventImpl) ModifierState(keyArg string) bool {
	return p.Call("getModifierState", keyArg).Bool()
}

// -------------8<---------------------------------------

type compositionEventImpl struct {
	*uiEventImpl
}

func wrapCompositionEvent(v Value) CompositionEvent {
	if v.Valid() {
		return &compositionEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *compositionEventImpl) Data() string {
	return p.Get("data").String()
}

// -------------8<---------------------------------------

type errorEventImpl struct {
	*eventImpl
}

func wrapErrorEvent(v Value) ErrorEvent {
	if v.Valid() {
		return &errorEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *errorEventImpl) Message() string {
	return p.Get("message").String()
}

func (p *errorEventImpl) Filename() string {
	return p.Get("filename").String()
}

func (p *errorEventImpl) Lineno() int {
	return p.Get("lineno").Int()
}

func (p *errorEventImpl) Colno() int {
	return p.Get("colno").Int()
}

func (p *errorEventImpl) Error() string {
	return p.Get("error").String() // TODO: test this
}

// -------------8<---------------------------------------

type transitionEventImpl struct {
	*eventImpl
}

func NewTransitionEvent(typ string, tei ...TransitionEventInit) TransitionEvent {
	if jsTe := jsGlobal.Get("TransitionEvent"); jsTe.Valid() {
		switch len(tei) {
		case 0:
			return wrapTransitionEvent(jsTe.New(typ))
		default:
			return wrapTransitionEvent(jsTe.New(typ, tei[0].toJSObject()))
		}
	}
	return nil
}

func wrapTransitionEvent(v Value) TransitionEvent {
	if v.Valid() {
		return &transitionEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *transitionEventImpl) PropertyName() string {
	return p.Get("propertyName").String()
}

func (p *transitionEventImpl) ElapsedTime() float64 {
	return p.Get("elapsedTime").Float()
}

func (p *transitionEventImpl) PseudoElement() string {
	return p.Get("pseudoElement").String()
}

// -------------8<---------------------------------------

type pointerEventImpl struct {
	*mouseEventImpl
}

func NewPointerEvent(typ string, pei ...PointerEventInit) PointerEvent {
	if jsPe := jsGlobal.Get("PointerEvent"); jsPe.Valid() {
		switch len(pei) {
		case 0:
			return wrapPointerEvent(jsPe.New(typ))
		default:
			return wrapPointerEvent(jsPe.New(typ, pei[0].toJSObject()))
		}
	}
	return nil
}

func wrapPointerEvent(v Value) PointerEvent {
	if v.Valid() {
		return &pointerEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *pointerEventImpl) PointerId() int {
	return p.Get("pointerId").Int()
}

func (p *pointerEventImpl) Width() float64 {
	return p.Get("width").Float()
}

func (p *pointerEventImpl) Height() float64 {
	return p.Get("height").Float()
}

func (p *pointerEventImpl) Pressure() float64 {
	return p.Get("pressure").Float()
}

func (p *pointerEventImpl) TangentialPressure() float64 {
	return p.Get("tangentialPressure").Float()
}

func (p *pointerEventImpl) TiltX() int {
	return p.Get("tiltX").Int()
}

func (p *pointerEventImpl) TiltY() int {
	return p.Get("tiltY").Int()
}

func (p *pointerEventImpl) Twist() int {
	return p.Get("twist").Int()
}

func (p *pointerEventImpl) PointerType() string {
	return p.Get("pointerType").String()
}

func (p *pointerEventImpl) IsPrimary() bool {
	return p.Get("isPrimary").Bool()
}
