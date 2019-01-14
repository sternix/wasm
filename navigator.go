// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/html52/webappapis.html#navigator
	Navigator interface {
		NavigatorID
		NavigatorLanguage
		NavigatorOnLine
		NavigatorContentUtils
		NavigatorCookies
		NavigatorConcurrentHardware //5.3

		// https://w3c.github.io/geolocation-api/#idl-index
		Geolocation() Geolocation

		// https://w3c.github.io/clipboard-apis/#idl-index
		Clipboard() Clipboard

		// https://www.w3.org/TR/pointerevents/#extensions-to-the-navigator-interface
		MaxTouchPoints() int
	}

	// https://www.w3.org/TR/html52/webappapis.html#navigatorid
	NavigatorID interface {
		AppCodeName() string
		AppName() string
		AppVersion() string
		Platform() string
		Product() string
		UserAgent() string
	}

	// https://www.w3.org/TR/html52/webappapis.html#navigatorlanguage
	NavigatorLanguage interface {
		Language() string
		Languages() []string
	}

	// https://www.w3.org/TR/html52/webappapis.html#navigatorcontentutils
	NavigatorContentUtils interface {
		RegisterProtocolHandler(string, string, string)
		RegisterContentHandler(string, string, string)
		IsProtocolHandlerRegistered(string, string) string
		IsContentHandlerRegistered(string, string) string
		UnregisterProtocolHandler(string, string)
		UnregisterContentHandler(string, string)
	}

	// https://www.w3.org/TR/html52/webappapis.html#navigatorcookies
	NavigatorCookies interface {
		CookieEnabled() bool
	}

	// https://html.spec.whatwg.org/multipage/system-state.html#mimetype
	/*
		MimeType interface {
			Type() string
			Description() string
			Suffixes() string
			EnabledPlugin() Plugin
		}
	*/

	// https://html.spec.whatwg.org/multipage/system-state.html#dom-plugin
	/*
		Plugin interface {
			Name() string
			Description() string
			Filename() string
			Length() int
			Item(int) MimeType
			NamedItem(string) MimeType
		}
	*/

	// https://html.spec.whatwg.org/multipage/system-state.html#mimetypearray
	/*
		MimeTypeArray interface {
			Length() int
			Item(int) MimeType
			NamedItem() MimeType
		}
	*/

	// https://html.spec.whatwg.org/multipage/system-state.html#pluginarray
	/*
		PluginArray interface {
			Refresh(...bool)
			Length() int
			Item(int) Plugin
			NamedItem(string) Plugin
		}
	*/

	// https://html.spec.whatwg.org/multipage/system-state.html#navigatorplugins
	/*
		NavigatorPlugins interface {
			Plugins() PluginArray
			MimeTypes() MimeTypeArray
			//JavaEnabled() bool // is this required ?
		}
	*/

	// https://html.spec.whatwg.org/multipage/imagebitmap-and-animations.html#imagebitmap
	ImageBitmap interface {
		Width() int
		Height() int
		Close()
	}

	/*
		typedef (HTMLImageElement or
			HTMLVideoElement or
			HTMLCanvasElement or
			Blob or
			ImageData or
			CanvasRenderingContext2D or
			ImageBitmap) ImageBitmapSource;
	*/

	// https://html.spec.whatwg.org/multipage/imagebitmap-and-animations.html#imagebitmapsource
	// typedef (CanvasImageSource or Blob or ImageData) ImageBitmapSource;
	ImageBitmapSource interface{}
)

type ImageOrientation string

const (
	ImageOrientationNone  ImageOrientation = "none"
	ImageOrientationFlipY ImageOrientation = "flipY"
)

type PremultiplyAlpha string

const (
	PremultiplyAlphaNone        PremultiplyAlpha = "none"
	PremultiplyAlphaPremultiply PremultiplyAlpha = "premultiply"
	PremultiplyAlphaDefault     PremultiplyAlpha = "default"
)

type ColorSpaceConversion string

const (
	ColorSpaceConversionNone    ColorSpaceConversion = "none"
	ColorSpaceConversionDefault ColorSpaceConversion = "default"
)

type ResizeQuality string

const (
	ResizeQualityPixelated ResizeQuality = "pixelated"
	ResizeQualityLow       ResizeQuality = "low"
	ResizeQualityMedium    ResizeQuality = "medium"
	ResizeQualityHigh      ResizeQuality = "high"
)

type ImageBitmapOptions struct {
	ImageOrientation     ImageOrientation
	PremultiplyAlpha     PremultiplyAlpha
	ColorSpaceConversion ColorSpaceConversion
	ResizeWidth          uint
	ResizeHeight         uint
	ResizeQuality        ResizeQuality
}

func (p ImageBitmapOptions) toJSObject() Value {
	o := jsObject.jsNew()

	if p.ImageOrientation != "none" {
		o.set("imageOrientation", string(p.ImageOrientation))
	}

	if p.PremultiplyAlpha != "default" {
		o.set("premultiplyAlpha", string(p.PremultiplyAlpha))
	}

	if p.ColorSpaceConversion != "default" {
		o.set("colorSpaceConversion", string(p.ColorSpaceConversion))
	}

	o.set("resizeWidth", p.ResizeWidth)
	o.set("resizeHeight", p.ResizeHeight)

	if p.ResizeQuality != "low" {
		o.set("resizeQuality", string(p.ResizeQuality))
	}

	return o
}
