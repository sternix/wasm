// +build js,wasm

package wasm

// -------------8<---------------------------------------

type headersImpl struct {
	Value
}

func wrapHeaders(v Value) Headers {
	if v.Valid() {
		return &headersImpl{
			Value: v,
		}
	}
	return nil
}

func (p *headersImpl) Append(name string, value string) {
	p.Call("append", name, value)
}
func (p *headersImpl) Delete(name string) {
	p.Call("delete", name)
}

func (p *headersImpl) Get(name string) string {
	return p.Call("get", name).String()
}

func (p *headersImpl) Has(name string) bool {
	return p.Call("has", name).Bool()
}

func (p *headersImpl) Set(name string, value string) {
	p.Call("set", name, value)
}

func (p *headersImpl) Entries() map[string]string {
	ret := make(map[string]string)

	it := p.Call("entries")
	for {
		n := it.Call("next")
		if n.Get("done").Bool() {
			break
		}

		pair := n.Get("value")

		key := pair.Index(0).String()
		value := pair.Index(1).String()

		ret[key] = value
	}
	return ret
}

// -------------8<---------------------------------------

type bodyImpl struct {
	Value
}

func wrapBody(v Value) Body {
	if p := newBodyImpl(v); p != nil {
		return p
	}
	return nil
}

func newBodyImpl(v Value) *bodyImpl {
	if v.Valid() {
		return &bodyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *bodyImpl) Body() ReadableStream {
	// TODO
	return nil
}

func (p *bodyImpl) BodyUsed() bool {
	return p.Get("bodyUsed").Bool()
}

func (p *bodyImpl) ArrayBuffer() func() (ArrayBuffer, bool) {
	return func() (ArrayBuffer, bool) {
		res, ok := await(p.Call("arrayBuffer"))
		if ok {
			return wrapArrayBuffer(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) Blob() func() (Blob, bool) {
	return func() (Blob, bool) {
		res, ok := await(p.Call("blob"))
		if ok {
			return wrapBlob(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) FormData() func() (FormData, bool) {
	return func() (FormData, bool) {
		res, ok := await(p.Call("formData"))
		if ok {
			return wrapFormData(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) JSON() func() ([]byte, bool) {
	return func() ([]byte, bool) {
		res, ok := await(p.Call("json"))
		if ok {
			return []byte(jsJSON.Call("stringify", res).String()), true
		}
		return nil, false
	}
}

func (p *bodyImpl) Text() func() (string, bool) {
	return func() (string, bool) {
		res, ok := await(p.Call("text"))
		if ok {
			return res.String(), true
		}
		return "", false
	}
}

// -------------8<---------------------------------------

type requestImpl struct {
	*bodyImpl
}

func wrapRequest(v Value) Request {
	if v.Valid() {
		return &requestImpl{
			bodyImpl: newBodyImpl(v),
		}
	}
	return nil
}

func (p *requestImpl) Method() string {
	return p.Get("method").String()
}

func (p *requestImpl) URL() string {
	return p.Get("url").String()
}

func (p *requestImpl) Headers() Headers {
	return wrapHeaders(p.Get("headers"))
}

func (p *requestImpl) Destination() RequestDestination {
	return RequestDestination(p.Get("destination").String())
}

func (p *requestImpl) Referrer() string {
	return p.Get("referrer").String()
}

func (p *requestImpl) ReferrerPolicy() ReferrerPolicy {
	return ReferrerPolicy(p.Get("referrerPolicy").String())
}

func (p *requestImpl) Mode() RequestMode {
	return RequestMode(p.Get("mode").String())
}

func (p *requestImpl) Credentials() RequestCredentials {
	return RequestCredentials(p.Get("credentials").String())
}

func (p *requestImpl) Cache() RequestCache {
	return RequestCache(p.Get("cache").String())
}

func (p *requestImpl) Redirect() RequestRedirect {
	return RequestRedirect(p.Get("redirect").String())
}

func (p *requestImpl) Integrity() string {
	return p.Get("integrity").String()
}

func (p *requestImpl) Keepalive() bool {
	return p.Get("keepalive").Bool()
}

func (p *requestImpl) IsReloadNavigation() bool {
	return p.Get("isReloadNavigation").Bool()
}

func (p *requestImpl) IsHistoryNavigation() bool {
	return p.Get("isHistoryNavigation").Bool()
}

func (p *requestImpl) Signal() AbortSignal {
	return wrapAbortSignal(p.Get("signal"))
}

func (p *requestImpl) Clone() Request {
	return wrapRequest(p.Call("clone"))
}

// -------------8<---------------------------------------

type responseImpl struct {
	*bodyImpl
}

func wrapResponse(v Value) Response {
	if v.Valid() {
		return &responseImpl{
			bodyImpl: newBodyImpl(v),
		}
	}
	return nil
}

func (p *responseImpl) Error() Response {
	return wrapResponse(p.Call("error"))
}

func (p *responseImpl) Redirect(url string, status ...int) Response {
	if len(status) > 0 {
		return wrapResponse(p.Call("redirect", url, status[0]))
	}

	return wrapResponse(p.Call("redirect", url))
}

func (p *responseImpl) Type() ResponseType {
	return ResponseType(p.Get("type").String())
}

func (p *responseImpl) URL() string {
	return p.Get("url").String()
}

func (p *responseImpl) Redirected() bool {
	return p.Get("redirected").Bool()
}

func (p *responseImpl) Status() int {
	return p.Get("status").Int()
}

func (p *responseImpl) Ok() bool {
	return p.Get("ok").Bool()
}

func (p *responseImpl) StatusText() string {
	return p.Get("statusText").String()
}

func (p *responseImpl) Headers() Headers {
	return wrapHeaders(p.Get("headers"))
}

func (p *responseImpl) Trailer() func() (Headers, bool) {
	return func() (Headers, bool) {
		res, ok := await(p.Call("trailer"))
		if ok {
			return wrapHeaders(res), true
		}
		return nil, false
	}
}

func (p *responseImpl) Clone() Response {
	return wrapResponse(p.Call("clone"))
}

// -------------8<---------------------------------------

type formDataImpl struct {
	Value
}

func NewFormData(form ...HTMLFormElement) FormData {
	if jsFormData := jsGlobal.Get("FormData"); jsFormData.Valid() {
		switch len(form) {
		case 0:
			return wrapFormData(jsFormData.New())
		default:
			return wrapFormData(jsFormData.New(JSValue(form[0])))
		}
	}
	return nil
}

func wrapFormData(v Value) FormData {
	if v.Valid() {
		return &formDataImpl{
			Value: v,
		}
	}
	return nil
}

func (p *formDataImpl) Append(name string, value interface{}, filename ...string) {
	switch x := value.(type) {
	case string:
		p.Call("append", x)
	case Blob:
		switch len(filename) {
		case 0:
			p.Call("append", JSValue(x))
		default:
			p.Call("append", JSValue(x), filename[0])
		}
	}
}

func (p *formDataImpl) Delete(name string) {
	p.Call("delete", name)
}

func (p *formDataImpl) Get(name string) FormDataEntryValue {
	return wrapFormDataEntryValue(p.Call("get", name))
}

func (p *formDataImpl) GetAll(name string) []FormDataEntryValue {
	if slc := p.Call("getAll", name).ToSlice(); slc != nil {
		var ret []FormDataEntryValue
		for _, v := range slc {
			if fd := wrapFormDataEntryValue(v); fd != nil {
				ret = append(ret, fd)
			}
		}
		return ret
	}
	return nil
}

func (p *formDataImpl) Has(name string) bool {
	return p.Call("has", name).Bool()
}

func (p *formDataImpl) Set(name string, value interface{}, filename ...string) {
	switch x := value.(type) {
	case string:
		p.Call("set", x)
	case Blob:
		switch len(filename) {
		case 0:
			p.Call("set", JSValue(x))
		default:
			p.Call("set", JSValue(x), filename[0])
		}
	}
}

func (p *formDataImpl) Values() []FormDataEntryValue {
	if slc := p.Call("values").ToSlice(); slc != nil {
		var ret []FormDataEntryValue
		for _, v := range slc {
			if fd := wrapFormDataEntryValue(v); fd != nil {
				ret = append(ret, fd)
			}
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

var jsFile = jsGlobal.Get("File")

func wrapFormDataEntryValue(v Value) FormDataEntryValue {
	if v.Type() == TypeString {
		return v.String()
	} else if v.InstanceOf(jsFile) {
		return wrapFile(v)
	}

	return nil
}

// -------------8<---------------------------------------

func NewRequest(info RequestInfo, ri ...RequestInit) Request {
	if request := jsGlobal.Get("Request"); request.Valid() {
		switch len(ri) {
		case 0:
			return wrapRequest(request.New(info))
		default:
			return wrapRequest(request.New(info, ri[0].toJSObject()))
		}
	}
	return nil
}

func NewHeaders(hi ...HeadersInit) Headers {
	if headers := jsGlobal.Get("Headers"); headers.Valid() {
		switch len(hi) {
		case 0:
			return wrapHeaders(headers.New())
		default:
			return wrapHeaders(headers.New(hi[0]))
		}
	}
	return nil
}

func NewResponse(args ...interface{}) Response {
	if response := jsGlobal.Get("Response"); response.Valid() {
		switch len(args) {
		case 0:
			return wrapResponse(response.New())
		case 1:
			if body, ok := args[0].(BodyInit); ok {
				return wrapResponse(response.New(body))
			}
		case 2:
			if body, ok := args[0].(BodyInit); ok {
				if ri, ok := args[1].(ResponseInit); ok {
					return wrapResponse(response.New(body, ri.toJSObject()))
				}
			}
		}
	}
	return nil
}
