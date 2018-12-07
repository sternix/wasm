// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://w3c.github.io/clipboard-apis/#idl-index

type (
	ClipboardEvent interface {
		Event

		ClipboardData() DataTransfer
	}

	Clipboard interface {
		EventTarget

		Read() func() (DataTransfer, error) //Promise<DataTransfer>
		ReadText() func() (string, bool)    // Promise<DOMString>
		Write(DataTransfer) func() bool     // Promise<void>
		WriteText(string) func() bool       // Promise<void>
	}
)

// -------------8<---------------------------------------

type ClipboardPermissionDescriptor struct {
	PermissionDescriptor

	AllowWithoutGesture bool
}

func (p ClipboardPermissionDescriptor) toDict() js.Value {
	o := p.PermissionDescriptor.toDict()
	o.Set("allowWithoutGesture", p.AllowWithoutGesture)
	return o
}

// -------------8<---------------------------------------

type ClipboardEventInit struct {
	EventInit

	ClipboardData DataTransfer
}

func (p ClipboardEventInit) toDict() js.Value {
	o := p.EventInit.toDict()
	o.Set("clipboardData", p.ClipboardData.JSValue())
	return o
}

// -------------8<---------------------------------------
