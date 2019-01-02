// +build js,wasm

package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func TestTabular() {
	titles := []string{"X", "Y", "Z"}
	table := wasm.NewHTMLTableElement()
	caption := table.CreateCaption()
	caption.SetInnerHTML("Go Wasm Table Caption")

	head := table.CreateTHead()
	hr := head.InsertRow()
	hc := hr.InsertCell()
	hc.SetColSpan(3)
	hc.SetInnerHTML("Go Wasm Table Head")

	row := table.InsertRow()
	for _, t := range titles {
		td := row.InsertCell()
		td.SetInnerText(t)
	}

	for i := 0; i < 15; i++ {
		row = table.InsertRow()
		for j := 0; j < len(titles); j++ {
			td := row.InsertCell()
			td.SetInnerText(fmt.Sprintf("%d-%d", i, j))
		}
	}

	foot := table.CreateTFoot()
	fr := foot.InsertRow()
	fc := fr.InsertCell()
	fc.SetColSpan(3)
	fc.SetInnerHTML("Go Wasm Table Foot")

	fmt.Printf("Table Row Size: %d\n", len(table.Rows()))

	wasm.CurrentDocument().Body().AppendChild(table)
}
