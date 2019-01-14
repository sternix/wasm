// +build js,wasm

package wasm

/*
https://www.w3.org/Style/CSS/all-properties#list
*/

type cssStyleHelper interface {
	AlignContent() string
	SetAlignContent(string)
	AlignItems() string
	SetAlignItems(string)
	AlignSelf() string
	SetAlignSelf(string)
	AlignmentBaseline() string
	SetAlignmentBaseline(string)
	All() string
	SetAll(string)
	Animation() string
	SetAnimation(string)
	AnimationDelay() string
	SetAnimationDelay(string)
	AnimationDirection() string
	SetAnimationDirection(string)
	AnimationDuration() string
	SetAnimationDuration(string)
	AnimationFillMode() string
	SetAnimationFillMode(string)
	AnimationIterationCount() string
	SetAnimationIterationCount(string)
	AnimationName() string
	SetAnimationName(string)
	AnimationPlayState() string
	SetAnimationPlayState(string)
	AnimationTimingFunction() string
	SetAnimationTimingFunction(string)
	Appearance() string
	SetAppearance(string)
	Azimuth() string
	SetAzimuth(string)
	Background() string
	SetBackground(string)
	BackgroundAttachment() string
	SetBackgroundAttachment(string)
	BackgroundBlendMode() string
	SetBackgroundBlendMode(string)
	BackgroundClip() string
	SetBackgroundClip(string)
	BackgroundColor() string
	SetBackgroundColor(string)
	BackgroundImage() string
	SetBackgroundImage(string)
	BackgroundOrigin() string
	SetBackgroundOrigin(string)
	BackgroundPosition() string
	SetBackgroundPosition(string)
	BackgroundRepeat() string
	SetBackgroundRepeat(string)
	BackgroundSize() string
	SetBackgroundSize(string)
	BaselineShift() string
	SetBaselineShift(string)
	BlockOverflow() string
	SetBlockOverflow(string)
	BlockSize() string
	SetBlockSize(string)
	BookmarkLabel() string
	SetBookmarkLabel(string)
	BookmarkLevel() string
	SetBookmarkLevel(string)
	BookmarkState() string
	SetBookmarkState(string)
	Border() string
	SetBorder(string)
	BorderBlock() string
	SetBorderBlock(string)
	BorderBlockColor() string
	SetBorderBlockColor(string)
	BorderBlockEnd() string
	SetBorderBlockEnd(string)
	BorderBlockEndColor() string
	SetBorderBlockEndColor(string)
	BorderBlockEndStyle() string
	SetBorderBlockEndStyle(string)
	BorderBlockEndWidth() string
	SetBorderBlockEndWidth(string)
	BorderBlockStart() string
	SetBorderBlockStart(string)
	BorderBlockStartColor() string
	SetBorderBlockStartColor(string)
	BorderBlockStartStyle() string
	SetBorderBlockStartStyle(string)
	BorderBlockStartWidth() string
	SetBorderBlockStartWidth(string)
	BorderBlockStyle() string
	SetBorderBlockStyle(string)
	BorderBlockWidth() string
	SetBorderBlockWidth(string)
	BorderBottom() string
	SetBorderBottom(string)
	BorderBottomColor() string
	SetBorderBottomColor(string)
	BorderBottomLeftRadius() string
	SetBorderBottomLeftRadius(string)
	BorderBottomRightRadius() string
	SetBorderBottomRightRadius(string)
	BorderBottomStyle() string
	SetBorderBottomStyle(string)
	BorderBottomWidth() string
	SetBorderBottomWidth(string)
	BorderBoundary() string
	SetBorderBoundary(string)
	BorderCollapse() string
	SetBorderCollapse(string)
	BorderColor() string
	SetBorderColor(string)
	BorderEndEndRadius() string
	SetBorderEndEndRadius(string)
	BorderEndStartRadius() string
	SetBorderEndStartRadius(string)
	BorderImage() string
	SetBorderImage(string)
	BorderImageOutset() string
	SetBorderImageOutset(string)
	BorderImageRepeat() string
	SetBorderImageRepeat(string)
	BorderImageSlice() string
	SetBorderImageSlice(string)
	BorderImageSource() string
	SetBorderImageSource(string)
	BorderImageWidth() string
	SetBorderImageWidth(string)
	BorderInline() string
	SetBorderInline(string)
	BorderInlineColor() string
	SetBorderInlineColor(string)
	BorderInlineEnd() string
	SetBorderInlineEnd(string)
	BorderInlineEndColor() string
	SetBorderInlineEndColor(string)
	BorderInlineEndStyle() string
	SetBorderInlineEndStyle(string)
	BorderInlineEndWidth() string
	SetBorderInlineEndWidth(string)
	BorderInlineStart() string
	SetBorderInlineStart(string)
	BorderInlineStartColor() string
	SetBorderInlineStartColor(string)
	BorderInlineStartStyle() string
	SetBorderInlineStartStyle(string)
	BorderInlineStartWidth() string
	SetBorderInlineStartWidth(string)
	BorderInlineStyle() string
	SetBorderInlineStyle(string)
	BorderInlineWidth() string
	SetBorderInlineWidth(string)
	BorderLeft() string
	SetBorderLeft(string)
	BorderLeftColor() string
	SetBorderLeftColor(string)
	BorderLeftStyle() string
	SetBorderLeftStyle(string)
	BorderLeftWidth() string
	SetBorderLeftWidth(string)
	BorderRadius() string
	SetBorderRadius(string)
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
	BorderStartEndRadius() string
	SetBorderStartEndRadius(string)
	BorderStartStartRadius() string
	SetBorderStartStartRadius(string)
	BorderStyle() string
	SetBorderStyle(string)
	BorderTop() string
	SetBorderTop(string)
	BorderTopColor() string
	SetBorderTopColor(string)
	BorderTopLeftRadius() string
	SetBorderTopLeftRadius(string)
	BorderTopRightRadius() string
	SetBorderTopRightRadius(string)
	BorderTopStyle() string
	SetBorderTopStyle(string)
	BorderTopWidth() string
	SetBorderTopWidth(string)
	BorderWidth() string
	SetBorderWidth(string)
	Bottom() string
	SetBottom(string)
	BoxDecorationBreak() string
	SetBoxDecorationBreak(string)
	BoxShadow() string
	SetBoxShadow(string)
	BoxSizing() string
	SetBoxSizing(string)
	BoxSnap() string
	SetBoxSnap(string)
	BreakAfter() string
	SetBreakAfter(string)
	BreakBefore() string
	SetBreakBefore(string)
	BreakInside() string
	SetBreakInside(string)
	CaptionSide() string
	SetCaptionSide(string)
	Caret() string
	SetCaret(string)
	CaretColor() string
	SetCaretColor(string)
	CaretShape() string
	SetCaretShape(string)
	Clear() string
	SetClear(string)
	Clip() string
	SetClip(string)
	ClipPath() string
	SetClipPath(string)
	ClipRule() string
	SetClipRule(string)
	Color() string
	SetColor(string)
	ColorInterpolationFilters() string
	SetColorInterpolationFilters(string)
	ColumnCount() string
	SetColumnCount(string)
	ColumnFill() string
	SetColumnFill(string)
	ColumnGap() string
	SetColumnGap(string)
	ColumnRule() string
	SetColumnRule(string)
	ColumnRuleColor() string
	SetColumnRuleColor(string)
	ColumnRuleStyle() string
	SetColumnRuleStyle(string)
	ColumnRuleWidth() string
	SetColumnRuleWidth(string)
	ColumnSpan() string
	SetColumnSpan(string)
	ColumnWidth() string
	SetColumnWidth(string)
	Columns() string
	SetColumns(string)
	Contain() string
	SetContain(string)
	Content() string
	SetContent(string)
	Continue() string
	SetContinue(string)
	CounterIncrement() string
	SetCounterIncrement(string)
	CounterReset() string
	SetCounterReset(string)
	CounterSet() string
	SetCounterSet(string)
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
	DominantBaseline() string
	SetDominantBaseline(string)
	Elevation() string
	SetElevation(string)
	EmptyCells() string
	SetEmptyCells(string)
	Filter() string
	SetFilter(string)
	Flex() string
	SetFlex(string)
	FlexBasis() string
	SetFlexBasis(string)
	FlexDirection() string
	SetFlexDirection(string)
	FlexFlow() string
	SetFlexFlow(string)
	FlexGrow() string
	SetFlexGrow(string)
	FlexShrink() string
	SetFlexShrink(string)
	FlexWrap() string
	SetFlexWrap(string)
	Float() string
	SetFloat(string)
	FloodColor() string
	SetFloodColor(string)
	FloodOpacity() string
	SetFloodOpacity(string)
	FlowFrom() string
	SetFlowFrom(string)
	FlowInto() string
	SetFlowInto(string)
	Font() string
	SetFont(string)
	FontFamily() string
	SetFontFamily(string)
	FontFeatureSettings() string
	SetFontFeatureSettings(string)
	FontKerning() string
	SetFontKerning(string)
	FontLanguageOverride() string
	SetFontLanguageOverride(string)
	FontMaxSize() string
	SetFontMaxSize(string)
	FontMinSize() string
	SetFontMinSize(string)
	FontOpticalSizing() string
	SetFontOpticalSizing(string)
	FontPalette() string
	SetFontPalette(string)
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
	FontSynthesisSmallCaps() string
	SetFontSynthesisSmallCaps(string)
	FontSynthesisStyle() string
	SetFontSynthesisStyle(string)
	FontSynthesisWeight() string
	SetFontSynthesisWeight(string)
	FontVariant() string
	SetFontVariant(string)
	FontVariantAlternates() string
	SetFontVariantAlternates(string)
	FontVariantCaps() string
	SetFontVariantCaps(string)
	FontVariantEastAsian() string
	SetFontVariantEastAsian(string)
	FontVariantEmoji() string
	SetFontVariantEmoji(string)
	FontVariantLigatures() string
	SetFontVariantLigatures(string)
	FontVariantNumeric() string
	SetFontVariantNumeric(string)
	FontVariantPosition() string
	SetFontVariantPosition(string)
	FontVariationSettings() string
	SetFontVariationSettings(string)
	FontWeight() string
	SetFontWeight(string)
	FootnoteDisplay() string
	SetFootnoteDisplay(string)
	FootnotePolicy() string
	SetFootnotePolicy(string)
	Gap() string
	SetGap(string)
	GlyphOrientationVertical() string
	SetGlyphOrientationVertical(string)
	Grid() string
	SetGrid(string)
	GridArea() string
	SetGridArea(string)
	GridAutoColumns() string
	SetGridAutoColumns(string)
	GridAutoFlow() string
	SetGridAutoFlow(string)
	GridAutoRows() string
	SetGridAutoRows(string)
	GridColumn() string
	SetGridColumn(string)
	GridColumnEnd() string
	SetGridColumnEnd(string)
	GridColumnStart() string
	SetGridColumnStart(string)
	GridRow() string
	SetGridRow(string)
	GridRowEnd() string
	SetGridRowEnd(string)
	GridRowStart() string
	SetGridRowStart(string)
	GridTemplate() string
	SetGridTemplate(string)
	GridTemplateAreas() string
	SetGridTemplateAreas(string)
	GridTemplateColumns() string
	SetGridTemplateColumns(string)
	GridTemplateRows() string
	SetGridTemplateRows(string)
	HangingPunctuation() string
	SetHangingPunctuation(string)
	Height() string
	SetHeight(string)
	HyphenateCharacter() string
	SetHyphenateCharacter(string)
	HyphenateLimitChars() string
	SetHyphenateLimitChars(string)
	HyphenateLimitLast() string
	SetHyphenateLimitLast(string)
	HyphenateLimitLines() string
	SetHyphenateLimitLines(string)
	HyphenateLimitZone() string
	SetHyphenateLimitZone(string)
	Hyphens() string
	SetHyphens(string)
	ImageOrientation() string
	SetImageOrientation(string)
	ImageResolution() string
	SetImageResolution(string)
	InitialLetters() string
	SetInitialLetters(string)
	InitialLettersAlign() string
	SetInitialLettersAlign(string)
	InitialLettersWrap() string
	SetInitialLettersWrap(string)
	InlineSize() string
	SetInlineSize(string)
	InlineSizing() string
	SetInlineSizing(string)
	Inset() string
	SetInset(string)
	InsetBlock() string
	SetInsetBlock(string)
	InsetBlockEnd() string
	SetInsetBlockEnd(string)
	InsetBlockStart() string
	SetInsetBlockStart(string)
	InsetInline() string
	SetInsetInline(string)
	InsetInlineEnd() string
	SetInsetInlineEnd(string)
	InsetInlineStart() string
	SetInsetInlineStart(string)
	Isolation() string
	SetIsolation(string)
	JustifyContent() string
	SetJustifyContent(string)
	JustifyItems() string
	SetJustifyItems(string)
	JustifySelf() string
	SetJustifySelf(string)
	Left() string
	SetLeft(string)
	LetterSpacing() string
	SetLetterSpacing(string)
	LightingColor() string
	SetLightingColor(string)
	LineBreak() string
	SetLineBreak(string)
	LineClamp() string
	SetLineClamp(string)
	LineGrid() string
	SetLineGrid(string)
	LineHeight() string
	SetLineHeight(string)
	LinePadding() string
	SetLinePadding(string)
	LineSnap() string
	SetLineSnap(string)
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
	MarginBlock() string
	SetMarginBlock(string)
	MarginBlockEnd() string
	SetMarginBlockEnd(string)
	MarginBlockStart() string
	SetMarginBlockStart(string)
	MarginBottom() string
	SetMarginBottom(string)
	MarginInline() string
	SetMarginInline(string)
	MarginInlineEnd() string
	SetMarginInlineEnd(string)
	MarginInlineStart() string
	SetMarginInlineStart(string)
	MarginLeft() string
	SetMarginLeft(string)
	MarginRight() string
	SetMarginRight(string)
	MarginTop() string
	SetMarginTop(string)
	MarginTrim() string
	SetMarginTrim(string)
	MarkerSide() string
	SetMarkerSide(string)
	Mask() string
	SetMask(string)
	MaskBorder() string
	SetMaskBorder(string)
	MaskBorderMode() string
	SetMaskBorderMode(string)
	MaskBorderOutset() string
	SetMaskBorderOutset(string)
	MaskBorderRepeat() string
	SetMaskBorderRepeat(string)
	MaskBorderSlice() string
	SetMaskBorderSlice(string)
	MaskBorderSource() string
	SetMaskBorderSource(string)
	MaskBorderWidth() string
	SetMaskBorderWidth(string)
	MaskClip() string
	SetMaskClip(string)
	MaskComposite() string
	SetMaskComposite(string)
	MaskImage() string
	SetMaskImage(string)
	MaskMode() string
	SetMaskMode(string)
	MaskOrigin() string
	SetMaskOrigin(string)
	MaskPosition() string
	SetMaskPosition(string)
	MaskRepeat() string
	SetMaskRepeat(string)
	MaskSize() string
	SetMaskSize(string)
	MaskType() string
	SetMaskType(string)
	MaxBlockSize() string
	SetMaxBlockSize(string)
	MaxHeight() string
	SetMaxHeight(string)
	MaxInlineSize() string
	SetMaxInlineSize(string)
	MaxLines() string
	SetMaxLines(string)
	MaxWidth() string
	SetMaxWidth(string)
	MinBlockSize() string
	SetMinBlockSize(string)
	MinHeight() string
	SetMinHeight(string)
	MinInlineSize() string
	SetMinInlineSize(string)
	MinWidth() string
	SetMinWidth(string)
	MixBlendMode() string
	SetMixBlendMode(string)
	NavDown() string
	SetNavDown(string)
	NavLeft() string
	SetNavLeft(string)
	NavRight() string
	SetNavRight(string)
	NavUp() string
	SetNavUp(string)
	ObjectFit() string
	SetObjectFit(string)
	ObjectPosition() string
	SetObjectPosition(string)
	Offset() string
	SetOffset(string)
	OffsetAfter() string
	SetOffsetAfter(string)
	OffsetAnchor() string
	SetOffsetAnchor(string)
	OffsetBefore() string
	SetOffsetBefore(string)
	OffsetDistance() string
	SetOffsetDistance(string)
	OffsetEnd() string
	SetOffsetEnd(string)
	OffsetPath() string
	SetOffsetPath(string)
	OffsetPosition() string
	SetOffsetPosition(string)
	OffsetRotate() string
	SetOffsetRotate(string)
	OffsetStart() string
	SetOffsetStart(string)
	Opacity() string
	SetOpacity(string)
	Order() string
	SetOrder(string)
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
	OverflowBlock() string
	SetOverflowBlock(string)
	OverflowInline() string
	SetOverflowInline(string)
	OverflowWrap() string
	SetOverflowWrap(string)
	OverflowX() string
	SetOverflowX(string)
	OverflowY() string
	SetOverflowY(string)
	Padding() string
	SetPadding(string)
	PaddingBlock() string
	SetPaddingBlock(string)
	PaddingBlockEnd() string
	SetPaddingBlockEnd(string)
	PaddingBlockStart() string
	SetPaddingBlockStart(string)
	PaddingBottom() string
	SetPaddingBottom(string)
	PaddingInline() string
	SetPaddingInline(string)
	PaddingInlineEnd() string
	SetPaddingInlineEnd(string)
	PaddingInlineStart() string
	SetPaddingInlineStart(string)
	PaddingLeft() string
	SetPaddingLeft(string)
	PaddingRight() string
	SetPaddingRight(string)
	PaddingTop() string
	SetPaddingTop(string)
	Page() string
	SetPage(string)
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
	PlaceContent() string
	SetPlaceContent(string)
	PlaceItems() string
	SetPlaceItems(string)
	PlaceSelf() string
	SetPlaceSelf(string)
	PlayDuring() string
	SetPlayDuring(string)
	Position() string
	SetPosition(string)
	Quotes() string
	SetQuotes(string)
	RegionFragment() string
	SetRegionFragment(string)
	Resize() string
	SetResize(string)
	Richness() string
	SetRichness(string)
	Right() string
	SetRight(string)
	RowGap() string
	SetRowGap(string)
	RubyAlign() string
	SetRubyAlign(string)
	RubyMerge() string
	SetRubyMerge(string)
	RubyPosition() string
	SetRubyPosition(string)
	Running() string
	SetRunning(string)
	ScrollBehavior() string
	SetScrollBehavior(string)
	ScrollMargin() string
	SetScrollMargin(string)
	ScrollMarginBlock() string
	SetScrollMarginBlock(string)
	ScrollMarginBlockEnd() string
	SetScrollMarginBlockEnd(string)
	ScrollMarginBlockStart() string
	SetScrollMarginBlockStart(string)
	ScrollMarginBottom() string
	SetScrollMarginBottom(string)
	ScrollMarginInline() string
	SetScrollMarginInline(string)
	ScrollMarginInlineEnd() string
	SetScrollMarginInlineEnd(string)
	ScrollMarginInlineStart() string
	SetScrollMarginInlineStart(string)
	ScrollMarginLeft() string
	SetScrollMarginLeft(string)
	ScrollMarginRight() string
	SetScrollMarginRight(string)
	ScrollMarginTop() string
	SetScrollMarginTop(string)
	ScrollPadding() string
	SetScrollPadding(string)
	ScrollPaddingBlock() string
	SetScrollPaddingBlock(string)
	ScrollPaddingBlockEnd() string
	SetScrollPaddingBlockEnd(string)
	ScrollPaddingBlockStart() string
	SetScrollPaddingBlockStart(string)
	ScrollPaddingBottom() string
	SetScrollPaddingBottom(string)
	ScrollPaddingInline() string
	SetScrollPaddingInline(string)
	ScrollPaddingInlineEnd() string
	SetScrollPaddingInlineEnd(string)
	ScrollPaddingInlineStart() string
	SetScrollPaddingInlineStart(string)
	ScrollPaddingLeft() string
	SetScrollPaddingLeft(string)
	ScrollPaddingRight() string
	SetScrollPaddingRight(string)
	ScrollPaddingTop() string
	SetScrollPaddingTop(string)
	ScrollSnapAlign() string
	SetScrollSnapAlign(string)
	ScrollSnapStop() string
	SetScrollSnapStop(string)
	ScrollSnapType() string
	SetScrollSnapType(string)
	ShapeImageThreshold() string
	SetShapeImageThreshold(string)
	ShapeInside() string
	SetShapeInside(string)
	ShapeMargin() string
	SetShapeMargin(string)
	ShapeOutside() string
	SetShapeOutside(string)
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
	StringSet() string
	SetStringSet(string)
	TabSize() string
	SetTabSize(string)
	TableLayout() string
	SetTableLayout(string)
	TextAlign() string
	SetTextAlign(string)
	TextAlignAll() string
	SetTextAlignAll(string)
	TextAlignLast() string
	SetTextAlignLast(string)
	TextCombineUpright() string
	SetTextCombineUpright(string)
	TextDecoration() string
	SetTextDecoration(string)
	TextDecorationColor() string
	SetTextDecorationColor(string)
	TextDecorationLine() string
	SetTextDecorationLine(string)
	TextDecorationStyle() string
	SetTextDecorationStyle(string)
	TextEmphasis() string
	SetTextEmphasis(string)
	TextEmphasisColor() string
	SetTextEmphasisColor(string)
	TextEmphasisPosition() string
	SetTextEmphasisPosition(string)
	TextEmphasisStyle() string
	SetTextEmphasisStyle(string)
	TextGroupAlign() string
	SetTextGroupAlign(string)
	TextIndent() string
	SetTextIndent(string)
	TextJustify() string
	SetTextJustify(string)
	TextOrientation() string
	SetTextOrientation(string)
	TextOverflow() string
	SetTextOverflow(string)
	TextShadow() string
	SetTextShadow(string)
	TextSpaceCollapse() string
	SetTextSpaceCollapse(string)
	TextSpaceTrim() string
	SetTextSpaceTrim(string)
	TextSpacing() string
	SetTextSpacing(string)
	TextTransform() string
	SetTextTransform(string)
	TextUnderlinePosition() string
	SetTextUnderlinePosition(string)
	TextWrap() string
	SetTextWrap(string)
	Top() string
	SetTop(string)
	Transform() string
	SetTransform(string)
	TransformBox() string
	SetTransformBox(string)
	TransformOrigin() string
	SetTransformOrigin(string)
	Transition() string
	SetTransition(string)
	TransitionDelay() string
	SetTransitionDelay(string)
	TransitionDuration() string
	SetTransitionDuration(string)
	TransitionProperty() string
	SetTransitionProperty(string)
	TransitionTimingFunction() string
	SetTransitionTimingFunction(string)
	UnicodeBidi() string
	SetUnicodeBidi(string)
	UserSelect() string
	SetUserSelect(string)
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
	WillChange() string
	SetWillChange(string)
	WordBreak() string
	SetWordBreak(string)
	WordSpacing() string
	SetWordSpacing(string)
	WordWrap() string
	SetWordWrap(string)
	WrapAfter() string
	SetWrapAfter(string)
	WrapBefore() string
	SetWrapBefore(string)
	WrapFlow() string
	SetWrapFlow(string)
	WrapInside() string
	SetWrapInside(string)
	WrapThrough() string
	SetWrapThrough(string)
	WritingMode() string
	SetWritingMode(string)
	ZIndex() string
	SetZIndex(string)
}

type cssStyleHelperImpl struct {
	*cssStyleDeclarationImpl
}

func newCSSStyleHelperImpl(v *cssStyleDeclarationImpl) *cssStyleHelperImpl {
	if v.valid() {
		return &cssStyleHelperImpl{
			cssStyleDeclarationImpl: v,
		}
	}
	return nil
}

func (p *cssStyleHelperImpl) AlignContent() string {
	return p.PropertyValue("align-content")
}

func (p *cssStyleHelperImpl) SetAlignContent(s string) {
	p.SetProperty("align-content", s)
}

func (p *cssStyleHelperImpl) AlignItems() string {
	return p.PropertyValue("align-items")
}

func (p *cssStyleHelperImpl) SetAlignItems(s string) {
	p.SetProperty("align-items", s)
}

func (p *cssStyleHelperImpl) AlignSelf() string {
	return p.PropertyValue("align-self")
}

func (p *cssStyleHelperImpl) SetAlignSelf(s string) {
	p.SetProperty("align-self", s)
}

func (p *cssStyleHelperImpl) AlignmentBaseline() string {
	return p.PropertyValue("alignment-baseline")
}

func (p *cssStyleHelperImpl) SetAlignmentBaseline(s string) {
	p.SetProperty("alignment-baseline", s)
}

func (p *cssStyleHelperImpl) All() string {
	return p.PropertyValue("all")
}

func (p *cssStyleHelperImpl) SetAll(s string) {
	p.SetProperty("all", s)
}

func (p *cssStyleHelperImpl) Animation() string {
	return p.PropertyValue("animation")
}

func (p *cssStyleHelperImpl) SetAnimation(s string) {
	p.SetProperty("animation", s)
}

func (p *cssStyleHelperImpl) AnimationDelay() string {
	return p.PropertyValue("animation-delay")
}

func (p *cssStyleHelperImpl) SetAnimationDelay(s string) {
	p.SetProperty("animation-delay", s)
}

func (p *cssStyleHelperImpl) AnimationDirection() string {
	return p.PropertyValue("animation-direction")
}

func (p *cssStyleHelperImpl) SetAnimationDirection(s string) {
	p.SetProperty("animation-direction", s)
}

func (p *cssStyleHelperImpl) AnimationDuration() string {
	return p.PropertyValue("animation-duration")
}

func (p *cssStyleHelperImpl) SetAnimationDuration(s string) {
	p.SetProperty("animation-duration", s)
}

func (p *cssStyleHelperImpl) AnimationFillMode() string {
	return p.PropertyValue("animation-fill-mode")
}

func (p *cssStyleHelperImpl) SetAnimationFillMode(s string) {
	p.SetProperty("animation-fill-mode", s)
}

func (p *cssStyleHelperImpl) AnimationIterationCount() string {
	return p.PropertyValue("animation-iteration-count")
}

func (p *cssStyleHelperImpl) SetAnimationIterationCount(s string) {
	p.SetProperty("animation-iteration-count", s)
}

func (p *cssStyleHelperImpl) AnimationName() string {
	return p.PropertyValue("animation-name")
}

func (p *cssStyleHelperImpl) SetAnimationName(s string) {
	p.SetProperty("animation-name", s)
}

func (p *cssStyleHelperImpl) AnimationPlayState() string {
	return p.PropertyValue("animation-play-state")
}

func (p *cssStyleHelperImpl) SetAnimationPlayState(s string) {
	p.SetProperty("animation-play-state", s)
}

func (p *cssStyleHelperImpl) AnimationTimingFunction() string {
	return p.PropertyValue("animation-timing-function")
}

func (p *cssStyleHelperImpl) SetAnimationTimingFunction(s string) {
	p.SetProperty("animation-timing-function", s)
}

func (p *cssStyleHelperImpl) Appearance() string {
	return p.PropertyValue("appearance")
}

func (p *cssStyleHelperImpl) SetAppearance(s string) {
	p.SetProperty("appearance", s)
}

func (p *cssStyleHelperImpl) Azimuth() string {
	return p.PropertyValue("azimuth")
}

func (p *cssStyleHelperImpl) SetAzimuth(s string) {
	p.SetProperty("azimuth", s)
}

func (p *cssStyleHelperImpl) Background() string {
	return p.PropertyValue("background")
}

func (p *cssStyleHelperImpl) SetBackground(s string) {
	p.SetProperty("background", s)
}

func (p *cssStyleHelperImpl) BackgroundAttachment() string {
	return p.PropertyValue("background-attachment")
}

func (p *cssStyleHelperImpl) SetBackgroundAttachment(s string) {
	p.SetProperty("background-attachment", s)
}

func (p *cssStyleHelperImpl) BackgroundBlendMode() string {
	return p.PropertyValue("background-blend-mode")
}

func (p *cssStyleHelperImpl) SetBackgroundBlendMode(s string) {
	p.SetProperty("background-blend-mode", s)
}

func (p *cssStyleHelperImpl) BackgroundClip() string {
	return p.PropertyValue("background-clip")
}

func (p *cssStyleHelperImpl) SetBackgroundClip(s string) {
	p.SetProperty("background-clip", s)
}

func (p *cssStyleHelperImpl) BackgroundColor() string {
	return p.PropertyValue("background-color")
}

func (p *cssStyleHelperImpl) SetBackgroundColor(s string) {
	p.SetProperty("background-color", s)
}

func (p *cssStyleHelperImpl) BackgroundImage() string {
	return p.PropertyValue("background-image")
}

func (p *cssStyleHelperImpl) SetBackgroundImage(s string) {
	p.SetProperty("background-image", s)
}

func (p *cssStyleHelperImpl) BackgroundOrigin() string {
	return p.PropertyValue("background-origin")
}

func (p *cssStyleHelperImpl) SetBackgroundOrigin(s string) {
	p.SetProperty("background-origin", s)
}

func (p *cssStyleHelperImpl) BackgroundPosition() string {
	return p.PropertyValue("background-position")
}

func (p *cssStyleHelperImpl) SetBackgroundPosition(s string) {
	p.SetProperty("background-position", s)
}

func (p *cssStyleHelperImpl) BackgroundRepeat() string {
	return p.PropertyValue("background-repeat")
}

func (p *cssStyleHelperImpl) SetBackgroundRepeat(s string) {
	p.SetProperty("background-repeat", s)
}

func (p *cssStyleHelperImpl) BackgroundSize() string {
	return p.PropertyValue("background-size")
}

func (p *cssStyleHelperImpl) SetBackgroundSize(s string) {
	p.SetProperty("background-size", s)
}

func (p *cssStyleHelperImpl) BaselineShift() string {
	return p.PropertyValue("baseline-shift")
}

func (p *cssStyleHelperImpl) SetBaselineShift(s string) {
	p.SetProperty("baseline-shift", s)
}

func (p *cssStyleHelperImpl) BlockOverflow() string {
	return p.PropertyValue("block-overflow")
}

func (p *cssStyleHelperImpl) SetBlockOverflow(s string) {
	p.SetProperty("block-overflow", s)
}

func (p *cssStyleHelperImpl) BlockSize() string {
	return p.PropertyValue("block-size")
}

func (p *cssStyleHelperImpl) SetBlockSize(s string) {
	p.SetProperty("block-size", s)
}

func (p *cssStyleHelperImpl) BookmarkLabel() string {
	return p.PropertyValue("bookmark-label")
}

func (p *cssStyleHelperImpl) SetBookmarkLabel(s string) {
	p.SetProperty("bookmark-label", s)
}

func (p *cssStyleHelperImpl) BookmarkLevel() string {
	return p.PropertyValue("bookmark-level")
}

func (p *cssStyleHelperImpl) SetBookmarkLevel(s string) {
	p.SetProperty("bookmark-level", s)
}

func (p *cssStyleHelperImpl) BookmarkState() string {
	return p.PropertyValue("bookmark-state")
}

func (p *cssStyleHelperImpl) SetBookmarkState(s string) {
	p.SetProperty("bookmark-state", s)
}

func (p *cssStyleHelperImpl) Border() string {
	return p.PropertyValue("border")
}

func (p *cssStyleHelperImpl) SetBorder(s string) {
	p.SetProperty("border", s)
}

func (p *cssStyleHelperImpl) BorderBlock() string {
	return p.PropertyValue("border-block")
}

func (p *cssStyleHelperImpl) SetBorderBlock(s string) {
	p.SetProperty("border-block", s)
}

func (p *cssStyleHelperImpl) BorderBlockColor() string {
	return p.PropertyValue("border-block-color")
}

func (p *cssStyleHelperImpl) SetBorderBlockColor(s string) {
	p.SetProperty("border-block-color", s)
}

func (p *cssStyleHelperImpl) BorderBlockEnd() string {
	return p.PropertyValue("border-block-end")
}

func (p *cssStyleHelperImpl) SetBorderBlockEnd(s string) {
	p.SetProperty("border-block-end", s)
}

func (p *cssStyleHelperImpl) BorderBlockEndColor() string {
	return p.PropertyValue("border-block-end-color")
}

func (p *cssStyleHelperImpl) SetBorderBlockEndColor(s string) {
	p.SetProperty("border-block-end-color", s)
}

func (p *cssStyleHelperImpl) BorderBlockEndStyle() string {
	return p.PropertyValue("border-block-end-style")
}

func (p *cssStyleHelperImpl) SetBorderBlockEndStyle(s string) {
	p.SetProperty("border-block-end-style", s)
}

func (p *cssStyleHelperImpl) BorderBlockEndWidth() string {
	return p.PropertyValue("border-block-end-width")
}

func (p *cssStyleHelperImpl) SetBorderBlockEndWidth(s string) {
	p.SetProperty("border-block-end-width", s)
}

func (p *cssStyleHelperImpl) BorderBlockStart() string {
	return p.PropertyValue("border-block-start")
}

func (p *cssStyleHelperImpl) SetBorderBlockStart(s string) {
	p.SetProperty("border-block-start", s)
}

func (p *cssStyleHelperImpl) BorderBlockStartColor() string {
	return p.PropertyValue("border-block-start-color")
}

func (p *cssStyleHelperImpl) SetBorderBlockStartColor(s string) {
	p.SetProperty("border-block-start-color", s)
}

func (p *cssStyleHelperImpl) BorderBlockStartStyle() string {
	return p.PropertyValue("border-block-start-style")
}

func (p *cssStyleHelperImpl) SetBorderBlockStartStyle(s string) {
	p.SetProperty("border-block-start-style", s)
}

func (p *cssStyleHelperImpl) BorderBlockStartWidth() string {
	return p.PropertyValue("border-block-start-width")
}

func (p *cssStyleHelperImpl) SetBorderBlockStartWidth(s string) {
	p.SetProperty("border-block-start-width", s)
}

func (p *cssStyleHelperImpl) BorderBlockStyle() string {
	return p.PropertyValue("border-block-style")
}

func (p *cssStyleHelperImpl) SetBorderBlockStyle(s string) {
	p.SetProperty("border-block-style", s)
}

func (p *cssStyleHelperImpl) BorderBlockWidth() string {
	return p.PropertyValue("border-block-width")
}

func (p *cssStyleHelperImpl) SetBorderBlockWidth(s string) {
	p.SetProperty("border-block-width", s)
}

func (p *cssStyleHelperImpl) BorderBottom() string {
	return p.PropertyValue("border-bottom")
}

func (p *cssStyleHelperImpl) SetBorderBottom(s string) {
	p.SetProperty("border-bottom", s)
}

func (p *cssStyleHelperImpl) BorderBottomColor() string {
	return p.PropertyValue("border-bottom-color")
}

func (p *cssStyleHelperImpl) SetBorderBottomColor(s string) {
	p.SetProperty("border-bottom-color", s)
}

func (p *cssStyleHelperImpl) BorderBottomLeftRadius() string {
	return p.PropertyValue("border-bottom-left-radius")
}

func (p *cssStyleHelperImpl) SetBorderBottomLeftRadius(s string) {
	p.SetProperty("border-bottom-left-radius", s)
}

func (p *cssStyleHelperImpl) BorderBottomRightRadius() string {
	return p.PropertyValue("border-bottom-right-radius")
}

func (p *cssStyleHelperImpl) SetBorderBottomRightRadius(s string) {
	p.SetProperty("border-bottom-right-radius", s)
}

func (p *cssStyleHelperImpl) BorderBottomStyle() string {
	return p.PropertyValue("border-bottom-style")
}

func (p *cssStyleHelperImpl) SetBorderBottomStyle(s string) {
	p.SetProperty("border-bottom-style", s)
}

func (p *cssStyleHelperImpl) BorderBottomWidth() string {
	return p.PropertyValue("border-bottom-width")
}

func (p *cssStyleHelperImpl) SetBorderBottomWidth(s string) {
	p.SetProperty("border-bottom-width", s)
}

func (p *cssStyleHelperImpl) BorderBoundary() string {
	return p.PropertyValue("border-boundary")
}

func (p *cssStyleHelperImpl) SetBorderBoundary(s string) {
	p.SetProperty("border-boundary", s)
}

func (p *cssStyleHelperImpl) BorderCollapse() string {
	return p.PropertyValue("border-collapse")
}

func (p *cssStyleHelperImpl) SetBorderCollapse(s string) {
	p.SetProperty("border-collapse", s)
}

func (p *cssStyleHelperImpl) BorderColor() string {
	return p.PropertyValue("border-color")
}

func (p *cssStyleHelperImpl) SetBorderColor(s string) {
	p.SetProperty("border-color", s)
}

func (p *cssStyleHelperImpl) BorderEndEndRadius() string {
	return p.PropertyValue("border-end-end-radius")
}

func (p *cssStyleHelperImpl) SetBorderEndEndRadius(s string) {
	p.SetProperty("border-end-end-radius", s)
}

func (p *cssStyleHelperImpl) BorderEndStartRadius() string {
	return p.PropertyValue("border-end-start-radius")
}

func (p *cssStyleHelperImpl) SetBorderEndStartRadius(s string) {
	p.SetProperty("border-end-start-radius", s)
}

func (p *cssStyleHelperImpl) BorderImage() string {
	return p.PropertyValue("border-image")
}

func (p *cssStyleHelperImpl) SetBorderImage(s string) {
	p.SetProperty("border-image", s)
}

func (p *cssStyleHelperImpl) BorderImageOutset() string {
	return p.PropertyValue("border-image-outset")
}

func (p *cssStyleHelperImpl) SetBorderImageOutset(s string) {
	p.SetProperty("border-image-outset", s)
}

func (p *cssStyleHelperImpl) BorderImageRepeat() string {
	return p.PropertyValue("border-image-repeat")
}

func (p *cssStyleHelperImpl) SetBorderImageRepeat(s string) {
	p.SetProperty("border-image-repeat", s)
}

func (p *cssStyleHelperImpl) BorderImageSlice() string {
	return p.PropertyValue("border-image-slice")
}

func (p *cssStyleHelperImpl) SetBorderImageSlice(s string) {
	p.SetProperty("border-image-slice", s)
}

func (p *cssStyleHelperImpl) BorderImageSource() string {
	return p.PropertyValue("border-image-source")
}

func (p *cssStyleHelperImpl) SetBorderImageSource(s string) {
	p.SetProperty("border-image-source", s)
}

func (p *cssStyleHelperImpl) BorderImageWidth() string {
	return p.PropertyValue("border-image-width")
}

func (p *cssStyleHelperImpl) SetBorderImageWidth(s string) {
	p.SetProperty("border-image-width", s)
}

func (p *cssStyleHelperImpl) BorderInline() string {
	return p.PropertyValue("border-inline")
}

func (p *cssStyleHelperImpl) SetBorderInline(s string) {
	p.SetProperty("border-inline", s)
}

func (p *cssStyleHelperImpl) BorderInlineColor() string {
	return p.PropertyValue("border-inline-color")
}

func (p *cssStyleHelperImpl) SetBorderInlineColor(s string) {
	p.SetProperty("border-inline-color", s)
}

func (p *cssStyleHelperImpl) BorderInlineEnd() string {
	return p.PropertyValue("border-inline-end")
}

func (p *cssStyleHelperImpl) SetBorderInlineEnd(s string) {
	p.SetProperty("border-inline-end", s)
}

func (p *cssStyleHelperImpl) BorderInlineEndColor() string {
	return p.PropertyValue("border-inline-end-color")
}

func (p *cssStyleHelperImpl) SetBorderInlineEndColor(s string) {
	p.SetProperty("border-inline-end-color", s)
}

func (p *cssStyleHelperImpl) BorderInlineEndStyle() string {
	return p.PropertyValue("border-inline-end-style")
}

func (p *cssStyleHelperImpl) SetBorderInlineEndStyle(s string) {
	p.SetProperty("border-inline-end-style", s)
}

func (p *cssStyleHelperImpl) BorderInlineEndWidth() string {
	return p.PropertyValue("border-inline-end-width")
}

func (p *cssStyleHelperImpl) SetBorderInlineEndWidth(s string) {
	p.SetProperty("border-inline-end-width", s)
}

func (p *cssStyleHelperImpl) BorderInlineStart() string {
	return p.PropertyValue("border-inline-start")
}

func (p *cssStyleHelperImpl) SetBorderInlineStart(s string) {
	p.SetProperty("border-inline-start", s)
}

func (p *cssStyleHelperImpl) BorderInlineStartColor() string {
	return p.PropertyValue("border-inline-start-color")
}

func (p *cssStyleHelperImpl) SetBorderInlineStartColor(s string) {
	p.SetProperty("border-inline-start-color", s)
}

func (p *cssStyleHelperImpl) BorderInlineStartStyle() string {
	return p.PropertyValue("border-inline-start-style")
}

func (p *cssStyleHelperImpl) SetBorderInlineStartStyle(s string) {
	p.SetProperty("border-inline-start-style", s)
}

func (p *cssStyleHelperImpl) BorderInlineStartWidth() string {
	return p.PropertyValue("border-inline-start-width")
}

func (p *cssStyleHelperImpl) SetBorderInlineStartWidth(s string) {
	p.SetProperty("border-inline-start-width", s)
}

func (p *cssStyleHelperImpl) BorderInlineStyle() string {
	return p.PropertyValue("border-inline-style")
}

func (p *cssStyleHelperImpl) SetBorderInlineStyle(s string) {
	p.SetProperty("border-inline-style", s)
}

func (p *cssStyleHelperImpl) BorderInlineWidth() string {
	return p.PropertyValue("border-inline-width")
}

func (p *cssStyleHelperImpl) SetBorderInlineWidth(s string) {
	p.SetProperty("border-inline-width", s)
}

func (p *cssStyleHelperImpl) BorderLeft() string {
	return p.PropertyValue("border-left")
}

func (p *cssStyleHelperImpl) SetBorderLeft(s string) {
	p.SetProperty("border-left", s)
}

func (p *cssStyleHelperImpl) BorderLeftColor() string {
	return p.PropertyValue("border-left-color")
}

func (p *cssStyleHelperImpl) SetBorderLeftColor(s string) {
	p.SetProperty("border-left-color", s)
}

func (p *cssStyleHelperImpl) BorderLeftStyle() string {
	return p.PropertyValue("border-left-style")
}

func (p *cssStyleHelperImpl) SetBorderLeftStyle(s string) {
	p.SetProperty("border-left-style", s)
}

func (p *cssStyleHelperImpl) BorderLeftWidth() string {
	return p.PropertyValue("border-left-width")
}

func (p *cssStyleHelperImpl) SetBorderLeftWidth(s string) {
	p.SetProperty("border-left-width", s)
}

func (p *cssStyleHelperImpl) BorderRadius() string {
	return p.PropertyValue("border-radius")
}

func (p *cssStyleHelperImpl) SetBorderRadius(s string) {
	p.SetProperty("border-radius", s)
}

func (p *cssStyleHelperImpl) BorderRight() string {
	return p.PropertyValue("border-right")
}

func (p *cssStyleHelperImpl) SetBorderRight(s string) {
	p.SetProperty("border-right", s)
}

func (p *cssStyleHelperImpl) BorderRightColor() string {
	return p.PropertyValue("border-right-color")
}

func (p *cssStyleHelperImpl) SetBorderRightColor(s string) {
	p.SetProperty("border-right-color", s)
}

func (p *cssStyleHelperImpl) BorderRightStyle() string {
	return p.PropertyValue("border-right-style")
}

func (p *cssStyleHelperImpl) SetBorderRightStyle(s string) {
	p.SetProperty("border-right-style", s)
}

func (p *cssStyleHelperImpl) BorderRightWidth() string {
	return p.PropertyValue("border-right-width")
}

func (p *cssStyleHelperImpl) SetBorderRightWidth(s string) {
	p.SetProperty("border-right-width", s)
}

func (p *cssStyleHelperImpl) BorderSpacing() string {
	return p.PropertyValue("border-spacing")
}

func (p *cssStyleHelperImpl) SetBorderSpacing(s string) {
	p.SetProperty("border-spacing", s)
}

func (p *cssStyleHelperImpl) BorderStartEndRadius() string {
	return p.PropertyValue("border-start-end-radius")
}

func (p *cssStyleHelperImpl) SetBorderStartEndRadius(s string) {
	p.SetProperty("border-start-end-radius", s)
}

func (p *cssStyleHelperImpl) BorderStartStartRadius() string {
	return p.PropertyValue("border-start-start-radius")
}

func (p *cssStyleHelperImpl) SetBorderStartStartRadius(s string) {
	p.SetProperty("border-start-start-radius", s)
}

func (p *cssStyleHelperImpl) BorderStyle() string {
	return p.PropertyValue("border-style")
}

func (p *cssStyleHelperImpl) SetBorderStyle(s string) {
	p.SetProperty("border-style", s)
}

func (p *cssStyleHelperImpl) BorderTop() string {
	return p.PropertyValue("border-top")
}

func (p *cssStyleHelperImpl) SetBorderTop(s string) {
	p.SetProperty("border-top", s)
}

func (p *cssStyleHelperImpl) BorderTopColor() string {
	return p.PropertyValue("border-top-color")
}

func (p *cssStyleHelperImpl) SetBorderTopColor(s string) {
	p.SetProperty("border-top-color", s)
}

func (p *cssStyleHelperImpl) BorderTopLeftRadius() string {
	return p.PropertyValue("border-top-left-radius")
}

func (p *cssStyleHelperImpl) SetBorderTopLeftRadius(s string) {
	p.SetProperty("border-top-left-radius", s)
}

func (p *cssStyleHelperImpl) BorderTopRightRadius() string {
	return p.PropertyValue("border-top-right-radius")
}

func (p *cssStyleHelperImpl) SetBorderTopRightRadius(s string) {
	p.SetProperty("border-top-right-radius", s)
}

func (p *cssStyleHelperImpl) BorderTopStyle() string {
	return p.PropertyValue("border-top-style")
}

func (p *cssStyleHelperImpl) SetBorderTopStyle(s string) {
	p.SetProperty("border-top-style", s)
}

func (p *cssStyleHelperImpl) BorderTopWidth() string {
	return p.PropertyValue("border-top-width")
}

func (p *cssStyleHelperImpl) SetBorderTopWidth(s string) {
	p.SetProperty("border-top-width", s)
}

func (p *cssStyleHelperImpl) BorderWidth() string {
	return p.PropertyValue("border-width")
}

func (p *cssStyleHelperImpl) SetBorderWidth(s string) {
	p.SetProperty("border-width", s)
}

func (p *cssStyleHelperImpl) Bottom() string {
	return p.PropertyValue("bottom")
}

func (p *cssStyleHelperImpl) SetBottom(s string) {
	p.SetProperty("bottom", s)
}

func (p *cssStyleHelperImpl) BoxDecorationBreak() string {
	return p.PropertyValue("box-decoration-break")
}

func (p *cssStyleHelperImpl) SetBoxDecorationBreak(s string) {
	p.SetProperty("box-decoration-break", s)
}

func (p *cssStyleHelperImpl) BoxShadow() string {
	return p.PropertyValue("box-shadow")
}

func (p *cssStyleHelperImpl) SetBoxShadow(s string) {
	p.SetProperty("box-shadow", s)
}

func (p *cssStyleHelperImpl) BoxSizing() string {
	return p.PropertyValue("box-sizing")
}

func (p *cssStyleHelperImpl) SetBoxSizing(s string) {
	p.SetProperty("box-sizing", s)
}

func (p *cssStyleHelperImpl) BoxSnap() string {
	return p.PropertyValue("box-snap")
}

func (p *cssStyleHelperImpl) SetBoxSnap(s string) {
	p.SetProperty("box-snap", s)
}

func (p *cssStyleHelperImpl) BreakAfter() string {
	return p.PropertyValue("break-after")
}

func (p *cssStyleHelperImpl) SetBreakAfter(s string) {
	p.SetProperty("break-after", s)
}

func (p *cssStyleHelperImpl) BreakBefore() string {
	return p.PropertyValue("break-before")
}

func (p *cssStyleHelperImpl) SetBreakBefore(s string) {
	p.SetProperty("break-before", s)
}

func (p *cssStyleHelperImpl) BreakInside() string {
	return p.PropertyValue("break-inside")
}

func (p *cssStyleHelperImpl) SetBreakInside(s string) {
	p.SetProperty("break-inside", s)
}

func (p *cssStyleHelperImpl) CaptionSide() string {
	return p.PropertyValue("caption-side")
}

func (p *cssStyleHelperImpl) SetCaptionSide(s string) {
	p.SetProperty("caption-side", s)
}

func (p *cssStyleHelperImpl) Caret() string {
	return p.PropertyValue("caret")
}

func (p *cssStyleHelperImpl) SetCaret(s string) {
	p.SetProperty("caret", s)
}

func (p *cssStyleHelperImpl) CaretColor() string {
	return p.PropertyValue("caret-color")
}

func (p *cssStyleHelperImpl) SetCaretColor(s string) {
	p.SetProperty("caret-color", s)
}

func (p *cssStyleHelperImpl) CaretShape() string {
	return p.PropertyValue("caret-shape")
}

func (p *cssStyleHelperImpl) SetCaretShape(s string) {
	p.SetProperty("caret-shape", s)
}

func (p *cssStyleHelperImpl) Clear() string {
	return p.PropertyValue("clear")
}

func (p *cssStyleHelperImpl) SetClear(s string) {
	p.SetProperty("clear", s)
}

func (p *cssStyleHelperImpl) Clip() string {
	return p.PropertyValue("clip")
}

func (p *cssStyleHelperImpl) SetClip(s string) {
	p.SetProperty("clip", s)
}

func (p *cssStyleHelperImpl) ClipPath() string {
	return p.PropertyValue("clip-path")
}

func (p *cssStyleHelperImpl) SetClipPath(s string) {
	p.SetProperty("clip-path", s)
}

func (p *cssStyleHelperImpl) ClipRule() string {
	return p.PropertyValue("clip-rule")
}

func (p *cssStyleHelperImpl) SetClipRule(s string) {
	p.SetProperty("clip-rule", s)
}

func (p *cssStyleHelperImpl) Color() string {
	return p.PropertyValue("color")
}

func (p *cssStyleHelperImpl) SetColor(s string) {
	p.SetProperty("color", s)
}

func (p *cssStyleHelperImpl) ColorInterpolationFilters() string {
	return p.PropertyValue("color-interpolation-filters")
}

func (p *cssStyleHelperImpl) SetColorInterpolationFilters(s string) {
	p.SetProperty("color-interpolation-filters", s)
}

func (p *cssStyleHelperImpl) ColumnCount() string {
	return p.PropertyValue("column-count")
}

func (p *cssStyleHelperImpl) SetColumnCount(s string) {
	p.SetProperty("column-count", s)
}

func (p *cssStyleHelperImpl) ColumnFill() string {
	return p.PropertyValue("column-fill")
}

func (p *cssStyleHelperImpl) SetColumnFill(s string) {
	p.SetProperty("column-fill", s)
}

func (p *cssStyleHelperImpl) ColumnGap() string {
	return p.PropertyValue("column-gap")
}

func (p *cssStyleHelperImpl) SetColumnGap(s string) {
	p.SetProperty("column-gap", s)
}

func (p *cssStyleHelperImpl) ColumnRule() string {
	return p.PropertyValue("column-rule")
}

func (p *cssStyleHelperImpl) SetColumnRule(s string) {
	p.SetProperty("column-rule", s)
}

func (p *cssStyleHelperImpl) ColumnRuleColor() string {
	return p.PropertyValue("column-rule-color")
}

func (p *cssStyleHelperImpl) SetColumnRuleColor(s string) {
	p.SetProperty("column-rule-color", s)
}

func (p *cssStyleHelperImpl) ColumnRuleStyle() string {
	return p.PropertyValue("column-rule-style")
}

func (p *cssStyleHelperImpl) SetColumnRuleStyle(s string) {
	p.SetProperty("column-rule-style", s)
}

func (p *cssStyleHelperImpl) ColumnRuleWidth() string {
	return p.PropertyValue("column-rule-width")
}

func (p *cssStyleHelperImpl) SetColumnRuleWidth(s string) {
	p.SetProperty("column-rule-width", s)
}

func (p *cssStyleHelperImpl) ColumnSpan() string {
	return p.PropertyValue("column-span")
}

func (p *cssStyleHelperImpl) SetColumnSpan(s string) {
	p.SetProperty("column-span", s)
}

func (p *cssStyleHelperImpl) ColumnWidth() string {
	return p.PropertyValue("column-width")
}

func (p *cssStyleHelperImpl) SetColumnWidth(s string) {
	p.SetProperty("column-width", s)
}

func (p *cssStyleHelperImpl) Columns() string {
	return p.PropertyValue("columns")
}

func (p *cssStyleHelperImpl) SetColumns(s string) {
	p.SetProperty("columns", s)
}

func (p *cssStyleHelperImpl) Contain() string {
	return p.PropertyValue("contain")
}

func (p *cssStyleHelperImpl) SetContain(s string) {
	p.SetProperty("contain", s)
}

func (p *cssStyleHelperImpl) Content() string {
	return p.PropertyValue("content")
}

func (p *cssStyleHelperImpl) SetContent(s string) {
	p.SetProperty("content", s)
}

func (p *cssStyleHelperImpl) Continue() string {
	return p.PropertyValue("continue")
}

func (p *cssStyleHelperImpl) SetContinue(s string) {
	p.SetProperty("continue", s)
}

func (p *cssStyleHelperImpl) CounterIncrement() string {
	return p.PropertyValue("counter-increment")
}

func (p *cssStyleHelperImpl) SetCounterIncrement(s string) {
	p.SetProperty("counter-increment", s)
}

func (p *cssStyleHelperImpl) CounterReset() string {
	return p.PropertyValue("counter-reset")
}

func (p *cssStyleHelperImpl) SetCounterReset(s string) {
	p.SetProperty("counter-reset", s)
}

func (p *cssStyleHelperImpl) CounterSet() string {
	return p.PropertyValue("counter-set")
}

func (p *cssStyleHelperImpl) SetCounterSet(s string) {
	p.SetProperty("counter-set", s)
}

func (p *cssStyleHelperImpl) Cue() string {
	return p.PropertyValue("cue")
}

func (p *cssStyleHelperImpl) SetCue(s string) {
	p.SetProperty("cue", s)
}

func (p *cssStyleHelperImpl) CueAfter() string {
	return p.PropertyValue("cue-after")
}

func (p *cssStyleHelperImpl) SetCueAfter(s string) {
	p.SetProperty("cue-after", s)
}

func (p *cssStyleHelperImpl) CueBefore() string {
	return p.PropertyValue("cue-before")
}

func (p *cssStyleHelperImpl) SetCueBefore(s string) {
	p.SetProperty("cue-before", s)
}

func (p *cssStyleHelperImpl) Cursor() string {
	return p.PropertyValue("cursor")
}

func (p *cssStyleHelperImpl) SetCursor(s string) {
	p.SetProperty("cursor", s)
}

func (p *cssStyleHelperImpl) Direction() string {
	return p.PropertyValue("direction")
}

func (p *cssStyleHelperImpl) SetDirection(s string) {
	p.SetProperty("direction", s)
}

func (p *cssStyleHelperImpl) Display() string {
	return p.PropertyValue("display")
}

func (p *cssStyleHelperImpl) SetDisplay(s string) {
	p.SetProperty("display", s)
}

func (p *cssStyleHelperImpl) DominantBaseline() string {
	return p.PropertyValue("dominant-baseline")
}

func (p *cssStyleHelperImpl) SetDominantBaseline(s string) {
	p.SetProperty("dominant-baseline", s)
}

func (p *cssStyleHelperImpl) Elevation() string {
	return p.PropertyValue("elevation")
}

func (p *cssStyleHelperImpl) SetElevation(s string) {
	p.SetProperty("elevation", s)
}

func (p *cssStyleHelperImpl) EmptyCells() string {
	return p.PropertyValue("empty-cells")
}

func (p *cssStyleHelperImpl) SetEmptyCells(s string) {
	p.SetProperty("empty-cells", s)
}

func (p *cssStyleHelperImpl) Filter() string {
	return p.PropertyValue("filter")
}

func (p *cssStyleHelperImpl) SetFilter(s string) {
	p.SetProperty("filter", s)
}

func (p *cssStyleHelperImpl) Flex() string {
	return p.PropertyValue("flex")
}

func (p *cssStyleHelperImpl) SetFlex(s string) {
	p.SetProperty("flex", s)
}

func (p *cssStyleHelperImpl) FlexBasis() string {
	return p.PropertyValue("flex-basis")
}

func (p *cssStyleHelperImpl) SetFlexBasis(s string) {
	p.SetProperty("flex-basis", s)
}

func (p *cssStyleHelperImpl) FlexDirection() string {
	return p.PropertyValue("flex-direction")
}

func (p *cssStyleHelperImpl) SetFlexDirection(s string) {
	p.SetProperty("flex-direction", s)
}

func (p *cssStyleHelperImpl) FlexFlow() string {
	return p.PropertyValue("flex-flow")
}

func (p *cssStyleHelperImpl) SetFlexFlow(s string) {
	p.SetProperty("flex-flow", s)
}

func (p *cssStyleHelperImpl) FlexGrow() string {
	return p.PropertyValue("flex-grow")
}

func (p *cssStyleHelperImpl) SetFlexGrow(s string) {
	p.SetProperty("flex-grow", s)
}

func (p *cssStyleHelperImpl) FlexShrink() string {
	return p.PropertyValue("flex-shrink")
}

func (p *cssStyleHelperImpl) SetFlexShrink(s string) {
	p.SetProperty("flex-shrink", s)
}

func (p *cssStyleHelperImpl) FlexWrap() string {
	return p.PropertyValue("flex-wrap")
}

func (p *cssStyleHelperImpl) SetFlexWrap(s string) {
	p.SetProperty("flex-wrap", s)
}

func (p *cssStyleHelperImpl) Float() string {
	return p.PropertyValue("float")
}

func (p *cssStyleHelperImpl) SetFloat(s string) {
	p.SetProperty("float", s)
}

func (p *cssStyleHelperImpl) FloodColor() string {
	return p.PropertyValue("flood-color")
}

func (p *cssStyleHelperImpl) SetFloodColor(s string) {
	p.SetProperty("flood-color", s)
}

func (p *cssStyleHelperImpl) FloodOpacity() string {
	return p.PropertyValue("flood-opacity")
}

func (p *cssStyleHelperImpl) SetFloodOpacity(s string) {
	p.SetProperty("flood-opacity", s)
}

func (p *cssStyleHelperImpl) FlowFrom() string {
	return p.PropertyValue("flow-from")
}

func (p *cssStyleHelperImpl) SetFlowFrom(s string) {
	p.SetProperty("flow-from", s)
}

func (p *cssStyleHelperImpl) FlowInto() string {
	return p.PropertyValue("flow-into")
}

func (p *cssStyleHelperImpl) SetFlowInto(s string) {
	p.SetProperty("flow-into", s)
}

func (p *cssStyleHelperImpl) Font() string {
	return p.PropertyValue("font")
}

func (p *cssStyleHelperImpl) SetFont(s string) {
	p.SetProperty("font", s)
}

func (p *cssStyleHelperImpl) FontFamily() string {
	return p.PropertyValue("font-family")
}

func (p *cssStyleHelperImpl) SetFontFamily(s string) {
	p.SetProperty("font-family", s)
}

func (p *cssStyleHelperImpl) FontFeatureSettings() string {
	return p.PropertyValue("font-feature-settings")
}

func (p *cssStyleHelperImpl) SetFontFeatureSettings(s string) {
	p.SetProperty("font-feature-settings", s)
}

func (p *cssStyleHelperImpl) FontKerning() string {
	return p.PropertyValue("font-kerning")
}

func (p *cssStyleHelperImpl) SetFontKerning(s string) {
	p.SetProperty("font-kerning", s)
}

func (p *cssStyleHelperImpl) FontLanguageOverride() string {
	return p.PropertyValue("font-language-override")
}

func (p *cssStyleHelperImpl) SetFontLanguageOverride(s string) {
	p.SetProperty("font-language-override", s)
}

func (p *cssStyleHelperImpl) FontMaxSize() string {
	return p.PropertyValue("font-max-size")
}

func (p *cssStyleHelperImpl) SetFontMaxSize(s string) {
	p.SetProperty("font-max-size", s)
}

func (p *cssStyleHelperImpl) FontMinSize() string {
	return p.PropertyValue("font-min-size")
}

func (p *cssStyleHelperImpl) SetFontMinSize(s string) {
	p.SetProperty("font-min-size", s)
}

func (p *cssStyleHelperImpl) FontOpticalSizing() string {
	return p.PropertyValue("font-optical-sizing")
}

func (p *cssStyleHelperImpl) SetFontOpticalSizing(s string) {
	p.SetProperty("font-optical-sizing", s)
}

func (p *cssStyleHelperImpl) FontPalette() string {
	return p.PropertyValue("font-palette")
}

func (p *cssStyleHelperImpl) SetFontPalette(s string) {
	p.SetProperty("font-palette", s)
}

func (p *cssStyleHelperImpl) FontSize() string {
	return p.PropertyValue("font-size")
}

func (p *cssStyleHelperImpl) SetFontSize(s string) {
	p.SetProperty("font-size", s)
}

func (p *cssStyleHelperImpl) FontSizeAdjust() string {
	return p.PropertyValue("font-size-adjust")
}

func (p *cssStyleHelperImpl) SetFontSizeAdjust(s string) {
	p.SetProperty("font-size-adjust", s)
}

func (p *cssStyleHelperImpl) FontStretch() string {
	return p.PropertyValue("font-stretch")
}

func (p *cssStyleHelperImpl) SetFontStretch(s string) {
	p.SetProperty("font-stretch", s)
}

func (p *cssStyleHelperImpl) FontStyle() string {
	return p.PropertyValue("font-style")
}

func (p *cssStyleHelperImpl) SetFontStyle(s string) {
	p.SetProperty("font-style", s)
}

func (p *cssStyleHelperImpl) FontSynthesis() string {
	return p.PropertyValue("font-synthesis")
}

func (p *cssStyleHelperImpl) SetFontSynthesis(s string) {
	p.SetProperty("font-synthesis", s)
}

func (p *cssStyleHelperImpl) FontSynthesisSmallCaps() string {
	return p.PropertyValue("font-synthesis-small-caps")
}

func (p *cssStyleHelperImpl) SetFontSynthesisSmallCaps(s string) {
	p.SetProperty("font-synthesis-small-caps", s)
}

func (p *cssStyleHelperImpl) FontSynthesisStyle() string {
	return p.PropertyValue("font-synthesis-style")
}

func (p *cssStyleHelperImpl) SetFontSynthesisStyle(s string) {
	p.SetProperty("font-synthesis-style", s)
}

func (p *cssStyleHelperImpl) FontSynthesisWeight() string {
	return p.PropertyValue("font-synthesis-weight")
}

func (p *cssStyleHelperImpl) SetFontSynthesisWeight(s string) {
	p.SetProperty("font-synthesis-weight", s)
}

func (p *cssStyleHelperImpl) FontVariant() string {
	return p.PropertyValue("font-variant")
}

func (p *cssStyleHelperImpl) SetFontVariant(s string) {
	p.SetProperty("font-variant", s)
}

func (p *cssStyleHelperImpl) FontVariantAlternates() string {
	return p.PropertyValue("font-variant-alternates")
}

func (p *cssStyleHelperImpl) SetFontVariantAlternates(s string) {
	p.SetProperty("font-variant-alternates", s)
}

func (p *cssStyleHelperImpl) FontVariantCaps() string {
	return p.PropertyValue("font-variant-caps")
}

func (p *cssStyleHelperImpl) SetFontVariantCaps(s string) {
	p.SetProperty("font-variant-caps", s)
}

func (p *cssStyleHelperImpl) FontVariantEastAsian() string {
	return p.PropertyValue("font-variant-east-asian")
}

func (p *cssStyleHelperImpl) SetFontVariantEastAsian(s string) {
	p.SetProperty("font-variant-east-asian", s)
}

func (p *cssStyleHelperImpl) FontVariantEmoji() string {
	return p.PropertyValue("font-variant-emoji")
}

func (p *cssStyleHelperImpl) SetFontVariantEmoji(s string) {
	p.SetProperty("font-variant-emoji", s)
}

func (p *cssStyleHelperImpl) FontVariantLigatures() string {
	return p.PropertyValue("font-variant-ligatures")
}

func (p *cssStyleHelperImpl) SetFontVariantLigatures(s string) {
	p.SetProperty("font-variant-ligatures", s)
}

func (p *cssStyleHelperImpl) FontVariantNumeric() string {
	return p.PropertyValue("font-variant-numeric")
}

func (p *cssStyleHelperImpl) SetFontVariantNumeric(s string) {
	p.SetProperty("font-variant-numeric", s)
}

func (p *cssStyleHelperImpl) FontVariantPosition() string {
	return p.PropertyValue("font-variant-position")
}

func (p *cssStyleHelperImpl) SetFontVariantPosition(s string) {
	p.SetProperty("font-variant-position", s)
}

func (p *cssStyleHelperImpl) FontVariationSettings() string {
	return p.PropertyValue("font-variation-settings")
}

func (p *cssStyleHelperImpl) SetFontVariationSettings(s string) {
	p.SetProperty("font-variation-settings", s)
}

func (p *cssStyleHelperImpl) FontWeight() string {
	return p.PropertyValue("font-weight")
}

func (p *cssStyleHelperImpl) SetFontWeight(s string) {
	p.SetProperty("font-weight", s)
}

func (p *cssStyleHelperImpl) FootnoteDisplay() string {
	return p.PropertyValue("footnote-display")
}

func (p *cssStyleHelperImpl) SetFootnoteDisplay(s string) {
	p.SetProperty("footnote-display", s)
}

func (p *cssStyleHelperImpl) FootnotePolicy() string {
	return p.PropertyValue("footnote-policy")
}

func (p *cssStyleHelperImpl) SetFootnotePolicy(s string) {
	p.SetProperty("footnote-policy", s)
}

func (p *cssStyleHelperImpl) Gap() string {
	return p.PropertyValue("gap")
}

func (p *cssStyleHelperImpl) SetGap(s string) {
	p.SetProperty("gap", s)
}

func (p *cssStyleHelperImpl) GlyphOrientationVertical() string {
	return p.PropertyValue("glyph-orientation-vertical")
}

func (p *cssStyleHelperImpl) SetGlyphOrientationVertical(s string) {
	p.SetProperty("glyph-orientation-vertical", s)
}

func (p *cssStyleHelperImpl) Grid() string {
	return p.PropertyValue("grid")
}

func (p *cssStyleHelperImpl) SetGrid(s string) {
	p.SetProperty("grid", s)
}

func (p *cssStyleHelperImpl) GridArea() string {
	return p.PropertyValue("grid-area")
}

func (p *cssStyleHelperImpl) SetGridArea(s string) {
	p.SetProperty("grid-area", s)
}

func (p *cssStyleHelperImpl) GridAutoColumns() string {
	return p.PropertyValue("grid-auto-columns")
}

func (p *cssStyleHelperImpl) SetGridAutoColumns(s string) {
	p.SetProperty("grid-auto-columns", s)
}

func (p *cssStyleHelperImpl) GridAutoFlow() string {
	return p.PropertyValue("grid-auto-flow")
}

func (p *cssStyleHelperImpl) SetGridAutoFlow(s string) {
	p.SetProperty("grid-auto-flow", s)
}

func (p *cssStyleHelperImpl) GridAutoRows() string {
	return p.PropertyValue("grid-auto-rows")
}

func (p *cssStyleHelperImpl) SetGridAutoRows(s string) {
	p.SetProperty("grid-auto-rows", s)
}

func (p *cssStyleHelperImpl) GridColumn() string {
	return p.PropertyValue("grid-column")
}

func (p *cssStyleHelperImpl) SetGridColumn(s string) {
	p.SetProperty("grid-column", s)
}

func (p *cssStyleHelperImpl) GridColumnEnd() string {
	return p.PropertyValue("grid-column-end")
}

func (p *cssStyleHelperImpl) SetGridColumnEnd(s string) {
	p.SetProperty("grid-column-end", s)
}

func (p *cssStyleHelperImpl) GridColumnStart() string {
	return p.PropertyValue("grid-column-start")
}

func (p *cssStyleHelperImpl) SetGridColumnStart(s string) {
	p.SetProperty("grid-column-start", s)
}

func (p *cssStyleHelperImpl) GridRow() string {
	return p.PropertyValue("grid-row")
}

func (p *cssStyleHelperImpl) SetGridRow(s string) {
	p.SetProperty("grid-row", s)
}

func (p *cssStyleHelperImpl) GridRowEnd() string {
	return p.PropertyValue("grid-row-end")
}

func (p *cssStyleHelperImpl) SetGridRowEnd(s string) {
	p.SetProperty("grid-row-end", s)
}

func (p *cssStyleHelperImpl) GridRowStart() string {
	return p.PropertyValue("grid-row-start")
}

func (p *cssStyleHelperImpl) SetGridRowStart(s string) {
	p.SetProperty("grid-row-start", s)
}

func (p *cssStyleHelperImpl) GridTemplate() string {
	return p.PropertyValue("grid-template")
}

func (p *cssStyleHelperImpl) SetGridTemplate(s string) {
	p.SetProperty("grid-template", s)
}

func (p *cssStyleHelperImpl) GridTemplateAreas() string {
	return p.PropertyValue("grid-template-areas")
}

func (p *cssStyleHelperImpl) SetGridTemplateAreas(s string) {
	p.SetProperty("grid-template-areas", s)
}

func (p *cssStyleHelperImpl) GridTemplateColumns() string {
	return p.PropertyValue("grid-template-columns")
}

func (p *cssStyleHelperImpl) SetGridTemplateColumns(s string) {
	p.SetProperty("grid-template-columns", s)
}

func (p *cssStyleHelperImpl) GridTemplateRows() string {
	return p.PropertyValue("grid-template-rows")
}

func (p *cssStyleHelperImpl) SetGridTemplateRows(s string) {
	p.SetProperty("grid-template-rows", s)
}

func (p *cssStyleHelperImpl) HangingPunctuation() string {
	return p.PropertyValue("hanging-punctuation")
}

func (p *cssStyleHelperImpl) SetHangingPunctuation(s string) {
	p.SetProperty("hanging-punctuation", s)
}

func (p *cssStyleHelperImpl) Height() string {
	return p.PropertyValue("height")
}

func (p *cssStyleHelperImpl) SetHeight(s string) {
	p.SetProperty("height", s)
}

func (p *cssStyleHelperImpl) HyphenateCharacter() string {
	return p.PropertyValue("hyphenate-character")
}

func (p *cssStyleHelperImpl) SetHyphenateCharacter(s string) {
	p.SetProperty("hyphenate-character", s)
}

func (p *cssStyleHelperImpl) HyphenateLimitChars() string {
	return p.PropertyValue("hyphenate-limit-chars")
}

func (p *cssStyleHelperImpl) SetHyphenateLimitChars(s string) {
	p.SetProperty("hyphenate-limit-chars", s)
}

func (p *cssStyleHelperImpl) HyphenateLimitLast() string {
	return p.PropertyValue("hyphenate-limit-last")
}

func (p *cssStyleHelperImpl) SetHyphenateLimitLast(s string) {
	p.SetProperty("hyphenate-limit-last", s)
}

func (p *cssStyleHelperImpl) HyphenateLimitLines() string {
	return p.PropertyValue("hyphenate-limit-lines")
}

func (p *cssStyleHelperImpl) SetHyphenateLimitLines(s string) {
	p.SetProperty("hyphenate-limit-lines", s)
}

func (p *cssStyleHelperImpl) HyphenateLimitZone() string {
	return p.PropertyValue("hyphenate-limit-zone")
}

func (p *cssStyleHelperImpl) SetHyphenateLimitZone(s string) {
	p.SetProperty("hyphenate-limit-zone", s)
}

func (p *cssStyleHelperImpl) Hyphens() string {
	return p.PropertyValue("hyphens")
}

func (p *cssStyleHelperImpl) SetHyphens(s string) {
	p.SetProperty("hyphens", s)
}

func (p *cssStyleHelperImpl) ImageOrientation() string {
	return p.PropertyValue("image-orientation")
}

func (p *cssStyleHelperImpl) SetImageOrientation(s string) {
	p.SetProperty("image-orientation", s)
}

func (p *cssStyleHelperImpl) ImageResolution() string {
	return p.PropertyValue("image-resolution")
}

func (p *cssStyleHelperImpl) SetImageResolution(s string) {
	p.SetProperty("image-resolution", s)
}

func (p *cssStyleHelperImpl) InitialLetters() string {
	return p.PropertyValue("initial-letters")
}

func (p *cssStyleHelperImpl) SetInitialLetters(s string) {
	p.SetProperty("initial-letters", s)
}

func (p *cssStyleHelperImpl) InitialLettersAlign() string {
	return p.PropertyValue("initial-letters-align")
}

func (p *cssStyleHelperImpl) SetInitialLettersAlign(s string) {
	p.SetProperty("initial-letters-align", s)
}

func (p *cssStyleHelperImpl) InitialLettersWrap() string {
	return p.PropertyValue("initial-letters-wrap")
}

func (p *cssStyleHelperImpl) SetInitialLettersWrap(s string) {
	p.SetProperty("initial-letters-wrap", s)
}

func (p *cssStyleHelperImpl) InlineSize() string {
	return p.PropertyValue("inline-size")
}

func (p *cssStyleHelperImpl) SetInlineSize(s string) {
	p.SetProperty("inline-size", s)
}

func (p *cssStyleHelperImpl) InlineSizing() string {
	return p.PropertyValue("inline-sizing")
}

func (p *cssStyleHelperImpl) SetInlineSizing(s string) {
	p.SetProperty("inline-sizing", s)
}

func (p *cssStyleHelperImpl) Inset() string {
	return p.PropertyValue("inset")
}

func (p *cssStyleHelperImpl) SetInset(s string) {
	p.SetProperty("inset", s)
}

func (p *cssStyleHelperImpl) InsetBlock() string {
	return p.PropertyValue("inset-block")
}

func (p *cssStyleHelperImpl) SetInsetBlock(s string) {
	p.SetProperty("inset-block", s)
}

func (p *cssStyleHelperImpl) InsetBlockEnd() string {
	return p.PropertyValue("inset-block-end")
}

func (p *cssStyleHelperImpl) SetInsetBlockEnd(s string) {
	p.SetProperty("inset-block-end", s)
}

func (p *cssStyleHelperImpl) InsetBlockStart() string {
	return p.PropertyValue("inset-block-start")
}

func (p *cssStyleHelperImpl) SetInsetBlockStart(s string) {
	p.SetProperty("inset-block-start", s)
}

func (p *cssStyleHelperImpl) InsetInline() string {
	return p.PropertyValue("inset-inline")
}

func (p *cssStyleHelperImpl) SetInsetInline(s string) {
	p.SetProperty("inset-inline", s)
}

func (p *cssStyleHelperImpl) InsetInlineEnd() string {
	return p.PropertyValue("inset-inline-end")
}

func (p *cssStyleHelperImpl) SetInsetInlineEnd(s string) {
	p.SetProperty("inset-inline-end", s)
}

func (p *cssStyleHelperImpl) InsetInlineStart() string {
	return p.PropertyValue("inset-inline-start")
}

func (p *cssStyleHelperImpl) SetInsetInlineStart(s string) {
	p.SetProperty("inset-inline-start", s)
}

func (p *cssStyleHelperImpl) Isolation() string {
	return p.PropertyValue("isolation")
}

func (p *cssStyleHelperImpl) SetIsolation(s string) {
	p.SetProperty("isolation", s)
}

func (p *cssStyleHelperImpl) JustifyContent() string {
	return p.PropertyValue("justify-content")
}

func (p *cssStyleHelperImpl) SetJustifyContent(s string) {
	p.SetProperty("justify-content", s)
}

func (p *cssStyleHelperImpl) JustifyItems() string {
	return p.PropertyValue("justify-items")
}

func (p *cssStyleHelperImpl) SetJustifyItems(s string) {
	p.SetProperty("justify-items", s)
}

func (p *cssStyleHelperImpl) JustifySelf() string {
	return p.PropertyValue("justify-self")
}

func (p *cssStyleHelperImpl) SetJustifySelf(s string) {
	p.SetProperty("justify-self", s)
}

func (p *cssStyleHelperImpl) Left() string {
	return p.PropertyValue("left")
}

func (p *cssStyleHelperImpl) SetLeft(s string) {
	p.SetProperty("left", s)
}

func (p *cssStyleHelperImpl) LetterSpacing() string {
	return p.PropertyValue("letter-spacing")
}

func (p *cssStyleHelperImpl) SetLetterSpacing(s string) {
	p.SetProperty("letter-spacing", s)
}

func (p *cssStyleHelperImpl) LightingColor() string {
	return p.PropertyValue("lighting-color")
}

func (p *cssStyleHelperImpl) SetLightingColor(s string) {
	p.SetProperty("lighting-color", s)
}

func (p *cssStyleHelperImpl) LineBreak() string {
	return p.PropertyValue("line-break")
}

func (p *cssStyleHelperImpl) SetLineBreak(s string) {
	p.SetProperty("line-break", s)
}

func (p *cssStyleHelperImpl) LineClamp() string {
	return p.PropertyValue("line-clamp")
}

func (p *cssStyleHelperImpl) SetLineClamp(s string) {
	p.SetProperty("line-clamp", s)
}

func (p *cssStyleHelperImpl) LineGrid() string {
	return p.PropertyValue("line-grid")
}

func (p *cssStyleHelperImpl) SetLineGrid(s string) {
	p.SetProperty("line-grid", s)
}

func (p *cssStyleHelperImpl) LineHeight() string {
	return p.PropertyValue("line-height")
}

func (p *cssStyleHelperImpl) SetLineHeight(s string) {
	p.SetProperty("line-height", s)
}

func (p *cssStyleHelperImpl) LinePadding() string {
	return p.PropertyValue("line-padding")
}

func (p *cssStyleHelperImpl) SetLinePadding(s string) {
	p.SetProperty("line-padding", s)
}

func (p *cssStyleHelperImpl) LineSnap() string {
	return p.PropertyValue("line-snap")
}

func (p *cssStyleHelperImpl) SetLineSnap(s string) {
	p.SetProperty("line-snap", s)
}

func (p *cssStyleHelperImpl) ListStyle() string {
	return p.PropertyValue("list-style")
}

func (p *cssStyleHelperImpl) SetListStyle(s string) {
	p.SetProperty("list-style", s)
}

func (p *cssStyleHelperImpl) ListStyleImage() string {
	return p.PropertyValue("list-style-image")
}

func (p *cssStyleHelperImpl) SetListStyleImage(s string) {
	p.SetProperty("list-style-image", s)
}

func (p *cssStyleHelperImpl) ListStylePosition() string {
	return p.PropertyValue("list-style-position")
}

func (p *cssStyleHelperImpl) SetListStylePosition(s string) {
	p.SetProperty("list-style-position", s)
}

func (p *cssStyleHelperImpl) ListStyleType() string {
	return p.PropertyValue("list-style-type")
}

func (p *cssStyleHelperImpl) SetListStyleType(s string) {
	p.SetProperty("list-style-type", s)
}

func (p *cssStyleHelperImpl) Margin() string {
	return p.PropertyValue("margin")
}

func (p *cssStyleHelperImpl) SetMargin(s string) {
	p.SetProperty("margin", s)
}

func (p *cssStyleHelperImpl) MarginBlock() string {
	return p.PropertyValue("margin-block")
}

func (p *cssStyleHelperImpl) SetMarginBlock(s string) {
	p.SetProperty("margin-block", s)
}

func (p *cssStyleHelperImpl) MarginBlockEnd() string {
	return p.PropertyValue("margin-block-end")
}

func (p *cssStyleHelperImpl) SetMarginBlockEnd(s string) {
	p.SetProperty("margin-block-end", s)
}

func (p *cssStyleHelperImpl) MarginBlockStart() string {
	return p.PropertyValue("margin-block-start")
}

func (p *cssStyleHelperImpl) SetMarginBlockStart(s string) {
	p.SetProperty("margin-block-start", s)
}

func (p *cssStyleHelperImpl) MarginBottom() string {
	return p.PropertyValue("margin-bottom")
}

func (p *cssStyleHelperImpl) SetMarginBottom(s string) {
	p.SetProperty("margin-bottom", s)
}

func (p *cssStyleHelperImpl) MarginInline() string {
	return p.PropertyValue("margin-inline")
}

func (p *cssStyleHelperImpl) SetMarginInline(s string) {
	p.SetProperty("margin-inline", s)
}

func (p *cssStyleHelperImpl) MarginInlineEnd() string {
	return p.PropertyValue("margin-inline-end")
}

func (p *cssStyleHelperImpl) SetMarginInlineEnd(s string) {
	p.SetProperty("margin-inline-end", s)
}

func (p *cssStyleHelperImpl) MarginInlineStart() string {
	return p.PropertyValue("margin-inline-start")
}

func (p *cssStyleHelperImpl) SetMarginInlineStart(s string) {
	p.SetProperty("margin-inline-start", s)
}

func (p *cssStyleHelperImpl) MarginLeft() string {
	return p.PropertyValue("margin-left")
}

func (p *cssStyleHelperImpl) SetMarginLeft(s string) {
	p.SetProperty("margin-left", s)
}

func (p *cssStyleHelperImpl) MarginRight() string {
	return p.PropertyValue("margin-right")
}

func (p *cssStyleHelperImpl) SetMarginRight(s string) {
	p.SetProperty("margin-right", s)
}

func (p *cssStyleHelperImpl) MarginTop() string {
	return p.PropertyValue("margin-top")
}

func (p *cssStyleHelperImpl) SetMarginTop(s string) {
	p.SetProperty("margin-top", s)
}

func (p *cssStyleHelperImpl) MarginTrim() string {
	return p.PropertyValue("margin-trim")
}

func (p *cssStyleHelperImpl) SetMarginTrim(s string) {
	p.SetProperty("margin-trim", s)
}

func (p *cssStyleHelperImpl) MarkerSide() string {
	return p.PropertyValue("marker-side")
}

func (p *cssStyleHelperImpl) SetMarkerSide(s string) {
	p.SetProperty("marker-side", s)
}

func (p *cssStyleHelperImpl) Mask() string {
	return p.PropertyValue("mask")
}

func (p *cssStyleHelperImpl) SetMask(s string) {
	p.SetProperty("mask", s)
}

func (p *cssStyleHelperImpl) MaskBorder() string {
	return p.PropertyValue("mask-border")
}

func (p *cssStyleHelperImpl) SetMaskBorder(s string) {
	p.SetProperty("mask-border", s)
}

func (p *cssStyleHelperImpl) MaskBorderMode() string {
	return p.PropertyValue("mask-border-mode")
}

func (p *cssStyleHelperImpl) SetMaskBorderMode(s string) {
	p.SetProperty("mask-border-mode", s)
}

func (p *cssStyleHelperImpl) MaskBorderOutset() string {
	return p.PropertyValue("mask-border-outset")
}

func (p *cssStyleHelperImpl) SetMaskBorderOutset(s string) {
	p.SetProperty("mask-border-outset", s)
}

func (p *cssStyleHelperImpl) MaskBorderRepeat() string {
	return p.PropertyValue("mask-border-repeat")
}

func (p *cssStyleHelperImpl) SetMaskBorderRepeat(s string) {
	p.SetProperty("mask-border-repeat", s)
}

func (p *cssStyleHelperImpl) MaskBorderSlice() string {
	return p.PropertyValue("mask-border-slice")
}

func (p *cssStyleHelperImpl) SetMaskBorderSlice(s string) {
	p.SetProperty("mask-border-slice", s)
}

func (p *cssStyleHelperImpl) MaskBorderSource() string {
	return p.PropertyValue("mask-border-source")
}

func (p *cssStyleHelperImpl) SetMaskBorderSource(s string) {
	p.SetProperty("mask-border-source", s)
}

func (p *cssStyleHelperImpl) MaskBorderWidth() string {
	return p.PropertyValue("mask-border-width")
}

func (p *cssStyleHelperImpl) SetMaskBorderWidth(s string) {
	p.SetProperty("mask-border-width", s)
}

func (p *cssStyleHelperImpl) MaskClip() string {
	return p.PropertyValue("mask-clip")
}

func (p *cssStyleHelperImpl) SetMaskClip(s string) {
	p.SetProperty("mask-clip", s)
}

func (p *cssStyleHelperImpl) MaskComposite() string {
	return p.PropertyValue("mask-composite")
}

func (p *cssStyleHelperImpl) SetMaskComposite(s string) {
	p.SetProperty("mask-composite", s)
}

func (p *cssStyleHelperImpl) MaskImage() string {
	return p.PropertyValue("mask-image")
}

func (p *cssStyleHelperImpl) SetMaskImage(s string) {
	p.SetProperty("mask-image", s)
}

func (p *cssStyleHelperImpl) MaskMode() string {
	return p.PropertyValue("mask-mode")
}

func (p *cssStyleHelperImpl) SetMaskMode(s string) {
	p.SetProperty("mask-mode", s)
}

func (p *cssStyleHelperImpl) MaskOrigin() string {
	return p.PropertyValue("mask-origin")
}

func (p *cssStyleHelperImpl) SetMaskOrigin(s string) {
	p.SetProperty("mask-origin", s)
}

func (p *cssStyleHelperImpl) MaskPosition() string {
	return p.PropertyValue("mask-position")
}

func (p *cssStyleHelperImpl) SetMaskPosition(s string) {
	p.SetProperty("mask-position", s)
}

func (p *cssStyleHelperImpl) MaskRepeat() string {
	return p.PropertyValue("mask-repeat")
}

func (p *cssStyleHelperImpl) SetMaskRepeat(s string) {
	p.SetProperty("mask-repeat", s)
}

func (p *cssStyleHelperImpl) MaskSize() string {
	return p.PropertyValue("mask-size")
}

func (p *cssStyleHelperImpl) SetMaskSize(s string) {
	p.SetProperty("mask-size", s)
}

func (p *cssStyleHelperImpl) MaskType() string {
	return p.PropertyValue("mask-type")
}

func (p *cssStyleHelperImpl) SetMaskType(s string) {
	p.SetProperty("mask-type", s)
}

func (p *cssStyleHelperImpl) MaxBlockSize() string {
	return p.PropertyValue("max-block-size")
}

func (p *cssStyleHelperImpl) SetMaxBlockSize(s string) {
	p.SetProperty("max-block-size", s)
}

func (p *cssStyleHelperImpl) MaxHeight() string {
	return p.PropertyValue("max-height")
}

func (p *cssStyleHelperImpl) SetMaxHeight(s string) {
	p.SetProperty("max-height", s)
}

func (p *cssStyleHelperImpl) MaxInlineSize() string {
	return p.PropertyValue("max-inline-size")
}

func (p *cssStyleHelperImpl) SetMaxInlineSize(s string) {
	p.SetProperty("max-inline-size", s)
}

func (p *cssStyleHelperImpl) MaxLines() string {
	return p.PropertyValue("max-lines")
}

func (p *cssStyleHelperImpl) SetMaxLines(s string) {
	p.SetProperty("max-lines", s)
}

func (p *cssStyleHelperImpl) MaxWidth() string {
	return p.PropertyValue("max-width")
}

func (p *cssStyleHelperImpl) SetMaxWidth(s string) {
	p.SetProperty("max-width", s)
}

func (p *cssStyleHelperImpl) MinBlockSize() string {
	return p.PropertyValue("min-block-size")
}

func (p *cssStyleHelperImpl) SetMinBlockSize(s string) {
	p.SetProperty("min-block-size", s)
}

func (p *cssStyleHelperImpl) MinHeight() string {
	return p.PropertyValue("min-height")
}

func (p *cssStyleHelperImpl) SetMinHeight(s string) {
	p.SetProperty("min-height", s)
}

func (p *cssStyleHelperImpl) MinInlineSize() string {
	return p.PropertyValue("min-inline-size")
}

func (p *cssStyleHelperImpl) SetMinInlineSize(s string) {
	p.SetProperty("min-inline-size", s)
}

func (p *cssStyleHelperImpl) MinWidth() string {
	return p.PropertyValue("min-width")
}

func (p *cssStyleHelperImpl) SetMinWidth(s string) {
	p.SetProperty("min-width", s)
}

func (p *cssStyleHelperImpl) MixBlendMode() string {
	return p.PropertyValue("mix-blend-mode")
}

func (p *cssStyleHelperImpl) SetMixBlendMode(s string) {
	p.SetProperty("mix-blend-mode", s)
}

func (p *cssStyleHelperImpl) NavDown() string {
	return p.PropertyValue("nav-down")
}

func (p *cssStyleHelperImpl) SetNavDown(s string) {
	p.SetProperty("nav-down", s)
}

func (p *cssStyleHelperImpl) NavLeft() string {
	return p.PropertyValue("nav-left")
}

func (p *cssStyleHelperImpl) SetNavLeft(s string) {
	p.SetProperty("nav-left", s)
}

func (p *cssStyleHelperImpl) NavRight() string {
	return p.PropertyValue("nav-right")
}

func (p *cssStyleHelperImpl) SetNavRight(s string) {
	p.SetProperty("nav-right", s)
}

func (p *cssStyleHelperImpl) NavUp() string {
	return p.PropertyValue("nav-up")
}

func (p *cssStyleHelperImpl) SetNavUp(s string) {
	p.SetProperty("nav-up", s)
}

func (p *cssStyleHelperImpl) ObjectFit() string {
	return p.PropertyValue("object-fit")
}

func (p *cssStyleHelperImpl) SetObjectFit(s string) {
	p.SetProperty("object-fit", s)
}

func (p *cssStyleHelperImpl) ObjectPosition() string {
	return p.PropertyValue("object-position")
}

func (p *cssStyleHelperImpl) SetObjectPosition(s string) {
	p.SetProperty("object-position", s)
}

func (p *cssStyleHelperImpl) Offset() string {
	return p.PropertyValue("offset")
}

func (p *cssStyleHelperImpl) SetOffset(s string) {
	p.SetProperty("offset", s)
}

func (p *cssStyleHelperImpl) OffsetAfter() string {
	return p.PropertyValue("offset-after")
}

func (p *cssStyleHelperImpl) SetOffsetAfter(s string) {
	p.SetProperty("offset-after", s)
}

func (p *cssStyleHelperImpl) OffsetAnchor() string {
	return p.PropertyValue("offset-anchor")
}

func (p *cssStyleHelperImpl) SetOffsetAnchor(s string) {
	p.SetProperty("offset-anchor", s)
}

func (p *cssStyleHelperImpl) OffsetBefore() string {
	return p.PropertyValue("offset-before")
}

func (p *cssStyleHelperImpl) SetOffsetBefore(s string) {
	p.SetProperty("offset-before", s)
}

func (p *cssStyleHelperImpl) OffsetDistance() string {
	return p.PropertyValue("offset-distance")
}

func (p *cssStyleHelperImpl) SetOffsetDistance(s string) {
	p.SetProperty("offset-distance", s)
}

func (p *cssStyleHelperImpl) OffsetEnd() string {
	return p.PropertyValue("offset-end")
}

func (p *cssStyleHelperImpl) SetOffsetEnd(s string) {
	p.SetProperty("offset-end", s)
}

func (p *cssStyleHelperImpl) OffsetPath() string {
	return p.PropertyValue("offset-path")
}

func (p *cssStyleHelperImpl) SetOffsetPath(s string) {
	p.SetProperty("offset-path", s)
}

func (p *cssStyleHelperImpl) OffsetPosition() string {
	return p.PropertyValue("offset-position")
}

func (p *cssStyleHelperImpl) SetOffsetPosition(s string) {
	p.SetProperty("offset-position", s)
}

func (p *cssStyleHelperImpl) OffsetRotate() string {
	return p.PropertyValue("offset-rotate")
}

func (p *cssStyleHelperImpl) SetOffsetRotate(s string) {
	p.SetProperty("offset-rotate", s)
}

func (p *cssStyleHelperImpl) OffsetStart() string {
	return p.PropertyValue("offset-start")
}

func (p *cssStyleHelperImpl) SetOffsetStart(s string) {
	p.SetProperty("offset-start", s)
}

func (p *cssStyleHelperImpl) Opacity() string {
	return p.PropertyValue("opacity")
}

func (p *cssStyleHelperImpl) SetOpacity(s string) {
	p.SetProperty("opacity", s)
}

func (p *cssStyleHelperImpl) Order() string {
	return p.PropertyValue("order")
}

func (p *cssStyleHelperImpl) SetOrder(s string) {
	p.SetProperty("order", s)
}

func (p *cssStyleHelperImpl) Orphans() string {
	return p.PropertyValue("orphans")
}

func (p *cssStyleHelperImpl) SetOrphans(s string) {
	p.SetProperty("orphans", s)
}

func (p *cssStyleHelperImpl) Outline() string {
	return p.PropertyValue("outline")
}

func (p *cssStyleHelperImpl) SetOutline(s string) {
	p.SetProperty("outline", s)
}

func (p *cssStyleHelperImpl) OutlineColor() string {
	return p.PropertyValue("outline-color")
}

func (p *cssStyleHelperImpl) SetOutlineColor(s string) {
	p.SetProperty("outline-color", s)
}

func (p *cssStyleHelperImpl) OutlineOffset() string {
	return p.PropertyValue("outline-offset")
}

func (p *cssStyleHelperImpl) SetOutlineOffset(s string) {
	p.SetProperty("outline-offset", s)
}

func (p *cssStyleHelperImpl) OutlineStyle() string {
	return p.PropertyValue("outline-style")
}

func (p *cssStyleHelperImpl) SetOutlineStyle(s string) {
	p.SetProperty("outline-style", s)
}

func (p *cssStyleHelperImpl) OutlineWidth() string {
	return p.PropertyValue("outline-width")
}

func (p *cssStyleHelperImpl) SetOutlineWidth(s string) {
	p.SetProperty("outline-width", s)
}

func (p *cssStyleHelperImpl) Overflow() string {
	return p.PropertyValue("overflow")
}

func (p *cssStyleHelperImpl) SetOverflow(s string) {
	p.SetProperty("overflow", s)
}

func (p *cssStyleHelperImpl) OverflowBlock() string {
	return p.PropertyValue("overflow-block")
}

func (p *cssStyleHelperImpl) SetOverflowBlock(s string) {
	p.SetProperty("overflow-block", s)
}

func (p *cssStyleHelperImpl) OverflowInline() string {
	return p.PropertyValue("overflow-inline")
}

func (p *cssStyleHelperImpl) SetOverflowInline(s string) {
	p.SetProperty("overflow-inline", s)
}

func (p *cssStyleHelperImpl) OverflowWrap() string {
	return p.PropertyValue("overflow-wrap")
}

func (p *cssStyleHelperImpl) SetOverflowWrap(s string) {
	p.SetProperty("overflow-wrap", s)
}

func (p *cssStyleHelperImpl) OverflowX() string {
	return p.PropertyValue("overflow-x")
}

func (p *cssStyleHelperImpl) SetOverflowX(s string) {
	p.SetProperty("overflow-x", s)
}

func (p *cssStyleHelperImpl) OverflowY() string {
	return p.PropertyValue("overflow-y")
}

func (p *cssStyleHelperImpl) SetOverflowY(s string) {
	p.SetProperty("overflow-y", s)
}

func (p *cssStyleHelperImpl) Padding() string {
	return p.PropertyValue("padding")
}

func (p *cssStyleHelperImpl) SetPadding(s string) {
	p.SetProperty("padding", s)
}

func (p *cssStyleHelperImpl) PaddingBlock() string {
	return p.PropertyValue("padding-block")
}

func (p *cssStyleHelperImpl) SetPaddingBlock(s string) {
	p.SetProperty("padding-block", s)
}

func (p *cssStyleHelperImpl) PaddingBlockEnd() string {
	return p.PropertyValue("padding-block-end")
}

func (p *cssStyleHelperImpl) SetPaddingBlockEnd(s string) {
	p.SetProperty("padding-block-end", s)
}

func (p *cssStyleHelperImpl) PaddingBlockStart() string {
	return p.PropertyValue("padding-block-start")
}

func (p *cssStyleHelperImpl) SetPaddingBlockStart(s string) {
	p.SetProperty("padding-block-start", s)
}

func (p *cssStyleHelperImpl) PaddingBottom() string {
	return p.PropertyValue("padding-bottom")
}

func (p *cssStyleHelperImpl) SetPaddingBottom(s string) {
	p.SetProperty("padding-bottom", s)
}

func (p *cssStyleHelperImpl) PaddingInline() string {
	return p.PropertyValue("padding-inline")
}

func (p *cssStyleHelperImpl) SetPaddingInline(s string) {
	p.SetProperty("padding-inline", s)
}

func (p *cssStyleHelperImpl) PaddingInlineEnd() string {
	return p.PropertyValue("padding-inline-end")
}

func (p *cssStyleHelperImpl) SetPaddingInlineEnd(s string) {
	p.SetProperty("padding-inline-end", s)
}

func (p *cssStyleHelperImpl) PaddingInlineStart() string {
	return p.PropertyValue("padding-inline-start")
}

func (p *cssStyleHelperImpl) SetPaddingInlineStart(s string) {
	p.SetProperty("padding-inline-start", s)
}

func (p *cssStyleHelperImpl) PaddingLeft() string {
	return p.PropertyValue("padding-left")
}

func (p *cssStyleHelperImpl) SetPaddingLeft(s string) {
	p.SetProperty("padding-left", s)
}

func (p *cssStyleHelperImpl) PaddingRight() string {
	return p.PropertyValue("padding-right")
}

func (p *cssStyleHelperImpl) SetPaddingRight(s string) {
	p.SetProperty("padding-right", s)
}

func (p *cssStyleHelperImpl) PaddingTop() string {
	return p.PropertyValue("padding-top")
}

func (p *cssStyleHelperImpl) SetPaddingTop(s string) {
	p.SetProperty("padding-top", s)
}

func (p *cssStyleHelperImpl) Page() string {
	return p.PropertyValue("page")
}

func (p *cssStyleHelperImpl) SetPage(s string) {
	p.SetProperty("page", s)
}

func (p *cssStyleHelperImpl) PageBreakAfter() string {
	return p.PropertyValue("page-break-after")
}

func (p *cssStyleHelperImpl) SetPageBreakAfter(s string) {
	p.SetProperty("page-break-after", s)
}

func (p *cssStyleHelperImpl) PageBreakBefore() string {
	return p.PropertyValue("page-break-before")
}

func (p *cssStyleHelperImpl) SetPageBreakBefore(s string) {
	p.SetProperty("page-break-before", s)
}

func (p *cssStyleHelperImpl) PageBreakInside() string {
	return p.PropertyValue("page-break-inside")
}

func (p *cssStyleHelperImpl) SetPageBreakInside(s string) {
	p.SetProperty("page-break-inside", s)
}

func (p *cssStyleHelperImpl) Pause() string {
	return p.PropertyValue("pause")
}

func (p *cssStyleHelperImpl) SetPause(s string) {
	p.SetProperty("pause", s)
}

func (p *cssStyleHelperImpl) PauseAfter() string {
	return p.PropertyValue("pause-after")
}

func (p *cssStyleHelperImpl) SetPauseAfter(s string) {
	p.SetProperty("pause-after", s)
}

func (p *cssStyleHelperImpl) PauseBefore() string {
	return p.PropertyValue("pause-before")
}

func (p *cssStyleHelperImpl) SetPauseBefore(s string) {
	p.SetProperty("pause-before", s)
}

func (p *cssStyleHelperImpl) Pitch() string {
	return p.PropertyValue("pitch")
}

func (p *cssStyleHelperImpl) SetPitch(s string) {
	p.SetProperty("pitch", s)
}

func (p *cssStyleHelperImpl) PitchRange() string {
	return p.PropertyValue("pitch-range")
}

func (p *cssStyleHelperImpl) SetPitchRange(s string) {
	p.SetProperty("pitch-range", s)
}

func (p *cssStyleHelperImpl) PlaceContent() string {
	return p.PropertyValue("place-content")
}

func (p *cssStyleHelperImpl) SetPlaceContent(s string) {
	p.SetProperty("place-content", s)
}

func (p *cssStyleHelperImpl) PlaceItems() string {
	return p.PropertyValue("place-items")
}

func (p *cssStyleHelperImpl) SetPlaceItems(s string) {
	p.SetProperty("place-items", s)
}

func (p *cssStyleHelperImpl) PlaceSelf() string {
	return p.PropertyValue("place-self")
}

func (p *cssStyleHelperImpl) SetPlaceSelf(s string) {
	p.SetProperty("place-self", s)
}

func (p *cssStyleHelperImpl) PlayDuring() string {
	return p.PropertyValue("play-during")
}

func (p *cssStyleHelperImpl) SetPlayDuring(s string) {
	p.SetProperty("play-during", s)
}

func (p *cssStyleHelperImpl) Position() string {
	return p.PropertyValue("position")
}

func (p *cssStyleHelperImpl) SetPosition(s string) {
	p.SetProperty("position", s)
}

func (p *cssStyleHelperImpl) Quotes() string {
	return p.PropertyValue("quotes")
}

func (p *cssStyleHelperImpl) SetQuotes(s string) {
	p.SetProperty("quotes", s)
}

func (p *cssStyleHelperImpl) RegionFragment() string {
	return p.PropertyValue("region-fragment")
}

func (p *cssStyleHelperImpl) SetRegionFragment(s string) {
	p.SetProperty("region-fragment", s)
}

func (p *cssStyleHelperImpl) Resize() string {
	return p.PropertyValue("resize")
}

func (p *cssStyleHelperImpl) SetResize(s string) {
	p.SetProperty("resize", s)
}

func (p *cssStyleHelperImpl) Richness() string {
	return p.PropertyValue("richness")
}

func (p *cssStyleHelperImpl) SetRichness(s string) {
	p.SetProperty("richness", s)
}

func (p *cssStyleHelperImpl) Right() string {
	return p.PropertyValue("right")
}

func (p *cssStyleHelperImpl) SetRight(s string) {
	p.SetProperty("right", s)
}

func (p *cssStyleHelperImpl) RowGap() string {
	return p.PropertyValue("row-gap")
}

func (p *cssStyleHelperImpl) SetRowGap(s string) {
	p.SetProperty("row-gap", s)
}

func (p *cssStyleHelperImpl) RubyAlign() string {
	return p.PropertyValue("ruby-align")
}

func (p *cssStyleHelperImpl) SetRubyAlign(s string) {
	p.SetProperty("ruby-align", s)
}

func (p *cssStyleHelperImpl) RubyMerge() string {
	return p.PropertyValue("ruby-merge")
}

func (p *cssStyleHelperImpl) SetRubyMerge(s string) {
	p.SetProperty("ruby-merge", s)
}

func (p *cssStyleHelperImpl) RubyPosition() string {
	return p.PropertyValue("ruby-position")
}

func (p *cssStyleHelperImpl) SetRubyPosition(s string) {
	p.SetProperty("ruby-position", s)
}

func (p *cssStyleHelperImpl) Running() string {
	return p.PropertyValue("running")
}

func (p *cssStyleHelperImpl) SetRunning(s string) {
	p.SetProperty("running", s)
}

func (p *cssStyleHelperImpl) ScrollBehavior() string {
	return p.PropertyValue("scroll-behavior")
}

func (p *cssStyleHelperImpl) SetScrollBehavior(s string) {
	p.SetProperty("scroll-behavior", s)
}

func (p *cssStyleHelperImpl) ScrollMargin() string {
	return p.PropertyValue("scroll-margin")
}

func (p *cssStyleHelperImpl) SetScrollMargin(s string) {
	p.SetProperty("scroll-margin", s)
}

func (p *cssStyleHelperImpl) ScrollMarginBlock() string {
	return p.PropertyValue("scroll-margin-block")
}

func (p *cssStyleHelperImpl) SetScrollMarginBlock(s string) {
	p.SetProperty("scroll-margin-block", s)
}

func (p *cssStyleHelperImpl) ScrollMarginBlockEnd() string {
	return p.PropertyValue("scroll-margin-block-end")
}

func (p *cssStyleHelperImpl) SetScrollMarginBlockEnd(s string) {
	p.SetProperty("scroll-margin-block-end", s)
}

func (p *cssStyleHelperImpl) ScrollMarginBlockStart() string {
	return p.PropertyValue("scroll-margin-block-start")
}

func (p *cssStyleHelperImpl) SetScrollMarginBlockStart(s string) {
	p.SetProperty("scroll-margin-block-start", s)
}

func (p *cssStyleHelperImpl) ScrollMarginBottom() string {
	return p.PropertyValue("scroll-margin-bottom")
}

func (p *cssStyleHelperImpl) SetScrollMarginBottom(s string) {
	p.SetProperty("scroll-margin-bottom", s)
}

func (p *cssStyleHelperImpl) ScrollMarginInline() string {
	return p.PropertyValue("scroll-margin-inline")
}

func (p *cssStyleHelperImpl) SetScrollMarginInline(s string) {
	p.SetProperty("scroll-margin-inline", s)
}

func (p *cssStyleHelperImpl) ScrollMarginInlineEnd() string {
	return p.PropertyValue("scroll-margin-inline-end")
}

func (p *cssStyleHelperImpl) SetScrollMarginInlineEnd(s string) {
	p.SetProperty("scroll-margin-inline-end", s)
}

func (p *cssStyleHelperImpl) ScrollMarginInlineStart() string {
	return p.PropertyValue("scroll-margin-inline-start")
}

func (p *cssStyleHelperImpl) SetScrollMarginInlineStart(s string) {
	p.SetProperty("scroll-margin-inline-start", s)
}

func (p *cssStyleHelperImpl) ScrollMarginLeft() string {
	return p.PropertyValue("scroll-margin-left")
}

func (p *cssStyleHelperImpl) SetScrollMarginLeft(s string) {
	p.SetProperty("scroll-margin-left", s)
}

func (p *cssStyleHelperImpl) ScrollMarginRight() string {
	return p.PropertyValue("scroll-margin-right")
}

func (p *cssStyleHelperImpl) SetScrollMarginRight(s string) {
	p.SetProperty("scroll-margin-right", s)
}

func (p *cssStyleHelperImpl) ScrollMarginTop() string {
	return p.PropertyValue("scroll-margin-top")
}

func (p *cssStyleHelperImpl) SetScrollMarginTop(s string) {
	p.SetProperty("scroll-margin-top", s)
}

func (p *cssStyleHelperImpl) ScrollPadding() string {
	return p.PropertyValue("scroll-padding")
}

func (p *cssStyleHelperImpl) SetScrollPadding(s string) {
	p.SetProperty("scroll-padding", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingBlock() string {
	return p.PropertyValue("scroll-padding-block")
}

func (p *cssStyleHelperImpl) SetScrollPaddingBlock(s string) {
	p.SetProperty("scroll-padding-block", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingBlockEnd() string {
	return p.PropertyValue("scroll-padding-block-end")
}

func (p *cssStyleHelperImpl) SetScrollPaddingBlockEnd(s string) {
	p.SetProperty("scroll-padding-block-end", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingBlockStart() string {
	return p.PropertyValue("scroll-padding-block-start")
}

func (p *cssStyleHelperImpl) SetScrollPaddingBlockStart(s string) {
	p.SetProperty("scroll-padding-block-start", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingBottom() string {
	return p.PropertyValue("scroll-padding-bottom")
}

func (p *cssStyleHelperImpl) SetScrollPaddingBottom(s string) {
	p.SetProperty("scroll-padding-bottom", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingInline() string {
	return p.PropertyValue("scroll-padding-inline")
}

func (p *cssStyleHelperImpl) SetScrollPaddingInline(s string) {
	p.SetProperty("scroll-padding-inline", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingInlineEnd() string {
	return p.PropertyValue("scroll-padding-inline-end")
}

func (p *cssStyleHelperImpl) SetScrollPaddingInlineEnd(s string) {
	p.SetProperty("scroll-padding-inline-end", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingInlineStart() string {
	return p.PropertyValue("scroll-padding-inline-start")
}

func (p *cssStyleHelperImpl) SetScrollPaddingInlineStart(s string) {
	p.SetProperty("scroll-padding-inline-start", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingLeft() string {
	return p.PropertyValue("scroll-padding-left")
}

func (p *cssStyleHelperImpl) SetScrollPaddingLeft(s string) {
	p.SetProperty("scroll-padding-left", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingRight() string {
	return p.PropertyValue("scroll-padding-right")
}

func (p *cssStyleHelperImpl) SetScrollPaddingRight(s string) {
	p.SetProperty("scroll-padding-right", s)
}

func (p *cssStyleHelperImpl) ScrollPaddingTop() string {
	return p.PropertyValue("scroll-padding-top")
}

func (p *cssStyleHelperImpl) SetScrollPaddingTop(s string) {
	p.SetProperty("scroll-padding-top", s)
}

func (p *cssStyleHelperImpl) ScrollSnapAlign() string {
	return p.PropertyValue("scroll-snap-align")
}

func (p *cssStyleHelperImpl) SetScrollSnapAlign(s string) {
	p.SetProperty("scroll-snap-align", s)
}

func (p *cssStyleHelperImpl) ScrollSnapStop() string {
	return p.PropertyValue("scroll-snap-stop")
}

func (p *cssStyleHelperImpl) SetScrollSnapStop(s string) {
	p.SetProperty("scroll-snap-stop", s)
}

func (p *cssStyleHelperImpl) ScrollSnapType() string {
	return p.PropertyValue("scroll-snap-type")
}

func (p *cssStyleHelperImpl) SetScrollSnapType(s string) {
	p.SetProperty("scroll-snap-type", s)
}

func (p *cssStyleHelperImpl) ShapeImageThreshold() string {
	return p.PropertyValue("shape-image-threshold")
}

func (p *cssStyleHelperImpl) SetShapeImageThreshold(s string) {
	p.SetProperty("shape-image-threshold", s)
}

func (p *cssStyleHelperImpl) ShapeInside() string {
	return p.PropertyValue("shape-inside")
}

func (p *cssStyleHelperImpl) SetShapeInside(s string) {
	p.SetProperty("shape-inside", s)
}

func (p *cssStyleHelperImpl) ShapeMargin() string {
	return p.PropertyValue("shape-margin")
}

func (p *cssStyleHelperImpl) SetShapeMargin(s string) {
	p.SetProperty("shape-margin", s)
}

func (p *cssStyleHelperImpl) ShapeOutside() string {
	return p.PropertyValue("shape-outside")
}

func (p *cssStyleHelperImpl) SetShapeOutside(s string) {
	p.SetProperty("shape-outside", s)
}

func (p *cssStyleHelperImpl) Speak() string {
	return p.PropertyValue("speak")
}

func (p *cssStyleHelperImpl) SetSpeak(s string) {
	p.SetProperty("speak", s)
}

func (p *cssStyleHelperImpl) SpeakHeader() string {
	return p.PropertyValue("speak-header")
}

func (p *cssStyleHelperImpl) SetSpeakHeader(s string) {
	p.SetProperty("speak-header", s)
}

func (p *cssStyleHelperImpl) SpeakNumeral() string {
	return p.PropertyValue("speak-numeral")
}

func (p *cssStyleHelperImpl) SetSpeakNumeral(s string) {
	p.SetProperty("speak-numeral", s)
}

func (p *cssStyleHelperImpl) SpeakPunctuation() string {
	return p.PropertyValue("speak-punctuation")
}

func (p *cssStyleHelperImpl) SetSpeakPunctuation(s string) {
	p.SetProperty("speak-punctuation", s)
}

func (p *cssStyleHelperImpl) SpeechRate() string {
	return p.PropertyValue("speech-rate")
}

func (p *cssStyleHelperImpl) SetSpeechRate(s string) {
	p.SetProperty("speech-rate", s)
}

func (p *cssStyleHelperImpl) Stress() string {
	return p.PropertyValue("stress")
}

func (p *cssStyleHelperImpl) SetStress(s string) {
	p.SetProperty("stress", s)
}

func (p *cssStyleHelperImpl) StringSet() string {
	return p.PropertyValue("string-set")
}

func (p *cssStyleHelperImpl) SetStringSet(s string) {
	p.SetProperty("string-set", s)
}

func (p *cssStyleHelperImpl) TabSize() string {
	return p.PropertyValue("tab-size")
}

func (p *cssStyleHelperImpl) SetTabSize(s string) {
	p.SetProperty("tab-size", s)
}

func (p *cssStyleHelperImpl) TableLayout() string {
	return p.PropertyValue("table-layout")
}

func (p *cssStyleHelperImpl) SetTableLayout(s string) {
	p.SetProperty("table-layout", s)
}

func (p *cssStyleHelperImpl) TextAlign() string {
	return p.PropertyValue("text-align")
}

func (p *cssStyleHelperImpl) SetTextAlign(s string) {
	p.SetProperty("text-align", s)
}

func (p *cssStyleHelperImpl) TextAlignAll() string {
	return p.PropertyValue("text-align-all")
}

func (p *cssStyleHelperImpl) SetTextAlignAll(s string) {
	p.SetProperty("text-align-all", s)
}

func (p *cssStyleHelperImpl) TextAlignLast() string {
	return p.PropertyValue("text-align-last")
}

func (p *cssStyleHelperImpl) SetTextAlignLast(s string) {
	p.SetProperty("text-align-last", s)
}

func (p *cssStyleHelperImpl) TextCombineUpright() string {
	return p.PropertyValue("text-combine-upright")
}

func (p *cssStyleHelperImpl) SetTextCombineUpright(s string) {
	p.SetProperty("text-combine-upright", s)
}

func (p *cssStyleHelperImpl) TextDecoration() string {
	return p.PropertyValue("text-decoration")
}

func (p *cssStyleHelperImpl) SetTextDecoration(s string) {
	p.SetProperty("text-decoration", s)
}

func (p *cssStyleHelperImpl) TextDecorationColor() string {
	return p.PropertyValue("text-decoration-color")
}

func (p *cssStyleHelperImpl) SetTextDecorationColor(s string) {
	p.SetProperty("text-decoration-color", s)
}

func (p *cssStyleHelperImpl) TextDecorationLine() string {
	return p.PropertyValue("text-decoration-line")
}

func (p *cssStyleHelperImpl) SetTextDecorationLine(s string) {
	p.SetProperty("text-decoration-line", s)
}

func (p *cssStyleHelperImpl) TextDecorationStyle() string {
	return p.PropertyValue("text-decoration-style")
}

func (p *cssStyleHelperImpl) SetTextDecorationStyle(s string) {
	p.SetProperty("text-decoration-style", s)
}

func (p *cssStyleHelperImpl) TextEmphasis() string {
	return p.PropertyValue("text-emphasis")
}

func (p *cssStyleHelperImpl) SetTextEmphasis(s string) {
	p.SetProperty("text-emphasis", s)
}

func (p *cssStyleHelperImpl) TextEmphasisColor() string {
	return p.PropertyValue("text-emphasis-color")
}

func (p *cssStyleHelperImpl) SetTextEmphasisColor(s string) {
	p.SetProperty("text-emphasis-color", s)
}

func (p *cssStyleHelperImpl) TextEmphasisPosition() string {
	return p.PropertyValue("text-emphasis-position")
}

func (p *cssStyleHelperImpl) SetTextEmphasisPosition(s string) {
	p.SetProperty("text-emphasis-position", s)
}

func (p *cssStyleHelperImpl) TextEmphasisStyle() string {
	return p.PropertyValue("text-emphasis-style")
}

func (p *cssStyleHelperImpl) SetTextEmphasisStyle(s string) {
	p.SetProperty("text-emphasis-style", s)
}

func (p *cssStyleHelperImpl) TextGroupAlign() string {
	return p.PropertyValue("text-group-align")
}

func (p *cssStyleHelperImpl) SetTextGroupAlign(s string) {
	p.SetProperty("text-group-align", s)
}

func (p *cssStyleHelperImpl) TextIndent() string {
	return p.PropertyValue("text-indent")
}

func (p *cssStyleHelperImpl) SetTextIndent(s string) {
	p.SetProperty("text-indent", s)
}

func (p *cssStyleHelperImpl) TextJustify() string {
	return p.PropertyValue("text-justify")
}

func (p *cssStyleHelperImpl) SetTextJustify(s string) {
	p.SetProperty("text-justify", s)
}

func (p *cssStyleHelperImpl) TextOrientation() string {
	return p.PropertyValue("text-orientation")
}

func (p *cssStyleHelperImpl) SetTextOrientation(s string) {
	p.SetProperty("text-orientation", s)
}

func (p *cssStyleHelperImpl) TextOverflow() string {
	return p.PropertyValue("text-overflow")
}

func (p *cssStyleHelperImpl) SetTextOverflow(s string) {
	p.SetProperty("text-overflow", s)
}

func (p *cssStyleHelperImpl) TextShadow() string {
	return p.PropertyValue("text-shadow")
}

func (p *cssStyleHelperImpl) SetTextShadow(s string) {
	p.SetProperty("text-shadow", s)
}

func (p *cssStyleHelperImpl) TextSpaceCollapse() string {
	return p.PropertyValue("text-space-collapse")
}

func (p *cssStyleHelperImpl) SetTextSpaceCollapse(s string) {
	p.SetProperty("text-space-collapse", s)
}

func (p *cssStyleHelperImpl) TextSpaceTrim() string {
	return p.PropertyValue("text-space-trim")
}

func (p *cssStyleHelperImpl) SetTextSpaceTrim(s string) {
	p.SetProperty("text-space-trim", s)
}

func (p *cssStyleHelperImpl) TextSpacing() string {
	return p.PropertyValue("text-spacing")
}

func (p *cssStyleHelperImpl) SetTextSpacing(s string) {
	p.SetProperty("text-spacing", s)
}

func (p *cssStyleHelperImpl) TextTransform() string {
	return p.PropertyValue("text-transform")
}

func (p *cssStyleHelperImpl) SetTextTransform(s string) {
	p.SetProperty("text-transform", s)
}

func (p *cssStyleHelperImpl) TextUnderlinePosition() string {
	return p.PropertyValue("text-underline-position")
}

func (p *cssStyleHelperImpl) SetTextUnderlinePosition(s string) {
	p.SetProperty("text-underline-position", s)
}

func (p *cssStyleHelperImpl) TextWrap() string {
	return p.PropertyValue("text-wrap")
}

func (p *cssStyleHelperImpl) SetTextWrap(s string) {
	p.SetProperty("text-wrap", s)
}

func (p *cssStyleHelperImpl) Top() string {
	return p.PropertyValue("top")
}

func (p *cssStyleHelperImpl) SetTop(s string) {
	p.SetProperty("top", s)
}

func (p *cssStyleHelperImpl) Transform() string {
	return p.PropertyValue("transform")
}

func (p *cssStyleHelperImpl) SetTransform(s string) {
	p.SetProperty("transform", s)
}

func (p *cssStyleHelperImpl) TransformBox() string {
	return p.PropertyValue("transform-box")
}

func (p *cssStyleHelperImpl) SetTransformBox(s string) {
	p.SetProperty("transform-box", s)
}

func (p *cssStyleHelperImpl) TransformOrigin() string {
	return p.PropertyValue("transform-origin")
}

func (p *cssStyleHelperImpl) SetTransformOrigin(s string) {
	p.SetProperty("transform-origin", s)
}

func (p *cssStyleHelperImpl) Transition() string {
	return p.PropertyValue("transition")
}

func (p *cssStyleHelperImpl) SetTransition(s string) {
	p.SetProperty("transition", s)
}

func (p *cssStyleHelperImpl) TransitionDelay() string {
	return p.PropertyValue("transition-delay")
}

func (p *cssStyleHelperImpl) SetTransitionDelay(s string) {
	p.SetProperty("transition-delay", s)
}

func (p *cssStyleHelperImpl) TransitionDuration() string {
	return p.PropertyValue("transition-duration")
}

func (p *cssStyleHelperImpl) SetTransitionDuration(s string) {
	p.SetProperty("transition-duration", s)
}

func (p *cssStyleHelperImpl) TransitionProperty() string {
	return p.PropertyValue("transition-property")
}

func (p *cssStyleHelperImpl) SetTransitionProperty(s string) {
	p.SetProperty("transition-property", s)
}

func (p *cssStyleHelperImpl) TransitionTimingFunction() string {
	return p.PropertyValue("transition-timing-function")
}

func (p *cssStyleHelperImpl) SetTransitionTimingFunction(s string) {
	p.SetProperty("transition-timing-function", s)
}

func (p *cssStyleHelperImpl) UnicodeBidi() string {
	return p.PropertyValue("unicode-bidi")
}

func (p *cssStyleHelperImpl) SetUnicodeBidi(s string) {
	p.SetProperty("unicode-bidi", s)
}

func (p *cssStyleHelperImpl) UserSelect() string {
	return p.PropertyValue("user-select")
}

func (p *cssStyleHelperImpl) SetUserSelect(s string) {
	p.SetProperty("user-select", s)
}

func (p *cssStyleHelperImpl) VerticalAlign() string {
	return p.PropertyValue("vertical-align")
}

func (p *cssStyleHelperImpl) SetVerticalAlign(s string) {
	p.SetProperty("vertical-align", s)
}

func (p *cssStyleHelperImpl) Visibility() string {
	return p.PropertyValue("visibility")
}

func (p *cssStyleHelperImpl) SetVisibility(s string) {
	p.SetProperty("visibility", s)
}

func (p *cssStyleHelperImpl) VoiceFamily() string {
	return p.PropertyValue("voice-family")
}

func (p *cssStyleHelperImpl) SetVoiceFamily(s string) {
	p.SetProperty("voice-family", s)
}

func (p *cssStyleHelperImpl) Volume() string {
	return p.PropertyValue("volume")
}

func (p *cssStyleHelperImpl) SetVolume(s string) {
	p.SetProperty("volume", s)
}

func (p *cssStyleHelperImpl) WhiteSpace() string {
	return p.PropertyValue("white-space")
}

func (p *cssStyleHelperImpl) SetWhiteSpace(s string) {
	p.SetProperty("white-space", s)
}

func (p *cssStyleHelperImpl) Widows() string {
	return p.PropertyValue("widows")
}

func (p *cssStyleHelperImpl) SetWidows(s string) {
	p.SetProperty("widows", s)
}

func (p *cssStyleHelperImpl) Width() string {
	return p.PropertyValue("width")
}

func (p *cssStyleHelperImpl) SetWidth(s string) {
	p.SetProperty("width", s)
}

func (p *cssStyleHelperImpl) WillChange() string {
	return p.PropertyValue("will-change")
}

func (p *cssStyleHelperImpl) SetWillChange(s string) {
	p.SetProperty("will-change", s)
}

func (p *cssStyleHelperImpl) WordBreak() string {
	return p.PropertyValue("word-break")
}

func (p *cssStyleHelperImpl) SetWordBreak(s string) {
	p.SetProperty("word-break", s)
}

func (p *cssStyleHelperImpl) WordSpacing() string {
	return p.PropertyValue("word-spacing")
}

func (p *cssStyleHelperImpl) SetWordSpacing(s string) {
	p.SetProperty("word-spacing", s)
}

func (p *cssStyleHelperImpl) WordWrap() string {
	return p.PropertyValue("word-wrap")
}

func (p *cssStyleHelperImpl) SetWordWrap(s string) {
	p.SetProperty("word-wrap", s)
}

func (p *cssStyleHelperImpl) WrapAfter() string {
	return p.PropertyValue("wrap-after")
}

func (p *cssStyleHelperImpl) SetWrapAfter(s string) {
	p.SetProperty("wrap-after", s)
}

func (p *cssStyleHelperImpl) WrapBefore() string {
	return p.PropertyValue("wrap-before")
}

func (p *cssStyleHelperImpl) SetWrapBefore(s string) {
	p.SetProperty("wrap-before", s)
}

func (p *cssStyleHelperImpl) WrapFlow() string {
	return p.PropertyValue("wrap-flow")
}

func (p *cssStyleHelperImpl) SetWrapFlow(s string) {
	p.SetProperty("wrap-flow", s)
}

func (p *cssStyleHelperImpl) WrapInside() string {
	return p.PropertyValue("wrap-inside")
}

func (p *cssStyleHelperImpl) SetWrapInside(s string) {
	p.SetProperty("wrap-inside", s)
}

func (p *cssStyleHelperImpl) WrapThrough() string {
	return p.PropertyValue("wrap-through")
}

func (p *cssStyleHelperImpl) SetWrapThrough(s string) {
	p.SetProperty("wrap-through", s)
}

func (p *cssStyleHelperImpl) WritingMode() string {
	return p.PropertyValue("writing-mode")
}

func (p *cssStyleHelperImpl) SetWritingMode(s string) {
	p.SetProperty("writing-mode", s)
}

func (p *cssStyleHelperImpl) ZIndex() string {
	return p.PropertyValue("z-index")
}

func (p *cssStyleHelperImpl) SetZIndex(s string) {
	p.SetProperty("z-index", s)
}
