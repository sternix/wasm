// +build js,wasm

package wasm

import (
	"time"
)

type rtcPeerConnectionImpl struct {
	*eventTargetImpl
}

func wrapRTCPeerConnection(v Value) RTCPeerConnection {
	if v.valid() {
		return &rtcPeerConnectionImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcPeerConnectionImpl) CreateOffer(options ...RTCOfferOptions) func() (RTCSessionDescriptionInit, error) {
	return func() (RTCSessionDescriptionInit, error) {
		var (
			result Value
			ok     bool
		)

		switch len(options) {
		case 0:
			result, ok = await(p.call("createOffer"))
		default:
			result, ok = await(p.call("createOffer", options[0].JSValue()))
		}

		if ok {
			return wrapRTCSessionDescriptionInit(result), nil
		}

		return RTCSessionDescriptionInit{}, wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) CreateAnswer(options ...RTCAnswerOptions) func() (RTCSessionDescriptionInit, error) {
	return func() (RTCSessionDescriptionInit, error) {
		var (
			result Value
			ok     bool
		)

		switch len(options) {
		case 0:
			result, ok = await(p.call("createAnswer"))
		default:
			result, ok = await(p.call("createAnswer", options[0].JSValue()))
		}

		if ok {
			return wrapRTCSessionDescriptionInit(result), nil
		}

		return RTCSessionDescriptionInit{}, wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) SetLocalDescription(description RTCSessionDescriptionInit) func() error {
	return func() error {
		result, ok := await(p.call("setLocalDescription", description.JSValue()))
		if ok {
			return nil
		}
		return wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) LocalDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("localDescription"))
}

func (p *rtcPeerConnectionImpl) CurrentLocalDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("currentLocalDescription"))
}

func (p *rtcPeerConnectionImpl) PendingLocalDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("pendingLocalDescription"))
}

func (p *rtcPeerConnectionImpl) SetRemoteDescription(description RTCSessionDescriptionInit) func() error {
	return func() error {
		result, ok := await(p.call("setRemoteDescription", description.JSValue()))
		if ok {
			return nil
		}
		return wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) RemoteDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("remoteDescription"))
}

func (p *rtcPeerConnectionImpl) CurrentRemoteDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("currentRemoteDescription"))
}

func (p *rtcPeerConnectionImpl) PendingRemoteDescription() RTCSessionDescription {
	return wrapRTCSessionDescription(p.get("pendingRemoteDescription"))
}

func (p *rtcPeerConnectionImpl) AddIceCandidate(candidate RTCIceCandidateInit) func() error {
	return func() error {
		result, ok := await(p.call("addIceCandidate", candidate.JSValue()))
		if ok {
			return nil
		}
		return wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) SignalingState() RTCSignalingState {
	return RTCSignalingState(p.get("signalingState").toString())
}

func (p *rtcPeerConnectionImpl) IceGatheringState() RTCIceGatheringState {
	return RTCIceGatheringState(p.get("iceGatheringState").toString())
}

func (p *rtcPeerConnectionImpl) IceConnectionState() RTCIceConnectionState {
	return RTCIceConnectionState(p.get("iceConnectionState").toString())
}

func (p *rtcPeerConnectionImpl) ConnectionState() RTCPeerConnectionState {
	return RTCPeerConnectionState(p.get("connectionState").toString())
}

func (p *rtcPeerConnectionImpl) CanTrickleIceCandidates() bool {
	return p.get("canTrickleIceCandidates").toBool()
}

func (p *rtcPeerConnectionImpl) DefaultIceServers() []RTCIceServer {
	if v := p.call("getDefaultIceServers"); v.valid() {
		return toRTCIceServerSlice(v)
	}
	return nil
}

func (p *rtcPeerConnectionImpl) Configuration() RTCConfiguration {
	return wrapRTCConfiguration(p.call("getConfiguration"))
}

func (p *rtcPeerConnectionImpl) SetConfiguration(configuration RTCConfiguration) {
	p.call("setConfiguration", configuration.JSValue())
}

func (p *rtcPeerConnectionImpl) Close() {
	p.call("close")
}

func (p *rtcPeerConnectionImpl) OnNegotiationNeeded(fn func(Event)) EventHandler {
	return p.On("negotiationneeded", fn)
}

func (p *rtcPeerConnectionImpl) OnIceCandidate(fn func(RTCPeerConnectionIceEvent)) EventHandler {
	return p.On("icecandidate", func(e Event) {
		if ie, ok := e.(RTCPeerConnectionIceEvent); ok {
			fn(ie)
		}
	})
}

func (p *rtcPeerConnectionImpl) OnIceCandidateError(fn func(RTCPeerConnectionIceErrorEvent)) EventHandler {
	return p.On("icecandidateerror", func(e Event) {
		if ie, ok := e.(RTCPeerConnectionIceErrorEvent); ok {
			fn(ie)
		}
	})
}

func (p *rtcPeerConnectionImpl) OnSignalingStateChange(fn func(Event)) EventHandler {
	return p.On("signalingstatechange", fn)
}

func (p *rtcPeerConnectionImpl) OnIceConnectionStateChange(fn func(Event)) EventHandler {
	return p.On("iceconnectionstatechange", fn)
}

func (p *rtcPeerConnectionImpl) OnIceGatheringStateChange(fn func(Event)) EventHandler {
	return p.On("icegatheringstatechange", fn)
}

func (p *rtcPeerConnectionImpl) OnConnectionStateChange(fn func(Event)) EventHandler {
	return p.On("connectionstatechange", fn)
}

// http://w3c.github.io/webrtc-pc/#sec.cert-mgmt
func (p *rtcPeerConnectionImpl) GenerateCertificate(keygenAlgorithm string) func() (RTCCertificate, error) { // static
	return func() (RTCCertificate, error) {
		result, ok := await(p.call("generateCertificate", keygenAlgorithm))
		if ok {
			return wrapRTCCertificate(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

// http://w3c.github.io/webrtc-pc/#rtp-media-api
func (p *rtcPeerConnectionImpl) Senders() []RTCRtpSender {
	if v := p.call("getSenders"); v.valid() {
		return toRTCRtpSenderSlice(v)
	}
	return nil
}

func (p *rtcPeerConnectionImpl) Receivers() []RTCRtpReceiver {
	if v := p.call("getReceivers"); v.valid() {
		return toRTCRtpReceiverSlice(v)
	}
	return nil
}

func (p *rtcPeerConnectionImpl) Transceivers() []RTCRtpTransceiver {
	if v := p.call("RTCRtpTransceiver"); v.valid() {
		return toRTCRtpTransceiverSlice(v)
	}
	return nil
}

func (p *rtcPeerConnectionImpl) AddTrack(track MediaStreamTrack, streams ...MediaStream) RTCRtpSender {
	switch len(streams) {
	case 0:
		return wrapRTCRtpSender(p.call("addTrack", JSValueOf(track)))
	default:
		return wrapRTCRtpSender(p.call("addTrack", JSValueOf(track), JSValueOf(streams[0])))
	}
}

func (p *rtcPeerConnectionImpl) RemoveTrack(sender RTCRtpSender) {
	p.call("removeTrack", JSValueOf(sender))
}

func (p *rtcPeerConnectionImpl) AddTransceiver(trackOrKind MediaStreamTrack, tinit ...RTCRtpTransceiverInit) RTCRtpTransceiver { // (MediaStreamTrack or DOMString) trackOrKind
	switch len(tinit) {
	case 0:
		return wrapRTCRtpTransceiver(p.call("addTransceiver", JSValueOf(trackOrKind)))
	default:
		return wrapRTCRtpTransceiver(p.call("addTransceiver", JSValueOf(trackOrKind), tinit[0].JSValue()))
	}
}

func (p *rtcPeerConnectionImpl) OnTrack(fn func(RTCTrackEvent)) EventHandler {
	return p.On("track", func(e Event) {
		if ie, ok := e.(RTCTrackEvent); ok {
			fn(ie)
		}
	})
}

// http://w3c.github.io/webrtc-pc/#rtcpeerconnection-interface-extensions-0
func (p *rtcPeerConnectionImpl) SCTP() RTCSctpTransport {
	return wrapRTCSctpTransport(p.get("sctp"))
}

func (p *rtcPeerConnectionImpl) CreateDataChannel(label string, dataChannelDict ...RTCDataChannelInit) RTCDataChannel {
	switch len(dataChannelDict) {
	case 0:
		return wrapRTCDataChannel(p.call("createDataChannel", label))
	default:
		return wrapRTCDataChannel(p.call("createDataChannel", label, dataChannelDict[0].JSValue()))
	}
}

func (p *rtcPeerConnectionImpl) OnDataChannel(fn func(RTCDataChannelEvent)) EventHandler {
	return p.On("datachannel", func(e Event) {
		if ie, ok := e.(RTCDataChannelEvent); ok {
			fn(ie)
		}
	})
}

// http://w3c.github.io/webrtc-pc/#rtcpeerconnection-interface-extensions-1
func (p *rtcPeerConnectionImpl) Stats(selector ...MediaStreamTrack) func() (RTCStatsReport, error) {
	return func() (RTCStatsReport, error) {
		var (
			result Value
			ok     bool
		)
		switch len(selector) {
		case 0:
			result, ok = await(p.call("getStats"))
		default:
			result, ok = await(p.call("getStats", JSValueOf(selector[0])))
		}

		if ok {
			return wrapRTCStatsReport(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

func (p *rtcPeerConnectionImpl) OnStatsEnded(fn func(RTCStatsEvent)) EventHandler {
	return p.On("statsended", func(e Event) {
		if ie, ok := e.(RTCStatsEvent); ok {
			fn(ie)
		}
	})
}

// -------------8<---------------------------------------

type rtcSessionDescriptionImpl struct {
	Value
}

func wrapRTCSessionDescription(v Value) RTCSessionDescription {
	if v.valid() {
		return &rtcSessionDescriptionImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcSessionDescriptionImpl) Type() RTCSdpType {
	return RTCSdpType(p.get("type").toString())
}

func (p *rtcSessionDescriptionImpl) Sdp() string {
	return p.get("sdp").toString()
}

func (p *rtcSessionDescriptionImpl) ToJSON() string {
	res := p.call("toJSON")
	return jsJSON.call("stringify", res).toString()
}

// -------------8<---------------------------------------

type rtcIceCandidateImpl struct {
	Value
}

func wrapRTCIceCandidate(v Value) RTCIceCandidate {
	if v.valid() {
		return &rtcIceCandidateImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcIceCandidateImpl) Candidate() string {
	return p.get("candidate").toString()
}

func (p *rtcIceCandidateImpl) SdpMid() string {
	return p.get("sdpMid").toString()
}

func (p *rtcIceCandidateImpl) SdpMLineIndex() uint16 {
	return p.get("sdpMLineIndex").toUint16()
}

func (p *rtcIceCandidateImpl) Foundation() string {
	return p.get("foundation").toString()
}

func (p *rtcIceCandidateImpl) Component() RTCIceComponent {
	return RTCIceComponent(p.get("component").toString())
}

func (p *rtcIceCandidateImpl) Priority() uint {
	return p.get("priority").toUint()
}

func (p *rtcIceCandidateImpl) Address() string {
	return p.get("address").toString()
}

func (p *rtcIceCandidateImpl) Protocol() RTCIceProtocol {
	return RTCIceProtocol(p.get("protocol").toString())
}

func (p *rtcIceCandidateImpl) Port() uint16 {
	return p.get("port").toUint16()
}

func (p *rtcIceCandidateImpl) Type() RTCIceCandidateType {
	return RTCIceCandidateType(p.get("type").toString())
}

func (p *rtcIceCandidateImpl) TcpType() RTCIceTcpCandidateType {
	return RTCIceTcpCandidateType(p.get("tcpType").toString())
}

func (p *rtcIceCandidateImpl) RelatedAddress() string {
	return p.get("relatedAddress").toString()
}

func (p *rtcIceCandidateImpl) RelatedPort() uint16 {
	return p.get("relatedPort").toUint16()
}

func (p *rtcIceCandidateImpl) UsernameFragment() string {
	return p.get("usernameFragment").toString()
}

func (p *rtcIceCandidateImpl) ToJSON() RTCIceCandidateInit {
	return wrapRTCIceCandidateInit(p.call("toJSON"))
}

// -------------8<---------------------------------------

type rtcPeerConnectionIceEventImpl struct {
	*eventImpl
}

func wrapRTCPeerConnectionIceEvent(v Value) RTCPeerConnectionIceEvent {
	if v.valid() {
		return &rtcPeerConnectionIceEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcPeerConnectionIceEventImpl) Candidate() RTCIceCandidate {
	return wrapRTCIceCandidate(p.get("candidate"))
}

func (p *rtcPeerConnectionIceEventImpl) URL() string {
	return p.get("url").toString()
}

// -------------8<---------------------------------------

type rtcPeerConnectionIceErrorEventImpl struct {
	*eventImpl
}

func wrapRTCPeerConnectionIceErrorEvent(v Value) RTCPeerConnectionIceErrorEvent {
	if v.valid() {
		return &rtcPeerConnectionIceErrorEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcPeerConnectionIceErrorEventImpl) HostCandidate() string {
	return p.get("hostCandidate").toString()
}

func (p *rtcPeerConnectionIceErrorEventImpl) URL() string {
	return p.get("url").toString()
}

func (p *rtcPeerConnectionIceErrorEventImpl) ErrorCode() uint16 {
	return p.get("errorCode").toUint16()
}

func (p *rtcPeerConnectionIceErrorEventImpl) ErrorText() string {
	return p.get("errorText").toString()
}

// -------------8<---------------------------------------

type rtcCertificateImpl struct {
	Value
}

func wrapRTCCertificate(v Value) RTCCertificate {
	if v.valid() {
		return &rtcCertificateImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcCertificateImpl) Expires() time.Time {
	return domTimeStampToTime(p.get("expires").toUint64())
}

func (p *rtcCertificateImpl) SupportedAlgorithms() []AlgorithmIdentifier { // static
	if s := p.call("getSupportedAlgorithms").toSlice(); s != nil {
		ret := make([]AlgorithmIdentifier, len(s))
		for i, a := range s {
			ret[i] = AlgorithmIdentifier(a.toString()) // TODO this is typedef Object or DOMString
		}
		return ret
	}
	return nil
}

func (p *rtcCertificateImpl) Fingerprints() []RTCDtlsFingerprint {
	if s := p.call(" getFingerprints").toSlice(); s != nil {
		ret := make([]RTCDtlsFingerprint, len(s))
		for i, f := range s {
			ret[i] = wrapRTCDtlsFingerprint(f)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type rtcRtpSenderImpl struct {
	Value
}

func wrapRTCRtpSender(v Value) RTCRtpSender {
	if v.valid() {
		return &rtcRtpSenderImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcRtpSenderImpl) Track() MediaStreamTrack {
	return wrapMediaStreamTrack(p.get("track"))
}

func (p *rtcRtpSenderImpl) Transport() RTCDtlsTransport {
	return wrapRTCDtlsTransport(p.get("transport"))
}

func (p *rtcRtpSenderImpl) RTCPTransport() RTCDtlsTransport {
	return wrapRTCDtlsTransport(p.get("rtcpTransport"))
}

func (p *rtcRtpSenderImpl) Capabilities(kind string) RTCRtpCapabilities { // static
	return wrapRTCRtpCapabilities(p.call("getCapabilities", kind))
}

func (p *rtcRtpSenderImpl) SetParameters(parameters RTCRtpSendParameters) func() error {
	return func() error {
		result, ok := await(p.call("setParameters", parameters.JSValue()))
		if ok {
			return nil
		}
		//TODO: RangeError, RTCError, InvalidStateError, InvalidModificationError, OperationError
		return wrapDOMException(result)
	}
}

func (p *rtcRtpSenderImpl) Parameters() RTCRtpSendParameters {
	return wrapRTCRtpSendParameters(p.call("getParameters"))
}

func (p *rtcRtpSenderImpl) ReplaceTrack(withTrack MediaStreamTrack) func() error {
	return func() error {
		result, ok := await(p.call("replaceTrack", JSValueOf(withTrack)))
		if ok {
			return nil
		}
		return wrapDOMException(result)
	}
}

func (p *rtcRtpSenderImpl) SetStreams(streams ...MediaStream) {
	switch len(streams) {
	case 0:
		p.call("setStreams")
	default: // TODO
		jsSlc := make([]jsValue, len(streams))
		for i, ms := range streams {
			jsSlc[i] = JSValueOf(ms)
		}
		ifaceSlice := ToIfaceSlice(jsSlc)
		p.call("setStreams", ifaceSlice...)
	}
}

func (p *rtcRtpSenderImpl) Stats() func() (RTCStatsReport, error) {
	return func() (RTCStatsReport, error) {
		result, ok := await(p.call("getStats"))
		if ok {
			return wrapRTCStatsReport(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

func (p *rtcRtpSenderImpl) DTMF() RTCDTMFSender {
	return wrapRTCDTMFSender(p.get("dtmf"))
}

// -------------8<---------------------------------------

type rtcRtpReceiverImpl struct {
	Value
}

func wrapRTCRtpReceiver(v Value) RTCRtpReceiver {
	if v.valid() {
		return &rtcRtpReceiverImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcRtpReceiverImpl) Track() MediaStreamTrack {
	return wrapMediaStreamTrack(p.get("track"))
}

func (p *rtcRtpReceiverImpl) Transport() RTCDtlsTransport {
	return wrapRTCDtlsTransport(p.get("transport"))
}

func (p *rtcRtpReceiverImpl) RTCPTransport() RTCDtlsTransport {
	return wrapRTCDtlsTransport(p.get("rtcpTransport"))
}

func (p *rtcRtpReceiverImpl) Capabilities(kind string) RTCRtpCapabilities { // static
	return wrapRTCRtpCapabilities(p.call("getCapabilities"))
}

func (p *rtcRtpReceiverImpl) Parameters() RTCRtpReceiveParameters {
	return wrapRTCRtpReceiveParameters(p.call("getParameters"))
}

func (p *rtcRtpReceiverImpl) ContributingSources() []RTCRtpContributingSource {
	return toRTCRtpContributingSourceSlice(p.call("getContributingSources"))
}

func (p *rtcRtpReceiverImpl) SynchronizationSources() []RTCRtpSynchronizationSource {
	return toRTCRtpSynchronizationSourceSlice(p.call("getSynchronizationSources"))
}

func (p *rtcRtpReceiverImpl) Stats() func() (RTCStatsReport, error) {
	return func() (RTCStatsReport, error) {
		result, ok := await(p.call("getStats"))
		if ok {
			return wrapRTCStatsReport(result), nil
		}
		return nil, wrapDOMException(result)
	}
}

// -------------8<---------------------------------------

type rtcRtpTransceiverImpl struct {
	Value
}

func wrapRTCRtpTransceiver(v Value) RTCRtpTransceiver {
	if v.valid() {
		return &rtcRtpTransceiverImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcRtpTransceiverImpl) Mid() string {
	return p.get("mid").toString()
}

func (p *rtcRtpTransceiverImpl) Sender() RTCRtpSender {
	return wrapRTCRtpSender(p.get("sender"))
}

func (p *rtcRtpTransceiverImpl) Receiver() RTCRtpReceiver {
	return wrapRTCRtpReceiver(p.get("receiver"))
}

func (p *rtcRtpTransceiverImpl) Stopped() bool {
	return p.get("stopped").toBool()
}

func (p *rtcRtpTransceiverImpl) Direction() RTCRtpTransceiverDirection {
	return RTCRtpTransceiverDirection(p.get("direction").toString())
}

func (p *rtcRtpTransceiverImpl) SetDirection(direction RTCRtpTransceiverDirection) {
	p.set("direction", string(direction))
}

func (p *rtcRtpTransceiverImpl) CurrentDirection() RTCRtpTransceiverDirection {
	return RTCRtpTransceiverDirection(p.get("currentDirection").toString())
}

func (p *rtcRtpTransceiverImpl) Stop() {
	p.call("stop")
}

func (p *rtcRtpTransceiverImpl) SetCodecPreferences(codecs []RTCRtpCodecCapability) {
	p.call("setCodecPreferences", ToJSArray(codecs))
}

// -------------8<---------------------------------------

type rtcDtlsTransportImpl struct {
	*eventTargetImpl
}

func wrapRTCDtlsTransport(v Value) RTCDtlsTransport {
	if v.valid() {
		return &rtcDtlsTransportImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcDtlsTransportImpl) IceTransport() RTCIceTransport {
	return wrapRTCIceTransport(p.get("iceTransport"))
}

func (p *rtcDtlsTransportImpl) State() RTCDtlsTransportState {
	return RTCDtlsTransportState(p.get("state").toString())
}

func (p *rtcDtlsTransportImpl) RemoteCertificates() []ArrayBuffer {
	return toArrayBufferSlice(p.call("getRemoteCertificates"))
}

func (p *rtcDtlsTransportImpl) OnStateChange(fn func(Event)) EventHandler {
	return p.On("statechange", fn)
}

func (p *rtcDtlsTransportImpl) OnError(fn func(RTCErrorEvent)) EventHandler {
	return p.On("error", func(e Event) {
		if ee, ok := e.(RTCErrorEvent); ok {
			fn(ee)
		}
	})
}

// -------------8<---------------------------------------

type rtcIceTransportImpl struct {
	*eventTargetImpl
}

func wrapRTCIceTransport(v Value) RTCIceTransport {
	if v.valid() {
		return &rtcIceTransportImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcIceTransportImpl) Role() RTCIceRole {
	return RTCIceRole(p.get("role").toString())
}

func (p *rtcIceTransportImpl) Component() RTCIceComponent {
	return RTCIceComponent(p.get("component").toString())
}

func (p *rtcIceTransportImpl) State() RTCIceTransportState {
	return RTCIceTransportState(p.get("state").toString())
}

func (p *rtcIceTransportImpl) GatheringState() RTCIceGathererState {
	return RTCIceGathererState(p.get("gatheringState").toString())
}

func (p *rtcIceTransportImpl) LocalCandidates() []RTCIceCandidate {
	if slc := p.call("getLocalCandidates").toSlice(); slc != nil {
		ret := make([]RTCIceCandidate, len(slc))
		for i, c := range slc {
			ret[i] = wrapRTCIceCandidate(c)
		}
		return ret
	}
	return nil
}

func (p *rtcIceTransportImpl) RemoteCandidates() []RTCIceCandidate {
	if slc := p.call("getRemoteCandidates").toSlice(); slc != nil {
		ret := make([]RTCIceCandidate, len(slc))
		for i, c := range slc {
			ret[i] = wrapRTCIceCandidate(c)
		}
		return ret
	}
	return nil
}

func (p *rtcIceTransportImpl) SelectedCandidatePair() RTCIceCandidatePair {
	return wrapRTCIceCandidatePair(p.call("getSelectedCandidatePair"))
}

func (p *rtcIceTransportImpl) LocalParameters() RTCIceParameters {
	return wrapRTCIceParameters(p.call("getLocalParameters"))
}

func (p *rtcIceTransportImpl) RemoteParameters() RTCIceParameters {
	return wrapRTCIceParameters(p.call("getRemoteParameters"))
}

func (p *rtcIceTransportImpl) OnStateChange(fn func(Event)) EventHandler {
	return p.On("statechange", fn)
}

func (p *rtcIceTransportImpl) OnGatheringStateChange(fn func(Event)) EventHandler {
	return p.On("gatheringstatechange", fn)
}

func (p *rtcIceTransportImpl) OnSelectedCandidatePairChange(fn func(Event)) EventHandler {
	return p.On("selectedcandidatepairchange", fn)
}

// -------------8<---------------------------------------

type rtcTrackEventImpl struct {
	*eventImpl
}

func wrapRTCTrackEvent(v Value) RTCTrackEvent {
	if v.valid() {
		return &rtcTrackEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcTrackEventImpl) Receiver() RTCRtpReceiver {
	return wrapRTCRtpReceiver(p.get("receiver"))
}

func (p *rtcTrackEventImpl) Track() MediaStreamTrack {
	return wrapMediaStreamTrack(p.get("track"))
}

func (p *rtcTrackEventImpl) Streams() []MediaStream {
	if slc := p.get("streams").toSlice(); slc != nil {
		ret := make([]MediaStream, len(slc))
		for i, s := range slc {
			ret[i] = wrapMediaStream(s)
		}
		return ret
	}
	return nil
}

func (p *rtcTrackEventImpl) Transceiver() RTCRtpTransceiver {
	return wrapRTCRtpTransceiver(p.get("transceiver"))
}

// -------------8<---------------------------------------

type rtcSctpTransportImpl struct {
	*eventTargetImpl
}

func wrapRTCSctpTransport(v Value) RTCSctpTransport {
	if v.valid() {
		return &rtcSctpTransportImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcSctpTransportImpl) Transport() RTCDtlsTransport {
	return wrapRTCDtlsTransport(p.get("transport"))
}

func (p *rtcSctpTransportImpl) State() RTCSctpTransportState {
	return RTCSctpTransportState(p.get("state").toString())
}

func (p *rtcSctpTransportImpl) MaxMessageSize() float64 {
	return p.get("maxMessageSize").toFloat64()
}

func (p *rtcSctpTransportImpl) MaxChannels() uint16 {
	return p.get("maxChannels").toUint16()
}

func (p *rtcSctpTransportImpl) OnStateChange(fn func(Event)) EventHandler {
	return p.On("statechange", fn)
}

// -------------8<---------------------------------------

type rtcDataChannelImpl struct {
	*eventTargetImpl
}

func wrapRTCDataChannel(v Value) RTCDataChannel {
	if v.valid() {
		return &rtcDataChannelImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcDataChannelImpl) Label() string {
	return p.get("label").toString()
}

func (p *rtcDataChannelImpl) Ordered() bool {
	return p.get("ordered").toBool()
}

func (p *rtcDataChannelImpl) MaxPacketLifeTime() uint16 {
	return p.get("maxPacketLifeTime").toUint16()
}

func (p *rtcDataChannelImpl) MaxRetransmits() uint16 {
	return p.get("maxRetransmits").toUint16()
}

func (p *rtcDataChannelImpl) Protocol() string {
	return p.get("protocol").toString()
}

func (p *rtcDataChannelImpl) Negotiated() bool {
	return p.get("negotiated").toBool()
}

func (p *rtcDataChannelImpl) Id() uint16 {
	return p.get("id").toUint16()
}

func (p *rtcDataChannelImpl) Priority() RTCPriorityType {
	return RTCPriorityType(p.get("priority").toString())
}

func (p *rtcDataChannelImpl) ReadyState() RTCDataChannelState {
	return RTCDataChannelState(p.get("readyState").toString())
}

func (p *rtcDataChannelImpl) BufferedAmount() uint {
	return p.get("bufferedAmount").toUint()
}

func (p *rtcDataChannelImpl) BufferedAmountLowThreshold() uint {
	return p.get("bufferedAmountLowThreshold").toUint()
}

func (p *rtcDataChannelImpl) OnOpen(fn func(Event)) EventHandler {
	return p.On("open", fn)
}

func (p *rtcDataChannelImpl) OnBufferedAmountLow(fn func(Event)) EventHandler {
	return p.On("bufferedamountlow", fn)
}

func (p *rtcDataChannelImpl) OnError(fn func(RTCErrorEvent)) EventHandler {
	return p.On("error", func(e Event) {
		if ee, ok := e.(RTCErrorEvent); ok {
			fn(ee)
		}
	})
}

func (p *rtcDataChannelImpl) OnClose(fn func(Event)) EventHandler {
	return p.On("close", fn)
}

func (p *rtcDataChannelImpl) Close() {
	p.call("close")
}

func (p *rtcDataChannelImpl) OnMessage(fn func(MessageEvent)) EventHandler {
	return p.On("message", func(e Event) {
		if me, ok := e.(MessageEvent); ok {
			fn(me)
		}
	})
}

func (p *rtcDataChannelImpl) BinaryType() string {
	return p.get("binaryType").toString()
}

func (p *rtcDataChannelImpl) Send(typ interface{}) { // string, Blob, ArrayBuffer, ArrayBufferView
	switch x := typ.(type) {
	case string:
		p.call("send", x)
	case []byte:
		ta := jsTypedArrayOf(x)
		blob := NewBlob(ta)
		p.call("send", JSValueOf(blob))
		ta.Release()
	case Blob, ArrayBuffer, ArrayBufferView:
		p.call("send", JSValueOf(x))
	}
}

// -------------8<---------------------------------------

type rtcDataChannelEventImpl struct {
	*eventImpl
}

func wrapRTCDataChannelEvent(v Value) RTCDataChannelEvent {
	if v.valid() {
		return &rtcDataChannelEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcDataChannelEventImpl) Channel() RTCDataChannel {
	return wrapRTCDataChannel(p.get("channel"))
}

// -------------8<---------------------------------------

type rtcDTMFSenderImpl struct {
	*eventTargetImpl
}

func wrapRTCDTMFSender(v Value) RTCDTMFSender {
	if v.valid() {
		return &rtcDTMFSenderImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *rtcDTMFSenderImpl) InsertDTMF(tones string, args ...uint) {
	switch len(args) {
	case 0:
		p.call("insertDTMF", tones)
	case 1:
		p.call("insertDTMF", tones, args[0]) // duration
	case 2:
		p.call("insertDTMF", tones, args[0], args[1]) // duration, interToneGap
	}
}

func (p *rtcDTMFSenderImpl) OnToneChange(fn func(RTCDTMFToneChangeEvent)) EventHandler {
	return p.On("tonechange", func(e Event) {
		if ce, ok := e.(RTCDTMFToneChangeEvent); ok {
			fn(ce)
		}
	})
}

func (p *rtcDTMFSenderImpl) CanInsertDTMF() bool {
	return p.get("canInsertDTMF").toBool()
}

func (p *rtcDTMFSenderImpl) ToneBuffer() string {
	return p.get("toneBuffer").toString()
}

// -------------8<---------------------------------------

type rtcDTMFToneChangeEventImpl struct {
	*eventImpl
}

func wrapRTCDTMFToneChangeEvent(v Value) RTCDTMFToneChangeEvent {
	if v.valid() {
		return &rtcDTMFToneChangeEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcDTMFToneChangeEventImpl) Tone() string {
	return p.get("tone").toString()
}

// -------------8<---------------------------------------

type rtcStatsReportImpl struct {
	Value
}

func wrapRTCStatsReport(v Value) RTCStatsReport {
	// TODO
	// https://www.w3.org/TR/webrtc-stats/
	return nil
}

// -------------8<---------------------------------------

type rtcStatsEventImpl struct {
	*eventImpl
}

func wrapRTCStatsEvent(v Value) RTCStatsEvent {
	if v.valid() {
		return &rtcStatsEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcStatsEventImpl) Report() RTCStatsReport {
	return wrapRTCStatsEvent(p.get("report"))
}

// -------------8<---------------------------------------

type rtcErrorImpl struct {
	Value
}

func wrapRTCError(v Value) RTCError {
	if v.valid() {
		return &rtcErrorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *rtcErrorImpl) ErrorDetail() RTCErrorDetailType {
	return RTCErrorDetailType(p.get("errorDetail").toString())
}

func (p *rtcErrorImpl) SDPLineNumber() int {
	return p.get("sdpLineNumber").toInt()
}

func (p *rtcErrorImpl) HttpRequestStatusCode() int {
	return p.get("httpRequestStatusCode").toInt()
}

func (p *rtcErrorImpl) SCTPCauseCode() int {
	return p.get("sctpCauseCode").toInt()
}

func (p *rtcErrorImpl) ReceivedAlert() uint {
	return p.get("receivedAlert").toUint()
}

func (p *rtcErrorImpl) SentAlert() uint {
	return p.get("sentAlert").toUint()
}

func (p *rtcErrorImpl) Message() string {
	return p.get("message").toString()
}

func (p *rtcErrorImpl) Name() string {
	return p.get("name").toString()
}

// -------------8<---------------------------------------

type rtcErrorEventImpl struct {
	*eventImpl
}

func wrapRTCErrorEvent(v Value) RTCErrorEvent {
	if v.valid() {
		return &rtcErrorEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *rtcErrorEventImpl) Error() RTCError {
	return wrapRTCError(p.get("error"))
}

// -------------8<---------------------------------------
