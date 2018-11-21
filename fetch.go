// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://fetch.spec.whatwg.org/#idl-index

type (

	// typedef (sequence<sequence<ByteString>> or record<ByteString, ByteString>) HeadersInit;
	HeadersInit interface{}

	Headers interface {
		js.Wrapper

		Append(string, string)
		Delete(string)
		Get(string) string
		Has(string) bool
		Set(string, string)
		Entries() map[string]string
	}

	// typedef (Blob or BufferSource or FormData or URLSearchParams or ReadableStream or USVString) BodyInit;
	BodyInit interface{}

	Body interface {
		Body() ReadableStream
		BodyUsed() bool
		ArrayBuffer() Promise
		Blob() Promise
		FormData() Promise
		JSON() Promise
		Text() Promise
	}

	// typedef (Request or USVString) RequestInfo;
	RequestInfo interface{}

	Request interface {
		Body

		Method() string
		URL() string
		Headers() Headers
		Destination() RequestDestination
		Referrer() string
		ReferrerPolicy() ReferrerPolicy
		Mode() RequestMode
		Credentials() RequestCredentials
		Cache() RequestCache
		Redirect() RequestRedirect
		Integrity() string
		Keepalive() bool
		IsReloadNavigation() bool
		IsHistoryNavigation() bool
		Signal() AbortSignal
		Clone() Request
	}

	RequestInit struct {
		Method         string             `json:"method"`
		Headers        HeadersInit        `json:"headers"`
		Body           BodyInit           `json:"body"`
		Referrer       string             `json:"referrer"`
		ReferrerPolicy ReferrerPolicy     `json:"referrerPolicy"`
		Mode           RequestMode        `json:"mode"`
		Credentials    RequestCredentials `json:"credentials"`
		Cache          RequestCache       `json:"cache"`
		Redirect       RequestRedirect    `json:"redirect"`
		Integrity      string             `json:"integrity"`
		Keepalive      bool               `json:"keepalive"`
		Signal         AbortSignal        `json:"signal"`
		Window         js.Value           `json:"window"`
	}

	Response interface {
		Body

		Error() Response
		Redirect(string, ...int) Response
		Type() ResponseType
		URL() string
		Redirected() bool
		Status() int
		Ok() bool
		StatusText() string
		Headers() Headers
		Trailer() Promise
		Clone() Response
	}

	ResponseInit struct {
		status     int         `json:"status"`
		StatusText string      `json:"statusText"`
		Headers    HeadersInit `json:"headers"`
	}
)

type RequestDestination string

const (
	RequestDestinationEmpty        RequestDestination = ""
	RequestDestinationAudio        RequestDestination = "audio"
	RequestDestinationAudioWorklet RequestDestination = "audioworklet"
	RequestDestinationDocument     RequestDestination = "document"
	RequestDestinationEmbed        RequestDestination = "embed"
	RequestDestinationFont         RequestDestination = "font"
	RequestDestinationImage        RequestDestination = "image"
	RequestDestinationManifest     RequestDestination = "manifest"
	RequestDestinationObject       RequestDestination = "object"
	RequestDestinationPaintWorklet RequestDestination = "paintworklet"
	RequestDestinationReport       RequestDestination = "report"
	RequestDestinationScript       RequestDestination = "script"
	RequestDestinationSharedWorker RequestDestination = "sharedworker"
	RequestDestinationStyle        RequestDestination = "style"
	RequestDestinationTrack        RequestDestination = "track"
	RequestDestinationVideo        RequestDestination = "video"
	RequestDestinationWorker       RequestDestination = "worker"
	RequestDestinationXSLT         RequestDestination = "xslt"
)

type RequestMode string

const (
	RequestModeNavigate   RequestMode = "navigate"
	RequestModeSameOrigin RequestMode = "same-origin"
	RequestModeNoCors     RequestMode = "no-cors"
	RequestModeCors       RequestMode = "cors"
)

type RequestCredentials string

const (
	RequestCredentialsOmit       RequestCredentials = "omit"
	RequestCredentialsSameOrigin RequestCredentials = "same-origin"
	RequestCredentialsInclude    RequestCredentials = "include"
)

type RequestCache string

const (
	RequestCacheDefault      RequestCache = "default"
	RequestCacheNoStore      RequestCache = "no-store"
	RequestCacheReload       RequestCache = "reload"
	RequestCacheNoCache      RequestCache = "no-cache"
	RequestCacheForceCache   RequestCache = "force-cache"
	RequestCacheOnlyIfCached RequestCache = "only-if-cached"
)

type RequestRedirect string

const (
	RequestRedirectFollow RequestRedirect = "follow"
	RequestRedirectError  RequestRedirect = "error"
	RequestRedirectManual RequestRedirect = "manual"
)

type ResponseType string

const (
	ResponseTypeBasic          ResponseType = "basic"
	ResponseTypeCors           ResponseType = "cors"
	ResponseTypeDefault        ResponseType = "default"
	ResponseTypeError          ResponseType = "error"
	ResponseTypeOpaque         ResponseType = "opaque"
	ResponseTypeOpaqueRedirect ResponseType = "opaqueredirect"
)

type ReferrerPolicy string

const (
	ReferrerPolicyEmpty                       ReferrerPolicy = ""
	ReferrerPolicyNoReferrer                  ReferrerPolicy = "no-referrer"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicy = "no-referrer-when-downgrade"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "same-origin"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strict-origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "origin-when-cross-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"
	ReferrerPolicyUnsafeUrl                   ReferrerPolicy = "unsafe-url"
)
