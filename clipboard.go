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

		Read() Promise              //Promise<DataTransfer>
		ReadText() Promise          // Promise<DOMString>
		Write(DataTransfer) Promise // Promise<void>
		WriteText(string) Promise   // Promise<void>
	}
)

// -------------8<---------------------------------------

type ClipboardPermissionDescriptor struct {
	PermissionDescriptor

	AllowWithoutGesture bool `json:"allowWithoutGesture"`
}

func (p ClipboardPermissionDescriptor) toDict() js.Value {
	o := p.PermissionDescriptor.toDict()
	o.Set("allowWithoutGesture", p.AllowWithoutGesture)
	return o
}

// -------------8<---------------------------------------

type ClipboardEventInit struct {
	EventInit

	ClipboardData DataTransfer `json:"clipboardData"`
}

func (p ClipboardEventInit) toDict() js.Value {
	o := p.EventInit.toDict()
	o.Set("clipboardData", p.ClipboardData.JSValue())
	return o
}

// -------------8<---------------------------------------
