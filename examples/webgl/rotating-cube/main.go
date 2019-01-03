// +build js,wasm

/*
Code forked from

https://github.com/bobcob7/wasm-rotating-cube

*/

package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sternix/wasm"
	"syscall/js"
	"unsafe"
)

//// BUFFERS + SHADERS ////
// Shamelessly copied from https://www.tutorialspoint.com/webgl/webgl_cube_rotation.htm //
var verticesNative = []float32{
	-1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1,
	-1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1, 1,
	-1, -1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1,
	1, -1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1,
	-1, -1, -1, -1, -1, 1, 1, -1, 1, 1, -1, -1,
	-1, 1, -1, -1, 1, 1, 1, 1, 1, 1, 1, -1,
}
var colorsNative = []float32{
	5, 3, 7, 5, 3, 7, 5, 3, 7, 5, 3, 7,
	1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3,
	0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0,
	1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0,
	0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0,
}
var indicesNative = []uint16{
	0, 1, 2, 0, 2, 3, 4, 5, 6, 4, 6, 7,
	8, 9, 10, 8, 10, 11, 12, 13, 14, 12, 14, 15,
	16, 17, 18, 16, 18, 19, 20, 21, 22, 20, 22, 23,
}

const vertShaderCode = `
	attribute vec3 position;
	uniform mat4 Pmatrix;
	uniform mat4 Vmatrix;
	uniform mat4 Mmatrix;
	attribute vec3 color;
	varying vec3 vColor;
	void main(void) {
		gl_Position = Pmatrix*Vmatrix*Mmatrix*vec4(position, 1.);
		vColor = color;
	}
	`

const fragShaderCode = `
	precision mediump float;
	varying vec3 vColor;
	void main(void) {
		gl_FragColor = vec4(vColor, 1.);
	}
	`

func main() {
	win := wasm.CurrentWindow()
	doc := wasm.CurrentDocument()
	canvasEl := doc.ElementById("gocanvas").(wasm.HTMLCanvasElement)
	width := doc.Body().ClientWidth()
	height := doc.Body().ClientHeight()
	canvasEl.SetWidth(width)
	canvasEl.SetHeight(height)

	gl := canvasEl.ContextWebGL()
	if gl == nil {
		win.Alert("Browser might support webgl")
		wasm.Exit()
	}

	// Convert buffers to JS TypedArrays
	var colors = js.TypedArrayOf(colorsNative)
	var vertices = js.TypedArrayOf(verticesNative)
	var indices = js.TypedArrayOf(indicesNative)

	// Create vertex buffer
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(wasm.ARRAY_BUFFER, vertexBuffer)
	gl.BufferDataSource(wasm.ARRAY_BUFFER, vertices, wasm.STATIC_DRAW)

	// Create color buffer
	colorBuffer := gl.CreateBuffer()
	gl.BindBuffer(wasm.ARRAY_BUFFER, colorBuffer)
	gl.BufferDataSource(wasm.ARRAY_BUFFER, colors, wasm.STATIC_DRAW)

	// Create index buffer
	indexBuffer := gl.CreateBuffer()
	gl.BindBuffer(wasm.ELEMENT_ARRAY_BUFFER, indexBuffer)
	gl.BufferDataSource(wasm.ELEMENT_ARRAY_BUFFER, indices, wasm.STATIC_DRAW)

	//// Shaders ////

	// Create a vertex shader object
	vertShader := gl.CreateShader(wasm.VERTEX_SHADER)
	gl.ShaderSource(vertShader, vertShaderCode)
	gl.CompileShader(vertShader)

	// Create fragment shader object
	fragShader := gl.CreateShader(wasm.FRAGMENT_SHADER)
	gl.ShaderSource(fragShader, fragShaderCode)
	gl.CompileShader(fragShader)

	// Create a shader program object to store
	// the combined shader program
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertShader)
	gl.AttachShader(shaderProgram, fragShader)
	gl.LinkProgram(shaderProgram)

	// Associate attributes to vertex shader
	PositionMatrix := gl.GetUniformLocation(shaderProgram, "Pmatrix")
	ViewMatrix := gl.GetUniformLocation(shaderProgram, "Vmatrix")
	ModelMatrix := gl.GetUniformLocation(shaderProgram, "Mmatrix")

	gl.BindBuffer(wasm.ARRAY_BUFFER, vertexBuffer)
	position := uint(gl.GetAttribLocation(shaderProgram, "position"))
	gl.VertexAttribPointer(position, 3, wasm.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(position)

	gl.BindBuffer(wasm.ARRAY_BUFFER, colorBuffer)
	color := uint(gl.GetAttribLocation(shaderProgram, "color"))
	gl.VertexAttribPointer(color, 3, wasm.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(color)

	gl.UseProgram(shaderProgram)

	// Set WeebGL properties
	gl.ClearColor(0.5, 0.5, 0.5, 0.9) // Color the screen is cleared to
	gl.ClearDepth(1.0)                // Z value that is set to the Depth buffer every frame
	gl.Viewport(0, 0, width, height)  // Viewport size
	gl.DepthFunc(wasm.LEQUAL)

	//// Create Matrixes ////
	ratio := float32(width) / float32(height)

	// Generate and apply projection matrix
	projMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), ratio, 1, 100.0)
	var projMatrixBuffer *[16]float32
	projMatrixBuffer = (*[16]float32)(unsafe.Pointer(&projMatrix))
	gl.UniformMatrix4fv(PositionMatrix, false, []float32((*projMatrixBuffer)[:]))

	// Generate and apply view matrix
	viewMatrix := mgl32.LookAtV(mgl32.Vec3{3.0, 3.0, 3.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 1.0, 0.0})
	var viewMatrixBuffer *[16]float32
	viewMatrixBuffer = (*[16]float32)(unsafe.Pointer(&viewMatrix))
	gl.UniformMatrix4fv(ViewMatrix, false, []float32((*viewMatrixBuffer)[:]))

	//// Drawing the Cube ////
	movMatrix := mgl32.Ident4()
	//Xvar renderFrame js.Callback
	var tmark float32
	var rotation = float32(0)

	// Bind to element array for draw function
	gl.BindBuffer(wasm.ELEMENT_ARRAY_BUFFER, indexBuffer)

	var rafcb int
	renderFrame := func(cb wasm.FrameRequestCallback, t float64) {
		// Calculate rotation rate
		now := float32(t)
		tdiff := now - tmark
		tmark = now
		rotation = rotation + float32(tdiff)/500

		// Do new model matrix calculations
		movMatrix = mgl32.HomogRotate3DX(0.5 * rotation)
		movMatrix = movMatrix.Mul4(mgl32.HomogRotate3DY(0.3 * rotation))
		movMatrix = movMatrix.Mul4(mgl32.HomogRotate3DZ(0.2 * rotation))

		// Convert model matrix to a JS TypedArray
		var modelMatrixBuffer *[16]float32
		modelMatrixBuffer = (*[16]float32)(unsafe.Pointer(&movMatrix))

		// Apply the model matrix
		gl.UniformMatrix4fv(ModelMatrix, false, []float32((*modelMatrixBuffer)[:]))

		// Clear the screen
		gl.Enable(wasm.DEPTH_TEST)
		gl.Clear(wasm.COLOR_BUFFER_BIT)
		gl.Clear(wasm.DEPTH_BUFFER_BIT)

		// Draw the cube
		gl.DrawElements(wasm.TRIANGLES, len(indicesNative), wasm.UNSIGNED_SHORT, 0)

		// Call next frame
		rafcb = win.RequestAnimationFrame(cb)
	}

	frcb := wasm.NewFrameRequestCallback(renderFrame)
	rafcb = win.RequestAnimationFrame(frcb)

	wasm.Loop()
	win.CancelAnimationFrame(rafcb)
	frcb.Release()
}
