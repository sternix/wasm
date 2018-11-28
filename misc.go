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
	}
)

// -------------8<---------------------------------------

type domExceptionImpl struct {
	js.Value
}

func newDOMException(v js.Value) DOMException {
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
		return newAbortController(v)
	case "AbortSignal":
		return newAbortSignal(v)
	case "AbstractRange":
		return newAbstractRange(v)
	case "AbstractWorker":
		return newAbstractWorker(v)
	case "ApplicationCache":
		return newApplicationCache(v)
	case "ArrayBuffer":
		return newArrayBuffer(v)
	case "ArrayBufferView":
		return newArrayBufferView(v)
	case "Attr":
		return newAttr(v)
	case "AudioTrack":
		return newAudioTrack(v)
	case "AudioTrackList":
		return newAudioTrackList(v)
	case "BarProp":
		return newBarProp(v)
	case "BeforeUnloadEvent":
		return newBeforeUnloadEvent(v)
	case "Blob":
		return newBlob(v)
	case "Body":
		return newBody(v)
	case "BroadcastChannel":
		return newBroadcastChannel(v)
	case "BufferSource":
		return newBufferSource(v)
	case "CDATASection":
		return newCDATASection(v)
	case "CanvasDrawingStyles":
		return newCanvasDrawingStyles(v)
	case "CanvasGradient":
		return newCanvasGradient(v)
	case "CanvasPathMethods":
		return newCanvasPathMethods(v)
	case "CanvasPattern":
		return newCanvasPattern(v)
	case "CanvasRenderingContext2D":
		return newCanvasRenderingContext2D(v)
	case "CaretPosition":
		return newCaretPosition(v)
	case "CharacterData":
		return newCharacterData(v)
	case "Clipboard":
		return newClipboard(v)
	case "ClipboardEvent":
		return newClipboardEvent(v)
	case "CloseEvent":
		return newCloseEvent(v)
	case "Comment":
		return newComment(v)
	case "CompositionEvent":
		return newCompositionEvent(v)
	case "Console":
		return newConsole(v)
	case "Coordinates":
		return newCoordinates(v)
	case "CustomEvent":
		return newCustomEvent(v)
	case "DOMException":
		return newDOMException(v)
	case "DOMImplementation":
		return newDOMImplementation(v)
	case "DOMMatrix":
		return newDOMMatrix(v)
	case "DOMMatrixReadOnly":
		return newDOMMatrixReadOnly(v)
	case "DOMPoint":
		return newDOMPoint(v)
	case "DOMPointReadOnly":
		return newDOMPointReadOnly(v)
	case "DOMQuad":
		return newDOMQuad(v)
	case "DOMRect":
		return newDOMRect(v)
	case "DOMRectReadOnly":
		return newDOMRectReadOnly(v)
	case "DOMStringMap":
		return newDOMStringMap(v)
	case "DOMTokenList":
		return newDOMTokenList(v)
	case "DataTransfer":
		return newDataTransfer(v)
	case "DataTransferItem":
		return newDataTransferItem(v)
	case "DataTransferItemList":
		return newDataTransferItemList(v)
	case "DataView":
		return newDataView(v)
	case "DedicatedWorkerGlobalScope":
		return newDedicatedWorkerGlobalScope(v)
	case "DocumentFragment":
		return newDocumentFragment(v)
	case "Document":
		return newDocument(v)
	case "DocumentType":
		return newDocumentType(v)
	case "DragEvent":
		return newDragEvent(v)
	case "ElementContentEditable":
		return newElementContentEditable(v)
	case "Element":
		return newElement(v)
	case "ErrorEvent":
		return newErrorEvent(v)
	case "Event":
		return newEvent(v)
	case "EventTarget":
		return newEventTarget(v)
	case "File":
		return newFile(v)
	case "FileReader":
		return newFileReader(v)
	case "FileReaderSync":
		return newFileReaderSync(v)
	case "FocusEvent":
		return newFocusEvent(v)
	case "GenericTransformStream":
		return newGenericTransformStream(v)
	case "Geolocation":
		return newGeolocation(v)
	case "GlobalEventHandlers":
		return newGlobalEventHandlersImpl(v)
	case "HTMLAnchorElement":
		return newHTMLAnchorElement(v)
	case "HTMLAreaElement":
		return newHTMLAreaElement(v)
	case "HTMLAudioElement":
		return newHTMLAudioElement(v)
	case "HTMLBRElement":
		return newHTMLBRElement(v)
	case "HTMLBaseElement":
		return newHTMLBaseElement(v)
	case "HTMLBodyElement":
		return newHTMLBodyElement(v)
	case "HTMLButtonElement":
		return newHTMLButtonElement(v)
	case "HTMLCanvasElement":
		return newHTMLCanvasElement(v)
	case "HTMLCollection":
		return newHTMLCollection(v)
	case "HTMLDListElement":
		return newHTMLDListElement(v)
	case "HTMLDataElement":
		return newHTMLDataElement(v)
	case "HTMLDataListElement":
		return newHTMLDataListElement(v)
	case "HTMLDetailsElement":
		return newHTMLDetailsElement(v)
	case "HTMLDialogElement":
		return newHTMLDialogElement(v)
	case "HTMLDivElement":
		return newHTMLDivElement(v)
	case "HTMLElement":
		return newHTMLElement(v)
	case "HTMLEmbedElement":
		return newHTMLEmbedElement(v)
	case "HTMLFieldSetElement":
		return newHTMLFieldSetElement(v)
	case "HTMLFormControl":
		return newHTMLFormControl(v)
	case "HTMLFormControlsCollection":
		return newHTMLFormControlsCollection(v)
	case "HTMLFormElement":
		return newHTMLFormElement(v)
	case "HTMLHRElement":
		return newHTMLHRElement(v)
	case "HTMLHeadElement":
		return newHTMLHeadElement(v)
	case "HTMLHeadingElement":
		return newHTMLHeadingElement(v)
	case "HTMLHtmlElement":
		return newHTMLHtmlElement(v)
	case "HTMLHyperlinkElementUtils":
		return newHTMLHyperlinkElementUtilsImpl(v)
	case "HTMLIFrameElement":
		return newHTMLIFrameElement(v)
	case "HTMLImageElement":
		return newHTMLImageElement(v)
	case "HTMLInputElement":
		return newHTMLInputElement(v)
	case "HTMLLIElement":
		return newHTMLLIElement(v)
	case "HTMLLabelElement":
		return newHTMLLabelElement(v)
	case "HTMLLegendElement":
		return newHTMLLegendElement(v)
	case "HTMLLinkElement":
		return newHTMLLinkElement(v)
	case "HTMLMapElement":
		return newHTMLMapElement(v)
	case "HTMLMediaElement":
		return newHTMLMediaElement(v)
	case "HTMLMetaElement":
		return newHTMLMetaElement(v)
	case "HTMLMeterElement":
		return newHTMLMeterElement(v)
	case "HTMLModElement":
		return newHTMLModElement(v)
	case "HTMLOListElement":
		return newHTMLOListElement(v)
	case "HTMLObjectElement":
		return newHTMLObjectElement(v)
	case "HTMLOptGroupElement":
		return newHTMLOptGroupElement(v)
	case "HTMLOptionElement":
		return newHTMLOptionElement(v)
	case "HTMLOptionsCollection":
		return newHTMLOptionsCollection(v)
	case "HTMLOrSVGElement":
		return newHTMLOrSVGElement(v)
	case "HTMLOrSVGScriptElement":
		return newHTMLOrSVGScriptElement(v)
	case "HTMLOutputElement":
		return newHTMLOutputElement(v)
	case "HTMLParagraphElement":
		return newHTMLParagraphElement(v)
	case "HTMLParamElement":
		return newHTMLParamElement(v)
	case "HTMLPictureElement":
		return newHTMLPictureElement(v)
	case "HTMLPreElement":
		return newHTMLPreElement(v)
	case "HTMLProgressElement":
		return newHTMLProgressElement(v)
	case "HTMLQuoteElement":
		return newHTMLQuoteElement(v)
	case "HTMLScriptElement":
		return newHTMLScriptElement(v)
	case "HTMLSelectElement":
		return newHTMLSelectElement(v)
	case "HTMLSlotElement":
		return newHTMLSlotElement(v)
	case "HTMLSourceElement":
		return newHTMLSourceElement(v)
	case "HTMLSpanElement":
		return newHTMLSpanElement(v)
	case "HTMLStyleElement":
		return newHTMLStyleElement(v)
	case "HTMLTableCaptionElement":
		return newHTMLTableCaptionElement(v)
	case "HTMLTableCellElement":
		return newHTMLTableCellElement(v)
	case "HTMLTableColElement":
		return newHTMLTableColElement(v)
	case "HTMLTableDataCellElement":
		return newHTMLTableDataCellElement(v)
	case "HTMLTableElement":
		return newHTMLTableElement(v)
	case "HTMLTableHeaderCellElement":
		return newHTMLTableHeaderCellElement(v)
	case "HTMLTableRowElement":
		return newHTMLTableRowElement(v)
	case "HTMLTableSectionElement":
		return newHTMLTableSectionElement(v)
	case "HTMLTemplateElement":
		return newHTMLTemplateElement(v)
	case "HTMLTextAreaElement":
		return newHTMLTextAreaElement(v)
	case "HTMLTimeElement":
		return newHTMLTimeElement(v)
	case "HTMLTitleElement":
		return newHTMLTitleElement(v)
	case "HTMLTrackElement":
		return newHTMLTrackElement(v)
	case "HTMLUListElement":
		return newHTMLUListElement(v)
	case "HTMLUnknownElement":
		return newHTMLUnknownElement(v)
	case "HTMLVideoElement":
		return newHTMLVideoElement(v)
	case "HashChangeEvent":
		return newHashChangeEvent(v)
	case "Headers":
		return newHeaders(v)
	case "History":
		return newHistory(v)
	case "IDBCursor":
		return newIDBCursor(v)
	case "IDBCursorSource":
		return newIDBCursorSource(v)
	case "IDBCursorWithValue":
		return newIDBCursorWithValue(v)
	case "IDBDatabase":
		return newIDBDatabase(v)
	case "IDBFactory":
		return newIDBFactory(v)
	case "IDBIndex":
		return newIDBIndex(v)
	case "IDBKeyRange":
		return newIDBKeyRange(v)
	case "IDBObjectStore":
		return newIDBObjectStore(v)
	case "IDBOpenDBRequest":
		return newIDBOpenDBRequest(v)
	case "IDBRequest":
		return newIDBRequest(v)
	case "IDBRequestSource":
		return newIDBRequestSource(v)
	case "IDBTransaction":
		return newIDBTransaction(v)
	case "IDBVersionChangeEvent":
		return newIDBVersionChangeEvent(v)
	case "ImageBitmap":
		return newImageBitmap(v)
	case "ImageData":
		return newImageData(v)
	case "InputEvent":
		return newInputEvent(v)
	case "KeyboardEvent":
		return newKeyboardEvent(v)
	case "LinkStyle":
		return newLinkStyle(v)
	case "Location":
		return newLocation(v)
	case "MediaError":
		return newMediaError(v)
	case "MediaList":
		return newMediaList(v)
	case "MediaProvider":
		return newMediaProvider(v)
	case "MediaQueryListEvent":
		return newMediaQueryListEvent(v)
	case "MediaQueryList":
		return newMediaQueryList(v)
	case "MediaStream":
		return newMediaStream(v)
	case "MediaStreamTrack":
		return newMediaStreamTrack(v)
	case "MessageChannel":
		return newMessageChannel(v)
	case "MessageEvent":
		return newMessageEvent(v)
	case "MessageEventSource":
		return newMessageEventSource(v)
	case "MessagePort":
		return newMessagePort(v)
	case "MouseEvent":
		return newMouseEvent(v)
	case "MutationObserver":
		return newMutationObserver(v)
	case "MutationRecord":
		return newMutationRecord(v)
	case "NamedNodeMap":
		return newNamedNodeMap(v)
	case "NavigatorConcurrentHardware":
		return newNavigatorConcurrentHardwareImpl(v)
	case "NavigatorContentUtils":
		return newNavigatorContentUtilsImpl(v)
	case "NavigatorCookies":
		return newNavigatorCookiesImpl(v)
	case "NavigatorID":
		return newNavigatorIDImpl(v)
	case "Navigator":
		return newNavigator(v)
	case "NavigatorLanguage":
		return newNavigatorLanguageImpl(v)
	case "NavigatorOnLine":
		return newNavigatorOnLineImpl(v)
	case "NodeFilter":
		return newNodeFilter(v)
	case "Node":
		return newNode(v)
	case "NodeIterator":
		return newNodeIterator(v)
	case "NodeList":
		return newNodeList(v)
	case "PageTransitionEvent":
		return newPageTransitionEvent(v)
	case "PopStateEvent":
		return newPopStateEvent(v)
	case "PositionError":
		return newPositionError(v)
	case "Position":
		return newPosition(v)
	case "ProcessingInstruction":
		return newProcessingInstruction(v)
	case "ProgressEvent":
		return newProgressEvent(v)
	case "RadioNodeList":
		return newRadioNodeList(v)
	case "Range":
		return newRange(v)
	case "ReadableStream":
		return newReadableStream(v)
	case "Request":
		return newRequest(v)
	case "Response":
		return newResponse(v)
	case "Screen":
		return newScreen(v)
	case "ShadowRoot":
		return newShadowRoot(v)
	case "SharedWorkerGlobalScope":
		return newSharedWorkerGlobalScope(v)
	case "SharedWorker":
		return newSharedWorker(v)
	case "SourceBuffer":
		return newSourceBuffer(v)
	case "StaticRange":
		return newStaticRange(v)
	case "StorageEvent":
		return newStorageEvent(v)
	case "Storage":
		return newStorage(v)
	case "StyleSheet":
		return newStyleSheet(v)
	case "TexImageSource":
		return newTexImageSource(v)
	case "TextDecoderCommon":
		return newTextDecoderCommon(v)
	case "TextDecoder":
		return newTextDecoder(v)
	case "TextDecoderStream":
		return newTextDecoderStream(v)
	case "TextEncoderCommon":
		return newTextEncoderCommon(v)
	case "TextEncoder":
		return newTextEncoder(v)
	case "TextEncoderStream":
		return newTextEncoderStream(v)
	case "Text":
		return newText(v)
	case "TextMetrics":
		return newTextMetrics(v)
	case "TextTrackCue":
		return newTextTrackCue(v)
	case "TextTrackCueList":
		return newTextTrackCueList(v)
	case "TextTrack":
		return newTextTrack(v)
	case "TextTrackList":
		return newTextTrackList(v)
	case "TimeRanges":
		return newTimeRanges(v)
	case "TouchEvent":
		return newTouchEvent(v)
	case "Touch":
		return newTouch(v)
	case "TreeWalker":
		return newTreeWalker(v)
	case "TransitionEvent":
		return newTransitionEvent(v)
	case "UIEvent":
		return newUIEvent(v)
	case "URL":
		return newURL(v)
	case "URLSearchParams":
		return newURLSearchParams(v)
	case "ValidityState":
		return newValidityState(v)
	case "VideoTrack":
		return newVideoTrack(v)
	case "VideoTrackList":
		return newVideoTrackList(v)
	case "WebGLActiveInfo":
		return newWebGLActiveInfo(v)
	case "WebGLBuffer":
		return newWebGLBuffer(v)
	case "WebGLObject":
		return newWebGLObject(v)
	case "WebGLProgram":
		return newWebGLProgram(v)
	case "WebGLRenderbuffer":
		return newWebGLRenderbuffer(v)
	case "WebGLRenderingContextBase":
		return newWebGLRenderingContextBaseImpl(v)
	case "WebGLRenderingContext":
		return newWebGLRenderingContext(v)
	case "WebGLShader":
		return newWebGLShader(v)
	case "WebGLShaderPrecisionFormat":
		return newWebGLShaderPrecisionFormat(v)
	case "WebGLTexture":
		return newWebGLTexture(v)
	case "WebGLUniformLocation":
		return newWebGLUniformLocation(v)
	case "WebSocket":
		return newWebSocket(v)
	case "WheelEvent":
		return newWheelEvent(v)
	case "WindowEventHandlers":
		return newWindowEventHandlersImpl(v)
	case "Window":
		return newWindow(v)
	case "WindowProxy":
		return newWindowProxy(v)
	case "WorkerGlobalScope":
		return newWorkerGlobalScopeImpl(v)
	case "Worker":
		return newWorker(v)
	case "WorkerLocation":
		return newWorkerLocation(v)
	case "WorkerNavigator":
		return newWorkerNavigator(v)
	case "WritableStream":
		return newWritableStream(v)
	case "XMLDocument":
		return newXMLDocument(v)
	default:
		fmt.Printf("Not supported type: %s", t)
		return nil
	}
}

func wrapElement(v js.Value) Element {
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

func wrapEvent(v js.Value) Event {
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

// -------------8<---------------------------------------
