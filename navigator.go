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

	// https://www.w3.org/TR/html52/webappapis.html#imagebitmap
	ImageBitmap interface {
		Width() int
		Height() int
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

	// https://www.w3.org/TR/html52/webappapis.html#typedefdef-imagebitmapsource
	ImageBitmapSource interface{}
)
