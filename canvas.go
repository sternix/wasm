// +build js,wasm

package wasm

type (
	// typedef (HTMLImageElement or HTMLVideoElement or HTMLCanvasElement) CanvasImageSource;
	CanvasImageSource interface{}

	// https://www.w3.org/TR/2dcontext/#canvasrenderingcontext2d
	CanvasRenderingContext2D interface {
		CanvasDrawingStyles
		CanvasPathMethods

		Canvas() HTMLCanvasElement
		Save()
		Restore()
		Scale(float64, float64)
		Rotate(float64)
		Translate(float64, float64)
		Transform(float64, float64, float64, float64, float64, float64)
		SetTransform(float64, float64, float64, float64, float64, float64)
		GlobalAlpha() float64
		SetGlobalAlpha(float64)
		GlobalCompositeOperation() string
		SetGlobalCompositeOperation(string)
		StrokeStyle() interface{}
		SetStrokeStyle(interface{})
		FillStyle() interface{}
		SetFillStyle(interface{})
		CreateLinearGradient(float64, float64, float64, float64) CanvasGradient
		CreateRadialGradient(float64, float64, float64, float64, float64, float64) CanvasGradient
		CreatePattern(CanvasImageSource, string)
		ShadowOffsetX() float64
		SetShadowOffsetX(float64)
		ShadowOffsetY() float64
		SetShadowOffsetY(float64)
		ShadowBlur() float64
		SetShadowBlur(float64)
		ShadowColor() string
		SetShadowColor(string)
		ClearRect(float64, float64, float64, float64)
		FillRect(float64, float64, float64, float64)
		StrokeRect(float64, float64, float64, float64)
		BeginPath()
		Fill()
		Stroke()
		DrawFocusIfNeeded(Element)
		Clip()
		IsPointInPath(float64, float64) bool
		FillText(string, float64, float64, ...float64)
		StrokeText(string, float64, float64, ...float64)
		MeasureText(string) TextMetrics
		DrawImage(CanvasImageSource, ...float64) // it works if parameter len is 2,4,8
		AddHitRegion(HitRegionOptions)
		RemoveHitRegion(string)
		ClearHitRegions()
		CreateImageData(float64, float64) ImageData
		CreateImageDataFromImageData(ImageData) ImageData
		ImageData(float64, float64, float64, float64) ImageData
		PutImageData(ImageData, float64, float64)
		PutImageDataDirty(ImageData, float64, float64, float64, float64, float64, float64)
	}

	// https://www.w3.org/TR/2dcontext/#canvasdrawingstyles
	CanvasDrawingStyles interface {
		LineWidth() float64
		SetLineWidth(float64)
		LineCap() string
		SetLineCap(string)
		LineJoin() string
		SetLineJoin(string)
		MiterLimit() float64
		SetMiterLimit(float64)
		LineDash() []float64
		SetLineDash(...float64) // we can use []float64
		LineDashOffset() float64
		Font() string
		SetFont(string)
		TextAlign() TextAlign
		SetTextAlign(TextAlign)
		TextBaseline() TextBaseline
		SetTextBaseline(TextBaseline)
	}

	// https://www.w3.org/TR/2dcontext/#canvaspathmethods
	CanvasPathMethods interface {
		ClosePath()
		MoveTo(float64, float64)
		LineTo(float64, float64)
		QuadraticCurveTo(float64, float64, float64, float64)
		BezierCurveTo(float64, float64, float64, float64, float64, float64)
		ArcTo(float64, float64, float64, float64, float64)
		Rect(float64, float64, float64, float64)
		Arc(float64, float64, float64, float64, float64, ...bool)
	}

	// https://www.w3.org/TR/2dcontext/#canvasgradient
	CanvasGradient interface {
		AddColorStop(float64, string)
	}

	// https://www.w3.org/TR/2dcontext/#canvaspattern
	CanvasPattern interface{}

	// https://www.w3.org/TR/2dcontext/#textmetrics
	TextMetrics interface {
		Width() float64
	}

	// https://www.w3.org/TR/2dcontext/#imagedata
	ImageData interface {
		Width() int
		Height() int
		Data() []byte // TODO Uint8ClampedArray
	}
)

// https://www.w3.org/TR/2dcontext/#hitregionoptions
type HitRegionOptions struct {
	Id      string
	Control Element
}

func (p HitRegionOptions) JSValue() jsValue {
	o := jsObject.New()
	o.Set("id", p.Id)
	o.Set("control", JSValueOf(p.Control))
	return o
}

// helpers for known types
// https://www.w3.org/TR/2dcontext/#dom-context-2d-globalcompositeoperation

const (
	GlobalCompositeOperationSourceAtop      string = "source-atop"
	GlobalCompositeOperationSourceIn        string = "source-in"
	GlobalCompositeOperationSourceAout      string = "source-out"
	GlobalCompositeOperationSourceOver      string = "source-over"
	GlobalCompositeOperationDestinationAtop string = "destination-atop"
	GlobalCompositeOperationDestinationIn   string = "destination-in"
	GlobalCompositeOperationDestinationOut  string = "destination-out"
	GlobalCompositeOperationDestinationOver string = "destination-over"
	GlobalCompositeOperationLighter         string = "lighter"
	GlobalCompositeOperationCopy            string = "copy"
	GlobalCompositeOperationXor             string = "xor"
)

// https://www.w3.org/TR/2dcontext/#dom-context-2d-textalign

type TextAlign string

const (
	TextAlignStart  TextAlign = "start"
	TextAlignEnd    TextAlign = "end"
	TextAlignLeft   TextAlign = "left"
	TextAlignRight  TextAlign = "right"
	TextAlignCenter TextAlign = "center"
)

type TextBaseline string

const (
	TextBaselineTop         TextBaseline = "top"
	TextBaselineHanging     TextBaseline = "hanging"
	TextBaselineMiddle      TextBaseline = "middle"
	TextBaselineAlphabetic  TextBaseline = "alphabetic"
	TextBaselineIdeographic TextBaseline = "ideographic"
	TextBaselineBottom      TextBaseline = "bottom"
)
