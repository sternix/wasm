// +build js,wasm

package wasm

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

func wrapHTMLTableElement(v Value) HTMLTableElement {
	if v.valid() {
		return &htmlTableElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableElementImpl) Caption() HTMLTableCaptionElement {
	return wrapHTMLTableCaptionElement(p.get("caption"))
}

func (p *htmlTableElementImpl) SetCaption(caption HTMLTableCaptionElement) {
	p.set("caption", JSValue(caption))
}

func (p *htmlTableElementImpl) CreateCaption() HTMLTableCaptionElement {
	return wrapHTMLTableCaptionElement(p.call("createCaption"))
}

func (p *htmlTableElementImpl) DeleteCaption() {
	p.call("deleteCaption")
}

func (p *htmlTableElementImpl) THead() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.get("tHead"))
}

func (p *htmlTableElementImpl) SetTHead(section HTMLTableSectionElement) {
	p.set("tHead", JSValue(section))
}

func (p *htmlTableElementImpl) CreateTHead() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.call("createTHead"))
}

func (p *htmlTableElementImpl) DeleteTHead() {
	p.call("deleteTHead")
}

func (p *htmlTableElementImpl) TFoot() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.get("tFoot"))
}

func (p *htmlTableElementImpl) SetTFoot(section HTMLTableSectionElement) {
	p.set("tFoot", JSValue(section))
}

func (p *htmlTableElementImpl) CreateTFoot() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.call("createTFoot"))
}

func (p *htmlTableElementImpl) DeleteTFoot() {
	p.call("deleteTFoot")
}

func (p *htmlTableElementImpl) TBodies() []HTMLTableSectionElement {
	if bodies := wrapHTMLCollection(p.get("tBodies")); bodies != nil && bodies.Length() > 0 {
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
	return wrapHTMLTableSectionElement(p.call("createTBody"))
}

func (p *htmlTableElementImpl) Rows() []HTMLTableRowElement {
	if rows := wrapHTMLCollection(p.get("rows")); rows != nil && rows.Length() > 0 {
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
		return wrapHTMLTableRowElement(p.call("insertRow"))
	default:
		return wrapHTMLTableRowElement(p.call("insertRow", index[0]))
	}

}

func (p *htmlTableElementImpl) DeleteRow(index int) {
	p.call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableCaptionElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableCaptionElement(v Value) HTMLTableCaptionElement {
	if v.valid() {
		return &htmlTableCaptionElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlTableSectionElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableSectionElement(v Value) HTMLTableSectionElement {
	if v.valid() {
		return &htmlTableSectionElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableSectionElementImpl) Rows() []HTMLTableRowElement {
	if rows := wrapHTMLCollection(p.get("rows")); rows != nil && rows.Length() > 0 {
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
		return wrapHTMLTableRowElement(p.call("insertRow"))
	default:
		return wrapHTMLTableRowElement(p.call("insertRow", index[0]))
	}
}

func (p *htmlTableSectionElementImpl) DeleteRow(index int) {
	p.call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableRowElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableRowElement(v Value) HTMLTableRowElement {
	if v.valid() {
		return &htmlTableRowElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableRowElementImpl) RowIndex() int {
	return p.get("rowIndex").toInt()
}

func (p *htmlTableRowElementImpl) SectionRowIndex() int {
	return p.get("sectionRowIndex").toInt()
}

func (p *htmlTableRowElementImpl) Cells() []HTMLTableCellElement {
	if cells := wrapHTMLCollection(p.get("cells")); cells != nil && cells.Length() > 0 {
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
		return wrapHTMLTableCellElement(p.call("insertCell"))
	default:
		return wrapHTMLTableCellElement(p.call("insertCell", index[0]))
	}
}

func (p *htmlTableRowElementImpl) DeleteCell(index int) {
	p.call("deleteCell", index)
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

func wrapHTMLTableColElement(v Value) HTMLTableColElement {
	if v.valid() {
		return &htmlTableColElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableColElementImpl) Span() int {
	return p.get("span").toInt()
}

func (p *htmlTableColElementImpl) SetSpan(span int) {
	p.set("span", span)
}

// -------------8<---------------------------------------

type htmlTableCellElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableCellElement(v Value) HTMLTableCellElement {
	if p := newHTMLTableCellElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLTableCellElementImpl(v Value) *htmlTableCellElementImpl {
	if v.valid() {
		return &htmlTableCellElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableCellElementImpl) ColSpan() int {
	return p.get("colSpan").toInt()
}

func (p *htmlTableCellElementImpl) SetColSpan(colspan int) {
	p.set("colSpan", colspan)
}

func (p *htmlTableCellElementImpl) RowSpan() int {
	return p.get("rowSpan").toInt()
}

func (p *htmlTableCellElementImpl) SetRowSpan(rowspan int) {
	p.set("rowSpan", rowspan)
}

func (p *htmlTableCellElementImpl) Headers() DOMTokenList {
	return wrapDOMTokenList(p.get("headers"))
}

func (p *htmlTableCellElementImpl) CellIndex() int {
	return p.get("cellIndex").toInt()
}

// -------------8<---------------------------------------

type htmlTableDataCellElementImpl struct {
	*htmlTableCellElementImpl
}

func wrapHTMLTableDataCellElement(v Value) HTMLTableDataCellElement {
	if v.valid() {
		return &htmlTableDataCellElementImpl{
			htmlTableCellElementImpl: newHTMLTableCellElementImpl(v),
		}
	}
	return nil
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

func wrapHTMLTableHeaderCellElement(v Value) HTMLTableHeaderCellElement {
	if v.valid() {
		return &htmlTableHeaderCellElementImpl{
			htmlTableCellElementImpl: newHTMLTableCellElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableHeaderCellElementImpl) Scope() string {
	return p.get("scope").toString()
}

func (p *htmlTableHeaderCellElementImpl) SetScope(scope string) {
	p.set("scope", scope)
}

func (p *htmlTableHeaderCellElementImpl) Abbr() string {
	return p.get("abbr").toString()
}

func (p *htmlTableHeaderCellElementImpl) SetAbbr(abbr string) {
	p.set("abbr", abbr)
}

// -------------8<---------------------------------------
