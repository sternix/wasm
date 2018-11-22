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

func (p ClipboardPermissionDescriptor) toMap() map[string]interface{} {
	m := p.PermissionDescriptor.toMap()
	m["allowWithoutGesture"] = p.AllowWithoutGesture
	return m
}

// -------------8<---------------------------------------

type ClipboardEventInit struct {
	EventInit

	ClipboardData DataTransfer `json:"clipboardData"`
}

func (p ClipboardEventInit) toMap() map[string]interface{} {
	m := p.EventInit.toMap()
	m["clipboardData"] = p.ClipboardData.JSValue()
	return m
}

// -------------8<---------------------------------------
