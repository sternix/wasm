// +build js,wasm
// Drag mouse on canvas
//Wasming
// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go

/*
Code forked from

https://github.com/stdiopt/gowasm-experiments/tree/master/splashy

*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"syscall/js"
	// this box2d throws some unexpected panics
	"github.com/ByteArena/box2d"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/sternix/wasm"
)

var (
	width      int
	height     int
	simSpeed   = 1.0
	worldScale = 0.0125
	resDiv     = 8
	maxBodies  = 120
	gl         wasm.WebGLRenderingContext
)

func main() {

	// Init Canvas stuff
	win := wasm.CurrentWindow()
	doc := wasm.CurrentDocument()
	canvasEl := doc.ElementById("mycanvas").(wasm.HTMLCanvasElement)
	width = doc.Body().ClientWidth()
	height = doc.Body().ClientHeight()
	canvasEl.SetWidth(width)
	canvasEl.SetHeight(height)

	gl = canvasEl.ContextWebGL()
	if gl == nil {
		win.Alert("browser might not support webgl")
		return
	}

	thing := Thing{}
	mouseDown := false

	mouseDownEvt := func(evt wasm.MouseEvent) {
		mouseDown = true
		if !wasm.Equal(evt.Target(), canvasEl) {
			return
		}
		mx := evt.ClientX() * worldScale
		my := evt.ClientY() * worldScale
		thing.AddCircle(mx, my)
	}

	mouseUpEvt := func(wasm.MouseEvent) {
		mouseDown = false
	}

	mouseMoveEvt := func(evt wasm.MouseEvent) {
		if !mouseDown {
			return
		}
		if !wasm.Equal(evt.Target(), canvasEl) {
			return
		}
		mx := evt.ClientX() * worldScale
		my := evt.ClientY() * worldScale
		thing.AddCircle(mx, my)
	}

	speedInputEvt := func(evt wasm.InputEvent) {
		rangeInput := evt.Target().(wasm.HTMLInputElement)
		fval, err := strconv.ParseFloat(rangeInput.Value(), 64)
		if err != nil {
			log.Println("Invalid value", err)
			return
		}
		simSpeed = fval
	}

	// Events
	doc.OnMouseDown(mouseDownEvt)
	doc.OnMouseUp(mouseUpEvt)
	doc.OnMouseMove(mouseMoveEvt)
	doc.ElementById("speed").(wasm.HTMLInputElement).OnInput(speedInputEvt)

	err := thing.Init(gl)
	if err != nil {
		log.Println("Err Initializing thing:", err)
		return
	}

	// Draw things
	var tmark float64
	var markCount = 0
	var tdiffSum float64

	fps := doc.ElementById("fps").(wasm.HTMLDivElement)

	var rafcb int
	renderFrame := func(cb wasm.FrameRequestCallback, now float64) {
		// Update the DOM less frequently TODO: func on this
		tdiff := now - tmark
		tdiffSum += tdiff
		markCount++
		if markCount > 10 {
			fps.SetInnerHTML(fmt.Sprintf("FPS: %.01f", 1000/(tdiffSum/float64(markCount))))
			tdiffSum, markCount = 0, 0
		}
		tmark = now
		// --
		thing.Render(tdiff / 1000)

		rafcb = win.RequestAnimationFrame(cb)
	}
	frcb := wasm.NewFrameRequestCallback(renderFrame)
	// Start running
	rafcb = win.RequestAnimationFrame(frcb)

	wasm.Loop()

	win.CancelAnimationFrame(rafcb)
	frcb.Release()
}

type Thing struct {
	// dot shaders
	prog        wasm.WebGLProgram
	aPosition   int
	uFragColor  wasm.WebGLUniformLocation
	uResolution wasm.WebGLUniformLocation

	dotBuf     wasm.WebGLBuffer
	qBlur      *QuadFX
	qThreshold *QuadFX

	rtTex [2]wasm.WebGLTexture     // render target Texture
	rt    [2]wasm.WebGLFramebuffer // framebuffer(render target)

	world box2d.B2World
}

func (t *Thing) Init(gl wasm.WebGLRenderingContext) error {
	// Drawing program
	var err error
	t.prog, err = programFromSrc(dotVertShader, dotFragShader)
	if err != nil {
		return err
	}
	t.aPosition = gl.GetAttribLocation(t.prog, "a_position")
	t.uFragColor = gl.GetUniformLocation(t.prog, "uFragColor")
	t.uResolution = gl.GetUniformLocation(t.prog, "uResolution")

	t.dotBuf = gl.CreateBuffer()
	//renderer targets
	for i := 0; i < 2; i++ {
		t.rtTex[i] = createTexture(width/resDiv, height/resDiv)
		t.rt[i] = createFB(t.rtTex[i])
	}

	t.qBlur = &QuadFX{}
	err = t.qBlur.Init(blurShader)
	if err != nil {
		log.Fatal("Error:", err)
	}
	t.qThreshold = &QuadFX{}
	err = t.qThreshold.Init(thresholdShader)
	if err != nil {
		log.Fatal("Error:", err)
	}

	//////////////////////////
	// Physics
	// ///////////
	t.world = box2d.MakeB2World(box2d.B2Vec2{X: 0, Y: 9.8})
	floor := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: 0, Y: float64(height+10) * worldScale},
		Active:   true,
	})
	floorShape := &box2d.B2PolygonShape{}
	floorShape.SetAsBox(float64(width)*worldScale, 20*worldScale)
	ft := floor.CreateFixture(floorShape, 1)
	ft.M_friction = 0.3

	// Walls
	wallShape := &box2d.B2PolygonShape{}
	wallShape.SetAsBox(20*worldScale, float64(height)*worldScale)

	wallL := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: 0, Y: 0},
		Active:   true,
	})
	wlf := wallL.CreateFixture(wallShape, 1)
	wlf.M_friction = 0.3

	wallR := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: float64(width) * worldScale, Y: 0},
		Active:   true,
	})
	wrt := wallR.CreateFixture(wallShape, 1)
	wrt.M_friction = 0.3

	for i := 0; i < 10; i++ {
		t.AddCircle(rand.Float64()*float64(width)*worldScale, rand.Float64()*float64(height)*worldScale)
	}

	return nil
}

func (t *Thing) Render(dtTime float64) {

	texWidth := width / resDiv
	texHeight := height / resDiv
	t.world.Step(dtTime*simSpeed, 3, 3)

	gl.BindFramebuffer(wasm.FRAMEBUFFER, t.rt[0])
	gl.Viewport(0, 0, texWidth, texHeight) //texSize
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(wasm.COLOR_BUFFER_BIT)

	// DotRenderer
	gl.UseProgram(t.prog)

	count := 0
	for curBody := t.world.GetBodyList(); curBody != nil; curBody = curBody.M_next {
		ft := curBody.M_fixtureList
		if _, ok := ft.M_shape.(*box2d.B2CircleShape); !ok {
			continue
		}
		x := float32(curBody.M_xf.P.X / (float64(width) * worldScale))  /* 0-1 */
		y := float32(curBody.M_xf.P.Y / (float64(height) * worldScale)) /*0-1*/

		col := colorful.Hsv(float64(360*count/maxBodies), 1, 1)
		gl.VertexAttrib2f(uint(t.aPosition), x, y)
		gl.Uniform4f(t.uFragColor, float32(col.R), float32(col.G), float32(col.B), 1.0)
		gl.DrawArrays(wasm.POINTS, 0, 1)

		count++
		// Stop processing
		if count > maxBodies {
			break
		}
	}

	/// FX Blurx4 TODO: better blur
	for i := 0; i < 4; i++ {
		gl.BindFramebuffer(wasm.FRAMEBUFFER, t.rt[1])
		gl.Viewport(0, 0, texWidth, texHeight)
		gl.BindTexture(wasm.TEXTURE_2D, t.rtTex[0])
		t.qBlur.Render()

		gl.BindFramebuffer(wasm.FRAMEBUFFER, t.rt[0])
		gl.Viewport(0, 0, texWidth, texHeight)
		gl.BindTexture(wasm.TEXTURE_2D, t.rtTex[1])
		t.qBlur.Render()
	}

	/// FX Threshold to Screen
	gl.BindFramebuffer(wasm.FRAMEBUFFER, nil)
	gl.Viewport(0, 0, width, height)
	gl.BindTexture(wasm.TEXTURE_2D, t.rtTex[0])
	t.qThreshold.Render()
}

func (t *Thing) AddCircle(mx, my float64) {
	if t.world.GetBodyCount() > maxBodies {
		// Check for the last on list and delete backwards:o
		var b *box2d.B2Body
		// theres is no M_last but we could cache it somewhere
		for b = t.world.GetBodyList(); b.M_next != nil; b = b.M_next {
		}
		// Search backwards for a circle (ignoring the walls/floors)
		for ; b != nil; b = b.M_prev {
			if _, ok := b.M_fixtureList.M_shape.(*box2d.B2CircleShape); ok {
				t.world.DestroyBody(b) // Destroy first found body
				break
			}
		}
	}
	obj1 := t.world.CreateBody(&box2d.B2BodyDef{
		Type:         box2d.B2BodyType.B2_dynamicBody,
		Position:     box2d.B2Vec2{X: mx, Y: my},
		Awake:        true,
		Active:       true,
		GravityScale: 1.0,
	})
	shape := box2d.NewB2CircleShape()
	shape.M_radius = 10 * worldScale
	ft := obj1.CreateFixture(shape, 1)
	ft.M_friction = 0.2
	ft.M_restitution = 0.6
}

//// SHADERS & Utils
const dotVertShader = `
attribute vec4 a_position;
void main () {
	vec4 lpos= vec4(a_position.xy*2.0-1.0, 0, 1);
	lpos.y = -lpos.y;
	gl_Position = lpos;
	gl_PointSize = 22.0/4.0;
}
`
const dotFragShader = `
precision mediump float;
uniform vec4 uFragColor;
void main () {
	vec2 pt = gl_PointCoord - vec2(0.5);
	if(pt.x*pt.x+pt.y*pt.y > 0.25)
	  discard;
	gl_FragColor = uFragColor;
}
`

const blurShader = `
precision mediump float;
uniform sampler2D u_image;
uniform vec2 u_textureSize;
varying vec2 v_texCoord;
void main() {
	vec2 onePixel = vec2(1,1) / u_textureSize;
	vec4 colorSum =
     texture2D(u_image, v_texCoord + onePixel * vec2(-1, -1)) + 
     texture2D(u_image, v_texCoord + onePixel * vec2( 0, -1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1, -1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2(-1,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 0,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2(-1,  1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 0,  1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1,  1));
  gl_FragColor = colorSum / 9.0;
}
`

const thresholdShader = `
precision mediump float;
uniform sampler2D u_image;
uniform vec2 u_textureSize;
varying vec2 v_texCoord;
void main() {
	float a;
	vec2 onePixel = vec2(1,1) / u_textureSize;
	vec4 col = texture2D(u_image,v_texCoord);
	if (col.a < 0.4) discard;
	if (col.a < 0.8 && col.a > 0.72) {
		a = texture2D(u_image, v_texCoord + onePixel * vec2(-1, 1)).a;
		if (a < col.a ) {
			col += 0.4;
		}
	} 
	gl_FragColor = vec4(col.rgb,1.0);
}
`

const vertQuad = `
attribute vec2 a_position;
attribute vec2 a_texCoord;
varying vec2 v_texCoord;
void main() {
   gl_Position = vec4((a_position * 2.0 - 1.0), 0, 1);
   v_texCoord = a_texCoord;
 }
`

type QuadFX struct {
	prog         wasm.WebGLProgram
	aPosition    uint
	aTexCoord    uint
	uTextureSize wasm.WebGLUniformLocation

	quadBuf wasm.WebGLBuffer

	vertexData js.TypedArray
}

func (q *QuadFX) Init(frag string) error {
	var err error
	q.prog, err = programFromSrc(vertQuad, frag)
	if err != nil {
		return err
	}
	q.vertexData = js.TypedArrayOf([]float32{
		0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 1.0, 0.0, 1.0, 1.0,
	})

	q.aPosition = uint(gl.GetAttribLocation(q.prog, "a_position"))
	q.aTexCoord = uint(gl.GetAttribLocation(q.prog, "a_texCoord"))
	q.uTextureSize = gl.GetUniformLocation(q.prog, "u_textureSize")

	q.quadBuf = gl.CreateBuffer()
	// texCoord/posCoord
	gl.BindBuffer(wasm.ARRAY_BUFFER, q.quadBuf)
	gl.BufferDataSource(wasm.ARRAY_BUFFER, q.vertexData, wasm.STATIC_DRAW)
	return nil

}
func (q *QuadFX) Render() {
	gl.UseProgram(q.prog)
	// Vertex
	gl.BindBuffer(wasm.ARRAY_BUFFER, q.quadBuf)

	gl.EnableVertexAttribArray(q.aPosition)
	gl.VertexAttribPointer(q.aPosition, 2, wasm.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(q.aTexCoord) // sabe buf
	gl.VertexAttribPointer(q.aTexCoord, 2, wasm.FLOAT, false, 0, 0)

	gl.Uniform2f(q.uTextureSize, float32(width/resDiv), float32(height/resDiv))

	gl.DrawArrays(wasm.TRIANGLES, 0 /*offset*/, 6 /*count*/)
	gl.DisableVertexAttribArray(q.aPosition)
	gl.DisableVertexAttribArray(q.aTexCoord)

}

func (q *QuadFX) Release() {
	q.vertexData.Release()
	// TODO: gl release
}

// Helper funcs

// Render to framebuffer first, then framebuffer to screen
func compileShader(shaderType wasm.GLenum, shaderSrc string) (wasm.WebGLShader, error) {
	var shader = gl.CreateShader(shaderType)
	gl.ShaderSource(shader, shaderSrc)
	gl.CompileShader(shader)

	if !gl.GetShaderParameter(shader, wasm.COMPILE_STATUS).(bool) {
		return nil, fmt.Errorf("could not compile shader: %s", gl.GetShaderInfoLog(shader))
	}
	return shader, nil
}

func linkProgram(vertexShader, fragmentShader wasm.WebGLShader) (wasm.WebGLProgram, error) {
	var program = gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	if !gl.GetProgramParameter(program, wasm.LINK_STATUS).(bool) {
		return nil, fmt.Errorf("could not link program: %s", gl.GetProgramInfoLog(program))
	}

	return program, nil
}

func programFromSrc(vertShaderSrc, fragShaderSrc string) (wasm.WebGLProgram, error) {
	vertexShader, err := compileShader(wasm.VERTEX_SHADER, vertShaderSrc)
	if err != nil {
		return nil, err
	}
	fragShader, err := compileShader(wasm.FRAGMENT_SHADER, fragShaderSrc)
	if err != nil {
		return nil, err
	}
	prog, err := linkProgram(vertexShader, fragShader)
	if err != nil {
		return nil, err
	}
	return prog, nil
}

func createTexture(width, height int) wasm.WebGLTexture {
	tex := gl.CreateTexture()
	gl.BindTexture(wasm.TEXTURE_2D, tex)
	// define size and format of level 0
	gl.TexImage2DBuffer(
		wasm.TEXTURE_2D,
		0,
		int(wasm.RGBA),
		width, height,
		0,
		wasm.RGBA,
		wasm.UNSIGNED_BYTE,
		nil,
	)

	// set the filtering so we don't need mips
	gl.TexParameteri(wasm.TEXTURE_2D, wasm.TEXTURE_MIN_FILTER, int(wasm.LINEAR))
	gl.TexParameteri(wasm.TEXTURE_2D, wasm.TEXTURE_MAG_FILTER, int(wasm.LINEAR))
	gl.TexParameteri(wasm.TEXTURE_2D, wasm.TEXTURE_WRAP_S, int(wasm.CLAMP_TO_EDGE))
	gl.TexParameteri(wasm.TEXTURE_2D, wasm.TEXTURE_WRAP_T, int(wasm.CLAMP_TO_EDGE))

	return tex
}

// Create framebuffer binded to texture
func createFB(tex wasm.WebGLTexture) wasm.WebGLFramebuffer {
	fb := gl.CreateFramebuffer()
	gl.BindFramebuffer(wasm.FRAMEBUFFER, fb)
	gl.FramebufferTexture2D(wasm.FRAMEBUFFER, wasm.COLOR_ATTACHMENT0, wasm.TEXTURE_2D, tex, 0)
	return fb
}
