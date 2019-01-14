// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewWorker(scriptURL string, wo ...WorkerOptions) Worker {
	if jsWorker := jsGlobal.get("Worker"); jsWorker.valid() {
		switch len(wo) {
		case 0:
			return wrapWorker(jsWorker.jsNew(scriptURL))
		default:
			return wrapWorker(jsWorker.jsNew(scriptURL, wo[0].toJSObject()))
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
	if v.valid() {
		return &workerGlobalScopeImpl{
			eventTargetImpl:               newEventTargetImpl(v),
			windowOrWorkerGlobalScopeImpl: newWindowOrWorkerGlobalScopeImpl(v),
			Value:                         v,
		}
	}
	return nil
}

func (p *workerGlobalScopeImpl) Self() WorkerGlobalScope {
	return wrapWorkerGlobalScope(p.get("self"))
}

func (p *workerGlobalScopeImpl) Location() WorkerLocation {
	return wrapWorkerLocation(p.get("location"))
}

func (p *workerGlobalScopeImpl) Navigator() WorkerNavigator {
	return wrapWorkerNavigator(p.get("navigator"))
}

func (p *workerGlobalScopeImpl) ImportScripts(urls ...string) {
	if len(urls) > 0 {
		var params []interface{}
		for _, url := range urls {
			params = append(params, url)
		}
		p.call("importScripts", params...)
	}
}

func (p *workerGlobalScopeImpl) Close() {
	p.call("close")
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
	if v.valid() {
		return &dedicatedWorkerGlobalScopeImpl{
			workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
		}
	}
	return nil
}

func (p *dedicatedWorkerGlobalScopeImpl) Name() string {
	return p.get("name").toString()
}

func (p *dedicatedWorkerGlobalScopeImpl) PostMessage(message interface{}) {
	p.call("postMessage", message)
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
	if v.valid() {
		return &sharedWorkerGlobalScopeImpl{
			workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
		}
	}
	return nil
}

func (p *workerGlobalScopeImpl) Name() string {
	return p.get("name").toString()
}

func (p *workerGlobalScopeImpl) ApplicationCache() ApplicationCache {
	return wrapApplicationCache(p.get("applicationCache"))
}

func (p *workerGlobalScopeImpl) OnConnect(fn func(Event)) EventHandler {
	return p.On("connect", fn)
}

// -------------8<---------------------------------------

type applicationCacheImpl struct {
	*eventTargetImpl
}

func wrapApplicationCache(v Value) ApplicationCache {
	if v.valid() {
		return &applicationCacheImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *applicationCacheImpl) Update() {
	p.call("update")
}

func (p *applicationCacheImpl) Abort() {
	p.call("abort")
}

func (p *applicationCacheImpl) SwapCache() {
	p.call("swapCache")
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
	if v.valid() {
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
	p.call("terminate")
}

func (p *workerImpl) PostMessage(message interface{}) {
	p.call("postMessage", message)
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
	if v.valid() {
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
	return wrapMessagePort(p.get("port"))
}

// -------------8<---------------------------------------

var _ NavigatorConcurrentHardware = &navigatorConcurrentHardwareImpl{}

type navigatorConcurrentHardwareImpl struct {
	Value
}

func newNavigatorConcurrentHardwareImpl(v Value) *navigatorConcurrentHardwareImpl {
	if v.valid() {
		return &navigatorConcurrentHardwareImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorConcurrentHardwareImpl) HardwareConcurrency() uint {
	return p.get("hardwareConcurrency").toUint()
}

// -------------8<---------------------------------------

type workerNavigatorImpl struct {
	*navigatorIDImpl
	*navigatorLanguageImpl
	*navigatorOnLineImpl
	*navigatorConcurrentHardwareImpl
}

func wrapWorkerNavigator(v Value) WorkerNavigator {
	if v.valid() {
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
	if v.valid() {
		return &workerLocationImpl{
			Value: v,
		}
	}
	return nil
}

func (p *workerLocationImpl) Href() string {
	return p.get("href").toString()
}

func (p *workerLocationImpl) Origin() string {
	return p.get("origin").toString()
}

func (p *workerLocationImpl) Protocol() string {
	return p.get("protocol").toString()
}

func (p *workerLocationImpl) Host() string {
	return p.get("host").toString()
}

func (p *workerLocationImpl) Hostname() string {
	return p.get("hostname").toString()
}

func (p *workerLocationImpl) Port() string {
	return p.get("port").toString()
}

func (p *workerLocationImpl) Pathname() string {
	return p.get("pathname").toString()
}

func (p *workerLocationImpl) Search() string {
	return p.get("search").toString()
}

func (p *workerLocationImpl) Hash() string {
	return p.get("hash").toString()
}

// -------------8<---------------------------------------
