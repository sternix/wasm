// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlSlotElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLSlotElement(v Value) HTMLSlotElement {
	if v.valid() {
		return &htmlSlotElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlSlotElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlSlotElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlSlotElementImpl) AssignedNodes(options ...AssignedNodesOptions) []Node {
	if len(options) > 0 {
		return nodeListToSlice(p.call("assignedNodes", options[0].toJSObject()))
	}
	return nodeListToSlice(p.call("assignedNodes"))
}

func (p *htmlSlotElementImpl) AssignedElements(options ...AssignedNodesOptions) []Element {
	if len(options) > 0 {
		return elementArrayToSlice(p.call("assignedElements", options[0].toJSObject()))
	}
	return elementArrayToSlice(p.call("assignedElements"))
}

// -------------8<---------------------------------------

type htmlOrSVGElementImpl struct {
	Value
}

func wrapHTMLOrSVGElement(v Value) HTMLOrSVGElement {
	if v.valid() {
		return &htmlOrSVGElementImpl{
			Value: v,
		}
	}
	return nil
}

func (p *htmlOrSVGElementImpl) DataSet() map[string]string {
	return domStringMapToMap(p.get("dataset"))
}

func (p *htmlOrSVGElementImpl) Nonce() string {
	return p.get("nonce").toString()
}

func (p *htmlOrSVGElementImpl) SetNonce(nonce string) {
	p.set("nonce", nonce)
}

func (p *htmlOrSVGElementImpl) TabIndex() int {
	return p.get("tabIndex").toInt()
}

func (p *htmlOrSVGElementImpl) SetTabIndex(index int) {
	p.set("tabIndex", index)
}

func (p *htmlOrSVGElementImpl) Focus(options ...FocusOptions) {
	if len(options) > 0 {
		p.call("focus", options[0].toJSObject())
	} else {
		p.call("focus")
	}
}

func (p *htmlOrSVGElementImpl) Blur() {
	p.call("blur")
}

// -------------8<---------------------------------------

type elementContentEditableImpl struct {
	Value
}

func wrapElementContentEditable(v Value) ElementContentEditable {
	if v.valid() {
		return &elementContentEditableImpl{
			Value: v,
		}
	}
	return nil
}

func (p *elementContentEditableImpl) ContentEditable() string {
	return p.get("contentEditable").toString()
}

func (p *elementContentEditableImpl) SetContentEditable(ce string) {
	p.set("contentEditable", ce)
}

func (p *elementContentEditableImpl) IsContentEditable() bool {
	return p.get("isContentEditable").toBool()
}

func (p *elementContentEditableImpl) InputMode() string {
	return p.get("inputMode").toString()
}

func (p *elementContentEditableImpl) SetInputMode(im string) {
	p.set("inputMode", im)
}

// -------------8<---------------------------------------

type abortSignalImpl struct {
	*eventTargetImpl
}

func wrapAbortSignal(v Value) AbortSignal {
	if v.valid() {
		return &abortSignalImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *abortSignalImpl) Aborted() bool {
	return p.get("aborted").toBool()
}

func (p *abortSignalImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

// -------------8<---------------------------------------

type abortControllerImpl struct {
	Value
}

func wrapAbortController(v Value) AbortController {
	if v.valid() {
		return &abortControllerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *abortControllerImpl) Signal() AbortSignal {
	return wrapAbortSignal(p.get("signal"))
}

func (p *abortControllerImpl) Abort() {
	p.call("abort")
}

// -------------8<---------------------------------------

func NewAbortController() AbortController {
	return wrapAbortController(jsGlobal.get("AbortController"))
}

// -------------8<---------------------------------------
