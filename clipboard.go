// +build js,wasm

package wasm

// https://w3c.github.io/clipboard-apis/#idl-index

type (
	ClipboardEventInit struct {
		EventInit

		ClipboardData DataTransfer `json:"clipboardData"`
	}

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

	ClipboardPermissionDescriptor struct {
		PermissionDescriptor

		AllowWithoutGesture bool `json:"allowWithoutGesture"`
	}
)
