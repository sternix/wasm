// +build js,wasm

// https://w3c.github.io/permissions/#idl-index
package wasm

type (
	PermissionStatus interface {
		EventTarget

		State() PermissionState
		OnChange(func(Event)) EventHandler
	}

	Permissions interface {
		Query(PermissionDescriptor) func() (PermissionStatus, error)
	}

	PermissionDescriptor interface {
		Name() PermissionName
		JSValue() jsValue
	}

	PushPermissionDescriptor interface {
		PermissionDescriptor

		UserVisibleOnly() bool
	}

	MidiPermissionDescriptor interface {
		PermissionDescriptor

		Sysex() bool
	}

	DevicePermissionDescriptor interface {
		PermissionDescriptor

		DeviceId() string
	}
)

type PermissionSetParameters struct {
	Descriptor PermissionDescriptor
	State      PermissionState
	OneRealm   bool
}

func (p PermissionSetParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("descriptor", p.Descriptor.JSValue())
	o.Set("state", string(p.State))
	o.Set("oneRealm", p.OneRealm)
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
	PermissionNameBackgroundFetch    PermissionName = "background-fetch"
	PermissionNameBackgroundSync     PermissionName = "background-sync"
	PermissionNameBluetooth          PermissionName = "bluetooth"
	PermissionNamePersistentStorage  PermissionName = "persistent-storage"
	PermissionNameAmbientLightSensor PermissionName = "ambient-light-sensor"
	PermissionNameAccelerometer      PermissionName = "accelerometer"
	PermissionNameGyroscope          PermissionName = "gyroscope"
	PermissionNameMagnetometer       PermissionName = "magnetometer"
	PermissionNameClipboard          PermissionName = "clipboard"
	PermissionNameDisplay            PermissionName = "display"
)

type PermissionState string

const (
	PermissionStateGranted PermissionState = "granted"
	PermissionStateDenied  PermissionState = "denied"
	PermissionStatePrompt  PermissionState = "prompt"
)
