// +build wasm,js

/*
https://www.w3schools.com/howto/howto_js_slideshow.asp
*/
package main

import (
	"github.com/sternix/wasm"
	"strings"
	"time"
)

var (
	slideIndex int
	doc        wasm.Document
	slides     []wasm.HTMLElement
	dots       []wasm.Element
)

func showSlides() {
	for i := 0; i < len(slides); i++ {
		slides[i].Style().SetProperty("display", "none")
	}
	slideIndex++
	if slideIndex > len(slides) {
		slideIndex = 1
	}

	for i := 0; i < len(dots); i++ {
		dots[i].SetClassName(strings.ReplaceAll(dots[i].ClassName(), " active", ""))
	}

	slides[slideIndex-1].Style().SetProperty("display", "block")
	dots[slideIndex-1].SetClassName(dots[slideIndex-1].ClassName() + " active")
}

func main() {
	doc = wasm.CurrentDocument()
	for _, e := range doc.ElementsByClassName("mySlides") {
		slides = append(slides, e.(wasm.HTMLElement))
	}

	dots = doc.ElementsByClassName("dot")

	showSlides()

	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for range ticker.C {
			showSlides()
		}
	}()

	wasm.Loop()
}
