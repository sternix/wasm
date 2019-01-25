// +build js,wasm

// http://w3c.github.io/webrtc-pc/
package wasm

import (
	"time"
)

// https://www.w3.org/TR/WebCryptoAPI/#dfn-AlgorithmIdentifier
// TODO
type AlgorithmIdentifier string

// https://heycam.github.io/webidl/#common-DOMTimeStamp
// typedef unsigned long long DOMTimeStamp;

func NewRTCPeerConnection(configuration ...RTCConfiguration) RTCPeerConnection {
	jsRTCPeerConnection := jsGlobal.get("RTCPeerConnection")
	switch len(configuration) {
	case 0:
		return wrapRTCPeerConnection(jsRTCPeerConnection.jsNew())
	default:
		return wrapRTCPeerConnection(jsRTCPeerConnection.jsNew(configuration[0].JSValue()))
	}
}

func NewRTCSessionDescription(descriptionInitDict RTCSessionDescriptionInit) RTCSessionDescription {
	if jsSD := jsGlobal.get("RTCSessionDescription"); jsSD.valid() {
		return wrapRTCSessionDescription(jsSD.jsNew(descriptionInitDict.JSValue()))
	}
	return nil
}

func NewRTCIceCandidate(candidateInitDict ...RTCIceCandidateInit) RTCIceCandidate {
	if jsIC := jsGlobal.get("RTCIceCandidate"); jsIC.valid() {
		switch len(candidateInitDict) {
		case 0:
			return wrapRTCIceCandidate(jsIC.jsNew())
		default:
			return wrapRTCIceCandidate(jsIC.jsNew(candidateInitDict[0].JSValue()))
		}
	}
	return nil
}

func NewRTCPeerConnectionIceEvent(typ string, eventInitDict ...RTCPeerConnectionIceEventInit) RTCPeerConnectionIceEvent {
	if jsIE := jsGlobal.get("RTCPeerConnectionIceEvent"); jsIE.valid() {
		switch len(eventInitDict) {
		case 0:
			return wrapRTCPeerConnectionIceEvent(jsIE.jsNew(typ))
		default:
			return wrapRTCPeerConnectionIceEvent(jsIE.jsNew(typ, eventInitDict[0].JSValue()))
		}
	}
	return nil
}

func NewRTCPeerConnectionIceErrorEvent(typ string, eventInitDict RTCPeerConnectionIceErrorEventInit) RTCPeerConnectionIceErrorEvent {
	if jsEE := jsGlobal.get("RTCPeerConnectionIceErrorEvent"); jsEE.valid() {
		return wrapRTCPeerConnectionIceErrorEvent(jsEE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

func NewRTCTrackEvent(typ string, eventInitDict RTCTrackEventInit) RTCTrackEvent {
	if jsTE := jsGlobal.get("RTCTrackEvent"); jsTE.valid() {
		return wrapRTCTrackEvent(jsTE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

func NewRTCDataChannelEvent(typ string, eventInitDict RTCDataChannelEventInit) RTCDataChannelEvent {
	if jsCE := jsGlobal.get("RTCDataChannelEvent"); jsCE.valid() {
		return wrapRTCDataChannelEvent(jsCE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

func NewRTCDTMFToneChangeEvent(typ string, eventInitDict RTCDTMFToneChangeEventInit) RTCDTMFToneChangeEvent {
	if jsTCE := jsGlobal.get("RTCDTMFToneChangeEvent"); jsTCE.valid() {
		return wrapRTCDTMFToneChangeEvent(jsTCE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

func NewRTCStatsEvent(typ string, eventInitDict RTCStatsEventInit) RTCStatsEvent {
	if jsSE := jsGlobal.get("RTCStatsEvent"); jsSE.valid() {
		return wrapRTCStatsEvent(jsSE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

func NewRTCError(detail RTCErrorDetailType, message string) RTCError {
	if jsE := jsGlobal.get("RTCError"); jsE.valid() {
		return wrapRTCError(jsE.jsNew(string(detail), message))
	}
	return nil
}

func NewRTCErrorEvent(typ string, eventInitDict RTCErrorEventInit) RTCErrorEvent {
	if jsEE := jsGlobal.get("RTCErrorEvent"); jsEE.valid() {
		return wrapRTCErrorEvent(jsEE.jsNew(typ, eventInitDict.JSValue()))
	}
	return nil
}

type (
	// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnection
	RTCPeerConnection interface {
		EventTarget

		CreateOffer(...RTCOfferOptions) func() (RTCSessionDescriptionInit, error)
		CreateAnswer(...RTCAnswerOptions) func() (RTCSessionDescriptionInit, error)
		SetLocalDescription(RTCSessionDescriptionInit) func() error
		LocalDescription() RTCSessionDescription
		CurrentLocalDescription() RTCSessionDescription
		PendingLocalDescription() RTCSessionDescription
		SetRemoteDescription(RTCSessionDescriptionInit) func() error
		RemoteDescription() RTCSessionDescription
		CurrentRemoteDescription() RTCSessionDescription
		PendingRemoteDescription() RTCSessionDescription
		AddIceCandidate(RTCIceCandidateInit) func() error
		SignalingState() RTCSignalingState
		IceGatheringState() RTCIceGatheringState
		IceConnectionState() RTCIceConnectionState
		ConnectionState() RTCPeerConnectionState
		CanTrickleIceCandidates() bool
		DefaultIceServers() []RTCIceServer
		Configuration() RTCConfiguration
		SetConfiguration(RTCConfiguration)
		Close()
		OnNegotiationNeeded(func(Event)) EventHandler
		OnIceCandidate(func(RTCPeerConnectionIceEvent)) EventHandler
		OnIceCandidateError(func(RTCPeerConnectionIceErrorEvent)) EventHandler
		OnSignalingStateChange(func(Event)) EventHandler
		OnIceConnectionStateChange(func(Event)) EventHandler
		OnIceGatheringStateChange(func(Event)) EventHandler
		OnConnectionStateChange(func(Event)) EventHandler

		// http://w3c.github.io/webrtc-pc/#sec.cert-mgmt
		//GenerateCertificate(AlgorithmIdentifier) func() (RTCCertificate, error) // static TODO
		GenerateCertificate(string) func() (RTCCertificate, error) // static

		// http://w3c.github.io/webrtc-pc/#rtp-media-api
		Senders() []RTCRtpSender
		Receivers() []RTCRtpReceiver
		Transceivers() []RTCRtpTransceiver
		AddTrack(MediaStreamTrack, ...MediaStream) RTCRtpSender
		RemoveTrack(RTCRtpSender)
		AddTransceiver(MediaStreamTrack, ...RTCRtpTransceiverInit) RTCRtpTransceiver // (MediaStreamTrack or DOMString) trackOrKind
		OnTrack(func(RTCTrackEvent)) EventHandler

		// http://w3c.github.io/webrtc-pc/#rtcpeerconnection-interface-extensions-0
		SCTP() RTCSctpTransport
		CreateDataChannel(string, ...RTCDataChannelInit) RTCDataChannel
		OnDataChannel(func(RTCDataChannelEvent)) EventHandler

		// http://w3c.github.io/webrtc-pc/#rtcpeerconnection-interface-extensions-1
		Stats(...MediaStreamTrack) func() (RTCStatsReport, error)
		OnStatsEnded(func(RTCStatsEvent)) EventHandler
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectionerrorcallback
	// callback RTCPeerConnectionErrorCallback = void (DOMException error);
	RTCPeerConnectionErrorCallback interface {
		Callback
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcsessiondescriptioncallback
	// callback RTCSessionDescriptionCallback = void (RTCSessionDescriptionInit description);
	RTCSessionDescriptionCallback interface {
		Callback
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcsessiondescription
	RTCSessionDescription interface {
		Type() RTCSdpType
		Sdp() string
		ToJSON() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcicecandidate
	RTCIceCandidate interface {
		Candidate() string
		SdpMid() string
		SdpMLineIndex() uint16
		Foundation() string
		Component() RTCIceComponent
		Priority() uint
		Address() string
		Protocol() RTCIceProtocol
		Port() uint16
		Type() RTCIceCandidateType
		TcpType() RTCIceTcpCandidateType
		RelatedAddress() string
		RelatedPort() uint16
		UsernameFragment() string
		ToJSON() RTCIceCandidateInit
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectioniceevent
	RTCPeerConnectionIceEvent interface {
		Event

		Candidate() RTCIceCandidate
		URL() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectioniceerrorevent
	RTCPeerConnectionIceErrorEvent interface {
		Event

		HostCandidate() string
		URL() string
		ErrorCode() uint16
		ErrorText() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtccertificate
	RTCCertificate interface {
		Expires() time.Time
		SupportedAlgorithms() []AlgorithmIdentifier // static
		Fingerprints() []RTCDtlsFingerprint

		JSValue() jsValue
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcrtpsender
	RTCRtpSender interface {
		Track() MediaStreamTrack
		Transport() RTCDtlsTransport
		RTCPTransport() RTCDtlsTransport
		Capabilities(string) RTCRtpCapabilities // static
		SetParameters(RTCRtpSendParameters) func() error
		Parameters() RTCRtpSendParameters
		ReplaceTrack(MediaStreamTrack) func() error
		SetStreams(...MediaStream)
		Stats() func() (RTCStatsReport, error)

		// http://w3c.github.io/webrtc-pc/#rtcrtpsender-interface-extensions
		DTMF() RTCDTMFSender
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcrtpreceiver
	RTCRtpReceiver interface {
		Track() MediaStreamTrack
		Transport() RTCDtlsTransport
		RTCPTransport() RTCDtlsTransport
		Capabilities(string) RTCRtpCapabilities // static
		Parameters() RTCRtpReceiveParameters
		ContributingSources() []RTCRtpContributingSource
		SynchronizationSources() []RTCRtpSynchronizationSource
		Stats() func() (RTCStatsReport, error)
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcrtptransceiver
	RTCRtpTransceiver interface {
		Mid() string
		Sender() RTCRtpSender
		Receiver() RTCRtpReceiver
		Stopped() bool
		Direction() RTCRtpTransceiverDirection
		SetDirection(RTCRtpTransceiverDirection)
		CurrentDirection() RTCRtpTransceiverDirection
		Stop()
		SetCodecPreferences([]RTCRtpCodecCapability)
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcdtlstransport
	RTCDtlsTransport interface {
		EventTarget

		IceTransport() RTCIceTransport
		State() RTCDtlsTransportState
		RemoteCertificates() []ArrayBuffer
		OnStateChange(func(Event)) EventHandler
		OnError(func(RTCErrorEvent)) EventHandler
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcicetransport
	RTCIceTransport interface {
		EventTarget

		Role() RTCIceRole
		Component() RTCIceComponent
		State() RTCIceTransportState
		GatheringState() RTCIceGathererState
		LocalCandidates() []RTCIceCandidate
		RemoteCandidates() []RTCIceCandidate
		SelectedCandidatePair() RTCIceCandidatePair
		LocalParameters() RTCIceParameters
		RemoteParameters() RTCIceParameters
		OnStateChange(func(Event)) EventHandler
		OnGatheringStateChange(func(Event)) EventHandler
		OnSelectedCandidatePairChange(func(Event)) EventHandler
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtctrackevent
	RTCTrackEvent interface {
		Event

		Receiver() RTCRtpReceiver
		Track() MediaStreamTrack
		Streams() []MediaStream
		Transceiver() RTCRtpTransceiver
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcsctptransport
	RTCSctpTransport interface {
		EventTarget // this isn't in standart

		Transport() RTCDtlsTransport
		State() RTCSctpTransportState
		MaxMessageSize() float64
		MaxChannels() uint16
		OnStateChange(func(Event)) EventHandler
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcdatachannel
	RTCDataChannel interface {
		EventTarget

		Label() string
		Ordered() bool
		MaxPacketLifeTime() uint16
		MaxRetransmits() uint16
		Protocol() string
		Negotiated() bool
		Id() uint16
		Priority() RTCPriorityType
		ReadyState() RTCDataChannelState
		BufferedAmount() uint
		BufferedAmountLowThreshold() uint
		OnOpen(func(Event)) EventHandler
		OnBufferedAmountLow(func(Event)) EventHandler
		OnError(func(RTCErrorEvent)) EventHandler
		OnClose(func(Event)) EventHandler
		Close()
		OnMessage(func(MessageEvent)) EventHandler
		BinaryType() string
		Send(interface{}) // string, Blob, ArrayBuffer, ArrayBufferView
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcdatachannelevent
	RTCDataChannelEvent interface {
		Event

		Channel() RTCDataChannel
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcdtmfsender
	RTCDTMFSender interface {
		EventTarget

		InsertDTMF(string, ...uint)
		OnToneChange(func(RTCDTMFToneChangeEvent)) EventHandler
		CanInsertDTMF() bool
		ToneBuffer() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcdtmftonechangeevent
	RTCDTMFToneChangeEvent interface {
		Event

		Tone() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcstatsreport
	RTCStatsReport interface {
		// TODO https://www.w3.org/TR/webrtc-stats/
		/*
			Map() map[string]RTCStats // TODO
			Get(string) RTCStats
			Has(string) RTCStats
			Values() []RTCStats
			Keys() []string
		*/
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcstatsevent
	RTCStatsEvent interface {
		Event

		Report() RTCStatsReport
	}

	// http://w3c.github.io/webrtc-pc/#dfn-rtcerror
	RTCError interface {
		ErrorDetail() RTCErrorDetailType
		SDPLineNumber() int
		HttpRequestStatusCode() int
		SCTPCauseCode() int
		ReceivedAlert() uint
		SentAlert() uint
		Message() string
		Name() string
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcerrorevent
	RTCErrorEvent interface {
		Event

		Error() RTCError
	}
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicecredentialtype
type RTCIceCredentialType string

const (
	RTCIceCredentialTypePassword RTCIceCredentialType = "password"
	RTCIceCredentialTypeOAuth    RTCIceCredentialType = "oauth"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicetransportpolicy
type RTCIceTransportPolicy string

const (
	RTCIceTransportPolicyRelay RTCIceTransportPolicy = "relay"
	RTCIceTransportPolicyAll   RTCIceTransportPolicy = "all"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcbundlepolicy
type RTCBundlePolicy string

const (
	RTCBundlePolicyBalanced  RTCBundlePolicy = "balanced"
	RTCBundlePolicyMaxCompat RTCBundlePolicy = "max-compat"
	RTCBundlePolicyMaxBundle RTCBundlePolicy = "max-bundle"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcrtcpmuxpolicy
type RTCRtcpMuxPolicy string

const (
	RTCRtcpMuxPolicyNegotiate RTCRtcpMuxPolicy = "negotiate"
	RTCRtcpMuxPolicyRequire   RTCRtcpMuxPolicy = "require"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcsignalingstate
type RTCSignalingState string

const (
	RTCSignalingStateStable             RTCSignalingState = "stable"
	RTCSignalingStateHaveLocalOffer     RTCSignalingState = "have-local-offer"
	RTCSignalingStateHaveRemoteOffer    RTCSignalingState = "have-remote-offer"
	RTCSignalingStateHaveLocalPranswer  RTCSignalingState = "have-local-pranswer"
	RTCSignalingStateHaveRemotePranswer RTCSignalingState = "have-remote-pranswer"
	RTCSignalingStateClosed             RTCSignalingState = "closed"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicegatheringstate
type RTCIceGatheringState string

const (
	RTCIceGatheringStateNew       RTCIceGatheringState = "new"
	RTCIceGatheringStateGathering RTCIceGatheringState = "gathering"
	RTCIceGatheringStateComplete  RTCIceGatheringState = "complete"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectionstate
type RTCPeerConnectionState string

const (
	RTCPeerConnectionStateClosed       RTCPeerConnectionState = "closed"
	RTCPeerConnectionStateFailed       RTCPeerConnectionState = "failed"
	RTCPeerConnectionStateDisconnected RTCPeerConnectionState = "disconnected"
	RTCPeerConnectionStateNew          RTCPeerConnectionState = "new"
	RTCPeerConnectionStateConnecting   RTCPeerConnectionState = "connecting"
	RTCPeerConnectionStateConnected    RTCPeerConnectionState = "connected"
)

// http://w3c.github.io/webrtc-pc/#dom-rtciceconnectionstate
type RTCIceConnectionState string

const (
	RTCIceConnectionStateClosed       RTCIceConnectionState = "closed"
	RTCIceConnectionStateFailed       RTCIceConnectionState = "failed"
	RTCIceConnectionStateDisconnected RTCIceConnectionState = "disconnected"
	RTCIceConnectionStateNew          RTCIceConnectionState = "new"
	RTCIceConnectionStateChecking     RTCIceConnectionState = "checking"
	RTCIceConnectionStateCompleted    RTCIceConnectionState = "completed"
	RTCIceConnectionStateConnected    RTCIceConnectionState = "connected"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcsdptype
type RTCSdpType string

const (
	RTCSdpTypeOffer    RTCSdpType = "offer"
	RTCSdpTypePranswer RTCSdpType = "pranswer"
	RTCSdpTypeAnswer   RTCSdpType = "answer"
	RTCSdpTypeRollback RTCSdpType = "rollback"
)

// http://w3c.github.io/webrtc-pc/#dom-rtciceprotocol
type RTCIceProtocol string

const (
	RTCIceProtocolUDP RTCIceProtocol = "udp"
	RTCIceProtocolTCP RTCIceProtocol = "tcp"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicetcpcandidatetype
type RTCIceTcpCandidateType string

const (
	RTCIceTcpCandidateTypeActive  RTCIceTcpCandidateType = "active"
	RTCIceTcpCandidateTypePassive RTCIceTcpCandidateType = "passive"
	RTCIceTcpCandidateTypeSo      RTCIceTcpCandidateType = "so"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicecandidatetype
type RTCIceCandidateType string

const (
	RTCIceCandidateTypeHost  RTCIceCandidateType = "host"
	RTCIceCandidateTypeSrflx RTCIceCandidateType = "srflx"
	RTCIceCandidateTypePrflx RTCIceCandidateType = "prflx"
	RTCIceCandidateTypeRelay RTCIceCandidateType = "relay"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcprioritytype
type RTCPriorityType string

const (
	RTCPriorityTypeVeryLow RTCPriorityType = "very-low"
	RTCPriorityTypeLow     RTCPriorityType = "low"
	RTCPriorityTypeMedium  RTCPriorityType = "medium"
	RTCPriorityTypeHigh    RTCPriorityType = "high"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcrtptransceiverdirection
type RTCRtpTransceiverDirection string

const (
	RTCRtpTransceiverDirectionSendRecv RTCRtpTransceiverDirection = "sendrecv"
	RTCRtpTransceiverDirectionSendOnly RTCRtpTransceiverDirection = "sendonly"
	RTCRtpTransceiverDirectionRecvOnly RTCRtpTransceiverDirection = "recvonly"
	RTCRtpTransceiverDirectionInactive RTCRtpTransceiverDirection = "inactive"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcdtxstatus
type RTCDtxStatus string

const (
	RTCDtxStatusDisabled RTCDtxStatus = "disabled"
	RTCDtxStatusEnabled  RTCDtxStatus = "enabled"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcdegradationpreference
type RTCDegradationPreference string

const (
	RTCDegradationPreferenceMaintainFramerate  RTCDegradationPreference = "maintain-framerate"
	RTCDegradationPreferenceMaintainResolution RTCDegradationPreference = "maintain-resolution"
	RTCDegradationPreferenceBalanced           RTCDegradationPreference = "balanced"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcdtlstransportstate
type RTCDtlsTransportState string

const (
	RTCDtlsTransportStateNew        RTCDtlsTransportState = "new"
	RTCDtlsTransportStateConnecting RTCDtlsTransportState = "connecting"
	RTCDtlsTransportStateConnected  RTCDtlsTransportState = "connected"
	RTCDtlsTransportStateClosed     RTCDtlsTransportState = "closed"
	RTCDtlsTransportStateFailed     RTCDtlsTransportState = "failed"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicegathererstate
type RTCIceGathererState string

const (
	RTCIceGathererStateNew       RTCIceGathererState = "new"
	RTCIceGathererStateGathering RTCIceGathererState = "gathering"
	RTCIceGathererStateComplete  RTCIceGathererState = "complete"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicetransportstate
type RTCIceTransportState string

const (
	RTCIceTransportStateNew          RTCIceTransportState = "new"
	RTCIceTransportStateChecking     RTCIceTransportState = "checking"
	RTCIceTransportStateConnected    RTCIceTransportState = "connected"
	RTCIceTransportStateCompleted    RTCIceTransportState = "completed"
	RTCIceTransportStateDisconnected RTCIceTransportState = "disconnected"
	RTCIceTransportStateFailed       RTCIceTransportState = "failed"
	RTCIceTransportStateClosed       RTCIceTransportState = "closed"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicerole
type RTCIceRole string

const (
	RTCIceRoleControlling RTCIceRole = "controlling"
	RTCIceRoleControlled  RTCIceRole = "controlled"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcicecomponent
type RTCIceComponent string

const (
	RTCIceComponentRTP  RTCIceComponent = "rtp"
	RTCIceComponentRTCP RTCIceComponent = "rtcp"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcsctptransportstate
type RTCSctpTransportState string

const (
	RTCSctpTransportStateConnecting RTCSctpTransportState = "connecting"
	RTCSctpTransportStateConnected  RTCSctpTransportState = "connected"
	RTCSctpTransportStateClosed     RTCSctpTransportState = "closed"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcdatachannelstate

type RTCDataChannelState string

const (
	RTCDataChannelStateConnecting RTCDataChannelState = "connecting"
	RTCDataChannelStateOpen       RTCDataChannelState = "open"
	RTCDataChannelStateClosing    RTCDataChannelState = "closing"
	RTCDataChannelStateClosed     RTCDataChannelState = "closed"
)

// http://w3c.github.io/webrtc-pc/#dom-rtcerrordetailtype
type RTCErrorDetailType string

const (
	RTCErrorDetailTypeDataChannelFailure          RTCErrorDetailType = "data-channel-failure"
	RTCErrorDetailTypeDTLSFailure                 RTCErrorDetailType = "dtls-failure"
	RTCErrorDetailTypeFingerprintFailure          RTCErrorDetailType = "fingerprint-failure"
	RTCErrorDetailTypeIdpBadScriptFailure         RTCErrorDetailType = "idp-bad-script-failure"
	RTCErrorDetailTypeIdpExecutionFailure         RTCErrorDetailType = "idp-execution-failure"
	RTCErrorDetailTypeIdpLoadFailure              RTCErrorDetailType = "idp-load-failure"
	RTCErrorDetailTypeIdpNeedLogin                RTCErrorDetailType = "idp-need-login"
	RTCErrorDetailTypeIdpTimeout                  RTCErrorDetailType = "idp-timeout"
	RTCErrorDetailTypeIdpTLSFailure               RTCErrorDetailType = "idp-tls-failure"
	RTCErrorDetailTypeIdpTokenExpired             RTCErrorDetailType = "idp-token-expired"
	RTCErrorDetailTypeIdpTokenInvalid             RTCErrorDetailType = "idp-token-invalid"
	RTCErrorDetailTypeSCTPFailure                 RTCErrorDetailType = "sctp-failure"
	RTCErrorDetailTypeSDPSyntaxError              RTCErrorDetailType = "sdp-syntax-error"
	RTCErrorDetailTypeHardwareEncoderNotAvailable RTCErrorDetailType = "hardware-encoder-not-available"
	RTCErrorDetailTypeHardwareEncoderError        RTCErrorDetailType = "hardware-encoder-error"
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

// http://w3c.github.io/webrtc-pc/#dom-rtcconfiguration
type RTCConfiguration struct {
	IceServers           []RTCIceServer
	IceTransportPolicy   RTCIceTransportPolicy // all
	BundlePolicy         RTCBundlePolicy       // balanced
	RTCPMuxPolicy        RTCRtcpMuxPolicy      // require
	PeerIdentity         string
	Certificates         []RTCCertificate
	IceCandidatePoolSize uint8 // 0
}

func wrapRTCConfiguration(v Value) RTCConfiguration {
	c := RTCConfiguration{}
	if v.valid() {
		c.IceServers = toRTCIceServerSlice(v.get("iceServers"))
		c.IceTransportPolicy = RTCIceTransportPolicy(v.get("iceTransportPolicy").toString())
		c.BundlePolicy = RTCBundlePolicy(v.get("bundlePolicy").toString())
		c.RTCPMuxPolicy = RTCRtcpMuxPolicy(v.get("rtcpMuxPolicy").toString())
		c.PeerIdentity = v.get("peerIdentity").toString()
		c.Certificates = toRTCCertificateSlice(v.get("certificates"))
		c.IceCandidatePoolSize = v.get("iceCandidatePoolSize").toUint8()
	}
	return c
}

func (p RTCConfiguration) JSValue() jsValue {
	o := jsObject.New()
	o.Set("iceServers", ToJSArray(p.IceServers))
	o.Set("iceTransportPolicy", string(p.IceTransportPolicy))
	o.Set("bundlePolicy", string(p.BundlePolicy))
	o.Set("rtcpMuxPolicy", string(p.RTCPMuxPolicy))
	o.Set("peerIdentity", p.PeerIdentity)
	o.Set("certificates", ToJSArray(p.Certificates))
	o.Set("iceCandidatePoolSize", p.IceCandidatePoolSize)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcoauthcredential
type RTCOAuthCredential struct {
	MacKey      string
	AccessToken string
}

func (p RTCOAuthCredential) JSValue() jsValue {
	o := jsObject.New()
	o.Set("macKey", p.MacKey)
	o.Set("accessToken", p.AccessToken)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtciceserver
type RTCIceServer struct {
	URLs           []string // required (DOMString or sequence<DOMString>)
	Username       string
	Credential     string               //  (DOMString or RTCOAuthCredential) TODO
	CredentialType RTCIceCredentialType // "password"
}

func wrapRTCIceServer(v Value) RTCIceServer {
	s := RTCIceServer{}
	if v.valid() {
		s.URLs = stringSequenceToSlice(v.get("urls"))
		s.Username = v.get("username").toString()
		s.Credential = v.get("credential").toString()
		s.CredentialType = RTCIceCredentialType(v.get("credentialType").toString())
	}
	return s
}

func (p RTCIceServer) JSValue() jsValue {
	o := jsObject.New()
	o.Set("urls", ToJSArray(p.URLs))
	o.Set("username", p.Username)
	o.Set("credential", p.Credential)
	o.Set("credentialType", string(p.CredentialType))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcofferansweroptions
type RTCOfferAnswerOptions struct {
	VoiceActivityDetection bool // true
}

func (p RTCOfferAnswerOptions) JSValue() jsValue {
	o := jsObject.New()
	o.Set("voiceActivityDetection", p.VoiceActivityDetection)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcofferoptions
type RTCOfferOptions struct {
	RTCOfferAnswerOptions

	IceRestart bool // false
}

func (p RTCOfferOptions) JSValue() jsValue {
	o := p.RTCOfferAnswerOptions.JSValue()
	o.Set("iceRestart", p.IceRestart)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcansweroptions
type RTCAnswerOptions struct {
	RTCOfferAnswerOptions
}

// http://w3c.github.io/webrtc-pc/#dom-rtccertificateexpiration
type RTCCertificateExpiration struct {
	Expires time.Time
}

func (p RTCCertificateExpiration) JSValue() jsValue {
	o := jsObject.New()
	o.Set("expires", p.Expires.UnixNano()/int64(time.Millisecond))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcsessiondescriptioninit
type RTCSessionDescriptionInit struct {
	Type RTCSdpType // required
	SDP  string     // ""
}

func wrapRTCSessionDescriptionInit(v Value) RTCSessionDescriptionInit {
	ret := RTCSessionDescriptionInit{}
	if v.valid() {
		ret.Type = RTCSdpType(v.get("type").toString())
		ret.SDP = v.get("sdp").toString()
	}
	return ret
}

func (p RTCSessionDescriptionInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("type", string(p.Type))
	o.Set("sdp", p.SDP)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcicecandidateinit
type RTCIceCandidateInit struct {
	Candidate        string
	SdpMid           string
	SdpMLineIndex    uint16
	UsernameFragment string
}

func wrapRTCIceCandidateInit(v Value) RTCIceCandidateInit {
	i := RTCIceCandidateInit{}
	if v.valid() {
		i.Candidate = v.get("candidate").toString()
		i.SdpMid = v.get("sdpMid").toString()
		i.SdpMLineIndex = v.get("sdpMLineIndex").toUint16()
		i.UsernameFragment = v.get("usernameFragment").toString()
	}
	return i
}

func (p RTCIceCandidateInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("candidate", p.Candidate)
	o.Set("sdpMid", p.SdpMid)
	o.Set("sdpMLineIndex", p.SdpMLineIndex)
	o.Set("usernameFragment", p.UsernameFragment)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectioniceeventinit
type RTCPeerConnectionIceEventInit struct {
	EventInit

	Candidate RTCIceCandidate
	URL       string
}

func (p RTCPeerConnectionIceEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("candidate", JSValueOf(p.Candidate))
	o.Set("url", p.URL)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcpeerconnectioniceerroreventinit
type RTCPeerConnectionIceErrorEventInit struct {
	EventInit

	HostCandidate string
	URL           string
	ErrorCode     uint16 // required
	StatusText    string
}

func (p RTCPeerConnectionIceErrorEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("hostCandidate", p.HostCandidate)
	o.Set("url", p.URL)
	o.Set("errorCode", p.ErrorCode)
	o.Set("statusText", p.StatusText)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtptransceiverinit
type RTCRtpTransceiverInit struct {
	Direction     RTCRtpTransceiverDirection
	Streams       []MediaStream
	SendEncodings []RTCRtpEncodingParameters
}

func (p RTCRtpTransceiverInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("direction", string(p.Direction))
	o.Set("streams", ToJSArray(p.Streams))
	o.Set("sendEncodings", ToJSArray(p.SendEncodings))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpparameters
type RTCRtpParameters struct {
	HeaderExtensions []RTCRtpHeaderExtensionParameters
	RTCP             RTCRtcpParameters
	Codecs           []RTCRtpCodecParameters
}

func wrapRTCRtpParameters(v Value) RTCRtpParameters {
	p := RTCRtpParameters{}
	if v.valid() {
		p.HeaderExtensions = toRTCRtpHeaderExtensionParametersSlice(v.get("headerExtensions"))
		p.RTCP = wrapRTCRtcpParameters(v.get("rtcp"))
		p.Codecs = toRTCRtpCodecParametersSlice(v.get("codecs"))
	}
	return p
}

func (p RTCRtpParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("headerExtensions", ToJSArray(p.HeaderExtensions))
	o.Set("rtcp", p.RTCP.JSValue())
	o.Set("codecs", ToJSArray(p.Codecs))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpsendparameters
type RTCRtpSendParameters struct {
	RTCRtpParameters

	TransactionId         string
	Encodings             []RTCRtpEncodingParameters
	DegradationPreference RTCDegradationPreference // balanced
	Priority              RTCPriorityType          // low
}

func wrapRTCRtpSendParameters(v Value) RTCRtpSendParameters {
	p := RTCRtpSendParameters{}
	if v.valid() {
		p.RTCRtpParameters = wrapRTCRtpParameters(v)
		p.TransactionId = v.get("transactionId").toString()
		p.Encodings = toRTCRtpEncodingParametersSlice(v.get("encodings"))
		p.DegradationPreference = RTCDegradationPreference(v.get("degradationPreference").toString())
		p.Priority = RTCPriorityType(v.get("priority").toString())
	}

	return p
}

func (p RTCRtpSendParameters) JSValue() jsValue {
	o := p.RTCRtpParameters.JSValue()
	o.Set("transactionId", p.TransactionId)
	o.Set("encodings", ToJSArray(p.Encodings))
	o.Set("degradationPreference", string(p.DegradationPreference))
	o.Set("priority", string(p.Priority))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpreceiveparameters
type RTCRtpReceiveParameters struct {
	RTCRtpParameters

	Encodings []RTCRtpDecodingParameters
}

func wrapRTCRtpReceiveParameters(v Value) RTCRtpReceiveParameters {
	p := RTCRtpReceiveParameters{}
	if v.valid() {
		p.Encodings = toRTCRtpDecodingParametersSlice(v)
	}
	return p
}

func (p RTCRtpReceiveParameters) JSValue() jsValue {
	o := p.RTCRtpParameters.JSValue()
	o.Set("encodings", ToJSArray(p.Encodings))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcodingparameters
type RTCRtpCodingParameters struct {
	RID string
}

func wrapRTCRtpCodingParameters(v Value) RTCRtpCodingParameters {
	p := RTCRtpCodingParameters{}
	if v.valid() {
		p.RID = v.get("rid").toString()
	}
	return p
}

func (p RTCRtpCodingParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("rid", p.RID)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpdecodingparameters
type RTCRtpDecodingParameters struct {
	RTCRtpCodingParameters
}

func wrapRTCRtpDecodingParameters(v Value) RTCRtpDecodingParameters {
	p := RTCRtpDecodingParameters{}
	if v.valid() {
		p.RTCRtpCodingParameters = wrapRTCRtpCodingParameters(v)
	}
	return p
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpencodingparameters
type RTCRtpEncodingParameters struct {
	RTCRtpCodingParameters

	CodecPayloadType      uint8
	DTX                   RTCDtxStatus
	Active                bool // true
	PTime                 uint
	MaxBitrate            uint
	MaxFramerate          float64
	ScaleResolutionDownBy float64
}

func wrapRTCRtpEncodingParameters(v Value) RTCRtpEncodingParameters {
	p := RTCRtpEncodingParameters{}
	if v.valid() {
		p.RTCRtpCodingParameters = wrapRTCRtpCodingParameters(v)
		p.CodecPayloadType = v.get("codecPayloadType").toUint8()
		p.DTX = RTCDtxStatus(v.get("dtx").toString())
		p.Active = v.get("active").toBool()
		p.PTime = v.get("ptime").toUint()
		p.MaxBitrate = v.get("maxBitrate").toUint()
		p.MaxFramerate = v.get("maxFramerate").toFloat64()
		p.ScaleResolutionDownBy = v.get("scaleResolutionDownBy").toFloat64()
	}
	return p
}

func (p RTCRtpEncodingParameters) JSValue() jsValue {
	o := p.RTCRtpCodingParameters.JSValue()
	o.Set("codecPayloadType", p.CodecPayloadType)
	o.Set("dtx", string(p.DTX))
	o.Set("active", p.Active)
	o.Set("ptime", p.PTime)
	o.Set("maxBitrate", p.MaxBitrate)
	o.Set("maxFramerate", p.MaxFramerate)
	o.Set("scaleResolutionDownBy", p.ScaleResolutionDownBy)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtcpparameters
type RTCRtcpParameters struct {
	CName       string
	ReducedSize bool
}

func wrapRTCRtcpParameters(v Value) RTCRtcpParameters {
	p := RTCRtcpParameters{}
	if v.valid() {
		p.CName = v.get("cname").toString()
		p.ReducedSize = v.get("reducedSize").toBool()
	}
	return p
}

func (p RTCRtcpParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("cname", p.CName)
	o.Set("reducedSize", p.ReducedSize)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpheaderextensionparameters
type RTCRtpHeaderExtensionParameters struct {
	URI       string
	Id        uint16
	Encrypted bool // false
}

func wrapRTCRtpHeaderExtensionParameters(v Value) RTCRtpHeaderExtensionParameters {
	p := RTCRtpHeaderExtensionParameters{}
	if v.valid() {
		p.URI = v.get("uri").toString()
		p.Id = v.get("id").toUint16()
		p.Encrypted = v.get("encrypted").toBool()
	}

	return p
}

func (p RTCRtpHeaderExtensionParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("uri", p.URI)
	o.Set("id", p.Id)
	o.Set("encrypted", p.Encrypted)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcodecparameters
type RTCRtpCodecParameters struct {
	PayloadType uint8
	MimeType    string
	ClockRate   uint
	Channels    uint16
	SDPFmtpLine string
}

func wrapRTCRtpCodecParameters(v Value) RTCRtpCodecParameters {
	p := RTCRtpCodecParameters{}
	if v.valid() {
		p.PayloadType = v.get("payloadType").toUint8()
		p.MimeType = v.get("mimeType").toString()
		p.ClockRate = v.get("clockRate").toUint()
		p.Channels = v.get("channels").toUint16()
		p.SDPFmtpLine = v.get("sdpFmtpLine").toString()
	}
	return p
}

func (p RTCRtpCodecParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("payloadType", p.PayloadType)
	o.Set("mimeType", p.MimeType)
	o.Set("clockRate", p.ClockRate)
	o.Set("channels", p.Channels)
	o.Set("sdpFmtpLine", p.SDPFmtpLine)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcapabilities
type RTCRtpCapabilities struct {
	Codecs           []RTCRtpCodecCapability
	HeaderExtensions []RTCRtpHeaderExtensionCapability
}

func wrapRTCRtpCapabilities(v Value) RTCRtpCapabilities {
	c := RTCRtpCapabilities{}
	if v.valid() {
		c.Codecs = toRTCRtpCodecCapabilitySlice(v.get("codecs"))
		c.HeaderExtensions = toRTCRtpHeaderExtensionCapabilitySlice(v.get("headerExtensions"))
	}
	return c
}

func (p RTCRtpCapabilities) JSValue() jsValue {
	o := jsObject.New()
	o.Set("codecs", ToJSArray(p.Codecs))
	o.Set("headerExtensions", ToJSArray(p.HeaderExtensions))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcodeccapability
type RTCRtpCodecCapability struct {
	MimeType    string
	ClockRate   uint
	Channels    uint16
	SDPFmtpLine string
}

func wrapRTCRtpCodecCapability(v Value) RTCRtpCodecCapability {
	c := RTCRtpCodecCapability{}
	if v.valid() {
		c.MimeType = v.get("mimeType").toString()
		c.ClockRate = v.get("clockRate").toUint()
		c.Channels = v.get("channels").toUint16()
		c.SDPFmtpLine = v.get("sdpFmtpLine").toString()
	}
	return c
}

func (p RTCRtpCodecCapability) JSValue() jsValue {
	o := jsObject.New()
	o.Set("mimeType", p.MimeType)
	o.Set("clockRate", p.ClockRate)
	o.Set("channels", p.Channels)
	o.Set("sdpFmtpLine", p.SDPFmtpLine)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpheaderextensioncapability
type RTCRtpHeaderExtensionCapability struct {
	URI string
}

func wrapRTCRtpHeaderExtensionCapability(v Value) RTCRtpHeaderExtensionCapability {
	c := RTCRtpHeaderExtensionCapability{}
	if v.valid() {
		c.URI = v.get("uri").toString()
	}
	return c
}

func (p RTCRtpHeaderExtensionCapability) JSValue() jsValue {
	o := jsObject.New()
	o.Set("uri", p.URI)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcontributingsource
type RTCRtpContributingSource struct {
	Timestamp  float64 // required TODO maybe time.Time
	Source     uint    // required
	AudioLevel float64
}

func wrapRTCRtpContributingSource(v Value) RTCRtpContributingSource {
	s := RTCRtpContributingSource{}
	if v.valid() {
		s.Timestamp = v.get("timestamp").toFloat64()
		s.Source = v.get("source").toUint()
		s.AudioLevel = v.get("audioLevel").toFloat64()
	}
	return s
}

func (p RTCRtpContributingSource) JSValue() jsValue {
	o := jsObject.New()
	o.Set("timestamp", p.Timestamp)
	o.Set("source", p.Source)
	o.Set("audioLevel", p.AudioLevel)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpsynchronizationsource
type RTCRtpSynchronizationSource struct {
	RTCRtpContributingSource

	VoiceActivityFlag bool
}

func wrapRTCRtpSynchronizationSource(v Value) RTCRtpSynchronizationSource {
	s := RTCRtpSynchronizationSource{}
	if v.valid() {
		s.RTCRtpContributingSource = wrapRTCRtpContributingSource(v)
		s.VoiceActivityFlag = v.get("voiceActivityFlag").toBool()
	}
	return s
}

func (p RTCRtpSynchronizationSource) JSValue() jsValue {
	o := p.RTCRtpContributingSource.JSValue()
	o.Set("voiceActivityFlag", p.VoiceActivityFlag)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcdtlsfingerprint
type RTCDtlsFingerprint struct {
	Algorithm string
	Value     string
}

func wrapRTCDtlsFingerprint(v Value) RTCDtlsFingerprint {
	ret := RTCDtlsFingerprint{}
	if v.valid() {
		ret.Algorithm = v.get("algorithm").toString()
		ret.Value = v.get("value").toString()
	}
	return ret
}

func (p RTCDtlsFingerprint) JSValue() jsValue {
	o := jsObject.New()
	o.Set("algorithm", p.Algorithm)
	o.Set("value", p.Value)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtciceparameters
type RTCIceParameters struct {
	UsernameFragment string
	Password         string
}

func wrapRTCIceParameters(v Value) RTCIceParameters {
	p := RTCIceParameters{}
	if v.valid() {
		p.UsernameFragment = v.get("usernameFragment").toString()
		p.Password = v.get("password").toString()
	}
	return p
}

func (p RTCIceParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("usernameFragment", p.UsernameFragment)
	o.Set("password", p.Password)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcicecandidatepair
type RTCIceCandidatePair struct {
	Local  RTCIceCandidate
	Remote RTCIceCandidate
}

func wrapRTCIceCandidatePair(v Value) RTCIceCandidatePair {
	p := RTCIceCandidatePair{}
	if v.valid() {
		p.Local = wrapRTCIceCandidate(v.get("local"))
		p.Remote = wrapRTCIceCandidate(v.get("remote"))
	}
	return p
}

func (p RTCIceCandidatePair) JSValue() jsValue {
	o := jsObject.New()
	o.Set("local", JSValueOf(p.Local))
	o.Set("remote", JSValueOf(p.Remote))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtctrackeventinit
type RTCTrackEventInit struct {
	EventInit

	Receiver    RTCRtpReceiver
	Track       MediaStreamTrack
	Streams     []MediaStream
	Transceiver RTCRtpTransceiver
}

func (p RTCTrackEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("receiver", JSValueOf(p.Receiver))
	o.Set("track", JSValueOf(p.Track))
	o.Set("streams", ToJSArray(p.Streams))
	o.Set("transceiver", JSValueOf(p.Transceiver))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcdatachannelinit
type RTCDataChannelInit struct {
	Ordered           bool //true
	MaxPacketLifeTime uint16
	MaxRetransmits    uint16
	Protocol          string // ""
	Negotiated        bool   // false
	Id                uint16
	Priority          RTCPriorityType // "low"
}

func (p RTCDataChannelInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("ordered", p.Ordered)
	o.Set("maxPacketLifeTime", p.MaxPacketLifeTime)
	o.Set("maxRetransmits", p.MaxRetransmits)
	o.Set("protocol", p.Protocol)
	o.Set("negotiated", p.Negotiated)
	o.Set("id", p.Id)
	o.Set("priority", string(p.Priority))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcdatachanneleventinit
type RTCDataChannelEventInit struct {
	EventInit

	Channel RTCDataChannel
}

func (p RTCDataChannelEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("channel", JSValueOf(p.Channel))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcdtmftonechangeeventinit
type RTCDTMFToneChangeEventInit struct {
	EventInit

	Tone string
}

func (p RTCDTMFToneChangeEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("tone", p.Tone)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcstats
type RTCStats struct {
	Timestamp float64      // required
	Type      RTCStatsType // required
	Id        string       // required
}

func (p RTCStats) JSValue() jsValue {
	o := jsObject.New()
	o.Set("timestamp", p.Timestamp)
	o.Set("type", string(p.Type))
	o.Set("id", p.Id)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcstatseventinit
type RTCStatsEventInit struct {
	EventInit

	Report RTCStatsReport
}

func (p RTCStatsEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("report", JSValueOf(p.Report))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcerroreventinit
type RTCErrorEventInit struct {
	EventInit

	Error RTCError
}

func (p RTCErrorEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("error", JSValueOf(p.Error))
	return o
}

/*
func wrapRTCIceServer(v Value) RTCIceServer {
	return RTCIceServer{}
}

func wrapRTCRtpSender(v Value) RTCRtpSender {
	return nil
}

func wrapRTCRtpReceiver(v Value) RTCRtpReceiver {
	return nil
}

func wrapRTCRtpTransceiver(v Value) RTCRtpTransceiver {
	return nil
}

func wrapRTCPeerConnection(v Value) RTCPeerConnection {
	return nil
}
*/
