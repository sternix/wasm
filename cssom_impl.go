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
	if v.valid() {
		return &linkStyleImpl{
			Value: v,
		}
	}
	return nil
}

func (p *linkStyleImpl) Sheet() StyleSheet {
	return wrapStyleSheet(p.get("sheet"))
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
	if v.valid() {
		return &styleSheetImpl{
			Value: v,
		}
	}
	return nil
}

func (p *styleSheetImpl) Type() string {
	return p.get("type").toString()
}

func (p *styleSheetImpl) Href() string {
	return p.get("href").toString()
}

func (p *styleSheetImpl) OwnerNode() Node {
	return wrapAsNode(p.get("ownerNode"))
}

func (p *styleSheetImpl) ParentStyleSheet() StyleSheet {
	return wrapStyleSheet(p.get("parentStyleSheet"))
}

func (p *styleSheetImpl) Title() string {
	return p.get("title").toString()
}

func (p *styleSheetImpl) Media() MediaList {
	return wrapMediaList(p.get("media"))
}

func (p *styleSheetImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *styleSheetImpl) SetDisabled(d bool) {
	p.set("disabled", d)
}

// -------------8<---------------------------------------

type mediaListImpl struct {
	Value
}

func wrapMediaList(v Value) MediaList {
	if v.valid() {
		return &mediaListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *mediaListImpl) MediaText() string {
	return p.get("mediaText").toString()
}

func (p *mediaListImpl) SetMediaText(text string) {
	p.set("mediaText", text)
}

func (p *mediaListImpl) Length() int {
	return p.get("length").toInt()
}

func (p *mediaListImpl) Item(index int) string {
	return p.call("item", index).toString()
}

func (p *mediaListImpl) AppendMedium(medium string) {
	p.call("appendMedium", medium)
}

func (p *mediaListImpl) DeleteMedium(medium string) {
	p.call("deleteMedium", medium)
}

// -------------8<---------------------------------------

type cssStyleSheetImpl struct {
	*styleSheetImpl
}

func wrapCSSStyleSheet(v Value) CSSStyleSheet {
	if v.valid() {
		return &cssStyleSheetImpl{
			styleSheetImpl: newStyleSheetImpl(v),
		}
	}
	return nil
}

func (p *cssStyleSheetImpl) OwnerRule() CSSRule {
	return wrapCSSRule(p.get("ownerRule"))
}

func (p *cssStyleSheetImpl) CSSRules() []CSSRule {
	if list := wrapCSSRuleList(p.get("cssRules")); list != nil && list.Length() > 0 {
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
		return p.call("insertRule", rule).toInt()
	default:
		return p.call("insertRule", rule, index[0]).toInt()
	}
}

func (p *cssStyleSheetImpl) DeleteRule(index int) {
	p.call("deleteRule", index)
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
	if v.valid() {
		return &cssRuleImpl{
			Value: v,
		}
	}
	return nil
}

func (p *cssRuleImpl) Type() CSSRuleType {
	return CSSRuleType(uint(p.get("type").toInt()))
}

func (p *cssRuleImpl) CSSText() string {
	return p.get("cssText").toString()
}

func (p *cssRuleImpl) SetCSSText(text string) {
	p.set("cssText", text)
}

func (p *cssRuleImpl) ParentRule() CSSRule {
	return wrapCSSRule(p.get("parentRule"))
}

func (p *cssRuleImpl) ParentStyleSheet() CSSStyleSheet {
	return wrapCSSStyleSheet(p.get("parentStyleSheet"))
}

// -------------8<---------------------------------------

type styleSheetListImpl struct {
	Value
}

func wrapStyleSheetList(v Value) StyleSheetList {
	if v.valid() {
		return &styleSheetListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *styleSheetListImpl) Item(index int) CSSStyleSheet {
	return wrapCSSStyleSheet(p.call("item", index))
}

func (p *styleSheetListImpl) Length() int {
	return p.get("length").toInt()
}

// -------------8<---------------------------------------

type cssRuleList struct {
	Value
}

func wrapCSSRuleList(v Value) CSSRuleList {
	if v.valid() {
		return &cssRuleList{
			Value: v,
		}
	}
	return nil
}

func (p *cssRuleList) Item(index int) CSSRule {
	return wrapCSSRule(p.call("item", index))
}

func (p *cssRuleList) Length() int {
	return p.get("length").toInt()
}

// -------------8<---------------------------------------

type cssStyleRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSStyleRule(v Value) CSSStyleRule {
	if v.valid() {
		return &cssStyleRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssStyleRuleImpl) SelectorText() string {
	return p.get("selectorText").toString()
}

func (p *cssStyleRuleImpl) SetSelectorText(text string) {
	p.set("selectorText", text)
}

func (p *cssStyleRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.get("style"))
}

// -------------8<---------------------------------------

type cssImportRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSImportRule(v Value) CSSImportRule {
	if v.valid() {
		return &cssImportRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssImportRuleImpl) Href() string {
	return p.get("href").toString()
}

func (p *cssImportRuleImpl) Media() MediaList {
	return wrapMediaList(p.get("media"))
}

func (p *cssImportRuleImpl) StyleSheet() CSSStyleSheet {
	return wrapCSSStyleSheet(p.get("styleSheet"))
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
	if v.valid() {
		return &cssGroupingRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssGroupingRuleImpl) CSSRules() []CSSRule {
	if list := wrapCSSRuleList(p.get("cssRules")); list != nil && list.Length() > 0 {
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
		return p.call("insertRule", rule).toInt()
	default:
		return p.call("insertRule", rule, index[0]).toInt()
	}
}

func (p *cssGroupingRuleImpl) DeleteRule(index int) {
	p.call("deleteRule", index)
}

// -------------8<---------------------------------------

type cssPageRuleImpl struct {
	*cssGroupingRuleImpl
}

func wrapCSSPageRule(v Value) CSSPageRule {
	if v.valid() {
		return &cssPageRuleImpl{
			cssGroupingRuleImpl: newCSSGroupingRuleImpl(v),
		}
	}
	return nil
}

func (p *cssPageRuleImpl) SelectorText() string {
	return p.get("selectorText").toString()
}

func (p *cssPageRuleImpl) SetSelectorText(text string) {
	p.set("selectorText", text)
}

func (p *cssPageRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.get("style"))
}

// -------------8<---------------------------------------

type cssMarginRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSMarginRule(v Value) CSSMarginRule {
	if v.valid() {
		return &cssMarginRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssMarginRuleImpl) Name() string {
	return p.get("name").toString()
}

func (p *cssMarginRuleImpl) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.get("style"))
}

// -------------8<---------------------------------------

type cssNamespaceRuleImpl struct {
	*cssRuleImpl
}

func wrapCSSNamespaceRule(v Value) CSSNamespaceRule {
	if v.valid() {
		return &cssNamespaceRuleImpl{
			cssRuleImpl: newCSSRuleImpl(v),
		}
	}
	return nil
}

func (p *cssNamespaceRuleImpl) NamespaceURI() string {
	return p.get("namespaceURI").toString()
}

func (p *cssNamespaceRuleImpl) Prefix() string {
	return p.get("prefix").toString()
}

// -------------8<---------------------------------------

type cssStyleDeclarationImpl struct {
	Value
	*cssStyleHelperImpl
}

func wrapCSSStyleDeclaration(v Value) CSSStyleDeclaration {
	if v.valid() {
		p := &cssStyleDeclarationImpl{
			Value: v,
		}
		p.cssStyleHelperImpl = newCSSStyleHelperImpl(p)
		return p
	}
	return nil
}

func (p *cssStyleDeclarationImpl) CSSText() string {
	return p.get("cssText").toString()
}

func (p *cssStyleDeclarationImpl) SetCSSText(text string) {
	p.set("cssText", text)
}

func (p *cssStyleDeclarationImpl) Length() int {
	return p.get("length").toInt()
}

func (p *cssStyleDeclarationImpl) Item(index int) string {
	return p.call("item", index).toString()
}

func (p *cssStyleDeclarationImpl) PropertyValue(property string) string {
	return p.call("getPropertyValue", property).toString()
}

func (p *cssStyleDeclarationImpl) PropertyPriority(property string) string {
	return p.call("getPropertyPriority", property).toString()
}

func (p *cssStyleDeclarationImpl) SetProperty(property string, value string, priority ...string) {
	switch len(priority) {
	case 0:
		p.call("setProperty", property, value)
	default:
		p.call("setProperty", property, value, priority[0])
	}
}

func (p *cssStyleDeclarationImpl) RemoveProperty(property string) string {
	return p.call("removeProperty", property).toString()
}

func (p *cssStyleDeclarationImpl) ParentRule() CSSRule {
	return wrapCSSRule(p.get("parentRule"))
}

func (p *cssStyleDeclarationImpl) CSSFloat() string {
	return p.get("cssFloat").toString()
}

func (p *cssStyleDeclarationImpl) SetCSSFloat(cf string) {
	p.set("cssFloat", cf)
}

// -------------8<---------------------------------------
