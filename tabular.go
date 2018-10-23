// +build js,wasm

package wasm

type (

	// https://www.w3.org/TR/html52/tabular-data.html#htmltableelement
	HTMLTableElement interface {
		HTMLElement

		Caption() HTMLTableCaptionElement
		SetCaption(HTMLTableCaptionElement)
		CreateCaption() HTMLTableCaptionElement
		DeleteCaption()
		THead() HTMLTableSectionElement
		SetTHead(HTMLTableSectionElement)
		CreateTHead() HTMLTableSectionElement
		DeleteTHead()
		TFoot() HTMLTableSectionElement
		SetTFoot(HTMLTableSectionElement)
		CreateTFoot() HTMLTableSectionElement
		DeleteTFoot()
		TBodies() HTMLCollection
		CreateTBody() HTMLTableSectionElement
		Rows() HTMLCollection
		InsertRow(...int) HTMLTableRowElement
		DeleteRow(int)
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltablecaptionelement
	HTMLTableCaptionElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltablesectionelement
	HTMLTableSectionElement interface {
		HTMLElement

		Rows() HTMLCollection
		InsertRow(...int) HTMLElement
		DeleteRow(int)
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltablerowelement
	HTMLTableRowElement interface {
		HTMLElement

		RowIndex() int
		SectionRowIndex() int
		Cells() HTMLCollection
		InsertCell(...int) HTMLElement // return <td> element
		DeleteCell(int)
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltablecolelement
	HTMLTableColElement interface {
		HTMLElement

		Span() int
		SetSpan(int)
	}

	//  https://www.w3.org/TR/html52/tabular-data.html#htmltablecellelement
	HTMLTableCellElement interface {
		HTMLElement

		ColSpan() int
		SetColSpan(int)
		RowSpan() int
		SetRowSpan(int)
		Headers() DOMTokenList
		CellIndex() int
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltabledatacellelement
	HTMLTableDataCellElement interface {
		HTMLTableCellElement
	}

	// https://www.w3.org/TR/html52/tabular-data.html#htmltableheadercellelement
	HTMLTableHeaderCellElement interface {
		HTMLTableCellElement

		Scope() string
		SetScope(string)
		Abbr() string
		SetAbbr(string)
	}
)
