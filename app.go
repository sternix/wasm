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
	CurrentWindow().On("beforeunload", func(Event) {
		doneCh <- true
	})

	<-doneCh
}

func Exit() {
	doneCh <- true
}

func CurrentWindow() Window {
	windowOnce.Do(func() {
		if currentWindow == nil {
			currentWindow = wrapWindow(js.Global())
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
	return wrapStorage(js.Global().Get("sessionStorage"))
}

func LocalStorage() Storage {
	return wrapStorage(js.Global().Get("localStorage"))
}
