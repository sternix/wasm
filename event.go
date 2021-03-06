// +build js,wasm

package wasm

import (
	"time"
)

// https://www.w3.org/TR/uievents/#idl-index

type (
	EventHandler interface {
		Type() string
		Handle(Event)
		Release()
		Remove()
	}

	EventTarget interface {
		On(string, func(Event)) EventHandler
		DispatchEvent(Event) bool

		JSValue() jsValue
	}

	// https://dom.spec.whatwg.org/#event
	Event interface {
		Type() string
		Target() EventTarget
		CurrentTarget() EventTarget
		ComposedPath() []EventTarget
		EventPhase() EventPhase
		StopPropagation()
		StopImmediatePropagation()
		Bubbles() bool
		Cancelable() bool
		PreventDefault()
		DefaultPrevented() bool
		Composed() bool
		IsTrusted() bool
		TimeStamp() time.Time

		JSValue() jsValue
	}

	// https://dom.spec.whatwg.org/#customevent
	CustomEvent interface {
		Event

		Detail() interface{}
		InitCustomEvent(string, ...interface{})
	}

	// https://html.spec.whatwg.org/multipage/webappapis.html#onerroreventhandler
	// TODO: can we use ErrorEvent for this
	OnErrorEventHandler func(OnErrorEventHandlerArg) interface{}

	// https://html.spec.whatwg.org/multipage/webappapis.html#onbeforeunloadeventhandler
	OnBeforeUnloadEventHandler func(Event) string

	OnErrorEventHandlerArg struct {
		Message Event
		Source  string
		LineNo  int
		ColNo   int
		Error   interface{}
	}

	// https://www.w3.org/TR/html52/webappapis.html#globaleventhandlers
	GlobalEventHandlers interface {
		OnAbort(func(Event)) EventHandler
		OnAuxClick(func(Event)) EventHandler
		OnBlur(func(Event)) EventHandler
		OnCancel(func(Event)) EventHandler
		OnCanPlay(func(Event)) EventHandler
		OnCanPlayThrough(func(Event)) EventHandler
		OnChange(func(Event)) EventHandler
		OnClick(func(MouseEvent)) EventHandler
		OnClose(func(CloseEvent)) EventHandler
		OnContextMenu(func(MouseEvent)) EventHandler
		OnCueChange(func(Event)) EventHandler
		OnDblClick(func(MouseEvent)) EventHandler
		OnDrag(func(DragEvent)) EventHandler
		OnDragEnd(func(DragEvent)) EventHandler
		OnDragEnter(func(DragEvent)) EventHandler
		OnDragExit(func(DragEvent)) EventHandler
		OnDragLeave(func(DragEvent)) EventHandler
		OnDragOver(func(DragEvent)) EventHandler
		OnDragStart(func(DragEvent)) EventHandler
		OnDrop(func(DragEvent)) EventHandler
		OnDurationChange(func(Event)) EventHandler
		OnEmptied(func(Event)) EventHandler
		OnEnded(func(Event)) EventHandler

		// TODO
		//OnError(func(Event)) OnErrorEventHandler
		OnError(func(ErrorEvent)) EventHandler
		OnFocus(func(FocusEvent)) EventHandler
		OnInput(func(InputEvent)) EventHandler
		OnInvalid(func(Event)) EventHandler
		OnKeyDown(func(KeyboardEvent)) EventHandler
		OnKeyPress(func(KeyboardEvent)) EventHandler
		OnKeyUp(func(KeyboardEvent)) EventHandler
		OnLoad(func(UIEvent)) EventHandler
		OnLoadedData(func(Event)) EventHandler
		OnLoadedMetadata(func(Event)) EventHandler
		OnLoadEnd(func(ProgressEvent)) EventHandler
		OnLoadStart(func(ProgressEvent)) EventHandler
		OnMouseDown(func(MouseEvent)) EventHandler
		OnMouseEnter(func(MouseEvent)) EventHandler
		OnMouseLeave(func(MouseEvent)) EventHandler
		OnMouseMove(func(MouseEvent)) EventHandler
		OnMouseOut(func(MouseEvent)) EventHandler
		OnMouseOver(func(MouseEvent)) EventHandler
		OnMouseUp(func(MouseEvent)) EventHandler
		OnWheel(func(WheelEvent)) EventHandler
		OnPause(func(Event)) EventHandler
		OnPlay(func(Event)) EventHandler
		OnPlaying(func(Event)) EventHandler
		OnProgress(func(ProgressEvent)) EventHandler
		OnRateChange(func(Event)) EventHandler
		OnReset(func(Event)) EventHandler
		OnResize(func(Event)) EventHandler
		OnScroll(func(Event)) EventHandler
		OnSecurityPolicyViolation(func(Event)) EventHandler
		OnSeeked(func(Event)) EventHandler
		OnSeeking(func(Event)) EventHandler
		OnSelect(func(Event)) EventHandler
		OnStalled(func(Event)) EventHandler
		OnSubmit(func(Event)) EventHandler
		OnSuspend(func(Event)) EventHandler
		OnTimeUpdate(func(Event)) EventHandler
		OnToggle(func(Event)) EventHandler
		OnVolumeChange(func(Event)) EventHandler
		OnWaiting(func(Event)) EventHandler

		// https://w3c.github.io/touch-events/#extensions-to-the-globaleventhandlers-interface
		OnTouchStart(func(TouchEvent)) EventHandler
		OnTouchEnd(func(TouchEvent)) EventHandler
		OnTouchMove(func(TouchEvent)) EventHandler
		OnTouchCancel(func(TouchEvent)) EventHandler

		//https://drafts.csswg.org/css-transitions/#interface-globaleventhandlers-idl
		OnTransitionRun(func(TransitionEvent)) EventHandler
		OnTransitionStart(func(TransitionEvent)) EventHandler
		OnTransitionEnd(func(TransitionEvent)) EventHandler
		OnTransitionCancel(func(TransitionEvent)) EventHandler

		// https://www.w3.org/TR/pointerevents/#extensions-to-the-globaleventhandlers-interface
		OnGotPointerCapture(func(PointerEvent)) EventHandler
		OnLostPointerCapture(func(PointerEvent)) EventHandler
		OnPointerDown(func(PointerEvent)) EventHandler
		OnPointerMove(func(PointerEvent)) EventHandler
		OnPointerUp(func(PointerEvent)) EventHandler
		OnPointerCancel(func(PointerEvent)) EventHandler
		OnPointerOver(func(PointerEvent)) EventHandler
		OnPointerOut(func(PointerEvent)) EventHandler
		OnPointerEnter(func(PointerEvent)) EventHandler
		OnPointerLeave(func(PointerEvent)) EventHandler
	}

	// https://www.w3.org/TR/html52/webappapis.html#windoweventhandlers
	WindowEventHandlers interface {
		OnAfterPrint(func(Event)) EventHandler
		OnBeforePrint(func(Event)) EventHandler

		// TODO
		//OnBeforeUnload(func(Event)) OnBeforeUnloadEventHandler
		OnBeforeUnload(func(BeforeUnloadEvent)) EventHandler
		OnHashChange(func(HashChangeEvent)) EventHandler
		OnLanguageChange(func(Event)) EventHandler
		OnMessage(func(MessageEvent)) EventHandler
		OnMessageError(func(Event)) EventHandler
		OnOffline(func(Event)) EventHandler
		OnOnline(func(Event)) EventHandler
		OnPageHide(func(PageTransitionEvent)) EventHandler
		OnPageShow(func(PageTransitionEvent)) EventHandler
		OnPopState(func(PopStateEvent)) EventHandler
		OnRejectionHandled(func(Event)) EventHandler
		OnStorage(func(StorageEvent)) EventHandler
		OnUnhandledRejection(func(Event)) EventHandler
		OnUnload(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/webappapis.html#documentandelementeventhandlers
	DocumentAndElementEventHandlers interface {
		OnCopy(func(Event)) EventHandler
		OnCut(func(Event)) EventHandler
		OnPaste(func(Event)) EventHandler
	}

	// https://html.spec.whatwg.org/multipage/webappapis.html#windoworworkerglobalscope-mixin
	WindowOrWorkerGlobalScope interface {
		Origin() string
		Btoa(string) string
		Atob(string) string
		SetTimeout(TimerCallback, int) int
		ClearTimeout(int)
		SetInterval(TimerCallback, int) int
		ClearInterval(int)
		QueueMicrotask(VoidFunction)
		CreateImageBitmap(ImageBitmapSource, ...ImageBitmapOptions) func() (ImageBitmap, error)
		CreateImageBitmapWithSize(ImageBitmapSource, int, int, int, int, ...ImageBitmapOptions) func() (ImageBitmap, error)

		// https://www.w3.org/TR/IndexedDB/#idbfactory
		IndexedDB() IDBFactory

		// https://fetch.spec.whatwg.org/#fetch-method
		Fetch(RequestInfo, ...RequestInit) func() (Response, error)
	}

	// https://www.w3.org/TR/uievents/#uievent-uievent
	UIEvent interface {
		Event

		View() Window
		Detail() int
	}

	// https://www.w3.org/TR/uievents/#focusevent
	FocusEvent interface {
		UIEvent

		RelatedTarget() EventTarget
	}

	// https://www.w3.org/TR/uievents/#mouseevent
	// https://www.w3.org/TR/cssom-view-1/#extensions-to-the-mouseevent-interface
	MouseEvent interface {
		UIEvent

		ScreenX() float64
		ScreenY() float64
		ClientX() float64
		ClientY() float64
		CtrlKey() bool
		ShiftKey() bool
		AltKey() bool
		MetaKey() bool
		Button() int
		Buttons() int
		RelatedTarget() EventTarget
		ModifierState(string) bool
		PageX() float64
		PageY() float64
		X() float64
		Y() float64
		OffsetX() float64
		OffsetY() float64
	}

	// https://www.w3.org/TR/uievents/#wheelevent
	WheelEvent interface {
		MouseEvent

		DeltaX() float64
		DeltaY() float64
		DeltaZ() float64
		DeltaMode() DeltaModeCode
	}

	// https://www.w3.org/TR/uievents/#inputevent
	InputEvent interface {
		UIEvent

		Data() string
		IsComposing() bool
	}

	// https://www.w3.org/TR/uievents/#keyboardevent
	KeyboardEvent interface {
		UIEvent

		Key() string
		Code() string
		Location() KeyLocationCode
		CtrlKey() bool
		ShiftKey() bool
		AltKey() bool
		MetaKey() bool
		Repeat() bool
		IsComposing() bool
		ModifierState(string) bool
	}

	// https://www.w3.org/TR/uievents/#compositionevent
	CompositionEvent interface {
		UIEvent

		Data() string
	}

	// https://html.spec.whatwg.org/multipage/webappapis.html#errorevent
	ErrorEvent interface {
		Event

		Message() string
		Filename() string
		Lineno() int
		Colno() int
		Error() string // TODO any
	}

	// https://drafts.csswg.org/css-transitions/#transitionevent
	TransitionEvent interface {
		Event

		PropertyName() string
		ElapsedTime() float64
		PseudoElement() string
	}

	// https://www.w3.org/TR/pointerevents/#pointerevent-interface
	PointerEvent interface {
		MouseEvent

		PointerId() int
		Width() float64
		Height() float64
		Pressure() float64
		TangentialPressure() float64
		TiltX() int
		TiltY() int
		Twist() int
		PointerType() string
		IsPrimary() bool
	}
)

type DeltaModeCode uint

const (
	DeltaModeCodePixel DeltaModeCode = 0x00
	DeltaModeCodeLine  DeltaModeCode = 0x01
	DeltaModeCodePage  DeltaModeCode = 0x02
)

type KeyLocationCode uint

const (
	KeyLocationCodeStandard KeyLocationCode = 0x00
	KeyLocationCodeLeft     KeyLocationCode = 0x01
	KeyLocationCodeRight    KeyLocationCode = 0x02
	KeyLocationCodeNumpad   KeyLocationCode = 0x03
)

type EventPhase uint16

const (
	EventPhaseNone      EventPhase = 0
	EventPhaseCapturing EventPhase = 1
	EventPhaseAtTarget  EventPhase = 2
	EventPhaseBubbling  EventPhase = 3
)

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-eventinit
type EventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
}

func (p EventInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("bubbles", p.Bubbles)
	o.Set("cancelable", p.Cancelable)
	o.Set("composed", p.Composed)
	return o
}

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-customeventinit
type CustomEventInit struct {
	EventInit

	Detail interface{}
}

func (p CustomEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("detail", p.Detail)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-uieventinit-uieventinit
type UIEventInit struct {
	EventInit

	View   Window
	Detail int
}

func (p UIEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("view", p.View)
	o.Set("detail", p.Detail)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-focuseventinit
type FocusEventInit struct {
	UIEventInit

	RelatedTarget EventTarget
}

func (p FocusEventInit) JSValue() jsValue {
	o := p.UIEventInit.JSValue()
	o.Set("relatedTarget", p.RelatedTarget)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#mouseevent
type MouseEventInit struct {
	EventModifierInit

	ScreenX       float64
	ScreenY       float64
	ClientX       float64
	ClientY       float64
	Button        int
	Buttons       int
	RelatedTarget EventTarget
}

func (p MouseEventInit) JSValue() jsValue {
	o := p.EventModifierInit.JSValue()
	o.Set("screenX", p.ScreenX)
	o.Set("screenY", p.ScreenY)
	o.Set("clientX", p.ClientX)
	o.Set("clientY", p.ClientY)
	o.Set("button", p.Button)
	o.Set("buttons", p.Buttons)
	o.Set("relatedTarget", JSValueOf(p.RelatedTarget))
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-eventmodifierinit
type EventModifierInit struct {
	UIEventInit

	CtrlKey            bool
	ShiftKey           bool
	AltKey             bool
	MetaKey            bool
	ModifierAltGraph   bool
	ModifierCapsLock   bool
	ModifierFn         bool
	ModifierFnLock     bool
	ModifierHyper      bool
	ModifierNumLock    bool
	ModifierScrollLock bool
	ModifierSuper      bool
	ModifierSymbol     bool
	ModifierSymbolLock bool
}

func (p EventModifierInit) JSValue() jsValue {
	o := p.UIEventInit.JSValue()
	o.Set("ctrlKey", p.CtrlKey)
	o.Set("shiftKey", p.ShiftKey)
	o.Set("altKey", p.AltKey)
	o.Set("metaKey", p.MetaKey)
	o.Set("modifierAltGraph", p.ModifierAltGraph)
	o.Set("modifierCapsLock", p.ModifierCapsLock)
	o.Set("modifierFn", p.ModifierFn)
	o.Set("modifierFnLock", p.ModifierFnLock)
	o.Set("modifierHyper", p.ModifierHyper)
	o.Set("modifierNumLock", p.ModifierNumLock)
	o.Set("modifierScrollLock", p.ModifierScrollLock)
	o.Set("modifierSuper", p.ModifierSuper)
	o.Set("modifierSymbol", p.ModifierSymbol)
	o.Set("modifierSymbolLock", p.ModifierSymbolLock)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-wheeleventinit
type WheelEventInit struct {
	MouseEventInit

	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode DeltaModeCode
}

func (p WheelEventInit) JSValue() jsValue {
	o := p.MouseEventInit.JSValue()
	o.Set("deltaX", p.DeltaX)
	o.Set("deltaY", p.DeltaY)
	o.Set("deltaZ", p.DeltaZ)
	o.Set("deltaMode", uint(p.DeltaMode))
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-inputeventinit
type InputEventInit struct {
	UIEventInit

	Data        string
	IsComposing bool
}

func (p InputEventInit) JSValue() jsValue {
	o := p.UIEventInit.JSValue()
	o.Set("data", p.Data)
	o.Set("isComposing", p.IsComposing)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-keyboardeventinit
type KeyboardEventInit struct {
	EventModifierInit

	Key         string
	Code        string
	Location    KeyLocationCode
	Repeat      bool
	IsComposing bool
}

func (p KeyboardEventInit) JSValue() jsValue {
	o := p.EventModifierInit.JSValue()
	o.Set("key", p.Key)
	o.Set("code", p.Code)
	o.Set("location", uint(p.Location))
	o.Set("repeat", p.Repeat)
	o.Set("isComposing", p.IsComposing)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/uievents/#dictdef-compositioneventinit
type CompositionEventInit struct {
	UIEventInit

	Data string
}

func (p CompositionEventInit) JSValue() jsValue {
	o := p.UIEventInit.JSValue()
	o.Set("data", p.Data)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/html52/webappapis.html#dictdef-erroreventinit
type ErrorEventInit struct {
	EventInit

	Message  string
	Filename string
	Lineno   int
	Colno    int
	Error    string // any
}

func (p ErrorEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("message", p.Message)
	o.Set("filename", p.Filename)
	o.Set("lineno", p.Lineno)
	o.Set("colno", p.Colno)
	o.Set("error", p.Error)
	return o
}

// -------------8<---------------------------------------

// https://drafts.csswg.org/css-transitions/#dictdef-transitioneventinit
type TransitionEventInit struct {
	EventInit

	PropertyName  string
	ElapsedTime   float64
	PseudoElement string
}

func (p TransitionEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("propertyName", p.PropertyName)
	o.Set("elapsedTime", p.ElapsedTime)
	o.Set("pseudoElement", p.PseudoElement)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/pointerevents/#pointerevent-interface
type PointerEventInit struct {
	MouseEventInit

	PointerId          int
	Width              float64
	Height             float64
	Pressure           float64
	TangentialPressure float64
	TiltX              int
	TiltY              int
	Twist              int
	PointerType        string
	IsPrimary          bool
}

func (p PointerEventInit) JSValue() jsValue {
	o := p.MouseEventInit.JSValue()
	o.Set("pointerId", p.PointerId)
	o.Set("width", p.Width)
	o.Set("height", p.Height)
	o.Set("pressure", p.Pressure)
	o.Set("tangentialPressure", p.TangentialPressure)
	o.Set("tiltX", p.TiltX)
	o.Set("tiltY", p.TiltY)
	o.Set("twist", p.Twist)
	o.Set("pointerType", p.PointerType)
	o.Set("isPrimary", p.IsPrimary)
	return o
}
