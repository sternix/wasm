// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/html52/editing.html#datatransfer
	DataTransfer interface {
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
		Length() uint
		Index(uint) DataTransferItem // Get
		Add(string, string) DataTransferItem
		AddFile(File) DataTransferItem
		Remove(uint)
		Clear()
	}

	// https://www.w3.org/TR/html52/editing.html#datatransferitem
	DataTransferItem interface {
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
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/html52/editing.html#dictdef-drageventinit
type DragEventInit struct {
	MouseEventInit

	DataTransfer DataTransfer
}

func (p DragEventInit) JSValue() jsValue {
	o := p.MouseEventInit.JSValue()
	o.Set("dataTransfer", JSValueOf(p.DataTransfer))
	return o
}
