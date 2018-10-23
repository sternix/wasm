// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

var _ partialElementFullscreen = &partialElementFullscreenImpl{}

type partialElementFullscreenImpl struct {
	js.Value
}

// KEEP
func newpartialElementFullscreenImpl(v js.Value) *partialElementFullscreenImpl {
	if isNil(v) {
		return nil
	}

	return &partialElementFullscreenImpl{
		Value: v,
	}
}

func (p *partialElementFullscreenImpl) RequestFullscreen(...FullscreenOptions) Promise {
	return newPromiseImpl(p.Call("requestFullscreen"))
}
func (p *partialElementFullscreenImpl) OnFullScreenChange(fn func(Event)) EventHandler {
	return On("fullscreenchange", fn)
}

func (p *partialElementFullscreenImpl) OnFullScreenError(fn func(Event)) EventHandler {
	return On("fullscreenerror", fn)
}

// -------------8<---------------------------------------

var _ partialDocumentFullscreen = &partialDocumentFullscreenImpl{}

type partialDocumentFullscreenImpl struct {
	js.Value
}

// KEEP
func newpartialDocumentFullscreenImpl(v js.Value) *partialDocumentFullscreenImpl {
	if isNil(v) {
		return nil
	}

	return &partialDocumentFullscreenImpl{
		Value: v,
	}
}

func (p *partialDocumentFullscreenImpl) FullscreenEnabled() bool {
	return p.Get("fullscreenEnabled").Bool()
}

func (p *partialDocumentFullscreenImpl) ExitFullscreen() Promise {
	return newPromiseImpl(p.Call("exitFullscreen"))
}

func (p *partialDocumentFullscreenImpl) OnFullscreenChange(fn func(Event)) EventHandler {
	return On("fullscreenchange", fn)
}

func (p *partialDocumentFullscreenImpl) OnFullscreenError(fn func(Event)) EventHandler {
	return On("fullscreenerror", fn)
}

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
