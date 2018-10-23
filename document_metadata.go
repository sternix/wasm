// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/html52/semantics.html#htmlhtmlelement
	HTMLHtmlElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmlheadelement
	HTMLHeadElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmltitleelement
	HTMLTitleElement interface {
		HTMLElement

		Text() string
		SetText(string)
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmlbaseelement
	HTMLBaseElement interface {
		HTMLElement

		Href() string
		SetHref(string)
		Target() string
		SetTarget(string)
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmllinkelement
	HTMLLinkElement interface {
		HTMLElement
		LinkStyle

		Href() string
		SetHref(string)
		CrossOrigin() string
		SetCrossOrigin(string)
		Rel() string
		SetRel(string)
		Rev() string
		SetRev(string)
		RelList() DOMTokenList
		Media() string
		SetMedia(string)
		Nonce() string
		SetNonce(string)
		HrefLang() string
		SetHrefLang(string)
		Type() string
		SetType(string)
		Sizes() DOMTokenList
		ReferrerPolicy() string
		SetReferrerPolicy(string)
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmlmetaelement
	HTMLMetaElement interface {
		HTMLElement

		Name() string
		SetName(string)
		HTTPEquiv() string
		SetHTTPEquiv(string)
		Content() string
		SetContent(string)
	}

	// https://www.w3.org/TR/html52/document-metadata.html#htmlstyleelement
	HTMLStyleElement interface {
		HTMLElement
		LinkStyle

		Media() string
		SetMedia(string)
		Nonce() string
		SetNonce(string)
		Type() string
		SetType(string)
	}
)
