// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/html52/browsers.html#the-window-object
	Window interface {
		EventTarget
		WindowOrWorkerGlobalScope
		GlobalEventHandlers
		WindowEventHandlers

		Console() Console

		Window() WindowProxy
		Self() WindowProxy
		Document() Document
		Name() string
		SetName(string)
		Location() Location
		History() History
		Locationbar() BarProp
		Menubar() BarProp
		Personalbar() BarProp
		Scrollbars() BarProp
		Statusbar() BarProp
		Toolbar() BarProp
		Status() string
		SetStatus(string)
		Close()
		Closed() bool
		Stop()
		Focus()
		Blur()
		Frames() WindowProxy
		Length() int
		Top() WindowProxy
		Opener() WindowProxy
		Parent() WindowProxy
		FrameElement() Element
		Open(...interface{}) WindowProxy
		Navigator() Navigator
		Alert(...string)
		Confirm(...string) bool
		Prompt(...string) string // message,default
		Print()
		RequestAnimationFrame(FrameRequestCallback) int
		CancelAnimationFrame(int)

		// https://www.w3.org/TR/cssom-view-1/#extensions-to-the-window-interface
		// partials
		MatchMedia(string) MediaQueryList
		Screen() Screen
		MoveTo(int, int)
		MoveBy(int, int)
		ResizeTo(int, int)
		ResizeBy(int, int)
		InnerWidth() int
		InnerHeight() int
		ScrollX() float64
		PageXOffset() float64
		ScrollY() float64
		PageYOffset() float64
		Scroll(...interface{})
		ScrollTo(...interface{})
		ScrollBy(...interface{})
		ScreenX() int
		ScreenY() int
		OuterWidth() int
		OuterHeight() int
		DevicePixelRatio() float64
	}

	// https://www.w3.org/TR/html52/browsers.html#barprop
	BarProp interface {
		Visible() bool
	}

	// https://www.w3.org/TR/html52/browsers.html#location
	Location interface {
		WorkerLocation // getters are same

		SetHref(string)
		SetProtocol(string)
		SetHost(string)
		SetHostname(string)
		SetPort(string)
		SetPathname(string)
		SetSearch(string)
		SetHash(string)
		Assign(string)
		Replace(string)
		Reload()
		AncestorOrigins() []string
	}

	// https://www.w3.org/TR/html52/browsers.html#windowproxy
	WindowProxy interface {
		Window
	}

	// https://www.w3.org/TR/html52/browsers.html#history
	History interface {
		Length() int
		ScrollRestoration() ScrollRestorationType
		SetScrollRestoration(ScrollRestorationType)
		State() interface{}
		Go(...int)
		Back()
		Forward()
		PushState(interface{}, string, ...string)
		ReplaceState(interface{}, string, ...string)
	}

	// https://www.w3.org/TR/html52/browsers.html#popstateevent
	PopStateEvent interface {
		Event

		State() interface{}
		SetState(interface{})
	}

	// https://www.w3.org/TR/html52/browsers.html#dictdef-popstateeventinit
	PopStateEventInit struct {
		EventInit

		State interface{} `json:"state"`
	}

	// https://www.w3.org/TR/html52/browsers.html#hashchangeevent
	HashChangeEvent interface {
		Event

		OldURL() string
		NewURL() string
	}

	// https://www.w3.org/TR/html52/browsers.html#dictdef-hashchangeeventinit
	HashChangeEventInit struct {
		EventInit

		OldUrl string `json:"oldURL"`
		NewURL string `json:"newURL"`
	}

	// https://www.w3.org/TR/html52/browsers.html#pagetransitionevent
	PageTransitionEvent interface {
		Event

		Persisted() bool
	}

	// https://www.w3.org/TR/html52/browsers.html#dictdef-pagetransitioneventinit
	PageTransitionEventInit struct {
		EventInit

		Persisted bool `json:"persisted"`
	}

	// https://www.w3.org/TR/html52/browsers.html#beforeunloadevent
	BeforeUnloadEvent interface {
		Event

		ReturnValue() string
		SetReturnValue(string)
	}

	// https://www.w3.org/TR/html52/browsers.html#navigatoronline
	NavigatorOnLine interface {
		OnLine() bool
	}
)

type ScrollRestorationType string

const (
	ScrollRestorationTypeAuto   ScrollRestorationType = "auto"
	ScrollRestorationTypeManual ScrollRestorationType = "manual"
)

type DocumentReadyState string

const (
	DocumentReadyStateLoading     DocumentReadyState = "loading"
	DocumentReadyStateInteractive DocumentReadyState = "interactive"
	DocumentReadyStateComplete    DocumentReadyState = "complete"
)
