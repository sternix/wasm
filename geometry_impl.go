// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewDOMPoint(dpi ...DOMPointInit) DOMPoint {
	if jsDOMPoint := jsGlobal.Get("DOMPoint"); jsDOMPoint.Valid() {
		switch len(dpi) {
		case 0:
			return wrapDOMPoint(jsDOMPoint.New())
		default:
			return wrapDOMPoint(jsDOMPoint.New(dpi[0].toJSObject()))
		}
	}
	return nil
}

func NewDOMPointReadOnly(x, y, z, w float64) DOMPointReadOnly {
	if jsDOMPointReadOnly := jsGlobal.Get("DOMPointReadOnly"); jsDOMPointReadOnly.Valid() {
		return wrapDOMPointReadOnly(jsDOMPointReadOnly.New(x, y, z, w))
	}
	return nil
}

func NewDOMRectReadOnly(x, y, width, height float64) DOMRectReadOnly {
	if jsDOMRectReadOnly := jsGlobal.Get("DOMRectReadOnly"); jsDOMRectReadOnly.Valid() {
		return wrapDOMRectReadOnly(jsDOMRectReadOnly.New(x, y, width, height))
	}
	return nil
}

func NewDOMRect(x, y, width, height float64) DOMRect {
	if jsDOMRect := jsGlobal.Get("DOMRect"); jsDOMRect.Valid() {
		return wrapDOMRect(jsDOMRect.New(x, y, width, height))
	}
	return nil
}

func NewDOMQuad(dri ...DOMRectInit) DOMQuad {
	if jsDOMQuad := jsGlobal.Get("DOMQuad"); jsDOMQuad.Valid() {
		switch len(dri) {
		case 0:
			return wrapDOMQuad(jsDOMQuad.New())
		default:
			return wrapDOMQuad(jsDOMQuad.New(dri[0].toJSObject()))
		}
	}
	return nil
}

func DOMQuadFromRect(other ...DOMRectInit) DOMQuad {
	switch len(other) {
	case 0:
		return wrapDOMQuad(jsGlobal.Invoke("DOMQuad.fromRect"))
	default:
		return wrapDOMQuad(jsGlobal.Invoke("DOMQuad.fromRect", other[0].toJSObject()))
	}
}

func DOMQuadFromQuad(other ...DOMQuadInit) DOMQuad {
	switch len(other) {
	case 0:
		return wrapDOMQuad(jsGlobal.Invoke("DOMQuad.fromQuad"))
	default:
		return wrapDOMQuad(jsGlobal.Invoke("DOMQuad.fromQuad", other[0].toJSObject()))
	}
}

// TODO: check this
func NewDOMMatrixReadOnly(numberSequence []float64) DOMMatrixReadOnly {
	if jsDOMMatrixReadOnly := jsGlobal.Get("DOMMatrixReadOnly"); jsDOMMatrixReadOnly.Valid() {
		var param []interface{}
		for _, n := range numberSequence {
			param = append(param, n)
		}
		return wrapDOMMatrixReadOnly(jsDOMMatrixReadOnly.New(param))
	}
	return nil
}

// TODO: check this
func NewDOMMatrix(numberSequence []float64) DOMMatrix {
	if jsDOMMatrix := jsGlobal.Get("DOMMatrix"); jsDOMMatrix.Valid() {
		var param []interface{}
		for _, n := range numberSequence {
			param = append(param, n)
		}
		return wrapDOMMatrix(jsDOMMatrix.New(param))
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
	if v.Valid() {
		return &domPointReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domPointReadOnlyImpl) X() float64 {
	return p.Get("x").Float()
}

func (p *domPointReadOnlyImpl) Y() float64 {
	return p.Get("y").Float()
}

func (p *domPointReadOnlyImpl) Z() float64 {
	return p.Get("z").Float()
}

func (p *domPointReadOnlyImpl) W() float64 {
	return p.Get("w").Float()
}

func (p *domPointReadOnlyImpl) MatrixTransform(matrix DOMMatrixReadOnly) DOMPoint {
	return wrapDOMPoint(p.Call("matrixTransform", JSValue(matrix)))
}

// -------------8<---------------------------------------

type domPointImpl struct {
	*domPointReadOnlyImpl
}

func wrapDOMPoint(v Value) DOMPoint {
	if v.Valid() {
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
	if v.Valid() {
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
	if v.Valid() {
		return &domRectReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domRectReadOnlyImpl) X() float64 {
	return p.Get("x").Float()
}

func (p *domRectReadOnlyImpl) Y() float64 {
	return p.Get("y").Float()
}

func (p *domRectReadOnlyImpl) Width() float64 {
	return p.Get("width").Float()
}

func (p *domRectReadOnlyImpl) Height() float64 {
	return p.Get("height").Float()
}

func (p *domRectReadOnlyImpl) Top() float64 {
	return p.Get("top").Float()
}

func (p *domRectReadOnlyImpl) Right() float64 {
	return p.Get("right").Float()
}

func (p *domRectReadOnlyImpl) Bottom() float64 {
	return p.Get("bottom").Float()
}

func (p *domRectReadOnlyImpl) Left() float64 {
	return p.Get("left").Float()
}

// -------------8<---------------------------------------

type domQuadImpl struct {
	Value
}

func wrapDOMQuad(v Value) DOMQuad {
	if v.Valid() {
		return &domQuadImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domQuadImpl) P1() DOMPoint {
	return wrapDOMPoint(p.Get("p1"))
}

func (p *domQuadImpl) P2() DOMPoint {
	return wrapDOMPoint(p.Get("p2"))
}

func (p *domQuadImpl) P3() DOMPoint {
	return wrapDOMPoint(p.Get("p3"))
}

func (p *domQuadImpl) P4() DOMPoint {
	return wrapDOMPoint(p.Get("p4"))
}

func (p *domQuadImpl) Bounds() DOMRectReadOnly {
	return wrapDOMRectReadOnly(p.Get("bounds"))
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
	if v.Valid() {
		return &domMatrixReadOnlyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *domMatrixReadOnlyImpl) A() float64 {
	return p.Get("a").Float()
}

func (p *domMatrixReadOnlyImpl) B() float64 {
	return p.Get("b").Float()
}

func (p *domMatrixReadOnlyImpl) C() float64 {
	return p.Get("c").Float()
}

func (p *domMatrixReadOnlyImpl) D() float64 {
	return p.Get("d").Float()
}

func (p *domMatrixReadOnlyImpl) E() float64 {
	return p.Get("e").Float()
}

func (p *domMatrixReadOnlyImpl) F() float64 {
	return p.Get("f").Float()
}

func (p *domMatrixReadOnlyImpl) M11() float64 {
	return p.Get("m11").Float()
}

func (p *domMatrixReadOnlyImpl) M12() float64 {
	return p.Get("m12").Float()
}

func (p *domMatrixReadOnlyImpl) M13() float64 {
	return p.Get("m13").Float()
}

func (p *domMatrixReadOnlyImpl) M14() float64 {
	return p.Get("m14").Float()
}

func (p *domMatrixReadOnlyImpl) M21() float64 {
	return p.Get("m21").Float()
}

func (p *domMatrixReadOnlyImpl) M22() float64 {
	return p.Get("m22").Float()
}

func (p *domMatrixReadOnlyImpl) M23() float64 {
	return p.Get("m23").Float()
}

func (p *domMatrixReadOnlyImpl) M24() float64 {
	return p.Get("m24").Float()
}

func (p *domMatrixReadOnlyImpl) M31() float64 {
	return p.Get("m31").Float()
}

func (p *domMatrixReadOnlyImpl) M32() float64 {
	return p.Get("m32").Float()
}

func (p *domMatrixReadOnlyImpl) M33() float64 {
	return p.Get("m33").Float()
}

func (p *domMatrixReadOnlyImpl) M34() float64 {
	return p.Get("m34").Float()
}

func (p *domMatrixReadOnlyImpl) M41() float64 {
	return p.Get("m41").Float()
}

func (p *domMatrixReadOnlyImpl) M42() float64 {
	return p.Get("m42").Float()
}

func (p *domMatrixReadOnlyImpl) M43() float64 {
	return p.Get("m43").Float()
}

func (p *domMatrixReadOnlyImpl) M44() float64 {
	return p.Get("m44").Float()
}

func (p *domMatrixReadOnlyImpl) Is2D() bool {
	return p.Get("is2D").Bool()
}

func (p *domMatrixReadOnlyImpl) IsIdentity() bool {
	return p.Get("isIdentity").Bool()
}

func (p *domMatrixReadOnlyImpl) Translate(tx float64, ty float64, tz ...float64) DOMMatrix {
	switch len(tz) {
	case 0:
		return wrapDOMMatrix(p.Call("translate", tx, ty))
	default:
		return wrapDOMMatrix(p.Call("translate", tx, ty, tz[0]))
	}
}

func (p *domMatrixReadOnlyImpl) Scale(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scale", scale))
	case 1:
		return wrapDOMMatrix(p.Call("scale", scale, args[0]))
	default:
		return wrapDOMMatrix(p.Call("scale", scale, args[0], args[1]))
	}
}

func (p *domMatrixReadOnlyImpl) Scale3d(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scale3d", scale))
	case 1:
		return wrapDOMMatrix(p.Call("scale3d", scale, args[0]))
	case 2:
		return wrapDOMMatrix(p.Call("scale3d", scale, args[0], args[1]))
	default:
		return wrapDOMMatrix(p.Call("scale3d", scale, args[0], args[1], args[2]))
	}
}

func (p *domMatrixReadOnlyImpl) ScaleNonUniform(scaleX float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX))
	case 1:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX, args[0]))
	case 2:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX, args[0], args[1]))
	case 3:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX, args[0], args[1], args[2]))
	case 4:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX, args[0], args[1], args[2], args[3]))
	default:
		return wrapDOMMatrix(p.Call("scaleNonUniform", scaleX, args[0], args[1], args[2], args[3], args[4]))
	}
}

func (p *domMatrixReadOnlyImpl) Rotate(angle float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("rotate", angle))
	case 1:
		return wrapDOMMatrix(p.Call("rotate", angle, args[0]))
	default:
		return wrapDOMMatrix(p.Call("rotate", angle, args[0], args[1]))
	}
}

func (p *domMatrixReadOnlyImpl) RotateFromVector(x float64, y float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("rotateFromVector", x, y))
}

func (p *domMatrixReadOnlyImpl) RotateAxisAngle(x float64, y float64, z float64, angle float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("rotateAxisAngle", x, y, z, angle))
}

func (p *domMatrixReadOnlyImpl) SkewX(sx float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("skewX", sx))
}

func (p *domMatrixReadOnlyImpl) SkewY(sy float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("skewY", sy))
}

func (p *domMatrixReadOnlyImpl) Multiply(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.Call("multiply", JSValue(other)))
}

func (p *domMatrixReadOnlyImpl) FlipX() DOMMatrix {
	return wrapDOMMatrix(p.Call("flipX"))
}

func (p *domMatrixReadOnlyImpl) FlipY() DOMMatrix {
	return wrapDOMMatrix(p.Call("flipY"))
}

func (p *domMatrixReadOnlyImpl) Inverse() DOMMatrix {
	return wrapDOMMatrix(p.Call("inverse"))
}

func (p *domMatrixReadOnlyImpl) TransformPoint(point ...DOMPointInit) DOMPoint {
	switch len(point) {
	case 0:
		return wrapDOMPoint(p.Call("transformPoint"))
	default:
		return wrapDOMPoint(p.Call("transformPoint", point[0].toJSObject()))
	}
}

func (p *domMatrixReadOnlyImpl) ToFloat32Array() []float32 {
	return toFloat32Slice(p.Call("toFloat32Array"))
}

func (p *domMatrixReadOnlyImpl) ToFloat64Array() []float64 {
	return toFloat64Slice(p.Call("toFloat64Array"))
}

func (p *domMatrixReadOnlyImpl) String() string {
	return p.Call("toString").String()
}

// -------------8<---------------------------------------

type domMatrixImpl struct {
	*domMatrixReadOnlyImpl
}

func wrapDOMMatrix(v Value) DOMMatrix {
	if v.Valid() {
		return &domMatrixImpl{
			domMatrixReadOnlyImpl: newDOMMatrixReadOnlyImpl(v),
		}
	}
	return nil
}

func (p *domMatrixImpl) MultiplySelf(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.Call("multiplySelf", JSValue(other)))
}

func (p *domMatrixImpl) PreMultiplySelf(other DOMMatrix) DOMMatrix {
	return wrapDOMMatrix(p.Call("preMultiplySelf", JSValue(other)))
}

func (p *domMatrixImpl) TranslateSelf(tx float64, ty float64, tz ...float64) DOMMatrix {
	switch len(tz) {
	case 0:
		return wrapDOMMatrix(p.Call("translateSelf", tx, ty))
	default:
		return wrapDOMMatrix(p.Call("translateSelf", tx, ty, tz[0]))
	}
}

func (p *domMatrixImpl) ScaleSelf(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scaleSelf", scale))
	case 1:
		return wrapDOMMatrix(p.Call("scaleSelf", scale, args[0]))
	default:
		return wrapDOMMatrix(p.Call("scaleSelf", scale, args[0], args[1]))
	}
}

func (p *domMatrixImpl) Scale3dSelf(scale float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scale3dSelf", scale))
	case 1:
		return wrapDOMMatrix(p.Call("scale3dSelf", scale, args[0]))
	case 2:
		return wrapDOMMatrix(p.Call("scale3dSelf", scale, args[0], args[1]))
	default:
		return wrapDOMMatrix(p.Call("scale3dSelf", scale, args[0], args[1], args[2]))
	}
}

func (p *domMatrixImpl) ScaleNonUniformSelf(scaleX float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX))
	case 1:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX, args[0]))
	case 2:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX, args[0], args[1]))
	case 3:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2]))
	case 4:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2], args[3]))
	default:
		return wrapDOMMatrix(p.Call("scaleNonUniformSelf", scaleX, args[0], args[1], args[2], args[3], args[4]))
	}
}

func (p *domMatrixImpl) RotateSelf(angle float64, args ...float64) DOMMatrix {
	switch len(args) {
	case 0:
		return wrapDOMMatrix(p.Call("rotateSelf", angle))
	case 1:
		return wrapDOMMatrix(p.Call("rotateSelf", angle, args[0]))
	default:
		return wrapDOMMatrix(p.Call("rotateSelf", angle, args[0], args[1]))
	}
}

func (p *domMatrixImpl) RotateFromVectorSelf(x float64, y float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("rotateFromVectorSelf", x, y))
}

func (p *domMatrixImpl) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("rotateAxisAngleSelf", x, y, z, angle))
}

func (p *domMatrixImpl) SkewXSelf(sx float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("skewXSelf", sx))
}

func (p *domMatrixImpl) SkewYSelf(sy float64) DOMMatrix {
	return wrapDOMMatrix(p.Call("skewYSelf", sy))
}

func (p *domMatrixImpl) InverseSelf() DOMMatrix {
	return wrapDOMMatrix(p.Call("invertSelf"))
}

func (p *domMatrixImpl) SetMatrixValue(transformList string) DOMMatrix {
	return wrapDOMMatrix(p.Call("setMatrixValue", transformList))
}

// -------------8<---------------------------------------
