// +build js,wasm

package wasm

import (
	"errors"
	"time"
)

// -------------8<---------------------------------------

var (
	errUnsupporttedType = errors.New("Unsupported Type")
	errInvalidType      = errors.New("Invalid Type")
	jsDate              = jsGlobal.get("Date")
	jsWindowProxy       = jsGlobal
	jsMessagePort       = jsGlobal.get("MessagePort")
	jsUint8Array        = jsGlobal.get("Uint8Array")
	jsJSON              = jsGlobal.get("JSON")
)

// -------------8<---------------------------------------

func nodeListToSlice(v Value) []Node {
	if v.valid() && v.length() > 0 {
		ret := make([]Node, v.length())
		for i := range ret {
			ret[i] = wrapAsNode(v.index(i))
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func elementArrayToSlice(v Value) []Element {
	ret := make([]Element, v.length())
	for i := range ret {
		ret[i] = wrapAsElement(v.index(i))
	}
	return ret
}

// -------------8<---------------------------------------

func domQuadArrayToSlice(v Value) []DOMQuad {
	if v.valid() && v.length() > 0 {
		ret := make([]DOMQuad, v.length())
		for i := range ret {
			ret[i] = wrapDOMQuad(v.index(i))
		}
		return ret
	}

	return nil
}

// -------------8<---------------------------------------

func stringSequenceToSlice(s Value) []string {
	if s.valid() && s.length() > 0 {
		ret := make([]string, s.length())
		for i := range ret {
			ret[i] = s.index(i).toString()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func boolSequenceToSlice(s Value) []bool {
	var slc []bool
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return v.toBool()
	})
	return slc
}

/*
func boolSequenceToSlice(s Value) []bool {
	if s.valid() && s.length() > 0 {
		ret := make([]bool, s.length())
		for i := range ret {
			ret[i] = s.index(i).toBool()
		}
		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func floatSequenceToSlice(s Value) []float64 {
	var slc []float64
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return v.toFloat64()
	})
	return slc
}

/*
func floatSequenceToSlice(s Value) []float64 {
	if s.valid() && s.length() > 0 {
		ret := make([]float64, s.length())
		for i := range ret {
			ret[i] = s.index(i).toFloat64()
		}
		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func mutationRecordSequenceToSlice(s Value) []MutationRecord {
	var slc []MutationRecord
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return wrapMutationRecord(v)
	})
	return slc
}

/*
func mutationRecordSequenceToSlice(v Value) []MutationRecord {
	if v.valid() && v.length() > 0 {
		ret := make([]MutationRecord, v.length())
		for i := range ret {
			ret[i] = wrapMutationRecord(v.index(i))
		}
		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func fileListToSlice(s Value) []File {
	var slc []File
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return wrapFile(v)
	})
	return slc
}

/*
func fileListToSlice(v Value) []File {
	ret := make([]File, v.length())
	for i := range ret {
		ret[i] = wrapFile(v.index(i))
	}
	return ret
}
*/

// -------------8<---------------------------------------

func keys(v Value) []string {
	if v.valid() {
		a := Value{jsObject.Call("keys", v)}
		ret := make([]string, a.length())
		for i := range ret {
			ret[i] = a.index(i).toString()
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func domStringMapToMap(v Value) map[string]string {
	m := make(map[string]string)

	for _, key := range keys(v) {
		m[key] = v.get(key).toString()
	}

	return m
}

// -------------8<---------------------------------------

// expects v is js Date object
func jsDateToTime(v Value) time.Time {
	ms := int64(v.call("getTime").toFloat64()) * int64(time.Millisecond)
	return time.Unix(0, ms)
}

// -------------8<---------------------------------------
// https://heycam.github.io/webidl/#DOMTimeStamp

func domTimeStampToTime(ts uint64) time.Time {
	return time.Unix(0, int64(ts)*int64(time.Millisecond))
}

// -------------8<---------------------------------------

func domStringListToSlice(dsl Value) []string {
	var slc []string
	dsl.AppendToSlice(&slc, func(v Value) interface{} {
		return v.call("item").toString()
	})
	return slc
}

/*
func domStringListToSlice(dsl Value) []string {
	if dsl.valid() && dsl.length() > 0 {
		ret := make([]string, dsl.length())
		for i := range ret {
			ret[i] = dsl.call("item").toString()
		}

		return ret
	}
	return nil
}
*/
// -------------8<---------------------------------------

func touchListToSlice(s Value) []Touch {
	var slc []Touch
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return wrapTouch(v)
	})
	return slc
}

/*
func touchListToSlice(v Value) []Touch {
	if v.valid() && v.length() > 0 {
		ret := make([]Touch, v.length())
		for i := range ret {
			ret[i] = wrapTouch(v.index(i))
		}
		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func toFloat32Slice(s Value) []float32 {
	var slc []float32
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return v.toFloat32()
	})
	return slc
}

/*
func toFloat32Slice(v Value) []float32 {
	if v.valid() && v.length() > 0 {
		ret := make([]float32, v.length())
		for i := range ret {
			ret[i] = v.index(i).toFloat32()
		}
		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func toFloat64Slice(s Value) []float64 {
	var slc []float64
	s.AppendToSlice(&slc, func(v Value) interface{} {
		return v.toFloat64()
	})
	return slc
}

/*
func toFloat64Slice(v Value) []float64 {
	if v.valid() && v.length() > 0 {
		ret := make([]float64, v.length())
		for i := range ret {
			ret[i] = v.index(i).toFloat64()
		}

		return ret
	}
	return nil
}
*/

// -------------8<---------------------------------------

func htmlCollectionToElementSlice(v Value) []Element {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		ret := make([]Element, c.Length())
		for i := uint(0); i < c.Length(); i++ {
			ret[i] = c.Item(i)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToHTMLElementSlice(v Value) []HTMLElement {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		var ret []HTMLElement
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func htmlCollectionToHTMLOptionElementSlice(v Value) []HTMLOptionElement {
	if c := wrapHTMLCollection(v); c != nil && c.Length() > 0 {
		var ret []HTMLOptionElement
		for i := uint(0); i < c.Length(); i++ {
			if el, ok := c.Item(i).(HTMLOptionElement); ok {
				ret = append(ret, el)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCIceServerSlice(v Value) []RTCIceServer {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCIceServer, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCIceServer(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCCertificateSlice(v Value) []RTCCertificate {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCCertificate, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCCertificate(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpSenderSlice(v Value) []RTCRtpSender {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpSender, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCRtpSender(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpReceiverSlice(v Value) []RTCRtpReceiver {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpReceiver, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCRtpReceiver(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpTransceiverSlice(v Value) []RTCRtpTransceiver {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpTransceiver, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCRtpTransceiver(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpCodecCapabilitySlice(v Value) []RTCRtpCodecCapability {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpCodecCapability, len(slc))
		for i, c := range slc {
			ret[i] = wrapRTCRtpCodecCapability(c)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpHeaderExtensionCapabilitySlice(v Value) []RTCRtpHeaderExtensionCapability {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpHeaderExtensionCapability, len(slc))
		for i, c := range slc {
			ret[i] = wrapRTCRtpHeaderExtensionCapability(c)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpHeaderExtensionParametersSlice(v Value) []RTCRtpHeaderExtensionParameters {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpHeaderExtensionParameters, len(slc))
		for i, p := range slc {
			ret[i] = wrapRTCRtpHeaderExtensionParameters(p)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpCodecParametersSlice(v Value) []RTCRtpCodecParameters {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpCodecParameters, len(slc))
		for i, p := range slc {
			ret[i] = wrapRTCRtpCodecParameters(p)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpContributingSourceSlice(v Value) []RTCRtpContributingSource {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpContributingSource, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCRtpContributingSource(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpDecodingParametersSlice(v Value) []RTCRtpDecodingParameters {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpDecodingParameters, len(slc))
		for i, p := range slc {
			ret[i] = wrapRTCRtpDecodingParameters(p)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toRTCRtpSynchronizationSourceSlice(v Value) []RTCRtpSynchronizationSource {
	if slc := v.toSlice(); slc != nil {
		ret := make([]RTCRtpSynchronizationSource, len(slc))
		for i, s := range slc {
			ret[i] = wrapRTCRtpSynchronizationSource(s)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

func toArrayBufferSlice(v Value) []ArrayBuffer {
	if slc := v.toSlice(); slc != nil {
		ret := make([]ArrayBuffer, len(slc))
		for i, a := range slc {
			ret[i] = wrapArrayBuffer(a)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type DOMError uint16

const (
	ErrIndexSize             DOMError = 1
	ErrDOMStringSize         DOMError = 2
	ErrHierarchyRequest      DOMError = 3
	ErrWrongDocument         DOMError = 4
	ErrInvalidCharacter      DOMError = 5
	ErrNoDataAllowed         DOMError = 6
	ErrNoModificationAllowed DOMError = 7
	ErrNotFound              DOMError = 8
	ErrNotSupported          DOMError = 9
	ErrInuseAttribute        DOMError = 10
	ErrInvalidState          DOMError = 11
	ErrSyntax                DOMError = 12
	ErrInvalidModification   DOMError = 13
	ErrNamespace             DOMError = 14
	ErrInvalidAccess         DOMError = 15
	ErrValidation            DOMError = 16
	ErrTypeMismatch          DOMError = 17
	ErrSecurity              DOMError = 18
	ErrNetwork               DOMError = 19
	ErrAbort                 DOMError = 20
	ErrURLMismatch           DOMError = 21
	ErrQuotaExceeded         DOMError = 22
	ErrTimeout               DOMError = 23
	ErrInvalidNodeType       DOMError = 24
	ErrDataClone             DOMError = 25
)

// -------------8<---------------------------------------
