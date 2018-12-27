// +build wasm,js

//https://www.w3schools.com/howto/tryit.asp?filename=tryhow_css_js_dropdown
package main

import (
	. "github.com/sternix/wasm"
)

func main() {
	win := CurrentWindow()
	doc := win.Document()

	dropBtns := doc.ElementsByClassName("dropbtn")
	myDropDown := doc.ElementById("myDropdown").(HTMLDivElement)

	for i := 0; i < len(dropBtns); i++ {
		btn := dropBtns[i].(HTMLButtonElement)
		btn.OnClick(func(MouseEvent) {
			myDropDown.ClassList().Toggle("show")
		})
	}

	win.OnClick(func(e MouseEvent) {
		if !e.Target().(HTMLElement).Matches(".dropbtn") {
			dropdowns := doc.ElementsByClassName("dropdown-content")
			for i := 0; i < len(dropdowns); i++ {
				openDropDown := dropdowns[i].(HTMLDivElement)
				if openDropDown.ClassList().Contains("show") {
					openDropDown.ClassList().Remove("show")
				}
			}
		}
	})

	Wait()
}
