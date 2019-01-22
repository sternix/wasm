// +build js,wasm

package wasm

// -------------8<---------------------------------------

type canvasRenderingContext2DImpl struct {
	*canvasDrawingStylesImpl
	*canvasPathMethodsImpl
	Value
}

func wrapCanvasRenderingContext2D(v Value) CanvasRenderingContext2D {
	if v.valid() {
		return &canvasRenderingContext2DImpl{
			canvasDrawingStylesImpl: newCanvasDrawingStylesImpl(v),
			canvasPathMethodsImpl:   newCanvasPathMethodsImpl(v),
			Value:                   v,
		}
	}
	return nil
}

func (p *canvasRenderingContext2DImpl) Canvas() HTMLCanvasElement {
	return wrapHTMLCanvasElement(p.get("canvas"))
}

func (p *canvasRenderingContext2DImpl) Save() {
	p.call("save")
}

func (p *canvasRenderingContext2DImpl) Restore() {
	p.call("restore")
}

func (p *canvasRenderingContext2DImpl) Scale(x float64, y float64) {
	p.call("scale", x, y)
}

func (p *canvasRenderingContext2DImpl) Rotate(angle float64) {
	p.call("rotate", angle)
}

func (p *canvasRenderingContext2DImpl) Translate(x float64, y float64) {
	p.call("translate", x, y)
}

func (p *canvasRenderingContext2DImpl) Transform(a float64, b float64, c float64, d float64, e float64, f float64) {
	p.call("transform", a, b, c, d, e, f)
}

func (p *canvasRenderingContext2DImpl) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) {
	p.call("setTransform", a, b, c, d, e, f)
}

func (p *canvasRenderingContext2DImpl) GlobalAlpha() float64 {
	return p.get("globalAlpha").toFloat64()
}

func (p *canvasRenderingContext2DImpl) SetGlobalAlpha(ga float64) {
	p.set("globalAlpha", ga)
}

func (p *canvasRenderingContext2DImpl) GlobalCompositeOperation() string {
	return p.get("globalCompositeOperation").toString()
}

func (p *canvasRenderingContext2DImpl) SetGlobalCompositeOperation(gco string) {
	p.set("globalCompositeOperation", gco)
}

func (p *canvasRenderingContext2DImpl) StrokeStyle() interface{} {
	return Wrap(p.get("strokeStyle"))
}

// attribute (DOMString or CanvasGradient or CanvasPattern) strokeStyle; // (default: "black")
func (p *canvasRenderingContext2DImpl) SetStrokeStyle(style interface{}) {
	switch x := style.(type) {
	case string:
		p.set("strokeStyle", x)
	case CanvasGradient, CanvasPattern:
		p.set("strokeStyle", JSValueOf(x))
	}
}

func (p *canvasRenderingContext2DImpl) FillStyle() interface{} {
	return Wrap(p.get("fillStyle"))
}

// attribute (DOMString or CanvasGradient or CanvasPattern) strokeStyle; // (default: "black")
func (p *canvasRenderingContext2DImpl) SetFillStyle(style interface{}) {
	switch x := style.(type) {
	case string:
		p.set("fillStyle", x)
	case CanvasGradient, CanvasPattern:
		p.set("fillStyle", JSValueOf(x))
	}
}

func (p *canvasRenderingContext2DImpl) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient {
	return wrapCanvasGradient(p.call("createLinearGradient", x0, y0, x1, y1))
}

func (p *canvasRenderingContext2DImpl) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient {
	return wrapCanvasGradient(p.call("createRadialGradient", x0, y0, r0, x1, y1, r1))
}

func (p *canvasRenderingContext2DImpl) CreatePattern(image CanvasImageSource, repetition string) {
	p.call("createPattern", JSValueOf(image), repetition)
}

func (p *canvasRenderingContext2DImpl) ShadowOffsetX() float64 {
	return p.get("shadowOffsetX").toFloat64()
}

func (p *canvasRenderingContext2DImpl) SetShadowOffsetX(so float64) {
	p.set("shadowOffsetX", so)
}

func (p *canvasRenderingContext2DImpl) ShadowOffsetY() float64 {
	return p.get("shadowOffsetY").toFloat64()
}

func (p *canvasRenderingContext2DImpl) SetShadowOffsetY(so float64) {
	p.set("shadowOffsetY", so)
}

func (p *canvasRenderingContext2DImpl) ShadowBlur() float64 {
	return p.get("shadowBlur").toFloat64()
}

func (p *canvasRenderingContext2DImpl) SetShadowBlur(sb float64) {
	p.set("shadowBlur", sb)
}

func (p *canvasRenderingContext2DImpl) ShadowColor() string {
	return p.get("shadowColor").toString()
}

func (p *canvasRenderingContext2DImpl) SetShadowColor(color string) {
	p.set("shadowColor", color)
}

func (p *canvasRenderingContext2DImpl) ClearRect(x float64, y float64, w float64, h float64) {
	p.call("clearRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) FillRect(x float64, y float64, w float64, h float64) {
	p.call("fillRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) StrokeRect(x float64, y float64, w float64, h float64) {
	p.call("strokeRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) BeginPath() {
	p.call("beginPath")
}

func (p *canvasRenderingContext2DImpl) Fill() {
	p.call("fill")
}

func (p *canvasRenderingContext2DImpl) Stroke() {
	p.call("stroke")
}

func (p *canvasRenderingContext2DImpl) DrawFocusIfNeeded(elm Element) {
	p.call("drawFocusIfNeeded", JSValueOf(elm))
}

func (p *canvasRenderingContext2DImpl) Clip() {
	p.call("clip")
}

func (p *canvasRenderingContext2DImpl) IsPointInPath(x float64, y float64) bool {
	return p.call("isPointInPath", x, y).toBool()
}

func (p *canvasRenderingContext2DImpl) FillText(text string, x float64, y float64, maxWidth ...float64) {
	switch len(maxWidth) {
	case 0:
		p.call("fillText", text, x, y)
	default:
		p.call("fillText", text, x, y, maxWidth[0])
	}
}

func (p *canvasRenderingContext2DImpl) StrokeText(text string, x float64, y float64, maxWidth ...float64) {
	switch len(maxWidth) {
	case 0:
		p.call("strokeText", text, x, y)
	default:
		p.call("strokeText", text, x, y, maxWidth[0])
	}
}

func (p *canvasRenderingContext2DImpl) MeasureText(text string) TextMetrics {
	return wrapTextMetrics(p.call("measureText", text))
}

func (p *canvasRenderingContext2DImpl) DrawImage(image CanvasImageSource, args ...float64) {
	switch len(args) {
	case 2:
		p.call("drawImage", JSValueOf(image), args[0], args[1])
	case 4:
		p.call("drawImage", JSValueOf(image), args[0], args[1], args[2], args[3])
	case 8:
		p.call("drawImage", JSValueOf(image), args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	}
}

func (p *canvasRenderingContext2DImpl) AddHitRegion(options HitRegionOptions) {
	p.call("addHitRegion", options.JSValue())
}

func (p *canvasRenderingContext2DImpl) RemoveHitRegion(id string) {
	p.call("removeHitRegion", id)
}

func (p *canvasRenderingContext2DImpl) ClearHitRegions() {
	p.call("clearHitRegions")
}

func (p *canvasRenderingContext2DImpl) CreateImageData(sw float64, sh float64) ImageData {
	return wrapImageData(p.call("createImageData", sw, sh))
}

func (p *canvasRenderingContext2DImpl) CreateImageDataFromImageData(imageData ImageData) ImageData {
	return wrapImageData(p.call("createImageData", JSValueOf(imageData)))
}

func (p *canvasRenderingContext2DImpl) ImageData(sx float64, sy float64, sw float64, sh float64) ImageData {
	return wrapImageData(p.call("getImageData", sx, sy, sw, sh))
}

func (p *canvasRenderingContext2DImpl) PutImageData(imageData ImageData, dx float64, dy float64) {
	p.call("putImageData", JSValueOf(imageData), dx, dy)
}

func (p *canvasRenderingContext2DImpl) PutImageDataDirty(imageData ImageData, dx float64, dy float64, dirtyX float64, dirtyY float64, dirtyWidth float64, dirtyHeight float64) {
	p.call("putImageData", JSValueOf(imageData), dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

// -------------8<---------------------------------------

type canvasDrawingStylesImpl struct {
	Value
}

func wrapCanvasDrawingStyles(v Value) CanvasDrawingStyles {
	if p := newCanvasDrawingStylesImpl(v); p != nil {
		return p
	}
	return nil
}

func newCanvasDrawingStylesImpl(v Value) *canvasDrawingStylesImpl {
	if v.valid() {
		return &canvasDrawingStylesImpl{
			Value: v,
		}
	}
	return nil
}

func (p *canvasDrawingStylesImpl) LineWidth() float64 {
	return p.get("lineWidth").toFloat64()
}

func (p *canvasDrawingStylesImpl) SetLineWidth(w float64) {
	p.set("lineWidth", w)
}

func (p *canvasDrawingStylesImpl) LineCap() string {
	return p.get("lineCap").toString()
}

func (p *canvasDrawingStylesImpl) SetLineCap(c string) {
	p.set("lineCap", c)
}

func (p *canvasDrawingStylesImpl) LineJoin() string {
	return p.get("lineJoin").toString()
}

func (p *canvasDrawingStylesImpl) SetLineJoin(j string) {
	p.set("lineJoin", j)
}

func (p *canvasDrawingStylesImpl) MiterLimit() float64 {
	return p.get("miterLimit").toFloat64()
}

func (p *canvasDrawingStylesImpl) SetMiterLimit(l float64) {
	p.set("miterLimit", l)
}

func (p *canvasDrawingStylesImpl) LineDash() []float64 {
	return floatSequenceToSlice(p.call("getLineDash"))
}

func (p *canvasDrawingStylesImpl) SetLineDash(arg ...float64) {
	// TODO check this
	p.call("setLineDash", ToJSArray(arg))
}

func (p *canvasDrawingStylesImpl) LineDashOffset() float64 {
	return p.get("lineDashOffset").toFloat64()
}

func (p *canvasDrawingStylesImpl) Font() string {
	return p.get("font").toString()
}

func (p *canvasDrawingStylesImpl) SetFont(font string) {
	p.set("font", font)
}

func (p *canvasDrawingStylesImpl) TextAlign() TextAlign {
	return TextAlign(p.get("textAlign").toString())
}

func (p *canvasDrawingStylesImpl) SetTextAlign(ta TextAlign) {
	p.set("textAlign", string(ta))
}

func (p *canvasDrawingStylesImpl) TextBaseline() TextBaseline {
	return TextBaseline(p.get("textBaseline").toString())
}

func (p *canvasDrawingStylesImpl) SetTextBaseline(tbl TextBaseline) {
	p.set("textBaseline", string(tbl))
}

// -------------8<---------------------------------------

type canvasPathMethodsImpl struct {
	Value
}

func wrapCanvasPathMethods(v Value) CanvasPathMethods {
	if p := newCanvasPathMethodsImpl(v); p != nil {
		return p
	}
	return nil
}

func newCanvasPathMethodsImpl(v Value) *canvasPathMethodsImpl {
	if v.valid() {
		return &canvasPathMethodsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *canvasPathMethodsImpl) ClosePath() {
	p.call("closePath")
}

func (p *canvasPathMethodsImpl) MoveTo(x float64, y float64) {
	p.call("moveTo", x, y)
}

func (p *canvasPathMethodsImpl) LineTo(x float64, y float64) {
	p.call("lineTo", x, y)
}

func (p *canvasPathMethodsImpl) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) {
	p.call("quadraticCurveTo", cpx, cpy, x, y)
}

func (p *canvasPathMethodsImpl) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	p.call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (p *canvasPathMethodsImpl) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	p.call("arcTo", x1, y1, x2, y2, radius)
}

func (p *canvasPathMethodsImpl) Rect(x float64, y float64, w float64, h float64) {
	p.call("rect", x, y, w, h)
}

func (p *canvasPathMethodsImpl) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterClockwise ...bool) {
	switch len(counterClockwise) {
	case 0:
		p.call("arc", x, y, radius, startAngle, endAngle)
	default:
		p.call("arc", x, y, radius, startAngle, endAngle, counterClockwise[0])
	}
}

// -------------8<---------------------------------------

type canvasGradientImpl struct {
	Value
}

func wrapCanvasGradient(v Value) CanvasGradient {
	if v.valid() {
		return &canvasGradientImpl{
			Value: v,
		}
	}
	return nil
}

func (p *canvasGradientImpl) AddColorStop(offset float64, color string) {
	p.call("addColorStop", offset, color)
}

// -------------8<---------------------------------------

type canvasPatternImpl struct {
	Value
}

func wrapCanvasPattern(v Value) CanvasPattern {
	if v.valid() {
		return &canvasPatternImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type textMetricsImpl struct {
	Value
}

func wrapTextMetrics(v Value) TextMetrics {
	if v.valid() {
		return &textMetricsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *textMetricsImpl) Width() float64 {
	return p.get("width").toFloat64()
}

// -------------8<---------------------------------------

type imageDataImpl struct {
	Value
}

func wrapImageData(v Value) ImageData {
	if v.valid() {
		return &imageDataImpl{
			Value: v,
		}
	}
	return nil
}

func (p *imageDataImpl) Width() int {
	return p.get("width").toInt()
}

func (p *imageDataImpl) Height() int {
	return p.get("height").toInt()
}

func (p *imageDataImpl) Data() []byte {
	// TODO  Uint8ClampedArray
	return nil
}

// -------------8<---------------------------------------
