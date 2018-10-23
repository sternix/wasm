// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

type htmlFormElementImpl struct {
	*htmlElementImpl
}

func newHTMLFormElement(v js.Value) HTMLFormElement {
	if isNil(v) {
		return nil
	}

	return &htmlFormElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlFormElementImpl) AcceptCharset() string {
	return p.Get("acceptCharset").String()
}

func (p *htmlFormElementImpl) SetAcceptCharset(ch string) {
	p.Set("acceptCharset", ch)
}

func (p *htmlFormElementImpl) Action() string {
	return p.Get("action").String()
}

func (p *htmlFormElementImpl) SetAction(a string) {
	p.Set("action", a)
}

func (p *htmlFormElementImpl) Autocomplete() string {
	return p.Get("autocomplete").String()
}

func (p *htmlFormElementImpl) SetAutocomplete(ac string) {
	p.Set("autocomplete", ac)
}

func (p *htmlFormElementImpl) Enctype() string {
	return p.Get("enctype").String()
}

func (p *htmlFormElementImpl) SetEnctype(et string) {
	p.Set("enctype", et)
}

func (p *htmlFormElementImpl) Encoding() string {
	return p.Get("encoding").String()
}

func (p *htmlFormElementImpl) SetEncoding(enc string) {
	p.Set("encoding", enc)
}

func (p *htmlFormElementImpl) Method() string {
	return p.Get("method").String()
}

func (p *htmlFormElementImpl) SetMethod(m string) {
	p.Set("method", m)
}

func (p *htmlFormElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlFormElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlFormElementImpl) NoValidate() bool {
	return p.Get("noValidate").Bool()
}

func (p *htmlFormElementImpl) SetNoValidate(b bool) {
	p.Set("noValidate", b)
}

func (p *htmlFormElementImpl) Target() string {
	return p.Get("target").String()
}

func (p *htmlFormElementImpl) SetTarget(t string) {
	p.Set("target", t)
}

func (p *htmlFormElementImpl) Elements() HTMLFormControlsCollection {
	return newHTMLFormControlsCollection(p.Get("elements"))
}

func (p *htmlFormElementImpl) Submit() {
	p.Call("submit")
}

func (p *htmlFormElementImpl) Reset() {
	p.Call("reset")
}

func (p *htmlFormElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlFormElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

// -------------8<---------------------------------------

type htmlFormControlsCollectionImpl struct {
	*htmlCollectionImpl
}

func newHTMLFormControlsCollection(v js.Value) HTMLFormControlsCollection {
	if isNil(v) {
		return nil
	}
	return &htmlFormControlsCollectionImpl{
		htmlCollectionImpl: newHTMLCollectionImpl(v),
	}
}

//NOTE overriden namedbyItem
func (p *htmlFormControlsCollectionImpl) ItemByName(item string) HTMLFormControl {
	return newHTMLFormControl(p.Call("namedItem", item))
}

// -------------8<---------------------------------------

type htmlFormControlImpl struct {
	js.Value
}

func newHTMLFormControl(v js.Value) HTMLFormControl {
	if isNil(v) {
		return nil
	}

	return &htmlFormControlImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type radioNodeListImpl struct {
	*nodeListImpl
}

func newRadioNodeList(v js.Value) RadioNodeList {
	if isNil(v) {
		return nil
	}

	return &radioNodeListImpl{
		nodeListImpl: newNodeListImpl(v),
	}
}

func (p *radioNodeListImpl) Value() string {
	return p.Get("value").String()
}

func (p *radioNodeListImpl) SetValue(value string) {
	p.Set("value", value)
}

// -------------8<---------------------------------------

type htmlLabelElementImpl struct {
	*htmlElementImpl
}

func newHTMLLabelElement(v js.Value) HTMLLabelElement {
	if isNil(v) {
		return nil
	}

	return &htmlLabelElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlLabelElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlLabelElementImpl) HtmlFor() string {
	return p.Get("htmlFor").String()
}

func (p *htmlLabelElementImpl) SetHtmlFor(hf string) {
	p.Set("htmlFor", hf)
}

func (p *htmlLabelElementImpl) Control() HTMLElement {
	return newHTMLElement(p.Get("control"))
}

// -------------8<---------------------------------------

type htmlInputElementImpl struct {
	*htmlElementImpl
}

func newHTMLInputElement(v js.Value) HTMLInputElement {
	if isNil(v) {
		return nil
	}

	return &htmlInputElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlInputElementImpl) Accept() string {
	return p.Get("accept").String()
}

func (p *htmlInputElementImpl) SetAccept(a string) {
	p.Set("accept", a)
}

func (p *htmlInputElementImpl) Alt() string {
	return p.Get("alt").String()
}

func (p *htmlInputElementImpl) SetAlt(a string) {
	p.Set("alt", a)
}

func (p *htmlInputElementImpl) Autocomplete() string {
	return p.Get("autocomplete").String()
}

func (p *htmlInputElementImpl) SetAutocomplete(ac string) {
	p.Set("autocomplete", ac)
}

func (p *htmlInputElementImpl) Autofocus() bool {
	return p.Get("autofocus").Bool()
}

func (p *htmlInputElementImpl) SetAutofocus(af bool) {
	p.Set("autofocus", af)
}

func (p *htmlInputElementImpl) DefaultChecked() bool {
	return p.Get("defaultChecked").Bool()
}

func (p *htmlInputElementImpl) SetDefaultChecked(dc bool) {
	p.Set("defaultChecked", dc)
}

func (p *htmlInputElementImpl) Checked() bool {
	return p.Get("checked").Bool()
}

func (p *htmlInputElementImpl) SetChecked(c bool) {
	p.Set("checked", c)
}

func (p *htmlInputElementImpl) DirName() string {
	return p.Get("dirName").String()
}

func (p *htmlInputElementImpl) SetDirName(dn string) {
	p.Set("dirName", dn)
}

func (p *htmlInputElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlInputElementImpl) SetDisabled(d bool) {
	p.Set("disabled", d)
}

func (p *htmlInputElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlInputElementImpl) Files() []File {
	return fileListToSlice(p.Get("files"))
}

func (p *htmlInputElementImpl) FormAction() string {
	return p.Get("formAction").String()
}

func (p *htmlInputElementImpl) SetFormAction(fa string) {
	p.Set("formAction", fa)
}

func (p *htmlInputElementImpl) FormEnctype() string {
	return p.Get("formEnctype").String()
}

func (p *htmlInputElementImpl) SetFormEnctype(fe string) {
	p.Set("formEnctype", fe)
}

func (p *htmlInputElementImpl) FormMethod() string {
	return p.Get("formMethod").String()
}

func (p *htmlInputElementImpl) SetFormMethod(fm string) {
	p.Set("formMethod", fm)
}

func (p *htmlInputElementImpl) FormNoValidate() bool {
	return p.Get("formNoValidate").Bool()
}

func (p *htmlInputElementImpl) SetFormNoValidate(b bool) {
	p.Set("formNoValidate", b)
}

func (p *htmlInputElementImpl) FormTarget() string {
	return p.Get("formTarget").String()
}

func (p *htmlInputElementImpl) SetFormTarget(ft string) {
	p.Set("formTarget", ft)
}

func (p *htmlInputElementImpl) Height() int {
	return p.Get("height").Int()
}

func (p *htmlInputElementImpl) SetHeight(h int) {
	p.Set("height", h)
}

func (p *htmlInputElementImpl) Indeterminate() bool {
	return p.Get("indeterminate").Bool()
}

func (p *htmlInputElementImpl) SetIndeterminate(b bool) {
	p.Set("indeterminate", b)
}

func (p *htmlInputElementImpl) List() HTMLElement {
	return newHTMLElement(p.Get("list"))
}

func (p *htmlInputElementImpl) Max() string {
	return p.Get("max").String()
}

func (p *htmlInputElementImpl) SetMax(m string) {
	p.Set("max", m)
}

func (p *htmlInputElementImpl) MaxLength() int {
	return p.Get("maxLength").Int()
}

func (p *htmlInputElementImpl) SetMaxLength(m int) {
	p.Set("maxLength", m)
}

func (p *htmlInputElementImpl) Min() string {
	return p.Get("min").String()
}

func (p *htmlInputElementImpl) SetMin(m string) {
	p.Set("min", m)
}

func (p *htmlInputElementImpl) MinLength() int {
	return p.Get("minLength").Int()
}

func (p *htmlInputElementImpl) SetMinLength(m int) {
	p.Set("minLength", m)
}

func (p *htmlInputElementImpl) Multiple() bool {
	return p.Get("multiple").Bool()
}

func (p *htmlInputElementImpl) SetMultiple(b bool) {
	p.Set("multiple", b)
}

func (p *htmlInputElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlInputElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlInputElementImpl) Pattern() string {
	return p.Get("pattern").String()
}

func (p *htmlInputElementImpl) SetPattern(pattern string) {
	p.Set("pattern", pattern)
}

func (p *htmlInputElementImpl) Placeholder() string {
	return p.Get("placeholder").String()
}

func (p *htmlInputElementImpl) SetPlaceholder(ph string) {
	p.Set("placeholder", ph)
}

func (p *htmlInputElementImpl) ReadOnly() bool {
	return p.Get("readOnly").Bool()
}

func (p *htmlInputElementImpl) SetReadOnly(b bool) {
	p.Set("readOnly", b)
}

func (p *htmlInputElementImpl) Required() bool {
	return p.Get("_required").Bool()
}

func (p *htmlInputElementImpl) SetRequired(b bool) {
	p.Set("_required", b)
}

func (p *htmlInputElementImpl) Size() int {
	return p.Get("size").Int()
}

func (p *htmlInputElementImpl) SetSize(s int) {
	p.Set("size", s)
}

func (p *htmlInputElementImpl) Src() string {
	return p.Get("src").String()
}

func (p *htmlInputElementImpl) SetSrc(src string) {
	p.Set("src", src)
}

func (p *htmlInputElementImpl) Step() string {
	return p.Get("step").String()
}

func (p *htmlInputElementImpl) SetStep(s string) {
	p.Set("step", s)
}

func (p *htmlInputElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlInputElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlInputElementImpl) DefaultValue() string {
	return p.Get("defaultValue").String()
}

func (p *htmlInputElementImpl) SetDefaultValue(dv string) {
	p.Set("defaultValue", dv)
}

func (p *htmlInputElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlInputElementImpl) SetValue(value string) {
	p.Set("value", value)
}

func (p *htmlInputElementImpl) ValueAsDate() time.Time {
	//TODO: test it
	return jsDateToTime(p.Get("valueAsDate"))
}

func (p *htmlInputElementImpl) SetValueAsDate(t time.Time) {
	// TODO: test it
	d := jsDate.New()
	d.Call("setTime", t.Unix())
	p.Set("valueAsDate", d)
}

func (p *htmlInputElementImpl) ValueAsNumber() float64 {
	return p.Get("valueAsNumber").Float()
}

func (p *htmlInputElementImpl) SetValueAsNumber(n float64) {
	p.Set("valueAsNumber", n)
}

func (p *htmlInputElementImpl) Width() int {
	return p.Get("width").Int()
}

func (p *htmlInputElementImpl) SetWidth(w int) {
	p.Set("width", w)
}

func (p *htmlInputElementImpl) StepUp(n ...int) {
	switch len(n) {
	case 0:
		p.Call("stepUp")
	default:
		p.Call("stepUp", n[0])
	}
}

func (p *htmlInputElementImpl) StepDown(n ...int) {
	switch len(n) {
	case 0:
		p.Call("stepDown")
	default:
		p.Call("stepDown", n[0])
	}
}

func (p *htmlInputElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlInputElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlInputElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlInputElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlInputElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlInputElementImpl) SetCustomValidity(cv string) {
	p.Call("setCustomValidity", cv)
}

func (p *htmlInputElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

func (p *htmlInputElementImpl) Select() {
	p.Call("select")
}

func (p *htmlInputElementImpl) SelectionStart() int {
	return p.Get("selectionStart").Int()
}

func (p *htmlInputElementImpl) SetSelectionStart(se int) {
	p.Set("selectionStart", se)
}

func (p *htmlInputElementImpl) SelectionEnd() int {
	return p.Get("selectionEnd").Int()
}

func (p *htmlInputElementImpl) SetSelectionEnd(se int) {
	p.Set("selectionEnd", se)
}

func (p *htmlInputElementImpl) SelectionDirection() string {
	return p.Get("selectionDirection").String()
}

func (p *htmlInputElementImpl) SetSelectionDirection(sd string) {
	p.Set("selectionDirection", sd)
}

func (p *htmlInputElementImpl) SetRangeText(r string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("setRangeText", r)
	case 2:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				p.Call("setRangeText", r, start, end)
			}
		}
	case 3:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				if selectionMode, ok := args[2].(SelectionMode); ok {
					p.Call("setRangeText", r, start, end, string(selectionMode))
				}
			}
		}
	}
}

func (p *htmlInputElementImpl) SetSelectionRange(start int, end int, direction ...string) {
	switch len(direction) {
	case 0:
		p.Call("setSelectionRange", start, end)
	default:
		p.Call("setSelectionRange", start, end, direction[0])

	}
}

// -------------8<---------------------------------------

type htmlButtonElementImpl struct {
	*htmlElementImpl
}

func newHTMLButtonElement(v js.Value) HTMLButtonElement {
	if isNil(v) {
		return nil
	}

	return &htmlButtonElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlButtonElementImpl) Autofocus() bool {
	return p.Get("autofocus").Bool()
}

func (p *htmlButtonElementImpl) SetAutofocus(b bool) {
	p.Set("autofocus", b)
}

func (p *htmlButtonElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlButtonElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlButtonElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlButtonElementImpl) FormAction() string {
	return p.Get("formAction").String()
}

func (p *htmlButtonElementImpl) SetFormAction(fa string) {
	p.Set("formAction", fa)
}

func (p *htmlButtonElementImpl) FormEnctype() string {
	return p.Get("formEnctype").String()
}

func (p *htmlButtonElementImpl) SetFormEnctype(fe string) {
	p.Set("formEnctype", fe)
}

func (p *htmlButtonElementImpl) FormMethod() string {
	return p.Get("formMethod").String()
}

func (p *htmlButtonElementImpl) SetFormMethod(fm string) {
	p.Set("formMethod", fm)
}

func (p *htmlButtonElementImpl) FormNoValidate() bool {
	return p.Get("formNoValidate").Bool()
}

func (p *htmlButtonElementImpl) SetFormNoValidate(b bool) {
	p.Set("formNoValidate", b)
}

func (p *htmlButtonElementImpl) FormTarget() string {
	return p.Get("formTarget").String()
}

func (p *htmlButtonElementImpl) SetFormTarget(ft string) {
	p.Set("formTarget", ft)
}

func (p *htmlButtonElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlButtonElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlButtonElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlButtonElementImpl) SetType(t string) {
	p.Set("type", t)
}

func (p *htmlButtonElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlButtonElementImpl) SetValue(v string) {
	p.Set("value", v)
}

func (p *htmlButtonElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlButtonElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlButtonElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlButtonElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlButtonElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlButtonElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

func (p *htmlButtonElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

// -------------8<---------------------------------------

type htmlSelectElementImpl struct {
	*htmlElementImpl
}

func newHTMLSelectElement(v js.Value) HTMLSelectElement {
	if isNil(v) {
		return nil
	}

	return &htmlSelectElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlSelectElementImpl) Autocomplete() string {
	return p.Get("autocomplete").String()
}

func (p *htmlSelectElementImpl) SetAutocomplete(ac string) {
	p.Set("autocomplete", ac)
}

func (p *htmlSelectElementImpl) Autofocus() bool {
	return p.Get("autofocus").Bool()
}

func (p *htmlSelectElementImpl) SetAutofocus(af bool) {
	p.Set("autofocus", af)
}

func (p *htmlSelectElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlSelectElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlSelectElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlSelectElementImpl) Multiple() bool {
	return p.Get("multiple").Bool()
}

func (p *htmlSelectElementImpl) SetMultiple(m bool) {
	p.Set("multiple", m)
}

func (p *htmlSelectElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlSelectElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlSelectElementImpl) Required() bool {
	return p.Get("_required").Bool()
}

func (p *htmlSelectElementImpl) SetRequired(b bool) {
	p.Set("_required", b)
}

func (p *htmlSelectElementImpl) Size() int {
	return p.Get("size").Int()
}

func (p *htmlSelectElementImpl) SetSize(s int) {
	p.Set("size", s)
}

func (p *htmlSelectElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlSelectElementImpl) Options() HTMLOptionsCollection {
	return newHTMLOptionsCollection(p.Get("options"))
}

func (p *htmlSelectElementImpl) Length() int {
	return p.Get("length").Int()
}

func (p *htmlSelectElementImpl) SetLength(l int) {
	p.Set("length", l)
}

func (p *htmlSelectElementImpl) Item(index int) Element {
	return newElement(p.Call("item", index))
}

func (p *htmlSelectElementImpl) NamedItem(name string) HTMLOptionElement {
	return newHTMLOptionElement(p.Call("namedItem", name))
}

func (p *htmlSelectElementImpl) Add(element HTMLElement, before ...interface{}) {
	switch len(before) {
	case 0:
		p.Call("add", element.JSValue())
	default:
		switch x := before[0].(type) {
		case HTMLElement:
			p.Call("add", element.JSValue(), x.JSValue())
		case int:
			p.Call("add", element.JSValue(), x)
		}
	}
}

func (p *htmlSelectElementImpl) RemoveByIndex(index int) {
	p.Call("remove", index)
}

func (p *htmlSelectElementImpl) SetByIndex(index int, option HTMLOptionElement) {
	// TODO
	// which method ??
	//  setter void (unsigned long index, HTMLOptionElement? option);
}

func (p *htmlSelectElementImpl) SelectedOptions() HTMLCollection {
	return newHTMLCollection(p.Get("selectedOptions"))
}

func (p *htmlSelectElementImpl) SelectedIndex() int {
	return p.Get("selectedIndex").Int()
}

func (p *htmlSelectElementImpl) SetSelectedIndex(index int) {
	p.Set("selectedIndex", index)
}

func (p *htmlSelectElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlSelectElementImpl) SetValue(v string) {
	p.Set("value", v)
}

func (p *htmlSelectElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlSelectElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlSelectElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlSelectElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlSelectElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlSelectElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

func (p *htmlSelectElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

// -------------8<---------------------------------------

type htmlOptionsCollectionImpl struct {
	*htmlCollectionImpl
}

func newHTMLOptionsCollection(v js.Value) HTMLOptionsCollection {
	if isNil(v) {
		return nil
	}

	return &htmlOptionsCollectionImpl{
		htmlCollectionImpl: newHTMLCollectionImpl(v),
	}
}

func (p *htmlOptionsCollectionImpl) Length() int {
	return p.Get("length").Int()
}

/* TODO setter
func (p *htmlOptionsCollectionImpl) Set(index int,option HTMLOptionElement) {
	p.Call("")
}
*/

func (p *htmlOptionsCollectionImpl) Add(element HTMLElement, before ...interface{}) {
	switch len(before) {
	case 0:
		p.Call("add", element.JSValue())
	default:
		switch x := before[0].(type) {
		case HTMLElement:
			p.Call("add", element.JSValue(), x.JSValue())
		case int:
			p.Call("add", element.JSValue(), x)
		}
	}
}

func (p *htmlOptionsCollectionImpl) Remove(index int) {
	p.Call("remove", index)
}

func (p *htmlOptionsCollectionImpl) SelectedIndex() int {
	return p.Get("selectedIndex").Int()
}

func (p *htmlOptionsCollectionImpl) SetSelectedIndex(i int) {
	p.Set("selectedIndex", i)
}

// -------------8<---------------------------------------

type htmlDataListElementImpl struct {
	*htmlElementImpl
}

func newHTMLDataListElement(v js.Value) HTMLDataListElement {
	if isNil(v) {
		return nil
	}

	return &htmlDataListElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlDataListElementImpl) Options() HTMLCollection {
	return newHTMLCollection(p.Get("options"))
}

// -------------8<---------------------------------------

type htmlOptGroupElementImpl struct {
	*htmlElementImpl
}

func newHTMLOptGroupElement(v js.Value) HTMLOptGroupElement {
	if isNil(v) {
		return nil
	}

	return &htmlOptGroupElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlOptGroupElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlOptGroupElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlOptGroupElementImpl) Label() string {
	return p.Get("label").String()
}

func (p *htmlOptGroupElementImpl) SetLabel(lbl string) {
	p.Set("label", lbl)
}

// -------------8<---------------------------------------

type htmlOptionElementImpl struct {
	*htmlElementImpl
}

func newHTMLOptionElement(v js.Value) HTMLOptionElement {
	if isNil(v) {
		return nil
	}

	return &htmlOptionElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlOptionElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlOptionElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlOptionElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlOptionElementImpl) Label() string {
	return p.Get("label").String()
}

func (p *htmlOptionElementImpl) SetLabel(lbl string) {
	p.Set("label", lbl)
}

func (p *htmlOptionElementImpl) DefaultSelected() bool {
	return p.Get("defaultSelected").Bool()
}

func (p *htmlOptionElementImpl) SetDefaultSelected(b bool) {
	p.Set("defaultSelected", b)
}

func (p *htmlOptionElementImpl) Selected() bool {
	return p.Get("selected").Bool()
}

func (p *htmlOptionElementImpl) SetSelected(b bool) {
	p.Set("selected", b)
}

func (p *htmlOptionElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlOptionElementImpl) SetValue(v string) {
	p.Set("value", v)
}

func (p *htmlOptionElementImpl) Text() string {
	return p.Get("text").String()
}

func (p *htmlOptionElementImpl) SetText(t string) {
	p.Set("text", t)
}

func (p *htmlOptionElementImpl) Index() int {
	return p.Get("index").Int()
}

// -------------8<---------------------------------------

type htmlTextAreaElementImpl struct {
	*htmlElementImpl
}

func newHTMLTextAreaElement(v js.Value) HTMLTextAreaElement {
	if isNil(v) {
		return nil
	}

	return &htmlTextAreaElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlTextAreaElementImpl) Autocomplete() string {
	return p.Get("autocomplete").String()
}

func (p *htmlTextAreaElementImpl) SetAutocomplete(ac string) {
	p.Set("autocomplete", ac)
}

func (p *htmlTextAreaElementImpl) Autofocus() bool {
	return p.Get("autofocus").Bool()
}

func (p *htmlTextAreaElementImpl) SetAutofocus(b bool) {
	p.Set("autofocus", b)
}

func (p *htmlTextAreaElementImpl) Cols() int {
	return p.Get("cols").Int()
}

func (p *htmlTextAreaElementImpl) SetCols(c int) {
	p.Set("cols", c)
}

func (p *htmlTextAreaElementImpl) DirName() string {
	return p.Get("dirName").String()
}

func (p *htmlTextAreaElementImpl) SetDirName(dn string) {
	p.Set("dirName", dn)
}

func (p *htmlTextAreaElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlTextAreaElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlTextAreaElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlTextAreaElementImpl) MaxLength() int {
	return p.Get("maxLength").Int()
}

func (p *htmlTextAreaElementImpl) SetMaxLength(m int) {
	p.Set("maxLength", m)
}

func (p *htmlTextAreaElementImpl) MinLength() int {
	return p.Get("minLength").Int()
}

func (p *htmlTextAreaElementImpl) SetMinLength(m int) {
	p.Set("minLength", m)
}

func (p *htmlTextAreaElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlTextAreaElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlTextAreaElementImpl) Placeholder() string {
	return p.Get("placeholder").String()
}

func (p *htmlTextAreaElementImpl) SetPlaceholder(h string) {
	p.Set("placeholder", h)
}

func (p *htmlTextAreaElementImpl) ReadOnly() bool {
	return p.Get("readOnly").Bool()
}

func (p *htmlTextAreaElementImpl) SetReadOnly(b bool) {
	p.Set("readOnly", b)
}

func (p *htmlTextAreaElementImpl) Required() bool {
	return p.Get("_required").Bool()
}

func (p *htmlTextAreaElementImpl) SetRequired(b bool) {
	p.Set("_required", b)
}

func (p *htmlTextAreaElementImpl) Rows() int {
	return p.Get("rows").Int()
}

func (p *htmlTextAreaElementImpl) SetRows(r int) {
	p.Set("rows", r)
}

func (p *htmlTextAreaElementImpl) Wrap() string {
	return p.Get("wrap").String()
}

func (p *htmlTextAreaElementImpl) SetWrap(w string) {
	p.Set("wrap", w)
}

func (p *htmlTextAreaElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlTextAreaElementImpl) DefaultValue() string {
	return p.Get("defaultValue").String()
}

func (p *htmlTextAreaElementImpl) SetDefaultValue(dv string) {
	p.Set("defaultValue", dv)
}

func (p *htmlTextAreaElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlTextAreaElementImpl) SetValue(v string) {
	p.Set("value", v)
}

func (p *htmlTextAreaElementImpl) TextLength() int {
	return p.Get("textLength").Int()
}

func (p *htmlTextAreaElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlTextAreaElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlTextAreaElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlTextAreaElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlTextAreaElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlTextAreaElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

func (p *htmlTextAreaElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

func (p *htmlTextAreaElementImpl) Select() {
	p.Call("select")
}

func (p *htmlTextAreaElementImpl) SelectionStart() int {
	return p.Get("selectionStart").Int()
}

func (p *htmlTextAreaElementImpl) SetSelectionStart(ss int) {
	p.Set("selectionStart", ss)
}

func (p *htmlTextAreaElementImpl) SelectionEnd() int {
	return p.Get("selectionEnd").Int()
}

func (p *htmlTextAreaElementImpl) SetSelectionEnd(se int) {
	p.Set("selectionEnd", se)
}

func (p *htmlTextAreaElementImpl) SelectionDirection() string {
	return p.Get("selectionDirection").String()
}

func (p *htmlTextAreaElementImpl) SetSelectionDirection(sd string) {
	p.Set("selectionDirection", sd)
}

func (p *htmlTextAreaElementImpl) SetRangeText(r string, args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("setRangeText", r)
	case 2:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				p.Call("setRangeText", r, start, end)
			}
		}
	case 3:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				if selectionMode, ok := args[2].(SelectionMode); ok {
					p.Call("setRangeText", r, start, end, string(selectionMode))
				}
			}
		}
	}
}

func (p *htmlTextAreaElementImpl) SetSelectionRange(start int, end int, direction ...string) {
	switch len(direction) {
	case 0:
		p.Call("setSelectionRange", start, end)
	default:
		p.Call("setSelectionRange", start, end, direction[0])
	}
}

// -------------8<---------------------------------------

type htmlOutputElementImpl struct {
	*htmlElementImpl
}

func newHTMLOutputElement(v js.Value) HTMLOutputElement {
	if isNil(v) {
		return nil
	}

	return &htmlOutputElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlOutputElementImpl) HtmlFor() DOMTokenList {
	return newDOMTokenList(p.Get("htmlFor"))
}

func (p *htmlOutputElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlOutputElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlOutputElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlOutputElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlOutputElementImpl) DefaultValue() string {
	return p.Get("defaultValue").String()
}

func (p *htmlOutputElementImpl) SetDefaultValue(dv string) {
	p.Set("defaultValue", dv)
}

func (p *htmlOutputElementImpl) Value() string {
	return p.Get("value").String()
}

func (p *htmlOutputElementImpl) SetValue(v string) {
	p.Set("value", v)
}

func (p *htmlOutputElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlOutputElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlOutputElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlOutputElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlOutputElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlOutputElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

func (p *htmlOutputElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

// -------------8<---------------------------------------

type htmlProgressElementImpl struct {
	*htmlElementImpl
}

func newHTMLProgressElement(v js.Value) HTMLProgressElement {
	if isNil(v) {
		return nil
	}

	return &htmlProgressElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlProgressElementImpl) Value() float64 {
	return p.Get("value").Float()
}

func (p *htmlProgressElementImpl) SetValue(v float64) {
	p.Set("value", v)
}

func (p *htmlProgressElementImpl) Max() float64 {
	return p.Get("max").Float()
}

func (p *htmlProgressElementImpl) SetMax(m float64) {
	p.Set("max", m)
}

func (p *htmlProgressElementImpl) Position() float64 {
	return p.Get("position").Float()
}

func (p *htmlProgressElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

// -------------8<---------------------------------------

type htmlMeterElementImpl struct {
	*htmlElementImpl
}

func newHTMLMeterElement(v js.Value) HTMLMeterElement {
	if isNil(v) {
		return nil
	}

	return &htmlMeterElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlMeterElementImpl) Value() float64 {
	return p.Get("value").Float()
}

func (p *htmlMeterElementImpl) SetValue(v float64) {
	p.Set("value", v)
}

func (p *htmlMeterElementImpl) Min() float64 {
	return p.Get("min").Float()
}

func (p *htmlMeterElementImpl) SetMin(m float64) {
	p.Set("min", m)
}

func (p *htmlMeterElementImpl) Max() float64 {
	return p.Get("max").Float()
}

func (p *htmlMeterElementImpl) SetMax(m float64) {
	p.Set("max", m)
}

func (p *htmlMeterElementImpl) Low() float64 {
	return p.Get("low").Float()
}

func (p *htmlMeterElementImpl) SetLow(v float64) {
	p.Set("low", v)
}

func (p *htmlMeterElementImpl) High() float64 {
	return p.Get("high").Float()
}

func (p *htmlMeterElementImpl) SetHigh(v float64) {
	p.Set("high", v)
}

func (p *htmlMeterElementImpl) Optimum() float64 {
	return p.Get("optimum").Float()
}

func (p *htmlMeterElementImpl) SetOptimum(v float64) {
	p.Set("optimum", v)
}

func (p *htmlMeterElementImpl) Labels() []Node {
	return nodeListToSlice(p.Get("labels"))
}

// -------------8<---------------------------------------

type htmlFieldSetElementImpl struct {
	*htmlElementImpl
}

func newHTMLFieldSetElement(v js.Value) HTMLFieldSetElement {
	if isNil(v) {
		return nil
	}

	return &htmlFieldSetElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlFieldSetElementImpl) Disabled() bool {
	return p.Get("disabled").Bool()
}

func (p *htmlFieldSetElementImpl) SetDisabled(b bool) {
	p.Set("disabled", b)
}

func (p *htmlFieldSetElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

func (p *htmlFieldSetElementImpl) Name() string {
	return p.Get("name").String()
}

func (p *htmlFieldSetElementImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *htmlFieldSetElementImpl) Type() string {
	return p.Get("type").String()
}

func (p *htmlFieldSetElementImpl) Elements() HTMLCollection {
	return newHTMLCollection(p.Get("elements"))
}

func (p *htmlFieldSetElementImpl) WillValidate() bool {
	return p.Get("willValidate").Bool()
}

func (p *htmlFieldSetElementImpl) Validity() ValidityState {
	return newValidityState(p.Get("validity"))
}

func (p *htmlFieldSetElementImpl) ValidationMessage() string {
	return p.Get("validationMessage").String()
}

func (p *htmlFieldSetElementImpl) CheckValidity() bool {
	return p.Call("checkValidity").Bool()
}

func (p *htmlFieldSetElementImpl) ReportValidity() bool {
	return p.Call("reportValidity").Bool()
}

func (p *htmlFieldSetElementImpl) SetCustomValidity(e string) {
	p.Call("setCustomValidity", e)
}

// -------------8<---------------------------------------

type htmlLegendElementImpl struct {
	*htmlElementImpl
}

func newHTMLLegendElement(v js.Value) HTMLLegendElement {
	if isNil(v) {
		return nil
	}

	return &htmlLegendElementImpl{
		htmlElementImpl: newHTMLElementImpl(v),
	}
}

func (p *htmlLegendElementImpl) Form() HTMLFormElement {
	return newHTMLFormElement(p.Get("form"))
}

// -------------8<---------------------------------------
