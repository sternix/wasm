// +build js,wasm

package wasm

// -------------8<---------------------------------------

type linkStyleImpl struct {
	Value
}

func wrapLinkStyle(v Value) LinkStyle {
	if p := newLinkStyleImpl(v); p != nil {
		return p
	}
	return nil
}

func newLinkStyleImpl(v Value) *linkStyleImpl {
	if v.Valid() {
		return &linkStyleImpl{
			Value: v,
		}
	}
	return nil
}

func (p *linkStyleImpl) Sheet() StyleSheet {
	return wrapStyleSheet(p.Get("sheet"))
}

// -------------8<---------------------------------------

type styleSheetImpl struct {
	Value
}

func wrapStyleSheet(v Value) StyleSheet {
	if p := newStyleSheetImpl(v); p != nil {
		return p
	}
	return nil
}

func newStyleSheetImpl(v Value) *styleSheetImpl {
	if v.Valid() {
		return &styleSheetImpl{
			Value: v,
		}
	}
	return nil
}

func (p *styleSheetImpl) Type() string {
	return p.Get("type").String()
}

func (p *styleSheetImpl) Href() string {
	return p.Get("href").String()
}

func (p *styleSheetImpl) OwnerNode() Node {
	return wrapAsNode(p.Get("ownerNode"))
}

func (p *styleSheetImpl) ParentStyleSheet() StyleSheet {
	return wrapStyleSheet(p.Get("parentStyleSheet"))
}

func (p *styleSheetImpl) Title() string {
	return p.Get("title").String()
}

func (p *styleSheetImpl) Media() MediaList {
	return wrapMediaList(p.Get("media"))
}

func (p *styleSheetImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *styleSheetImpl) SetDisabled(d bool) {
	p.Set("disabled", d)
}

// -------------8<---------------------------------------

type mediaListImpl struct {
	Value
}

func wrapMediaList(v Value) MediaList {
	if v.Valid() {
		return &mediaListImpl{
			Value: v,
		}
	}
	return nil
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

type cssStyleSheetImpl struct {
	*styleSheetImpl
}

func wrapCSSStyleSheet(v Value) CSSStyleSheet {
	if v.Valid() {
		return &cssStyleSheetImpl{
			styleSheetImpl: newStyleSheetImpl(v),
		}
	}
	return nil
}

func (p *cssStyleSheetImpl) OwnerRule() CSSRule {
	return wrapCSSRule(p.Get("ownerRule"))
}

func (p *cssStyleSheetImpl) CSSRules() []CSSRule {
	if list := wrapCSSRuleList(p.Get("cssRules")); list != nil && list.Length() > 0 {
		ret := make([]CSSRule, list.Length())
		for i := 0; i < list.Length(); i++ {
			ret[i] = list.Item(i)
		}
		return ret
	}
	return nil
}

func (p *cssStyleSheetImpl) InsertRule(rule string, index ...int) int {
	switch len(index) {
	case 0:
		return p.Call("insertRule", rule).Int()
	default:
		return p.Call("insertRule", rule, index[0]).Int()
	}
}

func (p *cssStyleSheetImpl) DeleteRule(index int) {
	p.Call("deleteRule", index)
}

// -------------8<---------------------------------------

type cssRuleImpl struct {
	Value
}

func wrapCSSRule(v Value) CSSRule {
	if p := newCSSRuleImpl(v); p != nil {
		return p
	}
	return nil
}

func newCSSRuleImpl(v Value) *cssRuleImpl {
	if v.Valid() {
		return &cssRuleImpl{
			Value: v,
		}
	}
	return nil
}

func (p *cssRuleImpl) Type() CSSRuleType {
	return CSSRuleType(uint(p.Get("type").Int()))
}

func (p *cssRuleImpl) CSSText() string {
	return p.Get("cssText").String()
}

func (p *cssRuleImpl) SetCSSText(text string) {
	p.Set("cssText", text)
}

func (p *cssRuleImpl) ParentRule() CSSRule {
	return wrapCSSRule(p.Get("parentRule"))
}

func (p *cssRuleImpl) ParentStyleSheet() CSSStyleSheet {
	return wrapCSSStyleSheet(p.Get("parentStyleSheet"))
}

// -------------8<---------------------------------------

type styleSheetListImpl struct {
	Value
}

func wrapStyleSheetList(v Value) StyleSheetList {
	if v.Valid() {
		return &styleSheetListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *styleSheetListImpl) Item(index int) CSSStyleSheet {
	return wrapCSSStyleSheet(p.Call("item", index))
}

func (p *styleSheetListImpl) Length() int {
	return p.Get("length").Int()
}

// -------------8<---------------------------------------

type cssRuleList struct {
	Value
}

func wrapCSSRuleList(v Value) CSSRuleList {
	if v.Valid() {
		return &cssRuleList{
			Value: v,
		}
	}
	return nil
}

func (p *cssRuleList) Item(index int) CSSRule {
	return wrapCSSRule(p.Call("item", index))
}

func (p *cssRuleList) Length() int {
	return p.Get("length").Int()
}

// -------------8<---------------------------------------

type cssStyleRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSStyleRule(v Value) CSSStyleRule {
	if v.Valid() {
		return &cssStyleRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssStyleRuleImpl) SelectorText() string {
	return p.Get("selectorText").String()
}

func (p *cssStyleRuleImpl) SetSelectorText(text string) {
	p.Set("selectorText", text)
}

func (p *cssStyleRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssImportRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSImportRule(v Value) CSSImportRule {
	if v.Valid() {
		return &cssImportRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssImportRuleImpl) Href() string {
	return p.Get("href").String()
}

func (p *cssImportRuleImpl) Media() MediaList {
	return wrapMediaList(p.Get("media"))
}

func (p *cssImportRuleImpl) StyleSheet() CSSStyleSheet {
	return wrapCSSStyleSheet(p.Get("styleSheet"))
}

// -------------8<---------------------------------------

type cssGroupingRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSGroupingRule(v Value) CSSGroupingRule {
	if p := newCSSGroupingRuleImpl(v); p != nil {
		return p
	}
	return nil
}

func newCSSGroupingRuleImpl(v Value) *cssGroupingRuleImpl {
	if v.Valid() {
		return &cssGroupingRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssGroupingRuleImpl) CSSRules() []CSSRule {
	if list := wrapCSSRuleList(p.Get("cssRules")); list != nil && list.Length() > 0 {
		ret := make([]CSSRule, list.Length())
		for i := 0; i < list.Length(); i++ {
			ret[i] = list.Item(i)
		}
		return ret
	}
	return nil
}

func (p *cssGroupingRuleImpl) InsertRule(rule string, index ...int) int {
	switch len(index) {
	case 0:
		return p.Call("insertRule", rule).Int()
	default:
		return p.Call("insertRule", rule, index[0]).Int()
	}
}

func (p *cssGroupingRuleImpl) DeleteRule(index int) {
	p.Call("deleteRule", index)
}

// -------------8<---------------------------------------

type cssPageRuleImpl struct {
	*cssGroupingRuleImpl
}

func wrapCSSPageRule(v Value) CSSPageRule {
	if v.Valid() {
		return &cssPageRuleImpl{
			cssGroupingRuleImpl: newCSSGroupingRuleImpl(v),
		}
	}
	return nil
}

func (p *cssPageRuleImpl) SelectorText() string {
	return p.Get("selectorText").String()
}

func (p *cssPageRuleImpl) SetSelectorText(text string) {
	p.Set("selectorText", text)
}

func (p *cssPageRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssMarginRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSMarginRule(v Value) CSSMarginRule {
	if v.Valid() {
		return &cssMarginRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssMarginRuleImpl) Name() string {
	return p.Get("name").String()
}

func (p *cssMarginRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssNamespaceRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSNamespaceRule(v Value) CSSNamespaceRule {
	if v.Valid() {
		return &cssNamespaceRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssNamespaceRuleImpl) NamespaceURI() string {
	return p.Get("namespaceURI").String()
}

func (p *cssNamespaceRuleImpl) Prefix() string {
	return p.Get("prefix").String()
}

// -------------8<---------------------------------------

type cssStyleDeclarationImpl struct {
	Value
	*cssStyleHelperImpl
}

func wrapCSSStyleDeclaration(v Value) CSSStyleDeclaration {
	if v.Valid() {
		p := &cssStyleDeclarationImpl{
			Value: v,
		}
		p.cssStyleHelperImpl = newCSSStyleHelperImpl(p)
		return p
	}
	return nil
}

func (p *cssStyleDeclarationImpl) CSSText() string {
	return p.Get("cssText").String()
}

func (p *cssStyleDeclarationImpl) SetCSSText(text string) {
	p.Set("cssText", text)
}

func (p *cssStyleDeclarationImpl) Length() int {
	return p.Get("length").Int()
}

func (p *cssStyleDeclarationImpl) Item(index int) string {
	return p.Call("item", index).String()
}

func (p *cssStyleDeclarationImpl) PropertyValue(property string) string {
	return p.Call("getPropertyValue", property).String()
}

func (p *cssStyleDeclarationImpl) PropertyPriority(property string) string {
	return p.Call("getPropertyPriority", property).String()
}

func (p *cssStyleDeclarationImpl) SetProperty(property string, value string, priority ...string) {
	switch len(priority) {
	case 0:
		p.Call("setProperty", property, value)
	default:
		p.Call("setProperty", property, value, priority[0])
	}
}

func (p *cssStyleDeclarationImpl) RemoveProperty(property string) string {
	return p.Call("removeProperty", property).String()
}

func (p *cssStyleDeclarationImpl) ParentRule() CSSRule {
	return wrapCSSRule(p.Get("parentRule"))
}

func (p *cssStyleDeclarationImpl) CSSFloat() string {
	return p.Get("cssFloat").String()
}

func (p *cssStyleDeclarationImpl) SetCSSFloat(cf string) {
	p.Set("cssFloat", cf)
}

// -------------8<---------------------------------------
