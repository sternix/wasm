
// +build js,wasm

package wasm

/*
https://www.w3.org/Style/CSS/all-properties#list
*/

type cssStyleHelper interface {
	
	Azimuth() string
	SetAzimuth(string)
	Background() string
	SetBackground(string)
	BackgroundAttachment() string
	SetBackgroundAttachment(string)
	BackgroundColor() string
	SetBackgroundColor(string)
	BackgroundImage() string
	SetBackgroundImage(string)
	BackgroundPosition() string
	SetBackgroundPosition(string)
	BackgroundRepeat() string
	SetBackgroundRepeat(string)
	Border() string
	SetBorder(string)
	BorderBottom() string
	SetBorderBottom(string)
	BorderBottomColor() string
	SetBorderBottomColor(string)
	BorderBottomStyle() string
	SetBorderBottomStyle(string)
	BorderBottomWidth() string
	SetBorderBottomWidth(string)
	BorderCollapse() string
	SetBorderCollapse(string)
	BorderColor() string
	SetBorderColor(string)
	BorderLeft() string
	SetBorderLeft(string)
	BorderLeftColor() string
	SetBorderLeftColor(string)
	BorderLeftStyle() string
	SetBorderLeftStyle(string)
	BorderLeftWidth() string
	SetBorderLeftWidth(string)
	BorderRight() string
	SetBorderRight(string)
	BorderRightColor() string
	SetBorderRightColor(string)
	BorderRightStyle() string
	SetBorderRightStyle(string)
	BorderRightWidth() string
	SetBorderRightWidth(string)
	BorderSpacing() string
	SetBorderSpacing(string)
	BorderStyle() string
	SetBorderStyle(string)
	BorderTop() string
	SetBorderTop(string)
	BorderTopColor() string
	SetBorderTopColor(string)
	BorderTopStyle() string
	SetBorderTopStyle(string)
	BorderTopWidth() string
	SetBorderTopWidth(string)
	BorderWidth() string
	SetBorderWidth(string)
	Bottom() string
	SetBottom(string)
	BoxSizing() string
	SetBoxSizing(string)
	CaptionSide() string
	SetCaptionSide(string)
	CaretColor() string
	SetCaretColor(string)
	Clear() string
	SetClear(string)
	Clip() string
	SetClip(string)
	Color() string
	SetColor(string)
	Content() string
	SetContent(string)
	CounterIncrement() string
	SetCounterIncrement(string)
	CounterReset() string
	SetCounterReset(string)
	Cue() string
	SetCue(string)
	CueAfter() string
	SetCueAfter(string)
	CueBefore() string
	SetCueBefore(string)
	Cursor() string
	SetCursor(string)
	Direction() string
	SetDirection(string)
	Display() string
	SetDisplay(string)
	Elevation() string
	SetElevation(string)
	EmptyCells() string
	SetEmptyCells(string)
	CssFloat() string
	SetCssFloat(string)
	Font() string
	SetFont(string)
	FontFamily() string
	SetFontFamily(string)
	FontFeatureSettings() string
	SetFontFeatureSettings(string)
	FontKerning() string
	SetFontKerning(string)
	FontSize() string
	SetFontSize(string)
	FontSizeAdjust() string
	SetFontSizeAdjust(string)
	FontStretch() string
	SetFontStretch(string)
	FontStyle() string
	SetFontStyle(string)
	FontSynthesis() string
	SetFontSynthesis(string)
	FontVariant() string
	SetFontVariant(string)
	FontVariantCaps() string
	SetFontVariantCaps(string)
	FontVariantEastAsian() string
	SetFontVariantEastAsian(string)
	FontVariantLigatures() string
	SetFontVariantLigatures(string)
	FontVariantNumeric() string
	SetFontVariantNumeric(string)
	FontVariantPosition() string
	SetFontVariantPosition(string)
	FontWeight() string
	SetFontWeight(string)
	Height() string
	SetHeight(string)
	Left() string
	SetLeft(string)
	LetterSpacing() string
	SetLetterSpacing(string)
	LineHeight() string
	SetLineHeight(string)
	ListStyle() string
	SetListStyle(string)
	ListStyleImage() string
	SetListStyleImage(string)
	ListStylePosition() string
	SetListStylePosition(string)
	ListStyleType() string
	SetListStyleType(string)
	Margin() string
	SetMargin(string)
	MarginBottom() string
	SetMarginBottom(string)
	MarginLeft() string
	SetMarginLeft(string)
	MarginRight() string
	SetMarginRight(string)
	MarginTop() string
	SetMarginTop(string)
	MaxHeight() string
	SetMaxHeight(string)
	MaxWidth() string
	SetMaxWidth(string)
	MinHeight() string
	SetMinHeight(string)
	MinWidth() string
	SetMinWidth(string)
	Opacity() string
	SetOpacity(string)
	Orphans() string
	SetOrphans(string)
	Outline() string
	SetOutline(string)
	OutlineColor() string
	SetOutlineColor(string)
	OutlineOffset() string
	SetOutlineOffset(string)
	OutlineStyle() string
	SetOutlineStyle(string)
	OutlineWidth() string
	SetOutlineWidth(string)
	Overflow() string
	SetOverflow(string)
	Padding() string
	SetPadding(string)
	PaddingBottom() string
	SetPaddingBottom(string)
	PaddingLeft() string
	SetPaddingLeft(string)
	PaddingRight() string
	SetPaddingRight(string)
	PaddingTop() string
	SetPaddingTop(string)
	PageBreakAfter() string
	SetPageBreakAfter(string)
	PageBreakBefore() string
	SetPageBreakBefore(string)
	PageBreakInside() string
	SetPageBreakInside(string)
	Pause() string
	SetPause(string)
	PauseAfter() string
	SetPauseAfter(string)
	PauseBefore() string
	SetPauseBefore(string)
	Pitch() string
	SetPitch(string)
	PitchRange() string
	SetPitchRange(string)
	PlayDuring() string
	SetPlayDuring(string)
	Position() string
	SetPosition(string)
	Quotes() string
	SetQuotes(string)
	Resize() string
	SetResize(string)
	Richness() string
	SetRichness(string)
	Right() string
	SetRight(string)
	Speak() string
	SetSpeak(string)
	SpeakHeader() string
	SetSpeakHeader(string)
	SpeakNumeral() string
	SetSpeakNumeral(string)
	SpeakPunctuation() string
	SetSpeakPunctuation(string)
	SpeechRate() string
	SetSpeechRate(string)
	Stress() string
	SetStress(string)
	TableLayout() string
	SetTableLayout(string)
	TextAlign() string
	SetTextAlign(string)
	TextDecoration() string
	SetTextDecoration(string)
	TextIndent() string
	SetTextIndent(string)
	TextOverflow() string
	SetTextOverflow(string)
	TextTransform() string
	SetTextTransform(string)
	Top() string
	SetTop(string)
	UnicodeBidi() string
	SetUnicodeBidi(string)
	VerticalAlign() string
	SetVerticalAlign(string)
	Visibility() string
	SetVisibility(string)
	VoiceFamily() string
	SetVoiceFamily(string)
	Volume() string
	SetVolume(string)
	WhiteSpace() string
	SetWhiteSpace(string)
	Widows() string
	SetWidows(string)
	Width() string
	SetWidth(string)
	WordSpacing() string
	SetWordSpacing(string)
	ZIndex() string
	SetZIndex(string)
}

type cssStyleHelperImpl struct {
	*cssStyleDeclarationImpl
}

func newCSSStyleHelperImpl(v *cssStyleDeclarationImpl) *cssStyleHelperImpl {
	if v.Valid() {
		return &cssStyleHelperImpl {
			cssStyleDeclarationImpl: v,
		}
	}
	return nil
}


func (p *cssStyleHelperImpl) Azimuth() string {
	return p.PropertyValue("azimuth")
}

func (p *cssStyleHelperImpl) SetAzimuth(s string) {
	p.SetProperty("azimuth",s)
}

func (p *cssStyleHelperImpl) Background() string {
	return p.PropertyValue("background")
}

func (p *cssStyleHelperImpl) SetBackground(s string) {
	p.SetProperty("background",s)
}

func (p *cssStyleHelperImpl) BackgroundAttachment() string {
	return p.PropertyValue("background-attachment")
}

func (p *cssStyleHelperImpl) SetBackgroundAttachment(s string) {
	p.SetProperty("background-attachment",s)
}

func (p *cssStyleHelperImpl) BackgroundColor() string {
	return p.PropertyValue("background-color")
}

func (p *cssStyleHelperImpl) SetBackgroundColor(s string) {
	p.SetProperty("background-color",s)
}

func (p *cssStyleHelperImpl) BackgroundImage() string {
	return p.PropertyValue("background-image")
}

func (p *cssStyleHelperImpl) SetBackgroundImage(s string) {
	p.SetProperty("background-image",s)
}

func (p *cssStyleHelperImpl) BackgroundPosition() string {
	return p.PropertyValue("background-position")
}

func (p *cssStyleHelperImpl) SetBackgroundPosition(s string) {
	p.SetProperty("background-position",s)
}

func (p *cssStyleHelperImpl) BackgroundRepeat() string {
	return p.PropertyValue("background-repeat")
}

func (p *cssStyleHelperImpl) SetBackgroundRepeat(s string) {
	p.SetProperty("background-repeat",s)
}

func (p *cssStyleHelperImpl) Border() string {
	return p.PropertyValue("border")
}

func (p *cssStyleHelperImpl) SetBorder(s string) {
	p.SetProperty("border",s)
}

func (p *cssStyleHelperImpl) BorderBottom() string {
	return p.PropertyValue("border-bottom")
}

func (p *cssStyleHelperImpl) SetBorderBottom(s string) {
	p.SetProperty("border-bottom",s)
}

func (p *cssStyleHelperImpl) BorderBottomColor() string {
	return p.PropertyValue("border-bottom-color")
}

func (p *cssStyleHelperImpl) SetBorderBottomColor(s string) {
	p.SetProperty("border-bottom-color",s)
}

func (p *cssStyleHelperImpl) BorderBottomStyle() string {
	return p.PropertyValue("border-bottom-style")
}

func (p *cssStyleHelperImpl) SetBorderBottomStyle(s string) {
	p.SetProperty("border-bottom-style",s)
}

func (p *cssStyleHelperImpl) BorderBottomWidth() string {
	return p.PropertyValue("border-bottom-width")
}

func (p *cssStyleHelperImpl) SetBorderBottomWidth(s string) {
	p.SetProperty("border-bottom-width",s)
}

func (p *cssStyleHelperImpl) BorderCollapse() string {
	return p.PropertyValue("border-collapse")
}

func (p *cssStyleHelperImpl) SetBorderCollapse(s string) {
	p.SetProperty("border-collapse",s)
}

func (p *cssStyleHelperImpl) BorderColor() string {
	return p.PropertyValue("border-color")
}

func (p *cssStyleHelperImpl) SetBorderColor(s string) {
	p.SetProperty("border-color",s)
}

func (p *cssStyleHelperImpl) BorderLeft() string {
	return p.PropertyValue("border-left")
}

func (p *cssStyleHelperImpl) SetBorderLeft(s string) {
	p.SetProperty("border-left",s)
}

func (p *cssStyleHelperImpl) BorderLeftColor() string {
	return p.PropertyValue("border-left-color")
}

func (p *cssStyleHelperImpl) SetBorderLeftColor(s string) {
	p.SetProperty("border-left-color",s)
}

func (p *cssStyleHelperImpl) BorderLeftStyle() string {
	return p.PropertyValue("border-left-style")
}

func (p *cssStyleHelperImpl) SetBorderLeftStyle(s string) {
	p.SetProperty("border-left-style",s)
}

func (p *cssStyleHelperImpl) BorderLeftWidth() string {
	return p.PropertyValue("border-left-width")
}

func (p *cssStyleHelperImpl) SetBorderLeftWidth(s string) {
	p.SetProperty("border-left-width",s)
}

func (p *cssStyleHelperImpl) BorderRight() string {
	return p.PropertyValue("border-right")
}

func (p *cssStyleHelperImpl) SetBorderRight(s string) {
	p.SetProperty("border-right",s)
}

func (p *cssStyleHelperImpl) BorderRightColor() string {
	return p.PropertyValue("border-right-color")
}

func (p *cssStyleHelperImpl) SetBorderRightColor(s string) {
	p.SetProperty("border-right-color",s)
}

func (p *cssStyleHelperImpl) BorderRightStyle() string {
	return p.PropertyValue("border-right-style")
}

func (p *cssStyleHelperImpl) SetBorderRightStyle(s string) {
	p.SetProperty("border-right-style",s)
}

func (p *cssStyleHelperImpl) BorderRightWidth() string {
	return p.PropertyValue("border-right-width")
}

func (p *cssStyleHelperImpl) SetBorderRightWidth(s string) {
	p.SetProperty("border-right-width",s)
}

func (p *cssStyleHelperImpl) BorderSpacing() string {
	return p.PropertyValue("border-spacing")
}

func (p *cssStyleHelperImpl) SetBorderSpacing(s string) {
	p.SetProperty("border-spacing",s)
}

func (p *cssStyleHelperImpl) BorderStyle() string {
	return p.PropertyValue("border-style")
}

func (p *cssStyleHelperImpl) SetBorderStyle(s string) {
	p.SetProperty("border-style",s)
}

func (p *cssStyleHelperImpl) BorderTop() string {
	return p.PropertyValue("border-top")
}

func (p *cssStyleHelperImpl) SetBorderTop(s string) {
	p.SetProperty("border-top",s)
}

func (p *cssStyleHelperImpl) BorderTopColor() string {
	return p.PropertyValue("border-top-color")
}

func (p *cssStyleHelperImpl) SetBorderTopColor(s string) {
	p.SetProperty("border-top-color",s)
}

func (p *cssStyleHelperImpl) BorderTopStyle() string {
	return p.PropertyValue("border-top-style")
}

func (p *cssStyleHelperImpl) SetBorderTopStyle(s string) {
	p.SetProperty("border-top-style",s)
}

func (p *cssStyleHelperImpl) BorderTopWidth() string {
	return p.PropertyValue("border-top-width")
}

func (p *cssStyleHelperImpl) SetBorderTopWidth(s string) {
	p.SetProperty("border-top-width",s)
}

func (p *cssStyleHelperImpl) BorderWidth() string {
	return p.PropertyValue("border-width")
}

func (p *cssStyleHelperImpl) SetBorderWidth(s string) {
	p.SetProperty("border-width",s)
}

func (p *cssStyleHelperImpl) Bottom() string {
	return p.PropertyValue("bottom")
}

func (p *cssStyleHelperImpl) SetBottom(s string) {
	p.SetProperty("bottom",s)
}

func (p *cssStyleHelperImpl) BoxSizing() string {
	return p.PropertyValue("box-sizing")
}

func (p *cssStyleHelperImpl) SetBoxSizing(s string) {
	p.SetProperty("box-sizing",s)
}

func (p *cssStyleHelperImpl) CaptionSide() string {
	return p.PropertyValue("caption-side")
}

func (p *cssStyleHelperImpl) SetCaptionSide(s string) {
	p.SetProperty("caption-side",s)
}

func (p *cssStyleHelperImpl) CaretColor() string {
	return p.PropertyValue("caret-color")
}

func (p *cssStyleHelperImpl) SetCaretColor(s string) {
	p.SetProperty("caret-color",s)
}

func (p *cssStyleHelperImpl) Clear() string {
	return p.PropertyValue("clear")
}

func (p *cssStyleHelperImpl) SetClear(s string) {
	p.SetProperty("clear",s)
}

func (p *cssStyleHelperImpl) Clip() string {
	return p.PropertyValue("clip")
}

func (p *cssStyleHelperImpl) SetClip(s string) {
	p.SetProperty("clip",s)
}

func (p *cssStyleHelperImpl) Color() string {
	return p.PropertyValue("color")
}

func (p *cssStyleHelperImpl) SetColor(s string) {
	p.SetProperty("color",s)
}

func (p *cssStyleHelperImpl) Content() string {
	return p.PropertyValue("content")
}

func (p *cssStyleHelperImpl) SetContent(s string) {
	p.SetProperty("content",s)
}

func (p *cssStyleHelperImpl) CounterIncrement() string {
	return p.PropertyValue("counter-increment")
}

func (p *cssStyleHelperImpl) SetCounterIncrement(s string) {
	p.SetProperty("counter-increment",s)
}

func (p *cssStyleHelperImpl) CounterReset() string {
	return p.PropertyValue("counter-reset")
}

func (p *cssStyleHelperImpl) SetCounterReset(s string) {
	p.SetProperty("counter-reset",s)
}

func (p *cssStyleHelperImpl) Cue() string {
	return p.PropertyValue("cue")
}

func (p *cssStyleHelperImpl) SetCue(s string) {
	p.SetProperty("cue",s)
}

func (p *cssStyleHelperImpl) CueAfter() string {
	return p.PropertyValue("cue-after")
}

func (p *cssStyleHelperImpl) SetCueAfter(s string) {
	p.SetProperty("cue-after",s)
}

func (p *cssStyleHelperImpl) CueBefore() string {
	return p.PropertyValue("cue-before")
}

func (p *cssStyleHelperImpl) SetCueBefore(s string) {
	p.SetProperty("cue-before",s)
}

func (p *cssStyleHelperImpl) Cursor() string {
	return p.PropertyValue("cursor")
}

func (p *cssStyleHelperImpl) SetCursor(s string) {
	p.SetProperty("cursor",s)
}

func (p *cssStyleHelperImpl) Direction() string {
	return p.PropertyValue("direction")
}

func (p *cssStyleHelperImpl) SetDirection(s string) {
	p.SetProperty("direction",s)
}

func (p *cssStyleHelperImpl) Display() string {
	return p.PropertyValue("display")
}

func (p *cssStyleHelperImpl) SetDisplay(s string) {
	p.SetProperty("display",s)
}

func (p *cssStyleHelperImpl) Elevation() string {
	return p.PropertyValue("elevation")
}

func (p *cssStyleHelperImpl) SetElevation(s string) {
	p.SetProperty("elevation",s)
}

func (p *cssStyleHelperImpl) EmptyCells() string {
	return p.PropertyValue("empty-cells")
}

func (p *cssStyleHelperImpl) SetEmptyCells(s string) {
	p.SetProperty("empty-cells",s)
}

func (p *cssStyleHelperImpl) CssFloat() string {
	return p.PropertyValue("css-float")
}

func (p *cssStyleHelperImpl) SetCssFloat(s string) {
	p.SetProperty("css-float",s)
}

func (p *cssStyleHelperImpl) Font() string {
	return p.PropertyValue("font")
}

func (p *cssStyleHelperImpl) SetFont(s string) {
	p.SetProperty("font",s)
}

func (p *cssStyleHelperImpl) FontFamily() string {
	return p.PropertyValue("font-family")
}

func (p *cssStyleHelperImpl) SetFontFamily(s string) {
	p.SetProperty("font-family",s)
}

func (p *cssStyleHelperImpl) FontFeatureSettings() string {
	return p.PropertyValue("font-feature-settings")
}

func (p *cssStyleHelperImpl) SetFontFeatureSettings(s string) {
	p.SetProperty("font-feature-settings",s)
}

func (p *cssStyleHelperImpl) FontKerning() string {
	return p.PropertyValue("font-kerning")
}

func (p *cssStyleHelperImpl) SetFontKerning(s string) {
	p.SetProperty("font-kerning",s)
}

func (p *cssStyleHelperImpl) FontSize() string {
	return p.PropertyValue("font-size")
}

func (p *cssStyleHelperImpl) SetFontSize(s string) {
	p.SetProperty("font-size",s)
}

func (p *cssStyleHelperImpl) FontSizeAdjust() string {
	return p.PropertyValue("font-size-adjust")
}

func (p *cssStyleHelperImpl) SetFontSizeAdjust(s string) {
	p.SetProperty("font-size-adjust",s)
}

func (p *cssStyleHelperImpl) FontStretch() string {
	return p.PropertyValue("font-stretch")
}

func (p *cssStyleHelperImpl) SetFontStretch(s string) {
	p.SetProperty("font-stretch",s)
}

func (p *cssStyleHelperImpl) FontStyle() string {
	return p.PropertyValue("font-style")
}

func (p *cssStyleHelperImpl) SetFontStyle(s string) {
	p.SetProperty("font-style",s)
}

func (p *cssStyleHelperImpl) FontSynthesis() string {
	return p.PropertyValue("font-synthesis")
}

func (p *cssStyleHelperImpl) SetFontSynthesis(s string) {
	p.SetProperty("font-synthesis",s)
}

func (p *cssStyleHelperImpl) FontVariant() string {
	return p.PropertyValue("font-variant")
}

func (p *cssStyleHelperImpl) SetFontVariant(s string) {
	p.SetProperty("font-variant",s)
}

func (p *cssStyleHelperImpl) FontVariantCaps() string {
	return p.PropertyValue("font-variant-caps")
}

func (p *cssStyleHelperImpl) SetFontVariantCaps(s string) {
	p.SetProperty("font-variant-caps",s)
}

func (p *cssStyleHelperImpl) FontVariantEastAsian() string {
	return p.PropertyValue("font-variant-east-asian")
}

func (p *cssStyleHelperImpl) SetFontVariantEastAsian(s string) {
	p.SetProperty("font-variant-east-asian",s)
}

func (p *cssStyleHelperImpl) FontVariantLigatures() string {
	return p.PropertyValue("font-variant-ligatures")
}

func (p *cssStyleHelperImpl) SetFontVariantLigatures(s string) {
	p.SetProperty("font-variant-ligatures",s)
}

func (p *cssStyleHelperImpl) FontVariantNumeric() string {
	return p.PropertyValue("font-variant-numeric")
}

func (p *cssStyleHelperImpl) SetFontVariantNumeric(s string) {
	p.SetProperty("font-variant-numeric",s)
}

func (p *cssStyleHelperImpl) FontVariantPosition() string {
	return p.PropertyValue("font-variant-position")
}

func (p *cssStyleHelperImpl) SetFontVariantPosition(s string) {
	p.SetProperty("font-variant-position",s)
}

func (p *cssStyleHelperImpl) FontWeight() string {
	return p.PropertyValue("font-weight")
}

func (p *cssStyleHelperImpl) SetFontWeight(s string) {
	p.SetProperty("font-weight",s)
}

func (p *cssStyleHelperImpl) Height() string {
	return p.PropertyValue("height")
}

func (p *cssStyleHelperImpl) SetHeight(s string) {
	p.SetProperty("height",s)
}

func (p *cssStyleHelperImpl) Left() string {
	return p.PropertyValue("left")
}

func (p *cssStyleHelperImpl) SetLeft(s string) {
	p.SetProperty("left",s)
}

func (p *cssStyleHelperImpl) LetterSpacing() string {
	return p.PropertyValue("letter-spacing")
}

func (p *cssStyleHelperImpl) SetLetterSpacing(s string) {
	p.SetProperty("letter-spacing",s)
}

func (p *cssStyleHelperImpl) LineHeight() string {
	return p.PropertyValue("line-height")
}

func (p *cssStyleHelperImpl) SetLineHeight(s string) {
	p.SetProperty("line-height",s)
}

func (p *cssStyleHelperImpl) ListStyle() string {
	return p.PropertyValue("list-style")
}

func (p *cssStyleHelperImpl) SetListStyle(s string) {
	p.SetProperty("list-style",s)
}

func (p *cssStyleHelperImpl) ListStyleImage() string {
	return p.PropertyValue("list-style-image")
}

func (p *cssStyleHelperImpl) SetListStyleImage(s string) {
	p.SetProperty("list-style-image",s)
}

func (p *cssStyleHelperImpl) ListStylePosition() string {
	return p.PropertyValue("list-style-position")
}

func (p *cssStyleHelperImpl) SetListStylePosition(s string) {
	p.SetProperty("list-style-position",s)
}

func (p *cssStyleHelperImpl) ListStyleType() string {
	return p.PropertyValue("list-style-type")
}

func (p *cssStyleHelperImpl) SetListStyleType(s string) {
	p.SetProperty("list-style-type",s)
}

func (p *cssStyleHelperImpl) Margin() string {
	return p.PropertyValue("margin")
}

func (p *cssStyleHelperImpl) SetMargin(s string) {
	p.SetProperty("margin",s)
}

func (p *cssStyleHelperImpl) MarginBottom() string {
	return p.PropertyValue("margin-bottom")
}

func (p *cssStyleHelperImpl) SetMarginBottom(s string) {
	p.SetProperty("margin-bottom",s)
}

func (p *cssStyleHelperImpl) MarginLeft() string {
	return p.PropertyValue("margin-left")
}

func (p *cssStyleHelperImpl) SetMarginLeft(s string) {
	p.SetProperty("margin-left",s)
}

func (p *cssStyleHelperImpl) MarginRight() string {
	return p.PropertyValue("margin-right")
}

func (p *cssStyleHelperImpl) SetMarginRight(s string) {
	p.SetProperty("margin-right",s)
}

func (p *cssStyleHelperImpl) MarginTop() string {
	return p.PropertyValue("margin-top")
}

func (p *cssStyleHelperImpl) SetMarginTop(s string) {
	p.SetProperty("margin-top",s)
}

func (p *cssStyleHelperImpl) MaxHeight() string {
	return p.PropertyValue("max-height")
}

func (p *cssStyleHelperImpl) SetMaxHeight(s string) {
	p.SetProperty("max-height",s)
}

func (p *cssStyleHelperImpl) MaxWidth() string {
	return p.PropertyValue("max-width")
}

func (p *cssStyleHelperImpl) SetMaxWidth(s string) {
	p.SetProperty("max-width",s)
}

func (p *cssStyleHelperImpl) MinHeight() string {
	return p.PropertyValue("min-height")
}

func (p *cssStyleHelperImpl) SetMinHeight(s string) {
	p.SetProperty("min-height",s)
}

func (p *cssStyleHelperImpl) MinWidth() string {
	return p.PropertyValue("min-width")
}

func (p *cssStyleHelperImpl) SetMinWidth(s string) {
	p.SetProperty("min-width",s)
}

func (p *cssStyleHelperImpl) Opacity() string {
	return p.PropertyValue("opacity")
}

func (p *cssStyleHelperImpl) SetOpacity(s string) {
	p.SetProperty("opacity",s)
}

func (p *cssStyleHelperImpl) Orphans() string {
	return p.PropertyValue("orphans")
}

func (p *cssStyleHelperImpl) SetOrphans(s string) {
	p.SetProperty("orphans",s)
}

func (p *cssStyleHelperImpl) Outline() string {
	return p.PropertyValue("outline")
}

func (p *cssStyleHelperImpl) SetOutline(s string) {
	p.SetProperty("outline",s)
}

func (p *cssStyleHelperImpl) OutlineColor() string {
	return p.PropertyValue("outline-color")
}

func (p *cssStyleHelperImpl) SetOutlineColor(s string) {
	p.SetProperty("outline-color",s)
}

func (p *cssStyleHelperImpl) OutlineOffset() string {
	return p.PropertyValue("outline-offset")
}

func (p *cssStyleHelperImpl) SetOutlineOffset(s string) {
	p.SetProperty("outline-offset",s)
}

func (p *cssStyleHelperImpl) OutlineStyle() string {
	return p.PropertyValue("outline-style")
}

func (p *cssStyleHelperImpl) SetOutlineStyle(s string) {
	p.SetProperty("outline-style",s)
}

func (p *cssStyleHelperImpl) OutlineWidth() string {
	return p.PropertyValue("outline-width")
}

func (p *cssStyleHelperImpl) SetOutlineWidth(s string) {
	p.SetProperty("outline-width",s)
}

func (p *cssStyleHelperImpl) Overflow() string {
	return p.PropertyValue("overflow")
}

func (p *cssStyleHelperImpl) SetOverflow(s string) {
	p.SetProperty("overflow",s)
}

func (p *cssStyleHelperImpl) Padding() string {
	return p.PropertyValue("padding")
}

func (p *cssStyleHelperImpl) SetPadding(s string) {
	p.SetProperty("padding",s)
}

func (p *cssStyleHelperImpl) PaddingBottom() string {
	return p.PropertyValue("padding-bottom")
}

func (p *cssStyleHelperImpl) SetPaddingBottom(s string) {
	p.SetProperty("padding-bottom",s)
}

func (p *cssStyleHelperImpl) PaddingLeft() string {
	return p.PropertyValue("padding-left")
}

func (p *cssStyleHelperImpl) SetPaddingLeft(s string) {
	p.SetProperty("padding-left",s)
}

func (p *cssStyleHelperImpl) PaddingRight() string {
	return p.PropertyValue("padding-right")
}

func (p *cssStyleHelperImpl) SetPaddingRight(s string) {
	p.SetProperty("padding-right",s)
}

func (p *cssStyleHelperImpl) PaddingTop() string {
	return p.PropertyValue("padding-top")
}

func (p *cssStyleHelperImpl) SetPaddingTop(s string) {
	p.SetProperty("padding-top",s)
}

func (p *cssStyleHelperImpl) PageBreakAfter() string {
	return p.PropertyValue("page-break-after")
}

func (p *cssStyleHelperImpl) SetPageBreakAfter(s string) {
	p.SetProperty("page-break-after",s)
}

func (p *cssStyleHelperImpl) PageBreakBefore() string {
	return p.PropertyValue("page-break-before")
}

func (p *cssStyleHelperImpl) SetPageBreakBefore(s string) {
	p.SetProperty("page-break-before",s)
}

func (p *cssStyleHelperImpl) PageBreakInside() string {
	return p.PropertyValue("page-break-inside")
}

func (p *cssStyleHelperImpl) SetPageBreakInside(s string) {
	p.SetProperty("page-break-inside",s)
}

func (p *cssStyleHelperImpl) Pause() string {
	return p.PropertyValue("pause")
}

func (p *cssStyleHelperImpl) SetPause(s string) {
	p.SetProperty("pause",s)
}

func (p *cssStyleHelperImpl) PauseAfter() string {
	return p.PropertyValue("pause-after")
}

func (p *cssStyleHelperImpl) SetPauseAfter(s string) {
	p.SetProperty("pause-after",s)
}

func (p *cssStyleHelperImpl) PauseBefore() string {
	return p.PropertyValue("pause-before")
}

func (p *cssStyleHelperImpl) SetPauseBefore(s string) {
	p.SetProperty("pause-before",s)
}

func (p *cssStyleHelperImpl) Pitch() string {
	return p.PropertyValue("pitch")
}

func (p *cssStyleHelperImpl) SetPitch(s string) {
	p.SetProperty("pitch",s)
}

func (p *cssStyleHelperImpl) PitchRange() string {
	return p.PropertyValue("pitch-range")
}

func (p *cssStyleHelperImpl) SetPitchRange(s string) {
	p.SetProperty("pitch-range",s)
}

func (p *cssStyleHelperImpl) PlayDuring() string {
	return p.PropertyValue("play-during")
}

func (p *cssStyleHelperImpl) SetPlayDuring(s string) {
	p.SetProperty("play-during",s)
}

func (p *cssStyleHelperImpl) Position() string {
	return p.PropertyValue("position")
}

func (p *cssStyleHelperImpl) SetPosition(s string) {
	p.SetProperty("position",s)
}

func (p *cssStyleHelperImpl) Quotes() string {
	return p.PropertyValue("quotes")
}

func (p *cssStyleHelperImpl) SetQuotes(s string) {
	p.SetProperty("quotes",s)
}

func (p *cssStyleHelperImpl) Resize() string {
	return p.PropertyValue("resize")
}

func (p *cssStyleHelperImpl) SetResize(s string) {
	p.SetProperty("resize",s)
}

func (p *cssStyleHelperImpl) Richness() string {
	return p.PropertyValue("richness")
}

func (p *cssStyleHelperImpl) SetRichness(s string) {
	p.SetProperty("richness",s)
}

func (p *cssStyleHelperImpl) Right() string {
	return p.PropertyValue("right")
}

func (p *cssStyleHelperImpl) SetRight(s string) {
	p.SetProperty("right",s)
}

func (p *cssStyleHelperImpl) Speak() string {
	return p.PropertyValue("speak")
}

func (p *cssStyleHelperImpl) SetSpeak(s string) {
	p.SetProperty("speak",s)
}

func (p *cssStyleHelperImpl) SpeakHeader() string {
	return p.PropertyValue("speak-header")
}

func (p *cssStyleHelperImpl) SetSpeakHeader(s string) {
	p.SetProperty("speak-header",s)
}

func (p *cssStyleHelperImpl) SpeakNumeral() string {
	return p.PropertyValue("speak-numeral")
}

func (p *cssStyleHelperImpl) SetSpeakNumeral(s string) {
	p.SetProperty("speak-numeral",s)
}

func (p *cssStyleHelperImpl) SpeakPunctuation() string {
	return p.PropertyValue("speak-punctuation")
}

func (p *cssStyleHelperImpl) SetSpeakPunctuation(s string) {
	p.SetProperty("speak-punctuation",s)
}

func (p *cssStyleHelperImpl) SpeechRate() string {
	return p.PropertyValue("speech-rate")
}

func (p *cssStyleHelperImpl) SetSpeechRate(s string) {
	p.SetProperty("speech-rate",s)
}

func (p *cssStyleHelperImpl) Stress() string {
	return p.PropertyValue("stress")
}

func (p *cssStyleHelperImpl) SetStress(s string) {
	p.SetProperty("stress",s)
}

func (p *cssStyleHelperImpl) TableLayout() string {
	return p.PropertyValue("table-layout")
}

func (p *cssStyleHelperImpl) SetTableLayout(s string) {
	p.SetProperty("table-layout",s)
}

func (p *cssStyleHelperImpl) TextAlign() string {
	return p.PropertyValue("text-align")
}

func (p *cssStyleHelperImpl) SetTextAlign(s string) {
	p.SetProperty("text-align",s)
}

func (p *cssStyleHelperImpl) TextDecoration() string {
	return p.PropertyValue("text-decoration")
}

func (p *cssStyleHelperImpl) SetTextDecoration(s string) {
	p.SetProperty("text-decoration",s)
}

func (p *cssStyleHelperImpl) TextIndent() string {
	return p.PropertyValue("text-indent")
}

func (p *cssStyleHelperImpl) SetTextIndent(s string) {
	p.SetProperty("text-indent",s)
}

func (p *cssStyleHelperImpl) TextOverflow() string {
	return p.PropertyValue("text-overflow")
}

func (p *cssStyleHelperImpl) SetTextOverflow(s string) {
	p.SetProperty("text-overflow",s)
}

func (p *cssStyleHelperImpl) TextTransform() string {
	return p.PropertyValue("text-transform")
}

func (p *cssStyleHelperImpl) SetTextTransform(s string) {
	p.SetProperty("text-transform",s)
}

func (p *cssStyleHelperImpl) Top() string {
	return p.PropertyValue("top")
}

func (p *cssStyleHelperImpl) SetTop(s string) {
	p.SetProperty("top",s)
}

func (p *cssStyleHelperImpl) UnicodeBidi() string {
	return p.PropertyValue("unicode-bidi")
}

func (p *cssStyleHelperImpl) SetUnicodeBidi(s string) {
	p.SetProperty("unicode-bidi",s)
}

func (p *cssStyleHelperImpl) VerticalAlign() string {
	return p.PropertyValue("vertical-align")
}

func (p *cssStyleHelperImpl) SetVerticalAlign(s string) {
	p.SetProperty("vertical-align",s)
}

func (p *cssStyleHelperImpl) Visibility() string {
	return p.PropertyValue("visibility")
}

func (p *cssStyleHelperImpl) SetVisibility(s string) {
	p.SetProperty("visibility",s)
}

func (p *cssStyleHelperImpl) VoiceFamily() string {
	return p.PropertyValue("voice-family")
}

func (p *cssStyleHelperImpl) SetVoiceFamily(s string) {
	p.SetProperty("voice-family",s)
}

func (p *cssStyleHelperImpl) Volume() string {
	return p.PropertyValue("volume")
}

func (p *cssStyleHelperImpl) SetVolume(s string) {
	p.SetProperty("volume",s)
}

func (p *cssStyleHelperImpl) WhiteSpace() string {
	return p.PropertyValue("white-space")
}

func (p *cssStyleHelperImpl) SetWhiteSpace(s string) {
	p.SetProperty("white-space",s)
}

func (p *cssStyleHelperImpl) Widows() string {
	return p.PropertyValue("widows")
}

func (p *cssStyleHelperImpl) SetWidows(s string) {
	p.SetProperty("widows",s)
}

func (p *cssStyleHelperImpl) Width() string {
	return p.PropertyValue("width")
}

func (p *cssStyleHelperImpl) SetWidth(s string) {
	p.SetProperty("width",s)
}

func (p *cssStyleHelperImpl) WordSpacing() string {
	return p.PropertyValue("word-spacing")
}

func (p *cssStyleHelperImpl) SetWordSpacing(s string) {
	p.SetProperty("word-spacing",s)
}

func (p *cssStyleHelperImpl) ZIndex() string {
	return p.PropertyValue("z-index")
}

func (p *cssStyleHelperImpl) SetZIndex(s string) {
	p.SetProperty("z-index",s)
}

