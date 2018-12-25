// +build js,wasm

package wasm

import (
	"fmt"
	"syscall/js"
)

type (
	// https://heycam.github.io/webidl/#idl-DOMException
	DOMException interface {
		Name() string
		Message() string
		Code() DOMError
		Error() string
	}
)

// -------------8<---------------------------------------

type domExceptionImpl struct {
	js.Value
}

func NewDOMException(args ...string) DOMException {
	jsEx := js.Global().Get("DOMException")
	if isNil(jsEx) {
		return nil
	}

	switch len(args) {
	case 0:
		return wrapDOMException(jsEx.New())
	case 1:
		return wrapDOMException(jsEx.New(args[0])) // message
	default:
		return wrapDOMException(jsEx.New(args[0], args[1])) // message, name
	}
}

func wrapDOMException(v js.Value) DOMException {
	if isNil(v) {
		return nil
	}

	return &domExceptionImpl{
		Value: v,
	}
}

func (p *domExceptionImpl) Name() string {
	return p.Get("name").String()
}

func (p *domExceptionImpl) Message() string {
	return p.Get("message").String()
}

func (p *domExceptionImpl) Code() DOMError {
	return DOMError(p.Get("code").Int())
}

func (p *domExceptionImpl) Error() string {
	return fmt.Sprintf("%s : %s\n", p.Name(), p.Message())
}

// -------------8<---------------------------------------

// it wraps known types
// TODO remove mixins and non js types
// TODO Array types
func Wrap(v js.Value) interface{} {
	switch v.Type() {
	case js.TypeUndefined, js.TypeNull, js.TypeSymbol, js.TypeFunction:
		return nil
	case js.TypeBoolean:
		return v.Bool()
	case js.TypeNumber:
		return v.Float()
	case js.TypeString:
		return v.String()
	default: // js.TypeObject
		return wrapObject(v)
	}
}

// -------------8<---------------------------------------

func wrapObject(v js.Value) interface{} {
	t := JSType(v)

	switch t {
	case "AbortController":
		return wrapAbortController(v)
	case "AbortSignal":
		return wrapAbortSignal(v)
	case "AbstractRange":
		return wrapAbstractRange(v)
		/*
			case "AbstractWorker":
				return wrapAbstractWorker(v)
		*/
	case "ApplicationCache":
		return wrapApplicationCache(v)
	case "ArrayBuffer":
		return wrapArrayBuffer(v)
	case "ArrayBufferView":
		return wrapArrayBufferView(v)
	case "Attr":
		return wrapAttr(v)
	case "AudioTrack":
		return wrapAudioTrack(v)
	case "AudioTrackList":
		return wrapAudioTrackList(v)
	case "BarProp":
		return wrapBarProp(v)
	case "BeforeUnloadEvent":
		return wrapBeforeUnloadEvent(v)
	case "Blob":
		return wrapBlob(v)
	case "Body":
		return wrapBody(v)
	case "BroadcastChannel":
		return wrapBroadcastChannel(v)
	case "BufferSource":
		return wrapBufferSource(v)
	case "CDATASection":
		return wrapCDATASection(v)
	case "CanvasDrawingStyles":
		return wrapCanvasDrawingStyles(v)
	case "CanvasGradient":
		return wrapCanvasGradient(v)
	case "CanvasPathMethods":
		return wrapCanvasPathMethods(v)
	case "CanvasPattern":
		return wrapCanvasPattern(v)
	case "CanvasRenderingContext2D":
		return wrapCanvasRenderingContext2D(v)
	case "CaretPosition":
		return wrapCaretPosition(v)
	case "CharacterData":
		return wrapCharacterData(v)
	case "Clipboard":
		return wrapClipboard(v)
	case "ClipboardEvent":
		return wrapClipboardEvent(v)
	case "CloseEvent":
		return wrapCloseEvent(v)
	case "Comment":
		return wrapComment(v)
	case "CompositionEvent":
		return wrapCompositionEvent(v)
	case "Console":
		return wrapConsole(v)
	case "Coordinates":
		return wrapCoordinates(v)
	case "CSSStyleSheet":
		return wrapCSSStyleSheet(v)
	case "CSSRule":
		return wrapCSSRule(v)
	case "CSSRuleList":
		return wrapCSSRuleList(v)
	case "CSSStyleRule":
		return wrapCSSStyleRule(v)
	case "CSSImportRule":
		return wrapCSSImportRule(v)
	case "CSSGroupingRule":
		return wrapCSSGroupingRule(v)
	case "CSSPageRule":
		return wrapCSSPageRule(v)
	case "CSSPseudoElement":
		return wrapCSSPseudoElement(v)
	case "CSSMarginRule":
		return wrapCSSMarginRule(v)
	case "CSSNamespaceRule":
		return wrapCSSNamespaceRule(v)
	case "CSSStyleDeclaration":
		return wrapCSSStyleDeclaration(v)
	case "CustomEvent":
		return wrapCustomEvent(v)
	case "DOMException":
		return wrapDOMException(v)
	case "DOMImplementation":
		return wrapDOMImplementation(v)
	case "DOMMatrix":
		return wrapDOMMatrix(v)
	case "DOMMatrixReadOnly":
		return wrapDOMMatrixReadOnly(v)
	case "DOMPoint":
		return wrapDOMPoint(v)
	case "DOMPointReadOnly":
		return wrapDOMPointReadOnly(v)
	case "DOMQuad":
		return wrapDOMQuad(v)
	case "DOMRect":
		return wrapDOMRect(v)
	case "DOMRectReadOnly":
		return wrapDOMRectReadOnly(v)
	case "DOMStringMap":
		return wrapDOMStringMap(v)
	case "DOMTokenList":
		return wrapDOMTokenList(v)
	case "DataTransfer":
		return wrapDataTransfer(v)
	case "DataTransferItem":
		return wrapDataTransferItem(v)
	case "DataTransferItemList":
		return wrapDataTransferItemList(v)
	case "DataView":
		return wrapDataView(v)
	case "DedicatedWorkerGlobalScope":
		return wrapDedicatedWorkerGlobalScope(v)
	case "DocumentFragment":
		return wrapDocumentFragment(v)
	case "Document":
		return wrapDocument(v)
	case "DocumentType":
		return wrapDocumentType(v)
	case "DragEvent":
		return wrapDragEvent(v)
	case "ElementContentEditable":
		return wrapElementContentEditable(v)
	case "Element":
		return wrapElement(v)
	case "ErrorEvent":
		return wrapErrorEvent(v)
	case "Event":
		return wrapEvent(v)
	case "EventTarget":
		return wrapEventTarget(v)
	case "File":
		return wrapFile(v)
	case "FileReader":
		return wrapFileReader(v)
	case "FileReaderSync":
		return wrapFileReaderSync(v)
	case "FocusEvent":
		return wrapFocusEvent(v)
	case "GenericTransformStream":
		return wrapGenericTransformStream(v)
	case "Geolocation":
		return wrapGeolocation(v)
		/*
			case "GlobalEventHandlers":
				return wrapGlobalEventHandlersImpl(v)
		*/
	case "HTMLAnchorElement":
		return wrapHTMLAnchorElement(v)
	case "HTMLAreaElement":
		return wrapHTMLAreaElement(v)
	case "HTMLAudioElement":
		return wrapHTMLAudioElement(v)
	case "HTMLBRElement":
		return wrapHTMLBRElement(v)
	case "HTMLBaseElement":
		return wrapHTMLBaseElement(v)
	case "HTMLBodyElement":
		return wrapHTMLBodyElement(v)
	case "HTMLButtonElement":
		return wrapHTMLButtonElement(v)
	case "HTMLCanvasElement":
		return wrapHTMLCanvasElement(v)
	case "HTMLCollection":
		return wrapHTMLCollection(v)
	case "HTMLDListElement":
		return wrapHTMLDListElement(v)
	case "HTMLDataElement":
		return wrapHTMLDataElement(v)
	case "HTMLDataListElement":
		return wrapHTMLDataListElement(v)
	case "HTMLDetailsElement":
		return wrapHTMLDetailsElement(v)
	case "HTMLDialogElement":
		return wrapHTMLDialogElement(v)
	case "HTMLDivElement":
		return wrapHTMLDivElement(v)
	case "HTMLElement":
		return wrapHTMLElement(v)
	case "HTMLEmbedElement":
		return wrapHTMLEmbedElement(v)
	case "HTMLFieldSetElement":
		return wrapHTMLFieldSetElement(v)
	case "HTMLFormControl":
		return wrapHTMLFormControl(v)
	case "HTMLFormControlsCollection":
		return wrapHTMLFormControlsCollection(v)
	case "HTMLFormElement":
		return wrapHTMLFormElement(v)
	case "HTMLHRElement":
		return wrapHTMLHRElement(v)
	case "HTMLHeadElement":
		return wrapHTMLHeadElement(v)
	case "HTMLHeadingElement":
		return wrapHTMLHeadingElement(v)
	case "HTMLHtmlElement":
		return wrapHTMLHtmlElement(v)
		/*
			case "HTMLHyperlinkElementUtils":
				return newHTMLHyperlinkElementUtilsImpl(v)
		*/
	case "HTMLIFrameElement":
		return wrapHTMLIFrameElement(v)
	case "HTMLImageElement":
		return wrapHTMLImageElement(v)
	case "HTMLInputElement":
		return wrapHTMLInputElement(v)
	case "HTMLLIElement":
		return wrapHTMLLIElement(v)
	case "HTMLLabelElement":
		return wrapHTMLLabelElement(v)
	case "HTMLLegendElement":
		return wrapHTMLLegendElement(v)
	case "HTMLLinkElement":
		return wrapHTMLLinkElement(v)
	case "HTMLMapElement":
		return wrapHTMLMapElement(v)
	case "HTMLMediaElement":
		return wrapHTMLMediaElement(v)
	case "HTMLMetaElement":
		return wrapHTMLMetaElement(v)
	case "HTMLMeterElement":
		return wrapHTMLMeterElement(v)
	case "HTMLModElement":
		return wrapHTMLModElement(v)
	case "HTMLOListElement":
		return wrapHTMLOListElement(v)
	case "HTMLObjectElement":
		return wrapHTMLObjectElement(v)
	case "HTMLOptGroupElement":
		return wrapHTMLOptGroupElement(v)
	case "HTMLOptionElement":
		return wrapHTMLOptionElement(v)
	case "HTMLOptionsCollection":
		return wrapHTMLOptionsCollection(v)
	case "HTMLOrSVGElement":
		return wrapHTMLOrSVGElement(v)
	case "HTMLOrSVGScriptElement":
		return wrapHTMLOrSVGScriptElement(v)
	case "HTMLOutputElement":
		return wrapHTMLOutputElement(v)
	case "HTMLParagraphElement":
		return wrapHTMLParagraphElement(v)
	case "HTMLParamElement":
		return wrapHTMLParamElement(v)
	case "HTMLPictureElement":
		return wrapHTMLPictureElement(v)
	case "HTMLPreElement":
		return wrapHTMLPreElement(v)
	case "HTMLProgressElement":
		return wrapHTMLProgressElement(v)
	case "HTMLQuoteElement":
		return wrapHTMLQuoteElement(v)
	case "HTMLScriptElement":
		return wrapHTMLScriptElement(v)
	case "HTMLSelectElement":
		return wrapHTMLSelectElement(v)
	case "HTMLSlotElement":
		return wrapHTMLSlotElement(v)
	case "HTMLSourceElement":
		return wrapHTMLSourceElement(v)
	case "HTMLSpanElement":
		return wrapHTMLSpanElement(v)
	case "HTMLStyleElement":
		return wrapHTMLStyleElement(v)
	case "HTMLTableCaptionElement":
		return wrapHTMLTableCaptionElement(v)
	case "HTMLTableCellElement":
		return wrapHTMLTableCellElement(v)
	case "HTMLTableColElement":
		return wrapHTMLTableColElement(v)
	case "HTMLTableDataCellElement":
		return wrapHTMLTableDataCellElement(v)
	case "HTMLTableElement":
		return wrapHTMLTableElement(v)
	case "HTMLTableHeaderCellElement":
		return wrapHTMLTableHeaderCellElement(v)
	case "HTMLTableRowElement":
		return wrapHTMLTableRowElement(v)
	case "HTMLTableSectionElement":
		return wrapHTMLTableSectionElement(v)
	case "HTMLTemplateElement":
		return wrapHTMLTemplateElement(v)
	case "HTMLTextAreaElement":
		return wrapHTMLTextAreaElement(v)
	case "HTMLTimeElement":
		return wrapHTMLTimeElement(v)
	case "HTMLTitleElement":
		return wrapHTMLTitleElement(v)
	case "HTMLTrackElement":
		return wrapHTMLTrackElement(v)
	case "HTMLUListElement":
		return wrapHTMLUListElement(v)
	case "HTMLUnknownElement":
		return wrapHTMLUnknownElement(v)
	case "HTMLVideoElement":
		return wrapHTMLVideoElement(v)
	case "HashChangeEvent":
		return wrapHashChangeEvent(v)
	case "Headers":
		return wrapHeaders(v)
	case "History":
		return wrapHistory(v)
	case "IDBCursor":
		return wrapIDBCursor(v)
	case "IDBCursorSource":
		return wrapIDBCursorSource(v)
	case "IDBCursorWithValue":
		return wrapIDBCursorWithValue(v)
	case "IDBDatabase":
		return wrapIDBDatabase(v)
	case "IDBFactory":
		return wrapIDBFactory(v)
	case "IDBIndex":
		return wrapIDBIndex(v)
	case "IDBKeyRange":
		return wrapIDBKeyRange(v)
	case "IDBObjectStore":
		return wrapIDBObjectStore(v)
	case "IDBOpenDBRequest":
		return wrapIDBOpenDBRequest(v)
	case "IDBRequest":
		return wrapIDBRequest(v)
	case "IDBRequestSource":
		return wrapIDBRequestSource(v)
	case "IDBTransaction":
		return wrapIDBTransaction(v)
	case "IDBVersionChangeEvent":
		return wrapIDBVersionChangeEvent(v)
	case "ImageBitmap":
		return wrapImageBitmap(v)
	case "ImageData":
		return wrapImageData(v)
	case "InputEvent":
		return wrapInputEvent(v)
	case "KeyboardEvent":
		return wrapKeyboardEvent(v)
	case "LinkStyle":
		return wrapLinkStyle(v)
	case "Location":
		return wrapLocation(v)
	case "MediaError":
		return wrapMediaError(v)
	case "MediaList":
		return wrapMediaList(v)
	case "MediaProvider":
		return wrapMediaProvider(v)
	case "MediaQueryListEvent":
		return wrapMediaQueryListEvent(v)
	case "MediaQueryList":
		return wrapMediaQueryList(v)
	case "MediaStream":
		return wrapMediaStream(v)
	case "MediaStreamTrack":
		return wrapMediaStreamTrack(v)
	case "MessageChannel":
		return wrapMessageChannel(v)
	case "MessageEvent":
		return wrapMessageEvent(v)
	case "MessageEventSource":
		return wrapMessageEventSource(v)
	case "MessagePort":
		return wrapMessagePort(v)
	case "MouseEvent":
		return wrapMouseEvent(v)
	case "MutationObserver":
		return wrapMutationObserver(v)
	case "MutationRecord":
		return wrapMutationRecord(v)
	case "NamedNodeMap":
		return wrapNamedNodeMap(v)
		/*
			case "NavigatorConcurrentHardware":
				return newNavigatorConcurrentHardwareImpl(v)
			case "NavigatorContentUtils":
				return newNavigatorContentUtilsImpl(v)
			case "NavigatorCookies":
				return newNavigatorCookiesImpl(v)
			case "NavigatorID":
				return newNavigatorIDImpl(v)
		*/
	case "Navigator":
		return wrapNavigator(v)
		/*
			case "NavigatorLanguage":
				return newNavigatorLanguageImpl(v)
			case "NavigatorOnLine":
				return newNavigatorOnLineImpl(v)
		*/
	case "NodeFilter":
		return wrapNodeFilter(v)
	case "Node":
		return wrapNode(v)
	case "NodeIterator":
		return wrapNodeIterator(v)
	case "NodeList":
		return wrapNodeList(v)
	case "PageTransitionEvent":
		return wrapPageTransitionEvent(v)
	case "PointerEvent":
		return wrapPointerEvent(v)
	case "PopStateEvent":
		return wrapPopStateEvent(v)
	case "PositionError":
		return wrapPositionError(v)
	case "Position":
		return wrapPosition(v)
	case "ProcessingInstruction":
		return wrapProcessingInstruction(v)
	case "ProgressEvent":
		return wrapProgressEvent(v)
	case "RadioNodeList":
		return wrapRadioNodeList(v)
	case "Range":
		return wrapRange(v)
	case "ReadableStream":
		return wrapReadableStream(v)
	case "Request":
		return wrapRequest(v)
	case "Response":
		return wrapResponse(v)
	case "Screen":
		return wrapScreen(v)
	case "ShadowRoot":
		return wrapShadowRoot(v)
	case "SharedWorkerGlobalScope":
		return wrapSharedWorkerGlobalScope(v)
	case "SharedWorker":
		return wrapSharedWorker(v)
	case "SourceBuffer":
		return wrapSourceBuffer(v)
	case "StaticRange":
		return wrapStaticRange(v)
	case "StorageEvent":
		return wrapStorageEvent(v)
	case "Storage":
		return wrapStorage(v)
	case "StyleSheet":
		return wrapStyleSheet(v)
	case "StyleSheetList":
		return wrapStyleSheetList(v)
	case "TexImageSource":
		return wrapTexImageSource(v)
	case "TextDecoderCommon":
		return wrapTextDecoderCommon(v)
	case "TextDecoder":
		return wrapTextDecoder(v)
	case "TextDecoderStream":
		return wrapTextDecoderStream(v)
	case "TextEncoderCommon":
		return wrapTextEncoderCommon(v)
	case "TextEncoder":
		return wrapTextEncoder(v)
	case "TextEncoderStream":
		return wrapTextEncoderStream(v)
	case "Text":
		return wrapText(v)
	case "TextMetrics":
		return wrapTextMetrics(v)
	case "TextTrackCue":
		return wrapTextTrackCue(v)
	case "TextTrackCueList":
		return wrapTextTrackCueList(v)
	case "TextTrack":
		return wrapTextTrack(v)
	case "TextTrackList":
		return wrapTextTrackList(v)
	case "TimeRanges":
		return wrapTimeRanges(v)
	case "TouchEvent":
		return wrapTouchEvent(v)
	case "Touch":
		return wrapTouch(v)
	case "TreeWalker":
		return wrapTreeWalker(v)
	case "TransitionEvent":
		return wrapTransitionEvent(v)
	case "UIEvent":
		return wrapUIEvent(v)
	case "URL":
		return wrapURL(v)
	case "URLSearchParams":
		return wrapURLSearchParams(v)
	case "ValidityState":
		return wrapValidityState(v)
	case "VideoTrack":
		return wrapVideoTrack(v)
	case "VideoTrackList":
		return wrapVideoTrackList(v)
	case "WebGLActiveInfo":
		return wrapWebGLActiveInfo(v)
	case "WebGLBuffer":
		return wrapWebGLBuffer(v)
	case "WebGLObject":
		return wrapWebGLObject(v)
	case "WebGLProgram":
		return wrapWebGLProgram(v)
	case "WebGLRenderbuffer":
		return wrapWebGLRenderbuffer(v)
		/*
			case "WebGLRenderingContextBase":
				return newWebGLRenderingContextBaseImpl(v)
		*/
	case "WebGLRenderingContext":
		return wrapWebGLRenderingContext(v)
	case "WebGLShader":
		return wrapWebGLShader(v)
	case "WebGLShaderPrecisionFormat":
		return wrapWebGLShaderPrecisionFormat(v)
	case "WebGLTexture":
		return wrapWebGLTexture(v)
	case "WebGLUniformLocation":
		return wrapWebGLUniformLocation(v)
	case "WebSocket":
		return wrapWebSocket(v)
	case "WheelEvent":
		return wrapWheelEvent(v)
		/*
			case "WindowEventHandlers":
				return newWindowEventHandlersImpl(v)
		*/
	case "Window":
		return wrapWindow(v)
	case "WindowProxy":
		return wrapWindowProxy(v)
	case "WorkerGlobalScope":
		return wrapWorkerGlobalScope(v)
	case "Worker":
		return wrapWorker(v)
	case "WorkerLocation":
		return wrapWorkerLocation(v)
	case "WorkerNavigator":
		return wrapWorkerNavigator(v)
	case "WritableStream":
		return wrapWritableStream(v)
	case "XMLDocument":
		return wrapXMLDocument(v)
	default:
		fmt.Printf("Not supported type: %s\n", t)
		return nil
	}
}

func wrapAsElement(v js.Value) Element {
	if isNil(v) {
		return nil
	}

	if o := wrapObject(v); o != nil {
		if e, ok := o.(Element); ok {
			return e
		}
	}

	return nil
}

func wrapAsEvent(v js.Value) Event {
	if isNil(v) {
		return nil
	}

	if o := wrapObject(v); o != nil {
		if e, ok := o.(Event); ok {
			return e
		}
	}

	return nil
}

func wrapAsEventTarget(v js.Value) EventTarget {
	if isNil(v) {
		return nil
	}

	if o := wrapObject(v); o != nil {
		if e, ok := o.(EventTarget); ok {
			return e
		}
	}

	return nil
}

func wrapAsNode(v js.Value) Node {
	if isNil(v) {
		return nil
	}

	if o := wrapObject(v); o != nil {
		if n, ok := o.(Node); ok {
			return n
		}
	}
	return nil
}

// -------------8<---------------------------------------
