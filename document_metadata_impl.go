// +build js,wasm

package wasm

import (
	"syscall/js"
)

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

func wrapHTMLHtmlElement(v js.Value) HTMLHtmlElement {
	if isNil(v) {
		return nil
	}

	return &htmlHtmlElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
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

func wrapHTMLHeadElement(v js.Value) HTMLHeadElement {
	if isNil(v) {
		return nil
	}

	return &htmlHeadElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
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

func wrapHTMLTitleElement(v js.Value) HTMLTitleElement {
	if isNil(v) {
		return nil
	}

	return &htmlTitleElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTitleElementImpl) Text() string {
	return p.Get("text").String()
}

func (p *htmlTitleElementImpl) SetText(text string) {
	p.Set("text", text)
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

func wrapHTMLBaseElement(v js.Value) HTMLBaseElement {
	if isNil(v) {
		return nil
	}

	return &htmlBaseElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlBaseElementImpl) Href() string {
	return p.Get("href").String()
}

func (p *htmlBaseElementImpl) SetHref(href string) {
	p.Set("href", href)
}

func (p *htmlBaseElementImpl) Target() string {
	return p.Get("target").String()
}

func (p *htmlBaseElementImpl) SetTarget(target string) {
	p.Set("target", target)
}

// -------------8<---------------------------------------

type htmlLinkElementImpl struct {
	*linkStyleImpl
	*htmlElementImpl
	js.Value
}

func NewHTMLLinkElement() HTMLLinkElement {
	if el := CurrentDocument().CreateElement("link"); el != nil {
		if link, ok := el.(HTMLLinkElement); ok {
			return link
		}
	}
	return nil
}

func wrapHTMLLinkElement(v js.Value) HTMLLinkElement {
	if isNil(v) {
		return nil
	}

	return &htmlLinkElementImpl{
		linkStyleImpl:   newLinkStyleImpl(v),
		htmlElementImpl: newHTMLElementImpl(v),
		Value:           v,
	}
}

func (p *htmlLinkElementImpl) Href() string {
	return p.Get("href").String()
}

func (p *htmlLinkElementImpl) SetHref(href string) {
	p.Set("href", href)
}

func (p *htmlLinkElementImpl) CrossOrigin() string {
	return p.Get("crossOrigin").String()
}

func (p *htmlLinkElementImpl) SetCrossOrigin(crossOrigin string) {
	p.Set("crossOrigin", crossOrigin)
}

func (p *htmlLinkElementImpl) Rel() string {
	return p.Get("rel").String()
}

func (p *htmlLinkElementImpl) SetRel(rel string) {
	p.Set("rel", rel)
}

func (p *htmlLinkElementImpl) Rev() string {
	return p.Get("rev").String()
}

func (p *htmlLinkElementImpl) SetRev(rev string) {
	p.Set("rev", rev)
}

func (p *htmlLinkElementImpl) RelList() DOMTokenList {
	return wrapDOMTokenList(p.Get("relList"))
}

func (p *htmlLinkElementImpl) Media() string {
	return p.Get("media").String()
}

func (p *htmlLinkElementImpl) SetMedia(media string) {
	p.Set("media", media)
}

func (p *htmlLinkElementImpl) Nonce() string {
	return p.Get("nonce").String()
}

func (p *htmlLinkElementImpl) SetNonce(nonce string) {
	p.Set("nonce", nonce)
}

func (p *htmlLinkElementImpl) HrefLang() string {
	return p.Get("hreflang").String()
}

func (p *htmlLinkElementImpl) SetHrefLang(hreflang string) {
	p.Set("hreflang", hreflang)
}

func (p *htmlLinkElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlLinkElementImpl) SetType(typ string) {
	p.Set("type", typ)
}

func (p *htmlLinkElementImpl) Sizes() DOMTokenList {
	return wrapDOMTokenList(p.Get("sizes"))
}

func (p *htmlLinkElementImpl) ReferrerPolicy() string {
	return p.Get("referrerPolicy").String()
}

func (p *htmlLinkElementImpl) SetReferrerPolicy(referrerPolicy string) {
	p.Set("referrerPolicy", referrerPolicy)
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

func wrapHTMLMetaElement(v js.Value) HTMLMetaElement {
	if isNil(v) {
		return nil
	}

	return &htmlMetaElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlMetaElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlMetaElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlMetaElementImpl) HTTPEquiv() string {
	return p.Get("httpEquiv").String()
}

func (p *htmlMetaElementImpl) SetHTTPEquiv(httpEquiv string) {
	p.Set("httpEquiv", httpEquiv)
}

func (p *htmlMetaElementImpl) Content() string {
	return p.Get("content").String()
}

func (p *htmlMetaElementImpl) SetContent(content string) {
	p.Set("content", content)
}

// -------------8<---------------------------------------

type htmlStyleElementImpl struct {
	*htmlElementImpl
	*linkStyleImpl
	js.Value
}

func NewHTMLStyleElement() HTMLStyleElement {
	if el := CurrentDocument().CreateElement("style"); el != nil {
		if style, ok := el.(HTMLStyleElement); ok {
			return style
		}
	}
	return nil
}

func wrapHTMLStyleElement(v js.Value) HTMLStyleElement {
	if isNil(v) {
		return nil
	}

	return &htmlStyleElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
		linkStyleImpl:   newLinkStyleImpl(v),
		Value:           v,
	}
}

func (p *htmlStyleElementImpl) Media() string {
	return p.Get("media").String()
}

func (p *htmlStyleElementImpl) SetMedia(media string) {
	p.Set("media", media)
}

func (p *htmlStyleElementImpl) Nonce() string {
	return p.Get("nonce").String()
}

func (p *htmlStyleElementImpl) SetNonce(nonce string) {
	p.Set("nonce", nonce)
}

func (p *htmlStyleElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlStyleElementImpl) SetType(typ string) {
	p.Set("type", typ)
}

// -------------8<---------------------------------------
