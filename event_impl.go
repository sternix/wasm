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
		return newEvent(jsEvent.New(typ, toJSONObject(ei[0])))
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
		return newCustomEvent(jsCustomEvent.New(typ, toJSONObject(cei[0])))
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
		return newFocusEvent(jsFocusEvent.New(typ, toJSONObject(ini[0])))
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
		return newMouseEvent(jsMouseEvent.New(typ, toJSONObject(ini[0])))
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
		return newWheelEvent(jsWheelEvent.New(typ, toJSONObject(ini[0])))
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
		return newInputEvent(jsInputEvent.New(typ, toJSONObject(ini[0])))
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
		return newKeyboardEvent(jsKeyboardEvent.New(typ, toJSONObject(ini[0])))
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
		return newErrorEvent(jsErrorEvent.New(typ, toJSONObject(eei[0])))
	}
}

// -------------8<---------------------------------------

type eventTargetImpl struct {
	*objectImpl
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
		objectImpl: newObjectImpl(v),
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
	*objectImpl
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
		objectImpl: newObjectImpl(v),
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

func (p *globalEventHandlersImpl) OnClick(fn func(Event)) EventHandler {
	return On("click", fn)
}

func (p *globalEventHandlersImpl) OnClose(fn func(Event)) EventHandler {
	return On("close", fn)
}

func (p *globalEventHandlersImpl) OnContextMenu(fn func(Event)) EventHandler {
	return On("contextmenu", fn)
}

func (p *globalEventHandlersImpl) OnCueChange(fn func(Event)) EventHandler {
	return On("cuechange", fn)
}

func (p *globalEventHandlersImpl) OnDblClick(fn func(Event)) EventHandler {
	return On("dblclick", fn)
}

func (p *globalEventHandlersImpl) OnDrag(fn func(Event)) EventHandler {
	return On("drag", fn)
}

func (p *globalEventHandlersImpl) OnDragEnd(fn func(Event)) EventHandler {
	return On("dragend", fn)
}

func (p *globalEventHandlersImpl) OnDragEnter(fn func(Event)) EventHandler {
	return On("dragenter", fn)
}

func (p *globalEventHandlersImpl) OnDragExit(fn func(Event)) EventHandler {
	return On("dragexit", fn)
}

func (p *globalEventHandlersImpl) OnDragLeave(fn func(Event)) EventHandler {
	return On("dragleave", fn)
}

func (p *globalEventHandlersImpl) OnDragOver(fn func(Event)) EventHandler {
	return On("dragover", fn)
}

func (p *globalEventHandlersImpl) OnDragStart(fn func(Event)) EventHandler {
	return On("dragstart", fn)
}

func (p *globalEventHandlersImpl) OnDrop(fn func(Event)) EventHandler {
	return On("drop", fn)
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

func (p *globalEventHandlersImpl) OnError(fn func(Event)) EventHandler {
	return On("error", fn)
}

func (p *globalEventHandlersImpl) OnFocus(fn func(Event)) EventHandler {
	return On("focus", fn)
}

func (p *globalEventHandlersImpl) OnInput(fn func(Event)) EventHandler {
	return On("input", fn)
}

func (p *globalEventHandlersImpl) OnInvalid(fn func(Event)) EventHandler {
	return On("invalid", fn)
}

func (p *globalEventHandlersImpl) OnKeyDown(fn func(Event)) EventHandler {
	return On("keydown", fn)
}

func (p *globalEventHandlersImpl) OnKeyPress(fn func(Event)) EventHandler {
	return On("keypress", fn)
}

func (p *globalEventHandlersImpl) OnKeyUp(fn func(Event)) EventHandler {
	return On("keyup", fn)
}

func (p *globalEventHandlersImpl) OnLoad(fn func(Event)) EventHandler {
	return On("load", fn)
}

func (p *globalEventHandlersImpl) OnLoadedData(fn func(Event)) EventHandler {
	return On("loadeddata", fn)
}

func (p *globalEventHandlersImpl) OnLoadedMetadata(fn func(Event)) EventHandler {
	return On("loadedmetadata", fn)
}

func (p *globalEventHandlersImpl) OnLoadEnd(fn func(Event)) EventHandler {
	return On("loadend", fn)
}

func (p *globalEventHandlersImpl) OnLoadStart(fn func(Event)) EventHandler {
	return On("loadstart", fn)
}

func (p *globalEventHandlersImpl) OnMouseDown(fn func(Event)) EventHandler {
	return On("mousedown", fn)
}

func (p *globalEventHandlersImpl) OnMouseEnter(fn func(Event)) EventHandler {
	return On("mouseenter", fn)
}

func (p *globalEventHandlersImpl) OnMouseLeave(fn func(Event)) EventHandler {
	return On("mouseleave", fn)
}

func (p *globalEventHandlersImpl) OnMouseMove(fn func(Event)) EventHandler {
	return On("mousemove", fn)
}

func (p *globalEventHandlersImpl) OnMouseOut(fn func(Event)) EventHandler {
	return On("mouseout", fn)
}

func (p *globalEventHandlersImpl) OnMouseOver(fn func(Event)) EventHandler {
	return On("mouseover", fn)
}

func (p *globalEventHandlersImpl) OnMouseUp(fn func(Event)) EventHandler {
	return On("mouseup", fn)
}

func (p *globalEventHandlersImpl) OnWheel(fn func(Event)) EventHandler {
	return On("wheel", fn)
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

func (p *globalEventHandlersImpl) OnProgress(fn func(Event)) EventHandler {
	return On("progress", fn)
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

func (p *windowEventHandlersImpl) OnBeforeUnload(fn func(Event)) EventHandler {
	return On("beforeunload", fn)
}

func (p *windowEventHandlersImpl) OnHashChange(fn func(Event)) EventHandler {
	return On("hashchange", fn)
}

func (p *windowEventHandlersImpl) OnLanguageChange(fn func(Event)) EventHandler {
	return On("languagechange", fn)
}

func (p *windowEventHandlersImpl) OnMessage(fn func(Event)) EventHandler {
	return On("message", fn)
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

func (p *windowEventHandlersImpl) OnPageHide(fn func(Event)) EventHandler {
	return On("pagehide", fn)
}

func (p *windowEventHandlersImpl) OnPageShow(fn func(Event)) EventHandler {
	return On("pageshow", fn)
}

func (p *windowEventHandlersImpl) OnPopState(fn func(Event)) EventHandler {
	return On("popstate", fn)
}

func (p *windowEventHandlersImpl) OnRejectionHandled(fn func(Event)) EventHandler {
	return On("rejectionhandled", fn)
}

func (p *windowEventHandlersImpl) OnStorage(fn func(Event)) EventHandler {
	return On("storage", fn)
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
