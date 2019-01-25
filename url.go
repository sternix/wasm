// +build js,wasm

package wasm

type (
	// https://url.spec.whatwg.org/#url
	URL interface {
		Href() string
		SetHref(string)
		Origin() string
		Protocol() string
		SetProtocol(string)
		Username() string
		SetUsername(string)
		Password() string
		SetPassword(string)
		Host() string
		SetHost(string)
		Hostname() string
		SetHostname(string)
		Port() string
		SetPort(string)
		Pathname() string
		SetPathname(string)
		Search() string
		SetSearch(string)
		SearchParams() URLSearchParams
		Hash() string
		SetHash(string)
		ToJSON() string
	}

	// https://url.spec.whatwg.org/#urlsearchparams
	URLSearchParams interface {
		Append(string, string)
		Delete(string)
		Get(string) string // USVString? , maybe return bool,string
		All(string) []string
		Has(string) bool
		Set(string, string)
		Sort()
		String() string

		JSValue() jsValue

		// TODO
		// iterable
	}
)
