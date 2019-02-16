// +build js,wasm

package wasm

// -------------8<---------------------------------------

type rtcStatsImpl struct {
	Value
}

func NewRTCStats(ts float64, typ RTCStatsType, id string) RTCStats {
	o := jsObject.New()
	o.Set("timestamp", ts)
	o.Set("type", string(typ))
	o.Set("id", id)
	return wrapRTCStats(Value{o})
}

func newRTCStatsImpl(v Value) rtcStatsImpl {
	return rtcStatsImpl{
		Value: v,
	}
}

func wrapRTCStats(v Value) RTCStats {
	return rtcStatsImpl{
		Value: v,
	}
}

func (p rtcStatsImpl) Timestamp() float64 {
	return p.get("timestamp").toFloat64()
}

func (p rtcStatsImpl) Type() RTCStatsType {
	return RTCStatsType(p.get("type").toString())
}

func (p rtcStatsImpl) Id() string {
	return p.get("id").toString()
}

// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
// -------------8<---------------------------------------
