// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type documentImpl struct {
	*nodeImpl
	*nonElementParentNodeImpl
	*documentOrShadowRootImpl
	*parentNodeImpl
	*globalEventHandlersImpl
	*documentAndElementEventHandlersImpl
	js.Value
}

func newDocument(v js.Value) Document {
	if p := newDocumentImpl(v); p != nil {
		return p
	}
	return nil
}

func newDocumentImpl(v js.Value) *documentImpl {
	if isNil(v) {
		return nil
	}

	return &documentImpl{
		nodeImpl:                            newNodeImpl(v),
		nonElementParentNodeImpl:            newNonElementParentNodeImpl(v),
		documentOrShadowRootImpl:            newDocumentOrShadowRootImpl(v),
		parentNodeImpl:                      newParentNodeImpl(v),
		globalEventHandlersImpl:             newGlobalEventHandlersImpl(v),
		documentAndElementEventHandlersImpl: newDocumentAndElementEventHandlersImpl(v),
		Value:                               v,
	}
}

func (p *documentImpl) Implementation() DOMImplementation {
	return newDOMImplementation(p.Get("implementation"))
}

func (p *documentImpl) URL() string {
	return p.Get("URL").String()
}

func (p *documentImpl) DocumentURI() string {
	return p.Get("documentURI").String()
}

func (p *documentImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *documentImpl) CompatMode() string {
	return p.Get("compatMode").String()
}

func (p *documentImpl) CharacterSet() string {
	return p.Get("characterSet").String()
}

func (p *documentImpl) ContentType() string {
	return p.Get("contentType").String()
}

func (p *documentImpl) DocType() DocumentType {
	return newDocumentType(p.Get("doctype"))
}

func (p *documentImpl) DocumentElement() Element {
	return newElement(p.Get("documentElement"))
}

func (p *documentImpl) ElementsByTagName(qualifiedName string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByTagName", qualifiedName))

}

func (p *documentImpl) ElementsByTagNameNS(namespace string, localName string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByTagNameNS", namespace, localName))
}

func (p *documentImpl) ElementsByClassName(classNames string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByClassName", classNames))
}

func (p *documentImpl) CreateElement(localName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapElement(p.Call("createElement", localName))
	default:
		return wrapElement(p.Call("createElement", localName, toJSONObject(options[0])))
	}
}

func (p *documentImpl) CreateElementNS(namespace string, qualifiedName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapElement(p.Call("createElementNS", namespace, qualifiedName))
	default:
		return wrapElement(p.Call("createElementNS", namespace, qualifiedName, toJSONObject(options[0])))
	}
}

func (p *documentImpl) CreateDocumentFragment() DocumentFragment {
	return newDocumentFragment(p.Call("createDocumentFragment"))
}

func (p *documentImpl) CreateTextNode(data string) Text {
	return newText(p.Call("createTextNode", data))
}

func (p *documentImpl) CreateCDATASection(data string) CDATASection {
	return newCDATASection(p.Call("createCDATASection", data))
}

func (p *documentImpl) CreateComment(data string) Comment {
	return newComment(p.Call("createComment", data))
}

func (p *documentImpl) CreateProcessingInstruction(target string, data string) ProcessingInstruction {
	return newProcessingInstruction(p.Call("createProcessingInstruction", target, data))
}

func (p *documentImpl) ImportNode(node Node, deep ...bool) Node {
	switch len(deep) {
	case 0:
		return newNode(p.Call("importNode", node.JSValue()))
	default:
		return newNode(p.Call("importNode", node.JSValue(), deep[0]))
	}
}

func (p *documentImpl) AdoptNode(node Node) Node {
	return newNode(p.Call("adoptNode", node.JSValue()))
}

func (p *documentImpl) CreateAttribute(localName string) Attr {
	return newAttr(p.Call("adoptNode", localName))
}

func (p *documentImpl) CreateAttributeNS(namespace string, qualifiedName string) Attr {
	return newAttr(p.Call("createAttributeNS", namespace, qualifiedName))
}

func (p *documentImpl) CreateRange() Range {
	return newRange(p.Call("createRange"))
}

func (p *documentImpl) CreateNodeIterator(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) NodeIterator {
	switch len(filter) {
	case 0:
		return newNodeIterator(p.Call("createNodeIterator", node.JSValue(), uint(whatToShow)))
	default:
		return newNodeIterator(p.Call("createNodeIterator", node.JSValue(), uint(whatToShow), filter[0].JSValue()))
	}
}

func (p *documentImpl) CreateTreeWalker(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) TreeWalker {
	switch len(filter) {
	case 0:
		return newTreeWalker(p.Call("createTreeWalker", node.JSValue(), uint(whatToShow)))
	default:
		return newTreeWalker(p.Call("createTreeWalker", node.JSValue(), uint(whatToShow), filter[0].JSValue()))
	}
}

func (p *documentImpl) FullscreenEnabled() bool {
	return p.Get("fullscreenEnabled").Bool()
}

func (p *documentImpl) ExitFullscreen() Promise {
	return newPromiseImpl(p.Call("exitFullscreen"))
}

func (p *documentImpl) OnFullscreenChange(fn func(Event)) EventHandler {
	return p.On("fullscreenchange", fn)
}

func (p *documentImpl) OnFullscreenError(fn func(Event)) EventHandler {
	return p.On("fullscreenerror", fn)
}

func (p *documentImpl) Location() Location {
	return newLocation(p.Get("location"))
}

func (p *documentImpl) Domain() string {
	return p.Get("domain").String()
}

func (p *documentImpl) SetDomain(domain string) {
	p.Set("domain", domain)
}

func (p *documentImpl) Referrer() string {
	return p.Get("referrer").String()
}

func (p *documentImpl) Cookie() string {
	return p.Get("cookie").String()
}

func (p *documentImpl) SetCookie(cookie string) {
	p.Set("cookie", cookie)
}

func (p *documentImpl) LastModified() string {
	return p.Get("lastModified").String()
}

func (p *documentImpl) ReadyState() DocumentReadyState {
	return DocumentReadyState(p.Get("readyState").String())
}

/*
func (p *documentImpl) ByName(string) js.Value {

}
*/

func (p *documentImpl) Title() string {
	return p.Get("title").String()
}

func (p *documentImpl) SetTitle(title string) {
	p.Set("title", title)
}

func (p *documentImpl) Dir() string {
	return p.Get("dir").String()
}

func (p *documentImpl) SetDir(dir string) {
	p.Set("dir", dir)
}

func (p *documentImpl) Body() HTMLBodyElement {
	return newHTMLBodyElement(p.Get("body"))
}

func (p *documentImpl) SetBody(body HTMLBodyElement) {
	p.Set("body", body.JSValue())
}

func (p *documentImpl) Head() HTMLHeadElement {
	return newHTMLHeadElement(p.Get("head"))
}

func (p *documentImpl) Images() HTMLCollection {
	return newHTMLCollection(p.Get("images"))
}

func (p *documentImpl) Embeds() HTMLCollection {
	return newHTMLCollection(p.Get("embeds"))
}

func (p *documentImpl) Plugins() HTMLCollection {
	return newHTMLCollection(p.Get("plugins"))
}

func (p *documentImpl) Links() HTMLCollection {
	return newHTMLCollection(p.Get("links"))
}

func (p *documentImpl) Forms() HTMLCollection {
	return newHTMLCollection(p.Get("forms"))
}

func (p *documentImpl) Scripts() HTMLCollection {
	return newHTMLCollection(p.Get("scripts"))
}

func (p *documentImpl) ElementsByName(name string) []Node {
	return nodeListToSlice(p.Call("getElementsByName", name))
}

func (p *documentImpl) CurrentScript() HTMLOrSVGScriptElement {
	return newHTMLOrSVGScriptElement(p.Get("currentScript"))
}

func (p *documentImpl) Open(args ...string) Document {
	switch len(args) {
	case 1:
		return newDocument(p.Call("open", args[0]))
	case 2:
		return newDocument(p.Call("open", args[0], args[1]))
	default:
		return newDocument(p.Call("open"))
	}
}

func (p *documentImpl) OpenURL(url string, name string, features string, replace ...bool) WindowProxy {
	switch len(replace) {
	case 0:
		return newWindowProxy(p.Call("open", url, name, features))
	default:
		return newWindowProxy(p.Call("open", url, name, features, replace[0]))
	}
}

func (p *documentImpl) Close() {
	p.Call("close")
}

func (p *documentImpl) Write(text ...string) {
	if len(text) > 0 {
		var params []interface{}
		for _, v := range text {
			params = append(params, v)
		}
		p.Call("write", params...)
	}
}

func (p *documentImpl) WriteLn(text ...string) {
	if len(text) > 0 {
		var params []interface{}
		for _, v := range text {
			params = append(params, v)
		}
		p.Call("writeln", params...)
	}
}

func (p *documentImpl) DefaultView() WindowProxy {
	return newWindowProxy(p.Get("defaultView"))
}

func (p *documentImpl) ActiveElement() Element {
	return newElement(p.Get("activeElement"))
}

func (p *documentImpl) HasFocus() bool {
	return p.Call("hasFocus").Bool()
}

func (p *documentImpl) DesignMode() string {
	return p.Get("designMode").String()
}

func (p *documentImpl) SetDesignMode(mode string) {
	p.Set("designMode", mode)
}

func (p *documentImpl) ExecCommand(commandId string, args ...interface{}) bool {
	switch len(args) {
	case 1:
		if showUI, ok := args[0].(bool); ok {
			return p.Call("execCommand", commandId, showUI).Bool()
		}
	case 2:
		if showUI, ok := args[0].(bool); ok {
			if value, ok := args[1].(string); ok {
				return p.Call("execCommand", commandId, showUI, value).Bool()
			}
		}
	}

	return p.Call("execCommand", commandId).Bool()
}

func (p *documentImpl) QueryCommandEnabled(commandId string) bool {
	return p.Call("queryCommandEnabled", commandId).Bool()
}

func (p *documentImpl) QueryCommandIndeterm(commandId string) bool {
	return p.Call("queryCommandIndeterm", commandId).Bool()
}

func (p *documentImpl) QueryCommandState(commandId string) bool {
	return p.Call("queryCommandState", commandId).Bool()
}

func (p *documentImpl) QueryCommandSupported(commandId string) bool {
	return p.Call("queryCommandSupported", commandId).Bool()
}

func (p *documentImpl) QueryCommandValue(commandId string) string {
	return p.Call("queryCommandValue", commandId).String()
}

func (p *documentImpl) OnReadyStateChange(fn func(Event)) EventHandler {
	return On("readystatechange", fn)
}

func (p *documentImpl) ElementFromPoint(x float64, y float64) Element {
	return newElement(p.Call("elementFromPoint", x, y))
}

func (p *documentImpl) ElementsFromPoint(x float64, y float64) []Element {
	var ret []Element

	sl := arrayToSlice(p.Call("elementsFromPoint", x, y))
	if sl != nil {
		for _, v := range sl {
			ret = append(ret, newElement(v))
		}
	}

	return ret
}

func (p *documentImpl) CaretPositionFromPoint(x float64, y float64) CaretPosition {
	return newCaretPosition(p.Call("caretPositionFromPoint", x, y))
}

func (p *documentImpl) ScrollingElement() Element {
	return newElement(p.Get("scrollingElement"))
}

// -------------8<---------------------------------------

type domImplementationImpl struct {
	js.Value
}

func newDOMImplementation(v js.Value) DOMImplementation {
	if isNil(v) {
		return nil
	}

	return &domImplementationImpl{
		Value: v,
	}
}

func (p *domImplementationImpl) CreateDocumentType(qualifiedName string, publicId string, systemId string) DocumentType {
	return newDocumentType(p.Call("createDocumentType", qualifiedName, publicId, systemId))
}

func (p *domImplementationImpl) CreateDocument(namespace string, qualifiedName string, doctype ...DocumentType) XMLDocument {
	switch len(doctype) {
	case 0:
		return newXMLDocument(p.Call("createDocument", namespace, qualifiedName))
	default:
		return newXMLDocument(p.Call("createDocument", namespace, qualifiedName, doctype[0].JSValue()))
	}
}

func (p *domImplementationImpl) CreateHTMLDocument(title ...string) Document {
	if len(title) > 0 {
		return newDocument(p.Call("createHTMLDocument", title[0]))
	}
	return newDocument(p.Call("createHTMLDocument"))
}

// -------------8<---------------------------------------

type xmlDocumentImpl struct {
	*documentImpl
}

func newXMLDocument(v js.Value) XMLDocument {
	if isNil(v) {
		return nil
	}
	return &xmlDocumentImpl{
		documentImpl: newDocumentImpl(v),
	}
}

// -------------8<---------------------------------------

type treeWalkerImpl struct {
	js.Value
}

func newTreeWalker(v js.Value) TreeWalker {
	if isNil(v) {
		return nil
	}

	return &treeWalkerImpl{
		Value: v,
	}
}

func (p *treeWalkerImpl) Root() Node {
	return newNode(p.Get("root"))
}

func (p *treeWalkerImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.Get("whatToShow").Int()))
}

func (p *treeWalkerImpl) Filter() NodeFilter {
	return newNodeFilter(p.Get("filter"))
}

func (p *treeWalkerImpl) CurrentNode() Node {
	return newNode(p.Get("currentNode"))
}

func (p *treeWalkerImpl) SetCurrentNode(node Node) {
	p.Set("currentNode", node.JSValue())
}

func (p *treeWalkerImpl) ParentNode() Node {
	return newNode(p.Call("parentNode"))
}

func (p *treeWalkerImpl) FirstChild() Node {
	return newNode(p.Call("firstChild"))
}

func (p *treeWalkerImpl) LastChild() Node {
	return newNode(p.Call("lastChild"))
}

func (p *treeWalkerImpl) PreviousSibling() Node {
	return newNode(p.Call("previousSibling"))
}

func (p *treeWalkerImpl) NextSibling() Node {
	return newNode(p.Call("nextSibling"))
}

func (p *treeWalkerImpl) PreviousNode() Node {
	return newNode(p.Call("previousNode"))
}

func (p *treeWalkerImpl) NextNode() Node {
	return newNode(p.Call("nextNode"))
}

// -------------8<---------------------------------------

type nodeFilterImpl struct {
	js.Value
}

func newNodeFilter(v js.Value) NodeFilter {
	if isNil(v) {
		return nil
	}
	return &nodeFilterImpl{
		Value: v,
	}
}

func (p *nodeFilterImpl) AcceptNode(node Node) NodeFilterResult {
	return NodeFilterResult(p.Call("acceptNode", node.JSValue()).Int())
}

// -------------8<---------------------------------------

type nodeIteratorImpl struct {
	js.Value
}

func newNodeIterator(v js.Value) NodeIterator {
	if isNil(v) {
		return nil
	}
	return &nodeIteratorImpl{
		Value: v,
	}
}

func (p *nodeIteratorImpl) Root() Node {
	return newNode(p.Get("root"))
}

func (p *nodeIteratorImpl) ReferenceNode() Node {
	return newNode(p.Get("referenceNode"))
}

func (p *nodeIteratorImpl) PointerBeforeReferenceNode() bool {
	return p.Get("pointerBeforeReferenceNode").Bool()
}

func (p *nodeIteratorImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.Get("whatToShow").Int()))
}

func (p *nodeIteratorImpl) Filter() NodeFilter {
	return newNodeFilter(p.Get("filter"))
}

func (p *nodeIteratorImpl) NextNode() Node {
	return newNode(p.Call("nextNode"))
}

func (p *nodeIteratorImpl) PreviousNode() Node {
	return newNode(p.Call("previousNode"))
}

func (p *nodeIteratorImpl) Detach() {
	p.Call("detach")
}

// -------------8<---------------------------------------

type rangeImpl struct {
	*abstractRangeImpl
}

func newRange(v js.Value) Range {
	if p := newRangeImpl(v); p != nil {
		return p
	}
	return nil
}

func newRangeImpl(v js.Value) *rangeImpl {
	if isNil(v) {
		return nil
	}
	return &rangeImpl{
		abstractRangeImpl: newAbstractRangeImpl(v),
	}
}

func (p *rangeImpl) CommonAncestorContainer() Node {
	return newNode(p.Get("commonAncestorContainer"))
}

func (p *rangeImpl) SetStart(node Node, offset int) {
	p.Call("setStart", node.JSValue(), offset)
}

func (p *rangeImpl) SetEnd(node Node, offset int) {
	p.Call("setEnd", node.JSValue(), offset)
}

func (p *rangeImpl) SetStartBefore(node Node) {
	p.Call("setStartBefore", node.JSValue())
}

func (p *rangeImpl) SetStartAfter(node Node) {
	p.Call("setStartAfter", node.JSValue())
}

func (p *rangeImpl) SetEndBefore(node Node) {
	p.Call("setEndBefore", node.JSValue())
}

func (p *rangeImpl) SetEndAfter(node Node) {
	p.Call("setEndAfter", node.JSValue())
}

func (p *rangeImpl) Collapse(toStart ...bool) {
	switch len(toStart) {
	case 0:
		p.Call("collapse")
	default:
		p.Call("collapse", toStart[0])
	}
}

func (p *rangeImpl) SelectNode(node Node) {
	p.Call("selectNode", node.JSValue())
}

func (p *rangeImpl) SelectNodeContents(node Node) {
	p.Call("selectNodeContents", node.JSValue())
}

func (p *rangeImpl) CompareBoundaryPoints(how RangeCompare, source Range) int {
	return p.Call("compareBoundaryPoints", int(how), source.JSValue()).Int()
}

func (p *rangeImpl) DeleteContents() {
	p.Call("deleteContents")
}

func (p *rangeImpl) ExtractContents() DocumentFragment {
	return newDocumentFragment(p.Call("extractContents"))
}

func (p *rangeImpl) CloneContents() DocumentFragment {
	return newDocumentFragment(p.Call("cloneContents"))
}

func (p *rangeImpl) InsertNode(node Node) {
	p.Call("insertNode", node.JSValue())
}

func (p *rangeImpl) SurroundContents(newParent Node) {
	p.Call("surroundContents", newParent.JSValue())
}

func (p *rangeImpl) CloneRange() Range {
	return newRange(p.Call("cloneRange"))
}

func (p *rangeImpl) Detach() {
	p.Call("detach")
}

func (p *rangeImpl) IsPointInRange(node Node, offset int) bool {
	return p.Call("isPointInRange", node.JSValue(), offset).Bool()
}

func (p *rangeImpl) ComparePoint(node Node, offset int) int {
	return p.Call("comparePoint", node.JSValue(), offset).Int()
}

func (p *rangeImpl) IntersectsNode(node Node) bool {
	return p.Call("intersectsNode", node.JSValue()).Bool()
}

func (p *rangeImpl) ClientRects() []DOMRect {
	rects := arrayToSlice(p.Call("getClientRects"))
	if rects == nil {
		return nil
	}

	var ret []DOMRect
	for _, rect := range rects {
		ret = append(ret, newDOMRect(rect))
	}
	return ret
}

func (p *rangeImpl) BoundingClientRect() DOMRect {
	return newDOMRect(p.Call("getBoundingClientRect"))
}

func (p *rangeImpl) CreateContextualFragment(fragment string) DocumentFragment {
	return newDocumentFragmentImpl(p.Call("createContextualFragment", fragment))
}

// -------------8<---------------------------------------

type abstractRangeImpl struct {
	js.Value
}

func newAbstractRange(v js.Value) AbstractRange {
	if p := newAbstractRangeImpl(v); p != nil {
		return p
	}
	return nil
}

func newAbstractRangeImpl(v js.Value) *abstractRangeImpl {
	if isNil(v) {
		return nil
	}
	return &abstractRangeImpl{
		Value: v,
	}
}

func (p *abstractRangeImpl) StartContainer() Node {
	return newNode(p.Get("startContainer"))
}

func (p *abstractRangeImpl) StartOffset() int {
	return p.Get("startOffset").Int()
}

func (p *abstractRangeImpl) EndContainer() Node {
	return newNode(p.Get("endContainer"))
}

func (p *abstractRangeImpl) EndOffset() int {
	return p.Get("endOffset").Int()
}

func (p *abstractRangeImpl) Collapsed() bool {
	return p.Get("collapsed").Bool()
}

// -------------8<---------------------------------------

type staticRangeImpl struct {
	*abstractRangeImpl
}

func newStaticRange(v js.Value) StaticRange {
	if isNil(v) {
		return nil
	}

	return &staticRangeImpl{
		abstractRangeImpl: newAbstractRangeImpl(v),
	}
}

// -------------8<---------------------------------------

type processingInstructionImpl struct {
	*characterDataImpl
}

func newProcessingInstruction(v js.Value) ProcessingInstruction {
	if isNil(v) {
		return nil
	}

	return &processingInstructionImpl{
		characterDataImpl: newCharacterDataImpl(v),
	}
}

func (p *processingInstructionImpl) Target() string {
	return p.Get("target").String()
}

// -------------8<---------------------------------------

type commentImpl struct {
	*characterDataImpl
}

func newComment(v js.Value) Comment {
	if isNil(v) {
		return nil
	}
	return &commentImpl{
		characterDataImpl: newCharacterDataImpl(v),
	}
}

// -------------8<---------------------------------------

type cDATASectionImpl struct {
	*textImpl
}

func newCDATASection(v js.Value) CDATASection {
	if isNil(v) {
		return nil
	}
	return &cDATASectionImpl{
		textImpl: newTextImpl(v),
	}
}

// -------------8<---------------------------------------

type textImpl struct {
	*characterDataImpl
	*slotableImpl
	js.Value
}

func newText(v js.Value) Text {
	if p := newTextImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextImpl(v js.Value) *textImpl {
	if isNil(v) {
		return nil
	}

	return &textImpl{
		characterDataImpl: newCharacterDataImpl(v),
		slotableImpl:      newSlotableImpl(v),
		Value:             v,
	}
}

func (p *textImpl) SplitText(offset int) Text {
	return newText(p.Call("splitText", offset))
}

func (p *textImpl) WholeText() string {
	return p.Get("wholeText").String()
}

func (p *textImpl) Length() int {
	return p.Get("length").Int()
}

// -------------8<---------------------------------------

type characterDataImpl struct {
	*nonDocumentTypeChildNodeImpl
	*childNodeImpl
	js.Value
}

func newCharacterData(v js.Value) CharacterData {
	if p := newCharacterDataImpl(v); p != nil {
		return p
	}
	return nil
}

func newCharacterDataImpl(v js.Value) *characterDataImpl {
	if isNil(v) {
		return nil
	}

	return &characterDataImpl{
		nonDocumentTypeChildNodeImpl: newNonDocumentTypeChildNodeImpl(v),
		childNodeImpl:                newChildNodeImpl(v),
		Value:                        v,
	}
}

func (p *characterDataImpl) Data() string {
	return p.Get("data").String()
}

func (p *characterDataImpl) SetData(data string) {
	p.Set("data", data)
}

func (p *characterDataImpl) Length() int {
	return p.Get("length").Int()
}

func (p *characterDataImpl) Substring(offset int, count int) string {
	return p.Call("substringData", offset, count).String()
}

func (p *characterDataImpl) Append(data string) {
	p.Call("appendData", data)
}

func (p *characterDataImpl) Insert(offset int, data string) {
	p.Call("insertData", offset, data)
}

func (p *characterDataImpl) Delete(offset int, count int) {
	p.Call("deleteData", offset, count)
}

func (p *characterDataImpl) Replace(offset int, count int, data string) {
	p.Call("replaceData", offset, count, data)
}

// -------------8<---------------------------------------

type documentTypeImpl struct {
	*nodeImpl
	*childNodeImpl
	js.Value
}

func newDocumentType(v js.Value) DocumentType {
	if isNil(v) {
		return nil
	}

	return &documentTypeImpl{
		nodeImpl:      newNodeImpl(v),
		childNodeImpl: newChildNodeImpl(v),
		Value:         v,
	}
}

func (p *documentTypeImpl) Name() string {
	return p.Get("name").String()
}

func (p *documentTypeImpl) PublicId() string {
	return p.Get("publicId").String()
}

func (p *documentTypeImpl) SystemId() string {
	return p.Get("systemId").String()
}

// -------------8<---------------------------------------

type nodeImpl struct {
	*eventTargetImpl
}

func newNode(v js.Value) Node {
	if p := newNodeImpl(v); p != nil {
		return p
	}
	return nil
}

func newNodeImpl(v js.Value) *nodeImpl {
	if isNil(v) {
		return nil
	}
	return &nodeImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *nodeImpl) NodeType() NodeType {
	return NodeType(p.Get("nodeType").Int())
}

func (p *nodeImpl) NodeName() string {
	return p.Get("nodeName").String()
}

func (p *nodeImpl) BaseURI() string {
	return p.Get("baseURI").String()
}

func (p *nodeImpl) IsConnected() bool {
	return p.Get("isConnected").Bool()
}

func (p *nodeImpl) OwnerDocument() Document {
	return newDocument(p.Get("ownerDocument"))
}

func (p *nodeImpl) RootNode(options ...RootNodeOptions) Node {
	if len(options) > 0 {
		return newNode(p.Call("getRootNode", toJSONObject(options[0])))
	}

	return newNode(p.Call("getRootNode"))
}

func (p *nodeImpl) ParentNode() Node {
	return newNode(p.Get("parentNode"))
}

func (p *nodeImpl) ParentElement() Element {
	return newElement(p.Get("parentElement"))
}

func (p *nodeImpl) HasChildNodes() bool {
	return p.Call("hasChildNodes").Bool()
}

func (p *nodeImpl) ChildNodes() []Node {
	return nodeListToSlice(p.Get("childNodes"))
}

func (p *nodeImpl) FirstChild() Node {
	return newNode(p.Get("firstChild"))
}

func (p *nodeImpl) LastChild() Node {
	return newNode(p.Get("lastChild"))
}

func (p *nodeImpl) PreviousSibling() Node {
	return newNode(p.Get("previousSibling"))
}

func (p *nodeImpl) NextSibling() Node {
	return newNode(p.Get("nextSibling"))
}

func (p *nodeImpl) NodeValue() string {
	return p.Get("nodeValue").String()
}

func (p *nodeImpl) SetNodeValue(nval string) {
	p.Set("nodeValue", nval)
}

func (p *nodeImpl) TextContent() string {
	return p.Get("textContent").String()
}

func (p *nodeImpl) SetTextContent(tc string) {
	p.Set("textContent", tc)
}

func (p *nodeImpl) Normalize() {
	p.Call("normalize")
}

func (p *nodeImpl) CloneNode(deep ...bool) Node {
	if len(deep) > 0 {
		return newNode(p.Call("cloneNode", deep[0]))
	}
	return newNode(p.Call("cloneNode"))
}

func (p *nodeImpl) IsEqualNode(otherNode Node) bool {
	return p.Call("isEqualNode", otherNode.JSValue()).Bool()
}

func (p *nodeImpl) IsSameNode(otherNode Node) bool {
	return p.Call("isSameNode", otherNode.JSValue()).Bool()
}

func (p *nodeImpl) CompareDocumentPosition(other Node) DocumentPosition {
	return DocumentPosition(p.Call("compareDocumentPosition", other.JSValue()).Int())
}

func (p *nodeImpl) Contains(other Node) bool {
	return p.Call("contains", other.JSValue()).Bool()
}

func (p *nodeImpl) LookupPrefix(namespace string) string {
	return p.Call("lookupPrefix", namespace).String()
}

func (p *nodeImpl) LookupNamespaceURI(prefix string) string {
	return p.Call("lookupNamespaceURI", prefix).String()
}

func (p *nodeImpl) IsDefaultNamespace(namespace string) bool {
	return p.Call("isDefaultNamespace", namespace).Bool()
}

func (p *nodeImpl) InsertBefore(node Node, child Node) Node {
	if child != nil {
		return newNode(p.Call("insertBefore", node.JSValue(), child.JSValue()))
	}
	return newNode(p.Call("insertBefore", node.JSValue()))
}

func (p *nodeImpl) AppendChild(node Node) Node {
	return newNode(p.Call("appendChild", node.JSValue()))
}

func (p *nodeImpl) ReplaceChild(node Node, child Node) Node {
	return newNode(p.Call("replaceChild", node.JSValue(), child.JSValue()))
}

func (p *nodeImpl) RemoveChild(child Node) Node {
	return newNode(p.Call("removeChild", child.JSValue()))
}

// -------------8<---------------------------------------

type elementImpl struct {
	*nodeImpl
	*parentNodeImpl
	*nonDocumentTypeChildNodeImpl
	*childNodeImpl
	*slotableImpl
	js.Value
}

func newElement(v js.Value) Element {
	if p := newElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newElementImpl(v js.Value) *elementImpl {
	if isNil(v) {
		return nil
	}

	return &elementImpl{
		nodeImpl:                     newNodeImpl(v),
		parentNodeImpl:               newParentNodeImpl(v),
		nonDocumentTypeChildNodeImpl: newNonDocumentTypeChildNodeImpl(v),
		childNodeImpl:                newChildNodeImpl(v),
		slotableImpl:                 newSlotableImpl(v),
		Value:                        v,
	}
}

func (p *elementImpl) NamespaceURI() string {
	return p.Get("namespaceURI").String()
}

func (p *elementImpl) Prefix() string {
	return p.Get("prefix").String()
}

func (p *elementImpl) LocalName() string {
	return p.Get("localName").String()
}

func (p *elementImpl) TagName() string {
	return p.Get("tagName").String()
}

func (p *elementImpl) Id() string {
	return p.Get("id").String()
}

func (p *elementImpl) SetId(id string) {
	p.Set("id", id)
}

func (p *elementImpl) ClassName() string {
	return p.Get("className").String()
}

func (p *elementImpl) SetClassName(name string) {
	p.Set("className", name)
}

func (p *elementImpl) ClassList() DOMTokenList {
	return newDOMTokenList(p.Get("classList"))
}

func (p *elementImpl) Slot() string {
	return p.Get("slot").String()
}

func (p *elementImpl) SetSlot(slot string) {
	p.Set("slot", slot)
}

func (p *elementImpl) HasAttributes() bool {
	return p.Call("hasAttributes").Bool()
}

func (p *elementImpl) Attributes() NamedNodeMap {
	return newNamedNodeMap(p.Get("attributes"))
}

func (p *elementImpl) AttributeNames() []string {
	return stringSequenceToSlice(p.Call("getAttributeNames"))
}

func (p *elementImpl) Attribute(name string) string {
	return p.Call("getAttribute", name).String()
}

func (p *elementImpl) AttributeNS(namespace string, localName string) string {
	return p.Call("getAttributeNS", namespace, localName).String()
}

func (p *elementImpl) SetAttribute(name string, value string) {
	p.Call("setAttribute", name, value)
}

func (p *elementImpl) SetAttributeNS(namespace string, name string, value string) {
	p.Call("setAttributeNS", namespace, name, value)
}

func (p *elementImpl) RemoveAttribute(name string) {
	p.Call("removeAttribute", name)
}

func (p *elementImpl) RemoveAttributeNS(namespace string, name string) {
	p.Call("removeAttributeNS", namespace, name)
}

func (p *elementImpl) ToggleAttribute(name string, force ...bool) bool {
	if len(force) > 0 {
		return p.Call("toggleAttribute", name, force[0]).Bool()
	}
	return p.Call("toggleAttribute", name).Bool()
}

func (p *elementImpl) HasAttribute(name string) bool {
	return p.Call("hasAttribute", name).Bool()
}

func (p *elementImpl) HasAttributeNS(namespace string, localName string) bool {
	return p.Call("hasAttributeNS", namespace, localName).Bool()
}

func (p *elementImpl) AttributeNode(name string) Attr {
	return newAttr(p.Call("getAttributeNode", name))
}

func (p *elementImpl) AttributeNodeNS(namespace string, name string) Attr {
	return newAttr(p.Call("getAttributeNodeNS", namespace, name))
}

func (p *elementImpl) SetAttributeNode(attr Attr) Attr {
	return newAttr(p.Call("setAttributeNode", attr.JSValue()))
}

func (p *elementImpl) SetAttributeNodeNS(attr Attr) Attr {
	return newAttr(p.Call("setAttributeNodeNS", attr.JSValue()))
}

func (p *elementImpl) RemoveAttributeNode(attr Attr) Attr {
	return newAttr(p.Call("removeAttributeNode", attr.JSValue()))
}

func (p *elementImpl) AttachShadow(si ShadowRootInit) ShadowRoot {
	return newShadowRoot(p.Call("attachShadow", toJSONObject(si)))
}

func (p *elementImpl) ShadowRoot() ShadowRoot {
	return newShadowRoot(p.Get("shadowRoot"))
}

func (p *elementImpl) Closest(selectors string) Element {
	return newElement(p.Call("closest"))
}

func (p *elementImpl) Matches(string) bool {
	return p.Call("matches").Bool()
}

func (p *elementImpl) ElementsByTagName(name string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByTagName", name))
}

func (p *elementImpl) ElementsByTagNameNS(namespace string, localName string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByTagNameNS", namespace, localName))
}

func (p *elementImpl) ElementsByClassName(names string) HTMLCollection {
	return newHTMLCollection(p.Call("getElementsByClassName", names))
}

func (p *elementImpl) ClientRects() []DOMRect {
	rects := arrayToSlice(p.Call("getClientRects"))
	if rects == nil {
		return nil
	}

	var ret []DOMRect
	for _, rect := range rects {
		ret = append(ret, newDOMRect(rect))
	}
	return ret
}

func (p *elementImpl) BoundingClientRect() DOMRect {
	return newDOMRect(p.Call("getBoundingClientRect"))
}

func (p *elementImpl) ScrollIntoView(arg ...interface{}) {
	switch len(arg) {
	case 0:
		p.Call("scrollIntoView")
	default:
		switch x := arg[0].(type) {
		case bool:
			p.Call("scrollIntoView", x)
		case ScrollIntoViewOptions:
			p.Call("scrollIntoView", toJSONObject(x))
		}
	}
}

func (p *elementImpl) Scroll(options ScrollToOptions) {
	p.Call("scroll", toJSONObject(options))
}

func (p *elementImpl) ScrollTo(options ScrollToOptions) {
	p.Call("scrollTo", toJSONObject(options))
}

func (p *elementImpl) ScrollBy(options ScrollToOptions) {
	p.Call("scrollBy", toJSONObject(options))
}

func (p *elementImpl) ScrollTop() float64 {
	return p.Get("scrollTop").Float()
}

func (p *elementImpl) SetScrollTop(st float64) {
	p.Set("scrollTop", st)
}

func (p *elementImpl) ScrollLeft() float64 {
	return p.Get("scrollLeft").Float()
}

func (p *elementImpl) SetScrollLeft(sl float64) {
	p.Set("scrollLeft", sl)
}

func (p *elementImpl) ScrollWidth() int {
	return p.Get("scrollWidth").Int()
}

func (p *elementImpl) ScrollHeight() int {
	return p.Get("scrollHeight").Int()
}

func (p *elementImpl) ClientTop() int {
	return p.Get("clientTop").Int()
}

func (p *elementImpl) ClientLeft() int {
	return p.Get("clientLeft").Int()
}

func (p *elementImpl) ClientWidth() int {
	return p.Get("clientWidth").Int()
}

func (p *elementImpl) ClientHeight() int {
	return p.Get("clientHeight").Int()
}

func (p *elementImpl) OnFullscreenChange(fn func(Event)) EventHandler {
	return p.On("fullscreenchange", fn)
}

func (p *elementImpl) OnFullscreenError(fn func(Event)) EventHandler {
	return p.On("fullscreenerror", fn)
}

func (p *elementImpl) InnerHTML() string {
	return p.Get("innerHTML").String()
}

func (p *elementImpl) SetInnerHTML(html string) {
	p.Set("innerHTML", html)
}

func (p *elementImpl) OuterHTML() string {
	return p.Get("outerHTML").String()
}

func (p *elementImpl) SetOuterHTML(html string) {
	p.Set("outerHTML", html)
}

func (p *elementImpl) InsertAdjacentHTML(position string, text string) {
	p.Call("insertAdjacentHTML", position, text)
}

func (p *elementImpl) RequestFullscreen(...FullscreenOptions) Promise {
	return newPromiseImpl(p.Call("requestFullscreen"))
}

func (p *elementImpl) OnFullScreenChange(fn func(Event)) EventHandler {
	return p.On("fullscreenchange", fn)
}

func (p *elementImpl) OnFullScreenError(fn func(Event)) EventHandler {
	return p.On("fullscreenerror", fn)
}

// -------------8<---------------------------------------

type shadowRootImpl struct {
	*documentFragmentImpl
	*documentOrShadowRootImpl
	*parentNodeImpl
	js.Value
}

func newShadowRoot(v js.Value) ShadowRoot {
	if isNil(v) {
		return nil
	}
	return &shadowRootImpl{
		documentFragmentImpl:     newDocumentFragmentImpl(v),
		documentOrShadowRootImpl: newDocumentOrShadowRootImpl(v),
		parentNodeImpl:           newParentNodeImpl(v),
		Value:                    v,
	}
}

func (p *shadowRootImpl) Mode() ShadowRootMode {
	return ShadowRootMode(p.Get("mode").String())
}

func (p *shadowRootImpl) Host() Element {
	return newElement(p.Get("host"))
}

// -------------8<---------------------------------------

type documentFragmentImpl struct {
	*nodeImpl
	*nonElementParentNodeImpl
}

func newDocumentFragment(v js.Value) DocumentFragment {
	if p := newDocumentFragmentImpl(v); p != nil {
		return p
	}
	return nil
}

func newDocumentFragmentImpl(v js.Value) *documentFragmentImpl {
	if isNil(v) {
		return nil
	}
	return &documentFragmentImpl{
		nodeImpl:                 newNodeImpl(v),
		nonElementParentNodeImpl: newNonElementParentNodeImpl(v),
	}
}

// -------------8<---------------------------------------

type domTokenListImpl struct {
	js.Value
}

func newDOMTokenList(v js.Value) DOMTokenList {
	if isNil(v) {
		return nil
	}
	return &domTokenListImpl{
		Value: v,
	}
}

func (p *domTokenListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *domTokenListImpl) Item(index int) string {
	return p.Get("item").String()
}

func (p *domTokenListImpl) Contains(token string) bool {
	return p.Call("contains", token).Bool()
}

func (p *domTokenListImpl) Add(tokens ...string) {
	if len(tokens) > 0 {
		var params []interface{}
		for _, v := range tokens {
			params = append(params, v)
		}
		p.Call("add", params...)
	}
}

func (p *domTokenListImpl) Remove(tokens ...string) {
	if len(tokens) > 0 {
		var params []interface{}
		for _, v := range tokens {
			params = append(params, v)
		}
		p.Call("remove", params...)
	}
}

func (p *domTokenListImpl) Toggle(token string, force ...bool) bool {
	if len(force) > 0 {
		return p.Call("toggle", token, force[0]).Bool()
	}
	return p.Call("toggle", token).Bool()
}

func (p *domTokenListImpl) Replace(token string, newToken string) bool {
	return p.Call("replace", token, newToken).Bool()
}

func (p *domTokenListImpl) Supports(token string) bool {
	return p.Call("supports", token).Bool()
}

func (p *domTokenListImpl) TokenValue() string {
	return p.Get("value").String()
}

func (p *domTokenListImpl) SetTokenValue(value string) {
	p.Set("value", value)
}

func (p *domTokenListImpl) TokenValues() []string {
	var ret []string
	it := p.Call("values")
	for {
		n := it.Call("next")
		if n.Get("done").Bool() {
			break
		}

		ret = append(ret, n.Get("value").String())
	}
	return ret
}

func (p *domTokenListImpl) String() string {
	return p.Call("toString").String()
}

// -------------8<---------------------------------------

type namedNodeMapImpl struct {
	js.Value
}

func newNamedNodeMap(v js.Value) NamedNodeMap {
	if isNil(v) {
		return nil
	}
	return &namedNodeMapImpl{
		Value: v,
	}
}

func (p *namedNodeMapImpl) Length() int {
	return p.Get("length").Int()
}

func (p *namedNodeMapImpl) Item(index int) Attr {
	return newAttr(p.Call("item", index))
}

func (p *namedNodeMapImpl) NamedItem(name string) Attr {
	return newAttr(p.Get("getNamedItem"))
}

func (p *namedNodeMapImpl) NamedItemNS(namespace string, name string) Attr {
	return newAttr(p.Call("getNamedItemNS", namespace, name))
}

func (p *namedNodeMapImpl) SetNamedItem(attr Attr) Attr {
	return newAttr(p.Call("setNamedItem", attr.JSValue()))
}

func (p *namedNodeMapImpl) SetNamedItemNS(attr Attr) Attr {
	return newAttr(p.Call("setNamedItemNS", attr.JSValue()))
}

func (p *namedNodeMapImpl) RemoveNamedItem(name string) Attr {
	return newAttr(p.Call("removeNamedItem", name))
}

func (p *namedNodeMapImpl) RemoveNamedItemNS(namespace string, name string) Attr {
	return newAttr(p.Call("removeNamedItemNS", namespace, name))
}

// -------------8<---------------------------------------

type attrImpl struct {
	*nodeImpl
}

func newAttr(v js.Value) Attr {
	if isNil(v) {
		return nil
	}
	return &attrImpl{
		nodeImpl: newNodeImpl(v),
	}
}

func (p *attrImpl) NamespaceURI() string {
	return p.Get("namespaceURI").String()
}

func (p *attrImpl) Prefix() string {
	return p.Get("prefix").String()
}

func (p *attrImpl) LocalName() string {
	return p.Get("localName").String()
}

func (p *attrImpl) Name() string {
	return p.Get("name").String()
}

func (p *attrImpl) Value() string {
	return p.Get("value").String()
}

func (p *attrImpl) SetValue(value string) {
	p.Set("value", value)
}

func (p *attrImpl) OwnerElement() Element {
	return newElement(p.Get("ownerElement"))
}

// -------------8<---------------------------------------

type htmlCollectionImpl struct {
	js.Value
}

func newHTMLCollection(v js.Value) HTMLCollection {
	if p := newHTMLCollectionImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLCollectionImpl(v js.Value) *htmlCollectionImpl {
	if isNil(v) {
		return nil
	}
	return &htmlCollectionImpl{
		Value: v,
	}
}

func (p *htmlCollectionImpl) Length() int {
	return p.Get("length").Int()
}

func (p *htmlCollectionImpl) Item(index int) Element {
	return newElement(p.Call("item", index))
}

func (p *htmlCollectionImpl) NamedItem(name string) Element {
	return newElement(p.Call("namedItem", name))
}

// -------------8<---------------------------------------

type mutationRecordImpl struct {
	js.Value
}

func newMutationRecord(v js.Value) MutationRecord {
	if isNil(v) {
		return nil
	}
	return &mutationRecordImpl{
		Value: v,
	}
}

func (p *mutationRecordImpl) Type() string {
	return p.Get("type").String()
}

func (p *mutationRecordImpl) Target() Node {
	return newNode(p.Get("target"))
}

func (p *mutationRecordImpl) AddedNodes() []Node {
	return nodeListToSlice(p.Get("addedNodes"))
}

func (p *mutationRecordImpl) RemovedNodes() []Node {
	return nodeListToSlice(p.Get("removedNodes"))
}

func (p *mutationRecordImpl) PreviousSibling() Node {
	return newNode(p.Get("previousSibling"))
}

func (p *mutationRecordImpl) NextSibling() Node {
	return newNode(p.Get("nextSibling"))
}

func (p *mutationRecordImpl) AttributeName() string {
	return p.Get("attributeName").String()
}

func (p *mutationRecordImpl) AttributeNamespace() string {
	return p.Get("attributeNamespace").String()
}

func (p *mutationRecordImpl) OldValue() string {
	return p.Get("oldValue").String()
}

// -------------8<---------------------------------------

type mutationObserverImpl struct {
	js.Value
}

func newMutationObserver(v js.Value) MutationObserver {
	if isNil(v) {
		return nil
	}
	return &mutationObserverImpl{
		Value: v,
	}
}

func (p *mutationObserverImpl) Observe(target Node, options ...MutationObserverInit) {
	switch len(options) {
	case 0:
		p.Call("observe", target.JSValue())
	default:
		p.Call("observe", target.JSValue(), toJSONObject(options[0]))
	}
}

func (p *mutationObserverImpl) Disconnect() {
	p.Call("disconnect")
}

func (p *mutationObserverImpl) TakeRecords() []MutationRecord {
	s := arrayToSlice(p.Call("takeRecords"))
	if s == nil {
		return nil
	}

	var ret []MutationRecord
	for _, v := range s {
		ret = append(ret, newMutationRecord(v))
	}
	return ret
}

// -------------8<---------------------------------------

type htmlUnknownElementImpl struct {
	*htmlElementImpl
}

func newHTMLUnknownElement(v js.Value) HTMLUnknownElement {
	if isNil(v) {
		return nil
	}
	return &htmlUnknownElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlElementImpl struct {
	*elementImpl
	*globalEventHandlersImpl
	*documentAndElementEventHandlersImpl
	js.Value
}

func newHTMLElement(v js.Value) HTMLElement {
	if p := newHTMLElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLElementImpl(v js.Value) *htmlElementImpl {
	if isNil(v) {
		return nil
	}
	return &htmlElementImpl{
		elementImpl:                         newElementImpl(v),
		globalEventHandlersImpl:             newGlobalEventHandlersImpl(v),
		documentAndElementEventHandlersImpl: newDocumentAndElementEventHandlersImpl(v),
		Value:                               v,
	}
}

func (p *htmlElementImpl) Title() string {
	return p.Get("title").String()
}

func (p *htmlElementImpl) SetTitle(title string) {
	p.Set("title", title)
}

func (p *htmlElementImpl) Lang() string {
	return p.Get("lang").String()
}

func (p *htmlElementImpl) SetLang(lang string) {
	p.Set("lang", lang)
}

func (p *htmlElementImpl) Translate() bool {
	return p.Get("translate").Bool()
}

func (p *htmlElementImpl) SetTranslate(tr bool) {
	p.Set("translate", tr)
}

func (p *htmlElementImpl) Dir() string {
	return p.Get("dir").String()
}

func (p *htmlElementImpl) SetDir(dir string) {
	p.Set("dir", dir)
}

func (p *htmlElementImpl) DataSet() DOMStringMap {
	return newDOMStringMap(p.Get("dataset"))
}

func (p *htmlElementImpl) Hidden() bool {
	return p.Get("hidden").Bool()
}

func (p *htmlElementImpl) SetHidden(hidden bool) {
	p.Set("hidden", hidden)
}

func (p *htmlElementImpl) Click() {
	p.Call("click")
}

func (p *htmlElementImpl) TabIndex() int {
	return p.Get("tabIndex").Int()
}

func (p *htmlElementImpl) SetTabIndex(index int) {
	p.Set("tabIndex", index)
}

func (p *htmlElementImpl) Focus() {
	p.Call("focus")
}

func (p *htmlElementImpl) Blur() {
	p.Call("blur")
}

func (p *htmlElementImpl) AccessKey() string {
	return p.Get("accessKey").String()
}

func (p *htmlElementImpl) SetAccessKey(key string) {
	p.Set("accessKey", key)
}

func (p *htmlElementImpl) AccessKeyLabel() string {
	return p.Get("accessKeyLabel").String()
}

func (p *htmlElementImpl) Draggable() bool {
	return p.Get("draggable").Bool()
}

func (p *htmlElementImpl) SetDraggable(d bool) {
	p.Set("draggable", d)
}

func (p *htmlElementImpl) SpellCheck() bool {
	return p.Get("spellcheck").Bool()
}

func (p *htmlElementImpl) SetSpellChack(s bool) {
	p.Set("spellcheck", s)
}

func (p *htmlElementImpl) ForceSpellCheck() {
	p.Call("forceSpellCheck")
}

/*
func (p *htmlElementImpl) Autocapitalize() string {

}


func (p *htmlElementImpl) SetAutocapitalize(string) {

}
*/

func (p *htmlElementImpl) InnerText() string {
	return p.Get("innerText").String()
}

func (p *htmlElementImpl) SetInnerText(text string) {
	p.Set("innerText", text)
}

func (p *htmlElementImpl) OffsetParent() Element {
	return newElement(p.Get("offsetParent"))
}

func (p *htmlElementImpl) OffsetTop() int {
	return p.Get("offsetTop").Int()
}

func (p *htmlElementImpl) OffsetLeft() int {
	return p.Get("offsetLeft").Int()
}

func (p *htmlElementImpl) OffsetWidth() int {
	return p.Get("offsetWidth").Int()
}

func (p *htmlElementImpl) OffsetHeight() int {
	return p.Get("offsetHeight").Int()
}

// -------------8<---------------------------------------

type domStringMapImpl struct {
	js.Value
}

func newDOMStringMap(v js.Value) DOMStringMap {
	if isNil(v) {
		return nil
	}

	return &domStringMapImpl{
		Value: v,
	}
}

func (p *domStringMapImpl) Get(name string) string {
	return p.Call("getDataAttr", name).String()
}

func (p *domStringMapImpl) Set(name string, value string) {
	p.Call("setDataAttr", name, value)
}

func (p *domStringMapImpl) Delete(name string) {
	p.Call("removeDataAttr", name)
}

// -------------8<---------------------------------------

type htmlOrSVGScriptElementImpl struct {
	js.Value
}

func newHTMLOrSVGScriptElement(v js.Value) HTMLOrSVGScriptElement {
	if isNil(v) {
		return nil
	}

	return &htmlOrSVGScriptElementImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type nodeListImpl struct {
	js.Value
}

func newNodeList(v js.Value) NodeList {
	if p := newNodeListImpl(v); p != nil {
		return p
	}
	return nil
}

func newNodeListImpl(v js.Value) *nodeListImpl {
	if isNil(v) {
		return nil
	}

	return &nodeListImpl{
		Value: v,
	}
}

func (p *nodeListImpl) Item(index int) Node {
	return newNode(p.Call("item", index))
}

func (p *nodeListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *nodeListImpl) Items() []Node {
	return nodeListToSlice(p.Call("entries"))
}

// -------------8<---------------------------------------

type domParserImpl struct {
	js.Value
}

func newDOMParserImpl(v js.Value) DOMParser {
	if isNil(v) {
		return nil
	}

	return &domParserImpl{
		Value: v,
	}
}

func (p *domParserImpl) ParseFromString(str string, typ SupportedType) Document {
	return newDocumentImpl(p.Call("parseFromString", str, string(typ)))
}

// -------------8<---------------------------------------

type xmlSerializerImpl struct {
	js.Value
}

func newXMLSerializerImpl(v js.Value) XMLSerializer {
	if isNil(v) {
		return nil
	}
	return &xmlSerializerImpl{
		Value: v,
	}
}

func (p *xmlSerializerImpl) SerializeToString(node Node) string {
	return p.Call("serializeToString", node.JSValue()).String()
}

// -------------8<---------------------------------------

func NewMutationObserver(cb MutationCallback) MutationObserver {
	jsMutationObserver := js.Global().Get("MutationObserver")
	if isNil(jsMutationObserver) {
		return nil
	}

	return newMutationObserver(jsMutationObserver.New(cb.jsCallback()))
}

// -------------8<---------------------------------------

func NewDOMParser() DOMParser {
	v := js.Global().Get("DOMParser")
	if isNil(v) {
		return nil
	}

	return newDOMParserImpl(v)
}

// -------------8<---------------------------------------

func NewXMLSerializer() XMLSerializer {
	v := js.Global().Get("XMLSerializer")
	if isNil(v) {
		return nil
	}

	return newXMLSerializerImpl(v)
}

// -------------8<---------------------------------------
