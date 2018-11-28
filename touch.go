// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://w3c.github.io/touch-events/

type (

	// https://w3c.github.io/touch-events/#idl-def-touch
	Touch interface {
		Identifier() int
		Target() EventTarget
		ScreenX() float64
		ScreenY() float64
		ClientX() float64
		ClientY() float64
		PageX() float64
		PageY() float64
		RadiusX() float64
		RadiusY() float64
		RotationAngle() float64
		Force() float64
		AltitudeAngle() float64
		AzimuthAngle() float64
		TouchType() TouchType
	}

	// https://w3c.github.io/touch-events/#idl-def-touchevent
	TouchEvent interface {
		UIEvent

		Touches() []Touch
		TargetTouches() []Touch
		ChangedTouches() []Touch
		AltKey() bool
		MetaKey() bool
		CtrlKey() bool
		ShiftKey() bool
	}
)

type TouchType string

const (
	TouchTypeDirect TouchType = "direct"
	TouchTypeStylus TouchType = "stylus"
)

// -------------8<---------------------------------------

// https://w3c.github.io/touch-events/#idl-def-touchinit
type TouchInit struct {
	Identifier    int         // required
	Target        EventTarget // required
	ClientX       float64
	ClientY       float64
	ScreenX       float64
	ScreenY       float64
	PageX         float64
	PageY         float64
	RadiusX       float64
	RadiusY       float64
	RotationAngle float64
	Force         float64
	AltitudeAngle float64
	AzimuthAngle  float64
	TouchType     TouchType
}

func (p TouchInit) toDict() js.Value {
	o := jsObject.New()
	o.Set("identifier", p.Identifier)
	o.Set("target", p.Target.JSValue())
	o.Set("clientX", p.ClientX)
	o.Set("clientY", p.ClientY)
	o.Set("screenX", p.ScreenX)
	o.Set("screenY", p.ScreenY)
	o.Set("pageX", p.PageX)
	o.Set("pageY", p.PageY)
	o.Set("radiusX", p.RadiusX)
	o.Set("radiusY", p.RadiusY)
	o.Set("rotationAngle", p.RotationAngle)
	o.Set("force", p.Force)
	o.Set("altitudeAngle", p.AltitudeAngle)
	o.Set("azimuthAngle", p.AzimuthAngle)
	o.Set("touchType", string(p.TouchType))
	return o
}

// -------------8<---------------------------------------
// https://w3c.github.io/touch-events/#idl-def-touchevent
type TouchEventInit struct {
	EventModifierInit

	Touches        []Touch
	TargetTouches  []Touch
	ChangedTouches []Touch
}

func (p TouchEventInit) toDict() js.Value {
	o := p.EventModifierInit.toDict()
	o.Set("touches", touchSliceToJsArray(p.Touches))
	o.Set("targetTouches", touchSliceToJsArray(p.TargetTouches))
	o.Set("changedTouches", touchSliceToJsArray(p.ChangedTouches))
	return o
}
