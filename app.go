// +build js,wasm

package wasm

import (
	"sync"
	"syscall/js"
)

var (
	doneCh          chan bool = make(chan bool)
	windowOnce      sync.Once
	docOnce         sync.Once
	currentWindow   Window
	currentDocument Document
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
	windowOnce.Do(func() {
		if currentWindow == nil {
			currentWindow = newWindow(js.Global())
		}
	})
	return currentWindow
}

func CurrentDocument() Document {
	docOnce.Do(func() {
		if currentDocument == nil {
			currentDocument = CurrentWindow().Document()
		}
	})
	return currentDocument
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
