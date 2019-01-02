// +build js,wasm

package main

/*
Original version

https://github.com/fewstera/go-wasm-ball-drop-game

*/

import (
	"fmt"
	"github.com/sternix/wasm"
	"math"
	"math/rand"
	"time"
)

const (
	holeSize             = 120
	distanceBetweenLines = 80
	playerRadius         = 20
	playerColor          = "red"
	lineColor            = "green"
)

var (
	width    float64
	height   float64
	doc      wasm.Document
	ctx      wasm.CanvasRenderingContext2D
	lines    []*line
	playerX  float64
	playerY  float64
	gameOver bool
	score    int
)

type line struct {
	// Y position of the line
	Y float64
	// Where the Hole in the line starts (x position)
	HoleStart float64
	// The width the Hole
	HoleWidth float64
}

func main() {
	done := make(chan struct{}, 0)

	wasm.CurrentWindow().On("beforeunload", func(wasm.Event) {
		done <- struct{}{}
	})

	doc = wasm.CurrentDocument()
	canvasEl := doc.ElementById("canvas").(wasm.HTMLCanvasElement)

	w := canvasEl.ClientWidth()
	h := canvasEl.ClientHeight()

	width = float64(w)
	height = float64(h)

	canvasEl.SetWidth(w)
	canvasEl.SetHeight(h)

	ctx = canvasEl.Context2D()
	playerX = width / 2
	playerY = height / 2

	eh := doc.OnKeyDown(keyDownHandler)

	go updateGame(done)
	go updateCanvas()

	fmt.Println("Starting")
	fmt.Printf("Width: %v, height: %v\n", width, height)

	<-done
	eh.Remove()
	gameOver = true
	updateCanvas()
}

func updateCanvas() {
	ctx.ClearRect(0.0, 0.0, width, height)

	for _, line := range lines {
		ctx.BeginPath()
		ctx.MoveTo(0.0, line.Y)
		// Draw first line segment
		ctx.LineTo(line.HoleStart, line.Y)
		// Skip hole
		ctx.MoveTo(line.HoleStart+line.HoleWidth, line.Y)
		// Draw end of line
		ctx.LineTo(width, line.Y)
		ctx.Stroke()
		ctx.ClosePath()
	}

	// Draw player circle
	ctx.BeginPath()
	ctx.SetFillStyle(playerColor)
	ctx.Arc(playerX, playerY, playerRadius, 0, math.Pi*2, true)
	ctx.ClosePath()
	ctx.Fill()

	if !gameOver {
		// Draw score
		ctx.SetFont("30px Helvetica bold")
		ctx.SetFillStyle(playerColor)
		ctx.FillText(fmt.Sprintf("%d", score), 20, 50)

		time.Sleep(time.Millisecond)
		updateCanvas()
	} else {
		finalTextLineOne := "YOU LOST."
		finalTextLineTwo := fmt.Sprintf("Final score: %d", score)

		centerX := width / 2
		centerY := height / 2

		ctx.SetFont("30px Helvetica bold")
		lineOneWidth := ctx.MeasureText(finalTextLineOne).Width()
		lineTwoWidth := ctx.MeasureText(finalTextLineTwo).Width()

		ctx.BeginPath()
		ctx.Rect(centerX-(lineTwoWidth/2)-20, centerY-80, lineTwoWidth+40, 160)
		ctx.SetFillStyle("grey")
		ctx.Fill()
		ctx.Stroke()
		ctx.ClosePath()

		ctx.BeginPath()
		ctx.SetFillStyle("black")
		ctx.FillText(finalTextLineOne, centerX-lineOneWidth/2, centerY-30)
		ctx.FillText(finalTextLineTwo, centerX-lineTwoWidth/2, centerY+30)
		ctx.ClosePath()
	}
}

func updateGame(done chan struct{}) {
	score++
	newestLinePosition := 0.0
	var closestLineToPlayer *line
	for _, line := range lines {
		line.Y = line.Y - 1
		newestLinePosition = line.Y

		if line.Y+5+playerRadius >= playerY && closestLineToPlayer == nil {
			closestLineToPlayer = line
		}
	}

	shouldAddNewLine := (height-newestLinePosition > distanceBetweenLines)
	if shouldAddNewLine {
		holePosition := float64(rand.Intn(int(width - holeSize)))
		lines = append(lines, &line{height, holePosition, holeSize})
	}

	shouldDeleteFirstLine := (lines[0].Y < 0)
	if shouldDeleteFirstLine {
		lines = lines[1:]
	}

	if closestLineToPlayer != nil {
		playerOnLine := math.Abs(float64(playerY-closestLineToPlayer.Y+playerRadius)) < 5

		holeMinX := closestLineToPlayer.HoleStart + playerRadius
		holeMaxX := closestLineToPlayer.HoleStart + closestLineToPlayer.HoleWidth - playerRadius
		playerIsInHole := playerX > holeMinX && playerX < holeMaxX

		if playerOnLine && !playerIsInHole {
			playerY = closestLineToPlayer.Y - playerRadius
		} else {
			playerY = playerY + 2
		}
	}

	if playerY+playerRadius > height {
		playerY = height - playerRadius
	}

	if playerY-playerRadius <= 0 {
		done <- struct{}{}
	} else {
		time.Sleep(time.Duration(10) * time.Millisecond)
		updateGame(done)
	}
}

func keyDownHandler(e wasm.KeyboardEvent) {
	switch e.Key() {
	case wasm.KeyArrowLeft, "a", "A", "4":
		playerX = playerX - 25
		if playerX < playerRadius {
			playerX = playerRadius
		}
	case wasm.KeyArrowRight, "d", "D", "6":
		playerX = playerX + 25
		if playerX+playerRadius-width >= 0 {
			playerX = width - playerRadius
		}
	}
}
