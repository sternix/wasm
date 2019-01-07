// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewWorker(scriptURL string, wo ...WorkerOptions) Worker {
	if jsWorker := jsGlobal.Get("Worker"); jsWorker.Valid() {
		switch len(wo) {
		case 0:
			return wrapWorker(jsWorker.New(scriptURL))
		default:
			return wrapWorker(jsWorker.New(scriptURL, wo[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type workerGlobalScopeImpl struct {
	*eventTargetImpl
	*windowOrWorkerGlobalScopeImpl
	Value
}

func wrapWorkerGlobalScope(v Value) WorkerGlobalScope {
	if p := newWorkerGlobalScopeImpl(v); p != nil {
		return p
	}
	return nil
}

func newWorkerGlobalScopeImpl(v Value) *workerGlobalScopeImpl {
	if v.Valid() {
		return &workerGlobalScopeImpl{
			eventTargetImpl:               newEventTargetImpl(v),
			windowOrWorkerGlobalScopeImpl: newWindowOrWorkerGlobalScopeImpl(v),
			Value:                         v,
		}
	}
	return nil
}

func (p *workerGlobalScopeImpl) Self() WorkerGlobalScope {
	return wrapWorkerGlobalScope(p.Get("self"))
}

func (p *workerGlobalScopeImpl) Location() WorkerLocation {
	return wrapWorkerLocation(p.Get("location"))
}

func (p *workerGlobalScopeImpl) Navigator() WorkerNavigator {
	return wrapWorkerNavigator(p.Get("navigator"))
}

func (p *workerGlobalScopeImpl) ImportScripts(urls ...string) {
	if len(urls) > 0 {
		var params []interface{}
		for _, url := range urls {
			params = append(params, url)
		}
		p.Call("importScripts", params...)
	}
}

func (p *workerGlobalScopeImpl) Close() {
	p.Call("close")
}

func (p *workerGlobalScopeImpl) OnError(fn func(Event)) EventHandler { // TODO OnErrorEventHandler
	return p.On("error", fn)
}

func (p *workerGlobalScopeImpl) OnLanguageChange(fn func(Event)) EventHandler {
	return p.On("languagechange", fn)
}

func (p *workerGlobalScopeImpl) OnOffline(fn func(Event)) EventHandler {
	return p.On("offline", fn)
}

func (p *workerGlobalScopeImpl) OnOnline(fn func(Event)) EventHandler {
	return p.On("online", fn)
}

func (p *workerGlobalScopeImpl) OnRejectionHandled(fn func(Event)) EventHandler {
	return p.On("rejectionhandled", fn)
}

func (p *workerGlobalScopeImpl) OnUnhandledRejection(fn func(Event)) EventHandler {
	return p.On("unhandledrejection", fn)
}

// -------------8<---------------------------------------

type dedicatedWorkerGlobalScopeImpl struct {
	*workerGlobalScopeImpl
}

func wrapDedicatedWorkerGlobalScope(v Value) DedicatedWorkerGlobalScope {
	if v.Valid() {
		return &dedicatedWorkerGlobalScopeImpl{
			workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
		}
	}
	return nil
}

func (p *dedicatedWorkerGlobalScopeImpl) Name() string {
	return p.Get("name").String()
}

func (p *dedicatedWorkerGlobalScopeImpl) PostMessage(message interface{}) {
	p.Call("postMessage", message)
}

func (p *dedicatedWorkerGlobalScopeImpl) OnMessage(fn func(Event)) EventHandler {
	return p.On("message", fn)
}

func (p *dedicatedWorkerGlobalScopeImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

// -------------8<---------------------------------------

type sharedWorkerGlobalScopeImpl struct {
	*workerGlobalScopeImpl
}

func wrapSharedWorkerGlobalScope(v Value) SharedWorkerGlobalScope {
	if v.Valid() {
		return &sharedWorkerGlobalScopeImpl{
			workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
		}
	}
	return nil
}

func (p *workerGlobalScopeImpl) Name() string {
	return p.Get("name").String()
}

func (p *workerGlobalScopeImpl) ApplicationCache() ApplicationCache {
	return wrapApplicationCache(p.Get("applicationCache"))
}

func (p *workerGlobalScopeImpl) OnConnect(fn func(Event)) EventHandler {
	return p.On("connect", fn)
}

// -------------8<---------------------------------------

type applicationCacheImpl struct {
	*eventTargetImpl
}

func wrapApplicationCache(v Value) ApplicationCache {
	if v.Valid() {
		return &applicationCacheImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *applicationCacheImpl) Update() {
	p.Call("update")
}

func (p *applicationCacheImpl) Abort() {
	p.Call("abort")
}

func (p *applicationCacheImpl) SwapCache() {
	p.Call("swapCache")
}

func (p *applicationCacheImpl) OnChecking(fn func(Event)) EventHandler {
	return p.On("checking", fn)
}

func (p *applicationCacheImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

func (p *applicationCacheImpl) OnNoUpdate(fn func(Event)) EventHandler {
	return p.On("noupdate", fn)
}

func (p *applicationCacheImpl) OnDownloading(fn func(Event)) EventHandler {
	return p.On("downloading", fn)
}

func (p *applicationCacheImpl) OnProgress(fn func(Event)) EventHandler {
	return p.On("progress", fn)
}

func (p *applicationCacheImpl) OnUpdateReady(fn func(Event)) EventHandler {
	return p.On("updateready", fn)
}

func (p *applicationCacheImpl) OnCached(fn func(Event)) EventHandler {
	return p.On("cached", fn)
}

func (p *applicationCacheImpl) OnObsolete(fn func(Event)) EventHandler {
	return p.On("obsolete", fn)
}

// -------------8<---------------------------------------

type abstractWorkerImpl struct {
	*eventTargetImpl
}

func newAbstractWorkerImpl(et *eventTargetImpl) *abstractWorkerImpl {
	return &abstractWorkerImpl{
		eventTargetImpl: et,
	}
}

func (p *abstractWorkerImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

// -------------8<---------------------------------------

type workerImpl struct {
	*eventTargetImpl
	*abstractWorkerImpl
	Value
}

func wrapWorker(v Value) Worker {
	if v.Valid() {
		wi := &workerImpl{
			eventTargetImpl: newEventTargetImpl(v),
			Value:           v,
		}
		wi.abstractWorkerImpl = newAbstractWorkerImpl(wi.eventTargetImpl)
		return wi
	}
	return nil
}

func (p *workerImpl) Terminate() {
	p.Call("terminate")
}

func (p *workerImpl) PostMessage(message interface{}) {
	p.Call("postMessage", message)
}

func (p *workerImpl) OnMessage(fn func(Event)) EventHandler {
	return p.On("message", fn)
}

func (p *workerImpl) OnMessageError(fn func(Event)) EventHandler {
	return p.On("messageerror", fn)
}

// -------------8<---------------------------------------

type sharedWorkerImpl struct {
	*eventTargetImpl
	*abstractWorkerImpl
	Value
}

func wrapSharedWorker(v Value) SharedWorker {
	if v.Valid() {
		swi := &sharedWorkerImpl{
			eventTargetImpl: newEventTargetImpl(v),
			Value:           v,
		}
		swi.abstractWorkerImpl = newAbstractWorkerImpl(swi.eventTargetImpl)
		return swi
	}
	return nil
}

func (p *sharedWorkerImpl) Port() MessagePort {
	return wrapMessagePort(p.Get("port"))
}

// -------------8<---------------------------------------

var _ NavigatorConcurrentHardware = &navigatorConcurrentHardwareImpl{}

type navigatorConcurrentHardwareImpl struct {
	Value
}

func newNavigatorConcurrentHardwareImpl(v Value) *navigatorConcurrentHardwareImpl {
	if v.Valid() {
		return &navigatorConcurrentHardwareImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorConcurrentHardwareImpl) HardwareConcurrency() int {
	return p.Get("hardwareConcurrency").Int()
}

// -------------8<---------------------------------------

type workerNavigatorImpl struct {
	*navigatorIDImpl
	*navigatorLanguageImpl
	*navigatorOnLineImpl
	*navigatorConcurrentHardwareImpl
}

func wrapWorkerNavigator(v Value) WorkerNavigator {
	if v.Valid() {
		return &workerNavigatorImpl{
			navigatorIDImpl:                 newNavigatorIDImpl(v),
			navigatorLanguageImpl:           newNavigatorLanguageImpl(v),
			navigatorOnLineImpl:             newNavigatorOnLineImpl(v),
			navigatorConcurrentHardwareImpl: newNavigatorConcurrentHardwareImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type workerLocationImpl struct {
	Value
}

func wrapWorkerLocation(v Value) WorkerLocation {
	if p := newWorkerLocationImpl(v); p != nil {
		return p
	}
	return nil
}

func newWorkerLocationImpl(v Value) *workerLocationImpl {
	if v.Valid() {
		return &workerLocationImpl{
			Value: v,
		}
	}
	return nil
}

func (p *workerLocationImpl) Href() string {
	return p.Get("href").String()
}

func (p *workerLocationImpl) Origin() string {
	return p.Get("origin").String()
}

func (p *workerLocationImpl) Protocol() string {
	return p.Get("protocol").String()
}

func (p *workerLocationImpl) Host() string {
	return p.Get("host").String()
}

func (p *workerLocationImpl) Hostname() string {
	return p.Get("hostname").String()
}

func (p *workerLocationImpl) Port() string {
	return p.Get("port").String()
}

func (p *workerLocationImpl) Pathname() string {
	return p.Get("pathname").String()
}

func (p *workerLocationImpl) Search() string {
	return p.Get("search").String()
}

func (p *workerLocationImpl) Hash() string {
	return p.Get("hash").String()
}

// -------------8<---------------------------------------
