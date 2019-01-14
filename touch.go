// +build js,wasm

package wasm

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

func (p TouchInit) toJSObject() Value {
	o := jsObject.jsNew()
	o.set("identifier", p.Identifier)
	o.set("target", JSValue(p.Target))
	o.set("clientX", p.ClientX)
	o.set("clientY", p.ClientY)
	o.set("screenX", p.ScreenX)
	o.set("screenY", p.ScreenY)
	o.set("pageX", p.PageX)
	o.set("pageY", p.PageY)
	o.set("radiusX", p.RadiusX)
	o.set("radiusY", p.RadiusY)
	o.set("rotationAngle", p.RotationAngle)
	o.set("force", p.Force)
	o.set("altitudeAngle", p.AltitudeAngle)
	o.set("azimuthAngle", p.AzimuthAngle)
	o.set("touchType", string(p.TouchType))
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

func (p TouchEventInit) toJSObject() Value {
	o := p.EventModifierInit.toJSObject()
	o.set("touches", sliceToJsArray(p.Touches))
	o.set("targetTouches", sliceToJsArray(p.TargetTouches))
	o.set("changedTouches", sliceToJsArray(p.ChangedTouches))
	return o
}
