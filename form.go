// +build js,wasm

package wasm

import (
	"time"
)

type (
	// https://www.w3.org/TR/html52/sec-forms.html#htmlformelement
	HTMLFormElement interface {
		HTMLElement

		AcceptCharset() string
		SetAcceptCharset(string)
		Action() string
		SetAction(string)
		Autocomplete() string
		SetAutocomplete(string)
		Enctype() string
		SetEnctype(string)
		Encoding() string
		SetEncoding(string)
		Method() string
		SetMethod(string)
		Name() string
		SetName(string)
		NoValidate() bool
		SetNoValidate(bool)
		Target() string
		SetTarget(string)
		Elements() HTMLFormControlsCollection
		Submit()
		Reset()
		CheckValidity() bool
		ReportValidity() bool
	}

	// https://www.w3.org/TR/html52/infrastructure.html#htmlformcontrolscollection
	HTMLFormControlsCollection interface {
		HTMLCollection
		ItemByName(string) HTMLFormControl
	}

	// Element or RadioNodeList
	HTMLFormControl interface{}

	// https://www.w3.org/TR/html52/infrastructure.html#radionodelist
	RadioNodeList interface {
		NodeList

		Value() string
		SetValue(string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmllabelelement
	HTMLLabelElement interface {
		HTMLElement

		Form() HTMLFormElement
		HtmlFor() string
		SetHtmlFor(string)
		Control() HTMLElement
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlinputelement
	HTMLInputElement interface {
		HTMLElement

		Accept() string
		SetAccept(string)
		Alt() string
		SetAlt(string)
		Autocomplete() string
		SetAutocomplete(string)
		Autofocus() bool
		SetAutofocus(bool)
		DefaultChecked() bool
		SetDefaultChecked(bool)
		Checked() bool
		SetChecked(bool)
		DirName() string
		SetDirName(string)
		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		Files() []File
		FormAction() string
		SetFormAction(string)
		FormEnctype() string
		SetFormEnctype(string)
		FormMethod() string
		SetFormMethod(string)
		FormNoValidate() bool
		SetFormNoValidate(bool)
		FormTarget() string
		SetFormTarget(string)
		Height() int
		SetHeight(int)
		Indeterminate() bool
		SetIndeterminate(bool)
		List() HTMLElement
		Max() string
		SetMax(string)
		MaxLength() int
		SetMaxLength(int)
		Min() string
		SetMin(string)
		MinLength() int
		SetMinLength(int)
		Multiple() bool
		SetMultiple(bool)
		Name() string
		SetName(string)
		Pattern() string
		SetPattern(string)
		Placeholder() string
		SetPlaceholder(string)
		ReadOnly() bool
		SetReadOnly(bool)
		Required() bool
		SetRequired(bool)
		Size() int
		SetSize(int)
		Src() string
		SetSrc(string)
		Step() string
		SetStep(string)
		Type() string
		SetType(string)
		DefaultValue() string
		SetDefaultValue(string)
		Value() string
		SetValue(string)
		ValueAsDate() time.Time
		SetValueAsDate(time.Time)
		ValueAsNumber() float64
		SetValueAsNumber(float64)
		Width() int
		SetWidth(int)
		StepUp(...int)
		StepDown(...int)
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
		Labels() []Node
		Select()
		SelectionStart() int
		SetSelectionStart(int)
		SelectionEnd() int
		SetSelectionEnd(int)
		SelectionDirection() string
		SetSelectionDirection(string)
		SetRangeText(string, ...interface{})
		SetSelectionRange(int, int, ...string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlbuttonelement
	HTMLButtonElement interface {
		HTMLElement

		Autofocus() bool
		SetAutofocus(bool)
		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		FormAction() string
		SetFormAction(string)
		FormEnctype() string
		SetFormEnctype(string)
		FormMethod() string
		SetFormMethod(string)
		FormNoValidate() bool
		SetFormNoValidate(bool)
		FormTarget() string
		SetFormTarget(string)
		Name() string
		SetName(string)
		Type() string
		SetType(string)
		Value() string
		SetValue(string)
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
		Labels() []Node
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlselectelement
	HTMLSelectElement interface {
		HTMLElement

		Autocomplete() string
		SetAutocomplete(string)
		Autofocus() bool
		SetAutofocus(bool)
		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		Multiple() bool
		SetMultiple(bool)
		Name() string
		SetName(string)
		Required() bool
		SetRequired(bool)
		Size() int
		SetSize(int)
		Type() string
		Options() HTMLOptionsCollection
		Length() int
		SetLength(int)
		Item(int) Element
		NamedItem(string) HTMLOptionElement
		Add(HTMLElement, ...interface{})
		// NOTE: Overload ChildNode.Remove() method
		RemoveByIndex(int) // remove(long)
		SetByIndex(int, HTMLOptionElement)
		SelectedOptions() HTMLCollection
		SelectedIndex() int
		SetSelectedIndex(int)
		Value() string
		SetValue(string)
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string

		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)

		Labels() []Node
	}

	// https://www.w3.org/TR/html52/infrastructure.html#htmloptionscollection
	HTMLOptionsCollection interface {
		HTMLCollection

		//Length() int
		//Set(int, HTMLOptionElement)
		Add(HTMLElement, ...interface{})
		Remove(int)
		SelectedIndex() int
		SetSelectedIndex(int)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmldatalistelement
	HTMLDataListElement interface {
		HTMLElement

		Options() HTMLCollection
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmloptgroupelement
	HTMLOptGroupElement interface {
		HTMLElement

		Disabled() bool
		SetDisabled(bool)
		Label() string
		SetLabel(string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmloptionelement
	HTMLOptionElement interface {
		HTMLElement

		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		Label() string
		SetLabel(string)
		DefaultSelected() bool
		SetDefaultSelected(bool)
		Selected() bool
		SetSelected(bool)
		Value() string
		SetValue(string)
		Text() string
		SetText(string)
		Index() int
	}

	// for NewOption
	OptionParams struct {
		Text            string
		Value           string
		DefaultSelected bool
		Selected        bool
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmltextareaelement
	HTMLTextAreaElement interface {
		HTMLElement

		Autocomplete() string
		SetAutocomplete(string)
		Autofocus() bool
		SetAutofocus(bool)
		Cols() int
		SetCols(int)
		DirName() string
		SetDirName(string)
		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		MaxLength() int
		SetMaxLength(int)
		MinLength() int
		SetMinLength(int)
		Name() string
		SetName(string)
		Placeholder() string
		SetPlaceholder(string)
		ReadOnly() bool
		SetReadOnly(bool)
		Required() bool
		SetRequired(bool)
		Rows() int
		SetRows(int)
		Wrap() string
		SetWrap(string)
		Type() string
		DefaultValue() string
		SetDefaultValue(string)
		Value() string
		SetValue(string)
		TextLength() int
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
		Labels() []Node
		Select()
		SelectionStart() int
		SetSelectionStart(int)
		SelectionEnd() int
		SetSelectionEnd(int)
		SelectionDirection() string
		SetSelectionDirection(string)
		SetRangeText(string, ...interface{})
		SetSelectionRange(int, int, ...string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmloutputelement
	HTMLOutputElement interface {
		HTMLElement

		HtmlFor() DOMTokenList
		Form() HTMLFormElement
		Name() string
		SetName(string)
		Type() string
		DefaultValue() string
		SetDefaultValue(string)
		Value() string
		SetValue(string)
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
		Labels() []Node
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlprogresselement
	HTMLProgressElement interface {
		HTMLElement

		Value() float64
		SetValue(float64)
		Max() float64
		SetMax(float64)
		Position() float64
		Labels() []Node
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlmeterelement
	HTMLMeterElement interface {
		HTMLElement

		Value() float64
		SetValue(float64)
		Min() float64
		SetMin(float64)
		Max() float64
		SetMax(float64)
		Low() float64
		SetLow(float64)
		High() float64
		SetHigh(float64)
		Optimum() float64
		SetOptimum(float64)
		Labels() []Node
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmlfieldsetelement
	HTMLFieldSetElement interface {
		HTMLElement

		Disabled() bool
		SetDisabled(bool)
		Form() HTMLFormElement
		Name() string
		SetName(string)
		Type() string
		Elements() HTMLCollection
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#htmllegendelement
	HTMLLegendElement interface {
		HTMLElement

		Form() HTMLFormElement
	}
)

// https://www.w3.org/TR/html52/sec-forms.html#enumdef-selectionmode
type SelectionMode string

const (
	SelectionModeSelect   SelectionMode = "select"
	SelectionModeStart    SelectionMode = "start"
	SelectionModeEnd      SelectionMode = "end"
	SelectionModePreserve SelectionMode = "preserve"
)

// https://www.w3.org/TR/html52/sec-forms.html#htmloptionelement
func NewOption(...OptionParams) HTMLOptionElement {
	return nil
}
