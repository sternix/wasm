// +build js,wasm

package wasm

func NewMediaStream(args ...interface{}) MediaStream {
	if jsMediaStream := jsGlobal.get("MediaStream"); jsMediaStream.valid() {
		switch len(args) {
		case 0:
			return wrapMediaStream(jsMediaStream.jsNew())
		default:
			switch x := args[0].(type) {
			case MediaStream:
				return wrapMediaStream(jsMediaStream.jsNew(JSValueOf(x)))
			case []MediaStreamTrack:
				a := jsArray.New()
				for i, m := range x {
					a.SetIndex(i, JSValueOf(m))
				}
				return wrapMediaStream(jsMediaStream.jsNew(a.JSValue()))
			}
		}
	}
	return nil
}

// -------------8<---------------------------------------

type mediaStreamImpl struct {
	*eventTargetImpl
}

func wrapMediaStream(v Value) MediaStream {
	if v.valid() {
		return &mediaStreamImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaStreamImpl) Id() string {
	return p.get("id").toString()
}

func (p *mediaStreamImpl) AudioTracks() []MediaStreamTrack {
	if s := p.call("getAudioTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) VideoTracks() []MediaStreamTrack {
	if s := p.call("getVideoTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) Tracks() []MediaStreamTrack {
	if s := p.call("getTracks").toSlice(); s != nil {
		ret := make([]MediaStreamTrack, len(s))
		for i, v := range s {
			ret[i] = wrapMediaStreamTrack(v)
		}
		return ret
	}
	return nil
}

func (p *mediaStreamImpl) TrackById(id string) MediaStreamTrack {
	return wrapMediaStreamTrack(p.call("getTrackById", id))
}

func (p *mediaStreamImpl) AddTrack(track MediaStreamTrack) {
	p.call("addTrack", JSValueOf(track))
}

func (p *mediaStreamImpl) RemoveTrack(track MediaStreamTrack) {
	p.call("removeTrack", JSValueOf(track))
}

func (p *mediaStreamImpl) Clone() MediaStream {
	return wrapMediaStream(p.call("clone"))
}

func (p *mediaStreamImpl) Active() bool {
	return p.get("active").toBool()
}

func (p *mediaStreamImpl) OnAddTrack(fn func(Event)) EventHandler {
	return p.On("addtrack", fn)
}

func (p *mediaStreamImpl) OnRemoveTrack(fn func(Event)) EventHandler {
	return p.On("removetrack", fn)
}

// -------------8<---------------------------------------

type mediaStreamTrackImpl struct {
	*eventTargetImpl
}

func wrapMediaStreamTrack(v Value) MediaStreamTrack {
	if v.valid() {
		return &mediaStreamTrackImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaStreamTrackImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *mediaStreamTrackImpl) Id() string {
	return p.get("id").toString()
}

func (p *mediaStreamTrackImpl) Label() string {
	return p.get("label").toString()
}

func (p *mediaStreamTrackImpl) Enabled() bool {
	return p.get("enabled").toBool()
}

func (p *mediaStreamTrackImpl) SetEnabled(b bool) {
	p.set("enabled", b)
}

func (p *mediaStreamTrackImpl) Muted() bool {
	return p.get("muted").toBool()
}

func (p *mediaStreamTrackImpl) OnMute(fn func(Event)) EventHandler {
	return p.On("mute", fn)
}

func (p *mediaStreamTrackImpl) OnUnMute(fn func(Event)) EventHandler {
	return p.On("unmute", fn)
}

func (p *mediaStreamTrackImpl) ReadyState() MediaStreamTrackState {
	return MediaStreamTrackState(p.get("readyState").toString())
}

func (p *mediaStreamTrackImpl) OnEnded(fn func(Event)) EventHandler {
	return p.On("ended", fn)
}

func (p *mediaStreamTrackImpl) Clone() MediaStreamTrack {
	return wrapMediaStreamTrack(p.call("clone"))
}

func (p *mediaStreamTrackImpl) Stop() {
	p.call("stop")
}

func (p *mediaStreamTrackImpl) Capabilities() MediaTrackCapabilities {
	return wrapMediaTrackCapabilities(p.call("getCapabilities"))
}

func (p *mediaStreamTrackImpl) Constraints() MediaTrackConstraints {
	return wrapMediaTrackConstraints(p.call("getConstraints"))
}

func (p *mediaStreamTrackImpl) Settings() MediaTrackSettings {
	return wrapMediaTrackSettings(p.call("getSettings"))
}

func (p *mediaStreamTrackImpl) ApplyConstraints(constraints ...MediaTrackConstraints) func() error {
	return func() error {
		var (
			res Value
			ok  bool
		)

		switch len(constraints) {
		case 0:
			res, ok = await(p.call("applyConstraints"))
		default:
			res, ok = await(p.call("applyConstraints", constraints[0].JSValue()))
		}

		if ok {
			return nil
		}
		return wrapDOMException(res)
	}
}

func (p *mediaStreamTrackImpl) OnOverConstrained(fn func(Event)) EventHandler {
	return p.On("overconstrained", fn)
}

// -------------8<---------------------------------------

type mediaDevicesImpl struct {
	*eventTargetImpl
}

func wrapMediaDevices(v Value) MediaDevices {
	if v.valid() {
		return &mediaDevicesImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *mediaDevicesImpl) OnDeviceChange(fn func(Event)) EventHandler {
	return p.On("devicechange", fn)
}

func (p *mediaDevicesImpl) EnumerateDevices() func() ([]MediaDeviceInfo, error) {
	return func() ([]MediaDeviceInfo, error) {
		result, ok := await(p.call("enumerateDevices"))
		if ok {
			return toMediaDeviceInfoSlice(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

func (p *mediaDevicesImpl) SupportedConstraints() MediaTrackSupportedConstraints {
	return wrapMediaTrackSupportedConstraints(p.call("getSupportedConstraints"))
}

func (p *mediaDevicesImpl) UserMedia(constraints ...MediaStreamConstraints) func() (MediaStream, error) {
	return func() (MediaStream, error) {
		var (
			result Value
			ok     bool
		)

		switch len(constraints) {
		case 0:
			result, ok = await(p.call("getUserMedia"))
		default:
			result, ok = await(p.call("getUserMedia", constraints[0].JSValue()))
		}

		if ok {
			return wrapMediaStream(result), nil
		}

		return nil, wrapDOMException(result)
	}
}

// -------------8<---------------------------------------

type mediaDeviceInfoImpl struct {
	Value
}

func newMediaDeviceInfoImpl(v Value) *mediaDeviceInfoImpl {
	if v.valid() {
		return &mediaDeviceInfoImpl{
			Value: v,
		}
	}
	return nil
}

func wrapMediaDeviceInfo(v Value) MediaDeviceInfo {
	if ret := newMediaDeviceInfoImpl(v); ret != nil {
		return ret
	}
	return nil
}

func (p *mediaDeviceInfoImpl) DeviceId() string {
	return p.get("deviceId").toString()
}

func (p *mediaDeviceInfoImpl) Kind() MediaDeviceKind {
	return MediaDeviceKind(p.get("kind").toString())
}

func (p *mediaDeviceInfoImpl) Label() string {
	return p.get("label").toString()
}

func (p *mediaDeviceInfoImpl) GroupId() string {
	return p.get("groupId").toString()
}

func (p *mediaDeviceInfoImpl) ToJSON() string {
	return jsJSON.call("stringify", p.call("toJSON")).toString()
}

// -------------8<---------------------------------------

type inputDeviceInfoImpl struct {
	*mediaDeviceInfoImpl
}

func wrapInputDeviceInfo(v Value) InputDeviceInfo {
	if v.valid() {
		return &inputDeviceInfoImpl{
			mediaDeviceInfoImpl: newMediaDeviceInfoImpl(v),
		}
	}
	return nil
}

func (p *inputDeviceInfoImpl) Capabilities() MediaTrackCapabilities {
	return wrapMediaTrackCapabilities(p.call("getCapabilities"))
}

// -------------8<---------------------------------------
