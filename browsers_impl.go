// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewPopStateEvent(typ string, p ...PopStateEventInit) PopStateEvent {
	jsPopStateEvent := js.Global().Get("PopStateEvent")
	if isNil(jsPopStateEvent) {
		return nil
	}

	if len(p) > 0 {
		return newPopStateEvent(jsPopStateEvent.New(typ, toJSONObject(p[0])))
	}

	return newPopStateEvent(jsPopStateEvent.New(typ))
}

func NewHashChangeEvent(typ string, p ...HashChangeEventInit) HashChangeEvent {
	jsHashChangeEvent := js.Global().Get("HashChangeEvent")
	if isNil(jsHashChangeEvent) {
		return nil
	}

	if len(p) > 0 {
		return newHashChangeEvent(jsHashChangeEvent.New(typ, toJSONObject(p[0])))
	}

	return newHashChangeEvent(jsHashChangeEvent.New(typ))
}

func NewPageTransitionEvent(typ string, p ...PageTransitionEventInit) PageTransitionEvent {
	jsPageTransitionEvent := js.Global().Get("PageTransitionEvent")
	if isNil(jsPageTransitionEvent) {
		return nil
	}

	if len(p) > 0 {
		return newPageTransitionEvent(jsPageTransitionEvent.New(typ, toJSONObject(p[0])))
	}

	return newPageTransitionEvent(jsPageTransitionEvent.New(typ))
}

// -------------8<---------------------------------------

type windowImpl struct {
	*eventTargetImpl
	*windowOrWorkerGlobalScopeImpl
	*globalEventHandlersImpl
	*windowEventHandlersImpl
	js.Value
}

func newWindow(v js.Value) Window {
	if p := newWindowImpl(v); p != nil {
		return p
	}
	return nil
}

func newWindowImpl(v js.Value) *windowImpl {
	if isNil(v) {
		return nil
	}

	return &windowImpl{
		eventTargetImpl:               newEventTargetImpl(v),
		windowOrWorkerGlobalScopeImpl: newWindowOrWorkerGlobalScopeImpl(v),
		globalEventHandlersImpl:       newGlobalEventHandlersImpl(v),
		windowEventHandlersImpl:       newWindowEventHandlersImpl(v),
		Value:                         v,
	}
}

func (p *windowImpl) Console() Console {
	return newConsole(p.Get("console"))
}

func (p *windowImpl) Window() WindowProxy {
	return newWindowProxy(p.Get("window"))
}

func (p *windowImpl) Self() WindowProxy {
	return newWindowProxy(p.Get("self"))
}

func (p *windowImpl) Document() Document {
	return newDocument(p.Get("document"))
}

func (p *windowImpl) Name() string {
	return p.Get("name").String()
}

func (p *windowImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *windowImpl) Location() Location {
	return newLocation(p.Get("location"))
}

func (p *windowImpl) History() History {
	return newHistory(p.Get("history"))
}

func (p *windowImpl) Locationbar() BarProp {
	return newBarProp(p.Get("locationbar"))
}

func (p *windowImpl) Menubar() BarProp {
	return newBarProp(p.Get("menubar"))
}

func (p *windowImpl) Personalbar() BarProp {
	return newBarProp(p.Get("personalbar"))
}

func (p *windowImpl) Scrollbars() BarProp {
	return newBarProp(p.Get("scrollbars"))
}

func (p *windowImpl) Statusbar() BarProp {
	return newBarProp(p.Get("statusbar"))
}

func (p *windowImpl) Toolbar() BarProp {
	return newBarProp(p.Get("toolbar"))
}

func (p *windowImpl) Status() string {
	return p.Get("status").String()
}

func (p *windowImpl) SetStatus(status string) {
	p.Set("status", status)
}

func (p *windowImpl) Close() {
	p.Call("close")
}

func (p *windowImpl) Closed() bool {
	return p.Get("closed").Bool()
}

func (p *windowImpl) Stop() {
	p.Call("stop")
}

func (p *windowImpl) Focus() {
	p.Call("focus")
}

func (p *windowImpl) Blur() {
	p.Call("blur")
}

func (p *windowImpl) Frames() WindowProxy {
	return newWindowProxy(p.Get("frames"))
}

func (p *windowImpl) Length() int {
	return p.Get("length").Int()
}

func (p *windowImpl) Top() WindowProxy {
	return newWindowProxy(p.Get("top"))
}

func (p *windowImpl) Opener() WindowProxy {
	return newWindowProxy(p.Get("opener"))
}

func (p *windowImpl) Parent() WindowProxy {
	return newWindowProxy(p.Get("parent"))
}

func (p *windowImpl) FrameElement() Element {
	return newElement(p.Get("frameElement"))
}

func (p *windowImpl) Open(args ...interface{}) WindowProxy {
	switch len(args) {
	case 1:
		if url, ok := args[0].(string); ok {
			return newWindowProxy(p.Call("open", url))
		}
	case 2:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				return newWindowProxy(p.Call("open", url, target))
			}
		}
	case 3:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				if features, ok := args[2].(string); ok {
					return newWindowProxy(p.Call("open", url, target, features))
				}
			}
		}
	case 4:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				if features, ok := args[2].(string); ok {
					if replace, ok := args[3].(bool); ok {
						return newWindowProxy(p.Call("open", url, target, features, replace))
					}
				}
			}
		}
	}

	return newWindowProxy(p.Call("open"))
}

func (p *windowImpl) Navigator() Navigator {
	return newNavigator(p.Get("navigator"))
}

func (p *windowImpl) Alert(msg ...string) {
	switch len(msg) {
	case 0:
		p.Call("alert")
	default:
		p.Call("alert", msg[0])
	}
}

func (p *windowImpl) Confirm(msg ...string) bool {
	switch len(msg) {
	case 0:
		return p.Call("confirm").Bool()
	default:
		return p.Call("confirm", msg[0]).Bool()
	}
}

func (p *windowImpl) Prompt(args ...string) string {
	switch len(args) {
	case 0:
		return p.Call("prompt").String()
	case 1:
		return p.Call("prompt", args[0]).String() // message
	default:
		return p.Call("prompt", args[0], args[1]).String() // message, default
	}
}

func (p *windowImpl) Print() {
	p.Call("print")
}

func (p *windowImpl) RequestAnimationFrame(cb FrameRequestCallback) int {
	return p.Call("requestAnimationFrame", cb.jsCallback).Int()
}

func (p *windowImpl) CancelAnimationFrame(handle int) {
	p.Call("cancelAnimationFrame", handle)
}

func (p *windowImpl) MatchMedia(query string) MediaQueryList {
	return newMediaQueryList(p.Call("matchMedia", query))
}

func (p *windowImpl) Screen() Screen {
	return newScreen(p.Get("screen"))
}

func (p *windowImpl) MoveTo(x int, y int) {
	p.Call("moveTo", x, y)
}

func (p *windowImpl) MoveBy(x int, y int) {
	p.Call("moveBy", x, y)
}

func (p *windowImpl) ResizeTo(x int, y int) {
	p.Call("resizeTo", x, y)
}

func (p *windowImpl) ResizeBy(x int, y int) {
	p.Call("resizeBy", x, y)
}

func (p *windowImpl) InnerWidth() int {
	return p.Get("innerWidth").Int()
}

func (p *windowImpl) InnerHeight() int {
	return p.Get("innerHeight").Int()
}

func (p *windowImpl) ScrollX() float64 {
	return p.Get("scrollX").Float()
}

func (p *windowImpl) PageXOffset() float64 {
	return p.Get("pageXOffset").Float()
}

func (p *windowImpl) ScrollY() float64 {
	return p.Get("scrollY").Float()
}

func (p *windowImpl) PageYOffset() float64 {
	return p.Get("pageYOffset").Float()
}

func (p *windowImpl) Scroll(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("scroll")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.Call("scroll", toJSONObject(options))
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.Call("scroll", x, y)
			}
		}
	}
}

func (p *windowImpl) ScrollTo(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("scrollTo")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.Call("scrollTo", toJSONObject(options))
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.Call("scrollTo", x, y)
			}
		}
	}
}

func (p *windowImpl) ScrollBy(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("scrollBy")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.Call("scrollBy", toJSONObject(options))
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.Call("scrollBy", x, y)
			}
		}
	}
}

func (p *windowImpl) ScreenX() int {
	return p.Get("screenX").Int()
}

func (p *windowImpl) ScreenY() int {
	return p.Get("screenY").Int()
}

func (p *windowImpl) OuterWidth() int {
	return p.Get("outerWidth").Int()
}

func (p *windowImpl) OuterHeight() int {
	return p.Get("outerHeight").Int()
}

func (p *windowImpl) DevicePixelRatio() float64 {
	return p.Get("devicePixelRatio").Float()
}

func (p *windowImpl) JSValue() js.Value {
	return p.Value
}

// -------------8<---------------------------------------

type barPropImpl struct {
	js.Value
}

func newBarProp(v js.Value) BarProp {
	if isNil(v) {
		return nil
	}

	return &barPropImpl{
		Value: v,
	}
}

func (p *barPropImpl) Visible() bool {
	return p.Get("visible").Bool()
}

// -------------8<---------------------------------------

type locationImpl struct {
	*workerLocationImpl
	js.Value
}

func newLocation(v js.Value) Location {
	if isNil(v) {
		return nil
	}
	return &locationImpl{
		workerLocationImpl: newWorkerLocationImpl(v),
		Value:              v,
	}
}

func (p *locationImpl) SetHref(href string) {
	p.Set("href", href)
}

func (p *locationImpl) SetProtocol(protocol string) {
	p.Set("protocol", protocol)
}

func (p *locationImpl) SetHost(host string) {
	p.Set("host", host)
}

func (p *locationImpl) SetHostname(hostname string) {
	p.Set("hostname", hostname)
}

func (p *locationImpl) SetPort(port string) {
	p.Set("port", port)
}

func (p *locationImpl) SetPathname(pathname string) {
	p.Set("pathname", pathname)
}

func (p *locationImpl) SetSearch(search string) {
	p.Set("search", search)
}

func (p *locationImpl) SetHash(hash string) {
	p.Set("hash", hash)
}

func (p *locationImpl) Assign(url string) {
	p.Call("assign", url)
}

func (p *locationImpl) Replace(url string) {
	p.Call("replace", url)
}

func (p *locationImpl) Reload() {
	p.Call("reload")
}

func (p *locationImpl) AncestorOrigins() []string {
	return stringSequenceToSlice(p.Get("ancestorOrigins"))
}

// -------------8<---------------------------------------

type historyImpl struct {
	js.Value
}

func newHistory(v js.Value) History {
	if isNil(v) {
		return nil
	}

	return &historyImpl{
		Value: v,
	}
}

func (p *historyImpl) Length() int {
	return p.Get("length").Int()
}

func (p *historyImpl) ScrollRestoration() ScrollRestorationType {
	return ScrollRestorationType(p.Get("scrollRestoration").String())
}

func (p *historyImpl) SetScrollRestoration(sr ScrollRestorationType) {
	p.Set("scrollRestoration", sr)
}

func (p *historyImpl) State() interface{} {
	return Wrap(p.Get("state"))
}

func (p *historyImpl) Go(delta ...int) {
	switch len(delta) {
	case 0:
		p.Call("go")
	default:
		p.Call("go", delta[0])
	}
}

func (p *historyImpl) Back() {
	p.Call("back")
}

func (p *historyImpl) Forward() {
	p.Call("forward")
}

func (p *historyImpl) PushState(data interface{}, title string, url ...string) {
	switch len(url) {
	case 0:
		p.Call("pushState", valueOf(data), title)
	default:
		p.Call("pushState", valueOf(data), title, url[0])
	}
}

func (p *historyImpl) ReplaceState(data interface{}, title string, url ...string) {
	switch len(url) {
	case 0:
		p.Call("replaceState", valueOf(data), title)
	default:
		p.Call("replaceState", valueOf(data), title, url[0])
	}
}

// -------------8<---------------------------------------

type popStateEventImpl struct {
	*eventImpl
}

func newPopStateEvent(v js.Value) PopStateEvent {
	if isNil(v) {
		return nil
	}

	return &popStateEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *popStateEventImpl) State() interface{} {
	return Wrap(p.Get("state"))
}

func (p *popStateEventImpl) SetState(state interface{}) {
	p.Set("state", state)
}

// -------------8<---------------------------------------

type hashChangeEventImpl struct {
	*eventImpl
}

func newHashChangeEvent(v js.Value) HashChangeEvent {
	if isNil(v) {
		return nil
	}

	return &hashChangeEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *hashChangeEventImpl) OldURL() string {
	return p.Get("oldURL").String()
}

func (p *hashChangeEventImpl) NewURL() string {
	return p.Get("newURL").String()
}

// -------------8<---------------------------------------

type pageTransitionEventImpl struct {
	*eventImpl
}

func newPageTransitionEvent(v js.Value) PageTransitionEvent {
	if isNil(v) {
		return nil
	}

	return &pageTransitionEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *pageTransitionEventImpl) Persisted() bool {
	return p.Get("persisted").Bool()
}

// -------------8<---------------------------------------

type windowProxyImpl struct {
	*windowImpl
}

func newWindowProxy(v js.Value) WindowProxy {
	if isNil(v) {
		return nil
	}

	return &windowProxyImpl{
		windowImpl: newWindowImpl(v),
	}
}

// -------------8<---------------------------------------

type navigatorOnLineImpl struct {
	js.Value
}

func newNavigatorOnLine(v js.Value) NavigatorOnLine {
	if p := newNavigatorOnLineImpl(v); p != nil {
		return p
	}
	return nil
}

func newNavigatorOnLineImpl(v js.Value) *navigatorOnLineImpl {
	if isNil(v) {
		return nil
	}

	return &navigatorOnLineImpl{
		Value: v,
	}
}

func (p *navigatorOnLineImpl) OnLine() bool {
	return p.Get("onLine").Bool()
}

// -------------8<---------------------------------------

type beforeUnloadEventImpl struct {
	*eventImpl
}

func newBeforeUnloadEvent(v js.Value) BeforeUnloadEvent {
	if isNil(v) {
		return nil
	}

	return &beforeUnloadEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *beforeUnloadEventImpl) ReturnValue() string {
	return p.Get("returnValue").String()
}

func (p *beforeUnloadEventImpl) SetReturnValue(retVal string) {
	p.Set("returnValue", retVal)
}

// -------------8<---------------------------------------
