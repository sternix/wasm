// +build js,wasm

package wasm

import (
	"time"
)

// -------------8<---------------------------------------

type htmlFormElementImpl struct {
	*htmlElementImpl
}

func NewHTMLFormElement() HTMLFormElement {
	if el := CurrentDocument().CreateElement("form"); el != nil {
		if form, ok := el.(HTMLFormElement); ok {
			return form
		}
	}
	return nil
}

func wrapHTMLFormElement(v Value) HTMLFormElement {
	if v.valid() {
		return &htmlFormElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlFormElementImpl) AcceptCharset() string {
	return p.get("acceptCharset").toString()
}

func (p *htmlFormElementImpl) SetAcceptCharset(ch string) {
	p.set("acceptCharset", ch)
}

func (p *htmlFormElementImpl) Action() string {
	return p.get("action").toString()
}

func (p *htmlFormElementImpl) SetAction(a string) {
	p.set("action", a)
}

func (p *htmlFormElementImpl) Autocomplete() string {
	return p.get("autocomplete").toString()
}

func (p *htmlFormElementImpl) SetAutocomplete(ac string) {
	p.set("autocomplete", ac)
}

func (p *htmlFormElementImpl) Enctype() string {
	return p.get("enctype").toString()
}

func (p *htmlFormElementImpl) SetEnctype(et string) {
	p.set("enctype", et)
}

func (p *htmlFormElementImpl) Encoding() string {
	return p.get("encoding").toString()
}

func (p *htmlFormElementImpl) SetEncoding(enc string) {
	p.set("encoding", enc)
}

func (p *htmlFormElementImpl) Method() string {
	return p.get("method").toString()
}

func (p *htmlFormElementImpl) SetMethod(m string) {
	p.set("method", m)
}

func (p *htmlFormElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlFormElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlFormElementImpl) NoValidate() bool {
	return p.get("noValidate").toBool()
}

func (p *htmlFormElementImpl) SetNoValidate(b bool) {
	p.set("noValidate", b)
}

func (p *htmlFormElementImpl) Target() string {
	return p.get("target").toString()
}

func (p *htmlFormElementImpl) SetTarget(t string) {
	p.set("target", t)
}

func (p *htmlFormElementImpl) Elements() HTMLFormControlsCollection {
	return wrapHTMLFormControlsCollection(p.get("elements"))
}

func (p *htmlFormElementImpl) Submit() {
	p.call("submit")
}

func (p *htmlFormElementImpl) Reset() {
	p.call("reset")
}

func (p *htmlFormElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlFormElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

// -------------8<---------------------------------------

type htmlFormControlsCollectionImpl struct {
	*htmlCollectionImpl
}

func wrapHTMLFormControlsCollection(v Value) HTMLFormControlsCollection {
	if v.valid() {
		return &htmlFormControlsCollectionImpl{
			htmlCollectionImpl: newHTMLCollectionImpl(v),
		}
	}
	return nil
}

//NOTE overriden namedbyItem
func (p *htmlFormControlsCollectionImpl) ItemByName(item string) HTMLFormControl {
	return wrapHTMLFormControl(p.call("namedItem", item))
}

// -------------8<---------------------------------------

type htmlFormControlImpl struct {
	Value
}

func wrapHTMLFormControl(v Value) HTMLFormControl {
	if v.valid() {
		return &htmlFormControlImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type radioNodeListImpl struct {
	*nodeListImpl
}

func wrapRadioNodeList(v Value) RadioNodeList {
	if v.valid() {
		return &radioNodeListImpl{
			nodeListImpl: newNodeListImpl(v),
		}
	}
	return nil
}

func (p *radioNodeListImpl) Value() string {
	return p.get("value").toString()
}

func (p *radioNodeListImpl) SetValue(value string) {
	p.set("value", value)
}

// -------------8<---------------------------------------

type htmlLabelElementImpl struct {
	*htmlElementImpl
}

func NewHTMLLabelElement() HTMLLabelElement {
	if el := CurrentDocument().CreateElement("label"); el != nil {
		if label, ok := el.(HTMLLabelElement); ok {
			return label
		}
	}
	return nil
}

func wrapHTMLLabelElement(v Value) HTMLLabelElement {
	if v.valid() {
		return &htmlLabelElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlLabelElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlLabelElementImpl) HtmlFor() string {
	return p.get("htmlFor").toString()
}

func (p *htmlLabelElementImpl) SetHtmlFor(hf string) {
	p.set("htmlFor", hf)
}

func (p *htmlLabelElementImpl) Control() HTMLElement {
	return wrapHTMLElement(p.get("control"))
}

// -------------8<---------------------------------------

type htmlInputElementImpl struct {
	*htmlElementImpl
}

var htmlInputElementTypeMap = map[string]string{
	"button":         "button",
	"checkbox":       "checkbox",
	"color":          "color",
	"date":           "date",
	"datetime-local": "datetime-local",
	"email":          "email",
	"file":           "file",
	"hidden":         "hidden",
	"image":          "image",
	"month":          "month",
	"number":         "number",
	"password":       "password",
	"radio":          "radio",
	"range":          "range",
	"reset":          "reset",
	"search":         "search",
	"submit":         "submit",
	"tel":            "tel",
	"text":           "text",
	"time":           "time",
	"url":            "url",
	"week":           "week",
}

func NewHTMLInputElement(typ ...string) HTMLInputElement {
	if el := CurrentDocument().CreateElement("input"); el != nil {
		if input, ok := el.(HTMLInputElement); ok {
			if len(typ) > 0 {
				if t := htmlInputElementTypeMap[typ[0]]; t != "" {
					input.SetType(t)
				}
			}
			return input
		}
	}
	return nil
}

func wrapHTMLInputElement(v Value) HTMLInputElement {
	if v.valid() {
		return &htmlInputElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlInputElementImpl) Accept() string {
	return p.get("accept").toString()
}

func (p *htmlInputElementImpl) SetAccept(a string) {
	p.set("accept", a)
}

func (p *htmlInputElementImpl) Alt() string {
	return p.get("alt").toString()
}

func (p *htmlInputElementImpl) SetAlt(a string) {
	p.set("alt", a)
}

func (p *htmlInputElementImpl) Autocomplete() string {
	return p.get("autocomplete").toString()
}

func (p *htmlInputElementImpl) SetAutocomplete(ac string) {
	p.set("autocomplete", ac)
}

func (p *htmlInputElementImpl) Autofocus() bool {
	return p.get("autofocus").toBool()
}

func (p *htmlInputElementImpl) SetAutofocus(af bool) {
	p.set("autofocus", af)
}

func (p *htmlInputElementImpl) DefaultChecked() bool {
	return p.get("defaultChecked").toBool()
}

func (p *htmlInputElementImpl) SetDefaultChecked(dc bool) {
	p.set("defaultChecked", dc)
}

func (p *htmlInputElementImpl) Checked() bool {
	return p.get("checked").toBool()
}

func (p *htmlInputElementImpl) SetChecked(c bool) {
	p.set("checked", c)
}

func (p *htmlInputElementImpl) DirName() string {
	return p.get("dirName").toString()
}

func (p *htmlInputElementImpl) SetDirName(dn string) {
	p.set("dirName", dn)
}

func (p *htmlInputElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlInputElementImpl) SetDisabled(d bool) {
	p.set("disabled", d)
}

func (p *htmlInputElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlInputElementImpl) Files() []File {
	return fileListToSlice(p.get("files"))
}

func (p *htmlInputElementImpl) FormAction() string {
	return p.get("formAction").toString()
}

func (p *htmlInputElementImpl) SetFormAction(fa string) {
	p.set("formAction", fa)
}

func (p *htmlInputElementImpl) FormEnctype() string {
	return p.get("formEnctype").toString()
}

func (p *htmlInputElementImpl) SetFormEnctype(fe string) {
	p.set("formEnctype", fe)
}

func (p *htmlInputElementImpl) FormMethod() string {
	return p.get("formMethod").toString()
}

func (p *htmlInputElementImpl) SetFormMethod(fm string) {
	p.set("formMethod", fm)
}

func (p *htmlInputElementImpl) FormNoValidate() bool {
	return p.get("formNoValidate").toBool()
}

func (p *htmlInputElementImpl) SetFormNoValidate(b bool) {
	p.set("formNoValidate", b)
}

func (p *htmlInputElementImpl) FormTarget() string {
	return p.get("formTarget").toString()
}

func (p *htmlInputElementImpl) SetFormTarget(ft string) {
	p.set("formTarget", ft)
}

func (p *htmlInputElementImpl) Height() uint {
	return p.get("height").toUint()
}

func (p *htmlInputElementImpl) SetHeight(h uint) {
	p.set("height", h)
}

func (p *htmlInputElementImpl) Indeterminate() bool {
	return p.get("indeterminate").toBool()
}

func (p *htmlInputElementImpl) SetIndeterminate(b bool) {
	p.set("indeterminate", b)
}

func (p *htmlInputElementImpl) List() HTMLElement {
	return wrapHTMLElement(p.get("list"))
}

func (p *htmlInputElementImpl) Max() string {
	return p.get("max").toString()
}

func (p *htmlInputElementImpl) SetMax(m string) {
	p.set("max", m)
}

func (p *htmlInputElementImpl) MaxLength() int {
	return p.get("maxLength").toInt()
}

func (p *htmlInputElementImpl) SetMaxLength(m int) {
	p.set("maxLength", m)
}

func (p *htmlInputElementImpl) Min() string {
	return p.get("min").toString()
}

func (p *htmlInputElementImpl) SetMin(m string) {
	p.set("min", m)
}

func (p *htmlInputElementImpl) MinLength() int {
	return p.get("minLength").toInt()
}

func (p *htmlInputElementImpl) SetMinLength(m int) {
	p.set("minLength", m)
}

func (p *htmlInputElementImpl) Multiple() bool {
	return p.get("multiple").toBool()
}

func (p *htmlInputElementImpl) SetMultiple(b bool) {
	p.set("multiple", b)
}

func (p *htmlInputElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlInputElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlInputElementImpl) Pattern() string {
	return p.get("pattern").toString()
}

func (p *htmlInputElementImpl) SetPattern(pattern string) {
	p.set("pattern", pattern)
}

func (p *htmlInputElementImpl) Placeholder() string {
	return p.get("placeholder").toString()
}

func (p *htmlInputElementImpl) SetPlaceholder(ph string) {
	p.set("placeholder", ph)
}

func (p *htmlInputElementImpl) ReadOnly() bool {
	return p.get("readOnly").toBool()
}

func (p *htmlInputElementImpl) SetReadOnly(b bool) {
	p.set("readOnly", b)
}

func (p *htmlInputElementImpl) Required() bool {
	return p.get("_required").toBool()
}

func (p *htmlInputElementImpl) SetRequired(b bool) {
	p.set("_required", b)
}

func (p *htmlInputElementImpl) Size() uint {
	return p.get("size").toUint()
}

func (p *htmlInputElementImpl) SetSize(s uint) {
	p.set("size", s)
}

func (p *htmlInputElementImpl) Src() string {
	return p.get("src").toString()
}

func (p *htmlInputElementImpl) SetSrc(src string) {
	p.set("src", src)
}

func (p *htmlInputElementImpl) Step() string {
	return p.get("step").toString()
}

func (p *htmlInputElementImpl) SetStep(s string) {
	p.set("step", s)
}

func (p *htmlInputElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlInputElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlInputElementImpl) DefaultValue() string {
	return p.get("defaultValue").toString()
}

func (p *htmlInputElementImpl) SetDefaultValue(dv string) {
	p.set("defaultValue", dv)
}

func (p *htmlInputElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlInputElementImpl) SetValue(value string) {
	p.set("value", value)
}

func (p *htmlInputElementImpl) ValueAsDate() time.Time {
	//TODO: test it
	return jsDateToTime(p.get("valueAsDate"))
}

func (p *htmlInputElementImpl) SetValueAsDate(t time.Time) {
	// TODO: test it
	d := jsDate.jsNew()
	d.call("setTime", t.Unix())
	p.set("valueAsDate", d)
}

func (p *htmlInputElementImpl) ValueAsNumber() float64 {
	return p.get("valueAsNumber").toFloat64()
}

func (p *htmlInputElementImpl) SetValueAsNumber(n float64) {
	p.set("valueAsNumber", n)
}

func (p *htmlInputElementImpl) Width() uint {
	return p.get("width").toUint()
}

func (p *htmlInputElementImpl) SetWidth(w uint) {
	p.set("width", w)
}

func (p *htmlInputElementImpl) StepUp(n ...int) {
	switch len(n) {
	case 0:
		p.call("stepUp")
	default:
		p.call("stepUp", n[0])
	}
}

func (p *htmlInputElementImpl) StepDown(n ...int) {
	switch len(n) {
	case 0:
		p.call("stepDown")
	default:
		p.call("stepDown", n[0])
	}
}

func (p *htmlInputElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlInputElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlInputElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlInputElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlInputElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlInputElementImpl) SetCustomValidity(cv string) {
	p.call("setCustomValidity", cv)
}

func (p *htmlInputElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

func (p *htmlInputElementImpl) Select() {
	p.call("select")
}

func (p *htmlInputElementImpl) SelectionStart() uint {
	return p.get("selectionStart").toUint()
}

func (p *htmlInputElementImpl) SetSelectionStart(se uint) {
	p.set("selectionStart", se)
}

func (p *htmlInputElementImpl) SelectionEnd() uint {
	return p.get("selectionEnd").toUint()
}

func (p *htmlInputElementImpl) SetSelectionEnd(se uint) {
	p.set("selectionEnd", se)
}

func (p *htmlInputElementImpl) SelectionDirection() string {
	return p.get("selectionDirection").toString()
}

func (p *htmlInputElementImpl) SetSelectionDirection(sd string) {
	p.set("selectionDirection", sd)
}

func (p *htmlInputElementImpl) SetRangeText(r string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("setRangeText", r)
	case 2:
		if start, ok := args[0].(uint); ok {
			if end, ok := args[1].(uint); ok {
				p.call("setRangeText", r, start, end)
			}
		}
	case 3:
		if start, ok := args[0].(uint); ok {
			if end, ok := args[1].(uint); ok {
				if selectionMode, ok := args[2].(SelectionMode); ok {
					p.call("setRangeText", r, start, end, string(selectionMode))
				}
			}
		}
	}
}

func (p *htmlInputElementImpl) SetSelectionRange(start uint, end uint, direction ...string) {
	switch len(direction) {
	case 0:
		p.call("setSelectionRange", start, end)
	default:
		p.call("setSelectionRange", start, end, direction[0])

	}
}

// -------------8<---------------------------------------

type htmlButtonElementImpl struct {
	*htmlElementImpl
}

func NewHTMLButtonElement() HTMLButtonElement {
	if el := CurrentDocument().CreateElement("button"); el != nil {
		if button, ok := el.(HTMLButtonElement); ok {
			return button
		}
	}
	return nil
}

func wrapHTMLButtonElement(v Value) HTMLButtonElement {
	if v.valid() {
		return &htmlButtonElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlButtonElementImpl) Autofocus() bool {
	return p.get("autofocus").toBool()
}

func (p *htmlButtonElementImpl) SetAutofocus(b bool) {
	p.set("autofocus", b)
}

func (p *htmlButtonElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlButtonElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlButtonElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlButtonElementImpl) FormAction() string {
	return p.get("formAction").toString()
}

func (p *htmlButtonElementImpl) SetFormAction(fa string) {
	p.set("formAction", fa)
}

func (p *htmlButtonElementImpl) FormEnctype() string {
	return p.get("formEnctype").toString()
}

func (p *htmlButtonElementImpl) SetFormEnctype(fe string) {
	p.set("formEnctype", fe)
}

func (p *htmlButtonElementImpl) FormMethod() string {
	return p.get("formMethod").toString()
}

func (p *htmlButtonElementImpl) SetFormMethod(fm string) {
	p.set("formMethod", fm)
}

func (p *htmlButtonElementImpl) FormNoValidate() bool {
	return p.get("formNoValidate").toBool()
}

func (p *htmlButtonElementImpl) SetFormNoValidate(b bool) {
	p.set("formNoValidate", b)
}

func (p *htmlButtonElementImpl) FormTarget() string {
	return p.get("formTarget").toString()
}

func (p *htmlButtonElementImpl) SetFormTarget(ft string) {
	p.set("formTarget", ft)
}

func (p *htmlButtonElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlButtonElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlButtonElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlButtonElementImpl) SetType(t string) {
	p.set("type", t)
}

func (p *htmlButtonElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlButtonElementImpl) SetValue(v string) {
	p.set("value", v)
}

func (p *htmlButtonElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlButtonElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlButtonElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlButtonElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlButtonElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlButtonElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

func (p *htmlButtonElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

// -------------8<---------------------------------------

type htmlSelectElementImpl struct {
	*htmlElementImpl
}

func NewHTMLSelectElement() HTMLSelectElement {
	if el := CurrentDocument().CreateElement("select"); el != nil {
		if sel, ok := el.(HTMLSelectElement); ok {
			return sel
		}
	}
	return nil
}

func wrapHTMLSelectElement(v Value) HTMLSelectElement {
	if v.valid() {
		return &htmlSelectElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlSelectElementImpl) Autocomplete() string {
	return p.get("autocomplete").toString()
}

func (p *htmlSelectElementImpl) SetAutocomplete(ac string) {
	p.set("autocomplete", ac)
}

func (p *htmlSelectElementImpl) Autofocus() bool {
	return p.get("autofocus").toBool()
}

func (p *htmlSelectElementImpl) SetAutofocus(af bool) {
	p.set("autofocus", af)
}

func (p *htmlSelectElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlSelectElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlSelectElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlSelectElementImpl) Multiple() bool {
	return p.get("multiple").toBool()
}

func (p *htmlSelectElementImpl) SetMultiple(m bool) {
	p.set("multiple", m)
}

func (p *htmlSelectElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlSelectElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlSelectElementImpl) Required() bool {
	return p.get("_required").toBool()
}

func (p *htmlSelectElementImpl) SetRequired(b bool) {
	p.set("_required", b)
}

func (p *htmlSelectElementImpl) Size() uint {
	return p.get("size").toUint()
}

func (p *htmlSelectElementImpl) SetSize(s uint) {
	p.set("size", s)
}

func (p *htmlSelectElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlSelectElementImpl) Options() HTMLOptionsCollection {
	return wrapHTMLOptionsCollection(p.get("options"))
}

func (p *htmlSelectElementImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *htmlSelectElementImpl) SetLength(l uint) {
	p.set("length", l)
}

func (p *htmlSelectElementImpl) Item(index uint) Element {
	return wrapAsElement(p.call("item", index))
}

func (p *htmlSelectElementImpl) NamedItem(name string) HTMLOptionElement {
	return wrapHTMLOptionElement(p.call("namedItem", name))
}

func (p *htmlSelectElementImpl) Add(element HTMLElement, before ...interface{}) {
	switch len(before) {
	case 0:
		p.call("add", JSValueOf(element))
	default:
		switch x := before[0].(type) {
		case HTMLElement:
			p.call("add", JSValueOf(element), JSValueOf(x))
		case int:
			p.call("add", JSValueOf(element), x)
		}
	}
}

func (p *htmlSelectElementImpl) RemoveByIndex(index int) {
	p.call("remove", index)
}

func (p *htmlSelectElementImpl) Set(index uint, option HTMLOptionElement) {
	// TODO
	// which method ??
	//  setter void (unsigned long index, HTMLOptionElement? option);
}

func (p *htmlSelectElementImpl) SelectedOptions() []HTMLOptionElement {
	return htmlCollectionToHTMLOptionElementSlice(p.get("selectedOptions"))
}

func (p *htmlSelectElementImpl) SelectedIndex() int {
	return p.get("selectedIndex").toInt()
}

func (p *htmlSelectElementImpl) SetSelectedIndex(index int) {
	p.set("selectedIndex", index)
}

func (p *htmlSelectElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlSelectElementImpl) SetValue(v string) {
	p.set("value", v)
}

func (p *htmlSelectElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlSelectElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlSelectElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlSelectElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlSelectElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlSelectElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

func (p *htmlSelectElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

// -------------8<---------------------------------------

type htmlOptionsCollectionImpl struct {
	*htmlCollectionImpl
}

func wrapHTMLOptionsCollection(v Value) HTMLOptionsCollection {
	if v.valid() {
		return &htmlOptionsCollectionImpl{
			htmlCollectionImpl: newHTMLCollectionImpl(v),
		}
	}
	return nil
}

func (p *htmlOptionsCollectionImpl) Length() uint {
	return p.get("length").toUint()
}

/* TODO setter
func (p *htmlOptionsCollectionImpl) Set(index uint,option HTMLOptionElement) {
	p.call("")
}
*/

func (p *htmlOptionsCollectionImpl) Add(element HTMLElement, before ...interface{}) {
	switch len(before) {
	case 0:
		p.call("add", JSValueOf(element))
	default:
		switch x := before[0].(type) {
		case HTMLElement:
			p.call("add", JSValueOf(element), JSValueOf(x))
		case int:
			p.call("add", JSValueOf(element), x)
		}
	}
}

func (p *htmlOptionsCollectionImpl) Remove(index int) {
	p.call("remove", index)
}

func (p *htmlOptionsCollectionImpl) SelectedIndex() int {
	return p.get("selectedIndex").toInt()
}

func (p *htmlOptionsCollectionImpl) SetSelectedIndex(i int) {
	p.set("selectedIndex", i)
}

// -------------8<---------------------------------------

type htmlDataListElementImpl struct {
	*htmlElementImpl
}

func NewHTMLDataListElement() HTMLDataListElement {
	if el := CurrentDocument().CreateElement("datalist"); el != nil {
		if datalist, ok := el.(HTMLDataListElement); ok {
			return datalist
		}
	}
	return nil
}

func wrapHTMLDataListElement(v Value) HTMLDataListElement {
	if v.valid() {
		return &htmlDataListElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlDataListElementImpl) Options() []HTMLOptionElement {
	return htmlCollectionToHTMLOptionElementSlice(p.get("options"))
}

// -------------8<---------------------------------------

type htmlOptGroupElementImpl struct {
	*htmlElementImpl
}

func NewHTMLOptGroupElement() HTMLOptGroupElement {
	if el := CurrentDocument().CreateElement("optgroup"); el != nil {
		if optgroup, ok := el.(HTMLOptGroupElement); ok {
			return optgroup
		}
	}
	return nil
}

func wrapHTMLOptGroupElement(v Value) HTMLOptGroupElement {
	if v.valid() {
		return &htmlOptGroupElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlOptGroupElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlOptGroupElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlOptGroupElementImpl) Label() string {
	return p.get("label").toString()
}

func (p *htmlOptGroupElementImpl) SetLabel(lbl string) {
	p.set("label", lbl)
}

// -------------8<---------------------------------------

type htmlOptionElementImpl struct {
	*htmlElementImpl
}

func NewHTMLOptionElement() HTMLOptionElement {
	if el := CurrentDocument().CreateElement("option"); el != nil {
		if option, ok := el.(HTMLOptionElement); ok {
			return option
		}
	}
	return nil
}

func wrapHTMLOptionElement(v Value) HTMLOptionElement {
	if v.valid() {
		return &htmlOptionElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlOptionElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlOptionElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlOptionElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlOptionElementImpl) Label() string {
	return p.get("label").toString()
}

func (p *htmlOptionElementImpl) SetLabel(lbl string) {
	p.set("label", lbl)
}

func (p *htmlOptionElementImpl) DefaultSelected() bool {
	return p.get("defaultSelected").toBool()
}

func (p *htmlOptionElementImpl) SetDefaultSelected(b bool) {
	p.set("defaultSelected", b)
}

func (p *htmlOptionElementImpl) Selected() bool {
	return p.get("selected").toBool()
}

func (p *htmlOptionElementImpl) SetSelected(b bool) {
	p.set("selected", b)
}

func (p *htmlOptionElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlOptionElementImpl) SetValue(v string) {
	p.set("value", v)
}

func (p *htmlOptionElementImpl) Text() string {
	return p.get("text").toString()
}

func (p *htmlOptionElementImpl) SetText(t string) {
	p.set("text", t)
}

func (p *htmlOptionElementImpl) Index() int {
	return p.get("index").toInt()
}

// -------------8<---------------------------------------

type htmlTextAreaElementImpl struct {
	*htmlElementImpl
}

func NewHTMLTextAreaElement() HTMLTextAreaElement {
	if el := CurrentDocument().CreateElement("textarea"); el != nil {
		if textarea, ok := el.(HTMLTextAreaElement); ok {
			return textarea
		}
	}
	return nil
}

func wrapHTMLTextAreaElement(v Value) HTMLTextAreaElement {
	if v.valid() {
		return &htmlTextAreaElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlTextAreaElementImpl) Autocomplete() string {
	return p.get("autocomplete").toString()
}

func (p *htmlTextAreaElementImpl) SetAutocomplete(ac string) {
	p.set("autocomplete", ac)
}

func (p *htmlTextAreaElementImpl) Autofocus() bool {
	return p.get("autofocus").toBool()
}

func (p *htmlTextAreaElementImpl) SetAutofocus(b bool) {
	p.set("autofocus", b)
}

func (p *htmlTextAreaElementImpl) Cols() uint {
	return p.get("cols").toUint()
}

func (p *htmlTextAreaElementImpl) SetCols(c uint) {
	p.set("cols", c)
}

func (p *htmlTextAreaElementImpl) DirName() string {
	return p.get("dirName").toString()
}

func (p *htmlTextAreaElementImpl) SetDirName(dn string) {
	p.set("dirName", dn)
}

func (p *htmlTextAreaElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlTextAreaElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlTextAreaElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlTextAreaElementImpl) MaxLength() int {
	return p.get("maxLength").toInt()
}

func (p *htmlTextAreaElementImpl) SetMaxLength(m int) {
	p.set("maxLength", m)
}

func (p *htmlTextAreaElementImpl) MinLength() int {
	return p.get("minLength").toInt()
}

func (p *htmlTextAreaElementImpl) SetMinLength(m int) {
	p.set("minLength", m)
}

func (p *htmlTextAreaElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlTextAreaElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlTextAreaElementImpl) Placeholder() string {
	return p.get("placeholder").toString()
}

func (p *htmlTextAreaElementImpl) SetPlaceholder(h string) {
	p.set("placeholder", h)
}

func (p *htmlTextAreaElementImpl) ReadOnly() bool {
	return p.get("readOnly").toBool()
}

func (p *htmlTextAreaElementImpl) SetReadOnly(b bool) {
	p.set("readOnly", b)
}

func (p *htmlTextAreaElementImpl) Required() bool {
	return p.get("_required").toBool()
}

func (p *htmlTextAreaElementImpl) SetRequired(b bool) {
	p.set("_required", b)
}

func (p *htmlTextAreaElementImpl) Rows() uint {
	return p.get("rows").toUint()
}

func (p *htmlTextAreaElementImpl) SetRows(r uint) {
	p.set("rows", r)
}

func (p *htmlTextAreaElementImpl) Wrap() string {
	return p.get("wrap").toString()
}

func (p *htmlTextAreaElementImpl) SetWrap(w string) {
	p.set("wrap", w)
}

func (p *htmlTextAreaElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlTextAreaElementImpl) DefaultValue() string {
	return p.get("defaultValue").toString()
}

func (p *htmlTextAreaElementImpl) SetDefaultValue(dv string) {
	p.set("defaultValue", dv)
}

func (p *htmlTextAreaElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlTextAreaElementImpl) SetValue(v string) {
	p.set("value", v)
}

func (p *htmlTextAreaElementImpl) TextLength() uint {
	return p.get("textLength").toUint()
}

func (p *htmlTextAreaElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlTextAreaElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlTextAreaElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlTextAreaElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlTextAreaElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlTextAreaElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

func (p *htmlTextAreaElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

func (p *htmlTextAreaElementImpl) Select() {
	p.call("select")
}

func (p *htmlTextAreaElementImpl) SelectionStart() uint {
	return p.get("selectionStart").toUint()
}

func (p *htmlTextAreaElementImpl) SetSelectionStart(ss uint) {
	p.set("selectionStart", ss)
}

func (p *htmlTextAreaElementImpl) SelectionEnd() uint {
	return p.get("selectionEnd").toUint()
}

func (p *htmlTextAreaElementImpl) SetSelectionEnd(se uint) {
	p.set("selectionEnd", se)
}

func (p *htmlTextAreaElementImpl) SelectionDirection() string {
	return p.get("selectionDirection").toString()
}

func (p *htmlTextAreaElementImpl) SetSelectionDirection(sd string) {
	p.set("selectionDirection", sd)
}

func (p *htmlTextAreaElementImpl) SetRangeText(r string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("setRangeText", r)
	case 2:
		if start, ok := args[0].(uint); ok {
			if end, ok := args[1].(uint); ok {
				p.call("setRangeText", r, start, end)
			}
		}
	case 3:
		if start, ok := args[0].(uint); ok {
			if end, ok := args[1].(uint); ok {
				if selectionMode, ok := args[2].(SelectionMode); ok {
					p.call("setRangeText", r, start, end, string(selectionMode))
				}
			}
		}
	}
}

func (p *htmlTextAreaElementImpl) SetSelectionRange(start uint, end uint, direction ...string) {
	switch len(direction) {
	case 0:
		p.call("setSelectionRange", start, end)
	default:
		p.call("setSelectionRange", start, end, direction[0])
	}
}

// -------------8<---------------------------------------

type htmlOutputElementImpl struct {
	*htmlElementImpl
}

func NewHTMLOutputElement() HTMLOutputElement {
	if el := CurrentDocument().CreateElement("output"); el != nil {
		if output, ok := el.(HTMLOutputElement); ok {
			return output
		}
	}
	return nil
}

func wrapHTMLOutputElement(v Value) HTMLOutputElement {
	if v.valid() {
		return &htmlOutputElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlOutputElementImpl) HtmlFor() DOMTokenList {
	return wrapDOMTokenList(p.get("htmlFor"))
}

func (p *htmlOutputElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlOutputElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlOutputElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlOutputElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlOutputElementImpl) DefaultValue() string {
	return p.get("defaultValue").toString()
}

func (p *htmlOutputElementImpl) SetDefaultValue(dv string) {
	p.set("defaultValue", dv)
}

func (p *htmlOutputElementImpl) Value() string {
	return p.get("value").toString()
}

func (p *htmlOutputElementImpl) SetValue(v string) {
	p.set("value", v)
}

func (p *htmlOutputElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlOutputElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlOutputElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlOutputElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlOutputElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlOutputElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

func (p *htmlOutputElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

// -------------8<---------------------------------------

type htmlProgressElementImpl struct {
	*htmlElementImpl
}

func NewHTMLProgressElement() HTMLProgressElement {
	if el := CurrentDocument().CreateElement("progress"); el != nil {
		if progress, ok := el.(HTMLProgressElement); ok {
			return progress
		}
	}
	return nil
}

func wrapHTMLProgressElement(v Value) HTMLProgressElement {
	if v.valid() {
		return &htmlProgressElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlProgressElementImpl) Value() float64 {
	return p.get("value").toFloat64()
}

func (p *htmlProgressElementImpl) SetValue(v float64) {
	p.set("value", v)
}

func (p *htmlProgressElementImpl) Max() float64 {
	return p.get("max").toFloat64()
}

func (p *htmlProgressElementImpl) SetMax(m float64) {
	p.set("max", m)
}

func (p *htmlProgressElementImpl) Position() float64 {
	return p.get("position").toFloat64()
}

func (p *htmlProgressElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

// -------------8<---------------------------------------

type htmlMeterElementImpl struct {
	*htmlElementImpl
}

func NewHTMLMeterElement() HTMLMeterElement {
	if el := CurrentDocument().CreateElement("meter"); el != nil {
		if meter, ok := el.(HTMLMeterElement); ok {
			return meter
		}
	}
	return nil
}

func wrapHTMLMeterElement(v Value) HTMLMeterElement {
	if v.valid() {
		return &htmlMeterElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlMeterElementImpl) Value() float64 {
	return p.get("value").toFloat64()
}

func (p *htmlMeterElementImpl) SetValue(v float64) {
	p.set("value", v)
}

func (p *htmlMeterElementImpl) Min() float64 {
	return p.get("min").toFloat64()
}

func (p *htmlMeterElementImpl) SetMin(m float64) {
	p.set("min", m)
}

func (p *htmlMeterElementImpl) Max() float64 {
	return p.get("max").toFloat64()
}

func (p *htmlMeterElementImpl) SetMax(m float64) {
	p.set("max", m)
}

func (p *htmlMeterElementImpl) Low() float64 {
	return p.get("low").toFloat64()
}

func (p *htmlMeterElementImpl) SetLow(v float64) {
	p.set("low", v)
}

func (p *htmlMeterElementImpl) High() float64 {
	return p.get("high").toFloat64()
}

func (p *htmlMeterElementImpl) SetHigh(v float64) {
	p.set("high", v)
}

func (p *htmlMeterElementImpl) Optimum() float64 {
	return p.get("optimum").toFloat64()
}

func (p *htmlMeterElementImpl) SetOptimum(v float64) {
	p.set("optimum", v)
}

func (p *htmlMeterElementImpl) Labels() []Node {
	return nodeListToSlice(p.get("labels"))
}

// -------------8<---------------------------------------

type htmlFieldSetElementImpl struct {
	*htmlElementImpl
}

func NewHTMLFieldSetElement() HTMLFieldSetElement {
	if el := CurrentDocument().CreateElement("fieldset"); el != nil {
		if fieldset, ok := el.(HTMLFieldSetElement); ok {
			return fieldset
		}
	}
	return nil
}

func wrapHTMLFieldSetElement(v Value) HTMLFieldSetElement {
	if v.valid() {
		return &htmlFieldSetElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlFieldSetElementImpl) Disabled() bool {
	return p.get("disabled").toBool()
}

func (p *htmlFieldSetElementImpl) SetDisabled(b bool) {
	p.set("disabled", b)
}

func (p *htmlFieldSetElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

func (p *htmlFieldSetElementImpl) Name() string {
	return p.get("name").toString()
}

func (p *htmlFieldSetElementImpl) SetName(name string) {
	p.set("name", name)
}

func (p *htmlFieldSetElementImpl) Type() string {
	return p.get("type").toString()
}

func (p *htmlFieldSetElementImpl) Elements() []HTMLElement {
	return htmlCollectionToHTMLElementSlice(p.get("elements"))
}

func (p *htmlFieldSetElementImpl) WillValidate() bool {
	return p.get("willValidate").toBool()
}

func (p *htmlFieldSetElementImpl) Validity() ValidityState {
	return wrapValidityState(p.get("validity"))
}

func (p *htmlFieldSetElementImpl) ValidationMessage() string {
	return p.get("validationMessage").toString()
}

func (p *htmlFieldSetElementImpl) CheckValidity() bool {
	return p.call("checkValidity").toBool()
}

func (p *htmlFieldSetElementImpl) ReportValidity() bool {
	return p.call("reportValidity").toBool()
}

func (p *htmlFieldSetElementImpl) SetCustomValidity(e string) {
	p.call("setCustomValidity", e)
}

// -------------8<---------------------------------------

type htmlLegendElementImpl struct {
	*htmlElementImpl
}

func NewHTMLLegendElement() HTMLLegendElement {
	if el := CurrentDocument().CreateElement("legend"); el != nil {
		if legend, ok := el.(HTMLLegendElement); ok {
			return legend
		}
	}
	return nil
}

func wrapHTMLLegendElement(v Value) HTMLLegendElement {
	if v.valid() {
		return &htmlLegendElementImpl{
			htmlElementImpl: newHTMLElementImpl(v),
		}
	}
	return nil
}

func (p *htmlLegendElementImpl) Form() HTMLFormElement {
	return wrapHTMLFormElement(p.get("form"))
}

// -------------8<---------------------------------------
