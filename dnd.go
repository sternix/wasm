// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://www.w3.org/TR/html52/editing.html#datatransfer
	DataTransfer interface {
		js.Wrapper

		DropEffect() string
		SetDropEffect(string)
		EffectAllowed() string
		SetEffectAllowed(string)
		Items() DataTransferItemList
		SetDragImage(Element, int, int)
		Types() []string
		Data(string) string
		SetData(string, string)
		ClearData(...string)
		Files() []File //FileList
	}

	// https://www.w3.org/TR/html52/editing.html#datatransferitemlist
	DataTransferItemList interface {
		js.Wrapper

		Length() int
		Index(int) DataTransferItem // Get
		Add(string, string) DataTransferItem
		AddFile(File) DataTransferItem
		Remove(int)
		Clear()
	}

	// https://www.w3.org/TR/html52/editing.html#datatransferitem
	DataTransferItem interface {
		js.Wrapper

		Kind() string
		Type() string
		AsString(FunctionStringCallback)
		AsFile() File
	}

	// https://www.w3.org/TR/html52/editing.html#the-dragevent-interface
	DragEvent interface {
		MouseEvent

		DataTransfer() DataTransfer
	}

	// https://www.w3.org/TR/html52/editing.html#dictdef-drageventinit
	DragEventInit struct {
		MouseEventInit

		DataTransfer DataTransfer `json:"dataTransfer"`
	}
)
