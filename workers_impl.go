// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewWorker(scriptURL string, wo ...WorkerOptions) Worker {
	jsWorker := js.Global().Get("Worker")
	if isNil(jsWorker) {
		return nil
	}

	switch len(wo) {
	case 0:
		return newWorker(jsWorker.New(scriptURL))
	default:
		return newWorker(jsWorker.New(scriptURL, toJSONObject(wo[0])))
	}
}

// -------------8<---------------------------------------

type workerGlobalScopeImpl struct {
	*eventTargetImpl
	*windowOrWorkerGlobalScopeImpl
	js.Value
}

func newWorkerGlobalScope(v js.Value) WorkerGlobalScope {
	if p := newWorkerGlobalScopeImpl(v); p != nil {
		return p
	}
	return nil
}

func newWorkerGlobalScopeImpl(v js.Value) *workerGlobalScopeImpl {
	if isNil(v) {
		return nil
	}

	return &workerGlobalScopeImpl{
		eventTargetImpl:               newEventTargetImpl(v),
		windowOrWorkerGlobalScopeImpl: newWindowOrWorkerGlobalScopeImpl(v),
		Value:                         v,
	}
}

func (p *workerGlobalScopeImpl) Self() WorkerGlobalScope {
	return newWorkerGlobalScope(p.Get("self"))
}

func (p *workerGlobalScopeImpl) Location() WorkerLocation {
	return newWorkerLocation(p.Get("location"))
}

func (p *workerGlobalScopeImpl) Navigator() WorkerNavigator {
	return newWorkerNavigator(p.Get("navigator"))
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

func newDedicatedWorkerGlobalScope(v js.Value) DedicatedWorkerGlobalScope {
	if isNil(v) {
		return nil
	}

	return &dedicatedWorkerGlobalScopeImpl{
		workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
	}
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

func newSharedWorkerGlobalScope(v js.Value) SharedWorkerGlobalScope {
	if isNil(v) {
		return nil
	}

	return &sharedWorkerGlobalScopeImpl{
		workerGlobalScopeImpl: newWorkerGlobalScopeImpl(v),
	}
}

func (p *workerGlobalScopeImpl) Name() string {
	return p.Get("name").String()
}

func (p *workerGlobalScopeImpl) ApplicationCache() ApplicationCache {
	return newApplicationCache(p.Get("applicationCache"))
}

func (p *workerGlobalScopeImpl) OnConnect(fn func(Event)) EventHandler {
	return p.On("connect", fn)
}

// -------------8<---------------------------------------

type applicationCacheImpl struct {
	*eventTargetImpl
}

func newApplicationCache(v js.Value) ApplicationCache {
	if isNil(v) {
		return nil
	}

	return &applicationCacheImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
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
	js.Value
}

func newAbstractWorker(v js.Value) AbstractWorker {
	if p := newAbstractWorkerImpl(v); p != nil {
		return p
	}
	return nil
}

func newAbstractWorkerImpl(v js.Value) *abstractWorkerImpl {
	if isNil(v) {
		return nil
	}

	return &abstractWorkerImpl{
		Value: v,
	}
}

func (p *abstractWorkerImpl) OnError(fn func(Event)) EventHandler {
	return On("error", fn)
}

// -------------8<---------------------------------------

type workerImpl struct {
	*eventTargetImpl
	*abstractWorkerImpl
	js.Value
}

func newWorker(v js.Value) Worker {
	if isNil(v) {
		return nil
	}

	return &workerImpl{
		eventTargetImpl:    newEventTargetImpl(v),
		abstractWorkerImpl: newAbstractWorkerImpl(v),
		Value:              v,
	}
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
	js.Value
}

func newSharedWorker(v js.Value) SharedWorker {
	if isNil(v) {
		return nil
	}

	return &sharedWorkerImpl{
		eventTargetImpl:    newEventTargetImpl(v),
		abstractWorkerImpl: newAbstractWorkerImpl(v),
		Value:              v,
	}
}

func (p *sharedWorkerImpl) Port() MessagePort {
	return newMessagePort(p.Get("port"))
}

// -------------8<---------------------------------------

var _ NavigatorConcurrentHardware = &navigatorConcurrentHardwareImpl{}

type navigatorConcurrentHardwareImpl struct {
	js.Value
}

func newNavigatorConcurrentHardwareImpl(v js.Value) *navigatorConcurrentHardwareImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorConcurrentHardwareImpl{
		Value: v,
	}
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

func newWorkerNavigator(v js.Value) WorkerNavigator {
	if isNil(v) {
		return nil
	}

	return &workerNavigatorImpl{
		navigatorIDImpl:                 newNavigatorIDImpl(v),
		navigatorLanguageImpl:           newNavigatorLanguageImpl(v),
		navigatorOnLineImpl:             newNavigatorOnLineImpl(v),
		navigatorConcurrentHardwareImpl: newNavigatorConcurrentHardwareImpl(v),
	}
}

// -------------8<---------------------------------------

type workerLocationImpl struct {
	js.Value
}

func newWorkerLocation(v js.Value) WorkerLocation {
	if p := newWorkerLocationImpl(v); p != nil {
		return p
	}
	return nil
}

func newWorkerLocationImpl(v js.Value) *workerLocationImpl {
	if isNil(v) {
		return nil
	}

	return &workerLocationImpl{
		Value: v,
	}
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
