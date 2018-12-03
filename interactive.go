// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/html52/interactive-elements.html#htmldetailselement
	HTMLDetailsElement interface {
		HTMLElement

		Open() bool
		SetOpen(bool)
	}

	/*
	   in firefox -> about:config -> dom.dialog_element.enabled -> true
	*/

	// https://www.w3.org/TR/html52/interactive-elements.html#htmldialogelement
	HTMLDialogElement interface {
		HTMLElement

		Open() bool
		SetOpen(bool)
		ReturnValue() string
		SetReturnValue(string)
		Show(...interface{})      // MouseEvent or Element
		ShowModal(...interface{}) // MouseEvent or Element
		Close(...string)          // optional returnValue
	}
)
