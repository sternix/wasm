// +build js,wasm

// https://w3c.github.io/mediacapture-main
package wasm

type (
	// https://w3c.github.io/mediacapture-main/#mediastream
	MediaStream interface {
		EventTarget

		Id() string
		AudioTracks() []MediaStreamTrack
		VideoTracks() []MediaStreamTrack
		Tracks() []MediaStreamTrack
		TrackById(string) MediaStreamTrack
		AddTrack(MediaStreamTrack)
		RemoveTrack(MediaStreamTrack)
		Clone() MediaStream
		Active() bool
		OnAddTrack(func(Event)) EventHandler
		OnRemoveTrack(func(Event)) EventHandler
	}

	// https://w3c.github.io/mediacapture-main/#mediastreamtrack
	MediaStreamTrack interface {
		EventTarget

		Kind() string
		Id() string
		Label() string
		Enabled() bool
		SetEnabled(bool)
		Muted() bool
		OnMute(func(Event)) EventHandler
		OnUnMute(func(Event)) EventHandler
		ReadyState() MediaStreamTrackState
		OnEnded(func(Event)) EventHandler
		Clone() MediaStreamTrack
		Stop()
		Capabilities() MediaTrackCapabilities
		Constraints() MediaTrackConstraints
		Settings() MediaTrackSettings
		ApplyConstraints(...MediaTrackConstraints) func() error
		OnOverConstrained(func(Event)) EventHandler
	}

	// https://w3c.github.io/mediacapture-main/#dom-mediadevices
	MediaDevices interface {
		EventTarget

		OnDeviceChange() EventHandler
		EnumerateDevices() func() ([]MediaDeviceInfo, error)

		// https://w3c.github.io/mediacapture-main/#mediadevices-interface-extensions
		SupportedConstraints() MediaTrackSupportedConstraints
		UserMedia(...MediaStreamConstraints) func() (MediaStream, error)
	}

	// https://w3c.github.io/mediacapture-main/#device-info
	MediaDeviceInfo interface {
		DeviceId() string
		Kind() MediaDeviceKind
		Label() string
		GroupId() string
		ToJSON() []string
	}

	// https://w3c.github.io/mediacapture-main/#dom-inputdeviceinfo
	InputDeviceInfo interface {
		MediaDeviceInfo

		Capabilities() MediaTrackCapabilities
	}
)

// https://w3c.github.io/mediacapture-main/#dom-mediastreamtrackstate
type MediaStreamTrackState string

const (
	MediaStreamTrackStateLive  MediaStreamTrackState = "live"
	MediaStreamTrackStateEnded MediaStreamTrackState = "ended"
)

// https://w3c.github.io/mediacapture-main/#dom-mediadevicekind
type MediaDeviceKind string

const (
	MediaDeviceKindAudioInput  MediaDeviceKind = "audioinput"
	MediaDeviceKindAudioOutput MediaDeviceKind = "audiooutput"
	MediaDeviceKindVideoInput  MediaDeviceKind = "videoinput"
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-mediatrackconstraints
type MediaTrackConstraints struct {
	MediaTrackConstraintSet

	Advanced []MediaTrackConstraintSet
}

func wrapMediaTrackConstraints(v Value) MediaTrackConstraints {
	c := MediaTrackConstraints{}
	if v.valid() {
		c.MediaTrackConstraintSet = wrapMediaTrackConstraintSet(v)
		c.Advanced = mediaTrackConstraintsSequenceToSlice(v.get("advanced"))
	}
	return c
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-mediatrackcapabilities
type MediaTrackCapabilities struct {
	Width            ULongRange
	Heigth           ULongRange
	AspectRatio      DoubleRange
	FrameRate        DoubleRange
	FacingMode       []string
	ResizeMode       []string
	Volume           DoubleRange
	SampleRate       ULongRange
	SampleSize       ULongRange
	EchoCancellation []bool
	AutoGainControl  []bool
	NoiseSuppression []bool
	Latency          DoubleRange
	ChannelCount     ULongRange
	DeviceId         string
	GroupId          string
}

func wrapMediaTrackCapabilities(v Value) MediaTrackCapabilities {
	c := MediaTrackCapabilities{}
	if v.valid() {
		c.Width = wrapULongRange(v.get("width"))
		c.Heigth = wrapULongRange(v.get("height"))
		c.AspectRatio = wrapDoubleRange(v.get("aspectRatio"))
		c.FrameRate = wrapDoubleRange(v.get("frameRate"))
		c.FacingMode = stringSequenceToSlice(v.get("facingMode"))
		c.ResizeMode = stringSequenceToSlice(v.get("resizeMode"))
		c.Volume = wrapDoubleRange(v.get("volume"))
		c.SampleRate = wrapULongRange(v.get("sampleRate"))
		c.SampleSize = wrapULongRange(v.get("sampleSize"))
		c.EchoCancellation = boolSequenceToSlice(v.get("echoCancellation"))
		c.AutoGainControl = boolSequenceToSlice(v.get("autoGainControl"))
		c.NoiseSuppression = boolSequenceToSlice(v.get("noiseSuppression"))
		c.Latency = wrapDoubleRange(v.get("latency"))
		c.ChannelCount = wrapULongRange(v.get("channelCount"))
		c.DeviceId = v.get("deviceId").toString()
		c.GroupId = v.get("groupId").toString()
	}
	return c
}

func (p MediaTrackCapabilities) JSValue() jsValue {
	o := jsObject.New()
	o.Set("width", p.Width.JSValue())
	o.Set("height", p.Heigth.JSValue())
	o.Set("aspectRatio", p.AspectRatio.JSValue())
	o.Set("frameRate", p.FrameRate.JSValue())
	o.Set("facingMode", ToJSArray(p.FacingMode))
	o.Set("resizeMode", ToJSArray(p.ResizeMode))
	o.Set("volume", p.Volume.JSValue())
	o.Set("sampleRate", p.SampleRate.JSValue())
	o.Set("sampleSize", p.SampleSize.JSValue())
	o.Set("echoCancellation", ToJSArray(p.EchoCancellation))
	o.Set("autoGainControl", ToJSArray(p.AutoGainControl))
	o.Set("noiseSuppression", ToJSArray(p.NoiseSuppression))
	o.Set("latency", p.Latency.JSValue())
	o.Set("channelCount", p.ChannelCount.JSValue())
	o.Set("deviceId", p.DeviceId)
	o.Set("groupId", p.GroupId)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-mediatrackconstraintset
type MediaTrackConstraintSet struct {
	Width            ConstrainULong
	Height           ConstrainULong
	AspectRatio      ConstrainDouble
	FrameRate        ConstrainDouble
	FacingMode       ConstrainDOMString
	ResizeMode       ConstrainDOMString
	Volume           ConstrainDouble
	SampleRate       ConstrainULong
	SampleSize       ConstrainULong
	EchoCancellation ConstrainBoolean
	AutoGainControl  ConstrainBoolean
	NoiseSuppression ConstrainBoolean
	Latency          ConstrainDouble
	ChannelCount     ConstrainULong
	DeviceId         ConstrainDOMString
	GroupId          ConstrainDOMString
}

func wrapMediaTrackConstraintSet(v Value) MediaTrackConstraintSet {
	s := MediaTrackConstraintSet{}
	if v.valid() {
		s.Width = ConstrainULong(v.get("width").toUint())
		s.Height = ConstrainULong(v.get("height").toUint())
		s.AspectRatio = wrapConstrainDouble(v.get("aspectRatio"))
		s.FrameRate = wrapConstrainDouble(v.get("frameRate"))
		s.FacingMode = wrapConstrainDOMString(v.get("facingMode"))
		s.ResizeMode = wrapConstrainDOMString(v.get("resizeMode"))
		s.Volume = wrapConstrainDouble(v.get("volume"))
		s.SampleRate = ConstrainULong(v.get("sampleRate").toUint())
		s.SampleSize = ConstrainULong(v.get("sampleSize").toUint())
		s.EchoCancellation = wrapConstrainBoolean(v.get("echoCancellation"))
		s.AutoGainControl = wrapConstrainBoolean(v.get("autoGainControl"))
		s.NoiseSuppression = wrapConstrainBoolean(v.get("noiseSuppression"))
		s.Latency = wrapConstrainDouble(v.get("latency"))
		s.ChannelCount = ConstrainULong(v.get("channelCount").toUint())
		s.DeviceId = wrapConstrainDOMString(v.get("deviceId"))
		s.GroupId = wrapConstrainDOMString(v.get("groupId"))
	}
	return s
}

func (p MediaTrackConstraintSet) JSValue() jsValue {
	o := jsObject.New()
	o.Set("width", uint(p.Width))
	o.Set("height", uint(p.Height))
	o.Set("aspectRatio", constrainDoubleJSValue(p.AspectRatio))
	o.Set("frameRate", constrainDoubleJSValue(p.FrameRate))
	o.Set("facingMode", constrainDOMStringJSValue(p.FacingMode))
	o.Set("resizeMode", constrainDOMStringJSValue(p.ResizeMode))
	o.Set("volume", constrainDoubleJSValue(p.Volume))
	o.Set("sampleRate", uint(p.SampleRate))
	o.Set("sampleSize", uint(p.SampleSize))
	o.Set("echoCancellation", constrainBooleanJSValue(p.EchoCancellation))
	o.Set("autoGainControl", constrainBooleanJSValue(p.AutoGainControl))
	o.Set("noiseSuppression", constrainBooleanJSValue(p.NoiseSuppression))
	o.Set("latency", constrainDoubleJSValue(p.Latency))
	o.Set("channelCount", uint(p.ChannelCount))
	o.Set("deviceId", constrainDOMStringJSValue(p.DeviceId))
	o.Set("groupId", constrainDOMStringJSValue(p.GroupId))
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-mediatracksettings
type MediaTrackSettings struct {
	Width            int
	Height           int
	AspectRatio      float64
	FrameRate        float64
	FacingMode       string
	ResizeMode       string
	Volume           float64
	SampleRate       int
	SampleSize       int
	EchoCancellation bool
	AutoGainControl  bool
	NoiseSuppression bool
	Latency          float64
	ChannelCount     int
	DeviceId         string
	GroupId          string
}

func wrapMediaTrackSettings(v Value) MediaTrackSettings {
	s := MediaTrackSettings{}
	if v.valid() {
		s.Width = v.get("width").toInt()
		s.Height = v.get("height").toInt()
		s.AspectRatio = v.get("aspectRatio").toFloat64()
		s.FrameRate = v.get("frameRate").toFloat64()
		s.FacingMode = v.get("facingMode").toString()
		s.ResizeMode = v.get("resizeMode").toString()
		s.Volume = v.get("volume").toFloat64()
		s.SampleRate = v.get("sampleRate").toInt()
		s.SampleSize = v.get("sampleSize").toInt()
		s.EchoCancellation = v.get("echoCancellation").toBool()
		s.AutoGainControl = v.get("autoGainControl").toBool()
		s.NoiseSuppression = v.get("noiseSuppression").toBool()
		s.Latency = v.get("latency").toFloat64()
		s.ChannelCount = v.get("channelCount").toInt()
		s.DeviceId = v.get("deviceId").toString()
		s.GroupId = v.get("groupId").toString()
	}
	return s
}

func (p MediaTrackSettings) JSValue() jsValue {
	o := jsObject.New()
	o.Set("width", p.Width)
	o.Set("height", p.Height)
	o.Set("aspectRatio", p.AspectRatio)
	o.Set("frameRate", p.FrameRate)
	o.Set("facingMode", p.FacingMode)
	o.Set("volume", p.Volume)
	o.Set("sampleRate", p.SampleRate)
	o.Set("sampleSize", p.SampleSize)
	o.Set("echoCancellation", p.EchoCancellation)
	o.Set("autoGainControl", p.AutoGainControl)
	o.Set("noiseSuppression", p.NoiseSuppression)
	o.Set("latency", p.Latency)
	o.Set("channelCount", p.ChannelCount)
	o.Set("deviceId", p.DeviceId)
	o.Set("groupId", p.GroupId)
	return o
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-ulongrange
type ULongRange struct {
	Max uint
	Min uint
}

func wrapULongRange(v Value) ULongRange {
	r := ULongRange{}
	if v.valid() {
		r.Max = v.get("max").toUint()
		r.Min = v.get("min").toUint()
	}
	return r
}

func (p ULongRange) JSValue() jsValue {
	o := jsObject.New()
	o.Set("max", p.Max)
	o.Set("min", p.Min)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-longrange
type LongRange struct {
	Max int
	Min int
}

func wrapLongRange(v Value) LongRange {
	l := LongRange{}
	if v.valid() {
		l.Max = v.get("max").toInt()
		l.Min = v.get("min").toInt()
	}
	return l
}

func (p LongRange) JSValue() jsValue {
	o := jsObject.New()
	o.Set("max", p.Max)
	o.Set("min", p.Min)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-doublerange
type DoubleRange struct {
	Max float64
	Min float64
}

func wrapDoubleRange(v Value) DoubleRange {
	d := DoubleRange{}
	if v.valid() {
		d.Max = v.get("max").toFloat64()
		d.Min = v.get("min").toFloat64()
	}
	return d
}

func (p DoubleRange) JSValue() jsValue {
	o := jsObject.New()
	o.Set("max", p.Max)
	o.Set("min", p.Min)
	return o
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-constrainulongrange
type ConstrainULongRange struct {
	ULongRange

	Exact uint
	Ideal uint
}

func (p ConstrainULongRange) JSValue() jsValue {
	o := p.ULongRange.JSValue()
	o.Set("exact", p.Exact)
	o.Set("ideal", p.Ideal)
	return o
}

// -------------8<---------------------------------------
// https://w3c.github.io/mediacapture-main/#dom-constrainulong
// typedef ([Clamp] unsigned long or ConstrainULongRange) ConstrainULong;

type ConstrainULong uint

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constrainlongrange
type ConstrainLongRange struct {
	LongRange

	Exact int
	Ideal int
}

func (p ConstrainLongRange) JSValue() jsValue {
	o := p.LongRange.JSValue()
	o.Set("exact", p.Exact)
	o.Set("ideal", p.Ideal)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constrainlong
type ConstrainLong ConstrainLongRange

func wrapConstrainLong(v Value) ConstrainLong {
	c := ConstrainLong{}
	if v.valid() {
		c.LongRange = wrapLongRange(v)
		c.Exact = v.get("exact").toInt()
		c.Ideal = v.get("ideal").toInt()
	}
	return c
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constraindoublerange
type ConstrainDoubleRange struct {
	DoubleRange

	Exact float64
	Ideal float64
}

func wrapConstrainDoubleRange(v Value) ConstrainDoubleRange {
	d := ConstrainDoubleRange{}
	if v.valid() {
		d.DoubleRange = wrapDoubleRange(v)
		d.Exact = v.get("exact").toFloat64()
		d.Ideal = v.get("ideal").toFloat64()
	}
	return d

}

func (p ConstrainDoubleRange) JSValue() jsValue {
	o := p.DoubleRange.JSValue()
	o.Set("exact", p.Exact)
	o.Set("ideal", p.Ideal)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constraindouble
// typedef (double or ConstrainDoubleRange) ConstrainDouble;
type ConstrainDouble interface{}

func wrapConstrainDouble(v Value) ConstrainDouble {
	if v.valid() {
		switch v.jsType() {
		case "double":
			return v.toFloat64()
		case "ConstrainDoubleRange": // TODO may be ConstrainDouble
			return ConstrainDoubleRange{
				DoubleRange: wrapDoubleRange(v),
				Exact:       v.get("exact").toFloat64(),
				Ideal:       v.get("ideal").toFloat64(),
			}
		}
	}
	return nil
}

func constrainDoubleJSValue(p ConstrainDouble) jsValue {
	switch x := p.(type) {
	case nil:
		return jsNull
	case float64:
		return JSValueOf(x)
	case ConstrainDoubleRange:
		return x.JSValue()
	default:
		return jsUndefined
	}
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constraindomstringparameters
type ConstrainDOMStringParameters struct {
	Exact string
	Ideal string
}

func (p ConstrainDOMStringParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("exact", p.Exact)
	o.Set("ideal", p.Ideal)
	return o
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-constraindomstring
// typedef (DOMString or sequence<DOMString> or ConstrainDOMStringParameters) ConstrainDOMString;

type ConstrainDOMString interface{}

func wrapConstrainDOMString(v Value) ConstrainDOMString {
	if v.valid() {
		switch v.jsType() {
		case "string":
		//case // Sequence // TODO sequence<DOMString> ,
		// The type name of a sequence type is the concatenation of the type name for T and the string "Sequence".
		case "ConstrainDOMStringParameters":
			return ConstrainDOMStringParameters{
				Exact: v.get("exact").toString(),
				Ideal: v.get("ideal").toString(),
			}
		}
	}
	return nil
}

func constrainDOMStringJSValue(p ConstrainDOMString) jsValue {
	switch x := p.(type) {
	case nil:
		return jsNull
	case string:
		return JSValueOf(x)
	case []string:
		return ToJSArray(x)
	case ConstrainDOMStringParameters:
		return x.JSValue()
	default:
		return jsUndefined
	}
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constrainbooleanparameters
type ConstrainBooleanParameters struct {
	Exact bool
	Ideal bool
}

func (p ConstrainBooleanParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("exact", p.Exact)
	o.Set("ideal", p.Ideal)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/mediacapture-streams/#dom-constrainboolean
// typedef (boolean or ConstrainBooleanParameters) ConstrainBoolean;
type ConstrainBoolean interface{}

func wrapConstrainBoolean(v Value) ConstrainBoolean {
	if v.valid() {
		switch v.jsType() {
		case "boolean":
			return v.toBool()
		case "ConstrainBooleanParameters": // TODO
			return ConstrainBooleanParameters{
				Exact: v.get("exact").toBool(),
				Ideal: v.get("ideal").toBool(),
			}
		}
	}
	return nil
}

func constrainBooleanJSValue(p ConstrainBoolean) jsValue {
	switch x := p.(type) {
	case nil:
		return jsNull
	case bool:
		return JSValueOf(x)
	case ConstrainBooleanParameters:
		return x.JSValue()
	default:
		return jsUndefined
	}
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-mediatracksupportedconstraints
type MediaTrackSupportedConstraints struct {
	Width            bool
	Height           bool
	AspectRatio      bool
	FrameRate        bool
	FacingMode       bool
	ResizeMode       bool
	Volume           bool
	SampleRate       bool
	SampleSize       bool
	EchoCancellation bool
	AutoGainControl  bool
	NoiseSuppression bool
	Latency          bool
	ChannelCount     bool
	DeviceId         bool
	GroupId          bool
}

func (p MediaTrackSupportedConstraints) JSValue() jsValue {
	o := jsObject.New()
	o.Set("width", p.Width)
	o.Set("height", p.Height)
	o.Set("aspectRatio", p.AspectRatio)
	o.Set("frameRate", p.FrameRate)
	o.Set("facingMode", p.FacingMode)
	o.Set("resizeMode", p.ResizeMode)
	o.Set("volume", p.Volume)
	o.Set("sampleRate", p.SampleRate)
	o.Set("sampleSize", p.SampleSize)
	o.Set("echoCancellation", p.EchoCancellation)
	o.Set("autoGainControl", p.AutoGainControl)
	o.Set("noiseSuppression", p.NoiseSuppression)
	o.Set("latency", p.Latency)
	o.Set("channelCount", p.ChannelCount)
	o.Set("deviceId", p.DeviceId)
	o.Set("groupId", p.GroupId)
	return o
}

// -------------8<---------------------------------------

// https://w3c.github.io/mediacapture-main/#dom-mediastreamconstraints
type MediaStreamConstraints struct {
	Video bool
	Audio bool
}

func (p MediaStreamConstraints) JSValue() jsValue {
	o := jsObject.New()
	o.Set("video", p.Video)
	o.Set("audio", p.Audio)
	return o
}

// -------------8<---------------------------------------
