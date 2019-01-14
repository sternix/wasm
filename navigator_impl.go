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
	if v.valid() {
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
	return wrapGeolocation(p.get("geolocation"))
}

func (p *navigatorImpl) Clipboard() Clipboard {
	return wrapClipboard(p.get("clipboard"))
}

func (p *navigatorImpl) MaxTouchPoints() int {
	return p.get("maxTouchPoints").toInt()
}

// -------------8<---------------------------------------

var _ NavigatorID = &navigatorIDImpl{}

type navigatorIDImpl struct {
	Value
}

func newNavigatorIDImpl(v Value) *navigatorIDImpl {
	if v.valid() {
		return &navigatorIDImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorIDImpl) AppCodeName() string {
	return p.get("appCodeName").toString()
}

func (p *navigatorIDImpl) AppName() string {
	return p.get("appName").toString()
}

func (p *navigatorIDImpl) AppVersion() string {
	return p.get("appVersion").toString()
}

func (p *navigatorIDImpl) Platform() string {
	return p.get("platform").toString()
}

func (p *navigatorIDImpl) Product() string {
	return p.get("product").toString()
}

func (p *navigatorIDImpl) UserAgent() string {
	return p.get("userAgent").toString()
}

// -------------8<---------------------------------------

var _ NavigatorLanguage = &navigatorLanguageImpl{}

type navigatorLanguageImpl struct {
	Value
}

func newNavigatorLanguageImpl(v Value) *navigatorLanguageImpl {
	if v.valid() {
		return &navigatorLanguageImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorLanguageImpl) Language() string {
	return p.get("language").toString()
}

func (p *navigatorLanguageImpl) Languages() []string {
	if s := p.get("languages").toSlice(); s != nil {
		ret := make([]string, len(s))
		for i, v := range s {
			ret[i] = v.toString()
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
	if v.valid() {
		return &navigatorContentUtilsImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorContentUtilsImpl) RegisterProtocolHandler(scheme string, url string, title string) {
	p.call("registerProtocolHandler", scheme, url, title)
}

func (p *navigatorContentUtilsImpl) RegisterContentHandler(mimeType string, url string, title string) {
	p.call("registerContentHandler", mimeType, url, title)
}

func (p *navigatorContentUtilsImpl) IsProtocolHandlerRegistered(scheme string, url string) string {
	return p.call("isProtocolHandlerRegistered", scheme, url).toString()
}

func (p *navigatorContentUtilsImpl) IsContentHandlerRegistered(mimeType string, url string) string {
	return p.call("isContentHandlerRegistered", mimeType, url).toString()
}

func (p *navigatorContentUtilsImpl) UnregisterProtocolHandler(scheme string, url string) {
	p.call("unregisterProtocolHandler", scheme, url)
}

func (p *navigatorContentUtilsImpl) UnregisterContentHandler(mimeType string, url string) {
	p.call("unregisterContentHandler", mimeType, url)
}

// -------------8<---------------------------------------

var _ NavigatorCookies = &navigatorCookiesImpl{}

type navigatorCookiesImpl struct {
	Value
}

func newNavigatorCookiesImpl(v Value) *navigatorCookiesImpl {
	if v.valid() {
		return &navigatorCookiesImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorCookiesImpl) CookieEnabled() bool {
	return p.get("cookieEnabled").toBool()
}

// -------------8<---------------------------------------

type imageBitmapImpl struct {
	Value
}

func wrapImageBitmap(v Value) ImageBitmap {
	if v.valid() {
		return &imageBitmapImpl{
			Value: v,
		}
	}
	return nil
}

func (p *imageBitmapImpl) Width() int {
	return p.get("width").toInt()
}

func (p *imageBitmapImpl) Height() int {
	return p.get("height").toInt()
}

func (p *imageBitmapImpl) Close() {
	p.call("close")
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
