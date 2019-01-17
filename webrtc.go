// +build js,wasm

// http://w3c.github.io/webrtc-pc/
package wasm

import (
	"time"
)

// https://www.w3.org/TR/WebCryptoAPI/#dfn-AlgorithmIdentifier
type AlgorithmIdentifier string

// https://heycam.github.io/webidl/#common-DOMTimeStamp
// typedef unsigned long long DOMTimeStamp;

func NewRTCPeerConnection(configuration ...RTCConfiguration) RTCPeerConnection {
	return nil
}

func NewRTCSessionDescription(descriptionInitDict RTCSessionDescriptionInit) RTCSessionDescription {
	return nil
}

func NewRTCIceCandidate(candidateInitDict ...RTCIceCandidateInit) RTCIceCandidate {
	return nil
}

func NewRTCPeerConnectionIceEvent(typ string, eventInitDict ...RTCPeerConnectionIceEventInit) RTCPeerConnectionIceEvent {
	return nil
}

func NewRTCPeerConnectionIceErrorEvent(typ string, eventInitDict RTCPeerConnectionIceErrorEventInit) RTCPeerConnectionIceErrorEvent {
	return nil
}

func NewRTCTrackEvent(typ string, eventInitDict RTCTrackEventInit) RTCTrackEvent {
	return nil
}

func NewRTCDataChannelEvent(typ string, eventInitDict RTCDataChannelEventInit) RTCDataChannelEvent {
	return nil
}

func NewRTCDTMFToneChangeEvent(typ string, eventInitDict RTCDTMFToneChangeEventInit) RTCDTMFToneChangeEvent {
	return nil
}

func NewRTCStatsEvent(typ string, eventInitDict RTCStatsEventInit) RTCStatsEvent {
	return nil
}

func NewRTCError(detail RTCErrorDetailType, message string) RTCError {
	return nil
}

func NewRTCErrorEvent(typ string, eventInitDict RTCErrorEventInit) RTCErrorEvent {
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
		GenerateCertificate(AlgorithmIdentifier) func() (RTCCertificate, error) // static

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
		Dtmf() RTCDTMFSender
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
		Map() map[string]RTCStats // TODO
	}

	// http://w3c.github.io/webrtc-pc/#dom-rtcstatsevent
	RTCStatsEvent interface {
		Event

		Report() RTCStatsReport
	}

	// http://w3c.github.io/webrtc-pc/#dfn-rtcerror
	RTCError interface {
		ErrorDetail() RTCErrorDetailType
		SDPLineNumber() uint
		HttpRequestStatusCode() uint
		SCTPCauseCode() uint
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

func (p RTCConfiguration) JSValue() jsValue {
	o := jsObject.New()
	o.Set("iceServers", p.IceServers) // TODO
	o.Set("iceTransportPolicy", string(p.IceTransportPolicy))
	o.Set("bundlePolicy", string(p.BundlePolicy))
	o.Set("rtcpMuxPolicy", string(p.RTCPMuxPolicy))
	o.Set("peerIdentity", p.PeerIdentity)
	o.Set("certificates", p.Certificates) // TODO
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

func (p RTCIceServer) JSValue() jsValue {
	o := jsObject.New()
	o.Set("urls", p.URLs) // TODO
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
	Sdp  string     // ""
}

func (p RTCSessionDescriptionInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("type", string(p.Type))
	o.Set("sdp", p.Sdp)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcicecandidateinit
type RTCIceCandidateInit struct {
	Candidate        string
	SdpMid           string
	SdpMLineIndex    uint16
	UsernameFragment string
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
	o.Set("streams", nil)       // TODO
	o.Set("sendEncodings", nil) // TODO
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpparameters
type RTCRtpParameters struct {
	HeaderExtensions []RTCRtpHeaderExtensionParameters
	RTCP             RTCRtcpParameters
	Codecs           []RTCRtpCodecParameters
}

func (p RTCRtpParameters) JSValue() jsValue {
	o := jsObject.New()
	o.Set("headerExtensions", p.HeaderExtensions) // TODO
	o.Set("rtcp", p.RTCP)
	o.Set("codecs", p.Codecs) // TODO
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

func (p RTCRtpSendParameters) JSValue() jsValue {
	o := p.RTCRtpParameters.JSValue()
	o.Set("transactionId", p.TransactionId)
	o.Set("encodings", p.Encodings) // TODO
	o.Set("degradationPreference", string(p.DegradationPreference))
	o.Set("priority", string(p.Priority))
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpreceiveparameters
type RTCRtpReceiveParameters struct {
	RTCRtpParameters

	Encodings []RTCRtpDecodingParameters
}

func (p RTCRtpReceiveParameters) JSValue() jsValue {
	o := p.RTCRtpParameters.JSValue()
	o.Set("encodings", p.Encodings) // TODO
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcodingparameters
type RTCRtpCodingParameters struct {
	RID string
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

func (p RTCRtpCapabilities) JSValue() jsValue {
	o := jsObject.New()
	o.Set("codecs", p.Codecs)                     // TODO
	o.Set("headerExtensions", p.HeaderExtensions) // TODO
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcodeccapability
type RTCRtpCodecCapability struct {
	MimeType    string
	ClockRate   string
	Channels    uint16
	SDPFmtpLine string
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

func (p RTCRtpHeaderExtensionCapability) JSValue() jsValue {
	o := jsObject.New()
	o.Set("uri", p.URI)
	return o
}

// http://w3c.github.io/webrtc-pc/#dom-rtcrtpcontributingsource
type RTCRtpContributingSource struct {
	Timestamp  float64 // required
	Source     uint    // required
	AudioLevel float64
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
	o.Set("streams", p.Streams) // TODO to js Array
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
