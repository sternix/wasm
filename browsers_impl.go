// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewPopStateEvent(typ string, p ...PopStateEventInit) PopStateEvent {
	jsPopStateEvent := jsGlobal.get("PopStateEvent")
	if jsPopStateEvent.valid() {
		if len(p) > 0 {
			return wrapPopStateEvent(jsPopStateEvent.jsNew(typ, p[0].toJSObject()))
		}

		return wrapPopStateEvent(jsPopStateEvent.jsNew(typ))
	}
	return nil

}

func NewHashChangeEvent(typ string, p ...HashChangeEventInit) HashChangeEvent {
	if jsHashChangeEvent := jsGlobal.get("HashChangeEvent"); jsHashChangeEvent.valid() {
		switch len(p) {
		case 0:
			return wrapHashChangeEvent(jsHashChangeEvent.jsNew(typ))
		default:
			return wrapHashChangeEvent(jsHashChangeEvent.jsNew(typ, p[0].toJSObject()))
		}
	}
	return nil
}

func NewPageTransitionEvent(typ string, p ...PageTransitionEventInit) PageTransitionEvent {
	if jsPageTransitionEvent := jsGlobal.get("PageTransitionEvent"); jsPageTransitionEvent.valid() {
		switch len(p) {
		case 0:
			return wrapPageTransitionEvent(jsPageTransitionEvent.jsNew(typ))
		default:
			return wrapPageTransitionEvent(jsPageTransitionEvent.jsNew(typ, p[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type windowImpl struct {
	*eventTargetImpl
	*windowOrWorkerGlobalScopeImpl
	*globalEventHandlersImpl
	*windowEventHandlersImpl
	Value
}

func wrapWindow(v Value) Window {
	if p := newWindowImpl(v); p != nil {
		return p
	}
	return nil
}

func newWindowImpl(v Value) *windowImpl {
	if v.valid() {
		wi := &windowImpl{
			eventTargetImpl:               newEventTargetImpl(v),
			windowOrWorkerGlobalScopeImpl: newWindowOrWorkerGlobalScopeImpl(v),
			Value:                         v,
		}
		wi.globalEventHandlersImpl = newGlobalEventHandlersImpl(wi.eventTargetImpl)
		wi.windowEventHandlersImpl = newWindowEventHandlersImpl(wi.eventTargetImpl)
		return wi
	}
	return nil
}

func (p *windowImpl) Console() Console {
	return wrapConsole(p.get("console"))
}

func (p *windowImpl) Window() WindowProxy {
	return wrapWindowProxy(p.get("window"))
}

func (p *windowImpl) Self() WindowProxy {
	return wrapWindowProxy(p.get("self"))
}

func (p *windowImpl) Document() Document {
	return wrapDocument(p.get("document"))
}

func (p *windowImpl) Name() string {
	return p.get("name").toString()
}

func (p *windowImpl) SetName(name string) {
	p.set("name", name)
}

func (p *windowImpl) Location() Location {
	return wrapLocation(p.get("location"))
}

func (p *windowImpl) History() History {
	return wrapHistory(p.get("history"))
}

func (p *windowImpl) Locationbar() BarProp {
	return wrapBarProp(p.get("locationbar"))
}

func (p *windowImpl) Menubar() BarProp {
	return wrapBarProp(p.get("menubar"))
}

func (p *windowImpl) Personalbar() BarProp {
	return wrapBarProp(p.get("personalbar"))
}

func (p *windowImpl) Scrollbars() BarProp {
	return wrapBarProp(p.get("scrollbars"))
}

func (p *windowImpl) Statusbar() BarProp {
	return wrapBarProp(p.get("statusbar"))
}

func (p *windowImpl) Toolbar() BarProp {
	return wrapBarProp(p.get("toolbar"))
}

func (p *windowImpl) Status() string {
	return p.get("status").toString()
}

func (p *windowImpl) SetStatus(status string) {
	p.set("status", status)
}

func (p *windowImpl) Close() {
	p.call("close")
}

func (p *windowImpl) Closed() bool {
	return p.get("closed").toBool()
}

func (p *windowImpl) Stop() {
	p.call("stop")
}

func (p *windowImpl) Focus() {
	p.call("focus")
}

func (p *windowImpl) Blur() {
	p.call("blur")
}

func (p *windowImpl) Frames() WindowProxy {
	return wrapWindowProxy(p.get("frames"))
}

func (p *windowImpl) Length() int {
	return p.get("length").toInt()
}

func (p *windowImpl) Top() WindowProxy {
	return wrapWindowProxy(p.get("top"))
}

func (p *windowImpl) Opener() WindowProxy {
	return wrapWindowProxy(p.get("opener"))
}

func (p *windowImpl) Parent() WindowProxy {
	return wrapWindowProxy(p.get("parent"))
}

func (p *windowImpl) FrameElement() Element {
	return wrapAsElement(p.get("frameElement"))
}

func (p *windowImpl) Open(args ...interface{}) WindowProxy {
	switch len(args) {
	case 1:
		if url, ok := args[0].(string); ok {
			return wrapWindowProxy(p.call("open", url))
		}
	case 2:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				return wrapWindowProxy(p.call("open", url, target))
			}
		}
	case 3:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				if features, ok := args[2].(string); ok {
					return wrapWindowProxy(p.call("open", url, target, features))
				}
			}
		}
	case 4:
		if url, ok := args[0].(string); ok {
			if target, ok := args[1].(string); ok {
				if features, ok := args[2].(string); ok {
					if replace, ok := args[3].(bool); ok {
						return wrapWindowProxy(p.call("open", url, target, features, replace))
					}
				}
			}
		}
	}

	return wrapWindowProxy(p.call("open"))
}

func (p *windowImpl) Navigator() Navigator {
	return wrapNavigator(p.get("navigator"))
}

func (p *windowImpl) Alert(msg ...string) {
	switch len(msg) {
	case 0:
		p.call("alert")
	default:
		p.call("alert", msg[0])
	}
}

func (p *windowImpl) Confirm(msg ...string) bool {
	switch len(msg) {
	case 0:
		return p.call("confirm").toBool()
	default:
		return p.call("confirm", msg[0]).toBool()
	}
}

func (p *windowImpl) Prompt(args ...string) string {
	switch len(args) {
	case 0:
		return p.call("prompt").toString()
	case 1:
		return p.call("prompt", args[0]).toString() // message
	default:
		return p.call("prompt", args[0], args[1]).toString() // message, default
	}
}

func (p *windowImpl) Print() {
	p.call("print")
}

func (p *windowImpl) RequestAnimationFrame(cb FrameRequestCallback) int {
	return p.call("requestAnimationFrame", cb.jsCallback()).toInt()
}

func (p *windowImpl) CancelAnimationFrame(handle int) {
	p.call("cancelAnimationFrame", handle)
}

func (p *windowImpl) MatchMedia(query string) MediaQueryList {
	return wrapMediaQueryList(p.call("matchMedia", query))
}

func (p *windowImpl) Screen() Screen {
	return wrapScreen(p.get("screen"))
}

func (p *windowImpl) MoveTo(x int, y int) {
	p.call("moveTo", x, y)
}

func (p *windowImpl) MoveBy(x int, y int) {
	p.call("moveBy", x, y)
}

func (p *windowImpl) ResizeTo(x int, y int) {
	p.call("resizeTo", x, y)
}

func (p *windowImpl) ResizeBy(x int, y int) {
	p.call("resizeBy", x, y)
}

func (p *windowImpl) InnerWidth() int {
	return p.get("innerWidth").toInt()
}

func (p *windowImpl) InnerHeight() int {
	return p.get("innerHeight").toInt()
}

func (p *windowImpl) ScrollX() float64 {
	return p.get("scrollX").toFloat64()
}

func (p *windowImpl) PageXOffset() float64 {
	return p.get("pageXOffset").toFloat64()
}

func (p *windowImpl) ScrollY() float64 {
	return p.get("scrollY").toFloat64()
}

func (p *windowImpl) PageYOffset() float64 {
	return p.get("pageYOffset").toFloat64()
}

func (p *windowImpl) Scroll(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("scroll")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.call("scroll", options.toJSObject())
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.call("scroll", x, y)
			}
		}
	}
}

func (p *windowImpl) ScrollTo(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("scrollTo")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.call("scrollTo", options.toJSObject())
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.call("scrollTo", x, y)
			}
		}
	}
}

func (p *windowImpl) ScrollBy(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("scrollBy")
	case 1:
		if options, ok := args[0].(ScrollToOptions); ok {
			p.call("scrollBy", options.toJSObject())
		}
	case 2:
		if x, ok := args[0].(float64); ok {
			if y, ok := args[1].(float64); ok {
				p.call("scrollBy", x, y)
			}
		}
	}
}

func (p *windowImpl) ScreenX() int {
	return p.get("screenX").toInt()
}

func (p *windowImpl) ScreenLeft() int {
	return p.get("screenLeft").toInt()
}

func (p *windowImpl) ScreenY() int {
	return p.get("screenY").toInt()
}

func (p *windowImpl) ScreenTop() int {
	return p.get("screenTop").toInt()
}

func (p *windowImpl) OuterWidth() int {
	return p.get("outerWidth").toInt()
}

func (p *windowImpl) OuterHeight() int {
	return p.get("outerHeight").toInt()
}

func (p *windowImpl) DevicePixelRatio() float64 {
	return p.get("devicePixelRatio").toFloat64()
}

func (p *windowImpl) ComputedStyle(Element, ...string) CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(p.call("getComputedStyle"))
}

func (p *windowImpl) PseudoElements(elt Element, typ string) []CSSPseudoElement {
	l := wrapCSSPseudoElementList(p.call("getPseudoElements", JSValue(elt), typ))
	if l != nil && l.Length() > 0 {
		ret := make([]CSSPseudoElement, l.Length())
		for i := range ret {
			ret[i] = l.Item(i)
		}
		return ret
	}
	return nil
}

// -------------8<---------------------------------------

type barPropImpl struct {
	Value
}

func wrapBarProp(v Value) BarProp {
	if v.valid() {
		return &barPropImpl{
			Value: v,
		}
	}
	return nil
}

func (p *barPropImpl) Visible() bool {
	return p.get("visible").toBool()
}

// -------------8<---------------------------------------

type locationImpl struct {
	*workerLocationImpl
}

func wrapLocation(v Value) Location {
	if v.valid() {
		return &locationImpl{
			workerLocationImpl: newWorkerLocationImpl(v),
		}
	}
	return nil
}

func (p *locationImpl) SetHref(href string) {
	p.set("href", href)
}

func (p *locationImpl) SetProtocol(protocol string) {
	p.set("protocol", protocol)
}

func (p *locationImpl) SetHost(host string) {
	p.set("host", host)
}

func (p *locationImpl) SetHostname(hostname string) {
	p.set("hostname", hostname)
}

func (p *locationImpl) SetPort(port string) {
	p.set("port", port)
}

func (p *locationImpl) SetPathname(pathname string) {
	p.set("pathname", pathname)
}

func (p *locationImpl) SetSearch(search string) {
	p.set("search", search)
}

func (p *locationImpl) SetHash(hash string) {
	p.set("hash", hash)
}

func (p *locationImpl) Assign(url string) {
	p.call("assign", url)
}

func (p *locationImpl) Replace(url string) {
	p.call("replace", url)
}

func (p *locationImpl) Reload() {
	p.call("reload")
}

func (p *locationImpl) AncestorOrigins() []string {
	return stringSequenceToSlice(p.get("ancestorOrigins"))
}

// -------------8<---------------------------------------

type historyImpl struct {
	Value
}

func wrapHistory(v Value) History {
	if v.valid() {
		return &historyImpl{
			Value: v,
		}
	}
	return nil
}

func (p *historyImpl) Length() int {
	return p.get("length").toInt()
}

func (p *historyImpl) ScrollRestoration() ScrollRestorationType {
	return ScrollRestorationType(p.get("scrollRestoration").toString())
}

func (p *historyImpl) SetScrollRestoration(sr ScrollRestorationType) {
	p.set("scrollRestoration", sr)
}

func (p *historyImpl) State() interface{} {
	return Wrap(p.get("state"))
}

func (p *historyImpl) Go(delta ...int) {
	switch len(delta) {
	case 0:
		p.call("go")
	default:
		p.call("go", delta[0])
	}
}

func (p *historyImpl) Back() {
	p.call("back")
}

func (p *historyImpl) Forward() {
	p.call("forward")
}

func (p *historyImpl) PushState(data interface{}, title string, url ...string) {
	switch len(url) {
	case 0:
		p.call("pushState", data, title)
	default:
		p.call("pushState", data, title, url[0])
	}
}

func (p *historyImpl) ReplaceState(data interface{}, title string, url ...string) {
	switch len(url) {
	case 0:
		p.call("replaceState", data, title)
	default:
		p.call("replaceState", data, title, url[0])
	}
}

// -------------8<---------------------------------------

type popStateEventImpl struct {
	*eventImpl
}

func wrapPopStateEvent(v Value) PopStateEvent {
	if !v.valid() {
		return nil
	}

	return &popStateEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *popStateEventImpl) State() interface{} {
	return Wrap(p.get("state"))
}

func (p *popStateEventImpl) SetState(state interface{}) {
	p.set("state", state)
}

// -------------8<---------------------------------------

type hashChangeEventImpl struct {
	*eventImpl
}

func wrapHashChangeEvent(v Value) HashChangeEvent {
	if v.valid() {
		return &hashChangeEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *hashChangeEventImpl) OldURL() string {
	return p.get("oldURL").toString()
}

func (p *hashChangeEventImpl) NewURL() string {
	return p.get("newURL").toString()
}

// -------------8<---------------------------------------

type pageTransitionEventImpl struct {
	*eventImpl
}

func wrapPageTransitionEvent(v Value) PageTransitionEvent {
	if v.valid() {
		return &pageTransitionEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *pageTransitionEventImpl) Persisted() bool {
	return p.get("persisted").toBool()
}

// -------------8<---------------------------------------

type windowProxyImpl struct {
	*windowImpl
}

func wrapWindowProxy(v Value) WindowProxy {
	if v.valid() {
		return &windowProxyImpl{
			windowImpl: newWindowImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type navigatorOnLineImpl struct {
	Value
}

func wrapNavigatorOnLine(v Value) NavigatorOnLine {
	if p := newNavigatorOnLineImpl(v); p != nil {
		return p
	}
	return nil
}

func newNavigatorOnLineImpl(v Value) *navigatorOnLineImpl {
	if v.valid() {
		return &navigatorOnLineImpl{
			Value: v,
		}
	}
	return nil
}

func (p *navigatorOnLineImpl) OnLine() bool {
	return p.get("onLine").toBool()
}

// -------------8<---------------------------------------

type beforeUnloadEventImpl struct {
	*eventImpl
}

func wrapBeforeUnloadEvent(v Value) BeforeUnloadEvent {
	if v.valid() {
		return &beforeUnloadEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *beforeUnloadEventImpl) ReturnValue() string {
	return p.get("returnValue").toString()
}

func (p *beforeUnloadEventImpl) SetReturnValue(retVal string) {
	p.set("returnValue", retVal)
}

// -------------8<---------------------------------------
