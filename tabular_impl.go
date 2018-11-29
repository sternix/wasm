// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type htmlTableElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTableElement() HTMLTableElement {
	if el := CurrentDocument().CreateElement("table"); el != nil {
		if table, ok := el.(HTMLTableElement); ok {
			return table
		}
	}
	return nil
}

func newHTMLTableElement(v js.Value) HTMLTableElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTableElementImpl) Caption() HTMLTableCaptionElement {
	return newHTMLTableCaptionElement(p.Get("caption"))
}

func (p *htmlTableElementImpl) SetCaption(caption HTMLTableCaptionElement) {
	p.Set("caption", caption.JSValue())
}

func (p *htmlTableElementImpl) CreateCaption() HTMLTableCaptionElement {
	return newHTMLTableCaptionElement(p.Call("createCaption"))
}

func (p *htmlTableElementImpl) DeleteCaption() {
	p.Call("deleteCaption")
}

func (p *htmlTableElementImpl) THead() HTMLTableSectionElement {
	return newHTMLTableSectionElement(p.Get("tHead"))
}

func (p *htmlTableElementImpl) SetTHead(section HTMLTableSectionElement) {
	p.Set("tHead", section.JSValue())
}

func (p *htmlTableElementImpl) CreateTHead() HTMLTableSectionElement {
	return newHTMLTableSectionElement(p.Call("createTHead"))
}

func (p *htmlTableElementImpl) DeleteTHead() {
	p.Call("deleteTHead")
}

func (p *htmlTableElementImpl) TFoot() HTMLTableSectionElement {
	return newHTMLTableSectionElement(p.Get("tFoot"))
}

func (p *htmlTableElementImpl) SetTFoot(section HTMLTableSectionElement) {
	p.Set("tFoot", section.JSValue())
}

func (p *htmlTableElementImpl) CreateTFoot() HTMLTableSectionElement {
	return newHTMLTableSectionElement(p.Call("createTFoot"))
}

func (p *htmlTableElementImpl) DeleteTFoot() {
	p.Call("deleteTFoot")
}

func (p *htmlTableElementImpl) TBodies() []HTMLTableSectionElement {
	bodies := newHTMLCollection(p.Get("tBodies"))
	if bodies != nil {
		var ret []HTMLTableSectionElement
		for i := 0; i < bodies.Length(); i++ {
			if r, ok := bodies.Item(i).(HTMLTableSectionElement); ok {
				ret = append(ret, r)
			}
		}
		return ret
	}
	return nil

}

func (p *htmlTableElementImpl) CreateTBody() HTMLTableSectionElement {
	return newHTMLTableSectionElement(p.Call("createTBody"))
}

func (p *htmlTableElementImpl) Rows() []HTMLTableRowElement {
	rows := newHTMLCollection(p.Get("rows"))
	if rows != nil {
		var ret []HTMLTableRowElement
		for i := 0; i < rows.Length(); i++ {
			if r, ok := rows.Item(i).(HTMLTableRowElement); ok {
				ret = append(ret, r)
			}
		}
		return ret
	}
	return nil
}

func (p *htmlTableElementImpl) InsertRow(index ...int) HTMLTableRowElement {
	switch len(index) {
	case 0:
		return newHTMLTableRowElement(p.Call("insertRow"))
	default:
		return newHTMLTableRowElement(p.Call("insertRow", index[0]))
	}

}

func (p *htmlTableElementImpl) DeleteRow(index int) {
	p.Call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableCaptionElementImpl struct {
	*htmlElementImpl
}

func newHTMLTableCaptionElement(v js.Value) HTMLTableCaptionElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableCaptionElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlTableSectionElementImpl struct {
	*htmlElementImpl
}

func newHTMLTableSectionElement(v js.Value) HTMLTableSectionElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableSectionElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTableSectionElementImpl) Rows() []HTMLTableRowElement {
	rows := newHTMLCollection(p.Get("rows"))
	if rows != nil {
		var ret []HTMLTableRowElement
		for i := 0; i < rows.Length(); i++ {
			if r, ok := rows.Item(i).(HTMLTableRowElement); ok {
				ret = append(ret, r)
			}
		}
		return ret
	}
	return nil
}

func (p *htmlTableSectionElementImpl) InsertRow(index ...int) HTMLTableRowElement {
	switch len(index) {
	case 0:
		return newHTMLTableRowElement(p.Call("insertRow"))
	default:
		return newHTMLTableRowElement(p.Call("insertRow", index[0]))
	}
}

func (p *htmlTableSectionElementImpl) DeleteRow(index int) {
	p.Call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableRowElementImpl struct {
	*htmlElementImpl
}

func newHTMLTableRowElement(v js.Value) HTMLTableRowElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableRowElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTableRowElementImpl) RowIndex() int {
	return p.Get("rowIndex").Int()
}

func (p *htmlTableRowElementImpl) SectionRowIndex() int {
	return p.Get("sectionRowIndex").Int()
}

func (p *htmlTableRowElementImpl) Cells() []HTMLTableCellElement {
	cells := newHTMLCollection(p.Get("cells"))
	if cells != nil {
		var ret []HTMLTableCellElement
		for i := 0; i < cells.Length(); i++ {
			if c, ok := cells.Item(i).(HTMLTableCellElement); ok {
				ret = append(ret, c)
			}
		}
		return ret
	}
	return nil

}

func (p *htmlTableRowElementImpl) InsertCell(index ...int) HTMLTableCellElement {
	switch len(index) {
	case 0:
		return newHTMLTableCellElement(p.Call("insertCell"))
	default:
		return newHTMLTableCellElement(p.Call("insertCell", index[0]))
	}
}

func (p *htmlTableRowElementImpl) DeleteCell(index int) {
	p.Call("deleteCell", index)
}

// -------------8<---------------------------------------

type htmlTableColElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTableColElement() HTMLTableColElement {
	if el := CurrentDocument().CreateElement("col"); el != nil {
		if col, ok := el.(HTMLTableColElement); ok {
			return col
		}
	}
	return nil
}

func NewHTMLTableColGroupElement() HTMLTableColElement {
	if el := CurrentDocument().CreateElement("colgroup"); el != nil {
		if colgroup, ok := el.(HTMLTableColElement); ok {
			return colgroup
		}
	}
	return nil
}

func newHTMLTableColElement(v js.Value) HTMLTableColElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableColElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTableColElementImpl) Span() int {
	return p.Get("span").Int()
}

func (p *htmlTableColElementImpl) SetSpan(span int) {
	p.Set("span", span)
}

// -------------8<---------------------------------------

type htmlTableCellElementImpl struct {
	*htmlElementImpl
}

func newHTMLTableCellElement(v js.Value) HTMLTableCellElement {
	if p := newHTMLTableCellElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLTableCellElementImpl(v js.Value) *htmlTableCellElementImpl {
	if isNil(v) {
		return nil
	}

	return &htmlTableCellElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTableCellElementImpl) ColSpan() int {
	return p.Get("colSpan").Int()
}

func (p *htmlTableCellElementImpl) SetColSpan(colspan int) {
	p.Set("colSpan", colspan)
}

func (p *htmlTableCellElementImpl) RowSpan() int {
	return p.Get("rowSpan").Int()
}

func (p *htmlTableCellElementImpl) SetRowSpan(rowspan int) {
	p.Set("rowSpan", rowspan)
}

func (p *htmlTableCellElementImpl) Headers() DOMTokenList {
	return newDOMTokenList(p.Get("headers"))
}

func (p *htmlTableCellElementImpl) CellIndex() int {
	return p.Get("cellIndex").Int()
}

// -------------8<---------------------------------------

type htmlTableDataCellElementImpl struct {
	*htmlTableCellElementImpl
}

func newHTMLTableDataCellElement(v js.Value) HTMLTableDataCellElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableDataCellElementImpl{
		htmlTableCellElementImpl: newHTMLTableCellElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlTableHeaderCellElementImpl struct {
	*htmlTableCellElementImpl
}

func NewHTMLTableHeaderCellElement() HTMLTableHeaderCellElement {
	if el := CurrentDocument().CreateElement("th"); el != nil {
		if th, ok := el.(HTMLTableHeaderCellElement); ok {
			return th
		}
	}
	return nil
}

func newHTMLTableHeaderCellElement(v js.Value) HTMLTableHeaderCellElement {
	if isNil(v) {
		return nil
	}

	return &htmlTableHeaderCellElementImpl{
		htmlTableCellElementImpl: newHTMLTableCellElementImpl(v),
	}
}

func (p *htmlTableHeaderCellElementImpl) Scope() string {
	return p.Get("scope").String()
}

func (p *htmlTableHeaderCellElementImpl) SetScope(scope string) {
	p.Set("scope", scope)
}

func (p *htmlTableHeaderCellElementImpl) Abbr() string {
	return p.Get("abbr").String()
}

func (p *htmlTableHeaderCellElementImpl) SetAbbr(abbr string) {
	p.Set("abbr", abbr)
}

// -------------8<---------------------------------------
