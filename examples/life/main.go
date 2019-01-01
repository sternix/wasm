// +build js,wasm

package main

/*
https://github.com/maded2/life
rewritten with wasm lib
*/

import (
	"fmt"
	"github.com/sternix/wasm"
	"math"
	"math/rand"
	"strings"
	"time"
)

var (
	gen = false
	win = wasm.CurrentWindow()
	doc = wasm.CurrentDocument()

	context wasm.CanvasRenderingContext2D
)

func main() {
	h, w := 80, 200

	win.OnResize(func(wasm.Event) {
		fmt.Printf("w=%d, h=%d\n", win.InnerWidth(), win.InnerHeight())
		gen = true
	})

	setSize(w, h)

	space2 := make([][]int8, h)
	for y := 0; y < h; y++ {
		space2[y] = make([]int8, w)
	}

	rand.Seed(time.Now().UnixNano())

	currentSpace := genSpace(h, w)
	nextSpace := space2

	t := time.Now().Add(time.Minute)
	for {
		//printSpace(currentSpace, h, w)
		drawSpace(currentSpace, h, w)
		runSpace(currentSpace, nextSpace, h, w)
		currentSpace, nextSpace = nextSpace, currentSpace
		time.Sleep(time.Millisecond * 50)
		if time.Now().After(t) || gen {
			t = time.Now().Add(time.Minute)
			currentSpace = genSpace(h, w)
			gen = false
		}
	}
}

func setSize(w, h int) {
	canvas := doc.QuerySelector("canvas").(wasm.HTMLCanvasElement)
	canvas.SetWidth(w * 8)
	canvas.SetHeight(h * 8)
	context = canvas.Context2D()
	clearScreen()
}

func clearScreen() {
	context.ClearRect(0, 0, float64(win.InnerWidth()), float64(win.InnerHeight()))
}

func drawCell(x, y float64) {
	context.BeginPath()
	context.Arc(x*8, y*8, 4.0, 0, math.Pi*2)
	context.Stroke()
}

func genSpace(h, w int) [][]int8 {
	space1 := make([][]int8, h)
	for y := 0; y < h; y++ {
		space1[y] = make([]int8, w)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			space1[y][x] = int8(rand.Intn(2))
		}
	}
	return space1
}

func runSpace(cs, ns [][]int8, h, w int) {
	ns[0] = make([]int8, w)
	ns[h-1] = make([]int8, w)
	for y := 1; y < h-1; y++ {
		ns[y] = make([]int8, w)
		for x := 1; x < w-1; x++ {
			currentCell := cs[y][x]
			neighbors := cs[y-1][x-1] + cs[y-1][x] + cs[y-1][x+1] + cs[y][x-1] + cs[y][x+1] + cs[y+1][x-1] + cs[y+1][x] + cs[y+1][x+1]
			if currentCell == 0 {
				if neighbors == 3 {
					ns[y][x] = 1
				}
			} else {
				if neighbors < 2 || neighbors > 3 {
					ns[y][x] = 0
				} else {
					ns[y][x] = 1
				}
			}
		}
	}
}

func printSpace(space [][]int8, h, w int) {
	sb := strings.Builder{}
	sb.WriteString("\033[H")
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if space[y][x] == 1 {
				sb.WriteString("#")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	fmt.Print(sb.String())
}

func drawSpace(space [][]int8, h, w int) {
	clearScreen()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if space[y][x] == 1 {
				drawCell(float64(x), float64(y))
			}
		}
	}
}
