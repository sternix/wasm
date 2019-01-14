// +build js,wasm

package wasm

// -------------8<---------------------------------------

type touchImpl struct {
	Value
}

func wrapTouch(v Value) Touch {
	if v.valid() {
		return &touchImpl{
			Value: v,
		}
	}
	return nil
}

func (p *touchImpl) Identifier() int {
	return p.get("identifier").toInt()
}

func (p *touchImpl) Target() EventTarget {
	return newEventTargetImpl(p.get("target"))
}

func (p *touchImpl) ScreenX() float64 {
	return p.get("screenX").toFloat64()
}

func (p *touchImpl) ScreenY() float64 {
	return p.get("screenY").toFloat64()
}

func (p *touchImpl) ClientX() float64 {
	return p.get("clientX").toFloat64()
}

func (p *touchImpl) ClientY() float64 {
	return p.get("clientY").toFloat64()
}

func (p *touchImpl) PageX() float64 {
	return p.get("pageX").toFloat64()
}

func (p *touchImpl) PageY() float64 {
	return p.get("pageY").toFloat64()
}

func (p *touchImpl) RadiusX() float64 {
	return p.get("radiusX").toFloat64()
}

func (p *touchImpl) RadiusY() float64 {
	return p.get("radiusY").toFloat64()
}

func (p *touchImpl) RotationAngle() float64 {
	return p.get("rotationAngle").toFloat64()
}

func (p *touchImpl) Force() float64 {
	return p.get("force").toFloat64()
}

func (p *touchImpl) AltitudeAngle() float64 {
	return p.get("altitudeAngle").toFloat64()
}

func (p *touchImpl) AzimuthAngle() float64 {
	return p.get("azimuthAngle").toFloat64()
}

func (p *touchImpl) TouchType() TouchType {
	return TouchType(p.get("touchType").toString())
}

// -------------8<---------------------------------------

type touchEventImpl struct {
	*uiEventImpl
}

func wrapTouchEvent(v Value) TouchEvent {
	if v.valid() {
		return &touchEventImpl{
			uiEventImpl: newUIEventImpl(v),
		}
	}
	return nil
}

func (p *touchEventImpl) Touches() []Touch {
	return touchListToSlice(p.get("touches"))
}

func (p *touchEventImpl) TargetTouches() []Touch {
	return touchListToSlice(p.get("targetTouches"))
}

func (p *touchEventImpl) ChangedTouches() []Touch {
	return touchListToSlice(p.get("changedTouches"))
}

func (p *touchEventImpl) AltKey() bool {
	return p.get("altKey").toBool()
}

func (p *touchEventImpl) MetaKey() bool {
	return p.get("metaKey").toBool()
}

func (p *touchEventImpl) CtrlKey() bool {
	return p.get("ctrlKey").toBool()
}

func (p *touchEventImpl) ShiftKey() bool {
	return p.get("shiftKey").toBool()
}

// -------------8<---------------------------------------

func NewTouch(ti TouchInit) Touch {
	if jsTouch := jsGlobal.get("Touch"); jsTouch.valid() {
		return wrapTouch(jsTouch.jsNew(ti.toJSObject()))
	}
	return nil
}

func NewTouchEvent(typ string, tei ...TouchEventInit) TouchEvent {
	if jsTouchEvent := jsGlobal.get("TouchEvent"); jsTouchEvent.valid() {
		switch len(tei) {
		case 0:
			return wrapTouchEvent(jsTouchEvent.jsNew(typ))
		default:
			return wrapTouchEvent(jsTouchEvent.jsNew(typ, tei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------
