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
	if jsDoc := jsGlobal.Get("Document"); jsDoc.Valid() {
		return wrapDocument(jsDoc.New())
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
	if v.Valid() {
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
	return wrapDOMImplementation(p.Get("implementation"))
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
	return wrapDocumentType(p.Get("doctype"))
}

func (p *documentImpl) DocumentElement() Element {
	return wrapAsElement(p.Get("documentElement"))
}

func (p *documentImpl) ElementsByTagName(qualifiedName string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByTagName", qualifiedName))
}

func (p *documentImpl) ElementsByTagNameNS(namespace string, localName string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByTagNameNS", namespace, localName))
}

func (p *documentImpl) ElementsByClassName(classNames string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByClassName", classNames))
}

func (p *documentImpl) HTMLElementsByClassName(classNames string) []HTMLElement {
	return htmlCollectionToHTMLElementSlice(p.Call("getElementsByClassName", classNames))
}

func (p *documentImpl) CreateElement(localName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapAsElement(p.Call("createElement", localName))
	default:
		return wrapAsElement(p.Call("createElement", localName, options[0].toJSObject()))
	}
}

func (p *documentImpl) CreateElementNS(namespace string, qualifiedName string, options ...ElementCreationOptions) Element {
	switch len(options) {
	case 0:
		return wrapAsElement(p.Call("createElementNS", namespace, qualifiedName))
	default:
		return wrapAsElement(p.Call("createElementNS", namespace, qualifiedName, options[0].toJSObject()))
	}
}

func (p *documentImpl) CreateDocumentFragment() DocumentFragment {
	return wrapDocumentFragment(p.Call("createDocumentFragment"))
}

func (p *documentImpl) CreateTextNode(data string) Text {
	return wrapText(p.Call("createTextNode", data))
}

func (p *documentImpl) CreateCDATASection(data string) CDATASection {
	return wrapCDATASection(p.Call("createCDATASection", data))
}

func (p *documentImpl) CreateComment(data string) Comment {
	return wrapComment(p.Call("createComment", data))
}

func (p *documentImpl) CreateProcessingInstruction(target string, data string) ProcessingInstruction {
	return wrapProcessingInstruction(p.Call("createProcessingInstruction", target, data))
}

func (p *documentImpl) ImportNode(node Node, deep ...bool) Node {
	switch len(deep) {
	case 0:
		return wrapAsNode(p.Call("importNode", JSValue(node)))
	default:
		return wrapAsNode(p.Call("importNode", JSValue(node), deep[0]))
	}
}

func (p *documentImpl) AdoptNode(node Node) Node {
	return wrapAsNode(p.Call("adoptNode", JSValue(node)))
}

func (p *documentImpl) CreateAttribute(localName string) Attr {
	return wrapAttr(p.Call("adoptNode", localName))
}

func (p *documentImpl) CreateAttributeNS(namespace string, qualifiedName string) Attr {
	return wrapAttr(p.Call("createAttributeNS", namespace, qualifiedName))
}

func (p *documentImpl) CreateRange() Range {
	return wrapRange(p.Call("createRange"))
}

func (p *documentImpl) CreateNodeIterator(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) NodeIterator {
	switch len(filter) {
	case 0:
		return wrapNodeIterator(p.Call("createNodeIterator", JSValue(node), uint(whatToShow)))
	default:
		return wrapNodeIterator(p.Call("createNodeIterator", JSValue(node), uint(whatToShow), JSValue(filter[0])))
	}
}

func (p *documentImpl) CreateTreeWalker(node Node, whatToShow NodeFilterShow, filter ...NodeFilter) TreeWalker {
	switch len(filter) {
	case 0:
		return wrapTreeWalker(p.Call("createTreeWalker", JSValue(node), uint(whatToShow)))
	default:
		return wrapTreeWalker(p.Call("createTreeWalker", JSValue(node), uint(whatToShow), JSValue(filter[0])))
	}
}

func (p *documentImpl) FullscreenEnabled() bool {
	return p.Get("fullscreenEnabled").Bool()
}

func (p *documentImpl) ExitFullscreen() func() error {
	return func() error {
		result, ok := await(p.Call("exitFullscreen"))
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
	return wrapLocation(p.Get("location"))
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
func (p *documentImpl) ByName(string) Value {

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
	return wrapHTMLBodyElement(p.Get("body"))
}

func (p *documentImpl) SetBody(body HTMLBodyElement) {
	p.Set("body", JSValue(body))
}

func (p *documentImpl) Head() HTMLHeadElement {
	return wrapHTMLHeadElement(p.Get("head"))
}

func (p *documentImpl) Images() []HTMLImageElement {
	if c := wrapHTMLCollection(p.Get("images")); c != nil && c.Length() > 0 {
		var ret []HTMLImageElement
		for i := 0; i < c.Length(); i++ {
			if img, ok := c.Item(i).(HTMLImageElement); ok {
				ret = append(ret, img)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) Embeds() []HTMLEmbedElement {
	if c := wrapHTMLCollection(p.Get("embeds")); c != nil && c.Length() > 0 {
		var ret []HTMLEmbedElement
		for i := 0; i < c.Length(); i++ {
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
	return htmlCollectionToHTMLElementSlice(p.Get("links"))
}

func (p *documentImpl) Forms() []HTMLFormElement {
	if c := wrapHTMLCollection(p.Get("form")); c != nil && c.Length() > 0 {
		var ret []HTMLFormElement
		for i := 0; i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLFormElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) Scripts() []HTMLScriptElement {
	if c := wrapHTMLCollection(p.Get("scripts")); c != nil && c.Length() > 0 {
		var ret []HTMLScriptElement
		for i := 0; i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLScriptElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

func (p *documentImpl) ElementsByName(name string) []Node {
	return nodeListToSlice(p.Call("getElementsByName", name))
}

func (p *documentImpl) CurrentScript() HTMLOrSVGScriptElement {
	return wrapHTMLOrSVGScriptElement(p.Get("currentScript"))
}

func (p *documentImpl) Open(args ...string) Document {
	switch len(args) {
	case 1:
		return wrapDocument(p.Call("open", args[0]))
	case 2:
		return wrapDocument(p.Call("open", args[0], args[1]))
	default:
		return wrapDocument(p.Call("open"))
	}
}

func (p *documentImpl) OpenURL(url string, name string, features string, replace ...bool) WindowProxy {
	switch len(replace) {
	case 0:
		return wrapWindowProxy(p.Call("open", url, name, features))
	default:
		return wrapWindowProxy(p.Call("open", url, name, features, replace[0]))
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
	return wrapWindowProxy(p.Get("defaultView"))
}

func (p *documentImpl) ActiveElement() Element {
	return wrapAsElement(p.Get("activeElement"))
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
	return p.On("readystatechange", fn)
}

func (p *documentImpl) ElementFromPoint(x float64, y float64) Element {
	return wrapAsElement(p.Call("elementFromPoint", x, y))
}

func (p *documentImpl) ElementsFromPoint(x float64, y float64) []Element {
	var ret []Element

	sl := p.Call("elementsFromPoint", x, y).ToSlice()
	if sl != nil {
		for _, v := range sl {
			ret = append(ret, wrapAsElement(v))
		}
	}

	return ret
}

func (p *documentImpl) CaretPositionFromPoint(x float64, y float64) CaretPosition {
	return wrapCaretPosition(p.Call("caretPositionFromPoint", x, y))
}

func (p *documentImpl) ScrollingElement() Element {
	return wrapAsElement(p.Get("scrollingElement"))
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
	if v.Valid() {
		return &domImplementationImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domImplementationImpl) CreateDocumentType(qualifiedName string, publicId string, systemId string) DocumentType {
	return wrapDocumentType(p.Call("createDocumentType", qualifiedName, publicId, systemId))
}

func (p *domImplementationImpl) CreateDocument(namespace string, qualifiedName string, doctype ...DocumentType) XMLDocument {
	switch len(doctype) {
	case 0:
		return wrapXMLDocument(p.Call("createDocument", namespace, qualifiedName))
	default:
		return wrapXMLDocument(p.Call("createDocument", namespace, qualifiedName, JSValue(doctype[0])))
	}
}

func (p *domImplementationImpl) CreateHTMLDocument(title ...string) Document {
	if len(title) > 0 {
		return wrapDocument(p.Call("createHTMLDocument", title[0]))
	}
	return wrapDocument(p.Call("createHTMLDocument"))
}

// -------------8<---------------------------------------

type xmlDocumentImpl struct {
	*documentImpl
}

func wrapXMLDocument(v Value) XMLDocument {
	if v.Valid() {
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
	if v.Valid() {
		return &treeWalkerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *treeWalkerImpl) Root() Node {
	return wrapAsNode(p.Get("root"))
}

func (p *treeWalkerImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.Get("whatToShow").Int()))
}

func (p *treeWalkerImpl) Filter() NodeFilter {
	return wrapNodeFilter(p.Get("filter"))
}

func (p *treeWalkerImpl) CurrentNode() Node {
	return wrapAsNode(p.Get("currentNode"))
}

func (p *treeWalkerImpl) SetCurrentNode(node Node) {
	p.Set("currentNode", JSValue(node))
}

func (p *treeWalkerImpl) ParentNode() Node {
	return wrapAsNode(p.Call("parentNode"))
}

func (p *treeWalkerImpl) FirstChild() Node {
	return wrapAsNode(p.Call("firstChild"))
}

func (p *treeWalkerImpl) LastChild() Node {
	return wrapAsNode(p.Call("lastChild"))
}

func (p *treeWalkerImpl) PreviousSibling() Node {
	return wrapAsNode(p.Call("previousSibling"))
}

func (p *treeWalkerImpl) NextSibling() Node {
	return wrapAsNode(p.Call("nextSibling"))
}

func (p *treeWalkerImpl) PreviousNode() Node {
	return wrapAsNode(p.Call("previousNode"))
}

func (p *treeWalkerImpl) NextNode() Node {
	return wrapAsNode(p.Call("nextNode"))
}

// -------------8<---------------------------------------

type nodeFilterImpl struct {
	Value
}

func wrapNodeFilter(v Value) NodeFilter {
	if v.Valid() {
		return &nodeFilterImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeFilterImpl) AcceptNode(node Node) NodeFilterResult {
	return NodeFilterResult(p.Call("acceptNode", JSValue(node)).Int())
}

// -------------8<---------------------------------------

type nodeIteratorImpl struct {
	Value
}

func wrapNodeIterator(v Value) NodeIterator {
	if v.Valid() {
		return &nodeIteratorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeIteratorImpl) Root() Node {
	return wrapAsNode(p.Get("root"))
}

func (p *nodeIteratorImpl) ReferenceNode() Node {
	return wrapAsNode(p.Get("referenceNode"))
}

func (p *nodeIteratorImpl) PointerBeforeReferenceNode() bool {
	return p.Get("pointerBeforeReferenceNode").Bool()
}

func (p *nodeIteratorImpl) WhatToShow() NodeFilterShow {
	return NodeFilterShow(uint(p.Get("whatToShow").Int()))
}

func (p *nodeIteratorImpl) Filter() NodeFilter {
	return wrapNodeFilter(p.Get("filter"))
}

func (p *nodeIteratorImpl) NextNode() Node {
	return wrapAsNode(p.Call("nextNode"))
}

func (p *nodeIteratorImpl) PreviousNode() Node {
	return wrapAsNode(p.Call("previousNode"))
}

func (p *nodeIteratorImpl) Detach() {
	p.Call("detach")
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
	if v.Valid() {
		return &rangeImpl{
			abstractRangeImpl: newAbstractRangeImpl(v),
		}
	}
	return nil
}

func (p *rangeImpl) CommonAncestorContainer() Node {
	return wrapAsNode(p.Get("commonAncestorContainer"))
}

func (p *rangeImpl) SetStart(node Node, offset int) {
	p.Call("setStart", JSValue(node), offset)
}

func (p *rangeImpl) SetEnd(node Node, offset int) {
	p.Call("setEnd", JSValue(node), offset)
}

func (p *rangeImpl) SetStartBefore(node Node) {
	p.Call("setStartBefore", JSValue(node))
}

func (p *rangeImpl) SetStartAfter(node Node) {
	p.Call("setStartAfter", JSValue(node))
}

func (p *rangeImpl) SetEndBefore(node Node) {
	p.Call("setEndBefore", JSValue(node))
}

func (p *rangeImpl) SetEndAfter(node Node) {
	p.Call("setEndAfter", JSValue(node))
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
	p.Call("selectNode", JSValue(node))
}

func (p *rangeImpl) SelectNodeContents(node Node) {
	p.Call("selectNodeContents", JSValue(node))
}

func (p *rangeImpl) CompareBoundaryPoints(how RangeCompare, source Range) int {
	return p.Call("compareBoundaryPoints", int(how), JSValue(source)).Int()
}

func (p *rangeImpl) DeleteContents() {
	p.Call("deleteContents")
}

func (p *rangeImpl) ExtractContents() DocumentFragment {
	return wrapDocumentFragment(p.Call("extractContents"))
}

func (p *rangeImpl) CloneContents() DocumentFragment {
	return wrapDocumentFragment(p.Call("cloneContents"))
}

func (p *rangeImpl) InsertNode(node Node) {
	p.Call("insertNode", JSValue(node))
}

func (p *rangeImpl) SurroundContents(newParent Node) {
	p.Call("surroundContents", JSValue(newParent))
}

func (p *rangeImpl) CloneRange() Range {
	return wrapRange(p.Call("cloneRange"))
}

func (p *rangeImpl) Detach() {
	p.Call("detach")
}

func (p *rangeImpl) IsPointInRange(node Node, offset int) bool {
	return p.Call("isPointInRange", JSValue(node), offset).Bool()
}

func (p *rangeImpl) ComparePoint(node Node, offset int) int {
	return p.Call("comparePoint", JSValue(node), offset).Int()
}

func (p *rangeImpl) IntersectsNode(node Node) bool {
	return p.Call("intersectsNode", JSValue(node)).Bool()
}

func (p *rangeImpl) ClientRects() []DOMRect {
	if rects := p.Call("getClientRects").ToSlice(); rects != nil {
		var ret []DOMRect
		for _, rect := range rects {
			ret = append(ret, wrapDOMRect(rect))
		}
		return ret
	}
	return nil
}

func (p *rangeImpl) BoundingClientRect() DOMRect {
	return wrapDOMRect(p.Call("getBoundingClientRect"))
}

func (p *rangeImpl) CreateContextualFragment(fragment string) DocumentFragment {
	return wrapDocumentFragment(p.Call("createContextualFragment", fragment))
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
	if v.Valid() {
		return &abstractRangeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *abstractRangeImpl) StartContainer() Node {
	return wrapAsNode(p.Get("startContainer"))
}

func (p *abstractRangeImpl) StartOffset() int {
	return p.Get("startOffset").Int()
}

func (p *abstractRangeImpl) EndContainer() Node {
	return wrapAsNode(p.Get("endContainer"))
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

func wrapStaticRange(v Value) StaticRange {
	if v.Valid() {
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
	if v.Valid() {
		return &processingInstructionImpl{
			characterDataImpl: newCharacterDataImpl(v),
			linkStyleImpl:     newLinkStyleImpl(v),
			Value:             v,
		}
	}
	return nil
}

func (p *processingInstructionImpl) Target() string {
	return p.Get("target").String()
}

func (p *processingInstructionImpl) Length() int {
	return p.characterDataImpl.Length()
}

// -------------8<---------------------------------------

type commentImpl struct {
	*characterDataImpl
}

func wrapComment(v Value) Comment {
	if v.Valid() {
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
	if v.Valid() {
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
	if jsText := jsGlobal.Get("Text"); jsText.Valid() {
		switch len(data) {
		case 0:
			return wrapText(jsText.New())
		default:
			return wrapText(jsText.New(data[0]))
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
	if v.Valid() {
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
	return wrapText(p.Call("splitText", offset))
}

func (p *textImpl) WholeText() string {
	return p.Get("wholeText").String()
}

func (p *textImpl) Length() int {
	return p.Get("length").Int()
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
	if v.Valid() {
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
	Value
}

func wrapDocumentType(v Value) DocumentType {
	if v.Valid() {
		return &documentTypeImpl{
			nodeImpl:      newNodeImpl(v),
			childNodeImpl: newChildNodeImpl(v),
			Value:         v,
		}
	}
	return nil
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

func wrapNode(v Value) Node {
	if p := newNodeImpl(v); p != nil {
		return p
	}
	return nil
}

func newNodeImpl(v Value) *nodeImpl {
	if v.Valid() {
		return &nodeImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
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
	return wrapDocument(p.Get("ownerDocument"))
}

func (p *nodeImpl) RootNode(options ...RootNodeOptions) Node {
	if len(options) > 0 {
		return wrapAsNode(p.Call("getRootNode", options[0].toJSObject()))
	}

	return wrapAsNode(p.Call("getRootNode"))
}

func (p *nodeImpl) ParentNode() Node {
	return wrapAsNode(p.Get("parentNode"))
}

func (p *nodeImpl) ParentElement() Element {
	return wrapAsElement(p.Get("parentElement"))
}

func (p *nodeImpl) HasChildNodes() bool {
	return p.Call("hasChildNodes").Bool()
}

func (p *nodeImpl) ChildNodes() []Node {
	return nodeListToSlice(p.Get("childNodes"))
}

func (p *nodeImpl) FirstChild() Node {
	return wrapAsNode(p.Get("firstChild"))
}

func (p *nodeImpl) LastChild() Node {
	return wrapAsNode(p.Get("lastChild"))
}

func (p *nodeImpl) PreviousSibling() Node {
	return wrapAsNode(p.Get("previousSibling"))
}

func (p *nodeImpl) NextSibling() Node {
	return wrapAsNode(p.Get("nextSibling"))
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
		return wrapAsNode(p.Call("cloneNode", deep[0]))
	}
	return wrapAsNode(p.Call("cloneNode"))
}

func (p *nodeImpl) IsEqualNode(otherNode Node) bool {
	return p.Call("isEqualNode", JSValue(otherNode)).Bool()
}

func (p *nodeImpl) IsSameNode(otherNode Node) bool {
	return p.Call("isSameNode", JSValue(otherNode)).Bool()
}

func (p *nodeImpl) CompareDocumentPosition(other Node) DocumentPosition {
	return DocumentPosition(p.Call("compareDocumentPosition", JSValue(other)).Int())
}

func (p *nodeImpl) Contains(other Node) bool {
	return p.Call("contains", JSValue(other)).Bool()
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
		return wrapAsNode(p.Call("insertBefore", JSValue(node), JSValue(child)))
	}
	return wrapAsNode(p.Call("insertBefore", JSValue(node)))
}

func (p *nodeImpl) AppendChild(node Node) Node {
	return wrapAsNode(p.Call("appendChild", JSValue(node)))
}

func (p *nodeImpl) ReplaceChild(node Node, child Node) Node {
	return wrapAsNode(p.Call("replaceChild", JSValue(node), JSValue(child)))
}

func (p *nodeImpl) RemoveChild(child Node) Node {
	return wrapAsNode(p.Call("removeChild", JSValue(child)))
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
	if v.Valid() {
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
	return wrapDOMTokenList(p.Get("classList"))
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
	return wrapNamedNodeMap(p.Get("attributes"))
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
	return wrapAttr(p.Call("getAttributeNode", name))
}

func (p *elementImpl) AttributeNodeNS(namespace string, name string) Attr {
	return wrapAttr(p.Call("getAttributeNodeNS", namespace, name))
}

func (p *elementImpl) SetAttributeNode(attr Attr) Attr {
	return wrapAttr(p.Call("setAttributeNode", JSValue(attr)))
}

func (p *elementImpl) SetAttributeNodeNS(attr Attr) Attr {
	return wrapAttr(p.Call("setAttributeNodeNS", JSValue(attr)))
}

func (p *elementImpl) RemoveAttributeNode(attr Attr) Attr {
	return wrapAttr(p.Call("removeAttributeNode", JSValue(attr)))
}

func (p *elementImpl) AttachShadow(si ShadowRootInit) ShadowRoot {
	return wrapShadowRoot(p.Call("attachShadow", si.toJSObject()))
}

func (p *elementImpl) ShadowRoot() ShadowRoot {
	return wrapShadowRoot(p.Get("shadowRoot"))
}

func (p *elementImpl) Closest(selectors string) Element {
	return wrapAsElement(p.Call("closest"))
}

func (p *elementImpl) Matches(selector string) bool {
	return p.Call("matches", selector).Bool()
}

func (p *elementImpl) ElementsByTagName(name string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByTagName", name))
}

func (p *elementImpl) ElementsByTagNameNS(namespace string, localName string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByTagNameNS", namespace, localName))
}

func (p *elementImpl) ElementsByClassName(names string) []Element {
	return htmlCollectionToElementSlice(p.Call("getElementsByClassName", names))
}

func (p *elementImpl) ClientRects() []DOMRect {
	if rects := p.Call("getClientRects").ToSlice(); rects != nil {
		var ret []DOMRect
		for _, rect := range rects {
			ret = append(ret, wrapDOMRect(rect))
		}
		return ret
	}
	return nil
}

func (p *elementImpl) BoundingClientRect() DOMRect {
	return wrapDOMRect(p.Call("getBoundingClientRect"))
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
			p.Call("scrollIntoView", x.toJSObject())
		}
	}
}

func (p *elementImpl) Scroll(options ScrollToOptions) {
	p.Call("scroll", options.toJSObject())
}

func (p *elementImpl) ScrollTo(options ScrollToOptions) {
	p.Call("scrollTo", options.toJSObject())
}

func (p *elementImpl) ScrollBy(options ScrollToOptions) {
	p.Call("scrollBy", options.toJSObject())
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

func (p *elementImpl) RequestFullscreen(...FullscreenOptions) func() error {
	return func() error {
		res, ok := await(p.Call("requestFullscreen"))
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
	p.Call("setPointerCapture", pointerId)
}

func (p *elementImpl) ReleasePointerCapture(pointerId int) {
	p.Call("releasePointerCapture", pointerId)
}

func (p *elementImpl) HasPointerCapture(pointerId int) bool {
	return p.Call("hasPointerCapture", pointerId).Bool()
}

// -------------8<---------------------------------------

type shadowRootImpl struct {
	*documentFragmentImpl
	*documentOrShadowRootImpl
	*parentNodeImpl
	Value
}

func wrapShadowRoot(v Value) ShadowRoot {
	if v.Valid() {
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
	return ShadowRootMode(p.Get("mode").String())
}

func (p *shadowRootImpl) Host() Element {
	return wrapAsElement(p.Get("host"))
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
	if v.Valid() {
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
	if v.Valid() {
		return &domTokenListImpl{
			Value: v,
		}
	}
	return nil
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
	Value
}

func wrapNamedNodeMap(v Value) NamedNodeMap {
	if v.Valid() {
		return &namedNodeMapImpl{
			Value: v,
		}
	}
	return nil
}

func (p *namedNodeMapImpl) Length() int {
	return p.Get("length").Int()
}

func (p *namedNodeMapImpl) Item(index int) Attr {
	return wrapAttr(p.Call("item", index))
}

func (p *namedNodeMapImpl) NamedItem(name string) Attr {
	return wrapAttr(p.Get("getNamedItem"))
}

func (p *namedNodeMapImpl) NamedItemNS(namespace string, name string) Attr {
	return wrapAttr(p.Call("getNamedItemNS", namespace, name))
}

func (p *namedNodeMapImpl) SetNamedItem(attr Attr) Attr {
	return wrapAttr(p.Call("setNamedItem", JSValue(attr)))
}

func (p *namedNodeMapImpl) SetNamedItemNS(attr Attr) Attr {
	return wrapAttr(p.Call("setNamedItemNS", JSValue(attr)))
}

func (p *namedNodeMapImpl) RemoveNamedItem(name string) Attr {
	return wrapAttr(p.Call("removeNamedItem", name))
}

func (p *namedNodeMapImpl) RemoveNamedItemNS(namespace string, name string) Attr {
	return wrapAttr(p.Call("removeNamedItemNS", namespace, name))
}

// -------------8<---------------------------------------

type attrImpl struct {
	*nodeImpl
}

func wrapAttr(v Value) Attr {
	if v.Valid() {
		return &attrImpl{
			nodeImpl: newNodeImpl(v),
		}
	}
	return nil
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
	return wrapAsElement(p.Get("ownerElement"))
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
	if v.Valid() {
		return &htmlCollectionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *htmlCollectionImpl) Length() int {
	return p.Get("length").Int()
}

func (p *htmlCollectionImpl) Item(index int) Element {
	return wrapAsElement(p.Call("item", index))
}

func (p *htmlCollectionImpl) NamedItem(name string) Element {
	return wrapAsElement(p.Call("namedItem", name))
}

// -------------8<---------------------------------------

type mutationRecordImpl struct {
	Value
}

func wrapMutationRecord(v Value) MutationRecord {
	if v.Valid() {
		return &mutationRecordImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mutationRecordImpl) Type() string {
	return p.Get("type").String()
}

func (p *mutationRecordImpl) Target() Node {
	return wrapAsNode(p.Get("target"))
}

func (p *mutationRecordImpl) AddedNodes() []Node {
	return nodeListToSlice(p.Get("addedNodes"))
}

func (p *mutationRecordImpl) RemovedNodes() []Node {
	return nodeListToSlice(p.Get("removedNodes"))
}

func (p *mutationRecordImpl) PreviousSibling() Node {
	return wrapAsNode(p.Get("previousSibling"))
}

func (p *mutationRecordImpl) NextSibling() Node {
	return wrapAsNode(p.Get("nextSibling"))
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
	Value
}

func wrapMutationObserver(v Value) MutationObserver {
	if v.Valid() {
		return &mutationObserverImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mutationObserverImpl) Observe(target Node, options ...MutationObserverInit) {
	switch len(options) {
	case 0:
		p.Call("observe", JSValue(target))
	default:
		p.Call("observe", JSValue(target), options[0].toJSObject())
	}
}

func (p *mutationObserverImpl) Disconnect() {
	p.Call("disconnect")
}

func (p *mutationObserverImpl) TakeRecords() []MutationRecord {
	if s := p.Call("takeRecords").ToSlice(); s != nil {
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
	if v.Valid() {
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
	if v.Valid() {
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
	return wrapDOMStringMap(p.Get("dataset"))
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
	return wrapAsElement(p.Get("offsetParent"))
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

func (p *htmlElementImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type domStringMapImpl struct {
	Value
}

func wrapDOMStringMap(v Value) DOMStringMap {
	if v.Valid() {
		return &domStringMapImpl{
			Value: v,
		}
	}
	return nil
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
	Value
}

func wrapHTMLOrSVGScriptElement(v Value) HTMLOrSVGScriptElement {
	if v.Valid() {
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
	if v.Valid() {
		return &nodeListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *nodeListImpl) Item(index int) Node {
	return wrapAsNode(p.Call("item", index))
}

func (p *nodeListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *nodeListImpl) Items() []Node {
	return nodeListToSlice(p.Call("entries"))
}

// -------------8<---------------------------------------

type domParserImpl struct {
	Value
}

func newDOMParserImpl(v Value) DOMParser {
	if v.Valid() {
		return &domParserImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domParserImpl) ParseFromString(str string, typ SupportedType) Document {
	return wrapDocument(p.Call("parseFromString", str, string(typ)))
}

// -------------8<---------------------------------------

type xmlSerializerImpl struct {
	Value
}

func newXMLSerializerImpl(v Value) XMLSerializer {
	if v.Valid() {
		return &xmlSerializerImpl{
			Value: v,
		}
	}
	return nil
}

func (p *xmlSerializerImpl) SerializeToString(node Node) string {
	return p.Call("serializeToString", JSValue(node)).String()
}

// -------------8<---------------------------------------

func NewMutationObserver(cb MutationCallback) MutationObserver {
	if jsMutationObserver := jsGlobal.Get("MutationObserver"); jsMutationObserver.Valid() {
		return wrapMutationObserver(jsMutationObserver.New(cb.jsCallback()))
	}
	return nil
}

// -------------8<---------------------------------------

func NewDOMParser() DOMParser {
	if v := jsGlobal.Get("DOMParser"); v.Valid() {
		return newDOMParserImpl(v)
	}
	return nil
}

// -------------8<---------------------------------------

func NewXMLSerializer() XMLSerializer {
	if v := jsGlobal.Get("XMLSerializer"); v.Valid() {
		return newXMLSerializerImpl(v)
	}
	return nil
}

// -------------8<---------------------------------------
