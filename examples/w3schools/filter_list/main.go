// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_filter_list
*/
package main

import (
	"github.com/sternix/wasm"
	"strings"
)

var (
	doc   wasm.Document
	input wasm.HTMLInputElement
	ul    wasm.Element
)

func main() {
	doc = wasm.CurrentDocument()
	input = doc.ElementById("myInput").(wasm.HTMLInputElement)
	ul = doc.ElementById("myUL")
	input.OnKeyUp(filterList)
	wasm.Loop()
}

func filterList(wasm.KeyboardEvent) {
	filter := strings.ToUpper(input.Value())
	lis := ul.ElementsByTagName("li")
	for _, li := range lis {
		a := li.ElementsByTagName("a")[0]
		txtValue := strings.ToUpper(a.TextContent())
		hli := li.(wasm.HTMLElement)
		if strings.Index(txtValue, filter) > -1 {
			hli.Style().SetProperty("display", "")
		} else {
			hli.Style().SetProperty("display", "none")
		}
	}
}
