// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://html.spec.whatwg.org/multipage/scripting.html#htmlslotelement
	HTMLSlotElement interface {
		HTMLElement

		Name() string
		SetName(string)
		AssignedNodes(...AssignedNodesOptions) []Node
		AssignedElements(...AssignedNodesOptions) []Element
	}

	// https://html.spec.whatwg.org/multipage/dom.html#htmlorsvgelement
	HTMLOrSVGElement interface {
		js.Wrapper

		DataSet() map[string]string // DomStringMap = map[string]string
		Nonce() string
		SetNonce(string)
		TabIndex() int
		SetTabIndex(int)
		Focus(...FocusOptions)
		Blur()
	}

	// https://html.spec.whatwg.org/multipage/interaction.html#elementcontenteditable
	ElementContentEditable interface {
		js.Wrapper

		ContentEditable() string
		SetContentEditable(string)
		IsContentEditable() bool
		InputMode() string
		SetInputMode(string)
	}

	// https://dom.spec.whatwg.org/#abortsignal
	AbortSignal interface {
		EventTarget

		Aborted() bool
		OnAbort(func(Event)) EventHandler
	}

	// https://dom.spec.whatwg.org/#abortcontroller
	AbortController interface {
		js.Wrapper

		Signal() AbortSignal
		Abort()
	}
)

// -------------8<---------------------------------------

// https://html.spec.whatwg.org/multipage/scripting.html#assignednodesoptions
type AssignedNodesOptions struct {
	Flatten bool `json:"flatten"` // default false
}

func (p AssignedNodesOptions) toDict() js.Value {
	if p.Flatten {
		o := jsObject.New()
		o.Set("flatten", p.Flatten)
		return o
	}
	return js.Null()
}

// -------------8<---------------------------------------

// https://html.spec.whatwg.org/multipage/interaction.html#focusoptions
type FocusOptions struct {
	PreventScroll bool `json:"preventScroll"`
}

func (p FocusOptions) toDict() js.Value {
	o := jsObject.New()
	o.Set("preventScroll", p.PreventScroll)
	return o
}
