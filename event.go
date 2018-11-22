// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

// https://www.w3.org/TR/uievents/#idl-index

type (
	EventHandler interface {
		Type() string
		Handle(Event)
		Release()
		Remove()
		Dispatch() bool
	}

	EventTarget interface {
		js.Wrapper

		On(string, func(ev Event)) EventHandler
	}

	// https://dom.spec.whatwg.org/#event
	Event interface {
		js.Wrapper

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
	}

	// https://dom.spec.whatwg.org/#customevent
	CustomEvent interface {
		Event

		Detail() interface{}
		InitCustomEvent(string, ...interface{})
	}

	// https://dom.spec.whatwg.org/#dictdef-customeventinit
	CustomEventInit struct {
		EventInit

		Detail interface{} `json:"detail"`
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
		touchEventHandlers

		OnAbort(func(Event)) EventHandler
		OnAuxClick(func(Event)) EventHandler
		OnBlur(func(Event)) EventHandler
		OnCancel(func(Event)) EventHandler
		OnCanPlay(func(Event)) EventHandler
		OnCanPlayThrough(func(Event)) EventHandler
		OnChange(func(Event)) EventHandler
		OnClick(func(Event)) EventHandler
		OnClose(func(Event)) EventHandler
		OnContextMenu(func(Event)) EventHandler
		OnCueChange(func(Event)) EventHandler
		OnDblClick(func(Event)) EventHandler
		OnDrag(func(Event)) EventHandler
		OnDragEnd(func(Event)) EventHandler
		OnDragEnter(func(Event)) EventHandler
		OnDragExit(func(Event)) EventHandler
		OnDragLeave(func(Event)) EventHandler
		OnDragOver(func(Event)) EventHandler
		OnDragStart(func(Event)) EventHandler
		OnDrop(func(Event)) EventHandler
		OnDurationChange(func(Event)) EventHandler
		OnEmptied(func(Event)) EventHandler
		OnEnded(func(Event)) EventHandler

		// TODO
		//OnError(func(Event)) OnErrorEventHandler
		OnError(func(Event)) EventHandler
		OnFocus(func(Event)) EventHandler
		OnInput(func(Event)) EventHandler
		OnInvalid(func(Event)) EventHandler
		OnKeyDown(func(Event)) EventHandler
		OnKeyPress(func(Event)) EventHandler
		OnKeyUp(func(Event)) EventHandler
		OnLoad(func(Event)) EventHandler
		OnLoadedData(func(Event)) EventHandler
		OnLoadedMetadata(func(Event)) EventHandler
		OnLoadEnd(func(Event)) EventHandler
		OnLoadStart(func(Event)) EventHandler
		OnMouseDown(func(Event)) EventHandler
		OnMouseEnter(func(Event)) EventHandler
		OnMouseLeave(func(Event)) EventHandler
		OnMouseMove(func(Event)) EventHandler
		OnMouseOut(func(Event)) EventHandler
		OnMouseOver(func(Event)) EventHandler
		OnMouseUp(func(Event)) EventHandler
		OnWheel(func(Event)) EventHandler
		OnPause(func(Event)) EventHandler
		OnPlay(func(Event)) EventHandler
		OnPlaying(func(Event)) EventHandler
		OnProgress(func(Event)) EventHandler
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
	}

	// https://www.w3.org/TR/html52/webappapis.html#windoweventhandlers
	WindowEventHandlers interface {
		OnAfterPrint(func(Event)) EventHandler
		OnBeforePrint(func(Event)) EventHandler

		// TODO
		//OnBeforeUnload(func(Event)) OnBeforeUnloadEventHandler
		OnBeforeUnload(func(Event)) EventHandler
		OnHashChange(func(Event)) EventHandler
		OnLanguageChange(func(Event)) EventHandler
		OnMessage(func(Event)) EventHandler
		OnMessageError(func(Event)) EventHandler
		OnOffline(func(Event)) EventHandler
		OnOnline(func(Event)) EventHandler
		OnPageHide(func(Event)) EventHandler
		OnPageShow(func(Event)) EventHandler
		OnPopState(func(Event)) EventHandler
		OnRejectionHandled(func(Event)) EventHandler
		OnStorage(func(Event)) EventHandler
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

		// https://www.w3.org/TR/IndexedDB/#idbfactory
		IndexedDB() IDBFactory

		/* TODO: Promise ???
		Promise<ImageBitmap> createImageBitmap(ImageBitmapSource image);
		Promise<ImageBitmap> createImageBitmap(ImageBitmapSource image, long sx, long sy, long sw, long sh);
		*/

		Fetch(RequestInfo, ...RequestInit) Promise
	}

	// https://www.w3.org/TR/uievents/#uievent-uievent
	UIEvent interface {
		Event

		View() Window
		Detail() int
	}

	// https://www.w3.org/TR/uievents/#dictdef-uieventinit-uieventinit
	UIEventInit struct {
		EventInit

		View   Window `json:"view"`
		Detail int    `json:"detail"`
	}

	// https://www.w3.org/TR/uievents/#focusevent
	FocusEvent interface {
		UIEvent

		RelatedTarget() EventTarget
	}

	// https://www.w3.org/TR/uievents/#dictdef-focuseventinit
	FocusEventInit struct {
		UIEventInit

		RelatedTarget EventTarget `json:"relatedTarget"`
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

	// https://www.w3.org/TR/uievents/#mouseevent
	MouseEventInit struct {
		EventModifierInit

		ScreenX       float64     `json:"screenX"`
		ScreenY       float64     `json:"screenY"`
		ClientX       float64     `json:"clientX"`
		ClientY       float64     `json:"clientY"`
		Button        int         `json:"button"`
		Buttons       int         `json:"buttons"`
		RelatedTarget EventTarget `json:"relatedTarget"`
	}

	// https://www.w3.org/TR/uievents/#dictdef-eventmodifierinit
	EventModifierInit struct {
		UIEventInit

		CtrlKey            bool `json:"ctrlKey"`
		ShiftKey           bool `json:"shiftKey"`
		AltKey             bool `json:"altKey"`
		MetaKey            bool `json:"metaKey"`
		ModifierAltGraph   bool `json:"modifierAltGraph"`
		ModifierCapsLock   bool `json:"modifierCapsLock"`
		ModifierFn         bool `json:"modifierFn"`
		ModifierFnLock     bool `json:"modifierFnLock"`
		ModifierHyper      bool `json:"modifierHyper"`
		ModifierNumLock    bool `json:"modifierNumLock"`
		ModifierScrollLock bool `json:"modifierScrollLock"`
		ModifierSuper      bool `json:"modifierSuper"`
		ModifierSymbol     bool `json:"modifierSymbol"`
		ModifierSymbolLock bool `json:"modifierSymbolLock"`
	}

	// https://www.w3.org/TR/uievents/#wheelevent
	WheelEvent interface {
		MouseEvent

		DeltaX() float64
		DeltaY() float64
		DeltaZ() float64
		DeltaMode() WheelEventDeltaMode
	}

	// https://www.w3.org/TR/uievents/#dictdef-wheeleventinit
	WheelEventInit struct {
		MouseEventInit

		DeltaX    float64             `json:"deltaX"`
		DeltaY    float64             `json:"deltaY"`
		DeltaZ    float64             `json:"deltaZ"`
		DeltaMode WheelEventDeltaMode `json:"deltaMode"`
	}

	// https://www.w3.org/TR/uievents/#inputevent
	InputEvent interface {
		UIEvent

		Data() string
		IsComposing() bool
	}

	// https://www.w3.org/TR/uievents/#dictdef-inputeventinit
	InputEventInit struct {
		UIEventInit

		Data        string `json:"data"`
		IsComposing bool   `json:"isComposing"`
	}

	// https://www.w3.org/TR/uievents/#keyboardevent-keyboardevent
	KeyboardEvent interface {
		UIEvent

		Key() string
		Code() string
		Location() int

		CtrlKey() bool
		ShiftKey() bool
		AltKey() bool
		MetaKey() bool
		Repeat() bool
		IsComposing() bool
		ModifierState(string) bool
	}

	// https://www.w3.org/TR/uievents/#dictdef-keyboardeventinit
	KeyboardEventInit struct {
		EventModifierInit

		Key         string `json:"key"`
		Code        string `json:"code"`
		Location    int    `json:"location"`
		Repeat      bool   `json:"repeat"`
		IsComposing bool   `json:"isComposing"`
	}

	// https://www.w3.org/TR/uievents/#compositionevent
	CompositionEvent interface {
		UIEvent

		Data() string
	}

	// https://www.w3.org/TR/uievents/#dictdef-compositioneventinit
	CompositionEventInit struct {
		UIEventInit

		Data string `json:"data"`
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

	// https://www.w3.org/TR/html52/webappapis.html#dictdef-erroreventinit
	ErrorEventInit struct {
		EventInit

		Message  string `json:"message"`
		Filename string `json:"filename"`
		Lineno   int    `json:"lineno"`
		Colno    int    `json:"colno"`
		Error    string `json:"error"` // any
	}
)

// https://dom.spec.whatwg.org/#dictdef-eventinit
type EventInit struct {
	Bubbles    bool `json:"bubbles"`
	Cancelable bool `json:"cancelable"`
	Composed   bool `json:"composed"`
}

func (p EventInit) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["bubbles"] = p.Bubbles
	m["cancelable"] = p.Cancelable
	m["composed"] = p.Composed
	return m
}

type WheelEventDeltaMode int

const (
	WheelEventDeltaModePixel WheelEventDeltaMode = 0x00
	WheelEventDeltaModeLine  WheelEventDeltaMode = 0x01
	WheelEventDeltaModePage  WheelEventDeltaMode = 0x02
)

type KeyboardEvenLocation int

const (
	KeyboardEvenLocationStandard KeyboardEvenLocation = 0x00
	KeyboardEvenLocationLeft     KeyboardEvenLocation = 0x01
	KeyboardEvenLocationRight    KeyboardEvenLocation = 0x02
	KeyboardEvenLocationNumpad   KeyboardEvenLocation = 0x03
)

type EventPhase int

const (
	EventPhaseNone      EventPhase = 0
	EventPhaseCapturing EventPhase = 1
	EventPhaseAtTarget  EventPhase = 2
	EventPhaseBubbling  EventPhase = 3
)
