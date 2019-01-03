// +build wasm,js

// https://www.w3schools.com/howto/tryit.asp?filename=tryhow_css_modal2
package main

import (
	"github.com/sternix/wasm"
)

func main() {
	win := wasm.CurrentWindow()
	doc := win.Document()

	modal := doc.ElementById("myModal").(wasm.HTMLDivElement)
	btn := doc.ElementById("myBtn").(wasm.HTMLButtonElement)
	span := doc.ElementsByClassName("close")[0].(wasm.HTMLElement)

	btn.OnClick(func(wasm.MouseEvent) {
		modal.Style().SetProperty("display", "block")
	})

	span.OnClick(func(wasm.MouseEvent) {
		modal.Style().SetProperty("display", "none")
	})

	win.OnClick(func(e wasm.MouseEvent) {
		if wasm.Equal(e.Target(), modal) {
			modal.Style().SetProperty("display", "none")
		}
	})

	wasm.Loop()
}
