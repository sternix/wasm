// +build js,wasm

package wasm

import (
	"time"
)

// -------------8<---------------------------------------

func NewAudio(src ...string) HTMLAudioElement {
	if jsHTMLAudioElement := jsGlobal.get("HTMLAudioElement"); jsHTMLAudioElement.valid() {
		switch len(src) {
		case 0:
			return wrapHTMLAudioElement(jsHTMLAudioElement.jsNew())
		default:
			return wrapHTMLAudioElement(jsHTMLAudioElement.jsNew(src[0]))
		}
	}
	return nil
}

func NewImage(args ...uint) HTMLImageElement {
	if jsHTMLImageElement := jsGlobal.get("HTMLImageElement"); jsHTMLImageElement.valid() {
		switch len(args) {
		case 0:
			return wrapHTMLImageElement(jsHTMLImageElement.jsNew())
		case 1:
			return wrapHTMLImageElement(jsHTMLImageElement.jsNew(args[0]))
		default:
			return wrapHTMLImageElement(jsHTMLImageElement.jsNew(args[0], args[1]))
		}
	}
	return nil
}

func NewMediaStream(args ...interface{}) MediaStream {
	if jsMediaStream := jsGlobal.get("MediaStream"); jsMediaStream.valid() {
		switch len(args) {
		case 0:
			return wrapMediaStream(jsMediaStream.jsNew())
		default:
			switch x := args[0].(type) {
			case MediaStream:
				return wrapMediaStream(jsMediaStream.jsNew(JSValue(x)))
			case []MediaStreamTrack:
				var s []Value
				for _, m := range x {
					s = append(s, JSValue(m))
				}
				return wrapMediaStream(jsMediaStream.jsNew(sliceToJsArray(s)))
			}
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlBodyElementImpl struct {
	*eventTargetImpl
	*htmlElementImpl
	*windowEventHandlersImpl
	Value
}

func NewHTMLBodyElement() HTMLBodyElement {
	if el := CurrentDocument().CreateElement("body"); el != nil {
		if body, ok := el.(HTMLBodyElement); ok {
			return body
		}
	}
	return nil
}

func wrapHTMLBodyElement(v Value) HTMLBodyElement {
	if v.valid() {
		hbi := &htmlBodyElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
			Value:           v,
		}

		hbi.eventTargetImpl = hbi.htmlElementImpl.eventTargetImpl
		hbi.windowEventHandlersImpl = newWindowEventHandlersImpl(hbi.eventTargetImpl)
		return hbi
	}
	return nil
}

// -------------8<---------------------------------------

type htmlHeadingElementImpl struct {
	*htmlElementImpl
}

var htmlHeadingTags = map[int]string{
	1: "h1",
	2: "h2",
	3: "h3",
	4: "h4",
	5: "h5",
	6: "h6",
}

func NewHTMLHeadingElement(rank int) HTMLHeadingElement {
	if tag := htmlHeadingTags[rank]; tag != "" {
		if el := CurrentDocument().CreateElement(tag); el != nil {
			if h, ok := el.(HTMLHeadingElement); ok {
				return h
			}
		}
	}
	return nil
}

func wrapHTMLHeadingElement(v Value) HTMLHeadingElement {
	if v.valid() {
		return &htmlHeadingElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlParagraphElementImpl struct {
	*htmlElementImpl
}

func NewHTMLParagraphElement() HTMLParagraphElement {
	if el := CurrentDocument().CreateElement("p"); el != nil {
		if p, ok := el.(HTMLParagraphElement); ok {
			return p
		}
	}
	return nil
}

func wrapHTMLParagraphElement(v Value) HTMLParagraphElement {
	if v.valid() {
		return &htmlParagraphElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlHRElementImpl struct {
	*htmlElementImpl
}

func NewHTMLHRElement() HTMLHRElement {
	if el := CurrentDocument().CreateElement("hr"); el != nil {
		if hr, ok := el.(HTMLHRElement); ok {
			return hr
		}
	}
	return nil
}

func wrapHTMLHRElement(v Value) HTMLHRElement {
	if v.valid() {
		return &htmlHRElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlPreElementImpl struct {
	*htmlElementImpl
}

func NewHTMLPreElement() HTMLPreElement {
	if el := CurrentDocument().CreateElement("pre"); el != nil {
		if pre, ok := el.(HTMLPreElement); ok {
			return pre
		}
	}
	return nil
}

func wrapHTMLPreElement(v Value) HTMLPreElement {
	if v.valid() {
		return &htmlPreElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlQuoteElementImpl struct {
	*htmlElementImpl
}

func NewHTMLQuoteElement(block ...bool) HTMLQuoteElement {
	var tag = "q"

	if len(block) > 0 && block[0] == true {
		tag = "blockquote"
	}

	if el := CurrentDocument().CreateElement(tag); el != nil {
		if q, ok := el.(HTMLQuoteElement); ok {
			return q
		}
	}
	return nil
}

func wrapHTMLQuoteElement(v Value) HTMLQuoteElement {
	if v.valid() {
		return &htmlQuoteElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlQuoteElementImpl) Cite() string {
	return p.get("cite").toString()
}

func (p *htmlQuoteElementImpl) SetCite(cite string) {
	p.set("cite", cite)
}

// -------------8<---------------------------------------

type htmlOListElementImpl struct {
	*htmlElementImpl
}

func NewHTMLOListElement() HTMLOListElement {
	if el := CurrentDocument().CreateElement("ol"); el != nil {
		if ol, ok := el.(HTMLOListElement); ok {
			return ol
		}
	}
	return nil
}

func wrapHTMLOListElement(v Value) HTMLOListElement {
	if v.valid() {
		return &htmlOListElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlOListElementImpl) Reversed() bool {
	return p.get("reversed").toBool()
}

func (p *htmlOListElementImpl) SetReversed(r bool) {
	p.set("reversed", r)
}

func (p *htmlOListElementImpl) Start() int {
	return p.get("start").toInt()
}

func (p *htmlOListElementImpl) SetStart(s int) {
	p.set("start", s)
}

func (p *htmlOListElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlOListElementImpl) SetType(t string) {
	p.set("type", t)
}

// -------------8<---------------------------------------

type htmlUListElementImpl struct {
	*htmlElementImpl
}

func NewHTMLUListElement() HTMLUListElement {
	if el := CurrentDocument().CreateElement("ul"); el != nil {
		if ul, ok := el.(HTMLUListElement); ok {
			return ul
		}
	}
	return nil
}

func wrapHTMLUListElement(v Value) HTMLUListElement {
	if v.valid() {
		return &htmlUListElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlLIElementImpl struct {
	*htmlElementImpl
}

func NewHTMLLIElement() HTMLLIElement {
	if el := CurrentDocument().CreateElement("li"); el != nil {
		if li, ok := el.(HTMLLIElement); ok {
			return li
		}
	}
	return nil
}

func wrapHTMLLIElement(v Value) HTMLLIElement {
	if v.valid() {
		return &htmlLIElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlLIElementImpl) Value() int {
	return p.get("value").toInt()
}

func (p *htmlLIElementImpl) SetValue(v int) {
	p.set("value", v)
}

// -------------8<---------------------------------------

type htmlDListElementImpl struct {
	*htmlElementImpl
}

func NewHTMLDListElement() HTMLDListElement {
	if el := CurrentDocument().CreateElement("dl"); el != nil {
		if dl, ok := el.(HTMLDListElement); ok {
			return dl
		}
	}
	return nil
}

func wrapHTMLDListElement(v Value) HTMLDListElement {
	if v.valid() {
		return &htmlDListElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlDivElementImpl struct {
	*htmlElementImpl
}

func NewHTMLDivElement() HTMLDivElement {
	if el := CurrentDocument().CreateElement("div"); el != nil {
		if div, ok := el.(HTMLDivElement); ok {
			return div
		}
	}
	return nil
}

func wrapHTMLDivElement(v Value) HTMLDivElement {
	if v.valid() {
		return &htmlDivElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlAnchorElementImpl struct {
	*htmlElementImpl
	*htmlHyperlinkElementUtilsImpl
	Value
}

func NewHTMLAnchorElement() HTMLAnchorElement {
	if el := CurrentDocument().CreateElement("a"); el != nil {
		if a, ok := el.(HTMLAnchorElement); ok {
			return a
		}
	}
	return nil
}

func wrapHTMLAnchorElement(v Value) HTMLAnchorElement {
	if v.valid() {
		return &htmlAnchorElementImpl{
			htmlElementImpl:               newHTMLElementImpl(v),
			htmlHyperlinkElementUtilsImpl: newHTMLHyperlinkElementUtilsImpl(v),
			Value:                         v,
		}
	}
	return nil
}

func (p *htmlAnchorElementImpl) Target() string {
	return p.get("target").toString()
}

func (p *htmlAnchorElementImpl) SetTarget(t string) {
	p.set("target", t)
}

func (p *htmlAnchorElementImpl) Download() string {
	return p.get("download").toString()
}

func (p *htmlAnchorElementImpl) SetDownload(d string) {
	p.set("download", d)
}

func (p *htmlAnchorElementImpl) Rel() string {
	return p.get("rel").toString()
}

func (p *htmlAnchorElementImpl) SetRel(r string) {
	p.set("rel", r)
}

func (p *htmlAnchorElementImpl) Rev() string {
	return p.get("rev").toString()
}

func (p *htmlAnchorElementImpl) SetRev(r string) {
	p.set("rev", r)
}

func (p *htmlAnchorElementImpl) RelList() DOMTokenList {
	return wrapDOMTokenList(p.get("relList"))
}

func (p *htmlAnchorElementImpl) HrefLang() string {
	return p.get("hreflang").toString()
}

func (p *htmlAnchorElementImpl) SetHrefLang(l string) {
	p.set("hreflang", l)
}

func (p *htmlAnchorElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlAnchorElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlAnchorElementImpl) Text() string {
	return p.get("text").toString()
}

func (p *htmlAnchorElementImpl) SetText(text string) {
	p.set("text", text)
}

func (p *htmlAnchorElementImpl) ReferrerPolicy() string {
	return p.get("referrerPolicy").toString()
}

func (p *htmlAnchorElementImpl) SetReferrerPolicy(policy string) {
	p.set("referrerPolicy", policy)
}

// -------------8<---------------------------------------

type htmlHyperlinkElementUtilsImpl struct {
	Value
}

func wrapHTMLHyperlinkElementUtils(v Value) HTMLHyperlinkElementUtils {
	if p := newHTMLHyperlinkElementUtilsImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLHyperlinkElementUtilsImpl(v Value) *htmlHyperlinkElementUtilsImpl {
	if v.valid() {
		return &htmlHyperlinkElementUtilsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *htmlHyperlinkElementUtilsImpl) Href() string {
	return p.get("href").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHref(href string) {
	p.set("href", href)
}

func (p *htmlHyperlinkElementUtilsImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) Protocol() string {
	return p.get("protocol").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetProtocol(protocol string) {
	p.set("protocol", protocol)
}

func (p *htmlHyperlinkElementUtilsImpl) Username() string {
	return p.get("username").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetUsername(username string) {
	p.set("username", username)
}

func (p *htmlHyperlinkElementUtilsImpl) Password() string {
	return p.get("password").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPassword(password string) {
	p.set("password", password)
}

func (p *htmlHyperlinkElementUtilsImpl) Host() string {
	return p.get("host").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHost(host string) {
	p.set("host", host)
}

func (p *htmlHyperlinkElementUtilsImpl) Hostname() string {
	return p.get("hostname").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHostname(hostname string) {
	p.set("hostname", hostname)
}

func (p *htmlHyperlinkElementUtilsImpl) Port() string {
	return p.get("port").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPort(port string) {
	p.set("port", port)
}

func (p *htmlHyperlinkElementUtilsImpl) Pathname() string {
	return p.get("pathname").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetPathname(pathname string) {
	p.set("pathname", pathname)
}

func (p *htmlHyperlinkElementUtilsImpl) Search() string {
	return p.get("search").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetSearch(search string) {
	p.set("search", search)
}

func (p *htmlHyperlinkElementUtilsImpl) Hash() string {
	return p.get("hash").toString()
}

func (p *htmlHyperlinkElementUtilsImpl) SetHash(hash string) {
	p.set("hash", hash)
}

// -------------8<---------------------------------------

type htmlDataElementImpl struct {
	*htmlElementImpl
}

func NewHTMLDataElement() HTMLDataElement {
	if el := CurrentDocument().CreateElement("data"); el != nil {
		if data, ok := el.(HTMLDataElement); ok {
			return data
		}
	}
	return nil
}

func wrapHTMLDataElement(v Value) HTMLDataElement {
	if v.valid() {
		return &htmlDataElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDataElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlDataElementImpl) SetValue(value string) {
	p.set("value", value)
}

// -------------8<---------------------------------------

type htmlTimeElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTimeElement() HTMLTimeElement {
	if el := CurrentDocument().CreateElement("time"); el != nil {
		if tim, ok := el.(HTMLTimeElement); ok {
			return tim
		}
	}
	return nil
}

func wrapHTMLTimeElement(v Value) HTMLTimeElement {
	if v.valid() {
		return &htmlTimeElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTimeElementImpl) DateTime() string {
	return p.get("dateTime").toString()
}

func (p *htmlTimeElementImpl) SetDateTime(dt string) {
	p.set("dateTime", dt)
}

// -------------8<---------------------------------------

type htmlSpanElementImpl struct {
	*htmlElementImpl
}

func NewHTMLSpanElement() HTMLSpanElement {
	if el := CurrentDocument().CreateElement("span"); el != nil {
		if span, ok := el.(HTMLSpanElement); ok {
			return span
		}
	}
	return nil
}

func wrapHTMLSpanElement(v Value) HTMLSpanElement {
	if v.valid() {
		return &htmlSpanElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlBRElementImpl struct {
	*htmlElementImpl
}

func NewHTMLBRElement() HTMLBRElement {
	if el := CurrentDocument().CreateElement("br"); el != nil {
		if br, ok := el.(HTMLBRElement); ok {
			return br
		}
	}
	return nil
}

func wrapHTMLBRElement(v Value) HTMLBRElement {
	if v.valid() {
		return &htmlBRElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlModElementImpl struct {
	*htmlElementImpl
}

func NewHTMLDelElement() HTMLModElement {
	if el := CurrentDocument().CreateElement("del"); el != nil {
		if d, ok := el.(HTMLModElement); ok {
			return d
		}
	}
	return nil
}

func NewHTMLInsElement() HTMLModElement {
	if el := CurrentDocument().CreateElement("ins"); el != nil {
		if d, ok := el.(HTMLModElement); ok {
			return d
		}
	}
	return nil
}

func wrapHTMLModElement(v Value) HTMLModElement {
	if v.valid() {
		return &htmlModElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlModElementImpl) Cite() string {
	return p.get("cite").toString()
}

func (p *htmlModElementImpl) SetCite(cite string) {
	p.set("cite", cite)
}

func (p *htmlModElementImpl) DateTime() string {
	return p.get("dateTime").toString()
}

func (p *htmlModElementImpl) SetDateTime(dt string) {
	p.set("dateTime", dt)
}

// -------------8<---------------------------------------

type htmlPictureElementImpl struct {
	*htmlElementImpl
}

func NewHTMLPictureElement() HTMLPictureElement {
	if el := CurrentDocument().CreateElement("picture"); el != nil {
		if picture, ok := el.(HTMLPictureElement); ok {
			return picture
		}
	}
	return nil
}

func wrapHTMLPictureElement(v Value) HTMLPictureElement {
	if v.valid() {
		return &htmlPictureElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlSourceElementImpl struct {
	*htmlElementImpl
}

func NewHTMLSourceElement() HTMLSourceElement {
	if el := CurrentDocument().CreateElement("source"); el != nil {
		if source, ok := el.(HTMLSourceElement); ok {
			return source
		}
	}
	return nil
}

func wrapHTMLSourceElement(v Value) HTMLSourceElement {
	if v.valid() {
		return &htmlSourceElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlSourceElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlSourceElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlSourceElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlSourceElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlSourceElementImpl) SrcSet() string {
	return p.get("srcset").toString()
}

func (p *htmlSourceElementImpl) SetSrcSet(srcset string) {
	p.set("srcset", srcset)
}

func (p *htmlSourceElementImpl) Sizes() string {
	return p.get("sizes").toString()
}

func (p *htmlSourceElementImpl) SetSizes(sizes string) {
	p.set("sizes", sizes)
}

func (p *htmlSourceElementImpl) Media() string {
	return p.get("media").toString()
}

func (p *htmlSourceElementImpl) SetMedia(media string) {
	p.set("media", media)
}

// -------------8<---------------------------------------

type htmlImageElementImpl struct {
	*htmlElementImpl
}

func NewHTMLImageElement() HTMLImageElement {
	if el := CurrentDocument().CreateElement("img"); el != nil {
		if img, ok := el.(HTMLImageElement); ok {
			return img
		}
	}
	return nil
}

func wrapHTMLImageElement(v Value) HTMLImageElement {
	if v.valid() {
		return &htmlImageElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlImageElementImpl) Alt() string {
	return p.get("alt").toString()
}

func (p *htmlImageElementImpl) SetAlt(alt string) {
	p.set("alt", alt)
}

func (p *htmlImageElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlImageElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlImageElementImpl) SrcSet() string {
	return p.get("srcset").toString()
}

func (p *htmlImageElementImpl) SetSrcSet(srcset string) {
	p.set("srcset", srcset)
}

func (p *htmlImageElementImpl) Sizes() string {
	return p.get("sizes").toString()
}

func (p *htmlImageElementImpl) SetSizes(sizes string) {
	p.set("sizes", sizes)
}

func (p *htmlImageElementImpl) CrossOrigin() string {
	return p.get("crossOrigin").toString()
}

func (p *htmlImageElementImpl) SetCrossOrigin(co string) {
	p.set("crossOrigin", co)
}

func (p *htmlImageElementImpl) UseMap() string {
	return p.get("useMap").toString()
}

func (p *htmlImageElementImpl) SetUseMap(um string) {
	p.set("useMap", um)
}

func (p *htmlImageElementImpl) LongDesc() string {
	return p.get("longDesc").toString()
}

func (p *htmlImageElementImpl) SetLongDesc(ld string) {
	p.set("longDesc", ld)
}

func (p *htmlImageElementImpl) IsMap() bool {
	return p.get("isMap").toBool()
}

func (p *htmlImageElementImpl) SetIsMap(b bool) {
	p.set("isMap", b)
}

func (p *htmlImageElementImpl) Width() int {
	return p.get("width").toInt()
}

func (p *htmlImageElementImpl) SetWidth(w int) {
	p.set("width", w)
}

func (p *htmlImageElementImpl) Height() int {
	return p.get("height").toInt()
}

func (p *htmlImageElementImpl) SetHeight(h int) {
	p.set("height", h)
}

func (p *htmlImageElementImpl) NaturalWidth() int {
	return p.get("naturalWidth").toInt()
}

func (p *htmlImageElementImpl) NaturalHeight() int {
	return p.get("naturalHeight").toInt()
}

func (p *htmlImageElementImpl) Complete() bool {
	return p.get("complete").toBool()
}

func (p *htmlImageElementImpl) CurrentSrc() string {
	return p.get("currentSrc").toString()
}

func (p *htmlImageElementImpl) ReferrerPolicy() string {
	return p.get("referrerPolicy").toString()
}

func (p *htmlImageElementImpl) SetReferrerPolicy(policy string) {
	p.set("referrerPolicy", policy)
}

func (p *htmlImageElementImpl) X() int {
	return p.get("x").toInt()
}

func (p *htmlImageElementImpl) Y() int {
	return p.get("y").toInt()
}

// -------------8<---------------------------------------

type htmlIFrameElementImpl struct {
	*htmlElementImpl
}

func NewHTMLIFrameElement() HTMLIFrameElement {
	if el := CurrentDocument().CreateElement("iframe"); el != nil {
		if iframe, ok := el.(HTMLIFrameElement); ok {
			return iframe
		}
	}
	return nil
}

func wrapHTMLIFrameElement(v Value) HTMLIFrameElement {
	if v.valid() {
		return &htmlIFrameElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlIFrameElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlIFrameElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlIFrameElementImpl) SrcDoc() string {
	return p.get("srcdoc").toString()
}

func (p *htmlIFrameElementImpl) SetSrcDoc(srcDoc string) {
	p.set("srcdoc", srcDoc)
}

func (p *htmlIFrameElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlIFrameElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlIFrameElementImpl) Sandbox() DOMTokenList {
	return wrapDOMTokenList(p.get("sandbox"))
}

func (p *htmlIFrameElementImpl) AllowFullScreen() bool {
	return p.get("allowFullscreen").toBool()
}

func (p *htmlIFrameElementImpl) SetAllowFullScreen(b bool) {
	p.set("allowFullscreen", b)
}

func (p *htmlIFrameElementImpl) AllowPaymentRequest() bool {
	return p.get("allowPaymentRequest").toBool()
}

func (p *htmlIFrameElementImpl) SetAllowPaymentRequest(b bool) {
	p.set("allowPaymentRequest", b)
}

func (p *htmlIFrameElementImpl) Width() string {
	return p.get("width").toString()
}

func (p *htmlIFrameElementImpl) SetWidth(w string) {
	p.set("width", w)
}

func (p *htmlIFrameElementImpl) Height() string {
	return p.get("height").toString()
}

func (p *htmlIFrameElementImpl) SetHeight(h string) {
	p.set("height", h)
}

func (p *htmlIFrameElementImpl) ReferrerPolicy() string {
	return p.get("referrerPolicy").toString()
}

func (p *htmlIFrameElementImpl) SetReferrerPolicy(policy string) {
	p.set("referrerPolicy", policy)
}

func (p *htmlIFrameElementImpl) ContentDocument() Document {
	return wrapDocument(p.get("contentDocument"))
}

func (p *htmlIFrameElementImpl) ContentWindow() WindowProxy {
	return wrapWindowProxy(p.get("contentWindow"))
}

// -------------8<---------------------------------------

type htmlEmbedElementImpl struct {
	*htmlElementImpl
}

func NewHTMLEmbedElement() HTMLEmbedElement {
	if el := CurrentDocument().CreateElement("embed"); el != nil {
		if embed, ok := el.(HTMLEmbedElement); ok {
			return embed
		}
	}
	return nil
}

func wrapHTMLEmbedElement(v Value) HTMLEmbedElement {
	if v.valid() {
		return &htmlEmbedElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlEmbedElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlEmbedElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlEmbedElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlEmbedElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlEmbedElementImpl) Width() string {
	return p.get("width").toString()
}

func (p *htmlEmbedElementImpl) SetWidth(w string) {
	p.set("width", w)
}

func (p *htmlEmbedElementImpl) Height() string {
	return p.get("height").toString()
}

func (p *htmlEmbedElementImpl) SetHeight(h string) {
	p.set("height", h)
}

// -------------8<---------------------------------------

type htmlObjectElementImpl struct {
	*htmlElementImpl
}

func NewHTMLObjectElement() HTMLObjectElement {
	if el := CurrentDocument().CreateElement("object"); el != nil {
		if obj, ok := el.(HTMLObjectElement); ok {
			return obj
		}
	}
	return nil
}

func wrapHTMLObjectElement(v Value) HTMLObjectElement {
	if v.valid() {
		return &htmlObjectElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlObjectElementImpl) Data() string {
	return p.get("data").toString()
}

func (p *htmlObjectElementImpl) SetData(d string) {
	p.set("data", d)
}

func (p *htmlObjectElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlObjectElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlObjectElementImpl) TypeMustMatch() bool {
	return p.get("typeMustMatch").toBool()
}

func (p *htmlObjectElementImpl) SetTypeMustMatch(b bool) {
	p.set("typeMustMatch", b)
}

func (p *htmlObjectElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlObjectElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlObjectElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlObjectElementImpl) Width() string {
	return p.get("width").toString()
}

func (p *htmlObjectElementImpl) SetWidth(w string) {
	p.set("width", w)
}

func (p *htmlObjectElementImpl) Height() string {
	return p.get("height").toString()
}

func (p *htmlObjectElementImpl) SetHeight(h string) {
	p.set("height", h)
}

func (p *htmlObjectElementImpl) ContentDocument() Document {
	return wrapDocument(p.get("contentDocument"))
}

func (p *htmlObjectElementImpl) ContentWindow() WindowProxy {
	return wrapWindowProxy(p.get("contentWindow"))
}

func (p *htmlObjectElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlObjectElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlObjectElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlObjectElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlObjectElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlObjectElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

// -------------8<---------------------------------------

type validityStateImpl struct {
	*htmlElementImpl
}

func wrapValidityState(v Value) ValidityState {
	if v.valid() {
		return &validityStateImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *validityStateImpl) ValueMissing() bool {
	return p.get("valueMissing").toBool()
}

func (p *validityStateImpl) TypeMismatch() bool {
	return p.get("typeMismatch").toBool()
}

func (p *validityStateImpl) PatternMismatch() bool {
	return p.get("patternMismatch").toBool()
}

func (p *validityStateImpl) TooLong() bool {
	return p.get("tooLong").toBool()
}

func (p *validityStateImpl) TooShort() bool {
	return p.get("tooShort").toBool()
}

func (p *validityStateImpl) RangeUnderflow() bool {
	return p.get("rangeUnderflow").toBool()
}

func (p *validityStateImpl) RangeOverflow() bool {
	return p.get("rangeOverflow").toBool()
}

func (p *validityStateImpl) StepMismatch() bool {
	return p.get("stepMismatch").toBool()
}

func (p *validityStateImpl) BadInput() bool {
	return p.get("badInput").toBool()
}

func (p *validityStateImpl) CustomError() bool {
	return p.get("customError").toBool()
}

func (p *validityStateImpl) Valid() bool {
	return p.get("valid").toBool()
}

// -------------8<---------------------------------------

type htmlParamElementImpl struct {
	*htmlElementImpl
}

func NewHTMLParamElement(name string, value string) HTMLParamElement {
	if el := CurrentDocument().CreateElement("param"); el != nil {
		if param, ok := el.(HTMLParamElement); ok {
			param.SetName(name)
			param.SetValue(value)
			return param
		}
	}
	return nil
}

func wrapHTMLParamElement(v Value) HTMLParamElement {
	if v.valid() {
		return &htmlParamElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlParamElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlParamElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlParamElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlParamElementImpl) SetValue(value string) {
	p.set("value", value)
}

// -------------8<---------------------------------------

type htmlVideoElementImpl struct {
	*htmlMediaElementImpl
}

func NewHTMLVideoElement() HTMLVideoElement {
	if el := CurrentDocument().CreateElement("video"); el != nil {
		if video, ok := el.(HTMLVideoElement); ok {
			return video
		}
	}
	return nil
}

func wrapHTMLVideoElement(v Value) HTMLVideoElement {
	if v.valid() {
		return &htmlVideoElementImpl{
			htmlMediaElementImpl: newHTMLMediaElementImpl(v),
		}
	}
	return nil
}

func (p *htmlVideoElementImpl) Width() int {
	return p.get("width").toInt()
}

func (p *htmlVideoElementImpl) SetWidth(w int) {
	p.set("width", w)
}

func (p *htmlVideoElementImpl) Height() int {
	return p.get("height").toInt()
}

func (p *htmlVideoElementImpl) SetHeight(h int) {
	p.set("height", h)
}

func (p *htmlVideoElementImpl) VideoWidth() int {
	return p.get("videoWidth").toInt()
}

func (p *htmlVideoElementImpl) VideoHeight() int {
	return p.get("videoHeight").toInt()
}

func (p *htmlVideoElementImpl) Poster() string {
	return p.get("poster").toString()
}

func (p *htmlVideoElementImpl) SetPoster(poster string) {
	p.set("poster", poster)
}

// -------------8<---------------------------------------

type htmlAudioElementImpl struct {
	*htmlMediaElementImpl
}

func NewHTMLAudioElement() HTMLAudioElement {
	if el := CurrentDocument().CreateElement("audio"); el != nil {
		if audio, ok := el.(HTMLAudioElement); ok {
			return audio
		}
	}
	return nil
}

func wrapHTMLAudioElement(v Value) HTMLAudioElement {
	if v.valid() {
		return &htmlAudioElementImpl{
			htmlMediaElementImpl: newHTMLMediaElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlTrackElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTrackElement() HTMLTrackElement {
	if el := CurrentDocument().CreateElement("track"); el != nil {
		if track, ok := el.(HTMLTrackElement); ok {
			return track
		}
	}
	return nil
}

func wrapHTMLTrackElement(v Value) HTMLTrackElement {
	if v.valid() {
		return &htmlTrackElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTrackElementImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *htmlTrackElementImpl) SetKind(k string) {
	p.set("kind", k)
}

func (p *htmlTrackElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlTrackElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlTrackElementImpl) SrcLang() string {
	return p.get("srclang").toString()
}

func (p *htmlTrackElementImpl) SetSrcLang(sl string) {
	p.set("srclang", sl)
}

func (p *htmlTrackElementImpl) Label() string {
	return p.get("label").toString()
}

func (p *htmlTrackElementImpl) SetLabel(lbl string) {
	p.set("label", lbl)
}

func (p *htmlTrackElementImpl) Default() bool {
	return p.get("default").toBool()
}

func (p *htmlTrackElementImpl) SetDefault(b bool) {
	p.set("default", b)
}

func (p *htmlTrackElementImpl) ReadyState() HTMLTrackElementReadyState {
	return HTMLTrackElementReadyState(p.get("readyState").toInt())
}

func (p *htmlTrackElementImpl) Track() TextTrack {
	return wrapTextTrack(p.get("track"))
}

// -------------8<---------------------------------------

type textTrackImpl struct {
	*eventTargetImpl
}

func wrapTextTrack(v Value) TextTrack {
	if v.valid() {
		return &textTrackImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *textTrackImpl) Kind() TextTrackKind {
	return TextTrackKind(p.get("kind").toString())
}

func (p *textTrackImpl) Label() string {
	return p.get("label").toString()
}

func (p *textTrackImpl) Language() string {
	return p.get("language").toString()
}

func (p *textTrackImpl) Id() string {
	return p.get("id").toString()
}

func (p *textTrackImpl) InBandMetadataTrackDispatchType() string {
	return p.get("inBandMetadataTrackDispatchType").toString()
}

func (p *textTrackImpl) Mode() TextTrackMode {
	return TextTrackMode(p.get("mode").toString())
}

func (p *textTrackImpl) SetMode(mode TextTrackMode) {
	p.set("mode", string(mode))
}

func (p *textTrackImpl) Cues() TextTrackCueList {
	return wrapTextTrackCueList(p.get("cues"))
}

func (p *textTrackImpl) ActiveCues() TextTrackCueList {
	return wrapTextTrackCueList(p.get("activeCues"))
}

func (p *textTrackImpl) AddCue(cue TextTrackCue) {
	p.call("addCue", JSValue(cue))
}

func (p *textTrackImpl) RemoveCue(cue TextTrackCue) {
	p.call("removeCue", JSValue(cue))
}

func (p *textTrackImpl) OnCueChange(fn func(Event)) EventHandler {
	return p.On("cuechange", fn)
}

// -------------8<---------------------------------------

type textTrackCueListImpl struct {
	Value
}

func wrapTextTrackCueList(v Value) TextTrackCueList {
	if v.valid() {
		return &textTrackCueListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textTrackCueListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *textTrackCueListImpl) Item(index uint) TextTrackCue {
	return wrapTextTrackCue(p.call("item", index))
}

func (p *textTrackCueListImpl) CueById(id string) TextTrackCue {
	return wrapTextTrackCue(p.call("getCueById", id))
}

// -------------8<---------------------------------------

type textTrackCueImpl struct {
	*eventTargetImpl
}

func wrapTextTrackCue(v Value) TextTrackCue {
	if v.valid() {
		return &textTrackCueImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *textTrackCueImpl) Track() TextTrack {
	return wrapTextTrack(p.get("track"))
}

func (p *textTrackCueImpl) Id() string {
	return p.get("id").toString()
}

func (p *textTrackCueImpl) SetId(id string) {
	p.set("id", id)
}

func (p *textTrackCueImpl) StartTime() float64 {
	return p.get("startTime").toFloat64()
}

func (p *textTrackCueImpl) SetStartTime(st float64) {
	p.set("startTime", st)
}

func (p *textTrackCueImpl) EndTime() float64 {
	return p.get("endTime").toFloat64()
}

func (p *textTrackCueImpl) SetEndTime(et float64) {
	p.set("endTime", et)
}

func (p *textTrackCueImpl) PauseOnExit() bool {
	return p.get("pauseOnExit").toBool()
}

func (p *textTrackCueImpl) SetPauseOnExit(poe bool) {
	p.set("pauseOnExit", poe)
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

func NewHTMLMapElement() HTMLMapElement {
	if el := CurrentDocument().CreateElement("map"); el != nil {
		if m, ok := el.(HTMLMapElement); ok {
			return m
		}
	}
	return nil
}

func wrapHTMLMapElement(v Value) HTMLMapElement {
	if v.valid() {
		return &htmlMapElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlMapElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlMapElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlMapElementImpl) Areas() []HTMLAreaElement {
	if c := wrapHTMLCollection(p.get("areas")); c != nil && c.Length() > 0 {
		var ret []HTMLAreaElement
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLAreaElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// <img> and <object> Elements
func (p *htmlMapElementImpl) Images() []HTMLElement {
	return htmlCollectionToHTMLElementSlice(p.get("images"))
}

// -------------8<---------------------------------------

type htmlAreaElementImpl struct {
	*htmlElementImpl
	*htmlHyperlinkElementUtilsImpl
	Value
}

func NewHTMLAreaElement() HTMLAreaElement {
	if el := CurrentDocument().CreateElement("area"); el != nil {
		if area, ok := el.(HTMLAreaElement); ok {
			return area
		}
	}
	return nil
}

func wrapHTMLAreaElement(v Value) HTMLAreaElement {
	if v.valid() {
		return &htmlAreaElementImpl{
			htmlElementImpl:               newHTMLElementImpl(v),
			htmlHyperlinkElementUtilsImpl: newHTMLHyperlinkElementUtilsImpl(v),
			Value:                         v,
		}
	}
	return nil
}

func (p *htmlAreaElementImpl) Alt() string {
	return p.get("alt").toString()
}

func (p *htmlAreaElementImpl) SetAlt(alt string) {
	p.set("alt", alt)
}

func (p *htmlAreaElementImpl) Coords() string {
	return p.get("coords").toString()
}

func (p *htmlAreaElementImpl) SetCoords(coords string) {
	p.set("coords", coords)
}

func (p *htmlAreaElementImpl) Shape() string {
	return p.get("shape").toString()
}

func (p *htmlAreaElementImpl) SetShape(shape string) {
	p.set("shape", shape)
}

func (p *htmlAreaElementImpl) Target() string {
	return p.get("target").toString()
}

func (p *htmlAreaElementImpl) SetTarget(target string) {
	p.set("target", target)
}

func (p *htmlAreaElementImpl) Download() string {
	return p.get("download").toString()
}

func (p *htmlAreaElementImpl) SetDownload(download string) {
	p.set("download", download)
}

func (p *htmlAreaElementImpl) Rel() string {
	return p.get("rel").toString()
}

func (p *htmlAreaElementImpl) SetRel(rel string) {
	p.set("rel", rel)
}

func (p *htmlAreaElementImpl) RelList() DOMTokenList {
	return wrapDOMTokenList(p.get("relList"))
}

func (p *htmlAreaElementImpl) HrefLang() string {
	return p.get("hreflang").toString()
}

func (p *htmlAreaElementImpl) SetHrefLang(hl string) {
	p.set("hreflang", hl)
}

func (p *htmlAreaElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlAreaElementImpl) SetType(typ string) {
	p.set("type", typ)
}

func (p *htmlAreaElementImpl) ReferrerPolicy() string {
	return p.get("referrerPolicy").toString()
}

func (p *htmlAreaElementImpl) SetReferrerPolicy(policy string) {
	p.set("referrerPolicy", policy)
}

// -------------8<---------------------------------------

type htmlMediaElementImpl struct {
	*htmlElementImpl
}

func wrapHTMLMediaElement(v Value) HTMLMediaElement {
	if p := newHTMLMediaElementImpl(v); p != nil {
		return p
	}
	return nil
}

func newHTMLMediaElementImpl(v Value) *htmlMediaElementImpl {
	if v.valid() {
		return &htmlMediaElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlMediaElementImpl) Error() MediaError {
	return wrapMediaError(p.get("error"))
}

func (p *htmlMediaElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlMediaElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlMediaElementImpl) SrcObject() MediaProvider {
	return wrapMediaProvider(p.get("srcObject"))
}

func (p *htmlMediaElementImpl) SetSrcObject(provider MediaProvider) {
	p.set("srcObject", JSValue(provider))
}

func (p *htmlMediaElementImpl) CurrentSrc() string {
	return p.get("currentSrc").toString()
}

func (p *htmlMediaElementImpl) CrossOrigin() string {
	return p.get("crossOrigin").toString()
}

func (p *htmlMediaElementImpl) SetCrossOrigin(co string) {
	p.set("crossOrigin", co)
}

func (p *htmlMediaElementImpl) NetworkState() MediaNetworkState {
	return MediaNetworkState(p.get("networkState").toInt())
}

func (p *htmlMediaElementImpl) Preload() string {
	return p.get("preload").toString()
}

func (p *htmlMediaElementImpl) SetPreload(preload string) {
	p.set("preload", preload)
}

func (p *htmlMediaElementImpl) Buffered() TimeRanges {
	return wrapTimeRanges(p.get("buffered"))
}

func (p *htmlMediaElementImpl) Load() {
	p.call("load")
}

func (p *htmlMediaElementImpl) CanPlayType(typ string) CanPlayTypeResult {
	return CanPlayTypeResult(p.call("canPlayType", typ).toString())
}

func (p *htmlMediaElementImpl) ReadyState() MediaReadyState {
	return MediaReadyState(p.get("readyState").toInt())
}

func (p *htmlMediaElementImpl) Seeking() bool {
	return p.get("seeking").toBool()
}

func (p *htmlMediaElementImpl) CurrentTime() float64 {
	return p.get("currentTime").toFloat64()
}

func (p *htmlMediaElementImpl) SetCurrentTime(ct float64) {
	p.set("currentTime", ct)
}

func (p *htmlMediaElementImpl) FastSeek(fs float64) {
	p.call("fastSeek", fs)
}

func (p *htmlMediaElementImpl) Duration() float64 {
	return p.get("duration").toFloat64()
}

func (p *htmlMediaElementImpl) StartDate() time.Time {
	return jsDateToTime(p.call("getStartDate"))
}

func (p *htmlMediaElementImpl) Paused() bool {
	return p.get("paused").toBool()
}

func (p *htmlMediaElementImpl) DefaultPlaybackRate() float64 {
	return p.get("defaultPlaybackRate").toFloat64()
}

func (p *htmlMediaElementImpl) SetDefaultPlaybackRate(rate float64) {
	p.set("defaultPlaybackRate", rate)
}

func (p *htmlMediaElementImpl) PlaybackRate() float64 {
	return p.get("playbackRate").toFloat64()
}

func (p *htmlMediaElementImpl) SetPlaybackRate(rate float64) {
	p.set("playbackRate", rate)
}

func (p *htmlMediaElementImpl) Played() TimeRanges {
	return wrapTimeRanges(p.get("played"))
}

func (p *htmlMediaElementImpl) Seekable() TimeRanges {
	return wrapTimeRanges(p.get("seekable"))
}

func (p *htmlMediaElementImpl) Ended() bool {
	return p.get("ended").toBool()
}

func (p *htmlMediaElementImpl) AutoPlay() bool {
	return p.get("autoplay").toBool()
}

func (p *htmlMediaElementImpl) SetAutoPlay(b bool) {
	p.set("autoplay", b)
}

func (p *htmlMediaElementImpl) Loop() bool {
	return p.get("loop").toBool()
}

func (p *htmlMediaElementImpl) SetLoop(b bool) {
	p.set("loop", b)
}

func (p *htmlMediaElementImpl) Play() {
	p.call("play")
}

func (p *htmlMediaElementImpl) Pause() {
	p.call("pause")
}

func (p *htmlMediaElementImpl) Controls() bool {
	return p.get("controls").toBool()
}

func (p *htmlMediaElementImpl) SetControls(b bool) {
	p.set("controls", b)
}

func (p *htmlMediaElementImpl) Volume() bool {
	return p.get("volume").toBool()
}

func (p *htmlMediaElementImpl) SetVolume(b bool) {
	p.set("volume", b)
}

func (p *htmlMediaElementImpl) Muted() bool {
	return p.get("muted").toBool()
}

func (p *htmlMediaElementImpl) SetMuted(b bool) {
	p.set("muted", b)
}

func (p *htmlMediaElementImpl) DefaultMuted() bool {
	return p.get("defaultMuted").toBool()
}

func (p *htmlMediaElementImpl) SetDefaultMuted(b bool) {
	p.set("defaultMuted", b)
}

func (p *htmlMediaElementImpl) AudioTracks() AudioTrackList {
	return wrapAudioTrackList(p.get("audioTracks"))
}

func (p *htmlMediaElementImpl) VideoTracks() VideoTrackList {
	return wrapVideoTrackList(p.get("videoTracks"))
}

func (p *htmlMediaElementImpl) TextTracks() TextTrackList {
	return wrapTextTrackList(p.get("textTracks"))
}

func (p *htmlMediaElementImpl) AddTextTrack(kind TextTrackKind, args ...string) TextTrack {
	switch len(args) {
	case 0:
		return wrapTextTrack(p.call("addTextTrack", string(kind)))
	case 1:
		return wrapTextTrack(p.call("addTextTrack", string(kind), args[0]))
	default:
		return wrapTextTrack(p.call("addTextTrack", string(kind), args[0], args[1]))
	}
}

// -------------8<---------------------------------------

type audioTrackListImpl struct {
	*eventTargetImpl
}

func wrapAudioTrackList(v Value) AudioTrackList {
	if v.valid() {
		return &audioTrackListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *audioTrackListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *audioTrackListImpl) Item(index uint) AudioTrack {
	return wrapAudioTrack(p.call("item", index))
}

func (p *audioTrackListImpl) TrackById(id string) AudioTrack {
	return wrapAudioTrack(p.call("getTrackById", id))
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
	Value
}

func wrapAudioTrack(v Value) AudioTrack {
	if v.valid() {
		return &audioTrackImpl{
			Value: v,
		}
	}
	return nil
}

func (p *audioTrackImpl) Id() string {
	return p.get("id").toString()
}

func (p *audioTrackImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *audioTrackImpl) Label() string {
	return p.get("label").toString()
}

func (p *audioTrackImpl) Language() string {
	return p.get("language").toString()
}

func (p *audioTrackImpl) Enabled() bool {
	return p.get("enabled").toBool()
}

func (p *audioTrackImpl) SetEnabled(b bool) {
	p.set("enabled", b)
}

// -------------8<---------------------------------------

type videoTrackListImpl struct {
	*eventTargetImpl
}

func wrapVideoTrackList(v Value) VideoTrackList {
	if v.valid() {
		return &videoTrackListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *videoTrackListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *videoTrackListImpl) Item(index uint) VideoTrack {
	return wrapVideoTrack(p.call("item", index))
}

func (p *videoTrackListImpl) TrackById(id string) VideoTrack {
	return wrapVideoTrack(p.call("getTrackById", id))
}

func (p *videoTrackListImpl) SelectedIndex() int {
	return p.get("selectedIndex").toInt()
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
	Value
}

func wrapVideoTrack(v Value) VideoTrack {
	if v.valid() {
		return &videoTrackImpl{
			Value: v,
		}
	}
	return nil
}

func (p *videoTrackImpl) Id() string {
	return p.get("id").toString()
}

func (p *videoTrackImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *videoTrackImpl) Label() string {
	return p.get("label").toString()
}

func (p *videoTrackImpl) Language() string {
	return p.get("language").toString()
}

func (p *videoTrackImpl) Selected() bool {
	return p.get("selected").toBool()
}

func (p *videoTrackImpl) SetSelected(b bool) {
	p.set("selected", b)
}

// -------------8<---------------------------------------

type textTrackListImpl struct {
	*eventTargetImpl
}

func wrapTextTrackList(v Value) TextTrackList {
	if v.valid() {
		return &textTrackListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *textTrackListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *textTrackListImpl) Item(index uint) TextTrack {
	return wrapTextTrack(p.call("item", index))
}

func (p *textTrackListImpl) TrackById(id string) TextTrack {
	return wrapTextTrack(p.call("getTrackById", id))
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
	Value
}

func wrapTimeRanges(v Value) TimeRanges {
	if v.valid() {
		return &timeRangesImpl{
			Value: v,
		}
	}
	return nil
}

func (p *timeRangesImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *timeRangesImpl) Start(index uint) float64 {
	return p.call("start", index).toFloat64()
}

func (p *timeRangesImpl) End(index uint) float64 {
	return p.call("end", index).toFloat64()
}

// -------------8<---------------------------------------

type mediaErrorImpl struct {
	Value
}

func wrapMediaError(v Value) MediaError {
	if v.valid() {
		return &mediaErrorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mediaErrorImpl) Code() MediaErrorCode {
	return MediaErrorCode(p.get("code").toInt())
}

// -------------8<---------------------------------------

type mediaProviderImpl struct {
	Value
}

func wrapMediaProvider(v Value) MediaProvider {
	if v.valid() {
		return &mediaProviderImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type mediaStreamImpl struct {
	*eventTargetImpl
}

func wrapMediaStream(v Value) MediaStream {
	if v.valid() {
		return &mediaStreamImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaStreamImpl) Id() string {
	return p.get("id").toString()
}

func (p *mediaStreamImpl) AudioTracks() []MediaStreamTrack {
	if s := p.call("getAudioTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) VideoTracks() []MediaStreamTrack {
	if s := p.call("getVideoTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) Tracks() []MediaStreamTrack {
	if s := p.call("getTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) TrackById(id string) MediaStreamTrack {
	return wrapMediaStreamTrack(p.call("getTrackById", id))
}

func (p *mediaStreamImpl) AddTrack(track MediaStreamTrack) {
	p.call("addTrack", JSValue(track))
}

func (p *mediaStreamImpl) RemoveTrack(track MediaStreamTrack) {
	p.call("removeTrack", JSValue(track))
}

func (p *mediaStreamImpl) Clone() MediaStream {
	return wrapMediaStream(p.call("clone"))
}

func (p *mediaStreamImpl) Active() bool {
	return p.get("active").toBool()
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

func wrapMediaStreamTrack(v Value) MediaStreamTrack {
	if v.valid() {
		return &mediaStreamTrackImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaStreamTrackImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *mediaStreamTrackImpl) Id() string {
	return p.get("id").toString()
}

func (p *mediaStreamTrackImpl) Label() string {
	return p.get("label").toString()
}

func (p *mediaStreamTrackImpl) Enabled() bool {
	return p.get("enabled").toBool()
}

func (p *mediaStreamTrackImpl) SetEnabled(b bool) {
	p.set("enabled", b)
}

func (p *mediaStreamTrackImpl) Muted() bool {
	return p.get("muted").toBool()
}

func (p *mediaStreamTrackImpl) OnMute(fn func(Event)) EventHandler {
	return p.On("mute", fn)
}

func (p *mediaStreamTrackImpl) OnUnMute(fn func(Event)) EventHandler {
	return p.On("unmute", fn)
}

func (p *mediaStreamTrackImpl) ReadyState() MediaStreamTrackState {
	return MediaStreamTrackState(p.get("readyState").toString())
}

func (p *mediaStreamTrackImpl) OnEnded(fn func(Event)) EventHandler {
	return p.On("ended", fn)
}

func (p *mediaStreamTrackImpl) Clone() MediaStreamTrack {
	return wrapMediaStreamTrack(p.call("clone"))
}

func (p *mediaStreamTrackImpl) Stop() {
	p.call("stop")
}

func (p *mediaStreamTrackImpl) Capabilities() MediaTrackCapabilities {
	if v := p.call("getCapabilities"); v.valid() {
		return MediaTrackCapabilities{
			Width:            wrapLongRange(v.get("width")),
			Heigth:           wrapLongRange(v.get("height")),
			AspectRatio:      wrapDoubleRange(v.get("aspectRatio")),
			FrameRate:        wrapDoubleRange(v.get("frameRate")),
			FacingMode:       stringSequenceToSlice(v.get("facingMode")),
			Volume:           wrapDoubleRange(v.get("volume")),
			SampleRate:       wrapLongRange(v.get("sampleRate")),
			SampleSize:       wrapLongRange(v.get("sampleSize")),
			EchoCancellation: boolSequenceToSlice(v.get("echoCancellation")),
			AutoGainControl:  boolSequenceToSlice(v.get("autoGainControl")),
			NoiseSuppression: boolSequenceToSlice(v.get("noiseSuppression")),
			Latency:          wrapDoubleRange(v.get("latency")),
			ChannelCount:     wrapLongRange(v.get("channelCount")),
			DeviceId:         v.get("deviceId").toString(),
			GroupId:          v.get("groupId").toString(),
		}
	}
	return MediaTrackCapabilities{}
}

func (p *mediaStreamTrackImpl) Constraints() MediaTrackConstraints {
	return wrapMediaTrackConstraints(p.call("getConstraints"))
}

func (p *mediaStreamTrackImpl) Settings() MediaTrackSettings {
	if v := p.call("getSettings"); v.valid() {
		return MediaTrackSettings{
			Width:            v.get("width").toInt(),
			Height:           v.get("height").toInt(),
			AspectRatio:      v.get("aspectRatio").toFloat64(),
			FrameRate:        v.get("frameRate").toFloat64(),
			FacingMode:       v.get("facingMode").toString(),
			Volume:           v.get("volume").toFloat64(),
			SampleRate:       v.get("sampleRate").toInt(),
			SampleSize:       v.get("sampleSize").toInt(),
			EchoCancellation: v.get("echoCancellation").toBool(),
			AutoGainControl:  v.get("autoGainControl").toBool(),
			NoiseSuppression: v.get("noiseSuppression").toBool(),
			Latency:          v.get("latency").toFloat64(),
			ChannelCount:     v.get("channelCount").toInt(),
			DeviceId:         v.get("deviceId").toString(),
			GroupId:          v.get("groupId").toString(),
		}
	}
	return MediaTrackSettings{}
}

func (p *mediaStreamTrackImpl) ApplyConstraints(constraints ...MediaTrackConstraints) func() error {
	return func() error {
		var (
			res Value
			ok  bool
		)

		switch len(constraints) {
		case 0:
			res, ok = await(p.call("applyConstraints"))
		default:
			res, ok = await(p.call("applyConstraints", constraints[0].toJSObject()))
		}

		if ok {
			return nil
		}
		return wrapDOMException(res)
	}
}

func (p *mediaStreamTrackImpl) OnOverConstrained(fn func(Event)) EventHandler {
	return p.On("overconstrained", fn)
}

// -------------8<---------------------------------------

func wrapMediaTrackConstraintSet(v Value) MediaTrackConstraintSet {
	if v.valid() {
		return MediaTrackConstraintSet{
			Width:            wrapConstrainLong(v.get("width")),
			Height:           wrapConstrainLong(v.get("height")),
			AspectRatio:      wrapConstrainDouble(v.get("aspectRatio")),
			FrameRate:        wrapConstrainDouble(v.get("frameRate")),
			FacingMode:       wrapConstrainDOMString(v.get("facingMode")),
			Volume:           wrapConstrainDouble(v.get("volume")),
			SampleRate:       wrapConstrainLong(v.get("sampleRate")),
			SampleSize:       wrapConstrainLong(v.get("sampleSize")),
			EchoCancellation: wrapConstrainBoolean(v.get("echoCancellation")),
			AutoGainControl:  wrapConstrainBoolean(v.get("autoGainControl")),
			NoiseSuppression: wrapConstrainBoolean(v.get("noiseSuppression")),
			Latency:          wrapConstrainDouble(v.get("latency")),
			ChannelCount:     wrapConstrainLong(v.get("channelCount")),
			DeviceId:         wrapConstrainDOMString(v.get("deviceId")),
			GroupId:          wrapConstrainDOMString(v.get("groupId")),
		}
	}
	return MediaTrackConstraintSet{}
}

func mediaTrackConstraintsSequenceToSlice(v Value) []MediaTrackConstraintSet {
	if v.valid() {
		ret := make([]MediaTrackConstraintSet, v.length())
		for i := range ret {
			ret[i] = wrapMediaTrackConstraintSet(v.index(i))
		}
		return ret
	}
	return nil
}

func wrapMediaTrackConstraints(v Value) MediaTrackConstraints {
	if v.valid() {
		return MediaTrackConstraints{
			MediaTrackConstraintSet: wrapMediaTrackConstraintSet(v),
			Advanced:                mediaTrackConstraintsSequenceToSlice(v.get("advanced")),
		}
	}
	return MediaTrackConstraints{}
}

// -------------8<---------------------------------------

func wrapLongRange(v Value) LongRange {
	if v.valid() {
		return LongRange{
			Max: v.get("max").toInt(),
			Min: v.get("min").toInt(),
		}
	}
	return LongRange{}
}

// -------------8<---------------------------------------

func wrapDoubleRange(v Value) DoubleRange {
	if v.valid() {
		return DoubleRange{
			Max: v.get("max").toFloat64(),
			Min: v.get("min").toFloat64(),
		}
	}
	return DoubleRange{}
}

// -------------8<---------------------------------------

func wrapConstrainLong(v Value) ConstrainLong {
	if v.valid() {
		return ConstrainLong{
			LongRange: wrapLongRange(v),
			Exact:     v.get("exact").toInt(),
			Ideal:     v.get("ideal").toInt(),
		}
	}
	return ConstrainLong{}
}

// -------------8<---------------------------------------

func wrapConstrainDouble(v Value) ConstrainDouble {
	if v.valid() {
		return ConstrainDouble{
			DoubleRange: wrapDoubleRange(v),
			Exact:       v.get("exact").toFloat64(),
			Ideal:       v.get("ideal").toFloat64(),
		}
	}
	return ConstrainDouble{}
}

// -------------8<---------------------------------------

func wrapConstrainBoolean(v Value) ConstrainBoolean {
	if v.valid() {
		return ConstrainBoolean{
			Exact: v.get("exact").toBool(),
			Ideal: v.get("ideal").toBool(),
		}
	}
	return ConstrainBoolean{}
}

// -------------8<---------------------------------------

func wrapConstrainDOMString(v Value) ConstrainDOMString {
	if v.valid() {
		return ConstrainDOMString{
			Exact: v.get("exact").toString(),
			Ideal: v.get("ideal").toString(),
		}
	}
	return ConstrainDOMString{}
}

// -------------8<---------------------------------------

type sourceBufferImpl struct {
	*eventTargetImpl
}

func wrapSourceBuffer(v Value) SourceBuffer {
	if v.valid() {
		return &sourceBufferImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *sourceBufferImpl) Mode() AppendMode {
	return AppendMode(p.get("mode").toString())
}

func (p *sourceBufferImpl) SetMode(mode AppendMode) {
	p.set("mode", string(mode))
}

func (p *sourceBufferImpl) Updating() bool {
	return p.get("updating").toBool()
}

func (p *sourceBufferImpl) Buffered() TimeRanges {
	return wrapTimeRanges(p.get("buffered"))
}

func (p *sourceBufferImpl) TimestampOffset() float64 {
	return p.get("timestampOffset").toFloat64()
}

func (p *sourceBufferImpl) SetTimestampOffset(t float64) {
	p.set("timestampOffset", t)
}

func (p *sourceBufferImpl) AudioTracks() AudioTrackList {
	return wrapAudioTrackList(p.get("audioTracks"))
}

func (p *sourceBufferImpl) VideoTracks() VideoTrackList {
	return wrapVideoTrackList(p.get("videoTracks"))
}

func (p *sourceBufferImpl) TextTracks() TextTrackList {
	return wrapTextTrackList(p.get("textTracks"))
}

func (p *sourceBufferImpl) AppendWindowStart() float64 {
	return p.get("appendWindowStart").toFloat64()
}

func (p *sourceBufferImpl) SetAppendWindowStart(ws float64) {
	p.set("appendWindowStart", ws)
}

func (p *sourceBufferImpl) AppendWindowEnd() float64 {
	return p.get("appendWindowEnd").toFloat64()
}

func (p *sourceBufferImpl) SetAppendWindowEnd(we float64) {
	p.set("appendWindowEnd", we)
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
	p.call("appendBuffer", JSValue(buffer))
}

func (p *sourceBufferImpl) Abort() {
	p.call("abort")
}

func (p *sourceBufferImpl) Remove(start float64, end float64) {
	p.call("remove", start, end)
}

// -------------8<---------------------------------------

type mediaSourceImpl struct {
	*eventTargetImpl
}

func wrapMediaSource(v Value) MediaSource {
	if v.valid() {
		return &mediaSourceImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaSourceImpl) SourceBuffers() SourceBufferList {
	return wrapSourceBufferList(p.get("sourceBuffers"))
}

func (p *mediaSourceImpl) ActiveSourceBuffers() SourceBufferList {
	return wrapSourceBufferList(p.get("activeSourceBuffers"))
}

func (p *mediaSourceImpl) ReadyState() ReadyState {
	return ReadyState(p.get("readyState").toString())
}

func (p *mediaSourceImpl) Duration() float64 {
	return p.get("duration").toFloat64()
}

func (p *mediaSourceImpl) SetDuration(d float64) {
	p.set("duration", d)
}

func (p *mediaSourceImpl) OnSourceOpen(fn func(Event)) EventHandler {
	return p.On("sourceopen", fn)
}

func (p *mediaSourceImpl) OnSourceEnded(fn func(Event)) EventHandler {
	return p.On("sourceended", fn)
}

func (p *mediaSourceImpl) OnSourceClose(fn func(Event)) EventHandler {
	return p.On("sourceclose", fn)
}

func (p *mediaSourceImpl) AddSourceBuffer(typ string) SourceBuffer {
	return wrapSourceBuffer(p.call("addSourceBuffer", typ))
}

func (p *mediaSourceImpl) RemoveSourceBuffer(sb SourceBuffer) {
	p.call("removeSourceBuffer", JSValue(sb))
}

func (p *mediaSourceImpl) EndOfStream(err ...EndOfStreamError) {
	switch len(err) {
	case 0:
		p.call("endOfStream")
	default:
		p.call("endOfStream", string(err[0]))
	}
}

func (p *mediaSourceImpl) SetLiveSeekableRange(start float64, end float64) {
	p.call("setLiveSeekableRange", start, end)
}

func (p *mediaSourceImpl) ClearLiveSeekableRange() {
	p.call("clearLiveSeekableRange")
}

func (p *mediaSourceImpl) IsTypeSupported(string) bool { // static
	return p.call("isTypeSupported").toBool()
}

// -------------8<---------------------------------------

type sourceBufferListImpl struct {
	*eventTargetImpl
}

func wrapSourceBufferList(v Value) SourceBufferList {
	if v.valid() {
		return &sourceBufferListImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *sourceBufferListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *sourceBufferListImpl) OnAddSourceBuffer(fn func(Event)) EventHandler {
	return p.On("addsourcebuffer", fn)
}

func (p *sourceBufferListImpl) OnRemoveSourceBuffer(fn func(Event)) EventHandler {
	return p.On("removesourcebuffer", fn)
}

func (p *sourceBufferListImpl) Item(index uint) SourceBuffer { // getter
	return wrapSourceBuffer(p.call("SourceBuffer", index))
}

// -------------8<---------------------------------------
