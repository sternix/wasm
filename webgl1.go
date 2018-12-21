// +build js,wasm

package wasm

import (
	"syscall/js"
)

// https://www.khronos.org/registry/webgl/specs/latest/1.0/

type (
	WebGLObject interface{}

	WebGLBuffer interface {
		WebGLObject
	}

	WebGLFramebuffer interface {
		WebGLObject
	}

	WebGLProgram interface {
		WebGLObject
	}

	WebGLRenderbuffer interface {
		WebGLObject
	}

	WebGLShader interface {
		WebGLObject
	}

	WebGLTexture interface {
		WebGLObject
	}

	WebGLUniformLocation interface{}

	WebGLActiveInfo interface {
		Size() int
		Type() GLenum
		Name() string
	}

	WebGLShaderPrecisionFormat interface {
		RangeMin() int
		RangeMax() int
		Precision() int
	}

	/*
			typedef (ImageBitmap or
		         ImageData or
		         HTMLImageElement or
		         HTMLCanvasElement or
		         HTMLVideoElement) TexImageSource;
	*/
	TexImageSource interface{}

	WebGLRenderingContextBase interface {
		Canvas() HTMLCanvasElement
		DrawingBufferWidth() int
		drawingBufferHeight() int
		GetContextAttributes() WebGLContextAttributes
		IsContextLost() bool
		GetSupportedExtensions() []string
		GetExtension(string) interface{}
		ActiveTexture(GLenum)
		AttachShader(WebGLProgram, WebGLShader)
		BindAttribLocation(WebGLProgram, int, string)
		BindBuffer(GLenum, WebGLBuffer)
		BindFramebuffer(GLenum, WebGLFramebuffer)
		BindRenderbuffer(GLenum, WebGLRenderbuffer)
		BindTexture(GLenum, WebGLTexture)
		BlendColor(float32, float32, float32, float32)
		BlendEquation(GLenum)
		BlendEquationSeparate(GLenum, GLenum)
		BlendFunc(GLenum, GLenum)
		BlendFuncSeparate(GLenum, GLenum, GLenum, GLenum)
		BufferData(GLenum, int, GLenum)
		BufferDataSource(GLenum, BufferSource, GLenum)
		BufferSubData(GLenum, int, BufferSource)
		CheckFramebufferStatus(GLenum) GLenum
		Clear(GLenum)
		ClearColor(float32, float32, float32, float32)
		ClearDepth(float32)
		ClearStencil(int)
		ColorMask(bool, bool, bool, bool)
		CompileShader(WebGLShader)
		CompressedTexImage2D(GLenum, int, GLenum, int, int, int, ArrayBufferView)
		CompressedTexSubImage2D(GLenum, int, int, int, int, int, GLenum, ArrayBufferView)
		CopyTexImage2D(GLenum, int, GLenum, int, int, int, int, int)
		CopyTexSubImage2D(GLenum, int, int, int, int, int, int, int)
		CreateBuffer() WebGLBuffer
		CreateFramebuffer() WebGLFramebuffer
		CreateProgram() WebGLProgram
		CreateRenderbuffer() WebGLRenderbuffer
		CreateShader(GLenum) WebGLShader
		CreateTexture() WebGLTexture
		CullFace(GLenum)
		DeleteBuffer(WebGLBuffer)
		DeleteFramebuffer(WebGLFramebuffer)
		DeleteProgram(WebGLProgram)
		DeleteRenderbuffer(WebGLRenderbuffer)
		DeleteShader(WebGLShader)
		DeleteTexture(WebGLTexture)
		DepthFunc(GLenum)
		DepthMask(bool)
		DepthRange(float32, float32)
		DetachShader(WebGLProgram, WebGLShader)
		Disable(GLenum)
		DisableVertexAttribArray(uint)
		DrawArrays(GLenum, int, int)
		DrawElements(GLenum, int, GLenum, int)
		Enable(GLenum)
		EnableVertexAttribArray(uint)
		Finish()
		Flush()
		FramebufferRenderbuffer(GLenum, GLenum, GLenum, WebGLRenderbuffer)
		FramebufferTexture2D(GLenum, GLenum, GLenum, WebGLTexture, int)
		FrontFace(GLenum)
		GenerateMipmap(GLenum)
		GetActiveAttrib(WebGLProgram, uint) WebGLActiveInfo
		GetActiveUniform(WebGLProgram, uint) WebGLActiveInfo
		GetAttachedShaders(WebGLProgram) []WebGLShader
		GetAttribLocation(WebGLProgram, string) int
		GetBufferParameter(GLenum, GLenum) interface{}
		GetParameter(GLenum) interface{}
		GetError() GLenum
		GetFramebufferAttachmentParameter(GLenum, GLenum, GLenum)
		GetProgramParameter(WebGLProgram, GLenum) interface{}
		GetProgramInfoLog(WebGLProgram) string
		GetRenderbufferParameter(GLenum, GLenum) interface{}
		GetShaderParameter(WebGLShader, GLenum) interface{}
		GetShaderPrecisionFormat(GLenum, GLenum) WebGLShaderPrecisionFormat
		GetShaderInfoLog(WebGLShader) string
		GetShaderSource(WebGLShader) string
		GetTexParameter(GLenum, GLenum) interface{}
		GetUniform(WebGLProgram, WebGLUniformLocation) interface{}
		GetUniformLocation(WebGLProgram, string) WebGLUniformLocation
		GetVertexAttrib(uint, GLenum) interface{}
		GetVertexAttribOffset(uint, GLenum) int
		Hint(GLenum, GLenum)
		IsBuffer(WebGLBuffer) bool
		IsEnabled(GLenum) bool
		IsFramebuffer(WebGLFramebuffer) bool
		IsProgram(WebGLProgram) bool
		IsRenderbuffer(WebGLRenderbuffer) bool
		IsShader(WebGLShader) bool
		IsTexture(WebGLTexture) bool
		LineWidth(float32)
		LinkProgram(WebGLProgram)
		PixelStorei(GLenum, int)
		PolygonOffset(float32, float32)
		ReadPixels(int, int, int, int, GLenum, GLenum, ArrayBufferView)
		RenderbufferStorage(GLenum, GLenum, int, int)
		SampleCoverage(float32, bool)
		Scissor(int, int, int, int)
		ShaderSource(WebGLShader, string)
		StencilFunc(GLenum, int, uint)
		StencilFuncSeparate(GLenum, GLenum, int, uint)
		StencilMask(uint)
		StencilMaskSeparate(GLenum, uint)
		StencilOp(GLenum, GLenum, GLenum)
		StencilOpSeparate(GLenum, GLenum, GLenum, GLenum)
		TexImage2DBuffer(GLenum, int, int, int, int, int, GLenum, GLenum, ArrayBufferView)
		TexImage2DSource(GLenum, int, int, GLenum, GLenum, TexImageSource)
		TexParameterf(GLenum, GLenum, float32)
		TexParameteri(GLenum, GLenum, int)
		TexSubImage2DBuffer(GLenum, int, int, int, int, int, GLenum, GLenum, ArrayBufferView)
		TexSubImage2DSource(GLenum, int, int, int, GLenum, GLenum, TexImageSource)
		Uniform1f(WebGLUniformLocation, float32)
		Uniform2f(WebGLUniformLocation, float32, float32)
		Uniform3f(WebGLUniformLocation, float32, float32, float32)
		Uniform4f(WebGLUniformLocation, float32, float32, float32, float32)
		Uniform1i(WebGLUniformLocation, int)
		Uniform2i(WebGLUniformLocation, int, int)
		Uniform3i(WebGLUniformLocation, int, int, int)
		Uniform4i(WebGLUniformLocation, int, int, int, int)
		Uniform1fv(WebGLUniformLocation, []float32)
		Uniform2fv(WebGLUniformLocation, []float32)
		Uniform3fv(WebGLUniformLocation, []float32)
		Uniform4fv(WebGLUniformLocation, []float32)
		Uniform1iv(WebGLUniformLocation, []int)
		Uniform2iv(WebGLUniformLocation, []int)
		Uniform3iv(WebGLUniformLocation, []int)
		Uniform4iv(WebGLUniformLocation, []int)
		UniformMatrix2fv(WebGLUniformLocation, bool, []float32)
		UniformMatrix3fv(WebGLUniformLocation, bool, []float32)
		UniformMatrix4fv(WebGLUniformLocation, bool, []float32)
		UseProgram(WebGLProgram)
		ValidateProgram(WebGLProgram)
		VertexAttrib1f(uint, float32)
		VertexAttrib2f(uint, float32, float32)
		VertexAttrib3f(uint, float32, float32, float32)
		VertexAttrib4f(uint, float32, float32, float32, float32)
		VertexAttrib1fv(uint, []float32)
		VertexAttrib2fv(uint, []float32)
		VertexAttrib3fv(uint, []float32)
		VertexAttrib4fv(uint, []float32)
		VertexAttribPointer(uint, int, GLenum, bool, int, int)
		Viewport(int, int, int, int)
	}

	WebGLRenderingContext interface {
		WebGLRenderingContextBase
	}
)

type WebGLPowerPreference string

const (
	WebGLPowerPreferenceDefault         WebGLPowerPreference = "default"
	WebGLPowerPreferenceLowPower        WebGLPowerPreference = "low-power"
	WebGLPowerPreferenceHighPerformance WebGLPowerPreference = "high-performance"
)

// -------------8<---------------------------------------

type WebGLContextAttributes struct {
	Alpha                        bool
	Depth                        bool
	Stencil                      bool
	Antialias                    bool
	PremultipliedAlpha           bool
	PreserveDrawingBuffer        bool
	PowerPreference              WebGLPowerPreference
	FailIfMajorPerformanceCaveat bool
}

func (p WebGLContextAttributes) toDict() js.Value {
	o := jsObject.New()
	o.Set("alpha", p.Alpha)
	o.Set("depth", p.Depth)
	o.Set("stencil", p.Stencil)
	o.Set("antialias", p.Antialias)
	o.Set("premultipliedAlpha", p.PremultipliedAlpha)
	o.Set("preserveDrawingBuffer", p.PreserveDrawingBuffer)
	o.Set("powerPreference", p.PowerPreference)
	o.Set("failIfMajorPerformanceCaveat", p.FailIfMajorPerformanceCaveat)
	return o
}
