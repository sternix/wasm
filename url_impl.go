// +build js,wasm

package wasm

// -------------8<---------------------------------------

func CreateObjectURL(source interface{}) (string, error) {
	jsURL := jsGlobal.get("URL")

	switch x := source.(type) {
	case File, Blob, MediaSource:
		return jsURL.call("createObjectURL", JSValueOf(x)).toString(), nil
	default:
		return "", errInvalidType
	}
}

func RevokeObjectURL(objectURL string) {
	jsGlobal.get("URL").call("revokeObjectURL", objectURL)
}

// -------------8<---------------------------------------

func NewURL(url string, base ...string) URL {
	if jsURL := jsGlobal.get("URL"); jsURL.valid() {
		switch len(base) {
		case 0:
			return wrapURL(jsURL.jsNew(url))
		default:
			return wrapURL(jsURL.jsNew(url, base[0]))
		}
	}
	return nil
}

func NewURLSearchParams(args ...string) URLSearchParams {
	if jsURLSearchParams := jsGlobal.get("URLSearchParams"); jsURLSearchParams.valid() {
		switch len(args) {
		case 0:
			return wrapURLSearchParams(jsURLSearchParams.jsNew())
		default:
			return wrapURLSearchParams(jsURLSearchParams.jsNew(args[0]))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type urlImpl struct {
	Value
}

func wrapURL(v Value) URL {
	if v.valid() {
		return &urlImpl{
			Value: v,
		}
	}
	return nil
}

func (p *urlImpl) Href() string {
	return p.get("href").toString()
}

func (p *urlImpl) SetHref(href string) {
	p.set("href", href)
}

func (p *urlImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *urlImpl) Protocol() string {
	return p.get("protocol").toString()
}

func (p *urlImpl) SetProtocol(protocol string) {
	p.set("protocol", protocol)
}

func (p *urlImpl) Username() string {
	return p.get("username").toString()
}

func (p *urlImpl) SetUsername(userName string) {
	p.set("username", userName)
}

func (p *urlImpl) Password() string {
	return p.get("password").toString()
}

func (p *urlImpl) SetPassword(password string) {
	p.set("password", password)
}

func (p *urlImpl) Host() string {
	return p.get("host").toString()
}

func (p *urlImpl) SetHost(host string) {
	p.set("host", host)
}

func (p *urlImpl) Hostname() string {
	return p.get("hostname").toString()
}

func (p *urlImpl) SetHostname(hostname string) {
	p.set("hostname", hostname)
}

func (p *urlImpl) Port() string {
	return p.get("port").toString()
}

func (p *urlImpl) SetPort(port string) {
	p.set("port", port)
}

func (p *urlImpl) Pathname() string {
	return p.get("pathname").toString()
}

func (p *urlImpl) SetPathname(pathname string) {
	p.set("pathname", pathname)
}

func (p *urlImpl) Search() string {
	return p.get("search").toString()
}

func (p *urlImpl) SetSearch(search string) {
	p.set("search", search)
}

func (p *urlImpl) SearchParams() URLSearchParams {
	return wrapURLSearchParams(p.get("searchParams"))
}

func (p *urlImpl) Hash() string {
	return p.get("hash").toString()
}

func (p *urlImpl) SetHash(hash string) {
	p.set("hash", hash)
}

func (p *urlImpl) ToJSON() string {
	return p.call("toJSON").toString()
}

// -------------8<---------------------------------------

type urlSearchParamsImpl struct {
	Value
}

func wrapURLSearchParams(v Value) URLSearchParams {
	if v.valid() {
		return &urlSearchParamsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *urlSearchParamsImpl) Append(name string, value string) {
	p.call("append", name, value)
}

func (p *urlSearchParamsImpl) Delete(name string) {
	p.call("delete", name)
}

func (p *urlSearchParamsImpl) Get(name string) string {
	return p.call("get", name).toString()
}

func (p *urlSearchParamsImpl) All(name string) []string {
	return stringSequenceToSlice(p.call("getAll", name))
}

func (p *urlSearchParamsImpl) Has(name string) bool {
	return p.call("has", name).toBool()
}

func (p *urlSearchParamsImpl) Set(name string, value string) {
	p.call("set", name, value)
}

func (p *urlSearchParamsImpl) Sort() {
	p.call("sort")
}

func (p *urlSearchParamsImpl) String() string {
	return p.call("toString").toString()
}

// -------------8<---------------------------------------
