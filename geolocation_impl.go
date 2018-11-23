// +build js,wasm

package wasm

import (
	"syscall/js"
	"time"
)

// -------------8<---------------------------------------

type geolocationImpl struct {
	js.Value
}

func newGeolocation(v js.Value) Geolocation {
	if isNil(v) {
		return nil
	}

	return &geolocationImpl{
		Value: v,
	}
}

func (p *geolocationImpl) CurrentPosition(cb PositionCallback, args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("getCurrentPosition", cb.jsCallback())
	case 1:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			p.Call("getCurrentPosition", cb.jsCallback(), peCb.jsCallback())
		}
	case 2:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			if po, ok := args[1].(PositionOptions); ok {
				p.Call("getCurrentPosition", cb.jsCallback(), peCb.jsCallback(), po.toDict())
			}
		}
	}
}

func (p *geolocationImpl) WatchPosition(cb PositionCallback, args ...interface{}) int {
	switch len(args) {
	case 1:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			return p.Call("watchPosition", cb.jsCallback(), peCb.jsCallback()).Int()
		}
	case 2:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			if po, ok := args[1].(PositionOptions); ok {
				return p.Call("watchPosition", cb.jsCallback(), peCb.jsCallback(), po.toDict()).Int()
			}
		}
	}

	return p.Call("watchPosition", cb.jsCallback()).Int()
}

func (p *geolocationImpl) ClearWatch(watchId int) {
	p.Call("clearWatch", watchId)
}

// -------------8<---------------------------------------

type positionImpl struct {
	js.Value
}

func newPosition(v js.Value) Position {
	if isNil(v) {
		return nil
	}

	return &positionImpl{
		Value: v,
	}
}

func (p *positionImpl) Coords() Coordinates {
	return newCoordinates(p.Get("coords"))
}

func (p *positionImpl) Timestamp() time.Time {
	return domTimeStampToTime(p.Get("timestamp").Int())
}

// -------------8<---------------------------------------

type coordinatesImpl struct {
	js.Value
}

func newCoordinates(v js.Value) Coordinates {
	if isNil(v) {
		return nil
	}

	return &coordinatesImpl{
		Value: v,
	}
}

func (p *coordinatesImpl) Latitude() float64 {
	return p.Get("latitude").Float()
}

func (p *coordinatesImpl) Longitude() float64 {
	return p.Get("longitude").Float()
}

func (p *coordinatesImpl) Altitude() float64 {
	return p.Get("altitude").Float()
}

func (p *coordinatesImpl) Accuracy() float64 {
	return p.Get("accuracy").Float()
}

func (p *coordinatesImpl) AltitudeAccuracy() float64 {
	return p.Get("altitudeAccuracy").Float()
}

func (p *coordinatesImpl) Heading() float64 {
	return p.Get("heading").Float()
}

func (p *coordinatesImpl) Speed() float64 {
	return p.Get("speed").Float()
}

// -------------8<---------------------------------------

type positionErrorImpl struct {
	js.Value
}

func newPositionError(v js.Value) PositionError {
	if isNil(v) {
		return nil
	}

	return &positionErrorImpl{
		Value: v,
	}
}

func (p *positionErrorImpl) Code() PositionErrorCode {
	return PositionErrorCode(p.Get("code").Int())
}

func (p *positionErrorImpl) Message() string {
	return p.Get("message").String()
}

// -------------8<---------------------------------------
