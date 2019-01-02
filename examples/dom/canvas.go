// +build js,wasm

package main

import (
	"github.com/sternix/wasm"
	"math"
)

func TestCanvas(doc wasm.Document) {
	// examples taken from
	// https://www.w3schools.com/html/html5_canvas.asp
	line := wasm.NewHTMLCanvasElement(100, 100)
	ctx := line.Context2D()
	ctx.MoveTo(0, 0)
	ctx.LineTo(100, 100)
	ctx.Stroke()
	doc.Body().AppendChild(line)

	circle := wasm.NewHTMLCanvasElement(200, 100)
	ctx = circle.Context2D()
	ctx.BeginPath()
	ctx.Arc(95, 50, 40, 0, 2*math.Pi)
	ctx.Stroke()
	doc.Body().AppendChild(circle)

	text := wasm.NewHTMLCanvasElement(200, 100)
	ctx = text.Context2D()
	ctx.SetFont("30px Arial")
	ctx.FillText("Hello World", 10, 50)
	doc.Body().AppendChild(text)

	stext := wasm.NewHTMLCanvasElement(200, 100)
	ctx = stext.Context2D()
	ctx.SetFont("30px Arial")
	ctx.StrokeText("Hello World", 10, 50)
	doc.Body().AppendChild(stext)

	gradient := wasm.NewHTMLCanvasElement(200, 100)
	ctx = gradient.Context2D()
	grd := ctx.CreateLinearGradient(0, 0, 200, 0)
	grd.AddColorStop(0, "red")
	grd.AddColorStop(1, "white")
	ctx.SetFillStyle(grd)
	ctx.FillRect(10, 10, 150, 80)
	doc.Body().AppendChild(gradient)

	cgradient := wasm.NewHTMLCanvasElement(200, 100)
	ctx = cgradient.Context2D()
	grd = ctx.CreateRadialGradient(75, 50, 5, 90, 60, 100)
	grd.AddColorStop(0, "red")
	grd.AddColorStop(1, "white")
	ctx.SetFillStyle(grd)
	ctx.FillRect(10, 10, 150, 80)
	doc.Body().AppendChild(cgradient)
}
