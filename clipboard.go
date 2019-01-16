// +build js,wasm

// https://w3c.github.io/clipboard-apis/#idl-index
package wasm

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

func (p ClipboardPermissionDescriptor) JSValue() jsValue {
	o := p.PermissionDescriptor.JSValue()
	o.Set("allowWithoutGesture", p.AllowWithoutGesture)
	return o
}

// -------------8<---------------------------------------

type ClipboardEventInit struct {
	EventInit

	ClipboardData DataTransfer
}

func (p ClipboardEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("clipboardData", JSValueOf(p.ClipboardData))
	return o
}

// -------------8<---------------------------------------
