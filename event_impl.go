// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

func NewEvent(typ string, ei ...EventInit) Event {
	jsEvent := js.Global().Get("Event")
	if isNil(jsEvent) {
		return nil
	}

	switch len(ei) {
	case 0:
		return newEvent(jsEvent.New(typ))
	default:
		return newEvent(jsEvent.New(typ, ei[0].toDict()))
	}
}

func NewCustomEvent(typ string, cei ...CustomEventInit) CustomEvent {
	jsCustomEvent := js.Global().Get("CustomEvent")
	if isNil(jsCustomEvent) {
		return nil
	}

	switch len(cei) {
	case 0:
		return newCustomEvent(jsCustomEvent.New(typ))
	default:
		return newCustomEvent(jsCustomEvent.New(typ, cei[0].toDict()))
	}
}

func NewFocusEvent(typ string, ini ...FocusEventInit) FocusEvent {
	jsFocusEvent := js.Global().Get("FocusEvent")
	if isNil(jsFocusEvent) {
		return nil
	}

	switch len(ini) {
	case 0:
		return newFocusEvent(jsFocusEvent.New(typ))
	default:
		return newFocusEvent(jsFocusEvent.New(typ, ini[0].toDict()))
	}
}

func NewMouseEvent(typ string, ini ...MouseEventInit) MouseEvent {
	jsMouseEvent := js.Global().Get("MouseEvent")
	if isNil(jsMouseEvent) {
		return nil
	}

	switch len(ini) {
	case 0:
		return newMouseEvent(jsMouseEvent.New(typ))
	default:
		return newMouseEvent(jsMouseEvent.New(typ, ini[0].toDict()))
	}
}

func NewWheelEvent(typ string, ini ...WheelEventInit) WheelEvent {
	jsWheelEvent := js.Global().Get("WheelEvent")
	if isNil(jsWheelEvent) {
		return nil
	}

	switch len(ini) {
	case 0:
		return newWheelEvent(jsWheelEvent.New(typ))
	default:
		return newWheelEvent(jsWheelEvent.New(typ, ini[0].toDict()))
	}
}

func NewInputEvent(typ string, ini ...InputEventInit) InputEvent {
	jsInputEvent := js.Global().Get("InputEvent")
	if isNil(jsInputEvent) {
		return nil
	}

	switch len(ini) {
	case 0:
		return newInputEvent(jsInputEvent.New(typ))
	default:
		return newInputEvent(jsInputEvent.New(typ, ini[0].toDict()))
	}
}

func NewKeyboardEvent(typ string, ini ...KeyboardEventInit) KeyboardEvent {
	jsKeyboardEvent := js.Global().Get("KeyboardEvent")
	if isNil(jsKeyboardEvent) {
		return nil
	}

	switch len(ini) {
	case 0:
		return newKeyboardEvent(jsKeyboardEvent.New(typ))
	default:
		return newKeyboardEvent(jsKeyboardEvent.New(typ, ini[0].toDict()))
	}
}

func NewErrorEvent(typ string, eei ...ErrorEventInit) ErrorEvent {
	jsErrorEvent := js.Global().Get("ErrorEvent")
	if isNil(jsErrorEvent) {
		return nil
	}

	switch len(eei) {
	case 0:
		return newErrorEvent(jsErrorEvent.New(typ))
	default:
		return newErrorEvent(jsErrorEvent.New(typ, eei[0].toDict()))
	}
}

// -------------8<---------------------------------------

type eventTargetImpl struct {
	js.Value
}

func newEventTarget(v js.Value) EventTarget {
	if p := newEventTargetImpl(v); p != nil {
		return p
	}
	return nil
}

func newEventTargetImpl(v js.Value) *eventTargetImpl {
	if isNil(v) {
		return nil
	}

	return &eventTargetImpl{
		Value: v,
	}
}

func (p *eventTargetImpl) On(event string, fn func(ev Event)) EventHandler {
	eh := &elementEventHandlerImpl{
		Value: p.Value,
		fn:    fn,
		typ:   event,
	}

	eh.jsCb = js.NewCallback(eh.jsFunc)
	p.Call("addEventListener", event, eh.jsCb)

	return eh
}

// -------------8<---------------------------------------

type eventImpl struct {
	js.Value
}

func newEvent(v js.Value) Event {
	if p := newEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newEventImpl(v js.Value) *eventImpl {
	if isNil(v) {
		return nil
	}

	return &eventImpl{
		Value: v,
	}
}

func (p *eventImpl) Type() string {
	return p.Get("type").String()
}

func (p *eventImpl) Target() EventTarget {
	return newEventTargetImpl(p.Get("target"))
}

func (p *eventImpl) CurrentTarget() EventTarget {
	return newEventTargetImpl(p.Get("currentTarget"))
}

func (p *eventImpl) ComposedPath() []EventTarget {
	s := arrayToSlice(p.Call("composedPath"))
	if s == nil {
		return nil
	}

	ret := make([]EventTarget, len(s))
	for i, v := range s {
		ret[i] = newEventTargetImpl(v)
	}
	return ret
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

func newCustomEvent(v js.Value) CustomEvent {
	if isNil(v) {
		return nil
	}

	return &customEventImpl{
		eventImpl: newEventImpl(v),
	}
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
	*touchEventHandlersImpl
}

func newGlobalEventHandlersImpl(v js.Value) *globalEventHandlersImpl {
	if isNil(v) {
		return nil
	}

	return &globalEventHandlersImpl{
		touchEventHandlersImpl: newTouchEventHandlersImpl(v),
	}
}

func (p *globalEventHandlersImpl) OnAbort(fn func(Event)) EventHandler {
	return On("abort", fn)
}

func (p *globalEventHandlersImpl) OnAuxClick(fn func(Event)) EventHandler {
	return On("auxclick", fn)
}

func (p *globalEventHandlersImpl) OnBlur(fn func(Event)) EventHandler {
	return On("blur", fn)
}

func (p *globalEventHandlersImpl) OnCancel(fn func(Event)) EventHandler {
	return On("cancel", fn)
}

func (p *globalEventHandlersImpl) OnCanPlay(fn func(Event)) EventHandler {
	return On("canplay", fn)
}

func (p *globalEventHandlersImpl) OnCanPlayThrough(fn func(Event)) EventHandler {
	return On("canplaythrough", fn)
}

func (p *globalEventHandlersImpl) OnChange(fn func(Event)) EventHandler {
	return On("change", fn)
}

func (p *globalEventHandlersImpl) OnClick(fn func(MouseEvent)) EventHandler {
	return On("click", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnClose(fn func(CloseEvent)) EventHandler {
	return On("close", func(e Event) {
		if ce, ok := e.(CloseEvent); ok {
			fn(ce)
		}
	})
}

func (p *globalEventHandlersImpl) OnContextMenu(fn func(MouseEvent)) EventHandler {
	return On("contextmenu", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnCueChange(fn func(Event)) EventHandler {
	return On("cuechange", fn)
}

func (p *globalEventHandlersImpl) OnDblClick(fn func(MouseEvent)) EventHandler {
	return On("dblclick", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnDrag(fn func(DragEvent)) EventHandler {
	return On("drag", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragEnd(fn func(DragEvent)) EventHandler {
	return On("dragend", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragEnter(fn func(DragEvent)) EventHandler {
	return On("dragenter", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragExit(fn func(DragEvent)) EventHandler {
	return On("dragexit", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragLeave(fn func(DragEvent)) EventHandler {
	return On("dragleave", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragOver(fn func(DragEvent)) EventHandler {
	return On("dragover", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDragStart(fn func(DragEvent)) EventHandler {
	return On("dragstart", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDrop(fn func(DragEvent)) EventHandler {
	return On("drop", func(e Event) {
		if de, ok := e.(DragEvent); ok {
			fn(de)
		}
	})
}

func (p *globalEventHandlersImpl) OnDurationChange(fn func(Event)) EventHandler {
	return On("durationchange", fn)
}

func (p *globalEventHandlersImpl) OnEmptied(fn func(Event)) EventHandler {
	return On("emptied", fn)
}

func (p *globalEventHandlersImpl) OnEnded(fn func(Event)) EventHandler {
	return On("ended", fn)
}

func (p *globalEventHandlersImpl) OnError(fn func(ErrorEvent)) EventHandler {
	return On("error", func(e Event) {
		if ee, ok := e.(ErrorEvent); ok {
			fn(ee)
		}
	})
}

func (p *globalEventHandlersImpl) OnFocus(fn func(FocusEvent)) EventHandler {
	return On("focus", func(e Event) {
		if fe, ok := e.(FocusEvent); ok {
			fn(fe)
		}
	})
}

func (p *globalEventHandlersImpl) OnInput(fn func(InputEvent)) EventHandler {
	return On("input", func(e Event) {
		if ie, ok := e.(InputEvent); ok {
			fn(ie)
		}
	})
}

func (p *globalEventHandlersImpl) OnInvalid(fn func(Event)) EventHandler {
	return On("invalid", fn)
}

func (p *globalEventHandlersImpl) OnKeyDown(fn func(KeyboardEvent)) EventHandler {
	return On("keydown", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnKeyPress(fn func(KeyboardEvent)) EventHandler {
	return On("keypress", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnKeyUp(fn func(KeyboardEvent)) EventHandler {
	return On("keyup", func(e Event) {
		if ke, ok := e.(KeyboardEvent); ok {
			fn(ke)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoad(fn func(UIEvent)) EventHandler {
	return On("load", func(e Event) {
		if ue, ok := e.(UIEvent); ok {
			fn(ue)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoadedData(fn func(Event)) EventHandler {
	return On("loadeddata", fn)
}

func (p *globalEventHandlersImpl) OnLoadedMetadata(fn func(Event)) EventHandler {
	return On("loadedmetadata", fn)
}

func (p *globalEventHandlersImpl) OnLoadEnd(fn func(ProgressEvent)) EventHandler {
	return On("loadend", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnLoadStart(fn func(ProgressEvent)) EventHandler {
	return On("loadstart", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseDown(fn func(MouseEvent)) EventHandler {
	return On("mousedown", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseEnter(fn func(MouseEvent)) EventHandler {
	return On("mouseenter", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseLeave(fn func(MouseEvent)) EventHandler {
	return On("mouseleave", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseMove(fn func(MouseEvent)) EventHandler {
	return On("mousemove", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseOut(fn func(MouseEvent)) EventHandler {
	return On("mouseout", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnMouseOver(fn func(MouseEvent)) EventHandler {
	return On("mouseover", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})

}

func (p *globalEventHandlersImpl) OnMouseUp(fn func(MouseEvent)) EventHandler {
	return On("mouseup", func(e Event) {
		if me, ok := e.(MouseEvent); ok {
			fn(me)
		}
	})
}

func (p *globalEventHandlersImpl) OnWheel(fn func(WheelEvent)) EventHandler {
	return On("wheel", func(e Event) {
		if we, ok := e.(WheelEvent); ok {
			fn(we)
		}
	})
}

func (p *globalEventHandlersImpl) OnPause(fn func(Event)) EventHandler {
	return On("pause", fn)
}

func (p *globalEventHandlersImpl) OnPlay(fn func(Event)) EventHandler {
	return On("play", fn)
}

func (p *globalEventHandlersImpl) OnPlaying(fn func(Event)) EventHandler {
	return On("playing", fn)
}

func (p *globalEventHandlersImpl) OnProgress(fn func(ProgressEvent)) EventHandler {
	return On("progress", func(e Event) {
		if pe, ok := e.(ProgressEvent); ok {
			fn(pe)
		}
	})
}

func (p *globalEventHandlersImpl) OnRateChange(fn func(Event)) EventHandler {
	return On("ratechange", fn)
}

func (p *globalEventHandlersImpl) OnReset(fn func(Event)) EventHandler {
	return On("reset", fn)
}

func (p *globalEventHandlersImpl) OnResize(fn func(Event)) EventHandler {
	return On("resize", fn)
}

func (p *globalEventHandlersImpl) OnScroll(fn func(Event)) EventHandler {
	return On("scroll", fn)
}

func (p *globalEventHandlersImpl) OnSecurityPolicyViolation(fn func(Event)) EventHandler {
	return On("securitypolicyviolation", fn)
}

func (p *globalEventHandlersImpl) OnSeeked(fn func(Event)) EventHandler {
	return On("seeked", fn)
}

func (p *globalEventHandlersImpl) OnSeeking(fn func(Event)) EventHandler {
	return On("seeking", fn)
}

func (p *globalEventHandlersImpl) OnSelect(fn func(Event)) EventHandler {
	return On("select", fn)
}

func (p *globalEventHandlersImpl) OnStalled(fn func(Event)) EventHandler {
	return On("stalled", fn)
}

func (p *globalEventHandlersImpl) OnSubmit(fn func(Event)) EventHandler {
	return On("submit", fn)
}

func (p *globalEventHandlersImpl) OnSuspend(fn func(Event)) EventHandler {
	return On("suspend", fn)
}

func (p *globalEventHandlersImpl) OnTimeUpdate(fn func(Event)) EventHandler {
	return On("timeupdate", fn)
}

func (p *globalEventHandlersImpl) OnToggle(fn func(Event)) EventHandler {
	return On("toggle", fn)
}

func (p *globalEventHandlersImpl) OnVolumeChange(fn func(Event)) EventHandler {
	return On("volumechange", fn)
}

func (p *globalEventHandlersImpl) OnWaiting(fn func(Event)) EventHandler {
	return On("waiting", fn)
}

// -------------8<---------------------------------------

var _ DocumentAndElementEventHandlers = &documentAndElementEventHandlersImpl{}

type documentAndElementEventHandlersImpl struct {
	js.Value
}

func newDocumentAndElementEventHandlersImpl(v js.Value) *documentAndElementEventHandlersImpl {
	if isNil(v) {
		return nil
	}
	return &documentAndElementEventHandlersImpl{
		Value: v,
	}
}

func (p *documentAndElementEventHandlersImpl) OnCopy(fn func(Event)) EventHandler {
	return On("copy", fn)
}

func (p *documentAndElementEventHandlersImpl) OnCut(fn func(Event)) EventHandler {
	return On("cut", fn)
}

func (p *documentAndElementEventHandlersImpl) OnPaste(fn func(Event)) EventHandler {
	return On("paste", fn)
}

// -------------8<---------------------------------------

var _ WindowOrWorkerGlobalScope = &windowOrWorkerGlobalScopeImpl{}

type windowOrWorkerGlobalScopeImpl struct {
	js.Value
}

func newWindowOrWorkerGlobalScopeImpl(v js.Value) *windowOrWorkerGlobalScopeImpl {
	if isNil(v) {
		return nil
	}

	return &windowOrWorkerGlobalScopeImpl{
		Value: v,
	}
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

func (p *windowOrWorkerGlobalScopeImpl) IndexedDB() IDBFactory {
	return newIDBFactory(p.Get("indexedDB"))
}

func (p *windowOrWorkerGlobalScopeImpl) Fetch(input RequestInfo, ri ...RequestInit) Promise {
	return newPromiseImpl(p.Call("fetch"))
}

// -------------8<---------------------------------------

var _ WindowEventHandlers = &windowEventHandlersImpl{}

type windowEventHandlersImpl struct {
	js.Value
}

func newWindowEventHandlersImpl(v js.Value) *windowEventHandlersImpl {
	if isNil(v) {
		return nil
	}

	return &windowEventHandlersImpl{
		Value: v,
	}
}

func (p *windowEventHandlersImpl) OnAfterPrint(fn func(Event)) EventHandler {
	return On("afterprint", fn)
}

func (p *windowEventHandlersImpl) OnBeforePrint(fn func(Event)) EventHandler {
	return On("beforeprint", fn)
}

func (p *windowEventHandlersImpl) OnBeforeUnload(fn func(BeforeUnloadEvent)) EventHandler {
	return On("beforeunload", func(e Event) {
		if be, ok := e.(BeforeUnloadEvent); ok {
			fn(be)
		}
	})
}

func (p *windowEventHandlersImpl) OnHashChange(fn func(HashChangeEvent)) EventHandler {
	return On("hashchange", func(e Event) {
		if hce, ok := e.(HashChangeEvent); ok {
			fn(hce)
		}
	})
}

func (p *windowEventHandlersImpl) OnLanguageChange(fn func(Event)) EventHandler {
	return On("languagechange", fn)
}

func (p *windowEventHandlersImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *windowEventHandlersImpl) OnMessageError(fn func(Event)) EventHandler {
	return On("messageerror", fn)
}

func (p *windowEventHandlersImpl) OnOffline(fn func(Event)) EventHandler {
	return On("offline", fn)
}

func (p *windowEventHandlersImpl) OnOnline(fn func(Event)) EventHandler {
	return On("online", fn)
}

func (p *windowEventHandlersImpl) OnPageHide(fn func(PageTransitionEvent)) EventHandler {
	return On("pagehide", func(e Event) {
		if te, ok := e.(PageTransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *windowEventHandlersImpl) OnPageShow(fn func(PageTransitionEvent)) EventHandler {
	return On("pageshow", func(e Event) {
		if te, ok := e.(PageTransitionEvent); ok {
			fn(te)
		}
	})
}

func (p *windowEventHandlersImpl) OnPopState(fn func(PopStateEvent)) EventHandler {
	return On("popstate", func(e Event) {
		if pse, ok := e.(PopStateEvent); ok {
			fn(pse)
		}
	})
}

func (p *windowEventHandlersImpl) OnRejectionHandled(fn func(Event)) EventHandler {
	return On("rejectionhandled", fn)
}

func (p *windowEventHandlersImpl) OnStorage(fn func(StorageEvent)) EventHandler {
	return On("storage", func(e Event) {
		if se, ok := e.(StorageEvent); ok {
			fn(se)
		}
	})
}

func (p *windowEventHandlersImpl) OnUnhandledRejection(fn func(Event)) EventHandler {
	return On("unhandledrejection", fn)
}

func (p *windowEventHandlersImpl) OnUnload(fn func(Event)) EventHandler {
	return On("unload", fn)
}

// -------------8<---------------------------------------

type uiEventImpl struct {
	*eventImpl
}

func newUIEvent(v js.Value) UIEvent {
	if p := newUIEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newUIEventImpl(v js.Value) *uiEventImpl {
	if isNil(v) {
		return nil
	}

	return &uiEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *uiEventImpl) View() Window {
	return newWindowImpl(p.Get("view"))
}

func (p *uiEventImpl) Detail() int {
	return p.Get("detail").Int()
}

// -------------8<---------------------------------------

type mouseEventImpl struct {
	*uiEventImpl
}

func newMouseEvent(v js.Value) MouseEvent {
	if p := newMouseEventImpl(v); p != nil {
		return p
	}
	return nil
}

func newMouseEventImpl(v js.Value) *mouseEventImpl {
	if isNil(v) {
		return nil
	}

	return &mouseEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
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
	return newEventTargetImpl(p.Get("relatedTarget"))
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

func newFocusEvent(v js.Value) FocusEvent {
	if isNil(v) {
		return nil
	}

	return &focusEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
}

func (p *focusEventImpl) RelatedTarget() EventTarget {
	return newEventTargetImpl(p.Get("relatedTarget"))
}

// -------------8<---------------------------------------

type wheelEventImpl struct {
	*mouseEventImpl
}

func newWheelEvent(v js.Value) WheelEvent {
	if isNil(v) {
		return nil
	}

	return &wheelEventImpl{
		mouseEventImpl: newMouseEventImpl(v),
	}
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

func newInputEvent(v js.Value) InputEvent {
	if isNil(v) {
		return nil
	}

	return &inputEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
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

func newKeyboardEvent(v js.Value) KeyboardEvent {
	if isNil(v) {
		return nil
	}

	return &keyboardEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
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

func newCompositionEvent(v js.Value) CompositionEvent {
	if isNil(v) {
		return nil
	}

	return &compositionEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
}

func (p *compositionEventImpl) Data() string {
	return p.Get("data").String()
}

// -------------8<---------------------------------------

type errorEventImpl struct {
	*eventImpl
}

func newErrorEvent(v js.Value) ErrorEvent {
	if isNil(v) {
		return nil
	}

	return &errorEventImpl{
		eventImpl: newEventImpl(v),
	}
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
