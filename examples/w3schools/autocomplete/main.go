// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_autocomplete
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
	"strings"
)

func main() {
	input := wasm.CurrentDocument().ElementById("myInput").(wasm.HTMLInputElement)
	Autocomplete(input, countries)
	wasm.Loop()
}

func Autocomplete(input wasm.HTMLInputElement, arr []string) {
	var (
		currentFocus int
		doc          = wasm.CurrentDocument()
	)

	closeAllLists := func(elm wasm.Element) {
		for _, item := range doc.ElementsByClassName("autocomplete-items") {
			if !wasm.Equal(elm, item) && !wasm.Equal(elm, input) {
				item.ParentNode().RemoveChild(item)
			}
		}
	}

	removeActive := func(x []wasm.Element) {
		for _, e := range x {
			e.ClassList().Remove("autocomplete-active")
		}
	}

	addActive := func(x []wasm.Element) {
		removeActive(x)

		if currentFocus >= len(x) {
			currentFocus = 0
		}

		if currentFocus < 0 {
			currentFocus = len(x) - 1
		}

		x[currentFocus].ClassList().Add("autocomplete-active")
	}

	input.OnKeyDown(func(e wasm.KeyboardEvent) {
		elm := doc.ElementById(input.Id() + "autocomplete-list")
		if elm == nil {
			return
		}

		x := elm.ElementsByTagName("div")

		switch e.Key() {
		case wasm.KeyArrowDown:
			currentFocus++
			addActive(x)
		case wasm.KeyArrowUp:
			currentFocus--
			addActive(x)
		case wasm.KeyEnter:
			e.PreventDefault()
			if currentFocus > -1 {
				if len(x) >= currentFocus {
					x[currentFocus].(wasm.HTMLElement).Click()
				}
			}
		}
	})

	input.OnInput(func(wasm.InputEvent) {
		closeAllLists(nil)
		val := input.Value()
		if val == "" {
			return
		}

		val = strings.ToUpper(val)

		currentFocus = -1
		div := wasm.NewHTMLDivElement()
		div.SetId(input.Id() + "autocomplete-list")
		div.SetClassName("autocomplete-items")
		input.ParentNode().AppendChild(div)

		lval := len(val)

		for _, item := range arr {
			if lval > len(item) {
				continue
			}

			uitem := strings.ToUpper(item[:len(val)])
			if uitem == val {
				b := wasm.NewHTMLDivElement()
				innerHTML := fmt.Sprintf("<strong>%s</strong>%s", item[:len(val)], item[len(val):])
				innerHTML += fmt.Sprintf("<input type='hidden' value='%s'>", item)
				b.SetInnerHTML(innerHTML)

				//TODO: this div's EventHandler's don't Released
				b.OnClick(func(wasm.MouseEvent) {
					input.SetValue(b.ElementsByTagName("input")[0].(wasm.HTMLInputElement).Value())
					closeAllLists(nil)
				})
				div.AppendChild(b)
			}
		}

	})

	doc.OnClick(func(e wasm.MouseEvent) {
		closeAllLists(e.Target().(wasm.Element))
	})
}
