// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_filter_table
*/
package main

import (
	"github.com/sternix/wasm"
	"strings"
)

var (
	doc   wasm.Document
	input wasm.HTMLInputElement
	table wasm.Element
)

func main() {
	doc = wasm.CurrentDocument()
	input = doc.ElementById("myInput").(wasm.HTMLInputElement)
	table = doc.ElementById("myTable")
	input.OnKeyUp(filterTable)
	wasm.Loop()
}

func filterTable(wasm.KeyboardEvent) {
	filter := strings.ToUpper(input.Value())
	trs := table.ElementsByTagName("tr")
	for _, tr := range trs {
		htr := tr.(wasm.HTMLElement)
		tds := tr.ElementsByTagName("td")
		if len(tds) > 0 {
			td := tds[0].(wasm.HTMLElement)
			txtValue := strings.ToUpper(td.TextContent())
			if strings.Index(txtValue, filter) > -1 {
				htr.Style().SetProperty("display", "")
			} else {
				htr.Style().SetProperty("display", "none")
			}
		}
	}
}
