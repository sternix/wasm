// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://www.w3.org/TR/cssom-view-1/#idl-index

type (
	// https://drafts.csswg.org/cssom-view/#screen
	Screen interface {
		AvailWidth() int
		AvailHeight() int
		Width() int
		Height() int
		ColorDepth() int
		PixelDepth() int
	}

	// https://www.w3.org/TR/cssom-view-1/#mediaquerylist
	MediaQueryList interface {
		EventTarget

		Media() string
		Matches() bool
		OnChange(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/cssom-view-1/#mediaquerylistevent
	MediaQueryListEvent interface {
		Event

		Media() string
		Matches() bool
	}

	// https://www.w3.org/TR/cssom-view-1/#caretposition
	CaretPosition interface {
		OffsetNode() Node
		Offset() int
		ClientRect() DOMRect
	}
)

// https://www.w3.org/TR/cssom-view-1/#enumdef-scrolllogicalposition
type ScrollLogicalPosition string

const (
	ScrollLogicalPositionStart   ScrollLogicalPosition = "start"
	ScrollLogicalPositionCenter  ScrollLogicalPosition = "center"
	ScrollLogicalPositionEnd     ScrollLogicalPosition = "end"
	ScrollLogicalPositionNearest ScrollLogicalPosition = "nearest"
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/cssom-view-1/#dictdef-scrollintoviewoptions
type ScrollIntoViewOptions struct {
	ScrollOptions

	Block  ScrollLogicalPosition // default "center"
	Inline ScrollLogicalPosition // default "center"
}

func (p ScrollIntoViewOptions) toDict() js.Value {
	o := p.ScrollOptions.toDict()
	o.Set("block", string(p.Block))
	o.Set("inline", string(p.Inline))
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/cssom-view-1/#dictdef-mediaquerylisteventinit
type MediaQueryListEventInit struct {
	EventInit

	Media   string
	Matches bool
}

func (p MediaQueryListEventInit) toDict() js.Value {
	o := p.EventInit.toDict()
	o.Set("media", p.Media)
	o.Set("matches", p.Matches)
	return o
}
