// +build js,wasm

package wasm

import (
	"time"
)

// -------------8<---------------------------------------

type geolocationImpl struct {
	Value
}

func wrapGeolocation(v Value) Geolocation {
	if v.valid() {
		return &geolocationImpl{
			Value: v,
		}
	}
	return nil
}

func (p *geolocationImpl) CurrentPosition(cb PositionCallback, args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("getCurrentPosition", cb.jsCallback())
	case 1:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			p.call("getCurrentPosition", cb.jsCallback(), peCb.jsCallback())
		}
	case 2:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			if po, ok := args[1].(PositionOptions); ok {
				p.call("getCurrentPosition", cb.jsCallback(), peCb.jsCallback(), po.JSValue())
			}
		}
	}
}

func (p *geolocationImpl) WatchPosition(cb PositionCallback, args ...interface{}) int {
	switch len(args) {
	case 1:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			return p.call("watchPosition", cb.jsCallback(), peCb.jsCallback()).toInt()
		}
	case 2:
		if peCb, ok := args[0].(PositionErrorCallback); ok {
			if po, ok := args[1].(PositionOptions); ok {
				return p.call("watchPosition", cb.jsCallback(), peCb.jsCallback(), po.JSValue()).toInt()
			}
		}
	}

	return p.call("watchPosition", cb.jsCallback()).toInt()
}

func (p *geolocationImpl) ClearWatch(watchId int) {
	p.call("clearWatch", watchId)
}

// -------------8<---------------------------------------

type positionImpl struct {
	Value
}

func wrapPosition(v Value) Position {
	if v.valid() {
		return &positionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *positionImpl) Coords() Coordinates {
	return wrapCoordinates(p.get("coords"))
}

func (p *positionImpl) Timestamp() time.Time {
	return domTimeStampToTime(p.get("timestamp").toInt())
}

// -------------8<---------------------------------------

type coordinatesImpl struct {
	Value
}

func wrapCoordinates(v Value) Coordinates {
	if v.valid() {
		return &coordinatesImpl{
			Value: v,
		}
	}
	return nil
}

func (p *coordinatesImpl) Latitude() float64 {
	return p.get("latitude").toFloat64()
}

func (p *coordinatesImpl) Longitude() float64 {
	return p.get("longitude").toFloat64()
}

func (p *coordinatesImpl) Altitude() float64 {
	return p.get("altitude").toFloat64()
}

func (p *coordinatesImpl) Accuracy() float64 {
	return p.get("accuracy").toFloat64()
}

func (p *coordinatesImpl) AltitudeAccuracy() float64 {
	return p.get("altitudeAccuracy").toFloat64()
}

func (p *coordinatesImpl) Heading() float64 {
	return p.get("heading").toFloat64()
}

func (p *coordinatesImpl) Speed() float64 {
	return p.get("speed").toFloat64()
}

// -------------8<---------------------------------------

type positionErrorImpl struct {
	Value
}

func wrapPositionError(v Value) PositionError {
	if v.valid() {
		return &positionErrorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *positionErrorImpl) Code() PositionErrorCode {
	return PositionErrorCode(p.get("code").toUint())
}

func (p *positionErrorImpl) Message() string {
	return p.get("message").toString()
}

// -------------8<---------------------------------------
