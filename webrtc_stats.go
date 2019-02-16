// +build js,wasm

// https://www.w3.org/TR/webrtc-stats
package wasm

type (
	// http://w3c.github.io/webrtc-pc/#dom-rtcstats
	// https://www.w3.org/TR/webrtc-stats/#dfn-stats-object
	RTCStats interface {
		Timestamp() float64 // required
		Type() RTCStatsType // required
		Id() string         // required
		JSValue() jsValue
	}

	// https://www.w3.org/TR/webrtc-stats/#dom-rtccodecstats
	RTCCodecStats interface {
		RTCStats

		PayloadType() uint
		CodecType() RTCCodecType
		TransportId() string
		MimeType() string
		ClockRate() uint
		Channels() uint
		SdpFmtpLine() string
		Implementation() string
	}
)

// https://www.w3.org/TR/webrtc-stats/#dom-rtcstatstype
type RTCStatsType string

const (
	RTCStatsTypeCodec             RTCStatsType = "codec"
	RTCStatsTypeInboundRtp        RTCStatsType = "inbound-rtp"
	RTCStatsTypeOutboundRtp       RTCStatsType = "outbound-rtp"
	RTCStatsTypeRemoteInboundRtp  RTCStatsType = "remote-inbound-rtp"
	RTCStatsTypeRemoteOutboundRtp RTCStatsType = "remote-outbound-rtp"
	RTCStatsTypeCSRC              RTCStatsType = "csrc"
	RTCStatsTypePeerConnection    RTCStatsType = "peer-connection"
	RTCStatsTypeDataChannel       RTCStatsType = "data-channel"
	RTCStatsTypeStream            RTCStatsType = "stream"
	RTCStatsTypeTrack             RTCStatsType = "track"
	RTCStatsTypeSender            RTCStatsType = "sender"
	RTCStatsTypeReceiver          RTCStatsType = "receiver"
	RTCStatsTypeTransport         RTCStatsType = "transport"
	RTCStatsTypeCandidatePair     RTCStatsType = "candidate-pair"
	RTCStatsTypeLocalCandidate    RTCStatsType = "local-candidate"
	RTCStatsTypeRemoteCandidate   RTCStatsType = "remote-candidate"
	RTCStatsTypeCertificate       RTCStatsType = "certificate"
)

// https://www.w3.org/TR/webrtc-stats/#dom-rtccodectype
type RTCCodecType string

const (
	RTCCodecTypeEncode RTCCodecType = "encode"
	RTCCodecTypeDecode RTCCodecType = "decode"
)
