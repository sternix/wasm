// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlDetailsElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLDetailsElement(v Value) HTMLDetailsElement {
	if v.valid() {
		return &htmlDetailsElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDetailsElementImpl) Open() bool {
	return p.get("open").toBool()
}

func (p *htmlDetailsElementImpl) SetOpen(b bool) {
	p.set("open", b)
}

// -------------8<---------------------------------------

type htmlDialogElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLDialogElement(v Value) HTMLDialogElement {
	if v.valid() {
		return &htmlDialogElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDialogElementImpl) Open() bool {
	return p.get("open").toBool()
}

func (p *htmlDialogElementImpl) SetOpen(b bool) {
	p.set("open", b)
}

func (p *htmlDialogElementImpl) ReturnValue() string {
	return p.get("returnValue").toString()
}

func (p *htmlDialogElementImpl) SetReturnValue(returnValue string) {
	p.set("returnValue", returnValue)
}

func (p *htmlDialogElementImpl) Show(anchor ...interface{}) {
	if len(anchor) > 0 {
		switch x := anchor[0].(type) {
		case MouseEvent, Element:
			p.call("show", JSValue(x))
		}
	} else {
		p.call("show")
	}
}

func (p *htmlDialogElementImpl) ShowModal(anchor ...interface{}) {
	if len(anchor) > 0 {
		switch x := anchor[0].(type) {
		case MouseEvent:
			p.call("showModal", JSValue(x))
		case Element:
			p.call("showModal", JSValue(x))
		}
	} else {
		p.call("showModal")
	}
}

func (p *htmlDialogElementImpl) Close(returnValue ...string) {
	switch len(returnValue) {
	case 0:
		p.call("close")
	default:
		p.call("close", returnValue[0])
	}
}

// -------------8<---------------------------------------
