// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewURL(url string, base ...string) URL {
	jsURL := js.Global().Get("URL")
	if isNil(jsURL) {
		return nil
	}

	switch len(base) {
	case 0:
		return newURL(jsURL.New(url))
	default:
		return newURL(jsURL.New(url, base[0]))
	}
}

func NewURLSearchParams(args ...string) URLSearchParams {
	jsURLSearchParams := js.Global().Get("URLSearchParams")
	if isNil(jsURLSearchParams) {
		return nil
	}

	switch len(args) {
	case 0:
		return newURLSearchParams(jsURLSearchParams.New())
	default:
		return newURLSearchParams(jsURLSearchParams.New(args[0]))
	}
}

// -------------8<---------------------------------------

type urlImpl struct {
	js.Value
}

func newURL(v js.Value) URL {
	if isNil(v) {
		return nil
	}

	return &urlImpl{
		Value: v,
	}
}

func (p *urlImpl) Href() string {
	return p.Get("href").String()
}

func (p *urlImpl) SetHref(href string) {
	p.Set("href", href)
}

func (p *urlImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *urlImpl) Protocol() string {
	return p.Get("protocol").String()
}

func (p *urlImpl) SetProtocol(protocol string) {
	p.Set("protocol", protocol)
}

func (p *urlImpl) Username() string {
	return p.Get("username").String()
}

func (p *urlImpl) SetUsername(userName string) {
	p.Set("username", userName)
}

func (p *urlImpl) Password() string {
	return p.Get("password").String()
}

func (p *urlImpl) SetPassword(password string) {
	p.Set("password", password)
}

func (p *urlImpl) Host() string {
	return p.Get("host").String()
}

func (p *urlImpl) SetHost(host string) {
	p.Set("host", host)
}

func (p *urlImpl) Hostname() string {
	return p.Get("hostname").String()
}

func (p *urlImpl) SetHostname(hostname string) {
	p.Set("hostname", hostname)
}

func (p *urlImpl) Port() string {
	return p.Get("port").String()
}

func (p *urlImpl) SetPort(port string) {
	p.Set("port", port)
}

func (p *urlImpl) Pathname() string {
	return p.Get("pathname").String()
}

func (p *urlImpl) SetPathname(pathname string) {
	p.Set("pathname", pathname)
}

func (p *urlImpl) Search() string {
	return p.Get("search").String()
}

func (p *urlImpl) SetSearch(search string) {
	p.Set("search", search)
}

func (p *urlImpl) SearchParams() URLSearchParams {
	return newURLSearchParams(p.Get("searchParams"))
}

func (p *urlImpl) Hash() string {
	return p.Get("hash").String()
}

func (p *urlImpl) SetHash(hash string) {
	p.Set("hash", hash)
}

func (p *urlImpl) ToJSON() string {
	return p.Call("toJSON").String()
}

// -------------8<---------------------------------------

type urlSearchParamsImpl struct {
	js.Value
}

func newURLSearchParams(v js.Value) URLSearchParams {
	if isNil(v) {
		return nil
	}

	return &urlSearchParamsImpl{
		Value: v,
	}
}

func (p *urlSearchParamsImpl) Append(name string, value string) {
	p.Call("append", name, value)
}

func (p *urlSearchParamsImpl) Delete(name string) {
	p.Call("delete", name)
}

func (p *urlSearchParamsImpl) Get(name string) string {
	return p.Call("get", name).String()
}

func (p *urlSearchParamsImpl) All(name string) []string {
	return stringSequenceToSlice(p.Call("getAll", name))
}

func (p *urlSearchParamsImpl) Has(name string) bool {
	return p.Call("has", name).Bool()
}

func (p *urlSearchParamsImpl) Set(name string, value string) {
	p.Call("set", name, value)
}

func (p *urlSearchParamsImpl) Sort() {
	p.Call("sort")
}

func (p *urlSearchParamsImpl) String() string {
	return p.Call("toString").String()
}

// -------------8<---------------------------------------
