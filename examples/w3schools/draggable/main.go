// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_draggable
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func main() {
	myDiv := wasm.CurrentDocument().ElementById("mydiv").(wasm.HTMLElement)
	dragElement(myDiv)
	wasm.Loop()
}

func dragElement(elm wasm.HTMLElement) {
	var (
		pos1, pos2, pos3, pos4 float64
		docMouseUpEH           wasm.EventHandler
		docMouseMoveEH         wasm.EventHandler
	)

	doc := wasm.CurrentDocument()

	elementDrag := func(e wasm.MouseEvent) {
		e.PreventDefault()
		pos1 = pos3 - e.ClientX()
		pos2 = pos4 - e.ClientY()
		pos3 = e.ClientX()
		pos4 = e.ClientY()
		elm.Style().SetProperty("top", fmt.Sprintf("%fpx", float64(elm.OffsetTop())-pos2))
		elm.Style().SetProperty("left", fmt.Sprintf("%fpx", float64(elm.OffsetLeft())-pos1))
	}

	closeDragElement := func(wasm.MouseEvent) {
		if docMouseUpEH != nil {
			docMouseUpEH.Remove()
			docMouseUpEH = nil
		}

		if docMouseMoveEH != nil {
			docMouseMoveEH.Remove()
			docMouseMoveEH = nil
		}
	}

	dragMouseDown := func(e wasm.MouseEvent) {
		e.PreventDefault()
		pos3 = e.ClientX()
		pos4 = e.ClientY()
		docMouseUpEH = doc.OnMouseUp(closeDragElement)
		docMouseMoveEH = doc.OnMouseMove(elementDrag)
	}

	if headerElm := doc.ElementById(elm.Id() + "header").(wasm.HTMLElement); headerElm != nil {
		headerElm.OnMouseDown(dragMouseDown)
	} else {
		elm.OnMouseDown(dragMouseDown)
	}
}
