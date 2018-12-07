// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type headersImpl struct {
	js.Value
}

func newHeaders(v js.Value) Headers {
	if isNil(v) {
		return nil
	}

	return &headersImpl{
		Value: v,
	}
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
	js.Value
}

func newBody(v js.Value) Body {
	if p := newBodyImpl(v); p != nil {
		return p
	}
	return nil
}

func newBodyImpl(v js.Value) *bodyImpl {
	if isNil(v) {
		return nil
	}

	return &bodyImpl{
		Value: v,
	}
}

func (p *bodyImpl) Body() ReadableStream {
	// TODO
	return nil
}

func (p *bodyImpl) BodyUsed() bool {
	return p.Get("bodyUsed").Bool()
}

func (p *bodyImpl) ArrayBuffer() Promise {
	return newPromiseImpl(p.Call("arrayBuffer"))
}

func (p *bodyImpl) Blob() Promise {
	return newPromiseImpl(p.Call("blob"))
}

func (p *bodyImpl) FormData() Promise {
	return newPromiseImpl(p.Call("formData"))
}

func (p *bodyImpl) JSON() Promise {
	return newPromiseImpl(p.Call("json"))
}

func (p *bodyImpl) Text() Promise {
	return newPromiseImpl(p.Call("text"))
}

// -------------8<---------------------------------------

type requestImpl struct {
	*bodyImpl
}

func newRequest(v js.Value) Request {
	if isNil(v) {
		return nil
	}

	return &requestImpl{
		bodyImpl: newBodyImpl(v),
	}
}

func (p *requestImpl) Method() string {
	return p.Get("method").String()
}

func (p *requestImpl) URL() string {
	return p.Get("url").String()
}

func (p *requestImpl) Headers() Headers {
	return newHeaders(p.Get("headers"))
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
	return newAbortSignal(p.Get("signal"))
}

func (p *requestImpl) Clone() Request {
	return newRequest(p.Call("clone"))
}

// -------------8<---------------------------------------

type responseImpl struct {
	*bodyImpl
}

func newResponse(v js.Value) Response {
	if isNil(v) {
		return nil
	}

	return &responseImpl{
		bodyImpl: newBodyImpl(v),
	}
}

func (p *responseImpl) Error() Response {
	return newResponse(p.Call("error"))
}

func (p *responseImpl) Redirect(url string, status ...int) Response {
	if len(status) > 0 {
		return newResponse(p.Call("redirect", url, status[0]))
	}

	return newResponse(p.Call("redirect", url))
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
	return newHeaders(p.Get("headers"))
}

func (p *responseImpl) Trailer() func() (Headers, bool) {
	return func() (Headers, bool) {
		res, ok := Await(p.Call("trailer"))
		if ok {
			return newHeaders(res), true
		}
		return nil, false
	}
}

func (p *responseImpl) Clone() Response {
	return newResponse(p.Call("clone"))
}

// -------------8<---------------------------------------

func NewRequest(info RequestInfo, ri ...RequestInit) Request {
	request := js.Global().Get("Request")
	if len(ri) > 0 {
		return newRequest(request.New(info, ri[0].toDict()))
	}

	return newRequest(request.New(info))
}

func NewHeaders(hi ...HeadersInit) Headers {
	headers := js.Global().Get("Headers")

	if len(hi) > 0 {
		return newHeaders(headers.New(hi[0]))
	}

	return newHeaders(headers.New())
}

func NewResponse(args ...interface{}) Response {
	response := js.Global().Get("Response")

	switch len(args) {
	case 1:
		if body, ok := args[0].(BodyInit); ok {
			return newResponse(response.New(body))
		}
	case 2:
		if body, ok := args[0].(BodyInit); ok {
			if ri, ok := args[1].(ResponseInit); ok {
				return newResponse(response.New(body, ri.toDict()))
			}
		}
	}

	return newResponse(response.New())
}
