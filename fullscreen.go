// +build js,wasm

package wasm

// https://fullscreen.spec.whatwg.org/

type (
	FullscreenOptions struct {
		NavigationUI FullscreenNavigationUI `json:"navigationUI"`
	}

	partialElementFullscreen interface {
		RequestFullscreen(...FullscreenOptions) Promise // Promise<void>
		OnFullScreenChange(func(Event)) EventHandler
		OnFullScreenError(func(Event)) EventHandler
	}

	partialDocumentFullscreen interface {
		FullscreenEnabled() bool
		ExitFullscreen() Promise // Promise<void>
		OnFullscreenChange(func(Event)) EventHandler
		OnFullscreenError(func(Event)) EventHandler
	}

	partialDocumentOrShadowFullscreen interface {
		FullscreenElement() Element
	}
)

type FullscreenNavigationUI string

const (
	FullscreenNavigationUIAuto FullscreenNavigationUI = "auto"
	FullscreenNavigationUIShow FullscreenNavigationUI = "show"
	FullscreenNavigationUIHide FullscreenNavigationUI = "hide"
)
