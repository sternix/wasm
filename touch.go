// +build js,wasm

package wasm

// https://w3c.github.io/touch-events/

type (

	// https://w3c.github.io/touch-events/#idl-def-touchinit
	TouchInit struct {
		Identifier    int         `json:"identifier"` // required
		Target        EventTarget `json:"target"`     // required
		ClientX       float64     `json:"clientX"`
		ClientY       float64     `json:"clientY"`
		ScreenX       float64     `json:"screenX"`
		ScreenY       float64     `json:"screenY"`
		PageX         float64     `json:"pageX"`
		PageY         float64     `json:"pageY"`
		RadiusX       float64     `json:"radiusX"`
		RadiusY       float64     `json:"radiusY"`
		RotationAngle float64     `json:"rotationAngle"`
		Force         float64     `json:"force"`
		AltitudeAngle float64     `json:"altitudeAngle"`
		AzimuthAngle  float64     `json:"azimuthAngle"`
		TouchType     TouchType   `json:"touchType"`
	}

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
	TouchEventInit struct {
		EventModifierInit

		Touches        []Touch `json:"touches"`
		TargetTouches  []Touch `json:"targetTouches"`
		ChangedTouches []Touch `json:"changedTouches"`
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

	// https://w3c.github.io/touch-events/#extensions-to-the-globaleventhandlers-interface
	touchEventHandlers interface {
		OnTouchStart(func(Event)) EventHandler
		OnTouchEnd(func(Event)) EventHandler
		OnTouchMove(func(Event)) EventHandler
		OnTouchCancel(func(Event)) EventHandler
	}
)

type TouchType string

const (
	TouchTypeDirect TouchType = "direct"
	TouchTypeStylus TouchType = "stylus"
)
