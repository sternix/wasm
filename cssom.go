// +build js,wasm

package wasm

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

	Block  ScrollLogicalPosition `json:"block"`  // default "center"
	Inline ScrollLogicalPosition `json:"inline"` // default "center"
}

func (p ScrollIntoViewOptions) toMap() map[string]interface{} {
	m := p.ScrollOptions.toMap()
	m["block"] = string(p.Block)
	m["inline"] = string(p.Inline)
	return m
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/cssom-view-1/#dictdef-mediaquerylisteventinit
type MediaQueryListEventInit struct {
	EventInit

	Media   string `json:"media"`
	Matches bool   `json:"matches"`
}

func (p MediaQueryListEventInit) toMap() map[string]interface{} {
	m := p.EventInit.toMap()
	m["media"] = p.Media
	m["matches"] = p.Matches
	return m
}
