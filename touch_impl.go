// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type touchImpl struct {
	js.Value
}

func newTouch(v js.Value) Touch {
	if isNil(v) {
		return nil
	}

	return &touchImpl{
		Value: v,
	}
}

func (p *touchImpl) Identifier() int {
	return p.Get("identifier").Int()
}

func (p *touchImpl) Target() EventTarget {
	return newEventTargetImpl(p.Get("target"))
}

func (p *touchImpl) ScreenX() float64 {
	return p.Get("screenX").Float()
}

func (p *touchImpl) ScreenY() float64 {
	return p.Get("screenY").Float()
}

func (p *touchImpl) ClientX() float64 {
	return p.Get("clientX").Float()
}

func (p *touchImpl) ClientY() float64 {
	return p.Get("clientY").Float()
}

func (p *touchImpl) PageX() float64 {
	return p.Get("pageX").Float()
}

func (p *touchImpl) PageY() float64 {
	return p.Get("pageY").Float()
}

func (p *touchImpl) RadiusX() float64 {
	return p.Get("radiusX").Float()
}

func (p *touchImpl) RadiusY() float64 {
	return p.Get("radiusY").Float()
}

func (p *touchImpl) RotationAngle() float64 {
	return p.Get("rotationAngle").Float()
}

func (p *touchImpl) Force() float64 {
	return p.Get("force").Float()
}

func (p *touchImpl) AltitudeAngle() float64 {
	return p.Get("altitudeAngle").Float()
}

func (p *touchImpl) AzimuthAngle() float64 {
	return p.Get("azimuthAngle").Float()
}

func (p *touchImpl) TouchType() TouchType {
	return TouchType(p.Get("touchType").String())
}

// -------------8<---------------------------------------

type touchEventImpl struct {
	*uiEventImpl
}

func newTouchEvent(v js.Value) TouchEvent {
	if isNil(v) {
		return nil
	}

	return &touchEventImpl{
		uiEventImpl: newUIEventImpl(v),
	}
}

func (p *touchEventImpl) Touches() []Touch {
	return touchListToSlice(p.Get("touches"))
}

func (p *touchEventImpl) TargetTouches() []Touch {
	return touchListToSlice(p.Get("targetTouches"))
}

func (p *touchEventImpl) ChangedTouches() []Touch {
	return touchListToSlice(p.Get("changedTouches"))
}

func (p *touchEventImpl) AltKey() bool {
	return p.Get("altKey").Bool()
}

func (p *touchEventImpl) MetaKey() bool {
	return p.Get("metaKey").Bool()
}

func (p *touchEventImpl) CtrlKey() bool {
	return p.Get("ctrlKey").Bool()
}

func (p *touchEventImpl) ShiftKey() bool {
	return p.Get("shiftKey").Bool()
}

// -------------8<---------------------------------------

func NewTouch(ti TouchInit) Touch {
	jsTouch := js.Global().Get("Touch")
	if isNil(jsTouch) {
		return nil
	}

	return newTouch(jsTouch.New(ti.toDict()))
}

func NewTouchEvent(typ string, tei ...TouchEventInit) TouchEvent {
	jsTouchEvent := js.Global().Get("TouchEvent")
	if isNil(jsTouchEvent) {
		return nil
	}

	switch len(tei) {
	case 0:
		return newTouchEvent(jsTouchEvent.New(typ))
	default:
		return newTouchEvent(jsTouchEvent.New(typ, tei[0].toDict()))
	}
}

// -------------8<---------------------------------------
