// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type navigatorImpl struct {
	*navigatorIDImpl
	*navigatorLanguageImpl
	*navigatorOnLineImpl
	*navigatorContentUtilsImpl
	*navigatorCookiesImpl
	*navigatorConcurrentHardwareImpl
	js.Value
}

func newNavigator(v js.Value) Navigator {
	if isNil(v) {
		return nil
	}

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

func (p *navigatorImpl) Geolocation() Geolocation {
	return newGeolocation(p.Get("geolocation"))
}

func (p *navigatorImpl) Clipboard() Clipboard {
	return newClipboard(p.Get("clipboard"))
}

// -------------8<---------------------------------------

var _ NavigatorID = &navigatorIDImpl{}

type navigatorIDImpl struct {
	js.Value
}

func newNavigatorIDImpl(v js.Value) *navigatorIDImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorIDImpl{
		Value: v,
	}
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
	js.Value
}

func newNavigatorLanguageImpl(v js.Value) *navigatorLanguageImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorLanguageImpl{
		Value: v,
	}
}

func (p *navigatorLanguageImpl) Language() string {
	return p.Get("language").String()
}

func (p *navigatorLanguageImpl) Languages() []string {
	if s := arrayToSlice(p.Get("languages")); s != nil {
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
	js.Value
}

func newNavigatorContentUtilsImpl(v js.Value) *navigatorContentUtilsImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorContentUtilsImpl{
		Value: v,
	}
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
	js.Value
}

func newNavigatorCookiesImpl(v js.Value) *navigatorCookiesImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorCookiesImpl{
		Value: v,
	}
}

func (p *navigatorCookiesImpl) CookieEnabled() bool {
	return p.Get("cookieEnabled").Bool()
}

// -------------8<---------------------------------------

type imageBitmapImpl struct {
	js.Value
}

func newImageBitmap(v js.Value) ImageBitmap {
	if isNil(v) {
		return nil
	}

	return &imageBitmapImpl{
		Value: v,
	}
}

func (p *imageBitmapImpl) Width() int {
	return p.Get("width").Int()
}

func (p *imageBitmapImpl) Height() int {
	return p.Get("height").Int()
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
