// +build js,wasm

package wasm

// https://www.w3.org/TR/cssom-view-1/#idl-index

type (
	// https://drafts.csswg.org/cssom/#medialist
	MediaList interface {
		MediaText() string
		SetMediaText(string)
		Length() int
		Item(int) string
		AppendMedium(string)
		DeleteMedium(string)
	}

	// https://drafts.csswg.org/cssom/#stylesheet
	StyleSheet interface {
		Type() string
		Href() string
		OwnerNode() Node // (Element or ProcessingInstruction) common interface is Node
		ParentStyleSheet() StyleSheet
		Title() string
		Media() MediaList
		Disabled() bool
		SetDisabled(bool)
	}

	// https://www.w3.org/TR/cssom-1/#linkstyle
	LinkStyle interface {
		Sheet() StyleSheet
	}

	// https://drafts.csswg.org/cssom/#cssstylesheet
	CSSStyleSheet interface {
		StyleSheet

		OwnerRule() CSSRule
		CSSRules() []CSSRule
		InsertRule(string, ...int) int
		DeleteRule(int)
	}

	// https://drafts.csswg.org/cssom/#cssrule
	CSSRule interface {
		Type() CSSRuleType
		CSSText() string
		SetCSSText(string)
		ParentRule() CSSRule
		ParentStyleSheet() CSSStyleSheet
	}

	// https://drafts.csswg.org/cssom/#stylesheetlist
	StyleSheetList interface {
		Item(int) CSSStyleSheet
		Length() int
	}

	// https://drafts.csswg.org/cssom/#cssrulelist
	CSSRuleList interface {
		Item(int) CSSRule
		Length() int
	}

	// https://drafts.csswg.org/cssom/#cssstylerule
	CSSStyleRule interface {
		CSSRule

		SelectorText() string
		SetSelectorText(string)
		Style() CSSStyleDeclaration
	}

	// https://drafts.csswg.org/cssom/#cssimportrule
	CSSImportRule interface {
		CSSRule

		Href() string
		Media() MediaList
		StyleSheet() CSSStyleSheet
	}

	// https://drafts.csswg.org/cssom/#cssgroupingrule
	CSSGroupingRule interface {
		CSSRule

		CSSRules() []CSSRule
		InsertRule(string, ...int) int
		DeleteRule(int)
	}

	// https://drafts.csswg.org/cssom/#csspagerule
	CSSPageRule interface {
		CSSGroupingRule

		SelectorText() string
		SetSelectorText(string)
		Style() CSSStyleDeclaration
	}

	// https://drafts.csswg.org/cssom/#cssmarginrule
	CSSMarginRule interface {
		CSSRule

		Name() string
		Style() CSSStyleDeclaration
	}

	// https://drafts.csswg.org/cssom/#cssnamespacerule
	CSSNamespaceRule interface {
		CSSRule

		NamespaceURI() string
		Prefix() string
	}

	// https://drafts.csswg.org/cssom/#cssstyledeclaration
	CSSStyleDeclaration interface {
		cssStyleHelper

		CSSText() string
		SetCSSText(string)
		Length() int
		Item(int) string
		PropertyValue(string) string
		PropertyPriority(string) string
		SetProperty(string, string, ...string)
		RemoveProperty(string) string
		ParentRule() CSSRule
		CSSFloat() string
		SetCSSFloat(string)
	}
)

// https://drafts.csswg.org/cssom/#namespacedef-css
func CSSEscape(ident string) string {
	return jsGlobal.invoke("CSS.escape", ident).toString()
}

// -------------8<---------------------------------------

type CSSRuleType uint

const (
	CSSRuleTypeStyle     CSSRuleType = 1
	CSSRuleTypeCharset   CSSRuleType = 2
	CSSRuleTypeImport    CSSRuleType = 3
	CSSRuleTypeMedia     CSSRuleType = 4
	CSSRuleTypeFontFace  CSSRuleType = 5
	CSSRuleTypePage      CSSRuleType = 6
	CSSRuleTypeMargin    CSSRuleType = 9
	CSSRuleTypeNamespace CSSRuleType = 10
)
