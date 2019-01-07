// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlDetailsElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLDetailsElement(v Value) HTMLDetailsElement {
	if v.Valid() {
		return &htmlDetailsElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDetailsElementImpl) Open() bool {
	return p.Get("open").Bool()
}

func (p *htmlDetailsElementImpl) SetOpen(b bool) {
	p.Set("open", b)
}

// -------------8<---------------------------------------

type htmlDialogElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLDialogElement(v Value) HTMLDialogElement {
	if v.Valid() {
		return &htmlDialogElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDialogElementImpl) Open() bool {
	return p.Get("open").Bool()
}

func (p *htmlDialogElementImpl) SetOpen(b bool) {
	p.Set("open", b)
}

func (p *htmlDialogElementImpl) ReturnValue() string {
	return p.Get("returnValue").String()
}

func (p *htmlDialogElementImpl) SetReturnValue(returnValue string) {
	p.Set("returnValue", returnValue)
}

func (p *htmlDialogElementImpl) Show(anchor ...interface{}) {
	if len(anchor) > 0 {
		switch x := anchor[0].(type) {
		case MouseEvent, Element:
			p.Call("show", JSValue(x))
		}
	} else {
		p.Call("show")
	}
}

func (p *htmlDialogElementImpl) ShowModal(anchor ...interface{}) {
	if len(anchor) > 0 {
		switch x := anchor[0].(type) {
		case MouseEvent:
			p.Call("showModal", JSValue(x))
		case Element:
			p.Call("showModal", JSValue(x))
		}
	} else {
		p.Call("showModal")
	}
}

func (p *htmlDialogElementImpl) Close(returnValue ...string) {
	switch len(returnValue) {
	case 0:
		p.Call("close")
	default:
		p.Call("close", returnValue[0])
	}
}

// -------------8<---------------------------------------
