// +build js,wasm

package wasm

// -------------8<---------------------------------------

type navigatorImpl struct {
	*navigatorIDImpl
	*navigatorLanguageImpl
	*navigatorOnLineImpl
	*navigatorContentUtilsImpl
	*navigatorCookiesImpl
	*navigatorConcurrentHardwareImpl
	Value
}

func wrapNavigator(v Value) Navigator {
	if v.Valid() {
		return &navigatorImpl{
			navigatorIDImpl:                 newNavigatorIDImpl(v),
			navigatorLanguageImpl:           newNavigatorLanguageImpl(v),
			navigatorOnLineImpl:             newNavigatorOnLineImpl(v),
			navigatorContentUtilsImpl:       newNavigatorContentUtilsImpl(v),
			navigatorCookiesImpl:            newNavigatorCookiesImpl(v),
			navigatorConcurrentHardwareImpl: newNavigatorConcurrentHardwareImpl(v),
			Value:                           v,
		}
	}
	return nil
}

func (p *navigatorImpl) Geolocation() Geolocation {
	return wrapGeolocation(p.Get("geolocation"))
}

func (p *navigatorImpl) Clipboard() Clipboard {
	return wrapClipboard(p.Get("clipboard"))
}

func (p *navigatorImpl) MaxTouchPoints() int {
	return p.Get("maxTouchPoints").Int()
}

// -------------8<---------------------------------------

var _ NavigatorID = &navigatorIDImpl{}

type navigatorIDImpl struct {
	Value
}

func newNavigatorIDImpl(v Value) *navigatorIDImpl {
	if v.Valid() {
		return &navigatorIDImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorIDImpl) AppCodeName() string {
	return p.Get("appCodeName").String()
}

func (p *navigatorIDImpl) AppName() string {
	return p.Get("appName").String()
}

func (p *navigatorIDImpl) AppVersion() string {
	return p.Get("appVersion").String()
}

func (p *navigatorIDImpl) Platform() string {
	return p.Get("platform").String()
}

func (p *navigatorIDImpl) Product() string {
	return p.Get("product").String()
}

func (p *navigatorIDImpl) UserAgent() string {
	return p.Get("userAgent").String()
}

// -------------8<---------------------------------------

var _ NavigatorLanguage = &navigatorLanguageImpl{}

type navigatorLanguageImpl struct {
	Value
}

func newNavigatorLanguageImpl(v Value) *navigatorLanguageImpl {
	if v.Valid() {
		return &navigatorLanguageImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorLanguageImpl) Language() string {
	return p.Get("language").String()
}

func (p *navigatorLanguageImpl) Languages() []string {
	if s := p.Get("languages").ToSlice(); s != nil {
		ret := make([]string, len(s))
		for i, v := range s {
			ret[i] = v.String()
		}
		return ret
	}

	return nil
}

// -------------8<---------------------------------------

var _ NavigatorContentUtils = &navigatorContentUtilsImpl{}

type navigatorContentUtilsImpl struct {
	Value
}

func newNavigatorContentUtilsImpl(v Value) *navigatorContentUtilsImpl {
	if v.Valid() {
		return &navigatorContentUtilsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorContentUtilsImpl) RegisterProtocolHandler(scheme string, url string, title string) {
	p.Call("registerProtocolHandler", scheme, url, title)
}

func (p *navigatorContentUtilsImpl) RegisterContentHandler(mimeType string, url string, title string) {
	p.Call("registerContentHandler", mimeType, url, title)
}

func (p *navigatorContentUtilsImpl) IsProtocolHandlerRegistered(scheme string, url string) string {
	return p.Call("isProtocolHandlerRegistered", scheme, url).String()
}

func (p *navigatorContentUtilsImpl) IsContentHandlerRegistered(mimeType string, url string) string {
	return p.Call("isContentHandlerRegistered", mimeType, url).String()
}

func (p *navigatorContentUtilsImpl) UnregisterProtocolHandler(scheme string, url string) {
	p.Call("unregisterProtocolHandler", scheme, url)
}

func (p *navigatorContentUtilsImpl) UnregisterContentHandler(mimeType string, url string) {
	p.Call("unregisterContentHandler", mimeType, url)
}

// -------------8<---------------------------------------

var _ NavigatorCookies = &navigatorCookiesImpl{}

type navigatorCookiesImpl struct {
	Value
}

func newNavigatorCookiesImpl(v Value) *navigatorCookiesImpl {
	if v.Valid() {
		return &navigatorCookiesImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorCookiesImpl) CookieEnabled() bool {
	return p.Get("cookieEnabled").Bool()
}

// -------------8<---------------------------------------

type imageBitmapImpl struct {
	Value
}

func wrapImageBitmap(v Value) ImageBitmap {
	if v.Valid() {
		return &imageBitmapImpl{
			Value: v,
		}
	}
	return nil
}

func (p *imageBitmapImpl) Width() int {
	return p.Get("width").Int()
}

func (p *imageBitmapImpl) Height() int {
	return p.Get("height").Int()
}

func (p *imageBitmapImpl) Close() {
	p.Call("close")
}

// -------------8<---------------------------------------

/*
	typedef (HTMLImageElement or
		HTMLVideoElement or
		HTMLCanvasElement or
		Blob or
		ImageData or
		CanvasRenderingContext2D or
		ImageBitmap) ImageBitmapSource;
*/

// https://www.w3.org/TR/html52/webappapis.html#typedefdef-imagebitmapsource
//	ImageBitmapSource interface{}
