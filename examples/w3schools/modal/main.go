// +build wasm,js

// https://www.w3schools.com/howto/tryit.asp?filename=tryhow_css_modal2
package main

import (
	. "github.com/sternix/wasm"
)

func main() {
	win := CurrentWindow()
	doc := win.Document()

	modal := doc.ElementById("myModal").(HTMLDivElement)
	btn := doc.ElementById("myBtn").(HTMLButtonElement)
	span := doc.ElementsByClassName("close")[0].(HTMLElement)

	btn.OnClick(func(MouseEvent) {
		modal.Style().SetProperty("display", "block")
	})

	span.OnClick(func(MouseEvent) {
		modal.Style().SetProperty("display", "none")
	})

	win.OnClick(func(e MouseEvent) {
		if Equal(e.Target(), modal) {
			modal.Style().SetProperty("display", "none")
		}
	})

	Wait()
}
