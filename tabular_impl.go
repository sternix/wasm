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
	if v.Valid() {
		return &htmlTableElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableElementImpl) Caption() HTMLTableCaptionElement {
	return wrapHTMLTableCaptionElement(p.Get("caption"))
}

func (p *htmlTableElementImpl) SetCaption(caption HTMLTableCaptionElement) {
	p.Set("caption", JSValue(caption))
}

func (p *htmlTableElementImpl) CreateCaption() HTMLTableCaptionElement {
	return wrapHTMLTableCaptionElement(p.Call("createCaption"))
}

func (p *htmlTableElementImpl) DeleteCaption() {
	p.Call("deleteCaption")
}

func (p *htmlTableElementImpl) THead() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.Get("tHead"))
}

func (p *htmlTableElementImpl) SetTHead(section HTMLTableSectionElement) {
	p.Set("tHead", JSValue(section))
}

func (p *htmlTableElementImpl) CreateTHead() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.Call("createTHead"))
}

func (p *htmlTableElementImpl) DeleteTHead() {
	p.Call("deleteTHead")
}

func (p *htmlTableElementImpl) TFoot() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.Get("tFoot"))
}

func (p *htmlTableElementImpl) SetTFoot(section HTMLTableSectionElement) {
	p.Set("tFoot", JSValue(section))
}

func (p *htmlTableElementImpl) CreateTFoot() HTMLTableSectionElement {
	return wrapHTMLTableSectionElement(p.Call("createTFoot"))
}

func (p *htmlTableElementImpl) DeleteTFoot() {
	p.Call("deleteTFoot")
}

func (p *htmlTableElementImpl) TBodies() []HTMLTableSectionElement {
	if bodies := wrapHTMLCollection(p.Get("tBodies")); bodies != nil && bodies.Length() > 0 {
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
	return wrapHTMLTableSectionElement(p.Call("createTBody"))
}

func (p *htmlTableElementImpl) Rows() []HTMLTableRowElement {
	if rows := wrapHTMLCollection(p.Get("rows")); rows != nil && rows.Length() > 0 {
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
		return wrapHTMLTableRowElement(p.Call("insertRow"))
	default:
		return wrapHTMLTableRowElement(p.Call("insertRow", index[0]))
	}

}

func (p *htmlTableElementImpl) DeleteRow(index int) {
	p.Call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableCaptionElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableCaptionElement(v Value) HTMLTableCaptionElement {
	if v.Valid() {
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
	if v.Valid() {
		return &htmlTableSectionElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableSectionElementImpl) Rows() []HTMLTableRowElement {
	if rows := wrapHTMLCollection(p.Get("rows")); rows != nil && rows.Length() > 0 {
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
		return wrapHTMLTableRowElement(p.Call("insertRow"))
	default:
		return wrapHTMLTableRowElement(p.Call("insertRow", index[0]))
	}
}

func (p *htmlTableSectionElementImpl) DeleteRow(index int) {
	p.Call("deleteRow", index)
}

// -------------8<---------------------------------------

type htmlTableRowElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLTableRowElement(v Value) HTMLTableRowElement {
	if v.Valid() {
		return &htmlTableRowElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTableRowElementImpl) RowIndex() int {
	return p.Get("rowIndex").Int()
}

func (p *htmlTableRowElementImpl) SectionRowIndex() int {
	return p.Get("sectionRowIndex").Int()
}

func (p *htmlTableRowElementImpl) Cells() []HTMLTableCellElement {
	if cells := wrapHTMLCollection(p.Get("cells")); cells != nil && cells.Length() > 0 {
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
		return wrapHTMLTableCellElement(p.Call("insertCell"))
	default:
		return wrapHTMLTableCellElement(p.Call("insertCell", index[0]))
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

func wrapHTMLTableColElement(v Value) HTMLTableColElement {
	if v.Valid() {
		return &htmlTableColElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
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

func wrapHTMLTableCellElement(v Value) HTMLTableCellElement {
	if p := newHTMLTableCellElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLTableCellElementImpl(v Value) *htmlTableCellElementImpl {
	if v.Valid() {
		return &htmlTableCellElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
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
	return wrapDOMTokenList(p.Get("headers"))
}

func (p *htmlTableCellElementImpl) CellIndex() int {
	return p.Get("cellIndex").Int()
}

// -------------8<---------------------------------------

type htmlTableDataCellElementImpl struct {
	*htmlTableCellElementImpl
}

func wrapHTMLTableDataCellElement(v Value) HTMLTableDataCellElement {
	if v.Valid() {
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
	if v.Valid() {
		return &htmlTableHeaderCellElementImpl{
			htmlTableCellElementImpl: newHTMLTableCellElementImpl(v),
		}
	}
	return nil
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
