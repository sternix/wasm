// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type canvasRenderingContext2DImpl struct {
	*canvasDrawingStylesImpl
	*canvasPathMethodsImpl
	js.Value
}

func wrapCanvasRenderingContext2D(v js.Value) CanvasRenderingContext2D {
	if isNil(v) {
		return nil
	}

	return &canvasRenderingContext2DImpl{
		canvasDrawingStylesImpl: newCanvasDrawingStylesImpl(v),
		canvasPathMethodsImpl:   newCanvasPathMethodsImpl(v),
		Value:                   v,
	}
}

func (p *canvasRenderingContext2DImpl) Canvas() HTMLCanvasElement {
	return wrapHTMLCanvasElement(p.Get("canvas"))
}

func (p *canvasRenderingContext2DImpl) Save() {
	p.Call("save")
}

func (p *canvasRenderingContext2DImpl) Restore() {
	p.Call("restore")
}

func (p *canvasRenderingContext2DImpl) Scale(x float64, y float64) {
	p.Call("scale", x, y)
}

func (p *canvasRenderingContext2DImpl) Rotate(angle float64) {
	p.Call("rotate", angle)
}

func (p *canvasRenderingContext2DImpl) Translate(x float64, y float64) {
	p.Call("translate", x, y)
}

func (p *canvasRenderingContext2DImpl) Transform(a float64, b float64, c float64, d float64, e float64, f float64) {
	p.Call("transform", a, b, c, d, e, f)
}

func (p *canvasRenderingContext2DImpl) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) {
	p.Call("setTransform", a, b, c, d, e, f)
}

func (p *canvasRenderingContext2DImpl) GlobalAlpha() float64 {
	return p.Get("globalAlpha").Float()
}

func (p *canvasRenderingContext2DImpl) SetGlobalAlpha(ga float64) {
	p.Set("globalAlpha", ga)
}

func (p *canvasRenderingContext2DImpl) GlobalCompositeOperation() string {
	return p.Get("globalCompositeOperation").String()
}

func (p *canvasRenderingContext2DImpl) SetGlobalCompositeOperation(gco string) {
	p.Set("globalCompositeOperation", gco)
}

func (p *canvasRenderingContext2DImpl) StrokeStyle() interface{} {
	return Wrap(p.Get("strokeStyle"))
}

// attribute (DOMString or CanvasGradient or CanvasPattern) strokeStyle; // (default: "black")
func (p *canvasRenderingContext2DImpl) SetStrokeStyle(style interface{}) {
	switch x := style.(type) {
	case string:
		p.Set("strokeStyle", x)
	case CanvasGradient:
		p.Set("strokeStyle", x.JSValue())
	case CanvasPattern:
		p.Set("strokeStyle", x.JSValue())
	}
}

func (p *canvasRenderingContext2DImpl) FillStyle() interface{} {
	return Wrap(p.Get("fillStyle"))
}

// attribute (DOMString or CanvasGradient or CanvasPattern) strokeStyle; // (default: "black")
func (p *canvasRenderingContext2DImpl) SetFillStyle(style interface{}) {
	switch x := style.(type) {
	case string:
		p.Set("fillStyle", x)
	case CanvasGradient:
		p.Set("fillStyle", x.JSValue())
	case CanvasPattern:
		p.Set("fillStyle", x.JSValue())
	}
}

func (p *canvasRenderingContext2DImpl) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient {
	return wrapCanvasGradient(p.Call("createLinearGradient", x0, y0, x1, y1))
}

func (p *canvasRenderingContext2DImpl) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient {
	return wrapCanvasGradient(p.Call("createRadialGradient", x0, y0, r0, x1, y1, r1))
}

func (p *canvasRenderingContext2DImpl) CreatePattern(image CanvasImageSource, repetition string) {
	p.Call("createPattern", image.JSValue(), repetition)
}

func (p *canvasRenderingContext2DImpl) ShadowOffsetX() float64 {
	return p.Get("shadowOffsetX").Float()
}

func (p *canvasRenderingContext2DImpl) SetShadowOffsetX(so float64) {
	p.Set("shadowOffsetX", so)
}

func (p *canvasRenderingContext2DImpl) ShadowOffsetY() float64 {
	return p.Get("shadowOffsetY").Float()
}

func (p *canvasRenderingContext2DImpl) SetShadowOffsetY(so float64) {
	p.Set("shadowOffsetY", so)
}

func (p *canvasRenderingContext2DImpl) ShadowBlur() float64 {
	return p.Get("shadowBlur").Float()
}

func (p *canvasRenderingContext2DImpl) SetShadowBlur(sb float64) {
	p.Set("shadowBlur", sb)
}

func (p *canvasRenderingContext2DImpl) ShadowColor() string {
	return p.Get("shadowColor").String()
}

func (p *canvasRenderingContext2DImpl) SetShadowColor(color string) {
	p.Set("shadowColor", color)
}

func (p *canvasRenderingContext2DImpl) ClearRect(x float64, y float64, w float64, h float64) {
	p.Call("clearRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) FillRect(x float64, y float64, w float64, h float64) {
	p.Call("fillRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) StrokeRect(x float64, y float64, w float64, h float64) {
	p.Call("strokeRect", x, y, w, h)
}

func (p *canvasRenderingContext2DImpl) BeginPath() {
	p.Call("beginPath")
}

func (p *canvasRenderingContext2DImpl) Fill() {
	p.Call("fill")
}

func (p *canvasRenderingContext2DImpl) Stroke() {
	p.Call("stroke")
}

func (p *canvasRenderingContext2DImpl) DrawFocusIfNeeded(elm Element) {
	p.Call("drawFocusIfNeeded", elm.JSValue())
}

func (p *canvasRenderingContext2DImpl) Clip() {
	p.Call("clip")
}

func (p *canvasRenderingContext2DImpl) IsPointInPath(x float64, y float64) bool {
	return p.Call("isPointInPath", x, y).Bool()
}

func (p *canvasRenderingContext2DImpl) FillText(text string, x float64, y float64, maxWidth ...float64) {
	switch len(maxWidth) {
	case 0:
		p.Call("fillText", text, x, y)
	default:
		p.Call("fillText", text, x, y, maxWidth[0])
	}
}

func (p *canvasRenderingContext2DImpl) StrokeText(text string, x float64, y float64, maxWidth ...float64) {
	switch len(maxWidth) {
	case 0:
		p.Call("strokeText", text, x, y)
	default:
		p.Call("strokeText", text, x, y, maxWidth[0])
	}
}

func (p *canvasRenderingContext2DImpl) MeasureText(text string) TextMetrics {
	return wrapTextMetrics(p.Call("measureText", text))
}

func (p *canvasRenderingContext2DImpl) DrawImage(image CanvasImageSource, args ...float64) {
	switch len(args) {
	case 2:
		p.Call("drawImage", image.JSValue(), args[0], args[1])
	case 4:
		p.Call("drawImage", image.JSValue(), args[0], args[1], args[2], args[3])
	case 8:
		p.Call("drawImage", image.JSValue(), args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	}
}

func (p *canvasRenderingContext2DImpl) AddHitRegion(options HitRegionOptions) {
	p.Call("addHitRegion", options.toDict())
}

func (p *canvasRenderingContext2DImpl) RemoveHitRegion(id string) {
	p.Call("removeHitRegion", id)
}

func (p *canvasRenderingContext2DImpl) ClearHitRegions() {
	p.Call("clearHitRegions")
}

func (p *canvasRenderingContext2DImpl) CreateImageData(sw float64, sh float64) ImageData {
	return wrapImageData(p.Call("createImageData", sw, sh))
}

func (p *canvasRenderingContext2DImpl) CreateImageDataFromImageData(imageData ImageData) ImageData {
	return wrapImageData(p.Call("createImageData", imageData.JSValue()))
}

func (p *canvasRenderingContext2DImpl) ImageData(sx float64, sy float64, sw float64, sh float64) ImageData {
	return wrapImageData(p.Call("getImageData", sx, sy, sw, sh))
}

func (p *canvasRenderingContext2DImpl) PutImageData(imageData ImageData, dx float64, dy float64) {
	p.Call("putImageData", imageData.JSValue(), dx, dy)
}

func (p *canvasRenderingContext2DImpl) PutImageDataDirty(imageData ImageData, dx float64, dy float64, dirtyX float64, dirtyY float64, dirtyWidth float64, dirtyHeight float64) {
	p.Call("putImageData", imageData.JSValue(), dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

// -------------8<---------------------------------------

type canvasDrawingStylesImpl struct {
	js.Value
}

func wrapCanvasDrawingStyles(v js.Value) CanvasDrawingStyles {
	if p := newCanvasDrawingStylesImpl(v); p != nil {
		return p
	}
	return nil
}

func newCanvasDrawingStylesImpl(v js.Value) *canvasDrawingStylesImpl {
	if isNil(v) {
		return nil
	}

	return &canvasDrawingStylesImpl{
		Value: v,
	}
}

func (p *canvasDrawingStylesImpl) LineWidth() float64 {
	return p.Get("lineWidth").Float()
}

func (p *canvasDrawingStylesImpl) SetLineWidth(w float64) {
	p.Set("lineWidth", w)
}

func (p *canvasDrawingStylesImpl) LineCap() string {
	return p.Get("lineCap").String()
}

func (p *canvasDrawingStylesImpl) SetLineCap(c string) {
	p.Set("lineCap", c)
}

func (p *canvasDrawingStylesImpl) LineJoin() string {
	return p.Get("lineJoin").String()
}

func (p *canvasDrawingStylesImpl) SetLineJoin(j string) {
	p.Set("lineJoin", j)
}

func (p *canvasDrawingStylesImpl) MiterLimit() float64 {
	return p.Get("miterLimit").Float()
}

func (p *canvasDrawingStylesImpl) SetMiterLimit(l float64) {
	p.Set("miterLimit", l)
}

func (p *canvasDrawingStylesImpl) LineDash() []float64 {
	return floatSequenceToSlice(p.Call("getLineDash"))
}

func (p *canvasDrawingStylesImpl) SetLineDash(arg ...float64) {
	// TODO check this
	p.Call("setLineDash", sliceToJsArray(arg))
}

func (p *canvasDrawingStylesImpl) LineDashOffset() float64 {
	return p.Get("lineDashOffset").Float()
}

func (p *canvasDrawingStylesImpl) Font() string {
	return p.Get("font").String()
}

func (p *canvasDrawingStylesImpl) SetFont(font string) {
	p.Set("font", font)
}

func (p *canvasDrawingStylesImpl) TextAlign() TextAlign {
	return TextAlign(p.Get("textAlign").String())
}

func (p *canvasDrawingStylesImpl) SetTextAlign(ta TextAlign) {
	p.Set("textAlign", string(ta))
}

func (p *canvasDrawingStylesImpl) TextBaseline() TextBaseline {
	return TextBaseline(p.Get("textBaseline").String())
}

func (p *canvasDrawingStylesImpl) SetTextBaseline(tbl TextBaseline) {
	p.Set("textBaseline", string(tbl))
}

// -------------8<---------------------------------------

type canvasPathMethodsImpl struct {
	js.Value
}

func wrapCanvasPathMethods(v js.Value) CanvasPathMethods {
	if p := newCanvasPathMethodsImpl(v); p != nil {
		return p
	}
	return nil
}

func newCanvasPathMethodsImpl(v js.Value) *canvasPathMethodsImpl {
	if isNil(v) {
		return nil
	}

	return &canvasPathMethodsImpl{
		Value: v,
	}
}

func (p *canvasPathMethodsImpl) ClosePath() {
	p.Call("closePath")
}

func (p *canvasPathMethodsImpl) MoveTo(x float64, y float64) {
	p.Call("moveTo", x, y)
}

func (p *canvasPathMethodsImpl) LineTo(x float64, y float64) {
	p.Call("lineTo", x, y)
}

func (p *canvasPathMethodsImpl) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) {
	p.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (p *canvasPathMethodsImpl) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	p.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (p *canvasPathMethodsImpl) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	p.Call("arcTo", x1, y1, x2, y2, radius)
}

func (p *canvasPathMethodsImpl) Rect(x float64, y float64, w float64, h float64) {
	p.Call("rect", x, y, w, h)
}

func (p *canvasPathMethodsImpl) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterClockwise ...bool) {
	switch len(counterClockwise) {
	case 0:
		p.Call("arc", x, y, radius, startAngle, endAngle)
	default:
		p.Call("arc", x, y, radius, startAngle, endAngle, counterClockwise[0])
	}
}

// -------------8<---------------------------------------

type canvasGradientImpl struct {
	js.Value
}

func wrapCanvasGradient(v js.Value) CanvasGradient {
	if isNil(v) {
		return nil
	}

	return &canvasGradientImpl{
		Value: v,
	}
}

func (p *canvasGradientImpl) AddColorStop(offset float64, color string) {
	p.Call("addColorStop", offset, color)
}

// -------------8<---------------------------------------

type canvasPatternImpl struct {
	js.Value
}

func wrapCanvasPattern(v js.Value) CanvasPattern {
	if isNil(v) {
		return nil
	}

	return &canvasPatternImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type textMetricsImpl struct {
	js.Value
}

func wrapTextMetrics(v js.Value) TextMetrics {
	if isNil(v) {
		return nil
	}

	return &textMetricsImpl{
		Value: v,
	}
}

func (p *textMetricsImpl) Width() float64 {
	return p.Get("width").Float()
}

// -------------8<---------------------------------------

type imageDataImpl struct {
	js.Value
}

func wrapImageData(v js.Value) ImageData {
	if isNil(v) {
		return nil
	}

	return &imageDataImpl{
		Value: v,
	}
}

func (p *imageDataImpl) Width() int {
	return p.Get("width").Int()
}

func (p *imageDataImpl) Height() int {
	return p.Get("height").Int()
}

func (p *imageDataImpl) Data() []byte {
	// TODO  Uint8ClampedArray
	return nil
}

// -------------8<---------------------------------------
