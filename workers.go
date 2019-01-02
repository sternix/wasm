// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://w3c.github.io/workers/#workerglobalscope0
	WorkerGlobalScope interface {
		EventTarget
		WindowOrWorkerGlobalScope

		Self() WorkerGlobalScope
		Location() WorkerLocation
		Navigator() WorkerNavigator
		ImportScripts(...string)
		Close()
		OnError(func(Event)) EventHandler //OnErrorEventHandler
		OnLanguageChange(func(Event)) EventHandler
		OnOffline(func(Event)) EventHandler
		OnOnline(func(Event)) EventHandler
		OnRejectionHandled(func(Event)) EventHandler
		OnUnhandledRejection(func(Event)) EventHandler
	}

	// https://w3c.github.io/workers/#dedicatedworkerglobalscope
	DedicatedWorkerGlobalScope interface {
		WorkerGlobalScope

		Name() string
		PostMessage(interface{}) // optional sequence<Transferable> transfer
		OnMessage(func(Event)) EventHandler
		OnMessageError(func(Event)) EventHandler
	}

	// https://w3c.github.io/workers/#sharedworkerglobalscope
	SharedWorkerGlobalScope interface {
		WorkerGlobalScope

		Name() string
		ApplicationCache() ApplicationCache
		OnConnect(func(Event)) EventHandler
	}

	// https://html.spec.whatwg.org/multipage/offline.html#applicationcache
	ApplicationCache interface {
		EventTarget

		Update()
		Abort()
		SwapCache()
		OnChecking(func(Event)) EventHandler
		OnError(func(Event)) EventHandler
		OnNoUpdate(func(Event)) EventHandler
		OnDownloading(func(Event)) EventHandler
		OnProgress(func(Event)) EventHandler
		OnUpdateReady(func(Event)) EventHandler
		OnCached(func(Event)) EventHandler
		OnObsolete(func(Event)) EventHandler
	}

	// https://w3c.github.io/workers/#abstractworker
	AbstractWorker interface {
		OnError(func(Event)) EventHandler
	}

	// https://w3c.github.io/workers/#worker
	Worker interface {
		EventTarget
		AbstractWorker

		Terminate()
		PostMessage(interface{}) // optional sequence<Transferable> transfer
		OnMessage(func(Event)) EventHandler
		OnMessageError(func(Event)) EventHandler
	}

	// https://w3c.github.io/workers/#sharedworker
	SharedWorker interface {
		EventTarget
		AbstractWorker

		Port() MessagePort
	}

	// https://w3c.github.io/workers/#navigatorconcurrenthardware
	NavigatorConcurrentHardware interface {
		HardwareConcurrency() int
	}

	// https://w3c.github.io/workers/#workernavigator
	WorkerNavigator interface {
		NavigatorID
		NavigatorLanguage
		NavigatorOnLine
		NavigatorConcurrentHardware
	}

	// https://w3c.github.io/workers/#workerlocation
	WorkerLocation interface {
		Href() string
		Origin() string
		Protocol() string
		Host() string
		Hostname() string
		Port() string
		Pathname() string
		Search() string
		Hash() string
	}
)

// https://w3c.github.io/workers/#enumdef-workertype
type WorkerType string

const (
	WorkerTypeClassic WorkerType = "classic"
	WorkerTypeModule  WorkerType = "module"
)

// -------------8<---------------------------------------

// https://w3c.github.io/workers/#dictdef-workeroptions
type WorkerOptions struct {
	Type        WorkerType         // default classic`
	Credentials RequestCredentials // default omit
	Name        string
}

func (p WorkerOptions) toJSObject() js.Value {
	o := jsObject.New()
	o.Set("type", string(p.Type))
	o.Set("credentials", string(p.Credentials))
	o.Set("name", p.Name)
	return o
}
