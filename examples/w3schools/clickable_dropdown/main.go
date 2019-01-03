// +build wasm,js

//https://www.w3schools.com/howto/tryit.asp?filename=tryhow_css_js_dropdown
package main

import (
	"github.com/sternix/wasm"
)

func main() {
	win := wasm.CurrentWindow()
	doc := win.Document()

	dropBtns := doc.ElementsByClassName("dropbtn")
	myDropDown := doc.ElementById("myDropdown").(wasm.HTMLDivElement)

	for i := 0; i < len(dropBtns); i++ {
		btn := dropBtns[i].(wasm.HTMLButtonElement)
		btn.OnClick(func(wasm.MouseEvent) {
			myDropDown.ClassList().Toggle("show")
		})
	}

	win.OnClick(func(e wasm.MouseEvent) {
		if !e.Target().(wasm.HTMLElement).Matches(".dropbtn") {
			dropdowns := doc.ElementsByClassName("dropdown-content")
			for i := 0; i < len(dropdowns); i++ {
				openDropDown := dropdowns[i].(wasm.HTMLDivElement)
				if openDropDown.ClassList().Contains("show") {
					openDropDown.ClassList().Remove("show")
				}
			}
		}
	})

	wasm.Loop()
}
