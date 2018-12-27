// +build wasm,js

//https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_accordion_symbol
package main

import (
	"fmt"
	. "github.com/sternix/wasm"
)

func main() {
	doc := CurrentDocument()
	acc := doc.ElementsByClassName("accordion")

	for i := 0; i < len(acc); i++ {
		btn := acc[i].(HTMLButtonElement)
		btn.OnClick(func(MouseEvent) {
			this := btn
			this.ClassList().Toggle("active")
			panel := this.NextElementSibling().(HTMLDivElement)
			if panel.Style().PropertyValue("max-height") != "" {
				panel.Style().SetProperty("max-height", "")
			} else {
				panel.Style().SetProperty("max-height", fmt.Sprintf("%dpx", panel.ScrollHeight()))
			}
		})
	}

	Wait()
}
