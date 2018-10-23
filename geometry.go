// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://www.w3.org/TR/geometry-1/#idl-index

type (
	// https://www.w3.org/TR/geometry-1/#dom-dompointreadonly
	DOMPointReadOnly interface {
		X() float64
		Y() float64
		Z() float64
		W() float64
		MatrixTransform(DOMMatrixReadOnly) DOMPoint
	}

	// https://www.w3.org/TR/geometry-1/#dom-dompoint
	DOMPoint interface {
		DOMPointReadOnly
	}

	// https://www.w3.org/TR/geometry-1/#dictdef-dompointinit
	DOMPointInit struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"` //default 1
	}

	// https://www.w3.org/TR/geometry-1/#dom-domrect
	DOMRect interface {
		DOMRectReadOnly
	}

	// https://www.w3.org/TR/geometry-1/#dom-domrectreadonly
	DOMRectReadOnly interface {
		X() float64
		Y() float64
		Width() float64
		Height() float64
		Top() float64
		Right() float64
		Bottom() float64
		Left() float64
	}

	// https://www.w3.org/TR/geometry-1/#dictdef-domrectinit
	DOMRectInit struct {
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	}

	// https://www.w3.org/TR/geometry-1/#dom-domquad
	DOMQuad interface {
		P1() DOMPoint
		P2() DOMPoint
		P3() DOMPoint
		P4() DOMPoint
		Bounds() DOMRectReadOnly
	}

	// https://www.w3.org/TR/geometry-1/#dom-dommatrixreadonly
	DOMMatrixReadOnly interface {
		A() float64
		B() float64
		C() float64
		D() float64
		E() float64
		F() float64

		M11() float64
		M12() float64
		M13() float64
		M14() float64
		M21() float64
		M22() float64
		M23() float64
		M24() float64
		M31() float64
		M32() float64
		M33() float64
		M34() float64
		M41() float64
		M42() float64
		M43() float64
		M44() float64
		Is2D() bool
		IsIdentity() bool
		Translate(float64, float64, ...float64) DOMMatrix
		Scale(float64, ...float64) DOMMatrix
		Scale3d(float64, ...float64) DOMMatrix
		ScaleNonUniform(float64, ...float64) DOMMatrix
		Rotate(float64, ...float64) DOMMatrix
		RotateFromVector(float64, float64) DOMMatrix
		RotateAxisAngle(float64, float64, float64, float64) DOMMatrix
		SkewX(float64) DOMMatrix
		SkewY(float64) DOMMatrix
		Multiply(DOMMatrix) DOMMatrix
		FlipX() DOMMatrix
		FlipY() DOMMatrix
		Inverse() DOMMatrix
		TransformPoint(...DOMPointInit) DOMPoint
		ToFloat32Array() []float32 //Float32Array
		ToFloat64Array() []float64 // Float64Array
		JSValue() js.Value
		String() string
	}

	// https://www.w3.org/TR/geometry-1/#dom-dommatrix
	DOMMatrix interface {
		DOMMatrixReadOnly

		MultiplySelf(DOMMatrix) DOMMatrix
		PreMultiplySelf(DOMMatrix) DOMMatrix
		TranslateSelf(float64, float64, ...float64) DOMMatrix
		ScaleSelf(float64, ...float64) DOMMatrix
		Scale3dSelf(float64, ...float64) DOMMatrix
		ScaleNonUniformSelf(float64, ...float64) DOMMatrix
		RotateSelf(float64, ...float64) DOMMatrix
		RotateFromVectorSelf(float64, float64) DOMMatrix
		RotateAxisAngleSelf(float64, float64, float64, float64) DOMMatrix
		SkewXSelf(float64) DOMMatrix
		SkewYSelf(float64) DOMMatrix
		InverseSelf() DOMMatrix
		SetMatrixValue(string) DOMMatrix
	}

	// https://drafts.csswg.org/cssom-view/#dictdef-boxquadoptions
	BoxQuadOptions struct {
		Box        CSSBoxType   `json:"box"` // default "border"
		RelativeTo GeometryNode `json:"relativeTo"`
	}

	// https://drafts.csswg.org/cssom-view/#dictdef-convertcoordinateoptions
	ConvertCoordinateOptions struct {
		FromBox CSSBoxType `json:"fromBox"` // default "border"
		ToBox   CSSBoxType `json:"toBox"`   // default "border"
	}

	// typedef (Text or Element or CSSPseudoElement or Document) GeometryNode
	GeometryNode interface{}
)

type CSSBoxType string

const (
	CSSBoxTypeMargin  CSSBoxType = "margin"
	CSSBoxTypeBorder  CSSBoxType = "border"
	CSSBoxTypePadding CSSBoxType = "padding"
	CSSBoxTypeContent CSSBoxType = "content"
)
