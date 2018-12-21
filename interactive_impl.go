// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type htmlDetailsElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLDetailsElement(v js.Value) HTMLDetailsElement {
	if isNil(v) {
		return nil
	}

	return &htmlDetailsElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
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

func wrapHTMLDialogElement(v js.Value) HTMLDialogElement {
	if isNil(v) {
		return nil
	}
	return &htmlDialogElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
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
