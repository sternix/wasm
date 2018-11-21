// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

func NewAudio(src ...string) HTMLAudioElement {
	jsHTMLAudioElement := js.Global().Get("HTMLAudioElement")
	if isNil(jsHTMLAudioElement) {
		return nil
	}

	switch len(src) {
	case 0:
		return newHTMLAudioElement(jsHTMLAudioElement.New())
	default:
		return newHTMLAudioElement(jsHTMLAudioElement.New(src[0]))
	}
}

func NewImage(args ...uint) HTMLImageElement {
	jsHTMLImageElement := js.Global().Get("HTMLImageElement")
	if isNil(jsHTMLImageElement) {
		return nil
	}

	switch len(args) {
	case 0:
		return newHTMLImageElement(jsHTMLImageElement.New())
	case 1:
		return newHTMLImageElement(jsHTMLImageElement.New(args[0]))
	default:
		return newHTMLImageElement(jsHTMLImageElement.New(args[0], args[1]))
	}
}

func NewMediaStream(args ...interface{}) MediaStream {
	jsMediaStream := js.Global().Get("MediaStream")
	if isNil(jsMediaStream) {
		return nil
	}

	if len(args) > 0 {
		switch x := args[0].(type) {
		case MediaStream:
			return newMediaStream(jsMediaStream.New(x.JSValue()))
		case []MediaStreamTrack:
			var s []js.Value
			for _, m := range x {
				s = append(s, m.JSValue())
			}
			return newMediaStream(jsMediaStream.New(sliceToJsArray(s)))
		}
	}

	return newMediaStream(jsMediaStream.New())
}

// -------------8<---------------------------------------

type htmlBodyElementImpl struct {
	*htmlElementImpl
	*windowEventHandlersImpl
	js.Value
}

func newHTMLBodyElement(v js.Value) HTMLBodyElement {
	if isNil(v) {
		return nil
	}

	return &htmlBodyElementImpl{
		htmlElementImpl:         newHTMLElementImpl(v),
		windowEventHandlersImpl: newWindowEventHandlersImpl(v),
		Value:                   v,
	}
}

// -------------8<---------------------------------------

type linkStyleImpl struct {
	js.Value
}

func newLinkStyle(v js.Value) LinkStyle {
	if p := newLinkStyleImpl(v); p != nil {
		return p
	}
	return nil
}

func newLinkStyleImpl(v js.Value) *linkStyleImpl {
	if isNil(v) {
		return nil
	}

	return &linkStyleImpl{
		Value: v,
	}
}

func (p *linkStyleImpl) Sheet() StyleSheet {
	return newStyleSheet(p.Get("sheet"))
}

func (p *linkStyleImpl) SetSheet(sheet StyleSheet) {
	p.Set("sheet", sheet.JSValue())
}

// -------------8<---------------------------------------

type styleSheetImpl struct {
	js.Value
}

func newStyleSheet(v js.Value) StyleSheet {
	if isNil(v) {
		return nil
	}

	return &styleSheetImpl{
		Value: v,
	}
}

func (p *styleSheetImpl) Type() string {
	return p.Get("type").String()
}

func (p *styleSheetImpl) Href() string {
	return p.Get("href").String()
}

func (p *styleSheetImpl) OwnerNode() Node {
	return newNode(p.Get("ownerNode"))
}

func (p *styleSheetImpl) ParentStyleSheet() StyleSheet {
	return newStyleSheet(p.Get("parentStyleSheet"))
}

func (p *styleSheetImpl) Title() string {
	return p.Get("title").String()
}

func (p *styleSheetImpl) Media() MediaList {
	return newMediaList(p.Get("media"))
}

func (p *styleSheetImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *styleSheetImpl) SetDisabled(d bool) {
	p.Set("disabled", d)
}

// -------------8<---------------------------------------

type mediaListImpl struct {
	js.Value
}

func newMediaList(v js.Value) MediaList {
	if isNil(v) {
		return nil
	}

	return &mediaListImpl{
		Value: v,
	}
}

func (p *mediaListImpl) MediaText() string {
	return p.Get("mediaText").String()
}

func (p *mediaListImpl) SetMediaText(text string) {
	p.Set("mediaText", text)
}

func (p *mediaListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *mediaListImpl) Item(index int) string {
	return p.Call("item", index).String()
}

func (p *mediaListImpl) AppendMedium(medium string) {
	p.Call("appendMedium", medium)
}

func (p *mediaListImpl) DeleteMedium(medium string) {
	p.Call("deleteMedium", medium)
}

// -------------8<---------------------------------------

type htmlHeadingElementImpl struct {
	*htmlElementImpl
}

func newHTMLHeadingElement(v js.Value) HTMLHeadingElement {
	if isNil(v) {
		return nil
	}

	return &htmlHeadingElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlParagraphElementImpl struct {
	*htmlElementImpl
}

func newHTMLParagraphElement(v js.Value) HTMLParagraphElement {
	if isNil(v) {
		return nil
	}

	return &htmlParagraphElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlHRElementImpl struct {
	*htmlElementImpl
}

func newHTMLHRElement(v js.Value) HTMLHRElement {
	if isNil(v) {
		return nil
	}

	return &htmlHRElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlPreElementImpl struct {
	*htmlElementImpl
}

func newHTMLPreElement(v js.Value) HTMLPreElement {
	if isNil(v) {
		return nil
	}

	return &htmlPreElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlQuoteElementImpl struct {
	*htmlElementImpl
}

func newHTMLQuoteElement(v js.Value) HTMLQuoteElement {
	if isNil(v) {
		return nil
	}

	return &htmlQuoteElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlQuoteElementImpl) Cite() string {
	return p.Get("cite").String()
}

func (p *htmlQuoteElementImpl) SetCite(cite string) {
	p.Set("cite", cite)
}

// -------------8<---------------------------------------

type htmlOListElementImpl struct {
	*htmlElementImpl
}

func newHTMLOListElement(v js.Value) HTMLOListElement {
	if isNil(v) {
		return nil
	}

	return &htmlOListElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlOListElementImpl) Reversed() bool {
	return p.Get("reversed").Bool()
}

func (p *htmlOListElementImpl) SetReversed(r bool) {
	p.Set("reversed", r)
}

func (p *htmlOListElementImpl) Start() int {
	return p.Get("start").Int()
}

func (p *htmlOListElementImpl) SetStart(s int) {
	p.Set("start", s)
}

func (p *htmlOListElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlOListElementImpl) SetType(t string) {
	p.Set("type", t)
}

// -------------8<---------------------------------------

type htmlUListElementImpl struct {
	*htmlElementImpl
}

func newHTMLUListElement(v js.Value) HTMLUListElement {
	if isNil(v) {
		return nil
	}

	return &htmlUListElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlLIElementImpl struct {
	*htmlElementImpl
}

func newHTMLLIElement(v js.Value) HTMLLIElement {
	if isNil(v) {
		return nil
	}

	return &htmlLIElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlLIElementImpl) Value() int {
	return p.Get("value").Int()
}

func (p *htmlLIElementImpl) SetValue(v int) {
	p.Set("value", v)
}

// -------------8<---------------------------------------

type htmlDListElementImpl struct {
	*htmlElementImpl
}

func newHTMLDListElement(v js.Value) HTMLDListElement {
	if isNil(v) {
		return nil
	}

	return &htmlDListElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlDivElementImpl struct {
	*htmlElementImpl
}

func newHTMLDivElement(v js.Value) HTMLDivElement {
	if isNil(v) {
		return nil
	}

	return &htmlDivElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlAnchorElementImpl struct {
	*htmlElementImpl
	*htmlHyperlinkElementUtilsImpl
	js.Value
}

func newHTMLAnchorElement(v js.Value) HTMLAnchorElement {
	if isNil(v) {
		return nil
	}

	return &htmlAnchorElementImpl{
		htmlElementImpl:               newHTMLElementImpl(v),
		htmlHyperlinkElementUtilsImpl: newHTMLHyperlinkElementUtilsImpl(v),
		Value:                         v,
	}
}

func (p *htmlAnchorElementImpl) Target() string {
	return p.Get("target").String()
}

func (p *htmlAnchorElementImpl) SetTarget(t string) {
	p.Set("target", t)
}

func (p *htmlAnchorElementImpl) Download() string {
	return p.Get("download").String()
}

func (p *htmlAnchorElementImpl) SetDownload(d string) {
	p.Set("download", d)
}

func (p *htmlAnchorElementImpl) Rel() string {
	return p.Get("rel").String()
}

func (p *htmlAnchorElementImpl) SetRel(r string) {
	p.Set("rel", r)
}

func (p *htmlAnchorElementImpl) Rev() string {
	return p.Get("rev").String()
}

func (p *htmlAnchorElementImpl) SetRev(r string) {
	p.Set("rev", r)
}

func (p *htmlAnchorElementImpl) RelList() DOMTokenList {
	return newDOMTokenList(p.Get("relList"))
}

func (p *htmlAnchorElementImpl) HrefLang() string {
	return p.Get("hreflang").String()
}

func (p *htmlAnchorElementImpl) SetHrefLang(l string) {
	p.Set("hreflang", l)
}

func (p *htmlAnchorElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlAnchorElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlAnchorElementImpl) Text() string {
	return p.Get("text").String()
}

func (p *htmlAnchorElementImpl) SetText(text string) {
	p.Set("text", text)
}

func (p *htmlAnchorElementImpl) ReferrerPolicy() string {
	return p.Get("referrerPolicy").String()
}

func (p *htmlAnchorElementImpl) SetReferrerPolicy(policy string) {
	p.Set("referrerPolicy", policy)
}

// -------------8<---------------------------------------

type htmlHyperlinkElementUtilsImpl struct {
	js.Value
}

func newHTMLHyperlinkElementUtils(v js.Value) HTMLHyperlinkElementUtils {
	if p := newHTMLHyperlinkElementUtilsImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLHyperlinkElementUtilsImpl(v js.Value) *htmlHyperlinkElementUtilsImpl {
	if isNil(v) {
		return nil
	}

	return &htmlHyperlinkElementUtilsImpl{
		Value: v,
	}
}

func (p *htmlHyperlinkElementUtilsImpl) Href() string {
	return p.Get("href").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHref(href string) {
	p.Set("href", href)
}

func (p *htmlHyperlinkElementUtilsImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *htmlHyperlinkElementUtilsImpl) Protocol() string {
	return p.Get("protocol").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetProtocol(protocol string) {
	p.Set("protocol", protocol)
}

func (p *htmlHyperlinkElementUtilsImpl) Username() string {
	return p.Get("username").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetUsername(username string) {
	p.Set("username", username)
}

func (p *htmlHyperlinkElementUtilsImpl) Password() string {
	return p.Get("password").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPassword(password string) {
	p.Set("password", password)
}

func (p *htmlHyperlinkElementUtilsImpl) Host() string {
	return p.Get("host").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHost(host string) {
	p.Set("host", host)
}

func (p *htmlHyperlinkElementUtilsImpl) Hostname() string {
	return p.Get("hostname").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHostname(hostname string) {
	p.Set("hostname", hostname)
}

func (p *htmlHyperlinkElementUtilsImpl) Port() string {
	return p.Get("port").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPort(port string) {
	p.Set("port", port)
}

func (p *htmlHyperlinkElementUtilsImpl) Pathname() string {
	return p.Get("pathname").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPathname(pathname string) {
	p.Set("pathname", pathname)
}

func (p *htmlHyperlinkElementUtilsImpl) Search() string {
	return p.Get("search").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetSearch(search string) {
	p.Set("search", search)
}

func (p *htmlHyperlinkElementUtilsImpl) Hash() string {
	return p.Get("hash").String()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHash(hash string) {
	p.Set("hash", hash)
}

// -------------8<---------------------------------------

type htmlDataElementImpl struct {
	*htmlElementImpl
}

func newHTMLDataElement(v js.Value) HTMLDataElement {
	if isNil(v) {
		return nil
	}

	return &htmlDataElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlDataElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlDataElementImpl) SetValue(value string) {
	p.Set("value", value)
}

// -------------8<---------------------------------------

type htmlTimeElementImpl struct {
	*htmlElementImpl
}

func newHTMLTimeElement(v js.Value) HTMLTimeElement {
	if isNil(v) {
		return nil
	}

	return &htmlTimeElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTimeElementImpl) DateTime() string {
	return p.Get("dateTime").String()
}

func (p *htmlTimeElementImpl) SetDateTime(dt string) {
	p.Set("dateTime", dt)
}

// -------------8<---------------------------------------

type htmlSpanElementImpl struct {
	*htmlElementImpl
}

func newHTMLSpanElement(v js.Value) HTMLSpanElement {
	if isNil(v) {
		return nil
	}

	return &htmlSpanElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlBRElementImpl struct {
	*htmlElementImpl
}

func newHTMLBRElement(v js.Value) HTMLBRElement {
	if isNil(v) {
		return nil
	}

	return &htmlBRElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlModElementImpl struct {
	*htmlElementImpl
}

func newHTMLModElement(v js.Value) HTMLModElement {
	if isNil(v) {
		return nil
	}

	return &htmlModElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlModElementImpl) Cite() string {
	return p.Get("cite").String()
}

func (p *htmlModElementImpl) SetCite(cite string) {
	p.Set("cite", cite)
}

func (p *htmlModElementImpl) DateTime() string {
	return p.Get("dateTime").String()
}

func (p *htmlModElementImpl) SetDateTime(dt string) {
	p.Set("dateTime", dt)
}

// -------------8<---------------------------------------

type htmlPictureElementImpl struct {
	*htmlElementImpl
}

func newHTMLPictureElement(v js.Value) HTMLPictureElement {
	if isNil(v) {
		return nil
	}

	return &htmlPictureElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlSourceElementImpl struct {
	*htmlElementImpl
}

func newHTMLSourceElement(v js.Value) HTMLSourceElement {
	if isNil(v) {
		return nil
	}

	return &htmlSourceElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlSourceElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlSourceElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlSourceElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlSourceElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlSourceElementImpl) SrcSet() string {
	return p.Get("srcset").String()
}

func (p *htmlSourceElementImpl) SetSrcSet(srcset string) {
	p.Set("srcset", srcset)
}

func (p *htmlSourceElementImpl) Sizes() string {
	return p.Get("sizes").String()
}

func (p *htmlSourceElementImpl) SetSizes(sizes string) {
	p.Set("sizes", sizes)
}

func (p *htmlSourceElementImpl) Media() string {
	return p.Get("media").String()
}

func (p *htmlSourceElementImpl) SetMedia(media string) {
	p.Set("media", media)
}

// -------------8<---------------------------------------

type htmlImageElementImpl struct {
	*htmlElementImpl
}

func newHTMLImageElement(v js.Value) HTMLImageElement {
	if isNil(v) {
		return nil
	}

	return &htmlImageElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlImageElementImpl) Alt() string {
	return p.Get("alt").String()
}

func (p *htmlImageElementImpl) SetAlt(alt string) {
	p.Set("alt", alt)
}

func (p *htmlImageElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlImageElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlImageElementImpl) SrcSet() string {
	return p.Get("srcset").String()
}

func (p *htmlImageElementImpl) SetSrcSet(srcset string) {
	p.Set("srcset", srcset)
}

func (p *htmlImageElementImpl) Sizes() string {
	return p.Get("sizes").String()
}

func (p *htmlImageElementImpl) SetSizes(sizes string) {
	p.Set("sizes", sizes)
}

func (p *htmlImageElementImpl) CrossOrigin() string {
	return p.Get("crossOrigin").String()
}

func (p *htmlImageElementImpl) SetCrossOrigin(co string) {
	p.Set("crossOrigin", co)
}

func (p *htmlImageElementImpl) UseMap() string {
	return p.Get("useMap").String()
}

func (p *htmlImageElementImpl) SetUseMap(um string) {
	p.Set("useMap", um)
}

func (p *htmlImageElementImpl) LongDesc() string {
	return p.Get("longDesc").String()
}

func (p *htmlImageElementImpl) SetLongDesc(ld string) {
	p.Set("longDesc", ld)
}

func (p *htmlImageElementImpl) IsMap() bool {
	return p.Get("isMap").Bool()
}

func (p *htmlImageElementImpl) SetIsMap(b bool) {
	p.Set("isMap", b)
}

func (p *htmlImageElementImpl) Width() int {
	return p.Get("width").Int()
}

func (p *htmlImageElementImpl) SetWidth(w int) {
	p.Set("width", w)
}

func (p *htmlImageElementImpl) Height() int {
	return p.Get("height").Int()
}

func (p *htmlImageElementImpl) SetHeight(h int) {
	p.Set("height", h)
}

func (p *htmlImageElementImpl) NaturalWidth() int {
	return p.Get("naturalWidth").Int()
}

func (p *htmlImageElementImpl) NaturalHeight() int {
	return p.Get("naturalHeight").Int()
}

func (p *htmlImageElementImpl) Complete() bool {
	return p.Get("complete").Bool()
}

func (p *htmlImageElementImpl) CurrentSrc() string {
	return p.Get("currentSrc").String()
}

func (p *htmlImageElementImpl) ReferrerPolicy() string {
	return p.Get("referrerPolicy").String()
}

func (p *htmlImageElementImpl) SetReferrerPolicy(policy string) {
	p.Set("referrerPolicy", policy)
}

func (p *htmlImageElementImpl) X() int {
	return p.Get("x").Int()
}

func (p *htmlImageElementImpl) Y() int {
	return p.Get("y").Int()
}

// -------------8<---------------------------------------

type htmlIFrameElementImpl struct {
	*htmlElementImpl
}

func newHTMLIFrameElement(v js.Value) HTMLIFrameElement {
	if isNil(v) {
		return nil
	}

	return &htmlIFrameElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlIFrameElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlIFrameElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlIFrameElementImpl) SrcDoc() string {
	return p.Get("srcdoc").String()
}

func (p *htmlIFrameElementImpl) SetSrcDoc(srcDoc string) {
	p.Set("srcdoc", srcDoc)
}

func (p *htmlIFrameElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlIFrameElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlIFrameElementImpl) Sandbox() DOMTokenList {
	return newDOMTokenList(p.Get("sandbox"))
}

func (p *htmlIFrameElementImpl) AllowFullScreen() bool {
	return p.Get("allowFullscreen").Bool()
}

func (p *htmlIFrameElementImpl) SetAllowFullScreen(b bool) {
	p.Set("allowFullscreen", b)
}

func (p *htmlIFrameElementImpl) AllowPaymentRequest() bool {
	return p.Get("allowPaymentRequest").Bool()
}

func (p *htmlIFrameElementImpl) SetAllowPaymentRequest(b bool) {
	p.Set("allowPaymentRequest", b)
}

func (p *htmlIFrameElementImpl) Width() string {
	return p.Get("width").String()
}

func (p *htmlIFrameElementImpl) SetWidth(w string) {
	p.Set("width", w)
}

func (p *htmlIFrameElementImpl) Height() string {
	return p.Get("height").String()
}

func (p *htmlIFrameElementImpl) SetHeight(h string) {
	p.Set("height", h)
}

func (p *htmlIFrameElementImpl) ReferrerPolicy() string {
	return p.Get("referrerPolicy").String()
}

func (p *htmlIFrameElementImpl) SetReferrerPolicy(policy string) {
	p.Set("referrerPolicy", policy)
}

func (p *htmlIFrameElementImpl) ContentDocument() Document {
	return newDocument(p.Get("contentDocument"))
}

func (p *htmlIFrameElementImpl) ContentWindow() WindowProxy {
	return newWindowProxy(p.Get("contentWindow"))
}

// -------------8<---------------------------------------

type htmlEmbedElementImpl struct {
	*htmlElementImpl
}

func newHTMLEmbedElement(v js.Value) HTMLEmbedElement {
	if isNil(v) {
		return nil
	}

	return &htmlEmbedElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlEmbedElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlEmbedElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlEmbedElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlEmbedElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlEmbedElementImpl) Width() string {
	return p.Get("width").String()
}

func (p *htmlEmbedElementImpl) SetWidth(w string) {
	p.Set("width", w)
}

func (p *htmlEmbedElementImpl) Height() string {
	return p.Get("height").String()
}

func (p *htmlEmbedElementImpl) SetHeight(h string) {
	p.Set("height", h)
}

// -------------8<---------------------------------------

type htmlObjectElementImpl struct {
	*htmlElementImpl
}

func newHTMLObjectElement(v js.Value) HTMLObjectElement {
	if isNil(v) {
		return nil
	}

	return &htmlObjectElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlObjectElementImpl) Data() string {
	return p.Get("data").String()
}

func (p *htmlObjectElementImpl) SetData(d string) {
	p.Set("data", d)
}

func (p *htmlObjectElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlObjectElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlObjectElementImpl) TypeMustMatch() bool {
	return p.Get("typeMustMatch").Bool()
}

func (p *htmlObjectElementImpl) SetTypeMustMatch(b bool) {
	p.Set("typeMustMatch", b)
}

func (p *htmlObjectElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlObjectElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlObjectElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlObjectElementImpl) Width() string {
	return p.Get("width").String()
}

func (p *htmlObjectElementImpl) SetWidth(w string) {
	p.Set("width", w)
}

func (p *htmlObjectElementImpl) Height() string {
	return p.Get("height").String()
}

func (p *htmlObjectElementImpl) SetHeight(h string) {
	p.Set("height", h)
}

func (p *htmlObjectElementImpl) ContentDocument() Document {
	return newDocument(p.Get("contentDocument"))
}

func (p *htmlObjectElementImpl) ContentWindow() WindowProxy {
	return newWindowProxy(p.Get("contentWindow"))
}

func (p *htmlObjectElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlObjectElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlObjectElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlObjectElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlObjectElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlObjectElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

// -------------8<---------------------------------------

type validityStateImpl struct {
	*htmlElementImpl
}

func newValidityState(v js.Value) ValidityState {
	if isNil(v) {
		return nil
	}

	return &validityStateImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *validityStateImpl) ValueMissing() bool {
	return p.Get("valueMissing").Bool()
}

func (p *validityStateImpl) TypeMismatch() bool {
	return p.Get("typeMismatch").Bool()
}

func (p *validityStateImpl) PatternMismatch() bool {
	return p.Get("patternMismatch").Bool()
}

func (p *validityStateImpl) TooLong() bool {
	return p.Get("tooLong").Bool()
}

func (p *validityStateImpl) TooShort() bool {
	return p.Get("tooShort").Bool()
}

func (p *validityStateImpl) RangeUnderflow() bool {
	return p.Get("rangeUnderflow").Bool()
}

func (p *validityStateImpl) RangeOverflow() bool {
	return p.Get("rangeOverflow").Bool()
}

func (p *validityStateImpl) StepMismatch() bool {
	return p.Get("stepMismatch").Bool()
}

func (p *validityStateImpl) BadInput() bool {
	return p.Get("badInput").Bool()
}

func (p *validityStateImpl) CustomError() bool {
	return p.Get("customError").Bool()
}

func (p *validityStateImpl) Valid() bool {
	return p.Get("valid").Bool()
}

// -------------8<---------------------------------------

type htmlParamElementImpl struct {
	*htmlElementImpl
}

func newHTMLParamElement(v js.Value) HTMLParamElement {
	if isNil(v) {
		return nil
	}

	return &htmlParamElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlParamElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlParamElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlParamElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlParamElementImpl) SetValue(value string) {
	p.Set("value", value)
}

// -------------8<---------------------------------------

type htmlVideoElementImpl struct {
	*htmlMediaElementImpl
}

func newHTMLVideoElement(v js.Value) HTMLVideoElement {
	if isNil(v) {
		return nil
	}

	return &htmlVideoElementImpl{
		htmlMediaElementImpl: newHTMLMediaElementImpl(v),
	}
}

func (p *htmlVideoElementImpl) Width() int {
	return p.Get("width").Int()
}

func (p *htmlVideoElementImpl) SetWidth(w int) {
	p.Set("width", w)
}

func (p *htmlVideoElementImpl) Height() int {
	return p.Get("height").Int()
}

func (p *htmlVideoElementImpl) SetHeight(h int) {
	p.Set("height", h)
}

func (p *htmlVideoElementImpl) VideoWidth() int {
	return p.Get("videoWidth").Int()
}

func (p *htmlVideoElementImpl) VideoHeight() int {
	return p.Get("videoHeight").Int()
}

func (p *htmlVideoElementImpl) Poster() string {
	return p.Get("poster").String()
}

func (p *htmlVideoElementImpl) SetPoster(poster string) {
	p.Set("poster", poster)
}

// -------------8<---------------------------------------

type htmlAudioElementImpl struct {
	*htmlMediaElementImpl
}

func newHTMLAudioElement(v js.Value) HTMLAudioElement {
	if isNil(v) {
		return nil
	}

	return &htmlAudioElementImpl{
		htmlMediaElementImpl: newHTMLMediaElementImpl(v),
	}
}

// -------------8<---------------------------------------

type htmlTrackElementImpl struct {
	*htmlElementImpl
}

func newHTMLTrackElement(v js.Value) HTMLTrackElement {
	if isNil(v) {
		return nil
	}

	return &htmlTrackElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTrackElementImpl) Kind() string {
	return p.Get("kind").String()
}

func (p *htmlTrackElementImpl) SetKind(k string) {
	p.Set("kind", k)
}

func (p *htmlTrackElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlTrackElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlTrackElementImpl) SrcLang() string {
	return p.Get("srclang").String()
}

func (p *htmlTrackElementImpl) SetSrcLang(sl string) {
	p.Set("srclang", sl)
}

func (p *htmlTrackElementImpl) Label() string {
	return p.Get("label").String()
}

func (p *htmlTrackElementImpl) SetLabel(lbl string) {
	p.Set("label", lbl)
}

func (p *htmlTrackElementImpl) Default() bool {
	return p.Get("default").Bool()
}

func (p *htmlTrackElementImpl) SetDefault(b bool) {
	p.Set("default", b)
}

func (p *htmlTrackElementImpl) ReadyState() HTMLTrackElementReadyState {
	return HTMLTrackElementReadyState(p.Get("readyState").Int())
}

func (p *htmlTrackElementImpl) Track() TextTrack {
	return newTextTrack(p.Get("track"))
}

// -------------8<---------------------------------------

type textTrackImpl struct {
	*eventTargetImpl
}

func newTextTrack(v js.Value) TextTrack {
	if isNil(v) {
		return nil
	}

	return &textTrackImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *textTrackImpl) Kind() TextTrackKind {
	return TextTrackKind(p.Get("kind").String())
}

func (p *textTrackImpl) Label() string {
	return p.Get("label").String()
}

func (p *textTrackImpl) Language() string {
	return p.Get("language").String()
}

func (p *textTrackImpl) Id() string {
	return p.Get("id").String()
}

func (p *textTrackImpl) InBandMetadataTrackDispatchType() string {
	return p.Get("inBandMetadataTrackDispatchType").String()
}

func (p *textTrackImpl) Mode() TextTrackMode {
	return TextTrackMode(p.Get("mode").String())
}

func (p *textTrackImpl) SetMode(mode TextTrackMode) {
	p.Set("mode", string(mode))
}

func (p *textTrackImpl) Cues() TextTrackCueList {
	return newTextTrackCueList(p.Get("cues"))
}

func (p *textTrackImpl) ActiveCues() TextTrackCueList {
	return newTextTrackCueList(p.Get("activeCues"))
}

func (p *textTrackImpl) AddCue(cue TextTrackCue) {
	p.Call("addCue", cue.JSValue())
}

func (p *textTrackImpl) RemoveCue(cue TextTrackCue) {
	p.Call("removeCue", cue.JSValue())
}

func (p *textTrackImpl) OnCueChange(fn func(Event)) EventHandler {
	return p.On("cuechange", fn)
}

// -------------8<---------------------------------------

type textTrackCueListImpl struct {
	js.Value
}

func newTextTrackCueList(v js.Value) TextTrackCueList {
	if isNil(v) {
		return nil
	}

	return &textTrackCueListImpl{
		Value: v,
	}
}

func (p *textTrackCueListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *textTrackCueListImpl) Item(index int) TextTrackCue {
	return newTextTrackCue(p.Call("item", index))
}

func (p *textTrackCueListImpl) CueById(id string) TextTrackCue {
	return newTextTrackCue(p.Call("getCueById", id))
}

// -------------8<---------------------------------------

type textTrackCueImpl struct {
	*eventTargetImpl
}

func newTextTrackCue(v js.Value) TextTrackCue {
	if isNil(v) {
		return nil
	}

	return &textTrackCueImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *textTrackCueImpl) Track() TextTrack {
	return newTextTrack(p.Get("track"))
}

func (p *textTrackCueImpl) Id() string {
	return p.Get("id").String()
}

func (p *textTrackCueImpl) SetId(id string) {
	p.Set("id", id)
}

func (p *textTrackCueImpl) StartTime() float64 {
	return p.Get("startTime").Float()
}

func (p *textTrackCueImpl) SetStartTime(st float64) {
	p.Set("startTime", st)
}

func (p *textTrackCueImpl) EndTime() float64 {
	return p.Get("endTime").Float()
}

func (p *textTrackCueImpl) SetEndTime(et float64) {
	p.Set("endTime", et)
}

func (p *textTrackCueImpl) PauseOnExit() bool {
	return p.Get("pauseOnExit").Bool()
}

func (p *textTrackCueImpl) SetPauseOnExit(poe bool) {
	p.Set("pauseOnExit", poe)
}

func (p *textTrackCueImpl) OnEnter(fn func(Event)) EventHandler {
	return p.On("enter", fn)
}

func (p *textTrackCueImpl) OnExit(fn func(Event)) EventHandler {
	return p.On("exit", fn)
}

// -------------8<---------------------------------------

type htmlMapElementImpl struct {
	*htmlElementImpl
}

func newHTMLMapElement(v js.Value) HTMLMapElement {
	if isNil(v) {
		return nil
	}

	return &htmlMapElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlMapElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlMapElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlMapElementImpl) Areas() HTMLCollection {
	return newHTMLCollection(p.Get("areas"))
}

func (p *htmlMapElementImpl) Images() HTMLCollection {
	return newHTMLCollection(p.Get("images"))
}

// -------------8<---------------------------------------

type htmlAreaElementImpl struct {
	*htmlElementImpl
	*htmlHyperlinkElementUtilsImpl
	js.Value
}

func newHTMLAreaElement(v js.Value) HTMLAreaElement {
	if isNil(v) {
		return nil
	}

	return &htmlAreaElementImpl{
		htmlElementImpl:               newHTMLElementImpl(v),
		htmlHyperlinkElementUtilsImpl: newHTMLHyperlinkElementUtilsImpl(v),
		Value:                         v,
	}
}

func (p *htmlAreaElementImpl) Alt() string {
	return p.Get("alt").String()
}

func (p *htmlAreaElementImpl) SetAlt(alt string) {
	p.Set("alt", alt)
}

func (p *htmlAreaElementImpl) Coords() string {
	return p.Get("coords").String()
}

func (p *htmlAreaElementImpl) SetCoords(coords string) {
	p.Set("coords", coords)
}

func (p *htmlAreaElementImpl) Shape() string {
	return p.Get("shape").String()
}

func (p *htmlAreaElementImpl) SetShape(shape string) {
	p.Set("shape", shape)
}

func (p *htmlAreaElementImpl) Target() string {
	return p.Get("target").String()
}

func (p *htmlAreaElementImpl) SetTarget(target string) {
	p.Set("target", target)
}

func (p *htmlAreaElementImpl) Download() string {
	return p.Get("download").String()
}

func (p *htmlAreaElementImpl) SetDownload(download string) {
	p.Set("download", download)
}

func (p *htmlAreaElementImpl) Rel() string {
	return p.Get("rel").String()
}

func (p *htmlAreaElementImpl) SetRel(rel string) {
	p.Set("rel", rel)
}

func (p *htmlAreaElementImpl) RelList() DOMTokenList {
	return newDOMTokenList(p.Get("relList"))
}

func (p *htmlAreaElementImpl) HrefLang() string {
	return p.Get("hreflang").String()
}

func (p *htmlAreaElementImpl) SetHrefLang(hl string) {
	p.Set("hreflang", hl)
}

func (p *htmlAreaElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlAreaElementImpl) SetType(typ string) {
	p.Set("type", typ)
}

func (p *htmlAreaElementImpl) ReferrerPolicy() string {
	return p.Get("referrerPolicy").String()
}

func (p *htmlAreaElementImpl) SetReferrerPolicy(policy string) {
	p.Set("referrerPolicy", policy)
}

// -------------8<---------------------------------------

type htmlMediaElementImpl struct {
	*htmlElementImpl
}

func newHTMLMediaElement(v js.Value) HTMLMediaElement {
	if p := newHTMLMediaElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLMediaElementImpl(v js.Value) *htmlMediaElementImpl {
	if isNil(v) {
		return nil
	}

	return &htmlMediaElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlMediaElementImpl) Error() MediaError {
	return newMediaError(p.Get("error"))
}

func (p *htmlMediaElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlMediaElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlMediaElementImpl) SrcObject() MediaProvider {
	return newMediaProvider(p.Get("srcObject"))
}

func (p *htmlMediaElementImpl) SetSrcObject(provider MediaProvider) {
	p.Set("srcObject", provider.JSValue())
}

func (p *htmlMediaElementImpl) CurrentSrc() string {
	return p.Get("currentSrc").String()
}

func (p *htmlMediaElementImpl) CrossOrigin() string {
	return p.Get("crossOrigin").String()
}

func (p *htmlMediaElementImpl) SetCrossOrigin(co string) {
	p.Set("crossOrigin", co)
}

func (p *htmlMediaElementImpl) NetworkState() MediaNetworkState {
	return MediaNetworkState(p.Get("networkState").Int())
}

func (p *htmlMediaElementImpl) Preload() string {
	return p.Get("preload").String()
}

func (p *htmlMediaElementImpl) SetPreload(preload string) {
	p.Set("preload", preload)
}

func (p *htmlMediaElementImpl) Buffered() TimeRanges {
	return newTimeRanges(p.Get("buffered"))
}

func (p *htmlMediaElementImpl) Load() {
	p.Call("load")
}

func (p *htmlMediaElementImpl) CanPlayType(typ string) CanPlayTypeResult {
	return CanPlayTypeResult(p.Call("canPlayType", typ).String())
}

func (p *htmlMediaElementImpl) ReadyState() MediaReadyState {
	return MediaReadyState(p.Get("readyState").Int())
}

func (p *htmlMediaElementImpl) Seeking() bool {
	return p.Get("seeking").Bool()
}

func (p *htmlMediaElementImpl) CurrentTime() float64 {
	return p.Get("currentTime").Float()
}

func (p *htmlMediaElementImpl) SetCurrentTime(ct float64) {
	p.Set("currentTime", ct)
}

func (p *htmlMediaElementImpl) FastSeek(fs float64) {
	p.Call("fastSeek", fs)
}

func (p *htmlMediaElementImpl) Duration() float64 {
	return p.Get("duration").Float()
}

func (p *htmlMediaElementImpl) StartDate() time.Time {
	return jsDateToTime(p.Call("getStartDate"))
}

func (p *htmlMediaElementImpl) Paused() bool {
	return p.Get("paused").Bool()
}

func (p *htmlMediaElementImpl) DefaultPlaybackRate() float64 {
	return p.Get("defaultPlaybackRate").Float()
}

func (p *htmlMediaElementImpl) SetDefaultPlaybackRate(rate float64) {
	p.Set("defaultPlaybackRate", rate)
}

func (p *htmlMediaElementImpl) PlaybackRate() float64 {
	return p.Get("playbackRate").Float()
}

func (p *htmlMediaElementImpl) SetPlaybackRate(rate float64) {
	p.Set("playbackRate", rate)
}

func (p *htmlMediaElementImpl) Played() TimeRanges {
	return newTimeRanges(p.Get("played"))
}

func (p *htmlMediaElementImpl) Seekable() TimeRanges {
	return newTimeRanges(p.Get("seekable"))
}

func (p *htmlMediaElementImpl) Ended() bool {
	return p.Get("ended").Bool()
}

func (p *htmlMediaElementImpl) AutoPlay() bool {
	return p.Get("autoplay").Bool()
}

func (p *htmlMediaElementImpl) SetAutoPlay(b bool) {
	p.Set("autoplay", b)
}

func (p *htmlMediaElementImpl) Loop() bool {
	return p.Get("loop").Bool()
}

func (p *htmlMediaElementImpl) SetLoop(b bool) {
	p.Set("loop", b)
}

func (p *htmlMediaElementImpl) Play() {
	p.Call("play")
}

func (p *htmlMediaElementImpl) Pause() {
	p.Call("pause")
}

func (p *htmlMediaElementImpl) Controls() bool {
	return p.Get("controls").Bool()
}

func (p *htmlMediaElementImpl) SetControls(b bool) {
	p.Set("controls", b)
}

func (p *htmlMediaElementImpl) Volume() bool {
	return p.Get("volume").Bool()
}

func (p *htmlMediaElementImpl) SetVolume(b bool) {
	p.Set("volume", b)
}

func (p *htmlMediaElementImpl) Muted() bool {
	return p.Get("muted").Bool()
}

func (p *htmlMediaElementImpl) SetMuted(b bool) {
	p.Set("muted", b)
}

func (p *htmlMediaElementImpl) DefaultMuted() bool {
	return p.Get("defaultMuted").Bool()
}

func (p *htmlMediaElementImpl) SetDefaultMuted(b bool) {
	p.Set("defaultMuted", b)
}

func (p *htmlMediaElementImpl) AudioTracks() AudioTrackList {
	return newAudioTrackList(p.Get("audioTracks"))
}

func (p *htmlMediaElementImpl) VideoTracks() VideoTrackList {
	return newVideoTrackList(p.Get("videoTracks"))
}

func (p *htmlMediaElementImpl) TextTracks() TextTrackList {
	return newTextTrackList(p.Get("textTracks"))
}

func (p *htmlMediaElementImpl) AddTextTrack(kind TextTrackKind, args ...string) TextTrack {
	switch len(args) {
	case 0:
		return newTextTrack(p.Call("addTextTrack", string(kind)))
	case 1:
		return newTextTrack(p.Call("addTextTrack", string(kind), args[0]))
	default:
		return newTextTrack(p.Call("addTextTrack", string(kind), args[0], args[1]))
	}
}

// -------------8<---------------------------------------

type audioTrackListImpl struct {
	*eventTargetImpl
}

func newAudioTrackList(v js.Value) AudioTrackList {
	if isNil(v) {
		return nil
	}

	return &audioTrackListImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *audioTrackListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *audioTrackListImpl) Item(index int) AudioTrack {
	return newAudioTrack(p.Call("item", index))
}

func (p *audioTrackListImpl) TrackById(id string) AudioTrack {
	return newAudioTrack(p.Call("getTrackById", id))
}

func (p *audioTrackListImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

func (p *audioTrackListImpl) OnAddTrack(fn func(Event)) EventHandler {
	return p.On("addtrack", fn)
}

func (p *audioTrackListImpl) OnRemoveTrack(fn func(Event)) EventHandler {
	return p.On("removetrack", fn)
}

// -------------8<---------------------------------------

type audioTrackImpl struct {
	js.Value
}

func newAudioTrack(v js.Value) AudioTrack {
	if isNil(v) {
		return nil
	}

	return &audioTrackImpl{
		Value: v,
	}
}

func (p *audioTrackImpl) Id() string {
	return p.Get("id").String()
}

func (p *audioTrackImpl) Kind() string {
	return p.Get("kind").String()
}

func (p *audioTrackImpl) Label() string {
	return p.Get("label").String()
}

func (p *audioTrackImpl) Language() string {
	return p.Get("language").String()
}

func (p *audioTrackImpl) Enabled() bool {
	return p.Get("enabled").Bool()
}

func (p *audioTrackImpl) SetEnabled(b bool) {
	p.Set("enabled", b)
}

// -------------8<---------------------------------------

type videoTrackListImpl struct {
	*eventTargetImpl
}

func newVideoTrackList(v js.Value) VideoTrackList {
	if isNil(v) {
		return nil
	}

	return &videoTrackListImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *videoTrackListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *videoTrackListImpl) Item(index int) VideoTrack {
	return newVideoTrack(p.Call("item", index))
}

func (p *videoTrackListImpl) TrackById(id string) VideoTrack {
	return newVideoTrack(p.Call("getTrackById", id))
}

func (p *videoTrackListImpl) SelectedIndex() int {
	return p.Get("selectedIndex").Int()
}

func (p *videoTrackListImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

func (p *videoTrackListImpl) OnAddTrack(fn func(Event)) EventHandler {
	return p.On("addtrack", fn)
}

func (p *videoTrackListImpl) OnRemoveTrack(fn func(Event)) EventHandler {
	return p.On("removetrack", fn)
}

// -------------8<---------------------------------------

type videoTrackImpl struct {
	js.Value
}

func newVideoTrack(v js.Value) VideoTrack {
	if isNil(v) {
		return nil
	}

	return &videoTrackImpl{
		v,
	}
}

func (p *videoTrackImpl) Id() string {
	return p.Get("id").String()
}

func (p *videoTrackImpl) Kind() string {
	return p.Get("kind").String()
}

func (p *videoTrackImpl) Label() string {
	return p.Get("label").String()
}

func (p *videoTrackImpl) Language() string {
	return p.Get("language").String()
}

func (p *videoTrackImpl) Selected() bool {
	return p.Get("selected").Bool()
}

func (p *videoTrackImpl) SetSelected(b bool) {
	p.Set("selected", b)
}

// -------------8<---------------------------------------

type textTrackListImpl struct {
	*eventTargetImpl
}

func newTextTrackList(v js.Value) TextTrackList {
	if isNil(v) {
		return nil
	}

	return &textTrackListImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *textTrackListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *textTrackListImpl) Item(index int) TextTrack {
	return newTextTrack(p.Call("item", index))
}

func (p *textTrackListImpl) TrackById(id string) TextTrack {
	return newTextTrack(p.Call("getTrackById", id))
}

func (p *textTrackListImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

func (p *textTrackListImpl) OnAddTrack(fn func(Event)) EventHandler {
	return p.On("addtrack", fn)
}

func (p *textTrackListImpl) OnRemoveTrack(fn func(Event)) EventHandler {
	return p.On("removetrack", fn)
}

// -------------8<---------------------------------------

type timeRangesImpl struct {
	js.Value
}

func newTimeRanges(v js.Value) TimeRanges {
	if isNil(v) {
		return nil
	}

	return &timeRangesImpl{
		Value: v,
	}
}

func (p *timeRangesImpl) Length() int {
	return p.Get("length").Int()
}

func (p *timeRangesImpl) Start(index int) float64 {
	return p.Call("start", index).Float()
}

func (p *timeRangesImpl) End(index int) float64 {
	return p.Call("end", index).Float()
}

// -------------8<---------------------------------------

type mediaErrorImpl struct {
	js.Value
}

func newMediaError(v js.Value) MediaError {
	if isNil(v) {
		return nil
	}

	return &mediaErrorImpl{
		Value: v,
	}
}

func (p *mediaErrorImpl) Code() MediaErrorCode {
	return MediaErrorCode(p.Get("code").Int())
}

// -------------8<---------------------------------------

type mediaProviderImpl struct {
	js.Value
}

func newMediaProvider(v js.Value) MediaProvider {
	if isNil(v) {
		return nil
	}

	return &mediaProviderImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type mediaStreamImpl struct {
	*eventTargetImpl
}

func newMediaStream(v js.Value) MediaStream {
	if isNil(v) {
		return nil
	}

	return &mediaStreamImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *mediaStreamImpl) Id() string {
	return p.Get("id").String()
}

func (p *mediaStreamImpl) AudioTracks() []MediaStreamTrack {
	if s := arrayToSlice(p.Call("getAudioTracks")); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = newMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) VideoTracks() []MediaStreamTrack {
	if s := arrayToSlice(p.Call("getVideoTracks")); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = newMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) Tracks() []MediaStreamTrack {
	if s := arrayToSlice(p.Call("getTracks")); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = newMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) TrackById(id string) MediaStreamTrack {
	return newMediaStreamTrack(p.Call("getTrackById", id))
}

func (p *mediaStreamImpl) AddTrack(track MediaStreamTrack) {
	p.Call("addTrack", track.JSValue())
}

func (p *mediaStreamImpl) RemoveTrack(track MediaStreamTrack) {
	p.Call("removeTrack", track.JSValue())
}

func (p *mediaStreamImpl) Clone() MediaStream {
	return newMediaStream(p.Call("clone"))
}

func (p *mediaStreamImpl) Active() bool {
	return p.Get("active").Bool()
}

func (p *mediaStreamImpl) OnAddTrack(fn func(Event)) EventHandler {
	return p.On("addtrack", fn)
}

func (p *mediaStreamImpl) OnRemoveTrack(fn func(Event)) EventHandler {
	return p.On("removetrack", fn)
}

// -------------8<---------------------------------------

type mediaStreamTrackImpl struct {
	*eventTargetImpl
}

func newMediaStreamTrack(v js.Value) MediaStreamTrack {
	if isNil(v) {
		return nil
	}

	return &mediaStreamTrackImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *mediaStreamTrackImpl) Kind() string {
	return p.Get("kind").String()
}

func (p *mediaStreamTrackImpl) Id() string {
	return p.Get("id").String()
}

func (p *mediaStreamTrackImpl) Label() string {
	return p.Get("label").String()
}

func (p *mediaStreamTrackImpl) Enabled() bool {
	return p.Get("enabled").Bool()
}

func (p *mediaStreamTrackImpl) SetEnabled(b bool) {
	p.Set("enabled", b)
}

func (p *mediaStreamTrackImpl) Muted() bool {
	return p.Get("muted").Bool()
}

func (p *mediaStreamTrackImpl) OnMute(fn func(Event)) EventHandler {
	return p.On("mute", fn)
}

func (p *mediaStreamTrackImpl) OnUnMute(fn func(Event)) EventHandler {
	return p.On("unmute", fn)
}

func (p *mediaStreamTrackImpl) ReadyState() MediaStreamTrackState {
	return MediaStreamTrackState(p.Get("readyState").String())
}

func (p *mediaStreamTrackImpl) OnEnded(fn func(Event)) EventHandler {
	return p.On("ended", fn)
}

func (p *mediaStreamTrackImpl) Clone() MediaStreamTrack {
	return newMediaStreamTrack(p.Call("clone"))
}

func (p *mediaStreamTrackImpl) Stop() {
	p.Call("stop")
}

func (p *mediaStreamTrackImpl) Capabilities() MediaTrackCapabilities {
	v := p.Call("getCapabilities")
	if isNil(v) {
		return MediaTrackCapabilities{}
	}

	return MediaTrackCapabilities{
		Width:            newLongRange(v.Get("width")),
		Heigth:           newLongRange(v.Get("height")),
		AspectRatio:      newDoubleRange(v.Get("aspectRatio")),
		FrameRate:        newDoubleRange(v.Get("frameRate")),
		FacingMode:       stringSequenceToSlice(v.Get("facingMode")),
		Volume:           newDoubleRange(v.Get("volume")),
		SampleRate:       newLongRange(v.Get("sampleRate")),
		SampleSize:       newLongRange(v.Get("sampleSize")),
		EchoCancellation: boolSequenceToSlice(v.Get("echoCancellation")),
		AutoGainControl:  boolSequenceToSlice(v.Get("autoGainControl")),
		NoiseSuppression: boolSequenceToSlice(v.Get("noiseSuppression")),
		Latency:          newDoubleRange(v.Get("latency")),
		ChannelCount:     newLongRange(v.Get("channelCount")),
		DeviceId:         v.Get("deviceId").String(),
		GroupId:          v.Get("groupId").String(),
	}
}

func (p *mediaStreamTrackImpl) Constraints() MediaTrackConstraints {
	return newMediaTrackConstraints(p.Call("getConstraints"))
}

func (p *mediaStreamTrackImpl) Settings() MediaTrackSettings {
	v := p.Call("getSettings")
	if isNil(v) {
		return MediaTrackSettings{}
	}

	return MediaTrackSettings{
		Width:            v.Get("width").Int(),
		Height:           v.Get("height").Int(),
		AspectRatio:      v.Get("aspectRatio").Float(),
		FrameRate:        v.Get("frameRate").Float(),
		FacingMode:       v.Get("facingMode").String(),
		Volume:           v.Get("volume").Float(),
		SampleRate:       v.Get("sampleRate").Int(),
		SampleSize:       v.Get("sampleSize").Int(),
		EchoCancellation: v.Get("echoCancellation").Bool(),
		AutoGainControl:  v.Get("autoGainControl").Bool(),
		NoiseSuppression: v.Get("noiseSuppression").Bool(),
		Latency:          v.Get("latency").Float(),
		ChannelCount:     v.Get("channelCount").Int(),
		DeviceId:         v.Get("deviceId").String(),
		GroupId:          v.Get("groupId").String(),
	}
}

func (p *mediaStreamTrackImpl) OnOverConstrained(fn func(Event)) EventHandler {
	return p.On("overconstrained", fn)
}

// -------------8<---------------------------------------

func newMediaTrackConstraintSet(v js.Value) MediaTrackConstraintSet {
	if isNil(v) {
		return MediaTrackConstraintSet{}
	}

	return MediaTrackConstraintSet{
		Width:            newConstrainLong(v.Get("width")),
		Height:           newConstrainLong(v.Get("height")),
		AspectRatio:      newConstrainDouble(v.Get("aspectRatio")),
		FrameRate:        newConstrainDouble(v.Get("frameRate")),
		FacingMode:       newConstrainDOMString(v.Get("facingMode")),
		Volume:           newConstrainDouble(v.Get("volume")),
		SampleRate:       newConstrainLong(v.Get("sampleRate")),
		SampleSize:       newConstrainLong(v.Get("sampleSize")),
		EchoCancellation: newConstrainBoolean(v.Get("echoCancellation")),
		AutoGainControl:  newConstrainBoolean(v.Get("autoGainControl")),
		NoiseSuppression: newConstrainBoolean(v.Get("noiseSuppression")),
		Latency:          newConstrainDouble(v.Get("latency")),
		ChannelCount:     newConstrainLong(v.Get("channelCount")),
		DeviceId:         newConstrainDOMString(v.Get("deviceId")),
		GroupId:          newConstrainDOMString(v.Get("groupId")),
	}
}

func mediaTrackConstraintsSequenceToSlice(v js.Value) []MediaTrackConstraintSet {
	if isNil(v) {
		return nil
	}

	ret := make([]MediaTrackConstraintSet, v.Length())
	for i := range ret {
		ret[i] = newMediaTrackConstraintSet(v.Index(i))
	}

	return ret
}

func newMediaTrackConstraints(v js.Value) MediaTrackConstraints {
	if isNil(v) {
		return MediaTrackConstraints{}
	}

	return MediaTrackConstraints{
		MediaTrackConstraintSet: newMediaTrackConstraintSet(v),
		Advanced:                mediaTrackConstraintsSequenceToSlice(v.Get("advanced")),
	}
}

// -------------8<---------------------------------------

func newLongRange(v js.Value) LongRange {
	if isNil(v) {
		return LongRange{}
	}

	return LongRange{
		Max: v.Get("max").Int(),
		Min: v.Get("min").Int(),
	}
}

// -------------8<---------------------------------------

func newDoubleRange(v js.Value) DoubleRange {
	if isNil(v) {
		return DoubleRange{}
	}

	return DoubleRange{
		Max: v.Get("max").Float(),
		Min: v.Get("min").Float(),
	}
}

// -------------8<---------------------------------------

func newConstrainLong(v js.Value) ConstrainLong {
	if isNil(v) {
		return ConstrainLong{}
	}

	return ConstrainLong{
		LongRange: newLongRange(v),
		Exact:     v.Get("exact").Int(),
		Ideal:     v.Get("ideal").Int(),
	}
}

// -------------8<---------------------------------------

func newConstrainDouble(v js.Value) ConstrainDouble {
	if isNil(v) {
		return ConstrainDouble{}
	}

	return ConstrainDouble{
		DoubleRange: newDoubleRange(v),
		Exact:       v.Get("exact").Float(),
		Ideal:       v.Get("ideal").Float(),
	}
}

// -------------8<---------------------------------------

func newConstrainBoolean(v js.Value) ConstrainBoolean {
	if isNil(v) {
		return ConstrainBoolean{}
	}

	return ConstrainBoolean{
		Exact: v.Get("exact").Bool(),
		Ideal: v.Get("ideal").Bool(),
	}
}

// -------------8<---------------------------------------

func newConstrainDOMString(v js.Value) ConstrainDOMString {
	if isNil(v) {
		return ConstrainDOMString{}
	}

	return ConstrainDOMString{
		Exact: v.Get("exact").String(),
		Ideal: v.Get("ideal").String(),
	}
}

// -------------8<---------------------------------------

type sourceBufferImpl struct {
	*eventTargetImpl
}

func newSourceBuffer(v js.Value) SourceBuffer {
	if isNil(v) {
		return nil
	}

	return &sourceBufferImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *sourceBufferImpl) Mode() AppendMode {
	return AppendMode(p.Get("mode").String())
}

func (p *sourceBufferImpl) SetMode(mode AppendMode) {
	p.Set("mode", string(mode))
}

func (p *sourceBufferImpl) Updating() bool {
	return p.Get("updating").Bool()
}

func (p *sourceBufferImpl) Buffered() TimeRanges {
	return newTimeRanges(p.Get("buffered"))
}

func (p *sourceBufferImpl) TimestampOffset() float64 {
	return p.Get("timestampOffset").Float()
}

func (p *sourceBufferImpl) SetTimestampOffset(t float64) {
	p.Set("timestampOffset", t)
}

func (p *sourceBufferImpl) AudioTracks() AudioTrackList {
	return newAudioTrackList(p.Get("audioTracks"))
}

func (p *sourceBufferImpl) VideoTracks() VideoTrackList {
	return newVideoTrackList(p.Get("videoTracks"))
}

func (p *sourceBufferImpl) TextTracks() TextTrackList {
	return newTextTrackList(p.Get("textTracks"))
}

func (p *sourceBufferImpl) AppendWindowStart() float64 {
	return p.Get("appendWindowStart").Float()
}

func (p *sourceBufferImpl) SetAppendWindowStart(ws float64) {
	p.Set("appendWindowStart", ws)
}

func (p *sourceBufferImpl) AppendWindowEnd() float64 {
	return p.Get("appendWindowEnd").Float()
}

func (p *sourceBufferImpl) SetAppendWindowEnd(we float64) {
	p.Set("appendWindowEnd", we)
}

func (p *sourceBufferImpl) OnUpdatestart(fn func(Event)) EventHandler {
	return p.On("updatestart", fn)
}

func (p *sourceBufferImpl) OnUpdate(fn func(Event)) EventHandler {
	return p.On("update", fn)
}

func (p *sourceBufferImpl) OnUpdateend(fn func(Event)) EventHandler {
	return p.On("updateend", fn)
}

func (p *sourceBufferImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

func (p *sourceBufferImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

func (p *sourceBufferImpl) AppendBuffer(buffer BufferSource) {
	p.Call("appendBuffer", buffer.JSValue())
}

func (p *sourceBufferImpl) Abort() {
	p.Call("abort")
}

func (p *sourceBufferImpl) Remove(start float64, end float64) {
	p.Call("remove", start, end)
}

// -------------8<---------------------------------------
