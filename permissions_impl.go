// +build js,wasm

package wasm

// -------------8<---------------------------------------

type permissionStatusImpl struct {
	*eventTargetImpl
}

func wrapPermissionStatus(v Value) PermissionStatus {
	if v.valid() {
		return &permissionStatusImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *permissionStatusImpl) State() PermissionState {
	return PermissionState(p.get("state").toString())
}

func (p *permissionStatusImpl) OnChange(fn func(Event)) EventHandler {
	return p.On("change", fn)
}

// -------------8<---------------------------------------

type permissionsImpl struct {
	Value
}

func wrapPermissions(v Value) Permissions {
	if v.valid() {
		return &permissionsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *permissionsImpl) Query(desc PermissionDescriptor) func() (PermissionStatus, error) {
	return func() (PermissionStatus, error) {
		result, ok := await(p.call("query", desc.JSValue()))
		if ok {
			return wrapPermissionStatus(result), nil
		}

		return nil, wrapDOMException(result)
	}
}

// -------------8<---------------------------------------

type permissionDescriptorImpl struct {
	Value
}

func NewPermissionDescriptor(name PermissionName) PermissionDescriptor {
	o := jsObject.New()
	o.Set("name", string(name))
	return wrapPermissionDescriptor(Value{o})
}

func newPermissionDescriptorImpl(v Value) permissionDescriptorImpl {
	return permissionDescriptorImpl{
		Value: v,
	}
}

func wrapPermissionDescriptor(v Value) PermissionDescriptor {
	return newPermissionDescriptorImpl(v)
}

func (p permissionDescriptorImpl) Name() PermissionName {
	return PermissionName(p.get("name").toString())
}

// -------------8<---------------------------------------

type pushPermissionDescriptorImpl struct {
	permissionDescriptorImpl
}

func NewPushPermissionDescriptor(name PermissionName, userVisibleOnly bool) PushPermissionDescriptor {
	o := jsObject.New()
	o.Set("name", string(name))
	o.Set("userVisibleOnly", userVisibleOnly)
	return wrapPushPermissionDescriptorImpl(Value{o})
}

func wrapPushPermissionDescriptorImpl(v Value) PushPermissionDescriptor {
	return pushPermissionDescriptorImpl{
		permissionDescriptorImpl: newPermissionDescriptorImpl(v),
	}
}

func (p pushPermissionDescriptorImpl) UserVisibleOnly() bool {
	return p.get("userVisibleOnly").toBool()
}

// -------------8<---------------------------------------

type midiPermissionDescriptorImpl struct {
	permissionDescriptorImpl
}

func NewMidiPermissionDescriptor(name PermissionName, sysex bool) MidiPermissionDescriptor {
	o := jsObject.New()
	o.Set("name", string(name))
	o.Set("sysex", sysex)
	return wrapMidiPermissionDescriptor(Value{o})
}

func wrapMidiPermissionDescriptor(v Value) MidiPermissionDescriptor {
	return midiPermissionDescriptorImpl{
		permissionDescriptorImpl: newPermissionDescriptorImpl(v),
	}
}

func (p midiPermissionDescriptorImpl) Sysex() bool {
	return p.get("sysex").toBool()
}

// -------------8<---------------------------------------

type devicePermissionDescriptorImpl struct {
	permissionDescriptorImpl
}

func NewDevicePermissionDescriptor(name PermissionName, deviceId string) DevicePermissionDescriptor {
	o := jsObject.New()
	o.Set("name", string(name))
	o.Set("deviceId", deviceId)
	return wrapDevicePermissionDescriptor(Value{o})
}

func wrapDevicePermissionDescriptor(v Value) DevicePermissionDescriptor {
	return devicePermissionDescriptorImpl{
		permissionDescriptorImpl: newPermissionDescriptorImpl(v),
	}
}

func (p devicePermissionDescriptorImpl) DeviceId() string {
	return p.get("deviceId").toString()
}

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
