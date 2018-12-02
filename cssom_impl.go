// +build js,wasm

package wasm

import (
	"syscall/js"
)

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

// -------------8<---------------------------------------

type styleSheetImpl struct {
	js.Value
}

func newStyleSheet(v js.Value) StyleSheet {
	if p := newStyleSheetImpl(v); p != nil {
		return p
	}
	return nil
}

func newStyleSheetImpl(v js.Value) *styleSheetImpl {
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

type cssStyleSheetImpl struct {
	*styleSheetImpl
}

func newCSSStyleSheet(v js.Value) CSSStyleSheet {
	if isNil(v) {
		return nil
	}

	return &cssStyleSheetImpl{
		styleSheetImpl: newStyleSheetImpl(v),
	}
}

func (p *cssStyleSheetImpl) OwnerRule() CSSRule {
	return newCSSRule(p.Get("ownerRule"))
}

func (p *cssStyleSheetImpl) CSSRules() []CSSRule {
	if list := newCSSRuleList(p.Get("cssRules")); list != nil && list.Length() > 0 {
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
	js.Value
}

func newCSSRule(v js.Value) CSSRule {
	if p := newCSSRuleImpl(v); p != nil {
		return p
	}
	return nil
}

func newCSSRuleImpl(v js.Value) *cssRuleImpl {
	if isNil(v) {
		return nil
	}

	return &cssRuleImpl{
		Value: v,
	}
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
	return newCSSRule(p.Get("parentRule"))
}

func (p *cssRuleImpl) ParentStyleSheet() CSSStyleSheet {
	return newCSSStyleSheet(p.Get("parentStyleSheet"))
}

// -------------8<---------------------------------------

type styleSheetListImpl struct {
	js.Value
}

func newStyleSheetList(v js.Value) StyleSheetList {
	if isNil(v) {
		return nil
	}

	return &styleSheetListImpl{
		Value: v,
	}
}

func (p *styleSheetListImpl) Item(index int) CSSStyleSheet {
	return newCSSStyleSheet(p.Call("item", index))
}

func (p *styleSheetListImpl) Length() int {
	return p.Get("length").Int()
}

// -------------8<---------------------------------------

type cssRuleList struct {
	js.Value
}

func newCSSRuleList(v js.Value) CSSRuleList {
	if isNil(v) {
		return nil
	}
	return &cssRuleList{
		Value: v,
	}
}

func (p *cssRuleList) Item(index int) CSSRule {
	return newCSSRule(p.Call("item", index))
}

func (p *cssRuleList) Length() int {
	return p.Get("length").Int()
}

// -------------8<---------------------------------------

type cssStyleRuleImpl struct {
	*cssRuleImpl
}

func newCSSStyleRule(v js.Value) CSSStyleRule {
	if isNil(v) {
		return nil
	}
	return &cssStyleRuleImpl{
		cssRuleImpl: newCSSRuleImpl(v),
	}
}

func (p *cssStyleRuleImpl) SelectorText() string {
	return p.Get("selectorText").String()
}

func (p *cssStyleRuleImpl) SetSelectorText(text string) {
	p.Set("selectorText", text)
}

func (p *cssStyleRuleImpl) Style() CSSStyleDeclaration {
	return newCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssImportRuleImpl struct {
	*cssRuleImpl
}

func newCSSImportRule(v js.Value) CSSImportRule {
	if isNil(v) {
		return nil
	}

	return &cssImportRuleImpl{
		cssRuleImpl: newCSSRuleImpl(v),
	}
}

func (p *cssImportRuleImpl) Href() string {
	return p.Get("href").String()
}

func (p *cssImportRuleImpl) Media() MediaList {
	return newMediaList(p.Get("media"))
}

func (p *cssImportRuleImpl) StyleSheet() CSSStyleSheet {
	return newCSSStyleSheet(p.Get("styleSheet"))
}

// -------------8<---------------------------------------

type cssGroupingRuleImpl struct {
	*cssRuleImpl
}

func newCSSGroupingRule(v js.Value) CSSGroupingRule {
	if p := newCSSGroupingRuleImpl(v); p != nil {
		return p
	}
	return nil
}

func newCSSGroupingRuleImpl(v js.Value) *cssGroupingRuleImpl {
	if isNil(v) {
		return nil
	}
	return &cssGroupingRuleImpl{
		cssRuleImpl: newCSSRuleImpl(v),
	}
}

func (p *cssGroupingRuleImpl) CSSRules() []CSSRule {
	if list := newCSSRuleList(p.Get("cssRules")); list != nil && list.Length() > 0 {
		ret := make([]CSSRule, list.Length())
		for i := 0; i < list.Length(); i++ {
			ret[i] = list.Item(i)
		}
		return ret
	}
	return nil
}

func (p *cssGroupingRuleImpl) insertRule(rule string, index ...int) int {
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

func newCSSPageRule(v js.Value) CSSPageRule {
	if isNil(v) {
		return nil
	}
	return &cssPageRuleImpl{
		cssGroupingRuleImpl: newCSSGroupingRuleImpl(v),
	}
}

func (p *cssPageRuleImpl) SelectorText() string {
	return p.Get("selectorText").String()
}

func (p *cssPageRuleImpl) SetSelectorText(text string) {
	p.Set("selectorText", text)
}

func (p *cssPageRuleImpl) Style() CSSStyleDeclaration {
	return newCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssMarginRuleImpl struct {
	*cssRuleImpl
}

func newCSSMarginRule(v js.Value) CSSMarginRule {
	if isNil(v) {
		return nil
	}
	return &cssMarginRuleImpl{
		cssRuleImpl: newCSSRuleImpl(v),
	}
}

func (p *cssMarginRuleImpl) Name() string {
	return p.Get("name").String()
}

func (p *cssMarginRuleImpl) Style() CSSStyleDeclaration {
	return newCSSStyleDeclaration(p.Get("style"))
}

// -------------8<---------------------------------------

type cssNamespaceRuleImpl struct {
	*cssRuleImpl
}

func newCSSNamespaceRule(v js.Value) CSSNamespaceRule {
	if isNil(v) {
		return nil
	}
	return &cssNamespaceRuleImpl{
		cssRuleImpl: newCSSRuleImpl(v),
	}
}

func (p *cssNamespaceRuleImpl) NamespaceURI() string {
	return p.Get("namespaceURI").String()
}

func (p *cssNamespaceRuleImpl) Prefix() string {
	return p.Get("prefix").String()
}

// -------------8<---------------------------------------

type cssStyleDeclarationImpl struct {
	js.Value
}

func newCSSStyleDeclaration(v js.Value) CSSStyleDeclaration {
	if isNil(v) {
		return nil
	}
	return &cssStyleDeclarationImpl{
		Value: v,
	}
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
	return newCSSRule(p.Get("parentRule"))
}

func (p *cssStyleDeclarationImpl) CSSFloat() string {
	return p.Get("cssFloat").String()
}

func (p *cssStyleDeclarationImpl) SetCSSFloat(cf string) {
	p.Set("cssFloat", cf)
}

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
