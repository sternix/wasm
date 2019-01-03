// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_close_list_items
*/
package main

import (
	"github.com/sternix/wasm"
)

func main() {
	doc := wasm.CurrentDocument()
	closeButtons := doc.ElementsByClassName("close")

	for _, btn := range closeButtons {
		if closeButton, ok := btn.(wasm.HTMLElement); ok {
			closeButton.OnClick(func(wasm.MouseEvent) {
				closeButton.ParentElement().(wasm.HTMLElement).Style().SetProperty("display", "none")
			})
		}
	}

	wasm.Loop()
}
