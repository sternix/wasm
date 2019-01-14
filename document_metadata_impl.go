// +build js,wasm

package wasm

// -------------8<---------------------------------------

type htmlHtmlElementImpl struct {
	*htmlElementImpl
}

func NewHTMLHtmlElement() HTMLHtmlElement {
	if el := CurrentDocument().CreateElement("html"); el != nil {
		if html, ok := el.(HTMLHtmlElement); ok {
			return html
		}
	}
	return nil
}

func wrapHTMLHtmlElement(v Value) HTMLHtmlElement {
	if v.valid() {
		return &htmlHtmlElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlHeadElementImpl struct {
	*htmlElementImpl
}

func NewHTMLHeadElement() HTMLHeadElement {
	if el := CurrentDocument().CreateElement("head"); el != nil {
		if head, ok := el.(HTMLHeadElement); ok {
			return head
		}
	}
	return nil
}

func wrapHTMLHeadElement(v Value) HTMLHeadElement {
	if v.valid() {
		return &htmlHeadElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type htmlTitleElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTitleElement() HTMLTitleElement {
	if el := CurrentDocument().CreateElement("title"); el != nil {
		if title, ok := el.(HTMLTitleElement); ok {
			return title
		}
	}
	return nil
}

func wrapHTMLTitleElement(v Value) HTMLTitleElement {
	if v.valid() {
		return &htmlTitleElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTitleElementImpl) Text() string {
	return p.get("text").toString()
}

func (p *htmlTitleElementImpl) SetText(text string) {
	p.set("text", text)
}

// -------------8<---------------------------------------

type htmlBaseElementImpl struct {
	*htmlElementImpl
}

func NewHTMLBaseElement() HTMLBaseElement {
	if el := CurrentDocument().CreateElement("base"); el != nil {
		if base, ok := el.(HTMLBaseElement); ok {
			return base
		}
	}
	return nil
}

func wrapHTMLBaseElement(v Value) HTMLBaseElement {
	if v.valid() {
		return &htmlBaseElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlBaseElementImpl) Href() string {
	return p.get("href").toString()
}

func (p *htmlBaseElementImpl) SetHref(href string) {
	p.set("href", href)
}

func (p *htmlBaseElementImpl) Target() string {
	return p.get("target").toString()
}

func (p *htmlBaseElementImpl) SetTarget(target string) {
	p.set("target", target)
}

// -------------8<---------------------------------------

type htmlLinkElementImpl struct {
	*linkStyleImpl
	*htmlElementImpl
	Value
}

func NewHTMLLinkElement() HTMLLinkElement {
	if el := CurrentDocument().CreateElement("link"); el != nil {
		if link, ok := el.(HTMLLinkElement); ok {
			return link
		}
	}
	return nil
}

func wrapHTMLLinkElement(v Value) HTMLLinkElement {
	if v.valid() {
		return &htmlLinkElementImpl{
			linkStyleImpl:   newLinkStyleImpl(v),
			htmlElementImpl: newHTMLElementImpl(v),
			Value:           v,
		}
	}
	return nil
}

func (p *htmlLinkElementImpl) Href() string {
	return p.get("href").toString()
}

func (p *htmlLinkElementImpl) SetHref(href string) {
	p.set("href", href)
}

func (p *htmlLinkElementImpl) CrossOrigin() string {
	return p.get("crossOrigin").toString()
}

func (p *htmlLinkElementImpl) SetCrossOrigin(crossOrigin string) {
	p.set("crossOrigin", crossOrigin)
}

func (p *htmlLinkElementImpl) Rel() string {
	return p.get("rel").toString()
}

func (p *htmlLinkElementImpl) SetRel(rel string) {
	p.set("rel", rel)
}

func (p *htmlLinkElementImpl) Rev() string {
	return p.get("rev").toString()
}

func (p *htmlLinkElementImpl) SetRev(rev string) {
	p.set("rev", rev)
}

func (p *htmlLinkElementImpl) RelList() DOMTokenList {
	return wrapDOMTokenList(p.get("relList"))
}

func (p *htmlLinkElementImpl) Media() string {
	return p.get("media").toString()
}

func (p *htmlLinkElementImpl) SetMedia(media string) {
	p.set("media", media)
}

func (p *htmlLinkElementImpl) Nonce() string {
	return p.get("nonce").toString()
}

func (p *htmlLinkElementImpl) SetNonce(nonce string) {
	p.set("nonce", nonce)
}

func (p *htmlLinkElementImpl) HrefLang() string {
	return p.get("hreflang").toString()
}

func (p *htmlLinkElementImpl) SetHrefLang(hreflang string) {
	p.set("hreflang", hreflang)
}

func (p *htmlLinkElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlLinkElementImpl) SetType(typ string) {
	p.set("type", typ)
}

func (p *htmlLinkElementImpl) Sizes() DOMTokenList {
	return wrapDOMTokenList(p.get("sizes"))
}

func (p *htmlLinkElementImpl) ReferrerPolicy() string {
	return p.get("referrerPolicy").toString()
}

func (p *htmlLinkElementImpl) SetReferrerPolicy(referrerPolicy string) {
	p.set("referrerPolicy", referrerPolicy)
}

// -------------8<---------------------------------------

type htmlMetaElementImpl struct {
	*htmlElementImpl
}

func NewHTMLMetaElement() HTMLMetaElement {
	if el := CurrentDocument().CreateElement("meta"); el != nil {
		if meta, ok := el.(HTMLMetaElement); ok {
			return meta
		}
	}
	return nil
}

func wrapHTMLMetaElement(v Value) HTMLMetaElement {
	if v.valid() {
		return &htmlMetaElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlMetaElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlMetaElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlMetaElementImpl) HTTPEquiv() string {
	return p.get("httpEquiv").toString()
}

func (p *htmlMetaElementImpl) SetHTTPEquiv(httpEquiv string) {
	p.set("httpEquiv", httpEquiv)
}

func (p *htmlMetaElementImpl) Content() string {
	return p.get("content").toString()
}

func (p *htmlMetaElementImpl) SetContent(content string) {
	p.set("content", content)
}

// -------------8<---------------------------------------

type htmlStyleElementImpl struct {
	*htmlElementImpl
	*linkStyleImpl
	Value
}

func NewHTMLStyleElement() HTMLStyleElement {
	if el := CurrentDocument().CreateElement("style"); el != nil {
		if style, ok := el.(HTMLStyleElement); ok {
			return style
		}
	}
	return nil
}

func wrapHTMLStyleElement(v Value) HTMLStyleElement {
	if v.valid() {
		return &htmlStyleElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
			linkStyleImpl:   newLinkStyleImpl(v),
			Value:           v,
		}
	}
	return nil
}

func (p *htmlStyleElementImpl) Media() string {
	return p.get("media").toString()
}

func (p *htmlStyleElementImpl) SetMedia(media string) {
	p.set("media", media)
}

func (p *htmlStyleElementImpl) Nonce() string {
	return p.get("nonce").toString()
}

func (p *htmlStyleElementImpl) SetNonce(nonce string) {
	p.set("nonce", nonce)
}

func (p *htmlStyleElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlStyleElementImpl) SetType(typ string) {
	p.set("type", typ)
}

// -------------8<---------------------------------------
