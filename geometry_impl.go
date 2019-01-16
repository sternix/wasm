// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewDOMPoint(dpi ...DOMPointInit) DOMPoint {
	if jsDOMPoint := jsGlobal.get("DOMPoint"); jsDOMPoint.valid() {
		switch len(dpi) {
		case 0:
			return wrapDOMPoint(jsDOMPoint.jsNew())
		default:
			return wrapDOMPoint(jsDOMPoint.jsNew(dpi[0].JSValue()))
		}
	}
	return nil
}

func NewDOMPointReadOnly(x, y, z, w float64) DOMPointReadOnly {
	if jsDOMPointReadOnly := jsGlobal.get("DOMPointReadOnly"); jsDOMPointReadOnly.valid() {
		return wrapDOMPointReadOnly(jsDOMPointReadOnly.jsNew(x, y, z, w))
	}
	return nil
}

func NewDOMRectReadOnly(x, y, width, height float64) DOMRectReadOnly {
	if jsDOMRectReadOnly := jsGlobal.get("DOMRectReadOnly"); jsDOMRectReadOnly.valid() {
		return wrapDOMRectReadOnly(jsDOMRectReadOnly.jsNew(x, y, width, height))
	}
	return nil
}

func NewDOMRect(x, y, width, height float64) DOMRect {
	if jsDOMRect := jsGlobal.get("DOMRect"); jsDOMRect.valid() {
		return wrapDOMRect(jsDOMRect.jsNew(x, y, width, height))
	}
	return nil
}

func NewDOMQuad(dri ...DOMRectInit) DOMQuad {
	if jsDOMQuad := jsGlobal.get("DOMQuad"); jsDOMQuad.valid() {
		switch len(dri) {
		case 0:
			return wrapDOMQuad(jsDOMQuad.jsNew())
		default:
			return wrapDOMQuad(jsDOMQuad.jsNew(dri[0].JSValue()))
		}
	}
	return nil
}

func DOMQuadFromRect(other ...DOMRectInit) DOMQuad {
	switch len(other) {
	case 0:
		return wrapDOMQuad(jsGlobal.invoke("DOMQuad.fromRect"))
	default:
		return wrapDOMQuad(jsGlobal.invoke("DOMQuad.fromRect", other[0].JSValue()))
	}
}

func DOMQuadFromQuad(other ...DOMQuadInit) DOMQuad {
	switch len(other) {
	case 0:
		return wrapDOMQuad(jsGlobal.invoke("DOMQuad.fromQuad"))
	default:
		return wrapDOMQuad(jsGlobal.invoke("DOMQuad.fromQuad", other[0].JSValue()))
	}
}

// TODO: check this
func NewDOMMatrixReadOnly(numberSequence []float64) DOMMatrixReadOnly {
	if jsDOMMatrixReadOnly := jsGlobal.get("DOMMatrixReadOnly"); jsDOMMatrixReadOnly.valid() {
		ta := jsTypedArrayOf(numberSequence)
		defer ta.Release()
		return wrapDOMMatrixReadOnly(jsDOMMatrixReadOnly.jsNew(ta))
	}
	return nil
}

// TODO: check this
func NewDOMMatrix(numberSequence []float64) DOMMatrix {
	if jsDOMMatrix := jsGlobal.get("DOMMatrix"); jsDOMMatrix.valid() {
		ta := jsTypedArrayOf(numberSequence)
		defer ta.Release()
		return wrapDOMMatrix(jsDOMMatrix.jsNew(ta))
	}
	return nil
}

// -------------8<---------------------------------------

type domPointReadOnlyImpl struct {
	Value
}

func wrapDOMPointReadOnly(v Value) DOMPointReadOnly {
	if p := newDOMPointReadOnlyImpl(v); p != nil {
		return p
	}
	return nil
}

func newDOMPointReadOnlyImpl(v Value) *domPointReadOnlyImpl {
	if v.valid() {
		return &domPointReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domPointReadOnlyImpl) X() float64 {
	return p.get("x").toFloat64()
}

func (p *domPointReadOnlyImpl) Y() float64 {
	return p.get("y").toFloat64()
}

func (p *domPointReadOnlyImpl) Z() float64 {
	return p.get("z").toFloat64()
}

func (p *domPointReadOnlyImpl) W() float64 {
	return p.get("w").toFloat64()
}

func (p *domPointReadOnlyImpl) MatrixTransform(matrix DOMMatrixReadOnly) DOMPoint {
	return wrapDOMPoint(p.call("matrixTransform", JSValueOf(matrix)))
}

// -------------8<---------------------------------------

type domPointImpl struct {
	*domPointReadOnlyImpl
}

func wrapDOMPoint(v Value) DOMPoint {
	if v.valid() {
		return &domPointImpl{
			domPointReadOnlyImpl: newDOMPointReadOnlyImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type domRectImpl struct {
	*domRectReadOnlyImpl
}

func wrapDOMRect(v Value) DOMRect {
	if v.valid() {
		return &domRectImpl{
			domRectReadOnlyImpl: newDOMRectReadOnlyImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type domRectReadOnlyImpl struct {
	Value
}

func wrapDOMRectReadOnly(v Value) DOMRectReadOnly {
	if p := newDOMRectReadOnlyImpl(v); p != nil {
		return p
	}
	return nil
}

func newDOMRectReadOnlyImpl(v Value) *domRectReadOnlyImpl {
	if v.valid() {
		return &domRectReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domRectReadOnlyImpl) X() float64 {
	return p.get("x").toFloat64()
}

func (p *domRectReadOnlyImpl) Y() float64 {
	return p.get("y").toFloat64()
}

func (p *domRectReadOnlyImpl) Width() float64 {
	return p.get("width").toFloat64()
}

func (p *domRectReadOnlyImpl) Height() float64 {
	return p.get("height").toFloat64()
}

func (p *domRectReadOnlyImpl) Top() float64 {
	return p.get("top").toFloat64()
}

func (p *domRectReadOnlyImpl) Right() float64 {
	return p.get("right").toFloat64()
}

func (p *domRectReadOnlyImpl) Bottom() float64 {
	return p.get("bottom").toFloat64()
}

func (p *domRectReadOnlyImpl) Left() float64 {
	return p.get("left").toFloat64()
}

// -------------8<---------------------------------------

type domQuadImpl struct {
	Value
}

func wrapDOMQuad(v Value) DOMQuad {
	if v.valid() {
		return &domQuadImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domQuadImpl) P1() DOMPoint {
	return wrapDOMPoint(p.get("p1"))
}

func (p *domQuadImpl) P2() DOMPoint {
	return wrapDOMPoint(p.get("p2"))
}

func (p *domQuadImpl) P3() DOMPoint {
	return wrapDOMPoint(p.get("p3"))
}

func (p *domQuadImpl) P4() DOMPoint {
	return wrapDOMPoint(p.get("p4"))
}

func (p *domQuadImpl) Bounds() DOMRectReadOnly {
	return wrapDOMRectReadOnly(p.get("bounds"))
}

// -------------8<---------------------------------------

type domMatrixReadOnlyImpl struct {
	Value
}

func wrapDOMMatrixReadOnly(v Value) DOMMatrixReadOnly {
	if p := newDOMMatrixReadOnlyImpl(v); p != nil {
		return p
	}
	return nil
}

func newDOMMatrixReadOnlyImpl(v Value) *domMatrixReadOnlyImpl {
	if v.valid() {
		return &domMatrixReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domMatrixReadOnlyImpl) A() float64 {
	return p.get("a").toFloat64()
}

func (p *domMatrixReadOnlyImpl) B() float64 {
	return p.get("b").toFloat64()
}

func (p *domMatrixReadOnlyImpl) C() float64 {
	return p.get("c").toFloat64()
}

func (p *domMatrixReadOnlyImpl) D() float64 {
	return p.get("d").toFloat64()
}

func (p *domMatrixReadOnlyImpl) E() float64 {
	return p.get("e").toFloat64()
}

func (p *domMatrixReadOnlyImpl) F() float64 {
	return p.get("f").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M11() float64 {
	return p.get("m11").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M12() float64 {
	return p.get("m12").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M13() float64 {
	return p.get("m13").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M14() float64 {
	return p.get("m14").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M21() float64 {
	return p.get("m21").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M22() float64 {
	return p.get("m22").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M23() float64 {
	return p.get("m23").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M24() float64 {
	return p.get("m24").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M31() float64 {
	return p.get("m31").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M32() float64 {
	return p.get("m32").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M33() float64 {
	return p.get("m33").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M34() float64 {
	return p.get("m34").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M41() float64 {
	return p.get("m41").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M42() float64 {
	return p.get("m42").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M43() float64 {
	return p.get("m43").toFloat64()
}

func (p *domMatrixReadOnlyImpl) M44() float64 {
	return p.get("m44").toFloat64()
}

func (p *domMatrixReadOnlyImpl) Is2D() bool {
	return p.get("is2D").toBool()
}

func (p *domMatrixReadOnlyImpl) IsIdentity() bool {
	return p.get("isIdentity").toBool()
}

func (p *domMatrixReadOnlyImpl) Translate(tx float64, ty float64, tz ...float64) DOMMatrix {
	switch len(tz) {
	case 0:
		return wrapDOMMatrix(p.call("translate", tx, ty))
	default:
		return wrapDOMMatrix(p.call("translate", tx, ty, tz[0]))
	}
}

func (p *domMatrixReadOnlyImpl) Scale(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scale", scale))
	case 1:
		return wrapDOMMatrix(p.call("scale", scale, args[0]))
	default:
		return wrapDOMMatrix(p.call("scale", scale, args[0], args[1]))
	}
}

func (p *domMatrixReadOnlyImpl) Scale3d(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scale3d", scale))
	case 1:
		return wrapDOMMatrix(p.call("scale3d", scale, args[0]))
	case 2:
		return wrapDOMMatrix(p.call("scale3d", scale, args[0], args[1]))
	default:
		return wrapDOMMatrix(p.call("scale3d", scale, args[0], args[1], args[2]))
	}
}

func (p *domMatrixReadOnlyImpl) ScaleNonUniform(scaleX float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX))
	case 1:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX, args[0]))
	case 2:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX, args[0], args[1]))
	case 3:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX, args[0], args[1], args[2]))
	case 4:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX, args[0], args[1], args[2], args[3]))
	default:
		return wrapDOMMatrix(p.call("scaleNonUniform", scaleX, args[0], args[1], args[2], args[3], args[4]))
	}
}

func (p *domMatrixReadOnlyImpl) Rotate(angle float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("rotate", angle))
	case 1:
		return wrapDOMMatrix(p.call("rotate", angle, args[0]))
	default:
		return wrapDOMMatrix(p.call("rotate", angle, args[0], args[1]))
	}
}

func (p *domMatrixReadOnlyImpl) RotateFromVector(x float64, y float64) DOMMatrix {
	return wrapDOMMatrix(p.call("rotateFromVector", x, y))
}

func (p *domMatrixReadOnlyImpl) RotateAxisAngle(x float64, y float64, z float64, angle float64) DOMMatrix {
	return wrapDOMMatrix(p.call("rotateAxisAngle", x, y, z, angle))
}

func (p *domMatrixReadOnlyImpl) SkewX(sx float64) DOMMatrix {
	return wrapDOMMatrix(p.call("skewX", sx))
}

func (p *domMatrixReadOnlyImpl) SkewY(sy float64) DOMMatrix {
	return wrapDOMMatrix(p.call("skewY", sy))
}

func (p *domMatrixReadOnlyImpl) Multiply(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.call("multiply", JSValueOf(other)))
}

func (p *domMatrixReadOnlyImpl) FlipX() DOMMatrix {
	return wrapDOMMatrix(p.call("flipX"))
}

func (p *domMatrixReadOnlyImpl) FlipY() DOMMatrix {
	return wrapDOMMatrix(p.call("flipY"))
}

func (p *domMatrixReadOnlyImpl) Inverse() DOMMatrix {
	return wrapDOMMatrix(p.call("inverse"))
}

func (p *domMatrixReadOnlyImpl) TransformPoint(point ...DOMPointInit) DOMPoint {
	switch len(point) {
	case 0:
		return wrapDOMPoint(p.call("transformPoint"))
	default:
		return wrapDOMPoint(p.call("transformPoint", point[0].JSValue()))
	}
}

func (p *domMatrixReadOnlyImpl) ToFloat32Array() []float32 {
	return toFloat32Slice(p.call("toFloat32Array"))
}

func (p *domMatrixReadOnlyImpl) ToFloat64Array() []float64 {
	return toFloat64Slice(p.call("toFloat64Array"))
}

func (p *domMatrixReadOnlyImpl) String() string {
	return p.call("toString").toString()
}

// -------------8<---------------------------------------

type domMatrixImpl struct {
	*domMatrixReadOnlyImpl
}

func wrapDOMMatrix(v Value) DOMMatrix {
	if v.valid() {
		return &domMatrixImpl{
			domMatrixReadOnlyImpl: newDOMMatrixReadOnlyImpl(v),
		}
	}
	return nil
}

func (p *domMatrixImpl) MultiplySelf(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.call("multiplySelf", JSValueOf(other)))
}

func (p *domMatrixImpl) PreMultiplySelf(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.call("preMultiplySelf", JSValueOf(other)))
}

func (p *domMatrixImpl) TranslateSelf(tx float64, ty float64, tz ...float64) DOMMatrix {
	switch len(tz) {
	case 0:
		return wrapDOMMatrix(p.call("translateSelf", tx, ty))
	default:
		return wrapDOMMatrix(p.call("translateSelf", tx, ty, tz[0]))
	}
}

func (p *domMatrixImpl) ScaleSelf(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scaleSelf", scale))
	case 1:
		return wrapDOMMatrix(p.call("scaleSelf", scale, args[0]))
	default:
		return wrapDOMMatrix(p.call("scaleSelf", scale, args[0], args[1]))
	}
}

func (p *domMatrixImpl) Scale3dSelf(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scale3dSelf", scale))
	case 1:
		return wrapDOMMatrix(p.call("scale3dSelf", scale, args[0]))
	case 2:
		return wrapDOMMatrix(p.call("scale3dSelf", scale, args[0], args[1]))
	default:
		return wrapDOMMatrix(p.call("scale3dSelf", scale, args[0], args[1], args[2]))
	}
}

func (p *domMatrixImpl) ScaleNonUniformSelf(scaleX float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX))
	case 1:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX, args[0]))
	case 2:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX, args[0], args[1]))
	case 3:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2]))
	case 4:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2], args[3]))
	default:
		return wrapDOMMatrix(p.call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2], args[3], args[4]))
	}
}

func (p *domMatrixImpl) RotateSelf(angle float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.call("rotateSelf", angle))
	case 1:
		return wrapDOMMatrix(p.call("rotateSelf", angle, args[0]))
	default:
		return wrapDOMMatrix(p.call("rotateSelf", angle, args[0], args[1]))
	}
}

func (p *domMatrixImpl) RotateFromVectorSelf(x float64, y float64) DOMMatrix {
	return wrapDOMMatrix(p.call("rotateFromVectorSelf", x, y))
}

func (p *domMatrixImpl) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) DOMMatrix {
	return wrapDOMMatrix(p.call("rotateAxisAngleSelf", x, y, z, angle))
}

func (p *domMatrixImpl) SkewXSelf(sx float64) DOMMatrix {
	return wrapDOMMatrix(p.call("skewXSelf", sx))
}

func (p *domMatrixImpl) SkewYSelf(sy float64) DOMMatrix {
	return wrapDOMMatrix(p.call("skewYSelf", sy))
}

func (p *domMatrixImpl) InverseSelf() DOMMatrix {
	return wrapDOMMatrix(p.call("invertSelf"))
}

func (p *domMatrixImpl) SetMatrixValue(transformList string) DOMMatrix {
	return wrapDOMMatrix(p.call("setMatrixValue", transformList))
}

// -------------8<---------------------------------------
