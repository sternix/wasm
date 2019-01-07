// +build js,wasm

package wasm

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

func (p ClipboardPermissionDescriptor) toJSObject() Value {
	o := p.PermissionDescriptor.toJSObject()
	o.Set("allowWithoutGesture", p.AllowWithoutGesture)
	return o
}

// -------------8<---------------------------------------

type ClipboardEventInit struct {
	EventInit

	ClipboardData DataTransfer
}

func (p ClipboardEventInit) toJSObject() Value {
	o := p.EventInit.toJSObject()
	o.Set("clipboardData", JSValue(p.ClipboardData))
	return o
}

// -------------8<---------------------------------------
