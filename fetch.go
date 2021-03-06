// +build js,wasm

package wasm

// https://fetch.spec.whatwg.org/#idl-index

type (
	Headers interface {
		Append(string, string)
		Delete(string)
		Get(string) string
		Has(string) bool
		Set(string, string)
		Entries() map[string]string
	}

	// https://fetch.spec.whatwg.org/#body
	Body interface {
		Body() ReadableStream
		BodyUsed() bool
		ArrayBuffer() func() (ArrayBuffer, bool) // Promise <ArrayBuffer>
		Blob() func() (Blob, bool)               // Promise<Blob>
		FormData() func() (FormData, bool)       // Promise<FormData>
		JSON() func() ([]byte, bool)             // Promise<any>
		Text() func() (string, bool)             // Promise<string>
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
		Trailer() func() (Headers, bool)
		Clone() Response
	}

	// https://xhr.spec.whatwg.org/#formdataentryvalue
	// typedef (File or USVString) FormDataEntryValue;
	FormDataEntryValue interface{}

	// https://xhr.spec.whatwg.org/#formdata
	FormData interface {
		Append(string, interface{}, ...string)
		Delete(string)
		Get(string) FormDataEntryValue
		GetAll(string) []FormDataEntryValue
		Has(string) bool
		Set(string, interface{}, ...string)
		Values() []FormDataEntryValue

		JSValue() jsValue
	}
)

// https://fetch.spec.whatwg.org/#typedefdef-headersinit
// typedef (sequence<sequence<ByteString>> or record<ByteString, ByteString>) HeadersInit;
// TODO
type HeadersInit interface{}

func headersInitJSValue(p HeadersInit) jsValue {
	switch x := p.(type) {
	case nil:
		return jsNull
	case []string:
		return ToJSArray(x)
	case map[string]string:
		o := jsObject.New()
		for k, v := range x {
			o.Set(k, v)
		}
		return o
	case string: // TODO ???
		return JSValueOf(x)
	default:
		return jsUndefined
	}
}

// typedef (Blob or BufferSource or FormData or URLSearchParams or ReadableStream or USVString) BodyInit;
type BodyInit interface{}

func bodyInitJSValue(p BodyInit) jsValue {
	switch x := p.(type) {
	case nil:
		return jsNull
	case Blob, BufferSource, FormData, URLSearchParams, ReadableStream, string:
		return JSValueOf(x)
	default:
		return jsUndefined
	}
}

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

// -------------8<---------------------------------------

// https://fetch.spec.whatwg.org/#requestinit
type RequestInit struct {
	Method         string
	Headers        HeadersInit
	Body           BodyInit
	Referrer       string
	ReferrerPolicy ReferrerPolicy
	Mode           RequestMode
	Credentials    RequestCredentials
	Cache          RequestCache
	Redirect       RequestRedirect
	Integrity      string
	Keepalive      bool
	Signal         AbortSignal
	// Window         interface{}
}

func (p RequestInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("method", p.Method)
	o.Set("headers", headersInitJSValue(p.Headers))
	o.Set("body", bodyInitJSValue(p.Body))
	o.Set("referrer", p.Referrer)
	o.Set("referrerPolicy", string(p.ReferrerPolicy))
	o.Set("mode", string(p.Mode))
	o.Set("credentials", string(p.Credentials))
	o.Set("cache", string(p.Cache))
	o.Set("redirect", string(p.Redirect))
	o.Set("integrity", p.Integrity)
	o.Set("keepalive", p.Keepalive)
	o.Set("signal", JSValueOf(p.Signal))
	// o.set("window", p.Window)
	return o
}

// -------------8<---------------------------------------

type ResponseInit struct {
	Status     int
	StatusText string
	Headers    HeadersInit
}

func (p ResponseInit) JSValue() jsValue {
	o := jsObject.New()
	o.Set("status", p.Status)
	o.Set("statusText", p.StatusText)
	o.Set("headers", headersInitJSValue(p.Headers))
	return o
}
