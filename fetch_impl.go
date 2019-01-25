// +build js,wasm

package wasm

// -------------8<---------------------------------------

type headersImpl struct {
	Value
}

func wrapHeaders(v Value) Headers {
	if v.valid() {
		return &headersImpl{
			Value: v,
		}
	}
	return nil
}

func (p *headersImpl) Append(name string, value string) {
	p.call("append", name, value)
}
func (p *headersImpl) Delete(name string) {
	p.call("delete", name)
}

func (p *headersImpl) Get(name string) string {
	return p.call("get", name).toString()
}

func (p *headersImpl) Has(name string) bool {
	return p.call("has", name).toBool()
}

func (p *headersImpl) Set(name string, value string) {
	p.call("set", name, value)
}

func (p *headersImpl) Entries() map[string]string {
	ret := make(map[string]string)

	it := p.call("entries")
	for {
		n := it.call("next")
		if n.get("done").toBool() {
			break
		}

		pair := n.get("value")

		key := pair.index(0).toString()
		value := pair.index(1).toString()

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
	if v.valid() {
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
	return p.get("bodyUsed").toBool()
}

func (p *bodyImpl) ArrayBuffer() func() (ArrayBuffer, bool) {
	return func() (ArrayBuffer, bool) {
		res, ok := await(p.call("arrayBuffer"))
		if ok {
			return wrapArrayBuffer(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) Blob() func() (Blob, bool) {
	return func() (Blob, bool) {
		res, ok := await(p.call("blob"))
		if ok {
			return wrapBlob(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) FormData() func() (FormData, bool) {
	return func() (FormData, bool) {
		res, ok := await(p.call("formData"))
		if ok {
			return wrapFormData(res), true
		}
		return nil, false
	}
}

func (p *bodyImpl) JSON() func() ([]byte, bool) {
	return func() ([]byte, bool) {
		res, ok := await(p.call("json"))
		if ok {
			return []byte(jsJSON.call("stringify", res).toString()), true
		}
		return nil, false
	}
}

func (p *bodyImpl) Text() func() (string, bool) {
	return func() (string, bool) {
		res, ok := await(p.call("text"))
		if ok {
			return res.toString(), true
		}
		return "", false
	}
}

// -------------8<---------------------------------------

type requestImpl struct {
	*bodyImpl
}

func wrapRequest(v Value) Request {
	if v.valid() {
		return &requestImpl{
			bodyImpl: newBodyImpl(v),
		}
	}
	return nil
}

func (p *requestImpl) Method() string {
	return p.get("method").toString()
}

func (p *requestImpl) URL() string {
	return p.get("url").toString()
}

func (p *requestImpl) Headers() Headers {
	return wrapHeaders(p.get("headers"))
}

func (p *requestImpl) Destination() RequestDestination {
	return RequestDestination(p.get("destination").toString())
}

func (p *requestImpl) Referrer() string {
	return p.get("referrer").toString()
}

func (p *requestImpl) ReferrerPolicy() ReferrerPolicy {
	return ReferrerPolicy(p.get("referrerPolicy").toString())
}

func (p *requestImpl) Mode() RequestMode {
	return RequestMode(p.get("mode").toString())
}

func (p *requestImpl) Credentials() RequestCredentials {
	return RequestCredentials(p.get("credentials").toString())
}

func (p *requestImpl) Cache() RequestCache {
	return RequestCache(p.get("cache").toString())
}

func (p *requestImpl) Redirect() RequestRedirect {
	return RequestRedirect(p.get("redirect").toString())
}

func (p *requestImpl) Integrity() string {
	return p.get("integrity").toString()
}

func (p *requestImpl) Keepalive() bool {
	return p.get("keepalive").toBool()
}

func (p *requestImpl) IsReloadNavigation() bool {
	return p.get("isReloadNavigation").toBool()
}

func (p *requestImpl) IsHistoryNavigation() bool {
	return p.get("isHistoryNavigation").toBool()
}

func (p *requestImpl) Signal() AbortSignal {
	return wrapAbortSignal(p.get("signal"))
}

func (p *requestImpl) Clone() Request {
	return wrapRequest(p.call("clone"))
}

// -------------8<---------------------------------------

type responseImpl struct {
	*bodyImpl
}

func wrapResponse(v Value) Response {
	if v.valid() {
		return &responseImpl{
			bodyImpl: newBodyImpl(v),
		}
	}
	return nil
}

func (p *responseImpl) Error() Response {
	return wrapResponse(p.call("error"))
}

func (p *responseImpl) Redirect(url string, status ...int) Response {
	if len(status) > 0 {
		return wrapResponse(p.call("redirect", url, status[0]))
	}

	return wrapResponse(p.call("redirect", url))
}

func (p *responseImpl) Type() ResponseType {
	return ResponseType(p.get("type").toString())
}

func (p *responseImpl) URL() string {
	return p.get("url").toString()
}

func (p *responseImpl) Redirected() bool {
	return p.get("redirected").toBool()
}

func (p *responseImpl) Status() int {
	return p.get("status").toInt()
}

func (p *responseImpl) Ok() bool {
	return p.get("ok").toBool()
}

func (p *responseImpl) StatusText() string {
	return p.get("statusText").toString()
}

func (p *responseImpl) Headers() Headers {
	return wrapHeaders(p.get("headers"))
}

func (p *responseImpl) Trailer() func() (Headers, bool) {
	return func() (Headers, bool) {
		res, ok := await(p.call("trailer"))
		if ok {
			return wrapHeaders(res), true
		}
		return nil, false
	}
}

func (p *responseImpl) Clone() Response {
	return wrapResponse(p.call("clone"))
}

// -------------8<---------------------------------------

type formDataImpl struct {
	Value
}

func NewFormData(form ...HTMLFormElement) FormData {
	if jsFormData := jsGlobal.get("FormData"); jsFormData.valid() {
		switch len(form) {
		case 0:
			return wrapFormData(jsFormData.jsNew())
		default:
			return wrapFormData(jsFormData.jsNew(JSValueOf(form[0])))
		}
	}
	return nil
}

func wrapFormData(v Value) FormData {
	if v.valid() {
		return &formDataImpl{
			Value: v,
		}
	}
	return nil
}

func (p *formDataImpl) Append(name string, value interface{}, filename ...string) {
	switch x := value.(type) {
	case string:
		p.call("append", x)
	case Blob:
		switch len(filename) {
		case 0:
			p.call("append", JSValueOf(x))
		default:
			p.call("append", JSValueOf(x), filename[0])
		}
	}
}

func (p *formDataImpl) Delete(name string) {
	p.call("delete", name)
}

func (p *formDataImpl) Get(name string) FormDataEntryValue {
	return wrapFormDataEntryValue(p.call("get", name))
}

func (p *formDataImpl) GetAll(name string) []FormDataEntryValue {
	if slc := p.call("getAll", name).toSlice(); slc != nil {
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
	return p.call("has", name).toBool()
}

func (p *formDataImpl) Set(name string, value interface{}, filename ...string) {
	switch x := value.(type) {
	case string:
		p.call("set", x)
	case Blob:
		switch len(filename) {
		case 0:
			p.call("set", JSValueOf(x))
		default:
			p.call("set", JSValueOf(x), filename[0])
		}
	}
}

func (p *formDataImpl) Values() []FormDataEntryValue {
	if slc := p.call("values").toSlice(); slc != nil {
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

func wrapFormDataEntryValue(v Value) FormDataEntryValue {
	switch v.jsType() {
	case "string":
		return v.toString()
	case "File":
		return wrapFile(v)
	default:
		return nil
	}
}

// -------------8<---------------------------------------

func NewRequest(info RequestInfo, ri ...RequestInit) Request {
	if request := jsGlobal.get("Request"); request.valid() {
		switch len(ri) {
		case 0:
			return wrapRequest(request.jsNew(info))
		default:
			return wrapRequest(request.jsNew(info, ri[0].JSValue()))
		}
	}
	return nil
}

func NewHeaders(hi ...HeadersInit) Headers {
	if headers := jsGlobal.get("Headers"); headers.valid() {
		switch len(hi) {
		case 0:
			return wrapHeaders(headers.jsNew())
		default:
			return wrapHeaders(headers.jsNew(headersInitJSValue(hi[0])))
		}
	}
	return nil
}

func NewResponse(args ...interface{}) Response {
	if response := jsGlobal.get("Response"); response.valid() {
		switch len(args) {
		case 0:
			return wrapResponse(response.jsNew())
		case 1:
			if body, ok := args[0].(BodyInit); ok {
				return wrapResponse(response.jsNew(bodyInitJSValue(body)))
			}
		case 2:
			if body, ok := args[0].(BodyInit); ok {
				if ri, ok := args[1].(ResponseInit); ok {
					return wrapResponse(response.jsNew(bodyInitJSValue(body), ri.JSValue()))
				}
			}
		}
	}
	return nil
}
