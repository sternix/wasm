// +build js,wasm

package wasm

import (
	"syscall/js"
)

var (
	doneCh        chan bool = make(chan bool)
	currentWindow Window
)

func Wait() {
	On("beforeunload", func(Event) {
		doneCh <- true
	})

	<-doneCh
}

func Exit() {
	doneCh <- true
}

func On(event string, fn func(ev Event)) EventHandler {
	return EventHandlerFunc(event, fn)
}

func CurrentWindow() Window {
	if currentWindow == nil {
		currentWindow = newWindow(js.Global())
	}
	return currentWindow
}

func CurrentDocument() Document {
	if w := CurrentWindow(); w != nil {
		return w.Document()
	}
	return nil
}

func SessionStorage() Storage {
	return newStorage(js.Global().Get("sessionStorage"))
}

func LocalStorage() Storage {
	return newStorage(js.Global().Get("localStorage"))
}

func CreateObjectURL(source interface{}) (string, error) {
	jsURL := js.Global().Get("URL")

	switch x := source.(type) {
	case File:
		return jsURL.Call("createObjectURL", x.JSValue()).String(), nil
	case Blob:
		return jsURL.Call("createObjectURL", x.JSValue()).String(), nil
	case MediaSource:
		return jsURL.Call("createObjectURL", x.JSValue()).String(), nil
	default:
		return "", errInvalidType
	}
}

func RevokeObjectURL(objectURL string) {
	js.Global().Get("URL").Call("revokeObjectURL", objectURL)
}
