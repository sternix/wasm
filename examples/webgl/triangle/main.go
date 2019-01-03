// +build js,wasm

/*
Code forked from

https://github.com/bobcob7/wasm-basic-triangle
*/

package main

import (
	"github.com/sternix/wasm"
	"syscall/js"
)

var (
	width  int
	height int
	gl     wasm.WebGLRenderingContext
)

func main() {
	win := wasm.CurrentWindow()
	doc := wasm.CurrentDocument()
	canvasEl := doc.ElementById("gocanvas").(wasm.HTMLCanvasElement)
	width = doc.Body().ClientWidth()
	height = doc.Body().ClientHeight()
	canvasEl.SetWidth(width)
	canvasEl.SetHeight(height)

	gl = canvasEl.ContextWebGL()
	if gl == nil {
		win.Alert("Browser might support webgl")
		wasm.Exit()
	}

	var verticesNative = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	var vertices = js.TypedArrayOf(verticesNative)
	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(wasm.ARRAY_BUFFER, vertexBuffer)

	// Pass data to buffer
	gl.BufferDataSource(wasm.ARRAY_BUFFER, vertices, wasm.STATIC_DRAW)

	// Unbind buffer
	gl.BindBuffer(wasm.ARRAY_BUFFER, nil)

	//// INDEX BUFFER ////
	var indicesNative = []uint32{
		2, 1, 0,
	}
	var indices = js.TypedArrayOf(indicesNative)

	// Create buffer
	indexBuffer := gl.CreateBuffer()

	// Bind to buffer
	gl.BindBuffer(wasm.ELEMENT_ARRAY_BUFFER, indexBuffer)

	// Pass data to buffer
	gl.BufferDataSource(wasm.ELEMENT_ARRAY_BUFFER, indices, wasm.STATIC_DRAW)

	// Unbind buffer
	gl.BindBuffer(wasm.ELEMENT_ARRAY_BUFFER, nil)

	//// Shaders ////

	// Vertex shader source code
	vertCode := `
		attribute vec3 coordinates;
		  
		void main(void) {
			gl_Position = vec4(coordinates, 1.0);
		}`

	// Create a vertex shader object
	vertShader := gl.CreateShader(wasm.VERTEX_SHADER)

	// Attach vertex shader source code
	gl.ShaderSource(vertShader, vertCode)

	// Compile the vertex shader
	gl.CompileShader(vertShader)

	//fragment shader source code
	fragCode := `
		void main(void) {
			gl_FragColor = vec4(0.0, 0.0, 1.0, 1.0);
		}`

	// Create fragment shader object
	fragShader := gl.CreateShader(wasm.FRAGMENT_SHADER)

	// Attach fragment shader source code
	gl.ShaderSource(fragShader, fragCode)

	// Compile the fragmentt shader
	gl.CompileShader(fragShader)

	// Create a shader program object to store
	// the combined shader program
	shaderProgram := gl.CreateProgram()

	// Attach a vertex shader
	gl.AttachShader(shaderProgram, vertShader)

	// Attach a fragment shader
	gl.AttachShader(shaderProgram, fragShader)

	// Link both the programs
	gl.LinkProgram(shaderProgram)

	// Use the combined shader program object
	gl.UseProgram(shaderProgram)

	//// Associating shaders to buffer objects ////

	// Bind vertex buffer object
	gl.BindBuffer(wasm.ARRAY_BUFFER, vertexBuffer)

	// Bind index buffer object
	gl.BindBuffer(wasm.ELEMENT_ARRAY_BUFFER, indexBuffer)

	// Get the attribute location
	coord := gl.GetAttribLocation(shaderProgram, "coordinates")

	// Point an attribute to the currently bound VBO
	gl.VertexAttribPointer(uint(coord), 3, wasm.FLOAT, false, 0, 0)

	// Enable the attribute
	gl.EnableVertexAttribArray(uint(coord))

	//// Drawing the triangle ////

	// Clear the canvas
	gl.ClearColor(0.5, 0.5, 0.5, 0.9)
	gl.Clear(wasm.COLOR_BUFFER_BIT)

	// Enable the depth test
	gl.Enable(wasm.DEPTH_TEST)

	// Set the view port
	gl.Viewport(0, 0, width, height)

	// Draw the triangle
	gl.DrawElements(wasm.TRIANGLES, len(indicesNative), wasm.UNSIGNED_SHORT, 0)

	wasm.Loop()
}
