// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlSlotElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLSlotElement(v Value) HTMLSlotElement {
	if v.Valid() {
		return &htmlSlotElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlSlotElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlSlotElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlSlotElementImpl) AssignedNodes(options ...AssignedNodesOptions) []Node {
	if len(options) > 0 {
		return nodeListToSlice(p.Call("assignedNodes", options[0].toJSObject()))
	}
	return nodeListToSlice(p.Call("assignedNodes"))
}

func (p *htmlSlotElementImpl) AssignedElements(options ...AssignedNodesOptions) []Element {
	if len(options) > 0 {
		return elementArrayToSlice(p.Call("assignedElements", options[0].toJSObject()))
	}
	return elementArrayToSlice(p.Call("assignedElements"))
}

// -------------8<---------------------------------------

type htmlOrSVGElementImpl struct {
	Value
}

func wrapHTMLOrSVGElement(v Value) HTMLOrSVGElement {
	if v.Valid() {
		return &htmlOrSVGElementImpl{
			Value: v,
		}
	}
	return nil
}

func (p *htmlOrSVGElementImpl) DataSet() map[string]string {
	return domStringMapToMap(p.Get("dataset"))
}

func (p *htmlOrSVGElementImpl) Nonce() string {
	return p.Get("nonce").String()
}

func (p *htmlOrSVGElementImpl) SetNonce(nonce string) {
	p.Set("nonce", nonce)
}

func (p *htmlOrSVGElementImpl) TabIndex() int {
	return p.Get("tabIndex").Int()
}

func (p *htmlOrSVGElementImpl) SetTabIndex(index int) {
	p.Set("tabIndex", index)
}

func (p *htmlOrSVGElementImpl) Focus(options ...FocusOptions) {
	if len(options) > 0 {
		p.Call("focus", options[0].toJSObject())
	} else {
		p.Call("focus")
	}
}

func (p *htmlOrSVGElementImpl) Blur() {
	p.Call("blur")
}

// -------------8<---------------------------------------

type elementContentEditableImpl struct {
	Value
}

func wrapElementContentEditable(v Value) ElementContentEditable {
	if v.Valid() {
		return &elementContentEditableImpl{
			Value: v,
		}
	}
	return nil
}

func (p *elementContentEditableImpl) ContentEditable() string {
	return p.Get("contentEditable").String()
}

func (p *elementContentEditableImpl) SetContentEditable(ce string) {
	p.Set("contentEditable", ce)
}

func (p *elementContentEditableImpl) IsContentEditable() bool {
	return p.Get("isContentEditable").Bool()
}

func (p *elementContentEditableImpl) InputMode() string {
	return p.Get("inputMode").String()
}

func (p *elementContentEditableImpl) SetInputMode(im string) {
	p.Set("inputMode", im)
}

// -------------8<---------------------------------------

type abortSignalImpl struct {
	*eventTargetImpl
}

func wrapAbortSignal(v Value) AbortSignal {
	if v.Valid() {
		return &abortSignalImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *abortSignalImpl) Aborted() bool {
	return p.Get("aborted").Bool()
}

func (p *abortSignalImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

// -------------8<---------------------------------------

type abortControllerImpl struct {
	Value
}

func wrapAbortController(v Value) AbortController {
	if v.Valid() {
		return &abortControllerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *abortControllerImpl) Signal() AbortSignal {
	return wrapAbortSignal(p.Get("signal"))
}

func (p *abortControllerImpl) Abort() {
	p.Call("abort")
}

// -------------8<---------------------------------------

func NewAbortController() AbortController {
	return wrapAbortController(jsGlobal.Get("AbortController"))
}

// -------------8<---------------------------------------
