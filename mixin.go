// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://dom.spec.whatwg.org/#nonelementparentnode
	NonElementParentNode interface {
		ElementById(string) Element
	}

	// https://dom.spec.whatwg.org/#documentorshadowroot
	DocumentOrShadowRoot interface {
		FullscreenElement() Element
	}

	// https://dom.spec.whatwg.org/#parentnode
	ParentNode interface {
		Children() HTMLCollection
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
	js.Value
}

func newNonElementParentNodeImpl(v js.Value) *nonElementParentNodeImpl {
	if isNil(v) {
		return nil
	}

	return &nonElementParentNodeImpl{
		Value: v,
	}
}

func (p *nonElementParentNodeImpl) ElementById(id string) Element {
	return newElement(p.Call("getElementById", id))
}

// -------------8<---------------------------------------

var _ NonDocumentTypeChildNode = &nonDocumentTypeChildNodeImpl{}

type nonDocumentTypeChildNodeImpl struct {
	js.Value
}

func newNonDocumentTypeChildNodeImpl(v js.Value) *nonDocumentTypeChildNodeImpl {
	if isNil(v) {
		return nil
	}

	return &nonDocumentTypeChildNodeImpl{
		Value: v,
	}
}

func (p *nonDocumentTypeChildNodeImpl) PreviousElementSibling() Element {
	return newElement(p.Get("previousElementSibling"))
}

func (p *nonDocumentTypeChildNodeImpl) NextElementSibling() Element {
	return newElement(p.Get("nextElementSibling"))
}

// -------------8<---------------------------------------

type childNodeImpl struct {
	js.Value
}

func newChildNode(v js.Value) ChildNode {
	if p := newChildNodeImpl(v); p != nil {
		return p
	}
	return nil
}

func newChildNodeImpl(v js.Value) *childNodeImpl {
	if isNil(v) {
		return nil
	}

	return &childNodeImpl{
		Value: v,
	}
}

func (p *childNodeImpl) Before(nodes ...interface{}) {
	var params []interface{}
	for _, node := range nodes {
		switch x := node.(type) {
		case Node:
			params = append(params, x.JSValue())
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
			params = append(params, x.JSValue())
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
			params = append(params, x.JSValue())
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
	js.Value
}

func newDocumentOrShadowRootImpl(v js.Value) *documentOrShadowRootImpl {
	if isNil(v) {
		return nil
	}

	return &documentOrShadowRootImpl{
		Value: v,
	}
}

func (p *documentOrShadowRootImpl) FullscreenElement() Element {
	return newElement(p.Get("fullscreenElement"))
}

// -------------8<---------------------------------------

type slotableImpl struct {
	js.Value
}

/*
func newSlotable(v js.Value) Slotable {
	if p := newSlotableImpl(v); p != nil {
		return p
	}
	return nil
}
*/

func newSlotableImpl(v js.Value) *slotableImpl {
	if isNil(v) {
		return nil
	}

	return &slotableImpl{
		Value: v,
	}
}

func (p *slotableImpl) AssignedSlot() HTMLSlotElement {
	return newHTMLSlotElement(p.Get("assignedSlot"))
}

// -------------8<---------------------------------------

var _ ParentNode = &parentNodeImpl{}

type parentNodeImpl struct {
	js.Value
}

func newParentNodeImpl(v js.Value) *parentNodeImpl {
	if isNil(v) {
		return nil
	}
	return &parentNodeImpl{
		Value: v,
	}
}

func (p *parentNodeImpl) Children() HTMLCollection {
	return newHTMLCollection(p.Get("children"))
}

func (p *parentNodeImpl) FirstElementChild() Element {
	return newElement(p.Get("firstElementChild"))
}

func (p *parentNodeImpl) LastElementChild() Element {
	return newElement(p.Get("lastElementChild"))
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
			params = append(params, x.JSValue())
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
			params = append(params, x.JSValue())
		}
	}
	if len(params) > 0 {
		p.Call("append", params...)
	}
}

func (p *parentNodeImpl) QuerySelector(selectors string) Element {
	return newElement(p.Call("querySelector", selectors))
}

func (p *parentNodeImpl) QuerySelectorAll(selectors string) []Node {
	return nodeListToSlice(p.Call("querySelectorAll", selectors))
}

// -------------8<---------------------------------------

type FullscreenOptions struct {
	NavigationUI FullscreenNavigationUI
}

func (p FullscreenOptions) toDict() js.Value {
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
