// +build js,wasm

package wasm

import (
	"time"
)

// https://w3c.github.io/geolocation-api/#idl-index

type (
	// https://w3c.github.io/geolocation-api/#dom-navigator
	partialNavigatorGeolocation interface {
		Geolocation() Geolocation
	}

	// https://w3c.github.io/geolocation-api/#dom-geolocation
	Geolocation interface {
		CurrentPosition(PositionCallback, ...interface{})
		WatchPosition(PositionCallback, ...interface{}) int
		ClearWatch(int)
	}

	// https://w3c.github.io/geolocation-api/#dom-positionoptions
	PositionOptions struct {
		EnableHighAccuracy bool `json:"enableHighAccuracy"`
		Timeout            int  `json:"timeout"`
		MaximumAge         int  `json:"maximumAge"`
	}

	// https://w3c.github.io/geolocation-api/#dom-position
	Position interface {
		Coords() Coordinates
		Timestamp() time.Time // DOMTimeStamp
	}

	// https://w3c.github.io/geolocation-api/#dom-coordinates
	Coordinates interface {
		Latitude() float64
		Longitude() float64
		Altitude() float64
		Accuracy() float64
		AltitudeAccuracy() float64
		Heading() float64
		Speed() float64
	}

	// https://w3c.github.io/geolocation-api/#dom-positionerror
	PositionError interface {
		Code() PositionErrorCode
		Message() string
	}
)

type PositionErrorCode int

const (
	PositionErrorCodePermissionDenied    PositionErrorCode = 1
	PositionErrorCodePositionUnavailable PositionErrorCode = 2
	PositionErrorCodeTimeout             PositionErrorCode = 3
)
