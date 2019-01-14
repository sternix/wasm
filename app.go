// +build js,wasm

package wasm

import (
	"sync"
)

var (
	doneCh          chan bool = make(chan bool)
	windowOnce      sync.Once
	docOnce         sync.Once
	currentWindow   Window
	currentDocument Document
)

func Loop() {
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
			currentWindow = wrapWindow(jsGlobal)
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
	return wrapStorage(jsGlobal.get("sessionStorage"))
}

func LocalStorage() Storage {
	return wrapStorage(jsGlobal.get("localStorage"))
}
