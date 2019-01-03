// +build js,wasm

package main

import (
	"fmt"
	"github.com/sternix/wasm"
	"time"
)

func errx(msg string) {
	fmt.Printf("Wasm Error: %s\n", msg)
	wasm.Exit()
}

func main() {
	fmt.Println("Hello from WASM")
	win := wasm.CurrentWindow()
	if win == nil {
		errx("Window is NULL")
	}

	console := win.Console()
	if console == nil {
		errx("Window.Console is NULL")
	}

	TestConsole(console)

	doc := win.Document()
	if doc == nil {
		errx("Window.Document is NULL")
	}

	css := wasm.NewHTMLLinkElement()
	css.SetRel("stylesheet")
	css.SetType("text/css")
	css.SetHref("wasm.css") // http://XYZ/wasm.css
	doc.Head().AppendChild(css)

	TestDocument(doc)

	clock := wasm.NewHTMLDivElement()
	if clock == nil {
		errx("wasm.NewHTMLDivElement == NULL")
	}
	doc.Body().AppendChild(clock)

	var rafcb int
	fn := func(cb wasm.FrameRequestCallback, t float64) {
		clock.SetInnerHTML(time.Now().Format("15:04:05"))
		rafcb = win.RequestAnimationFrame(cb)
	}

	frcb := wasm.NewFrameRequestCallback(fn)

	rafcb = win.RequestAnimationFrame(frcb)

	TestWebSocket(doc)

	TestCanvas(doc)

	wp := win.Window()
	if wp == nil {
		errx("Window.Window is NULL")
	}

	self := win.Self()
	if self == nil {
		errx("Window.Self is NULL")
	}

	win.SetName("Wasm")

	fmt.Printf("Window.Name : %s\n", win.Name())
	if win.Name() != "Wasm" {
		errx("Window.SetName failed")
	}

	location := win.Location()
	if location == nil {
		errx("Window.Location is NULL")
	}
	TestLocation(location)

	history := win.History()
	if history == nil {
		errx("Window.History is NULL")
	}

	fmt.Printf("Window.History.Length: %d\n", history.Length())
	fmt.Printf("Window.History.ScrollRestoration %s\n", history.ScrollRestoration())
	fmt.Printf("Window.History.State: %v\n", history.State())

	locationBar := win.Locationbar()
	if locationBar == nil {
		errx("Window.LocationBar is NULL")
	}

	if locationBar.Visible() {
		fmt.Println("Window.Locationbar is Visible")
	} else {
		fmt.Println("Window.Locationbar is not Visible")
	}

	menuBar := win.Menubar()
	if menuBar == nil {
		errx("Window.Menubar is NULL")
	}

	if menuBar.Visible() {
		fmt.Println("Window.Menubar is Visible")
	} else {
		fmt.Println("Window.Menubar is not Visible")
	}

	personalBar := win.Personalbar()
	if personalBar == nil {
		errx("Window.Personalbar is NULL")
	}

	if personalBar.Visible() {
		fmt.Println("Window.Personalbar is Visible")
	} else {
		fmt.Println("Window.Personalbar is not Visible")
	}

	scrollbars := win.Scrollbars()
	if scrollbars == nil {
		errx("Window.Scrollbars is NULL")
	}

	if scrollbars.Visible() {
		fmt.Println("Window.Scrollbars is Visible")
	} else {
		fmt.Println("Window.Scrollbars is not Visible")
	}

	statusBar := win.Statusbar()
	if statusBar == nil {
		errx("Window.StatusBar is NULL")
	}

	if statusBar.Visible() {
		fmt.Println("Window.Statusbar is Visible")
	} else {
		fmt.Println("Window.Statusbar is not Visible")
	}

	toolBar := win.Toolbar()
	if toolBar == nil {
		errx("Window.Toolbar is NULL")
	}

	if toolBar.Visible() {
		fmt.Println("Window.Toolbar is Visible")
	} else {
		fmt.Println("Window.Toolbar is not Visible")
	}

	stStr := "This value is visible in StatusBar"
	win.SetStatus(stStr)
	fmt.Printf("Window.Status: %s\n", win.Status())
	if win.Status() != stStr {
		errx("Window.SetStatus failed")
	}

	if win.Closed() {
		errx("Window.Closed must be false")
	}

	win.Focus()

	if frame := win.Frames(); frame != nil {
		fmt.Printf("Window.Frame.Name: %s\n", frame.Name())
	}

	if win.FrameElement() != nil {
		errx("Window has not any frame element")
	}

	fmt.Printf("Window.Length: %d\n", win.Length())

	parent := win.Parent()
	if parent != nil {
		fmt.Printf("Window.Parent.Name: %s\n", parent.Name())
	}

	fmt.Printf("Window.InnerWidth: %d\n", win.InnerWidth())
	fmt.Printf("Window.InnerHeight: %d\n", win.InnerHeight())
	fmt.Printf("Window.ScrollX: %f\n", win.ScrollX())
	fmt.Printf("Window.ScrollY: %f\n", win.ScrollY())
	fmt.Printf("Window.PageXOffset: %f\n", win.PageXOffset())
	fmt.Printf("Window.PageYOffset: %f\n", win.PageYOffset())
	fmt.Printf("Window.ScreenX: %d\n", win.ScreenX())
	fmt.Printf("Window.ScreenY: %d\n", win.ScreenY())
	fmt.Printf("Window.OuterWidth: %d\n", win.OuterWidth())
	fmt.Printf("Window.OuterHeight: %d\n", win.OuterHeight())
	fmt.Printf("Window.DevicePixelRatio: %f\n", win.DevicePixelRatio())

	mql := win.MatchMedia("(min-width: 700px)")
	if mql.Matches() {
		fmt.Printf("MediaQueryList.Media: %s\n", mql.Media())
	} else {
		fmt.Println("Media does not matches")
	}

	mql = win.MatchMedia("(min-width: 9000px)")
	if mql.Matches() {
		fmt.Printf("MediaQueryList.Media: %s\n", mql.Media())
	} else {
		fmt.Println("Media does not matches")
	}

	navigator := win.Navigator()
	if navigator == nil {
		errx("Window.Navigator is NULL")
	}

	TestNavigator(navigator)

	screen := win.Screen()
	if screen == nil {
		errx("Window.Screen is NULL")
	}

	fmt.Printf("Screen.AvailWidth: %d\n", screen.AvailWidth())
	fmt.Printf("Screen.AvailHeight: %d\n", screen.AvailHeight())
	fmt.Printf("Screen.Width: %d\n", screen.Width())
	fmt.Printf("Screen.Height: %d\n", screen.Height())
	fmt.Printf("Screen.ColorDepth: %d\n", screen.ColorDepth())
	fmt.Printf("Screen.PixelDepth: %d\n", screen.PixelDepth())

	TestTabular()

	fmt.Println("All Tests completed")

	wasm.Loop()

	win.CancelAnimationFrame(rafcb)
	frcb.Release()
}

// skip for now
func TestWindowInteractive(win wasm.Window) {
	win.Alert("Hello World from WASM")

	if win.Confirm("Are you sure ?") {
		fmt.Println("Yes you sure")
	} else {
		fmt.Println("You arent sure")
	}

	res := win.Prompt("Message", "This is Default Value")
	fmt.Printf("You entered: %s\n", res)

	goWin := win.Open("https://golang.org/")
	if goWin.Opener() != nil {
		fmt.Printf("Window.Opener.Name: %s\n", goWin.Opener().Name())
	}

	win.Print()

	win.Blur()
}

func TestLocation(location wasm.Location) {
	fmt.Printf("Window.Location.Href: %s\n", location.Href())
	fmt.Printf("Window.Location.Origin: %s\n", location.Origin())
	fmt.Printf("Window.Location.Protocol: %s\n", location.Protocol())
	fmt.Printf("Window.Location.Host: %s\n", location.Host())
	fmt.Printf("Window.Location.Hostname: %s\n", location.Hostname())
	fmt.Printf("Window.Location.Port: %s\n", location.Port())
	fmt.Printf("Window.Location.Pathname: %s\n", location.Pathname())
	fmt.Printf("Window.Location.Search: %s\n", location.Search())
	fmt.Printf("Window.Location.Hash: %s\n", location.Hash())
}
