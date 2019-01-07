// +build js,wasm

package wasm

type (
	// https://dom.spec.whatwg.org/#nonelementparentnode
	NonElementParentNode interface {
		ElementById(string) Element
		// non standart helper method, eliminites type asserts to HTMLElement
		HTMLElementById(string) HTMLElement
	}

	// https://dom.spec.whatwg.org/#documentorshadowroot
	DocumentOrShadowRoot interface {
		FullscreenElement() Element
		StyleSheets() []CSSStyleSheet
	}

	// https://dom.spec.whatwg.org/#parentnode
	ParentNode interface {
		Children() []Element
		FirstElementChild() Element
		LastElementChild() Element
		ChildElementCount() int
		Prepend(...interface{})
		Append(...interface{})
		QuerySelector(string) Element
		QuerySelectorAll(string) []Node
	}

	// https://dom.spec.whatwg.org/#nondocumenttypechildnode
	NonDocumentTypeChildNode interface {
		PreviousElementSibling() Element
		NextElementSibling() Element
	}

	// https://dom.spec.whatwg.org/#childnode
	ChildNode interface {
		Before(...interface{})
		After(...interface{})
		ReplaceWith(...interface{})
		Remove()
	}

	// https://dom.spec.whatwg.org/#slotable
	Slotable interface {
		AssignedSlot() HTMLSlotElement
	}
)

// -------------8<---------------------------------------

var _ NonElementParentNode = &nonElementParentNodeImpl{}

type nonElementParentNodeImpl struct {
	Value
}

func newNonElementParentNodeImpl(v Value) *nonElementParentNodeImpl {
	if v.Valid() {
		return &nonElementParentNodeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nonElementParentNodeImpl) ElementById(id string) Element {
	return wrapAsElement(p.Call("getElementById", id))
}

func (p *nonElementParentNodeImpl) HTMLElementById(id string) HTMLElement {
	return wrapAsHTMLElement(p.Call("getElementById", id))
}

// -------------8<---------------------------------------

var _ NonDocumentTypeChildNode = &nonDocumentTypeChildNodeImpl{}

type nonDocumentTypeChildNodeImpl struct {
	Value
}

func newNonDocumentTypeChildNodeImpl(v Value) *nonDocumentTypeChildNodeImpl {
	if v.Valid() {
		return &nonDocumentTypeChildNodeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nonDocumentTypeChildNodeImpl) PreviousElementSibling() Element {
	return wrapAsElement(p.Get("previousElementSibling"))
}

func (p *nonDocumentTypeChildNodeImpl) NextElementSibling() Element {
	return wrapAsElement(p.Get("nextElementSibling"))
}

// -------------8<---------------------------------------

type childNodeImpl struct {
	Value
}

func wrapChildNode(v Value) ChildNode {
	if p := newChildNodeImpl(v); p != nil {
		return p
	}
	return nil
}

func newChildNodeImpl(v Value) *childNodeImpl {
	if v.Valid() {
		return &childNodeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *childNodeImpl) Before(nodes ...interface{}) {
	var params []interface{}
	for _, node := range nodes {
		switch x := node.(type) {
		case Node:
			params = append(params, JSValue(x))
		case string:
			params = append(params, x)
		}
	}
	if len(params) > 0 {
		p.Call("before", params...)
	}
}

func (p *childNodeImpl) After(nodes ...interface{}) {
	var params []interface{}
	for _, node := range nodes {
		switch x := node.(type) {
		case Node:
			params = append(params, JSValue(x))
		case string:
			params = append(params, x)
		}
	}

	if len(params) > 0 {
		p.Call("after", params...)
	}
}

func (p *childNodeImpl) ReplaceWith(nodes ...interface{}) {
	var params []interface{}
	for _, node := range nodes {
		switch x := node.(type) {
		case Node:
			params = append(params, JSValue(x))
		case string:
			params = append(params, x)
		}
	}

	if len(params) > 0 {
		p.Call("replaceWith", params...)
	}
}

func (p *childNodeImpl) Remove() {
	p.Call("remove")
}

// -------------8<---------------------------------------

var _ DocumentOrShadowRoot = &documentOrShadowRootImpl{}

type documentOrShadowRootImpl struct {
	Value
}

func newDocumentOrShadowRootImpl(v Value) *documentOrShadowRootImpl {
	if v.Valid() {
		return &documentOrShadowRootImpl{
			Value: v,
		}
	}
	return nil
}

func (p *documentOrShadowRootImpl) FullscreenElement() Element {
	return wrapAsElement(p.Get("fullscreenElement"))
}

func (p *documentOrShadowRootImpl) StyleSheets() []CSSStyleSheet {
	if list := wrapStyleSheetList(p.Get("styleSheets")); list != nil && list.Length() > 0 {
		ret := make([]CSSStyleSheet, list.Length())
		for i := 0; i < list.Length(); i++ {
			ret[i] = list.Item(i)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type slotableImpl struct {
	Value
}

/*
func wrapSlotable(v Value) Slotable {
	if p := newSlotableImpl(v); p != nil {
		return p
	}
	return nil
}
*/

func newSlotableImpl(v Value) *slotableImpl {
	if v.Valid() {
		return &slotableImpl{
			Value: v,
		}
	}
	return nil
}

func (p *slotableImpl) AssignedSlot() HTMLSlotElement {
	return wrapHTMLSlotElement(p.Get("assignedSlot"))
}

// -------------8<---------------------------------------

var _ ParentNode = &parentNodeImpl{}

type parentNodeImpl struct {
	Value
}

func newParentNodeImpl(v Value) *parentNodeImpl {
	if v.Valid() {
		return &parentNodeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *parentNodeImpl) Children() []Element {
	return htmlCollectionToElementSlice(p.Get("children"))
}

func (p *parentNodeImpl) FirstElementChild() Element {
	return wrapAsElement(p.Get("firstElementChild"))
}

func (p *parentNodeImpl) LastElementChild() Element {
	return wrapAsElement(p.Get("lastElementChild"))
}

func (p *parentNodeImpl) ChildElementCount() int {
	return p.Get("childElementCount").Int()
}

func (p *parentNodeImpl) Prepend(nodes ...interface{}) {
	var params []interface{}
	for _, n := range nodes {
		switch x := n.(type) {
		case string:
			params = append(params, x)
		case Node:
			params = append(params, JSValue(x))
		}
	}

	if len(params) > 0 {
		p.Call("prepend", params...)
	}
}

func (p *parentNodeImpl) Append(nodes ...interface{}) {
	var params []interface{}
	for _, n := range nodes {
		switch x := n.(type) {
		case string:
			params = append(params, x)
		case Node:
			params = append(params, JSValue(x))
		}
	}
	if len(params) > 0 {
		p.Call("append", params...)
	}
}

func (p *parentNodeImpl) QuerySelector(selectors string) Element {
	return wrapAsElement(p.Call("querySelector", selectors))
}

func (p *parentNodeImpl) QuerySelectorAll(selectors string) []Node {
	return nodeListToSlice(p.Call("querySelectorAll", selectors))
}

// -------------8<---------------------------------------

type FullscreenOptions struct {
	NavigationUI FullscreenNavigationUI
}

func (p FullscreenOptions) toJSObject() Value {
	o := jsObject.New()
	o.Set("navigationUI", string(p.NavigationUI))
	return o
}

// -------------8<---------------------------------------

type FullscreenNavigationUI string

const (
	FullscreenNavigationUIAuto FullscreenNavigationUI = "auto"
	FullscreenNavigationUIShow FullscreenNavigationUI = "show"
	FullscreenNavigationUIHide FullscreenNavigationUI = "hide"
)

// -------------8<---------------------------------------
