// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://drafts.csswg.org/cssom-view/

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

	// https://drafts.csswg.org/cssom-view/#mediaquerylist
	MediaQueryList interface {
		EventTarget

		Media() string
		Matches() bool
		OnChange(func(Event)) EventHandler
	}

	// https://drafts.csswg.org/cssom-view/#mediaquerylistevent
	MediaQueryListEvent interface {
		Event

		Media() string
		Matches() bool
	}

	// https://drafts.csswg.org/cssom-view/#caretposition
	CaretPosition interface {
		OffsetNode() Node
		Offset() int
		ClientRect() DOMRect
	}

	// (Element or ProcessingInstruction)
	//StyleSheetOwnerNode Node

	// typedef (Text or Element or CSSPseudoElement or Document) GeometryNode;
	// https://drafts.csswg.org/cssom-view/#typedefdef-geometrynode
	GeometryNode interface {
		js.Wrapper
	}

	// https://drafts.csswg.org/cssom-view/#geometryutils
	GeometryUtils interface {
		BoxQuads(...BoxQuadOptions) []DOMQuad
		ConvertQuadFromNode(DOMQuadInit, GeometryNode, ...ConvertCoordinateOptions) DOMQuad
		ConvertRectFromNode(DOMRectReadOnly, GeometryNode, ...ConvertCoordinateOptions) DOMQuad
		ConvertPointFromNode(DOMPointInit, GeometryNode, ...ConvertCoordinateOptions) DOMPoint
	}

	// https://drafts.csswg.org/css-pseudo-4/#csspseudoelement
	CSSPseudoElement interface {
		EventTarget

		Type() string
		Style() CSSStyleDeclaration
	}

	// https://drafts.csswg.org/css-pseudo-4/#csspseudoelementlist
	CSSPseudoElementList interface {
		Length() int
		Item(int) CSSPseudoElement
		ByType(string) CSSPseudoElement
	}
)

// -------------8<---------------------------------------

// https://drafts.csswg.org/cssom-view/#enumdef-cssboxtype
type CSSBoxType string

const (
	CSSBoxTypeMargin  CSSBoxType = "margin"
	CSSBoxTypeBorder  CSSBoxType = "border"
	CSSBoxTypePadding CSSBoxType = "padding"
	CSSBoxTypeContent CSSBoxType = "content"
)

// https://www.w3.org/TR/cssom-view-1/#enumdef-scrollbehavior
type ScrollBehavior string

const (
	ScrollBehaviorAuto    ScrollBehavior = "auto"
	ScrollBehaviorInstant ScrollBehavior = "instant"
	ScrollBehaviorSmooth  ScrollBehavior = "smooth"
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

// https://www.w3.org/TR/cssom-view-1/#dictdef-scrolloptions
type ScrollOptions struct {
	Behavior ScrollBehavior // default auto
}

func (p ScrollOptions) toDict() js.Value {
	o := jsObject.New()
	if p.Behavior != "" && p.Behavior != ScrollBehaviorAuto {
		o.Set("behavior", string(p.Behavior))
	}
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/cssom-view-1/#dictdef-scrolltooptions
type ScrollToOptions struct {
	ScrollOptions

	Left float64
	Top  float64
}

func (p ScrollToOptions) toDict() js.Value {
	o := p.ScrollOptions.toDict()
	o.Set("left", p.Left)
	o.Set("top", p.Top)
	return o
}

// -------------8<---------------------------------------

// https://drafts.csswg.org/cssom-view/#dictdef-scrollintoviewoptions
type ScrollIntoViewOptions struct {
	ScrollOptions

	Block  ScrollLogicalPosition // default "start"
	Inline ScrollLogicalPosition // default "nearest"
}

func (p ScrollIntoViewOptions) toDict() js.Value {
	o := p.ScrollOptions.toDict()
	if p.Block != "" && p.Block != ScrollLogicalPositionStart {
		o.Set("block", string(p.Block))
	}
	if p.Inline != "" && p.Inline != ScrollLogicalPositionNearest {
		o.Set("inline", string(p.Inline))
	}
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

// -------------8<---------------------------------------

// https://drafts.csswg.org/cssom-view/#dictdef-boxquadoptions
type BoxQuadOptions struct {
	Box        CSSBoxType // default "border"
	RelativeTo GeometryNode
}

func (p BoxQuadOptions) toDict() js.Value {
	o := jsObject.New()
	o.Set("box", string(p.Box))
	o.Set("relativeTo", p.RelativeTo.JSValue())
	return o
}

// -------------8<---------------------------------------

// https://drafts.csswg.org/cssom-view/#dictdef-convertcoordinateoptions
type ConvertCoordinateOptions struct {
	FromBox CSSBoxType // default "border"
	ToBox   CSSBoxType // default "border"
}

func (p ConvertCoordinateOptions) toDict() js.Value {
	o := jsObject.New()
	if p.FromBox != "" && p.FromBox != CSSBoxTypeBorder {
		o.Set("fromBox", string(p.FromBox))
	}

	if p.ToBox != "" && p.ToBox != CSSBoxTypeBorder {
		o.Set("toBox", string(p.ToBox))
	}
	return o
}
