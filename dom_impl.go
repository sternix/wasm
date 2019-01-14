// +build js,wasm

package wasm

// -------------8<---------------------------------------

type documentImpl struct {
	*eventTargetImpl
	*nodeImpl
	*nonElementParentNodeImpl
	*documentOrShadowRootImpl
	*parentNodeImpl
	*globalEventHandlersImpl
	*documentAndElementEventHandlersImpl
	*geometryUtilsImpl
	Value
}

func NewDocument() Document {
	if jsDoc := jsGlobal.get("Document"); jsDoc.valid() {
		return wrapDocument(jsDoc.jsNew())
	}
	return nil
}

func wrapDocument(v Value) Document {
	if p := newDocumentImpl(v); p != nil {
		return p
	}
	return nil
}

func newDocumentImpl(v Value) *documentImpl {
	if v.valid() {
		di := &documentImpl{
			nodeImpl:                 newNodeImpl(v),
			nonElementParentNodeImpl: newNonElementParentNodeImpl(v),
			documentOrShadowRootImpl: newDocumentOrShadowRootImpl(v),
			parentNodeImpl:           newParentNodeImpl(v),
			geometryUtilsImpl:        newGeometryUtilsImpl(v),
			Value:                    v,
		}
		di.eventTargetImpl = di.nodeImpl.eventTargetImpl
		di.globalEventHandlersImpl = newGlobalEventHandlersImpl(di.eventTargetImpl)
		di.documentAndElementEventHandlersImpl = newDocumentAndElementEventHandlersImpl(di.eventTargetImpl)
		return di
	}
	return nil
}

func (p *documentImpl) Implementation() DOMImplementation {
	return wrapDOMImplementation(p.get("implementation"))
}

func (p *documentImpl) URL() string {
	return p.get("URL").toString()
}

func (p *documentImpl) DocumentURI() string {
	return p.get("documentURI").toString()
}

func (p *documentImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *documentImpl) CompatMode() string {
	return p.get("compatMode").toString()
}

func (p *documentImpl) CharacterSet() string {
	return p.get("characterSet").toString()
}

func (p *documentImpl) ContentType() string {
	return p.get("contentType").toString()
}

func (p *documentImpl) DocType() DocumentType {
	return wrapDocumentType(p.get("doctype"))
}

func (p *documentImpl) DocumentElement() Element {
	return wrapAsElement(p.get("documentElement"))
}

func (p *documentImpl) ElementsByTagName(qualifiedName string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByTagName", qualifiedName))
}

func (p *documentImpl) ElementsByTagNameNS(namespace string, localName string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByTagNameNS", namespace, localName))
}

func (p *documentImpl) ElementsByClassName(classNames string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByClassName", classNames))
}

func (p *documentImpl) HTMLElementsByClassName(classNames string) []HTMLElement {
	return htmlCollectionToHTMLElementSlice(p.call("getElementsByClassName", classNames))
}

func (p *documentImpl) CreateElement(localName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapAsElement(p.call("createElement", localName))
	default:
		return wrapAsElement(p.call("createElement", localName, options[0].toJSObject()))
	}
}

func (p *documentImpl) CreateElementNS(namespace string, qualifiedName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapAsElement(p.call("createElementNS", namespace, qualifiedName))
	default:
		return wrapAsElement(p.call("createElementNS", namespace, qualifiedName, options[0].toJSObject()))
	}
}

func (p *documentImpl) CreateDocumentFragment() DocumentFragment {
	return wrapDocumentFragment(p.call("createDocumentFragment"))
}

func (p *documentImpl) CreateTextNode(data string) Text {
	return wrapText(p.call("createTextNode", data))
}

func (p *documentImpl) CreateCDATASection(data string) CDATASection {
	return wrapCDATASection(p.call("createCDATASection", data))
}

func (p *documentImpl) CreateComment(data string) Comment {
	return wrapComment(p.call("createComment", data))
}

func (p *documentImpl) CreateProcessingInstruction(target string, data string) ProcessingInstruction {
	return wrapProcessingInstruction(p.call("createProcessingInstruction", target, data))
}

func (p *documentImpl) ImportNode(node Node, deep ...bool) Node {
	switch len(deep) {
	case 0:
		return wrapAsNode(p.call("importNode", JSValue(node)))
	default:
		return wrapAsNode(p.call("importNode", JSValue(node), deep[0]))
	}
}

func (p *documentImpl) AdoptNode(node Node) Node {
	return wrapAsNode(p.call("adoptNode", JSValue(node)))
}

func (p *documentImpl) CreateAttribute(localName string) Attr {
	return wrapAttr(p.call("adoptNode", localName))
}

func (p *documentImpl) CreateAttributeNS(namespace string, qualifiedName string) Attr {
	return wrapAttr(p.call("createAttributeNS", namespace, qualifiedName))
}

func (p *documentImpl) CreateRange() Range {
	return wrapRange(p.call("createRange"))
}

func (p *documentImpl) CreateNodeIterator(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) NodeIterator {
	switch len(filter) {
	case 0:
		return wrapNodeIterator(p.call("createNodeIterator", JSValue(node), uint(whatToShow)))
	default:
		return wrapNodeIterator(p.call("createNodeIterator", JSValue(node), uint(whatToShow), JSValue(filter[0])))
	}
}

func (p *documentImpl) CreateTreeWalker(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) TreeWalker {
	switch len(filter) {
	case 0:
		return wrapTreeWalker(p.call("createTreeWalker", JSValue(node), uint(whatToShow)))
	default:
		return wrapTreeWalker(p.call("createTreeWalker", JSValue(node), uint(whatToShow), JSValue(filter[0])))
	}
}

func (p *documentImpl) FullscreenEnabled() bool {
	return p.get("fullscreenEnabled").toBool()
}

func (p *documentImpl) ExitFullscreen() func() error {
	return func() error {
		result, ok := await(p.call("exitFullscreen"))
		if ok {
			return nil
		}
		return wrapDOMException(result)
	}
}

func (p *documentImpl) OnFullscreenChange(fn func(Event)) EventHandler {
	return p.nodeImpl.eventTargetImpl.On("fullscreenchange", fn)
}

func (p *documentImpl) OnFullscreenError(fn func(Event)) EventHandler {
	return p.nodeImpl.eventTargetImpl.On("fullscreenerror", fn)
}

func (p *documentImpl) Location() Location {
	return wrapLocation(p.get("location"))
}

func (p *documentImpl) Domain() string {
	return p.get("domain").toString()
}

func (p *documentImpl) SetDomain(domain string) {
	p.set("domain", domain)
}

func (p *documentImpl) Referrer() string {
	return p.get("referrer").toString()
}

func (p *documentImpl) Cookie() string {
	return p.get("cookie").toString()
}

func (p *documentImpl) SetCookie(cookie string) {
	p.set("cookie", cookie)
}

func (p *documentImpl) LastModified() string {
	return p.get("lastModified").toString()
}

func (p *documentImpl) ReadyState() DocumentReadyState {
	return DocumentReadyState(p.get("readyState").toString())
}

/*
func (p *documentImpl) ByName(string) Value {

}
*/

func (p *documentImpl) Title() string {
	return p.get("title").toString()
}

func (p *documentImpl) SetTitle(title string) {
	p.set("title", title)
}

func (p *documentImpl) Dir() string {
	return p.get("dir").toString()
}

func (p *documentImpl) SetDir(dir string) {
	p.set("dir", dir)
}

func (p *documentImpl) Body() HTMLBodyElement {
	return wrapHTMLBodyElement(p.get("body"))
}

func (p *documentImpl) SetBody(body HTMLBodyElement) {
	p.set("body", JSValue(body))
}

func (p *documentImpl) Head() HTMLHeadElement {
	return wrapHTMLHeadElement(p.get("head"))
}

func (p *documentImpl) Images() []HTMLImageElement {
	if c := wrapHTMLCollection(p.get("images")); c != nil && c.Length() > 0 {
		var ret []HTMLImageElement
		for i := uint(0); i < c.Length(); i++ {
			if img, ok := c.Item(i).(HTMLImageElement); ok {
				ret = append(ret, img)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) Embeds() []HTMLEmbedElement {
	if c := wrapHTMLCollection(p.get("embeds")); c != nil && c.Length() > 0 {
		var ret []HTMLEmbedElement
		for i := uint(0); i < c.Length(); i++ {
			if embed, ok := c.Item(i).(HTMLEmbedElement); ok {
				ret = append(ret, embed)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) Plugins() []HTMLEmbedElement {
	// The plugins attribute must return the same object as that returned by the embeds attribute.
	return p.Embeds()
}

// returns <a> and <area> elements with href attributes, common interface is HTMLElement
func (p *documentImpl) Links() []HTMLElement {
	return htmlCollectionToHTMLElementSlice(p.get("links"))
}

func (p *documentImpl) Forms() []HTMLFormElement {
	if c := wrapHTMLCollection(p.get("form")); c != nil && c.Length() > 0 {
		var ret []HTMLFormElement
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLFormElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) Scripts() []HTMLScriptElement {
	if c := wrapHTMLCollection(p.get("scripts")); c != nil && c.Length() > 0 {
		var ret []HTMLScriptElement
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLScriptElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) ElementsByName(name string) []Node {
	return nodeListToSlice(p.call("getElementsByName", name))
}

func (p *documentImpl) CurrentScript() HTMLOrSVGScriptElement {
	return wrapHTMLOrSVGScriptElement(p.get("currentScript"))
}

func (p *documentImpl) Open(args ...string) Document {
	switch len(args) {
	case 1:
		return wrapDocument(p.call("open", args[0]))
	case 2:
		return wrapDocument(p.call("open", args[0], args[1]))
	default:
		return wrapDocument(p.call("open"))
	}
}

func (p *documentImpl) OpenURL(url string, name string, features string, replace ...bool) WindowProxy {
	switch len(replace) {
	case 0:
		return wrapWindowProxy(p.call("open", url, name, features))
	default:
		return wrapWindowProxy(p.call("open", url, name, features, replace[0]))
	}
}

func (p *documentImpl) Close() {
	p.call("close")
}

func (p *documentImpl) Write(text ...string) {
	if len(text) > 0 {
		var params []interface{}
		for _, v := range text {
			params = append(params, v)
		}
		p.call("write", params...)
	}
}

func (p *documentImpl) WriteLn(text ...string) {
	if len(text) > 0 {
		var params []interface{}
		for _, v := range text {
			params = append(params, v)
		}
		p.call("writeln", params...)
	}
}

func (p *documentImpl) DefaultView() WindowProxy {
	return wrapWindowProxy(p.get("defaultView"))
}

func (p *documentImpl) ActiveElement() Element {
	return wrapAsElement(p.get("activeElement"))
}

func (p *documentImpl) HasFocus() bool {
	return p.call("hasFocus").toBool()
}

func (p *documentImpl) DesignMode() string {
	return p.get("designMode").toString()
}

func (p *documentImpl) SetDesignMode(mode string) {
	p.set("designMode", mode)
}

func (p *documentImpl) ExecCommand(commandId string, args ...interface{}) bool {
	switch len(args) {
	case 1:
		if showUI, ok := args[0].(bool); ok {
			return p.call("execCommand", commandId, showUI).toBool()
		}
	case 2:
		if showUI, ok := args[0].(bool); ok {
			if value, ok := args[1].(string); ok {
				return p.call("execCommand", commandId, showUI, value).toBool()
			}
		}
	}

	return p.call("execCommand", commandId).toBool()
}

func (p *documentImpl) QueryCommandEnabled(commandId string) bool {
	return p.call("queryCommandEnabled", commandId).toBool()
}

func (p *documentImpl) QueryCommandIndeterm(commandId string) bool {
	return p.call("queryCommandIndeterm", commandId).toBool()
}

func (p *documentImpl) QueryCommandState(commandId string) bool {
	return p.call("queryCommandState", commandId).toBool()
}

func (p *documentImpl) QueryCommandSupported(commandId string) bool {
	return p.call("queryCommandSupported", commandId).toBool()
}

func (p *documentImpl) QueryCommandValue(commandId string) string {
	return p.call("queryCommandValue", commandId).toString()
}

func (p *documentImpl) OnReadyStateChange(fn func(Event)) EventHandler {
	return p.On("readystatechange", fn)
}

func (p *documentImpl) ElementFromPoint(x float64, y float64) Element {
	return wrapAsElement(p.call("elementFromPoint", x, y))
}

func (p *documentImpl) ElementsFromPoint(x float64, y float64) []Element {
	var ret []Element

	sl := p.call("elementsFromPoint", x, y).toSlice()
	if sl != nil {
		for _, v := range sl {
			ret = append(ret, wrapAsElement(v))
		}
	}

	return ret
}

func (p *documentImpl) CaretPositionFromPoint(x float64, y float64) CaretPosition {
	return wrapCaretPosition(p.call("caretPositionFromPoint", x, y))
}

func (p *documentImpl) ScrollingElement() Element {
	return wrapAsElement(p.get("scrollingElement"))
}

// helper function type assert on new Element return HTMLElement
func (p *documentImpl) CreateHTMLElement(tag string) HTMLElement {
	if el := p.CreateElement(tag); el != nil {
		if htmlEl, ok := el.(HTMLElement); ok {
			return htmlEl
		}
	}
	return nil
}

// -------------8<---------------------------------------

type domImplementationImpl struct {
	Value
}

func wrapDOMImplementation(v Value) DOMImplementation {
	if v.valid() {
		return &domImplementationImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domImplementationImpl) CreateDocumentType(qualifiedName string, publicId string, systemId string) DocumentType {
	return wrapDocumentType(p.call("createDocumentType", qualifiedName, publicId, systemId))
}

func (p *domImplementationImpl) CreateDocument(namespace string, qualifiedName string, doctype ...DocumentType) XMLDocument {
	switch len(doctype) {
	case 0:
		return wrapXMLDocument(p.call("createDocument", namespace, qualifiedName))
	default:
		return wrapXMLDocument(p.call("createDocument", namespace, qualifiedName, JSValue(doctype[0])))
	}
}

func (p *domImplementationImpl) CreateHTMLDocument(title ...string) Document {
	if len(title) > 0 {
		return wrapDocument(p.call("createHTMLDocument", title[0]))
	}
	return wrapDocument(p.call("createHTMLDocument"))
}

// -------------8<---------------------------------------

type xmlDocumentImpl struct {
	*documentImpl
}

func wrapXMLDocument(v Value) XMLDocument {
	if v.valid() {
		return &xmlDocumentImpl{
			documentImpl: newDocumentImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type treeWalkerImpl struct {
	Value
}

func wrapTreeWalker(v Value) TreeWalker {
	if v.valid() {
		return &treeWalkerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *treeWalkerImpl) Root() Node {
	return wrapAsNode(p.get("root"))
}

func (p *treeWalkerImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.get("whatToShow").toInt()))
}

func (p *treeWalkerImpl) Filter() NodeFilter {
	return wrapNodeFilter(p.get("filter"))
}

func (p *treeWalkerImpl) CurrentNode() Node {
	return wrapAsNode(p.get("currentNode"))
}

func (p *treeWalkerImpl) SetCurrentNode(node Node) {
	p.set("currentNode", JSValue(node))
}

func (p *treeWalkerImpl) ParentNode() Node {
	return wrapAsNode(p.call("parentNode"))
}

func (p *treeWalkerImpl) FirstChild() Node {
	return wrapAsNode(p.call("firstChild"))
}

func (p *treeWalkerImpl) LastChild() Node {
	return wrapAsNode(p.call("lastChild"))
}

func (p *treeWalkerImpl) PreviousSibling() Node {
	return wrapAsNode(p.call("previousSibling"))
}

func (p *treeWalkerImpl) NextSibling() Node {
	return wrapAsNode(p.call("nextSibling"))
}

func (p *treeWalkerImpl) PreviousNode() Node {
	return wrapAsNode(p.call("previousNode"))
}

func (p *treeWalkerImpl) NextNode() Node {
	return wrapAsNode(p.call("nextNode"))
}

// -------------8<---------------------------------------

type nodeFilterImpl struct {
	Value
}

func wrapNodeFilter(v Value) NodeFilter {
	if v.valid() {
		return &nodeFilterImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeFilterImpl) AcceptNode(node Node) NodeFilterResult {
	return NodeFilterResult(p.call("acceptNode", JSValue(node)).toUint16())
}

// -------------8<---------------------------------------

type nodeIteratorImpl struct {
	Value
}

func wrapNodeIterator(v Value) NodeIterator {
	if v.valid() {
		return &nodeIteratorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeIteratorImpl) Root() Node {
	return wrapAsNode(p.get("root"))
}

func (p *nodeIteratorImpl) ReferenceNode() Node {
	return wrapAsNode(p.get("referenceNode"))
}

func (p *nodeIteratorImpl) PointerBeforeReferenceNode() bool {
	return p.get("pointerBeforeReferenceNode").toBool()
}

func (p *nodeIteratorImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.get("whatToShow").toInt()))
}

func (p *nodeIteratorImpl) Filter() NodeFilter {
	return wrapNodeFilter(p.get("filter"))
}

func (p *nodeIteratorImpl) NextNode() Node {
	return wrapAsNode(p.call("nextNode"))
}

func (p *nodeIteratorImpl) PreviousNode() Node {
	return wrapAsNode(p.call("previousNode"))
}

func (p *nodeIteratorImpl) Detach() {
	p.call("detach")
}

// -------------8<---------------------------------------

type rangeImpl struct {
	*abstractRangeImpl
}

func wrapRange(v Value) Range {
	if p := newRangeImpl(v); p != nil {
		return p
	}
	return nil
}

func newRangeImpl(v Value) *rangeImpl {
	if v.valid() {
		return &rangeImpl{
			abstractRangeImpl: newAbstractRangeImpl(v),
		}
	}
	return nil
}

func (p *rangeImpl) CommonAncestorContainer() Node {
	return wrapAsNode(p.get("commonAncestorContainer"))
}

func (p *rangeImpl) SetStart(node Node, offset int) {
	p.call("setStart", JSValue(node), offset)
}

func (p *rangeImpl) SetEnd(node Node, offset int) {
	p.call("setEnd", JSValue(node), offset)
}

func (p *rangeImpl) SetStartBefore(node Node) {
	p.call("setStartBefore", JSValue(node))
}

func (p *rangeImpl) SetStartAfter(node Node) {
	p.call("setStartAfter", JSValue(node))
}

func (p *rangeImpl) SetEndBefore(node Node) {
	p.call("setEndBefore", JSValue(node))
}

func (p *rangeImpl) SetEndAfter(node Node) {
	p.call("setEndAfter", JSValue(node))
}

func (p *rangeImpl) Collapse(toStart ...bool) {
	switch len(toStart) {
	case 0:
		p.call("collapse")
	default:
		p.call("collapse", toStart[0])
	}
}

func (p *rangeImpl) SelectNode(node Node) {
	p.call("selectNode", JSValue(node))
}

func (p *rangeImpl) SelectNodeContents(node Node) {
	p.call("selectNodeContents", JSValue(node))
}

func (p *rangeImpl) CompareBoundaryPoints(how RangeCompare, source Range) int {
	return p.call("compareBoundaryPoints", int(how), JSValue(source)).toInt()
}

func (p *rangeImpl) DeleteContents() {
	p.call("deleteContents")
}

func (p *rangeImpl) ExtractContents() DocumentFragment {
	return wrapDocumentFragment(p.call("extractContents"))
}

func (p *rangeImpl) CloneContents() DocumentFragment {
	return wrapDocumentFragment(p.call("cloneContents"))
}

func (p *rangeImpl) InsertNode(node Node) {
	p.call("insertNode", JSValue(node))
}

func (p *rangeImpl) SurroundContents(newParent Node) {
	p.call("surroundContents", JSValue(newParent))
}

func (p *rangeImpl) CloneRange() Range {
	return wrapRange(p.call("cloneRange"))
}

func (p *rangeImpl) Detach() {
	p.call("detach")
}

func (p *rangeImpl) IsPointInRange(node Node, offset int) bool {
	return p.call("isPointInRange", JSValue(node), offset).toBool()
}

func (p *rangeImpl) ComparePoint(node Node, offset int) int {
	return p.call("comparePoint", JSValue(node), offset).toInt()
}

func (p *rangeImpl) IntersectsNode(node Node) bool {
	return p.call("intersectsNode", JSValue(node)).toBool()
}

func (p *rangeImpl) ClientRects() []DOMRect {
	if rects := p.call("getClientRects").toSlice(); rects != nil {
		var ret []DOMRect
		for _, rect := range rects {
			ret = append(ret, wrapDOMRect(rect))
		}
		return ret
	}
	return nil
}

func (p *rangeImpl) BoundingClientRect() DOMRect {
	return wrapDOMRect(p.call("getBoundingClientRect"))
}

func (p *rangeImpl) CreateContextualFragment(fragment string) DocumentFragment {
	return wrapDocumentFragment(p.call("createContextualFragment", fragment))
}

// -------------8<---------------------------------------

type abstractRangeImpl struct {
	Value
}

func wrapAbstractRange(v Value) AbstractRange {
	if p := newAbstractRangeImpl(v); p != nil {
		return p
	}
	return nil
}

func newAbstractRangeImpl(v Value) *abstractRangeImpl {
	if v.valid() {
		return &abstractRangeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *abstractRangeImpl) StartContainer() Node {
	return wrapAsNode(p.get("startContainer"))
}

func (p *abstractRangeImpl) StartOffset() int {
	return p.get("startOffset").toInt()
}

func (p *abstractRangeImpl) EndContainer() Node {
	return wrapAsNode(p.get("endContainer"))
}

func (p *abstractRangeImpl) EndOffset() int {
	return p.get("endOffset").toInt()
}

func (p *abstractRangeImpl) Collapsed() bool {
	return p.get("collapsed").toBool()
}

// -------------8<---------------------------------------

type staticRangeImpl struct {
	*abstractRangeImpl
}

func wrapStaticRange(v Value) StaticRange {
	if v.valid() {
		return &staticRangeImpl{
			abstractRangeImpl: newAbstractRangeImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type processingInstructionImpl struct {
	*characterDataImpl
	*linkStyleImpl
	Value
}

func wrapProcessingInstruction(v Value) ProcessingInstruction {
	if v.valid() {
		return &processingInstructionImpl{
			characterDataImpl: newCharacterDataImpl(v),
			linkStyleImpl:     newLinkStyleImpl(v),
			Value:             v,
		}
	}
	return nil
}

func (p *processingInstructionImpl) Target() string {
	return p.get("target").toString()
}

// -------------8<---------------------------------------

type commentImpl struct {
	*characterDataImpl
}

func wrapComment(v Value) Comment {
	if v.valid() {
		return &commentImpl{
			characterDataImpl: newCharacterDataImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type cDATASectionImpl struct {
	*textImpl
}

func wrapCDATASection(v Value) CDATASection {
	if v.valid() {
		return &cDATASectionImpl{
			textImpl: newTextImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type textImpl struct {
	*characterDataImpl
	*slotableImpl
	*geometryUtilsImpl
	Value
}

func NewText(data ...string) Text {
	if jsText := jsGlobal.get("Text"); jsText.valid() {
		switch len(data) {
		case 0:
			return wrapText(jsText.jsNew())
		default:
			return wrapText(jsText.jsNew(data[0]))
		}
	}
	return nil
}

func wrapText(v Value) Text {
	if p := newTextImpl(v); p != nil {
		return p
	}
	return nil
}

func newTextImpl(v Value) *textImpl {
	if v.valid() {
		return &textImpl{
			characterDataImpl: newCharacterDataImpl(v),
			slotableImpl:      newSlotableImpl(v),
			geometryUtilsImpl: newGeometryUtilsImpl(v),
			Value:             v,
		}
	}
	return nil
}

func (p *textImpl) SplitText(offset int) Text {
	return wrapText(p.call("splitText", offset))
}

func (p *textImpl) WholeText() string {
	return p.get("wholeText").toString()
}

// -------------8<---------------------------------------

type characterDataImpl struct {
	*nodeImpl
	*nonDocumentTypeChildNodeImpl
	*childNodeImpl
	Value
}

func wrapCharacterData(v Value) CharacterData {
	if p := newCharacterDataImpl(v); p != nil {
		return p
	}
	return nil
}

func newCharacterDataImpl(v Value) *characterDataImpl {
	if v.valid() {
		return &characterDataImpl{
			nodeImpl:                     newNodeImpl(v),
			nonDocumentTypeChildNodeImpl: newNonDocumentTypeChildNodeImpl(v),
			childNodeImpl:                newChildNodeImpl(v),
			Value:                        v,
		}
	}
	return nil
}

func (p *characterDataImpl) Data() string {
	return p.get("data").toString()
}

func (p *characterDataImpl) SetData(data string) {
	p.set("data", data)
}

func (p *characterDataImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *characterDataImpl) Substring(offset uint, count uint) string {
	return p.call("substringData", offset, count).toString()
}

func (p *characterDataImpl) Append(data string) {
	p.call("appendData", data)
}

func (p *characterDataImpl) Insert(offset int, data string) {
	p.call("insertData", offset, data)
}

func (p *characterDataImpl) Delete(offset int, count int) {
	p.call("deleteData", offset, count)
}

func (p *characterDataImpl) Replace(offset int, count int, data string) {
	p.call("replaceData", offset, count, data)
}

// -------------8<---------------------------------------

type documentTypeImpl struct {
	*nodeImpl
	*childNodeImpl
	Value
}

func wrapDocumentType(v Value) DocumentType {
	if v.valid() {
		return &documentTypeImpl{
			nodeImpl:      newNodeImpl(v),
			childNodeImpl: newChildNodeImpl(v),
			Value:         v,
		}
	}
	return nil
}

func (p *documentTypeImpl) Name() string {
	return p.get("name").toString()
}

func (p *documentTypeImpl) PublicId() string {
	return p.get("publicId").toString()
}

func (p *documentTypeImpl) SystemId() string {
	return p.get("systemId").toString()
}

// -------------8<---------------------------------------

type nodeImpl struct {
	*eventTargetImpl
}

func wrapNode(v Value) Node {
	if p := newNodeImpl(v); p != nil {
		return p
	}
	return nil
}

func newNodeImpl(v Value) *nodeImpl {
	if v.valid() {
		return &nodeImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *nodeImpl) NodeType() NodeType {
	return NodeType(p.get("nodeType").toInt())
}

func (p *nodeImpl) NodeName() string {
	return p.get("nodeName").toString()
}

func (p *nodeImpl) BaseURI() string {
	return p.get("baseURI").toString()
}

func (p *nodeImpl) IsConnected() bool {
	return p.get("isConnected").toBool()
}

func (p *nodeImpl) OwnerDocument() Document {
	return wrapDocument(p.get("ownerDocument"))
}

func (p *nodeImpl) RootNode(options ...RootNodeOptions) Node {
	if len(options) > 0 {
		return wrapAsNode(p.call("getRootNode", options[0].toJSObject()))
	}

	return wrapAsNode(p.call("getRootNode"))
}

func (p *nodeImpl) ParentNode() Node {
	return wrapAsNode(p.get("parentNode"))
}

func (p *nodeImpl) ParentElement() Element {
	return wrapAsElement(p.get("parentElement"))
}

func (p *nodeImpl) HasChildNodes() bool {
	return p.call("hasChildNodes").toBool()
}

func (p *nodeImpl) ChildNodes() []Node {
	return nodeListToSlice(p.get("childNodes"))
}

func (p *nodeImpl) FirstChild() Node {
	return wrapAsNode(p.get("firstChild"))
}

func (p *nodeImpl) LastChild() Node {
	return wrapAsNode(p.get("lastChild"))
}

func (p *nodeImpl) PreviousSibling() Node {
	return wrapAsNode(p.get("previousSibling"))
}

func (p *nodeImpl) NextSibling() Node {
	return wrapAsNode(p.get("nextSibling"))
}

func (p *nodeImpl) NodeValue() string {
	return p.get("nodeValue").toString()
}

func (p *nodeImpl) SetNodeValue(nval string) {
	p.set("nodeValue", nval)
}

func (p *nodeImpl) TextContent() string {
	return p.get("textContent").toString()
}

func (p *nodeImpl) SetTextContent(tc string) {
	p.set("textContent", tc)
}

func (p *nodeImpl) Normalize() {
	p.call("normalize")
}

func (p *nodeImpl) CloneNode(deep ...bool) Node {
	if len(deep) > 0 {
		return wrapAsNode(p.call("cloneNode", deep[0]))
	}
	return wrapAsNode(p.call("cloneNode"))
}

func (p *nodeImpl) IsEqualNode(otherNode Node) bool {
	return p.call("isEqualNode", JSValue(otherNode)).toBool()
}

func (p *nodeImpl) IsSameNode(otherNode Node) bool {
	return p.call("isSameNode", JSValue(otherNode)).toBool()
}

func (p *nodeImpl) CompareDocumentPosition(other Node) DocumentPosition {
	return DocumentPosition(p.call("compareDocumentPosition", JSValue(other)).toInt())
}

func (p *nodeImpl) Contains(other Node) bool {
	return p.call("contains", JSValue(other)).toBool()
}

func (p *nodeImpl) LookupPrefix(namespace string) string {
	return p.call("lookupPrefix", namespace).toString()
}

func (p *nodeImpl) LookupNamespaceURI(prefix string) string {
	return p.call("lookupNamespaceURI", prefix).toString()
}

func (p *nodeImpl) IsDefaultNamespace(namespace string) bool {
	return p.call("isDefaultNamespace", namespace).toBool()
}

func (p *nodeImpl) InsertBefore(node Node, child Node) Node {
	if child != nil {
		return wrapAsNode(p.call("insertBefore", JSValue(node), JSValue(child)))
	}
	return wrapAsNode(p.call("insertBefore", JSValue(node)))
}

func (p *nodeImpl) AppendChild(node Node) Node {
	return wrapAsNode(p.call("appendChild", JSValue(node)))
}

func (p *nodeImpl) ReplaceChild(node Node, child Node) Node {
	return wrapAsNode(p.call("replaceChild", JSValue(node), JSValue(child)))
}

func (p *nodeImpl) RemoveChild(child Node) Node {
	return wrapAsNode(p.call("removeChild", JSValue(child)))
}

// -------------8<---------------------------------------

type elementImpl struct {
	*nodeImpl
	*parentNodeImpl
	*nonDocumentTypeChildNodeImpl
	*childNodeImpl
	*slotableImpl
	*geometryUtilsImpl
	Value
}

func wrapElement(v Value) Element {
	if p := newElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newElementImpl(v Value) *elementImpl {
	if v.valid() {
		return &elementImpl{
			nodeImpl:                     newNodeImpl(v),
			parentNodeImpl:               newParentNodeImpl(v),
			nonDocumentTypeChildNodeImpl: newNonDocumentTypeChildNodeImpl(v),
			childNodeImpl:                newChildNodeImpl(v),
			slotableImpl:                 newSlotableImpl(v),
			geometryUtilsImpl:            newGeometryUtilsImpl(v),
			Value:                        v,
		}
	}
	return nil
}

func (p *elementImpl) NamespaceURI() string {
	return p.get("namespaceURI").toString()
}

func (p *elementImpl) Prefix() string {
	return p.get("prefix").toString()
}

func (p *elementImpl) LocalName() string {
	return p.get("localName").toString()
}

func (p *elementImpl) TagName() string {
	return p.get("tagName").toString()
}

func (p *elementImpl) Id() string {
	return p.get("id").toString()
}

func (p *elementImpl) SetId(id string) {
	p.set("id", id)
}

func (p *elementImpl) ClassName() string {
	return p.get("className").toString()
}

func (p *elementImpl) SetClassName(name string) {
	p.set("className", name)
}

func (p *elementImpl) ClassList() DOMTokenList {
	return wrapDOMTokenList(p.get("classList"))
}

func (p *elementImpl) Slot() string {
	return p.get("slot").toString()
}

func (p *elementImpl) SetSlot(slot string) {
	p.set("slot", slot)
}

func (p *elementImpl) HasAttributes() bool {
	return p.call("hasAttributes").toBool()
}

func (p *elementImpl) Attributes() NamedNodeMap {
	return wrapNamedNodeMap(p.get("attributes"))
}

func (p *elementImpl) AttributeNames() []string {
	return stringSequenceToSlice(p.call("getAttributeNames"))
}

func (p *elementImpl) Attribute(name string) string {
	return p.call("getAttribute", name).toString()
}

func (p *elementImpl) AttributeNS(namespace string, localName string) string {
	return p.call("getAttributeNS", namespace, localName).toString()
}

func (p *elementImpl) SetAttribute(name string, value string) {
	p.call("setAttribute", name, value)
}

func (p *elementImpl) SetAttributeNS(namespace string, name string, value string) {
	p.call("setAttributeNS", namespace, name, value)
}

func (p *elementImpl) RemoveAttribute(name string) {
	p.call("removeAttribute", name)
}

func (p *elementImpl) RemoveAttributeNS(namespace string, name string) {
	p.call("removeAttributeNS", namespace, name)
}

func (p *elementImpl) ToggleAttribute(name string, force ...bool) bool {
	if len(force) > 0 {
		return p.call("toggleAttribute", name, force[0]).toBool()
	}
	return p.call("toggleAttribute", name).toBool()
}

func (p *elementImpl) HasAttribute(name string) bool {
	return p.call("hasAttribute", name).toBool()
}

func (p *elementImpl) HasAttributeNS(namespace string, localName string) bool {
	return p.call("hasAttributeNS", namespace, localName).toBool()
}

func (p *elementImpl) AttributeNode(name string) Attr {
	return wrapAttr(p.call("getAttributeNode", name))
}

func (p *elementImpl) AttributeNodeNS(namespace string, name string) Attr {
	return wrapAttr(p.call("getAttributeNodeNS", namespace, name))
}

func (p *elementImpl) SetAttributeNode(attr Attr) Attr {
	return wrapAttr(p.call("setAttributeNode", JSValue(attr)))
}

func (p *elementImpl) SetAttributeNodeNS(attr Attr) Attr {
	return wrapAttr(p.call("setAttributeNodeNS", JSValue(attr)))
}

func (p *elementImpl) RemoveAttributeNode(attr Attr) Attr {
	return wrapAttr(p.call("removeAttributeNode", JSValue(attr)))
}

func (p *elementImpl) AttachShadow(si ShadowRootInit) ShadowRoot {
	return wrapShadowRoot(p.call("attachShadow", si.toJSObject()))
}

func (p *elementImpl) ShadowRoot() ShadowRoot {
	return wrapShadowRoot(p.get("shadowRoot"))
}

func (p *elementImpl) Closest(selectors string) Element {
	return wrapAsElement(p.call("closest"))
}

func (p *elementImpl) Matches(selector string) bool {
	return p.call("matches", selector).toBool()
}

func (p *elementImpl) ElementsByTagName(name string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByTagName", name))
}

func (p *elementImpl) ElementsByTagNameNS(namespace string, localName string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByTagNameNS", namespace, localName))
}

func (p *elementImpl) ElementsByClassName(names string) []Element {
	return htmlCollectionToElementSlice(p.call("getElementsByClassName", names))
}

func (p *elementImpl) ClientRects() []DOMRect {
	if rects := p.call("getClientRects").toSlice(); rects != nil {
		var ret []DOMRect
		for _, rect := range rects {
			ret = append(ret, wrapDOMRect(rect))
		}
		return ret
	}
	return nil
}

func (p *elementImpl) BoundingClientRect() DOMRect {
	return wrapDOMRect(p.call("getBoundingClientRect"))
}

func (p *elementImpl) ScrollIntoView(arg ...interface{}) {
	switch len(arg) {
	case 0:
		p.call("scrollIntoView")
	default:
		switch x := arg[0].(type) {
		case bool:
			p.call("scrollIntoView", x)
		case ScrollIntoViewOptions:
			p.call("scrollIntoView", x.toJSObject())
		}
	}
}

func (p *elementImpl) Scroll(options ScrollToOptions) {
	p.call("scroll", options.toJSObject())
}

func (p *elementImpl) ScrollTo(options ScrollToOptions) {
	p.call("scrollTo", options.toJSObject())
}

func (p *elementImpl) ScrollBy(options ScrollToOptions) {
	p.call("scrollBy", options.toJSObject())
}

func (p *elementImpl) ScrollTop() float64 {
	return p.get("scrollTop").toFloat64()
}

func (p *elementImpl) SetScrollTop(st float64) {
	p.set("scrollTop", st)
}

func (p *elementImpl) ScrollLeft() float64 {
	return p.get("scrollLeft").toFloat64()
}

func (p *elementImpl) SetScrollLeft(sl float64) {
	p.set("scrollLeft", sl)
}

func (p *elementImpl) ScrollWidth() int {
	return p.get("scrollWidth").toInt()
}

func (p *elementImpl) ScrollHeight() int {
	return p.get("scrollHeight").toInt()
}

func (p *elementImpl) ClientTop() int {
	return p.get("clientTop").toInt()
}

func (p *elementImpl) ClientLeft() int {
	return p.get("clientLeft").toInt()
}

func (p *elementImpl) ClientWidth() int {
	return p.get("clientWidth").toInt()
}

func (p *elementImpl) ClientHeight() int {
	return p.get("clientHeight").toInt()
}

func (p *elementImpl) OnFullscreenChange(fn func(Event)) EventHandler {
	return p.On("fullscreenchange", fn)
}

func (p *elementImpl) OnFullscreenError(fn func(Event)) EventHandler {
	return p.On("fullscreenerror", fn)
}

func (p *elementImpl) InnerHTML() string {
	return p.get("innerHTML").toString()
}

func (p *elementImpl) SetInnerHTML(html string) {
	p.set("innerHTML", html)
}

func (p *elementImpl) OuterHTML() string {
	return p.get("outerHTML").toString()
}

func (p *elementImpl) SetOuterHTML(html string) {
	p.set("outerHTML", html)
}

func (p *elementImpl) InsertAdjacentHTML(position string, text string) {
	p.call("insertAdjacentHTML", position, text)
}

func (p *elementImpl) RequestFullscreen(...FullscreenOptions) func() error {
	return func() error {
		res, ok := await(p.call("requestFullscreen"))
		if ok {
			return nil
		}
		return wrapDOMException(res)
	}
}

func (p *elementImpl) OnFullScreenChange(fn func(Event)) EventHandler {
	return p.On("fullscreenchange", fn)
}

func (p *elementImpl) OnFullScreenError(fn func(Event)) EventHandler {
	return p.On("fullscreenerror", fn)
}

func (p *elementImpl) SetPointerCapture(pointerId int) {
	p.call("setPointerCapture", pointerId)
}

func (p *elementImpl) ReleasePointerCapture(pointerId int) {
	p.call("releasePointerCapture", pointerId)
}

func (p *elementImpl) HasPointerCapture(pointerId int) bool {
	return p.call("hasPointerCapture", pointerId).toBool()
}

// -------------8<---------------------------------------

type shadowRootImpl struct {
	*documentFragmentImpl
	*documentOrShadowRootImpl
	*parentNodeImpl
	Value
}

func wrapShadowRoot(v Value) ShadowRoot {
	if v.valid() {
		return &shadowRootImpl{
			documentFragmentImpl:     newDocumentFragmentImpl(v),
			documentOrShadowRootImpl: newDocumentOrShadowRootImpl(v),
			parentNodeImpl:           newParentNodeImpl(v),
			Value:                    v,
		}
	}
	return nil
}

func (p *shadowRootImpl) Mode() ShadowRootMode {
	return ShadowRootMode(p.get("mode").toString())
}

func (p *shadowRootImpl) Host() Element {
	return wrapAsElement(p.get("host"))
}

// -------------8<---------------------------------------

type documentFragmentImpl struct {
	*nodeImpl
	*nonElementParentNodeImpl
}

func wrapDocumentFragment(v Value) DocumentFragment {
	if p := newDocumentFragmentImpl(v); p != nil {
		return p
	}
	return nil
}

func newDocumentFragmentImpl(v Value) *documentFragmentImpl {
	if v.valid() {
		return &documentFragmentImpl{
			nodeImpl:                 newNodeImpl(v),
			nonElementParentNodeImpl: newNonElementParentNodeImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type domTokenListImpl struct {
	Value
}

func wrapDOMTokenList(v Value) DOMTokenList {
	if v.valid() {
		return &domTokenListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domTokenListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *domTokenListImpl) Item(index uint) string {
	return p.get("item").toString()
}

func (p *domTokenListImpl) Contains(token string) bool {
	return p.call("contains", token).toBool()
}

func (p *domTokenListImpl) Add(tokens ...string) {
	if len(tokens) > 0 {
		var params []interface{}
		for _, v := range tokens {
			params = append(params, v)
		}
		p.call("add", params...)
	}
}

func (p *domTokenListImpl) Remove(tokens ...string) {
	if len(tokens) > 0 {
		var params []interface{}
		for _, v := range tokens {
			params = append(params, v)
		}
		p.call("remove", params...)
	}
}

func (p *domTokenListImpl) Toggle(token string, force ...bool) bool {
	if len(force) > 0 {
		return p.call("toggle", token, force[0]).toBool()
	}
	return p.call("toggle", token).toBool()
}

func (p *domTokenListImpl) Replace(token string, newToken string) bool {
	return p.call("replace", token, newToken).toBool()
}

func (p *domTokenListImpl) Supports(token string) bool {
	return p.call("supports", token).toBool()
}

func (p *domTokenListImpl) TokenValue() string {
	return p.get("value").toString()
}

func (p *domTokenListImpl) SetTokenValue(value string) {
	p.set("value", value)
}

func (p *domTokenListImpl) TokenValues() []string {
	var ret []string
	it := p.call("values")
	for {
		n := it.call("next")
		if n.get("done").toBool() {
			break
		}

		ret = append(ret, n.get("value").toString())
	}
	return ret
}

func (p *domTokenListImpl) String() string {
	return p.call("toString").toString()
}

// -------------8<---------------------------------------

type namedNodeMapImpl struct {
	Value
}

func wrapNamedNodeMap(v Value) NamedNodeMap {
	if v.valid() {
		return &namedNodeMapImpl{
			Value: v,
		}
	}
	return nil
}

func (p *namedNodeMapImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *namedNodeMapImpl) Item(index uint) Attr {
	return wrapAttr(p.call("item", index))
}

func (p *namedNodeMapImpl) NamedItem(name string) Attr {
	return wrapAttr(p.get("getNamedItem"))
}

func (p *namedNodeMapImpl) NamedItemNS(namespace string, name string) Attr {
	return wrapAttr(p.call("getNamedItemNS", namespace, name))
}

func (p *namedNodeMapImpl) SetNamedItem(attr Attr) Attr {
	return wrapAttr(p.call("setNamedItem", JSValue(attr)))
}

func (p *namedNodeMapImpl) SetNamedItemNS(attr Attr) Attr {
	return wrapAttr(p.call("setNamedItemNS", JSValue(attr)))
}

func (p *namedNodeMapImpl) RemoveNamedItem(name string) Attr {
	return wrapAttr(p.call("removeNamedItem", name))
}

func (p *namedNodeMapImpl) RemoveNamedItemNS(namespace string, name string) Attr {
	return wrapAttr(p.call("removeNamedItemNS", namespace, name))
}

// -------------8<---------------------------------------

type attrImpl struct {
	*nodeImpl
}

func wrapAttr(v Value) Attr {
	if v.valid() {
		return &attrImpl{
			nodeImpl: newNodeImpl(v),
		}
	}
	return nil
}

func (p *attrImpl) NamespaceURI() string {
	return p.get("namespaceURI").toString()
}

func (p *attrImpl) Prefix() string {
	return p.get("prefix").toString()
}

func (p *attrImpl) LocalName() string {
	return p.get("localName").toString()
}

func (p *attrImpl) Name() string {
	return p.get("name").toString()
}

func (p *attrImpl) Value() string {
	return p.get("value").toString()
}

func (p *attrImpl) SetValue(value string) {
	p.set("value", value)
}

func (p *attrImpl) OwnerElement() Element {
	return wrapAsElement(p.get("ownerElement"))
}

// -------------8<---------------------------------------

type htmlCollectionImpl struct {
	Value
}

func wrapHTMLCollection(v Value) HTMLCollection {
	if p := newHTMLCollectionImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLCollectionImpl(v Value) *htmlCollectionImpl {
	if v.valid() {
		return &htmlCollectionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *htmlCollectionImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *htmlCollectionImpl) Item(index uint) Element {
	return wrapAsElement(p.call("item", index))
}

func (p *htmlCollectionImpl) NamedItem(name string) Element {
	return wrapAsElement(p.call("namedItem", name))
}

// -------------8<---------------------------------------

type mutationRecordImpl struct {
	Value
}

func wrapMutationRecord(v Value) MutationRecord {
	if v.valid() {
		return &mutationRecordImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mutationRecordImpl) Type() string {
	return p.get("type").toString()
}

func (p *mutationRecordImpl) Target() Node {
	return wrapAsNode(p.get("target"))
}

func (p *mutationRecordImpl) AddedNodes() []Node {
	return nodeListToSlice(p.get("addedNodes"))
}

func (p *mutationRecordImpl) RemovedNodes() []Node {
	return nodeListToSlice(p.get("removedNodes"))
}

func (p *mutationRecordImpl) PreviousSibling() Node {
	return wrapAsNode(p.get("previousSibling"))
}

func (p *mutationRecordImpl) NextSibling() Node {
	return wrapAsNode(p.get("nextSibling"))
}

func (p *mutationRecordImpl) AttributeName() string {
	return p.get("attributeName").toString()
}

func (p *mutationRecordImpl) AttributeNamespace() string {
	return p.get("attributeNamespace").toString()
}

func (p *mutationRecordImpl) OldValue() string {
	return p.get("oldValue").toString()
}

// -------------8<---------------------------------------

type mutationObserverImpl struct {
	Value
}

func wrapMutationObserver(v Value) MutationObserver {
	if v.valid() {
		return &mutationObserverImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mutationObserverImpl) Observe(target Node, options ...MutationObserverInit) {
	switch len(options) {
	case 0:
		p.call("observe", JSValue(target))
	default:
		p.call("observe", JSValue(target), options[0].toJSObject())
	}
}

func (p *mutationObserverImpl) Disconnect() {
	p.call("disconnect")
}

func (p *mutationObserverImpl) TakeRecords() []MutationRecord {
	if s := p.call("takeRecords").toSlice(); s != nil {
		var ret []MutationRecord
		for _, v := range s {
			ret = append(ret, wrapMutationRecord(v))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type htmlUnknownElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLUnknownElement(v Value) HTMLUnknownElement {
	if v.valid() {
		return &htmlUnknownElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlElementImpl struct {
	*eventTargetImpl
	*elementImpl
	*globalEventHandlersImpl
	*documentAndElementEventHandlersImpl
	Value
}

func wrapHTMLElement(v Value) HTMLElement {
	if p := newHTMLElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLElementImpl(v Value) *htmlElementImpl {
	if v.valid() {
		ei := &htmlElementImpl{
			elementImpl: newElementImpl(v),
			Value:       v,
		}
		ei.eventTargetImpl = ei.elementImpl.eventTargetImpl
		ei.globalEventHandlersImpl = newGlobalEventHandlersImpl(ei.eventTargetImpl)
		ei.documentAndElementEventHandlersImpl = newDocumentAndElementEventHandlersImpl(ei.eventTargetImpl)
		return ei
	}
	return nil
}

func (p *htmlElementImpl) Title() string {
	return p.get("title").toString()
}

func (p *htmlElementImpl) SetTitle(title string) {
	p.set("title", title)
}

func (p *htmlElementImpl) Lang() string {
	return p.get("lang").toString()
}

func (p *htmlElementImpl) SetLang(lang string) {
	p.set("lang", lang)
}

func (p *htmlElementImpl) Translate() bool {
	return p.get("translate").toBool()
}

func (p *htmlElementImpl) SetTranslate(tr bool) {
	p.set("translate", tr)
}

func (p *htmlElementImpl) Dir() string {
	return p.get("dir").toString()
}

func (p *htmlElementImpl) SetDir(dir string) {
	p.set("dir", dir)
}

func (p *htmlElementImpl) DataSet() DOMStringMap {
	return wrapDOMStringMap(p.get("dataset"))
}

func (p *htmlElementImpl) Hidden() bool {
	return p.get("hidden").toBool()
}

func (p *htmlElementImpl) SetHidden(hidden bool) {
	p.set("hidden", hidden)
}

func (p *htmlElementImpl) Click() {
	p.call("click")
}

func (p *htmlElementImpl) TabIndex() int {
	return p.get("tabIndex").toInt()
}

func (p *htmlElementImpl) SetTabIndex(index int) {
	p.set("tabIndex", index)
}

func (p *htmlElementImpl) Focus() {
	p.call("focus")
}

func (p *htmlElementImpl) Blur() {
	p.call("blur")
}

func (p *htmlElementImpl) AccessKey() string {
	return p.get("accessKey").toString()
}

func (p *htmlElementImpl) SetAccessKey(key string) {
	p.set("accessKey", key)
}

func (p *htmlElementImpl) AccessKeyLabel() string {
	return p.get("accessKeyLabel").toString()
}

func (p *htmlElementImpl) Draggable() bool {
	return p.get("draggable").toBool()
}

func (p *htmlElementImpl) SetDraggable(d bool) {
	p.set("draggable", d)
}

func (p *htmlElementImpl) SpellCheck() bool {
	return p.get("spellcheck").toBool()
}

func (p *htmlElementImpl) SetSpellChack(s bool) {
	p.set("spellcheck", s)
}

func (p *htmlElementImpl) ForceSpellCheck() {
	p.call("forceSpellCheck")
}

/*
func (p *htmlElementImpl) Autocapitalize() string {

}


func (p *htmlElementImpl) SetAutocapitalize(string) {

}
*/

func (p *htmlElementImpl) InnerText() string {
	return p.get("innerText").toString()
}

func (p *htmlElementImpl) SetInnerText(text string) {
	p.set("innerText", text)
}

func (p *htmlElementImpl) OffsetParent() Element {
	return wrapAsElement(p.get("offsetParent"))
}

func (p *htmlElementImpl) OffsetTop() int {
	return p.get("offsetTop").toInt()
}

func (p *htmlElementImpl) OffsetLeft() int {
	return p.get("offsetLeft").toInt()
}

func (p *htmlElementImpl) OffsetWidth() int {
	return p.get("offsetWidth").toInt()
}

func (p *htmlElementImpl) OffsetHeight() int {
	return p.get("offsetHeight").toInt()
}

func (p *htmlElementImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.get("style"))
}

// -------------8<---------------------------------------

type domStringMapImpl struct {
	Value
}

func wrapDOMStringMap(v Value) DOMStringMap {
	if v.valid() {
		return &domStringMapImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domStringMapImpl) Get(name string) string {
	return p.call("getDataAttr", name).toString()
}

func (p *domStringMapImpl) Set(name string, value string) {
	p.call("setDataAttr", name, value)
}

func (p *domStringMapImpl) Delete(name string) {
	p.call("removeDataAttr", name)
}

// -------------8<---------------------------------------

type htmlOrSVGScriptElementImpl struct {
	Value
}

func wrapHTMLOrSVGScriptElement(v Value) HTMLOrSVGScriptElement {
	if v.valid() {
		return &htmlOrSVGScriptElementImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type nodeListImpl struct {
	Value
}

func wrapNodeList(v Value) NodeList {
	if p := newNodeListImpl(v); p != nil {
		return p
	}
	return nil
}

func newNodeListImpl(v Value) *nodeListImpl {
	if v.valid() {
		return &nodeListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeListImpl) Item(index uint) Node {
	return wrapAsNode(p.call("item", index))
}

func (p *nodeListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *nodeListImpl) Items() []Node {
	return nodeListToSlice(p.call("entries"))
}

// -------------8<---------------------------------------

type domParserImpl struct {
	Value
}

func newDOMParserImpl(v Value) DOMParser {
	if v.valid() {
		return &domParserImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domParserImpl) ParseFromString(str string, typ SupportedType) Document {
	return wrapDocument(p.call("parseFromString", str, string(typ)))
}

// -------------8<---------------------------------------

type xmlSerializerImpl struct {
	Value
}

func newXMLSerializerImpl(v Value) XMLSerializer {
	if v.valid() {
		return &xmlSerializerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *xmlSerializerImpl) SerializeToString(node Node) string {
	return p.call("serializeToString", JSValue(node)).toString()
}

// -------------8<---------------------------------------

func NewMutationObserver(cb MutationCallback) MutationObserver {
	if jsMutationObserver := jsGlobal.get("MutationObserver"); jsMutationObserver.valid() {
		return wrapMutationObserver(jsMutationObserver.jsNew(cb.jsCallback()))
	}
	return nil
}

// -------------8<---------------------------------------

func NewDOMParser() DOMParser {
	if v := jsGlobal.get("DOMParser"); v.valid() {
		return newDOMParserImpl(v)
	}
	return nil
}

// -------------8<---------------------------------------

func NewXMLSerializer() XMLSerializer {
	if v := jsGlobal.get("XMLSerializer"); v.valid() {
		return newXMLSerializerImpl(v)
	}
	return nil
}

// -------------8<---------------------------------------
