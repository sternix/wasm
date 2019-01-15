// +build js,wasm

package wasm

import (
	"fmt"
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

func NewEvent(typ string, ei ...EventInit) Event {
	if jsEvent := jsGlobal.get("Event"); jsEvent.valid() {
		switch len(ei) {
		case 0:
			return wrapEvent(jsEvent.jsNew(typ))
		default:
			return wrapEvent(jsEvent.jsNew(typ, ei[0].toJSObject()))
		}
	}
	return nil
}

func NewCustomEvent(typ string, cei ...CustomEventInit) CustomEvent {
	if jsCustomEvent := jsGlobal.get("CustomEvent"); jsCustomEvent.valid() {
		switch len(cei) {
		case 0:
			return wrapCustomEvent(jsCustomEvent.jsNew(typ))
		default:
			return wrapCustomEvent(jsCustomEvent.jsNew(typ, cei[0].toJSObject()))
		}
	}
	return nil
}

func NewFocusEvent(typ string, ini ...FocusEventInit) FocusEvent {
	if jsFocusEvent := jsGlobal.get("FocusEvent"); jsFocusEvent.valid() {
		switch len(ini) {
		case 0:
			return wrapFocusEvent(jsFocusEvent.jsNew(typ))
		default:
			return wrapFocusEvent(jsFocusEvent.jsNew(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewMouseEvent(typ string, ini ...MouseEventInit) MouseEvent {
	if jsMouseEvent := jsGlobal.get("MouseEvent"); jsMouseEvent.valid() {
		switch len(ini) {
		case 0:
			return wrapMouseEvent(jsMouseEvent.jsNew(typ))
		default:
			return wrapMouseEvent(jsMouseEvent.jsNew(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewWheelEvent(typ string, ini ...WheelEventInit) WheelEvent {
	if jsWheelEvent := jsGlobal.get("WheelEvent"); jsWheelEvent.valid() {
		switch len(ini) {
		case 0:
			return wrapWheelEvent(jsWheelEvent.jsNew(typ))
		default:
			return wrapWheelEvent(jsWheelEvent.jsNew(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewInputEvent(typ string, ini ...InputEventInit) InputEvent {
	if jsInputEvent := jsGlobal.get("InputEvent"); jsInputEvent.valid() {
		switch len(ini) {
		case 0:
			return wrapInputEvent(jsInputEvent.jsNew(typ))
		default:
			return wrapInputEvent(jsInputEvent.jsNew(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewKeyboardEvent(typ string, ini ...KeyboardEventInit) KeyboardEvent {
	if jsKeyboardEvent := jsGlobal.get("KeyboardEvent"); jsKeyboardEvent.valid() {
		switch len(ini) {
		case 0:
			return wrapKeyboardEvent(jsKeyboardEvent.jsNew(typ))
		default:
			return wrapKeyboardEvent(jsKeyboardEvent.jsNew(typ, ini[0].toJSObject()))
		}
	}
	return nil
}

func NewErrorEvent(typ string, eei ...ErrorEventInit) ErrorEvent {
	if jsErrorEvent := jsGlobal.get("ErrorEvent"); jsErrorEvent.valid() {
		switch len(eei) {
		case 0:
			return wrapErrorEvent(jsErrorEvent.jsNew(typ))
		default:
			return wrapErrorEvent(jsErrorEvent.jsNew(typ, eei[0].toJSObject()))
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
	p.call("removeEventListener", p.typ, p.jsCb)
	p.Release()
}

// -------------8<---------------------------------------

type eventTargetImpl struct {
	Value
}

func NewEventTarget() EventTarget {
	return wrapEventTarget(jsGlobal.get("EventTarget"))
}

func wrapEventTarget(v Value) EventTarget {
	if p := newEventTargetImpl(v); p != nil {
		return p
	}
	return nil
}

func newEventTargetImpl(v Value) *eventTargetImpl {
	if v.valid() {
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
	p.call("addEventListener", event, eh.jsCb)

	return eh
}

func (p *eventTargetImpl) DispatchEvent(e Event) bool {
	return p.call("dispatchEvent", JSValue(e)).toBool()
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
	if v.valid() {
		return &eventImpl{
			Value: v,
		}
	}
	return nil
}

func (p *eventImpl) Type() string {
	return p.get("type").toString()
}

func (p *eventImpl) Target() EventTarget {
	return wrapAsEventTarget(p.get("target"))
}

func (p *eventImpl) CurrentTarget() EventTarget {
	return wrapAsEventTarget(p.get("currentTarget"))
}

func (p *eventImpl) ComposedPath() []EventTarget {
	if s := p.call("composedPath").toSlice(); s != nil {
		ret := make([]EventTarget, len(s))
		for i, v := range s {
			ret[i] = wrapAsEventTarget(v)
		}
		return ret
	}
	return nil
}

func (p *eventImpl) EventPhase() EventPhase {
	return EventPhase(p.get("eventPhase").toUint16())
}

func (p *eventImpl) StopPropagation() {
	p.call("stopPropagation")
}

func (p *eventImpl) StopImmediatePropagation() {
	p.call("stopImmediatePropagation")
}

func (p *eventImpl) Bubbles() bool {
	return p.get("bubbles").toBool()
}

func (p *eventImpl) Cancelable() bool {
	return p.get("cancelable").toBool()
}

func (p *eventImpl) PreventDefault() {
	p.call("preventDefault")
}

func (p *eventImpl) DefaultPrevented() bool {
	return p.get("defaultPrevented").toBool()
}

func (p *eventImpl) Composed() bool {
	return p.get("composed").toBool()
}

func (p *eventImpl) IsTrusted() bool {
	return p.get("isTrusted").toBool()
}

func (p *eventImpl) TimeStamp() time.Time {
	ms := int64(p.get("timeStamp").toFloat64())
	return time.Unix(0, ms*int64(time.Millisecond))
}

// -------------8<---------------------------------------

type customEventImpl struct {
	*eventImpl
}

func wrapCustomEvent(v Value) CustomEvent {
	if v.valid() {
		return &customEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *customEventImpl) Detail() interface{} {
	return Wrap(p.get("detail"))
}

func (p *customEventImpl) InitCustomEvent(typ string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("initCustomEvent", typ)
	case 1:
		if bubbles, ok := args[0].(bool); ok {
			p.call("initCustomEvent", typ, bubbles)
		}
	case 2:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.call("initCustomEvent", typ, bubbles, cancelable)
			}
		}
	case 3:
		if bubbles, ok := args[0].(bool); ok {
			if cancelable, ok := args[1].(bool); ok {
				p.call("initCustomEvent", typ, bubbles, cancelable, args[2])
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
	if v.valid() {
		return &windowOrWorkerGlobalScopeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *windowOrWorkerGlobalScopeImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *windowOrWorkerGlobalScopeImpl) Btoa(btoa string) string {
	return p.call("btoa", btoa).toString()
}

func (p *windowOrWorkerGlobalScopeImpl) Atob(atob string) string {
	return p.call("atob", atob).toString()
}

func (p *windowOrWorkerGlobalScopeImpl) SetTimeout(cb TimerCallback, delay int) int {
	return p.call("setTimeout", cb.jsCallback(), delay).toInt()
}

func (p *windowOrWorkerGlobalScopeImpl) ClearTimeout(handle int) {
	p.call("clearTimeout", handle)
}

func (p *windowOrWorkerGlobalScopeImpl) SetInterval(cb TimerCallback, delay int) int {
	return p.call("setInterval", cb.jsCallback(), delay).toInt()
}

func (p *windowOrWorkerGlobalScopeImpl) ClearInterval(handle int) {
	p.call("clearInterval", handle)
}

func (p *windowOrWorkerGlobalScopeImpl) QueueMicrotask(vfn VoidFunction) {
	p.call("queueMicrotask", vfn.jsCallback())
}

func (p *windowOrWorkerGlobalScopeImpl) CreateImageBitmap(image ImageBitmapSource, options ...ImageBitmapOptions) func() (ImageBitmap, error) {
	return func() (ImageBitmap, error) {
		var (
			result Value
			ok     bool
		)

		switch len(options) {
		case 0:
			result, ok = await(p.call("createImageBitmap", JSValue(image)))
		default:
			result, ok = await(p.call("createImageBitmap", JSValue(image), options[0].toJSObject()))
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
			result, ok = await(p.call("createImageBitmap", JSValue(image), sx, sy, sw, sh))
		default:
			result, ok = await(p.call("createImageBitmap", JSValue(image), sx, sy, sw, sh, options[0].toJSObject()))
		}

		if ok {
			return wrapImageBitmap(result), nil
		}

		return nil, wrapDOMException(result)
	}
}

func (p *windowOrWorkerGlobalScopeImpl) IndexedDB() IDBFactory {
	return wrapIDBFactory(p.get("indexedDB"))
}

func (p *windowOrWorkerGlobalScopeImpl) Fetch(input RequestInfo, ri ...RequestInit) func() (Response, error) {
	return func() (Response, error) {
		var in js.Value
		switch x := input.(type) {
		case string, Request:
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
			result, ok = await(p.call("fetch", in))
		default:
			result, ok = await(p.call("fetch", in, ri[0].toJSObject()))
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
	if v.valid() {
		return &uiEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *uiEventImpl) View() Window {
	return wrapWindow(p.get("view"))
}

func (p *uiEventImpl) Detail() int {
	return p.get("detail").toInt()
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
	if v.valid() {
		return &mouseEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *mouseEventImpl) ScreenX() float64 {
	return p.get("screenX").toFloat64()
}

func (p *mouseEventImpl) ScreenY() float64 {
	return p.get("screenY").toFloat64()
}

func (p *mouseEventImpl) ClientX() float64 {
	return p.get("clientX").toFloat64()
}

func (p *mouseEventImpl) ClientY() float64 {
	return p.get("clientY").toFloat64()
}

func (p *mouseEventImpl) CtrlKey() bool {
	return p.get("ctrlKey").toBool()
}

func (p *mouseEventImpl) ShiftKey() bool {
	return p.get("shiftKey").toBool()
}

func (p *mouseEventImpl) AltKey() bool {
	return p.get("altKey").toBool()
}

func (p *mouseEventImpl) MetaKey() bool {
	return p.get("metaKey").toBool()
}

func (p *mouseEventImpl) Button() int {
	return p.get("button").toInt()
}

func (p *mouseEventImpl) Buttons() int {
	return int(p.get("buttons").toInt())
}

func (p *mouseEventImpl) RelatedTarget() EventTarget {
	return wrapAsEventTarget(p.get("relatedTarget"))
}

func (p *mouseEventImpl) ModifierState(keyArg string) bool {
	return p.call("getModifierState", keyArg).toBool()
}

func (p *mouseEventImpl) PageX() float64 {
	return p.get("pageX").toFloat64()
}

func (p *mouseEventImpl) PageY() float64 {
	return p.get("pageY").toFloat64()
}

func (p *mouseEventImpl) X() float64 {
	return p.get("x").toFloat64()
}

func (p *mouseEventImpl) Y() float64 {
	return p.get("y").toFloat64()
}

func (p *mouseEventImpl) OffsetX() float64 {
	return p.get("offsetX").toFloat64()
}

func (p *mouseEventImpl) OffsetY() float64 {
	return p.get("offsetY").toFloat64()
}

// -------------8<---------------------------------------

type focusEventImpl struct {
	*uiEventImpl
}

func wrapFocusEvent(v Value) FocusEvent {
	if v.valid() {
		return &focusEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *focusEventImpl) RelatedTarget() EventTarget {
	return newEventTargetImpl(p.get("relatedTarget"))
}

// -------------8<---------------------------------------

type wheelEventImpl struct {
	*mouseEventImpl
}

func wrapWheelEvent(v Value) WheelEvent {
	if v.valid() {
		return &wheelEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *wheelEventImpl) DeltaX() float64 {
	return p.get("deltaX").toFloat64()
}

func (p *wheelEventImpl) DeltaY() float64 {
	return p.get("deltaY").toFloat64()
}

func (p *wheelEventImpl) DeltaZ() float64 {
	return p.get("deltaZ").toFloat64()
}

func (p *wheelEventImpl) DeltaMode() WheelEventDeltaMode {
	return WheelEventDeltaMode(p.get("deltaMode").toInt())
}

// -------------8<---------------------------------------

type inputEventImpl struct {
	*uiEventImpl
}

func wrapInputEvent(v Value) InputEvent {
	if v.valid() {
		return &inputEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *inputEventImpl) Data() string {
	return p.get("data").toString()
}

func (p *inputEventImpl) IsComposing() bool {
	return p.get("isComposing").toBool()
}

// -------------8<---------------------------------------

type keyboardEventImpl struct {
	*uiEventImpl
}

func wrapKeyboardEvent(v Value) KeyboardEvent {
	if v.valid() {
		return &keyboardEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *keyboardEventImpl) Key() string {
	return p.get("key").toString()
}

func (p *keyboardEventImpl) Code() string {
	return p.get("code").toString()
}

func (p *keyboardEventImpl) Location() KeyLocationCode {
	return KeyLocationCode(p.get("location").toUint())
}

func (p *keyboardEventImpl) CtrlKey() bool {
	return p.get("ctrlKey").toBool()
}

func (p *keyboardEventImpl) ShiftKey() bool {
	return p.get("shiftKey").toBool()
}

func (p *keyboardEventImpl) AltKey() bool {
	return p.get("altKey").toBool()
}

func (p *keyboardEventImpl) MetaKey() bool {
	return p.get("metaKey").toBool()
}

func (p *keyboardEventImpl) Repeat() bool {
	return p.get("repeat").toBool()
}

func (p *keyboardEventImpl) IsComposing() bool {
	return p.get("isComposing").toBool()
}

func (p *keyboardEventImpl) ModifierState(keyArg string) bool {
	return p.call("getModifierState", keyArg).toBool()
}

// -------------8<---------------------------------------

type compositionEventImpl struct {
	*uiEventImpl
}

func wrapCompositionEvent(v Value) CompositionEvent {
	if v.valid() {
		return &compositionEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *compositionEventImpl) Data() string {
	return p.get("data").toString()
}

// -------------8<---------------------------------------

type errorEventImpl struct {
	*eventImpl
}

func wrapErrorEvent(v Value) ErrorEvent {
	if v.valid() {
		return &errorEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *errorEventImpl) Message() string {
	return p.get("message").toString()
}

func (p *errorEventImpl) Filename() string {
	return p.get("filename").toString()
}

func (p *errorEventImpl) Lineno() int {
	return p.get("lineno").toInt()
}

func (p *errorEventImpl) Colno() int {
	return p.get("colno").toInt()
}

func (p *errorEventImpl) Error() string {
	return p.get("error").toString() // TODO: test this
}

// -------------8<---------------------------------------

type transitionEventImpl struct {
	*eventImpl
}

func NewTransitionEvent(typ string, tei ...TransitionEventInit) TransitionEvent {
	if jsTe := jsGlobal.get("TransitionEvent"); jsTe.valid() {
		switch len(tei) {
		case 0:
			return wrapTransitionEvent(jsTe.jsNew(typ))
		default:
			return wrapTransitionEvent(jsTe.jsNew(typ, tei[0].toJSObject()))
		}
	}
	return nil
}

func wrapTransitionEvent(v Value) TransitionEvent {
	if v.valid() {
		return &transitionEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *transitionEventImpl) PropertyName() string {
	return p.get("propertyName").toString()
}

func (p *transitionEventImpl) ElapsedTime() float64 {
	return p.get("elapsedTime").toFloat64()
}

func (p *transitionEventImpl) PseudoElement() string {
	return p.get("pseudoElement").toString()
}

// -------------8<---------------------------------------

type pointerEventImpl struct {
	*mouseEventImpl
}

func NewPointerEvent(typ string, pei ...PointerEventInit) PointerEvent {
	if jsPe := jsGlobal.get("PointerEvent"); jsPe.valid() {
		switch len(pei) {
		case 0:
			return wrapPointerEvent(jsPe.jsNew(typ))
		default:
			return wrapPointerEvent(jsPe.jsNew(typ, pei[0].toJSObject()))
		}
	}
	return nil
}

func wrapPointerEvent(v Value) PointerEvent {
	if v.valid() {
		return &pointerEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *pointerEventImpl) PointerId() int {
	return p.get("pointerId").toInt()
}

func (p *pointerEventImpl) Width() float64 {
	return p.get("width").toFloat64()
}

func (p *pointerEventImpl) Height() float64 {
	return p.get("height").toFloat64()
}

func (p *pointerEventImpl) Pressure() float64 {
	return p.get("pressure").toFloat64()
}

func (p *pointerEventImpl) TangentialPressure() float64 {
	return p.get("tangentialPressure").toFloat64()
}

func (p *pointerEventImpl) TiltX() int {
	return p.get("tiltX").toInt()
}

func (p *pointerEventImpl) TiltY() int {
	return p.get("tiltY").toInt()
}

func (p *pointerEventImpl) Twist() int {
	return p.get("twist").toInt()
}

func (p *pointerEventImpl) PointerType() string {
	return p.get("pointerType").toString()
}

func (p *pointerEventImpl) IsPrimary() bool {
	return p.get("isPrimary").toBool()
}
