// +build js,wasm

package wasm

// https://w3c.github.io/permissions/#idl-index

/*
TODO
*/

import (
	"syscall/js"
)

type PermissionDescriptor struct {
	Name PermissionName
}

func (p PermissionDescriptor) toDict() js.Value {
	o := jsObject.New()
	o.Set("name", string(p.Name))
	return o
}

type PermissionName string

const (
	PermissionNameGeolocation        PermissionName = "geolocation"
	PermissionNameNotifications      PermissionName = "notifications"
	PermissionNamePush               PermissionName = "push"
	PermissionNameMidi               PermissionName = "midi"
	PermissionNameCamera             PermissionName = "camera"
	PermissionNameMicrophone         PermissionName = "microphone"
	PermissionNameSpeaker            PermissionName = "speaker"
	PermissionNameDeviceInfo         PermissionName = "device-info"
	PermissionNameBackgroundSync     PermissionName = "background-sync"
	PermissionNameBluetooth          PermissionName = "bluetooth"
	PermissionNamePersistentStorage  PermissionName = "persistent-storage"
	PermissionNameAmbientLightSensor PermissionName = "ambient-light-sensor"
	PermissionNameAccelerometer      PermissionName = "accelerometer"
	PermissionNameGyroscope          PermissionName = "gyroscope"
	PermissionNameMagnetometer       PermissionName = "magnetometer"
)
