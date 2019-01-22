// +build js,wasm

package wasm

type (
	// https://dom.spec.whatwg.org/#document
	Document interface {
		Node
		NonElementParentNode
		DocumentOrShadowRoot
		ParentNode
		GeometryUtils
		GlobalEventHandlers
		DocumentAndElementEventHandlers

		Implementation() DOMImplementation
		URL() string
		DocumentURI() string
		Origin() string
		CompatMode() string
		CharacterSet() string
		ContentType() string
		DocType() DocumentType
		DocumentElement() Element
		ElementsByTagName(string) []Element
		ElementsByTagNameNS(string, string) []Element
		ElementsByClassName(string) []Element
		HTMLElementsByClassName(string) []HTMLElement
		CreateElement(string, ...ElementCreationOptions) Element
		CreateElementNS(string, string, ...ElementCreationOptions) Element
		CreateDocumentFragment() DocumentFragment
		CreateTextNode(string) Text
		CreateCDATASection(string) CDATASection
		CreateComment(string) Comment
		CreateProcessingInstruction(string, string) ProcessingInstruction
		ImportNode(Node, ...bool) Node
		AdoptNode(Node) Node
		CreateAttribute(string) Attr
		CreateAttributeNS(string, string) Attr
		CreateRange() Range
		CreateNodeIterator(Node, NodeFilterShow, ...NodeFilter) NodeIterator
		CreateTreeWalker(Node, NodeFilterShow, ...NodeFilter) TreeWalker

		// https://fullscreen.spec.whatwg.org/
		FullscreenEnabled() bool
		ExitFullscreen() func() error // Promise<void>
		OnFullscreenChange(func(Event)) EventHandler
		OnFullscreenError(func(Event)) EventHandler

		// https://www.w3.org/TR/html52/dom.html#elementdef-document
		Location() Location
		Domain() string
		SetDomain(string)
		Referrer() string
		Cookie() string
		SetCookie(string)
		LastModified() string
		ReadyState() DocumentReadyState
		//ByName(string) Value // ???
		// getter object (DOMString name);
		Title() string
		SetTitle(string)
		Dir() string
		SetDir(string)
		Body() HTMLBodyElement
		SetBody(HTMLBodyElement)
		Head() HTMLHeadElement
		Images() []HTMLImageElement
		Embeds() []HTMLEmbedElement
		Plugins() []HTMLEmbedElement
		Links() []HTMLElement
		Forms() []HTMLFormElement
		Scripts() []HTMLScriptElement
		ElementsByName(string) []Node
		CurrentScript() HTMLOrSVGScriptElement
		Open(...string) Document
		OpenURL(string, string, string, ...bool) WindowProxy
		Close()
		Write(...string)
		WriteLn(...string)
		DefaultView() WindowProxy
		ActiveElement() Element
		HasFocus() bool
		DesignMode() string
		SetDesignMode(string)
		ExecCommand(string, ...interface{}) bool
		QueryCommandEnabled(string) bool
		QueryCommandIndeterm(string) bool
		QueryCommandState(string) bool
		QueryCommandSupported(string) bool
		QueryCommandValue(string) string
		OnReadyStateChange(func(Event)) EventHandler

		// https://drafts.csswg.org/cssom-view/#extensions-to-the-document-interface
		ElementFromPoint(float64, float64) Element
		ElementsFromPoint(float64, float64) []Element
		CaretPositionFromPoint(float64, float64) CaretPosition
		ScrollingElement() Element

		// helpers
		CreateHTMLElement(string) HTMLElement
	}

	// https://dom.spec.whatwg.org/#domimplementation
	DOMImplementation interface {
		CreateDocumentType(string, string, string) DocumentType
		CreateDocument(string, string, ...DocumentType) XMLDocument
		CreateHTMLDocument(...string) Document
	}

	// https://dom.spec.whatwg.org/#xmldocument
	XMLDocument interface {
		Document
	}

	// https://dom.spec.whatwg.org/#treewalker
	TreeWalker interface {
		Root() Node
		WhatToShow() NodeFilterShow
		Filter() NodeFilter
		CurrentNode() Node
		SetCurrentNode(Node)
		ParentNode() Node
		FirstChild() Node
		LastChild() Node
		PreviousSibling() Node
		NextSibling() Node
		PreviousNode() Node
		NextNode() Node
	}

	// https://dom.spec.whatwg.org/#callbackdef-nodefilter
	NodeFilter interface {
		AcceptNode(Node) NodeFilterResult
	}

	// https://dom.spec.whatwg.org/#nodeiterator
	NodeIterator interface {
		Root() Node
		ReferenceNode() Node
		PointerBeforeReferenceNode() bool
		WhatToShow() NodeFilterShow
		Filter() NodeFilter
		NextNode() Node
		PreviousNode() Node
		Detach()
	}

	// https://dom.spec.whatwg.org/#range
	Range interface {
		AbstractRange

		CommonAncestorContainer() Node
		SetStart(Node, uint)
		SetEnd(Node, uint)
		SetStartBefore(Node)
		SetStartAfter(Node)
		SetEndBefore(Node)
		SetEndAfter(Node)
		Collapse(...bool)
		SelectNode(Node)
		SelectNodeContents(Node)
		CompareBoundaryPoints(RangeCompare, Range) int
		DeleteContents()
		ExtractContents() DocumentFragment
		CloneContents() DocumentFragment
		InsertNode(Node)
		SurroundContents(Node)
		CloneRange() Range
		Detach()
		IsPointInRange(Node, uint) bool
		ComparePoint(Node, uint) int
		IntersectsNode(Node) bool

		// https://drafts.csswg.org/cssom-view/#extensions-to-the-range-interface
		ClientRects() []DOMRect
		BoundingClientRect() DOMRect

		// https://www.w3.org/TR/DOM-Parsing/
		CreateContextualFragment(string) DocumentFragment
	}

	// https://dom.spec.whatwg.org/#abstractrange
	AbstractRange interface {
		StartContainer() Node
		StartOffset() int
		EndContainer() Node
		EndOffset() int
		Collapsed() bool
	}

	// https://dom.spec.whatwg.org/#staticrange
	StaticRange interface {
		AbstractRange
	}

	// https://dom.spec.whatwg.org/#processinginstruction
	ProcessingInstruction interface {
		CharacterData
		LinkStyle

		Target() string
	}

	// https://dom.spec.whatwg.org/#comment
	Comment interface {
		CharacterData
	}

	// https://dom.spec.whatwg.org/#cdatasection
	CDATASection interface {
		Text
	}

	// https://dom.spec.whatwg.org/#text
	Text interface {
		CharacterData
		Slotable
		GeometryUtils

		SplitText(int) Text
		WholeText() string
	}

	// https://dom.spec.whatwg.org/#characterdata
	CharacterData interface {
		Node
		NonDocumentTypeChildNode
		ChildNode

		Data() string
		SetData(string)
		Length() uint
		SubstringData(uint, uint) string
		AppendData(string)
		InsertData(uint, string)
		DeleteData(uint, uint)
		ReplaceData(uint, uint, string)
	}

	// https://dom.spec.whatwg.org/#documenttype
	DocumentType interface {
		Node
		ChildNode

		Name() string
		PublicId() string
		SystemId() string
	}

	// https://dom.spec.whatwg.org/#node
	Node interface {
		EventTarget

		NodeType() NodeType
		NodeName() string
		BaseURI() string
		IsConnected() bool
		OwnerDocument() Document
		RootNode(...RootNodeOptions) Node
		ParentNode() Node
		ParentElement() Element
		HasChildNodes() bool
		ChildNodes() []Node
		FirstChild() Node
		LastChild() Node
		PreviousSibling() Node
		NextSibling() Node
		NodeValue() string
		SetNodeValue(string)
		TextContent() string
		SetTextContent(string)
		Normalize()
		CloneNode(...bool) Node // deep bool = default false
		IsEqualNode(Node) bool
		IsSameNode(Node) bool
		CompareDocumentPosition(Node) DocumentPosition
		Contains(Node) bool
		LookupPrefix(string) string
		LookupNamespaceURI(string) string
		IsDefaultNamespace(string) bool
		InsertBefore(Node, Node) Node
		AppendChild(Node) Node
		ReplaceChild(Node, Node) Node
		RemoveChild(Node) Node
	}

	// https://dom.spec.whatwg.org/#element
	Element interface {
		Node
		ParentNode
		NonDocumentTypeChildNode
		ChildNode
		Slotable
		GeometryUtils

		NamespaceURI() string
		Prefix() string
		LocalName() string
		TagName() string
		Id() string
		SetId(string)
		ClassName() string
		SetClassName(string)
		ClassList() DOMTokenList
		Slot() string
		SetSlot(string)
		HasAttributes() bool
		Attributes() NamedNodeMap
		AttributeNames() []string
		Attribute(string) string
		AttributeNS(string, string) string
		SetAttribute(string, string)
		SetAttributeNS(string, string, string)
		RemoveAttribute(string)
		RemoveAttributeNS(string, string)
		ToggleAttribute(string, ...bool) bool
		HasAttribute(string) bool
		HasAttributeNS(string, string) bool
		AttributeNode(string) Attr
		AttributeNodeNS(string, string) Attr
		SetAttributeNode(Attr) Attr
		SetAttributeNodeNS(Attr) Attr
		RemoveAttributeNode(Attr) Attr
		AttachShadow(ShadowRootInit) ShadowRoot
		ShadowRoot() ShadowRoot
		Closest(string) Element
		Matches(string) bool
		ElementsByTagName(string) []Element
		ElementsByTagNameNS(string, string) []Element
		ElementsByClassName(string) []Element

		// https://www.w3.org/TR/cssom-view-1/#extension-to-the-element-interface
		ClientRects() []DOMRect
		BoundingClientRect() DOMRect
		ScrollIntoView(...interface{}) // empty, boolean or object
		Scroll(ScrollToOptions)
		ScrollTo(ScrollToOptions)
		ScrollBy(ScrollToOptions)
		ScrollTop() float64
		SetScrollTop(float64)
		ScrollLeft() float64
		SetScrollLeft(float64)
		ScrollWidth() int
		ScrollHeight() int
		ClientTop() int
		ClientLeft() int
		ClientWidth() int
		ClientHeight() int

		// https://www.w3.org/TR/DOM-Parsing/
		InnerHTML() string
		SetInnerHTML(string)
		OuterHTML() string
		SetOuterHTML(string)
		InsertAdjacentHTML(string, string)

		// https://fullscreen.spec.whatwg.org/
		RequestFullscreen(...FullscreenOptions) func() error // Promise<void>
		OnFullScreenChange(func(Event)) EventHandler
		OnFullScreenError(func(Event)) EventHandler

		// https://www.w3.org/TR/pointerevents/#extensions-to-the-element-interface
		SetPointerCapture(int)
		ReleasePointerCapture(int)
		HasPointerCapture(int) bool
	}

	// https://dom.spec.whatwg.org/#shadowroot
	ShadowRoot interface {
		DocumentFragment
		DocumentOrShadowRoot
		ParentNode

		Mode() ShadowRootMode
		Host() Element
	}

	// https://dom.spec.whatwg.org/#documentfragment
	DocumentFragment interface {
		Node
		NonElementParentNode
	}

	// https://dom.spec.whatwg.org/#domtokenlist
	DOMTokenList interface {
		Length() uint
		Item(uint) string
		Contains(string) bool
		Add(...string)
		Remove(...string)
		Toggle(string, ...bool) bool
		Replace(string, string) bool
		Supports(string) bool
		TokenValue() string
		SetTokenValue(string)
		TokenValues() []string
		String() string
	}

	//https://dom.spec.whatwg.org/#namednodemap
	NamedNodeMap interface {
		Length() uint
		Item(uint) Attr
		NamedItem(string) Attr
		NamedItemNS(string, string) Attr
		SetNamedItem(Attr) Attr
		SetNamedItemNS(Attr) Attr
		RemoveNamedItem(string) Attr
		RemoveNamedItemNS(string, string) Attr
	}

	// https://dom.spec.whatwg.org/#attr
	Attr interface {
		Node

		NamespaceURI() string
		Prefix() string
		LocalName() string
		Name() string
		Value() string
		SetValue(string)
		OwnerElement() Element
	}

	// https://dom.spec.whatwg.org/#htmlcollection
	HTMLCollection interface {
		Length() uint
		Item(uint) Element
		NamedItem(string) Element
	}

	// https://dom.spec.whatwg.org/#mutationrecord
	MutationRecord interface {
		Type() string
		Target() Node
		AddedNodes() []Node
		RemovedNodes() []Node
		PreviousSibling() Node
		NextSibling() Node
		AttributeName() string
		AttributeNamespace() string
		OldValue() string
	}

	// https://dom.spec.whatwg.org/#mutationobserver
	MutationObserver interface {
		Observe(Node, ...MutationObserverInit)
		Disconnect()
		TakeRecords() []MutationRecord
	}

	// https://www.w3.org/TR/html52/dom.html#typedefdef-htmlorsvgscriptelement
	// typedef (HTMLScriptElement or SVGScriptElement) HTMLOrSVGScriptElement;
	HTMLOrSVGScriptElement interface{}

	// https://html.spec.whatwg.org/multipage/dom.html#htmlelement
	HTMLUnknownElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/dom.html#htmlelement
	HTMLElement interface {
		Element
		GlobalEventHandlers
		DocumentAndElementEventHandlers

		Title() string
		SetTitle(string)
		Lang() string
		SetLang(string)
		Translate() bool
		SetTranslate(bool)
		Dir() string
		SetDir(string)
		DataSet() DOMStringMap
		Hidden() bool
		SetHidden(bool)
		Click()
		TabIndex() int
		SetTabIndex(int)
		Focus()
		Blur()
		AccessKey() string
		SetAccessKey(string)
		AccessKeyLabel() string //5.3
		Draggable() bool
		SetDraggable(bool)
		SpellCheck() bool
		SetSpellChack(bool)
		ForceSpellCheck()
		//Autocapitalize() string   //5.3
		//SetAutocapitalize(string) //5.3
		InnerText() string
		SetInnerText(string)

		// https://www.w3.org/TR/cssom-view-1/#extensions-to-the-htmlelement-interface
		OffsetParent() Element
		OffsetTop() int
		OffsetLeft() int
		OffsetWidth() int
		OffsetHeight() int

		// https://drafts.csswg.org/cssom/#elementcssinlinestyle
		Style() CSSStyleDeclaration
	}

	// https://www.w3.org/TR/html52/dom.html#domstringmap
	DOMStringMap interface {
		Get(string) string
		Set(string, string)
		Delete(string)
	}

	// https://dom.spec.whatwg.org/#nodelist
	NodeList interface {
		Item(uint) Node
		Length() uint
		Items() []Node
	}

	// https://www.w3.org/TR/DOM-Parsing/
	DOMParser interface {
		ParseFromString(string, SupportedType) Document
	}

	// https://www.w3.org/TR/DOM-Parsing/
	XMLSerializer interface {
		SerializeToString(Node) string
	}
)

type NodeType uint

const (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDATASection          NodeType = 4
	NodeTypeEntityReference       NodeType = 5
	NodeTypeEntity                NodeType = 6
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
	NodeTypeNotation              NodeType = 12
)

// https://dom.spec.whatwg.org/#dom-node-comparedocumentposition
type DocumentPosition uint

const (
	DocumentPositionDisconnected           DocumentPosition = 0x01
	DocumentPositionPreceding              DocumentPosition = 0x02
	DocumentPositionFollowing              DocumentPosition = 0x04
	DocumentPositionContains               DocumentPosition = 0x08
	DocumentPositionContainedBy            DocumentPosition = 0x10
	DocumentPositionImplementationSpecific DocumentPosition = 0x20
)

type ShadowRootMode string

const (
	ShadowRootModeOpen   ShadowRootMode = "open"
	ShadowRootModeClosed ShadowRootMode = "closed"
)

type RangeCompare uint16

const (
	RangeCompareStartToStart RangeCompare = 0
	RangeCompareStartToEnd   RangeCompare = 1
	RangeCompareEndToEnd     RangeCompare = 2
	RangeCompareEndToStart   RangeCompare = 3
)

// https://dom.spec.whatwg.org/#interface-nodefilter
type NodeFilterResult uint16

const (
	NodeFilterResultAccept NodeFilterResult = 1
	NodeFilterResultReject NodeFilterResult = 2
	NodeFilterResultSkip   NodeFilterResult = 3
)

// https://dom.spec.whatwg.org/#interface-nodefilter
type NodeFilterShow uint

const (
	NodeFilterShowAll                   NodeFilterShow = 0xFFFFFFFF
	NodeFilterShowElement               NodeFilterShow = 0x1
	NodeFilterShowAttribute             NodeFilterShow = 0x2
	NodeFilterShowText                  NodeFilterShow = 0x4
	NodeFilterShowCDATASection          NodeFilterShow = 0x8
	NodeFilterShowEntityReference       NodeFilterShow = 0x10
	NodeFilterShowEntity                NodeFilterShow = 0x20
	NodeFilterShowProcessingInstruction NodeFilterShow = 0x40
	NodeFilterShowComment               NodeFilterShow = 0x80
	NodeFilterShowDocument              NodeFilterShow = 0x100
	NodeFilterShowDocumentType          NodeFilterShow = 0x200
	NodeFilterShowDocumentFragment      NodeFilterShow = 0x400
	NodeFilterShowNotation              NodeFilterShow = 0x800
)

// https://www.w3.org/TR/DOM-Parsing/
type SupportedType string

const (
	SupportedType_Text_HTML             SupportedType = "text/html"
	SupportedType_Text_XML              SupportedType = "text/xml"
	SupportedType_Application_XML       SupportedType = "application/xml"
	SupportedType_Application_XHTML_XML SupportedType = "application/xhtml+xml"
	SupportedType_Image_SVG_XML         SupportedType = "image/svg+xml"
)

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-elementcreationoptions
type ElementCreationOptions struct {
	Is string
}

func (p ElementCreationOptions) JSValue() jsValue {
	o := jsObject.New()
	o.Set("is", p.Is)
	return o
}

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-getrootnodeoptions
type RootNodeOptions struct {
	Composed bool
}

func (p RootNodeOptions) JSValue() jsValue {
	o := jsObject.New()
	o.Set("composed", p.Composed)
	return o
}

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-shadowrootinit
type ShadowRootInit struct {
	Mode ShadowRootMode
}

func (p ShadowRootInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("mode", p.Mode)
	return o
}

// -------------8<---------------------------------------

// https://dom.spec.whatwg.org/#dictdef-mutationobserverinit
type MutationObserverInit struct {
	ChildList             bool
	Attributes            bool
	CharacterData         bool
	SubTree               bool
	AttributeOldValue     bool
	CharacterDataOldValue bool
	AttributeFilter       []string
}

func (p MutationObserverInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("childList", p.ChildList)
	o.Set("attributes", p.Attributes)
	o.Set("characterData", p.CharacterData)
	o.Set("subtree", p.SubTree)
	o.Set("attributeOldValue", p.AttributeOldValue)
	o.Set("characterDataOldValue", p.CharacterDataOldValue)
	o.Set("attributeFilter", ToJSArray(p.AttributeFilter))
	return o
}

// -------------8<---------------------------------------
