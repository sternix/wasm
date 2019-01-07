// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type webGLObjectImpl struct {
	Value
}

func wrapWebGLObject(v Value) WebGLObject {
	if p := newWebGLObjectImpl(v); p != nil {
		return p
	}
	return nil
}

func newWebGLObjectImpl(v Value) *webGLObjectImpl {
	if v.Valid() {
		return &webGLObjectImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLBufferImpl struct {
	*webGLObjectImpl
}

func wrapWebGLBuffer(v Value) WebGLBuffer {
	if v.Valid() {
		return &webGLBufferImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLFramebufferImpl struct {
	*webGLObjectImpl
}

func wrapWebGLFramebuffer(v Value) WebGLFramebuffer {
	if v.Valid() {
		return &webGLFramebufferImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLProgramImpl struct {
	*webGLObjectImpl
}

func wrapWebGLProgram(v Value) WebGLProgram {
	if v.Valid() {
		return &webGLProgramImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLRenderbufferImpl struct {
	*webGLObjectImpl
}

func wrapWebGLRenderbuffer(v Value) WebGLRenderbuffer {
	if v.Valid() {
		return &webGLRenderbufferImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLShaderImpl struct {
	*webGLObjectImpl
}

func wrapWebGLShader(v Value) WebGLShader {
	if v.Valid() {
		return &webGLShaderImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLTextureImpl struct {
	*webGLObjectImpl
}

func wrapWebGLTexture(v Value) WebGLTexture {
	if v.Valid() {
		return &webGLTextureImpl{
			webGLObjectImpl: newWebGLObjectImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLUniformLocationImpl struct {
	Value
}

func wrapWebGLUniformLocation(v Value) WebGLUniformLocation {
	if v.Valid() {
		return &webGLUniformLocationImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type webGLActiveInfoImpl struct {
	Value
}

func wrapWebGLActiveInfo(v Value) WebGLActiveInfo {
	if v.Valid() {
		return &webGLActiveInfoImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLActiveInfoImpl) Size() int {
	return p.Get("size").Int()
}

func (p *webGLActiveInfoImpl) Type() GLenum {
	return GLenum(p.Get("type").Int())
}

func (p *webGLActiveInfoImpl) Name() string {
	return p.Get("name").String()
}

// -------------8<---------------------------------------

type webGLShaderPrecisionFormatImpl struct {
	Value
}

func wrapWebGLShaderPrecisionFormat(v Value) WebGLShaderPrecisionFormat {
	if v.Valid() {
		return &webGLShaderPrecisionFormatImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLShaderPrecisionFormatImpl) RangeMin() int {
	return p.Get("rangeMin").Int()
}

func (p *webGLShaderPrecisionFormatImpl) RangeMax() int {
	return p.Get("rangeMax").Int()
}

func (p *webGLShaderPrecisionFormatImpl) Precision() int {
	return p.Get("precision").Int()
}

// -------------8<---------------------------------------

func wrapWebGLContextAttributes(v Value) WebGLContextAttributes {
	if v.Valid() {
		return WebGLContextAttributes{
			Alpha:                        v.Get("alpha").Bool(),
			Depth:                        v.Get("depth").Bool(),
			Stencil:                      v.Get("stencil").Bool(),
			Antialias:                    v.Get("antialias").Bool(),
			PremultipliedAlpha:           v.Get("premultipliedAlpha").Bool(),
			PreserveDrawingBuffer:        v.Get("preserveDrawingBuffer").Bool(),
			PowerPreference:              WebGLPowerPreference(v.Get("powerPreference").String()),
			FailIfMajorPerformanceCaveat: v.Get("failIfMajorPerformanceCaveat").Bool(),
		}
	}
	return WebGLContextAttributes{}
}

// -------------8<---------------------------------------

type texImageSourceImpl struct {
	Value
}

func wrapTexImageSource(v Value) TexImageSource {
	if v.Valid() {
		return &texImageSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

var _ WebGLRenderingContextBase = &webGLRenderingContextBaseImpl{}

type webGLRenderingContextBaseImpl struct {
	Value
}

func newWebGLRenderingContextBaseImpl(v Value) *webGLRenderingContextBaseImpl {
	if v.Valid() {
		return &webGLRenderingContextBaseImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLRenderingContextBaseImpl) Canvas() HTMLCanvasElement {
	return wrapHTMLCanvasElement(p.Get("canvas"))
}

func (p *webGLRenderingContextBaseImpl) DrawingBufferWidth() int {
	return p.Get("drawingBufferWidth").Int()
}

func (p *webGLRenderingContextBaseImpl) drawingBufferHeight() int {
	return p.Get("drawingBufferHeight").Int()
}

func (p *webGLRenderingContextBaseImpl) GetContextAttributes() WebGLContextAttributes {
	return wrapWebGLContextAttributes(p.Call("getContextAttributes"))
}

func (p *webGLRenderingContextBaseImpl) IsContextLost() bool {
	return p.Call("isContextLost").Bool()
}

func (p *webGLRenderingContextBaseImpl) GetSupportedExtensions() []string {
	return stringSequenceToSlice(p.Call("getSupportedExtensions"))
}

func (p *webGLRenderingContextBaseImpl) GetExtension(string) interface{} {
	return Wrap(p.Call("getExtension"))
}

func (p *webGLRenderingContextBaseImpl) ActiveTexture(texture GLenum) {
	p.Call("activeTexture", uint(texture))
}

func (p *webGLRenderingContextBaseImpl) AttachShader(program WebGLProgram, shader WebGLShader) {
	p.Call("attachShader", JSValue(program), JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) BindAttribLocation(program WebGLProgram, index int, name string) {
	p.Call("bindAttribLocation", JSValue(program), index, name)
}

func (p *webGLRenderingContextBaseImpl) BindBuffer(target GLenum, buffer WebGLBuffer) {
	var b Value
	if buffer != nil {
		b = JSValue(buffer)
	} else {
		b = jsNull
	}

	p.Call("bindBuffer", uint(target), b)
}

func (p *webGLRenderingContextBaseImpl) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) {
	if framebuffer != nil {
		p.Call("bindFramebuffer", uint(target), JSValue(framebuffer))
	} else {
		p.Call("bindFramebuffer", uint(target), js.Null())
	}
}

func (p *webGLRenderingContextBaseImpl) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) {
	p.Call("bindRenderbuffer", uint(target), JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) BindTexture(target GLenum, texture WebGLTexture) {
	p.Call("bindTexture", uint(target), JSValue(texture))
}

func (p *webGLRenderingContextBaseImpl) BlendColor(red, green, blue, alpha float32) {
	p.Call("blendColor", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) BlendEquation(mode GLenum) {
	p.Call("blendEquation", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) {
	p.Call("blendEquationSeparate", uint(modeRGB), uint(modeAlpha))
}

func (p *webGLRenderingContextBaseImpl) BlendFunc(sfactor GLenum, dfactor GLenum) {
	p.Call("blendFunc", uint(sfactor), uint(dfactor))
}

func (p *webGLRenderingContextBaseImpl) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLenum) {
	p.Call("blendFuncSeparate", uint(srcRGB), uint(dstRGB), uint(srcAlpha), uint(dstAlpha))
}

func (p *webGLRenderingContextBaseImpl) BufferData(target GLenum, size int, usage GLenum) {
	p.Call("bufferData", uint(target), size, uint(usage))
}

func (p *webGLRenderingContextBaseImpl) BufferDataSource(target GLenum, data BufferSource, usage GLenum) {
	p.Call("bufferData", uint(target), JSValue(data), uint(usage))
}

func (p *webGLRenderingContextBaseImpl) BufferSubData(target GLenum, offset int, data BufferSource) {
	p.Call("bufferSubData", uint(target), offset, JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(p.Call("checkFramebufferStatus", uint(target)).Int())
}

func (p *webGLRenderingContextBaseImpl) Clear(mask GLenum) {
	p.Call("clear", uint(mask))
}

func (p *webGLRenderingContextBaseImpl) ClearColor(red, green, blue, alpha float32) {
	p.Call("clearColor", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) ClearDepth(depth float32) {
	p.Call("clearDepth", depth)
}

func (p *webGLRenderingContextBaseImpl) ClearStencil(s int) {
	p.Call("clearStencil", s)
}

func (p *webGLRenderingContextBaseImpl) ColorMask(red, green, blue, alpha bool) {
	p.Call("colorMask", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) CompileShader(shader WebGLShader) {
	p.Call("compileShader", JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) CompressedTexImage2D(target GLenum, level int, internalFormat GLenum, width int, height int, border int, data ArrayBufferView) {
	p.Call("compressedTexImage2D", uint(target), level, uint(internalFormat), width, height, border, JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CompressedTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, data ArrayBufferView) {
	p.Call("compressedTexSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	p.Call("copyTexImage2D", uint(target), level, uint(internalFormat), x, y, width, height, border)
}

func (p *webGLRenderingContextBaseImpl) CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	p.Call("copyTexSubImage2D", uint(target), level, xoffset, yoffset, x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) CreateBuffer() WebGLBuffer {
	return wrapWebGLBuffer(p.Call("createBuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateFramebuffer() WebGLFramebuffer {
	return wrapWebGLFramebuffer(p.Call("createFramebuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateProgram() WebGLProgram {
	return wrapWebGLProgram(p.Call("createProgram"))
}

func (p *webGLRenderingContextBaseImpl) CreateRenderbuffer() WebGLRenderbuffer {
	return wrapWebGLRenderbuffer(p.Call("createRenderbuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateShader(typ GLenum) WebGLShader {
	return wrapWebGLShader(p.Call("createShader", uint(typ)))
}

func (p *webGLRenderingContextBaseImpl) CreateTexture() WebGLTexture {
	return wrapWebGLTexture(p.Call("createTexture"))
}

func (p *webGLRenderingContextBaseImpl) CullFace(mode GLenum) {
	p.Call("cullFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) DeleteBuffer(buffer WebGLBuffer) {
	p.Call("deleteBuffer", JSValue(buffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteFramebuffer(frameBuffer WebGLFramebuffer) {
	p.Call("deleteFramebuffer", JSValue(frameBuffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteProgram(program WebGLProgram) {
	p.Call("deleteProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) {
	p.Call("deleteRenderbuffer", JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteShader(shader WebGLShader) {
	p.Call("deleteShader", JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) DeleteTexture(texture WebGLTexture) {
	p.Call("deleteTexture", JSValue(texture))
}

func (p *webGLRenderingContextBaseImpl) DepthFunc(fn GLenum) {
	p.Call("depthFunc", uint(fn))
}

func (p *webGLRenderingContextBaseImpl) DepthMask(flag bool) {
	p.Call("depthMask", flag)
}

func (p *webGLRenderingContextBaseImpl) DepthRange(zNear, zFar float32) {
	p.Call("depthRange", zNear, zFar)
}

func (p *webGLRenderingContextBaseImpl) DetachShader(program WebGLProgram, shader WebGLShader) {
	p.Call("detachShader", JSValue(program), JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) Disable(c GLenum) {
	p.Call("disable", uint(c))
}

func (p *webGLRenderingContextBaseImpl) DisableVertexAttribArray(index uint) {
	p.Call("disableVertexAttribArray", index)
}

func (p *webGLRenderingContextBaseImpl) DrawArrays(mode GLenum, first int, count int) {
	p.Call("drawArrays", uint(mode), first, count)
}

func (p *webGLRenderingContextBaseImpl) DrawElements(mode GLenum, count int, typ GLenum, offset int) {
	p.Call("drawElements", uint(mode), count, uint(typ), offset)
}

func (p *webGLRenderingContextBaseImpl) Enable(c GLenum) {
	p.Call("enable", uint(c))
}

func (p *webGLRenderingContextBaseImpl) EnableVertexAttribArray(index uint) {
	p.Call("enableVertexAttribArray", index)
}

func (p *webGLRenderingContextBaseImpl) Finish() {
	p.Call("finish")
}

func (p *webGLRenderingContextBaseImpl) Flush() {
	p.Call("flush")
}

func (p *webGLRenderingContextBaseImpl) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) {
	p.Call("framebufferRenderbuffer", uint(target), uint(attachment), uint(renderbuffertarget), JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level int) {
	p.Call("framebufferTexture2D", uint(target), uint(attachment), uint(textarget), JSValue(texture), level)
}

func (p *webGLRenderingContextBaseImpl) FrontFace(mode GLenum) {
	p.Call("frontFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) GenerateMipmap(target GLenum) {
	p.Call("generateMipmap", uint(target))
}

func (p *webGLRenderingContextBaseImpl) GetActiveAttrib(program WebGLProgram, index uint) WebGLActiveInfo {
	return wrapWebGLActiveInfo(p.Call("getActiveAttrib", JSValue(program), index))
}

func (p *webGLRenderingContextBaseImpl) GetActiveUniform(program WebGLProgram, index uint) WebGLActiveInfo {
	return wrapWebGLActiveInfo(p.Call("getActiveUniform", JSValue(program), index))
}

func (p *webGLRenderingContextBaseImpl) GetAttachedShaders(program WebGLProgram) []WebGLShader {
	if s := p.Call("getAttachedShaders", JSValue(program)).ToSlice(); s != nil {
		ret := make([]WebGLShader, len(s))
		for i, v := range s {
			ret[i] = wrapWebGLShader(v)
		}
		return ret
	}
	return nil
}

func (p *webGLRenderingContextBaseImpl) GetAttribLocation(program WebGLProgram, name string) int {
	return p.Call("getAttribLocation", JSValue(program), name).Int()
}

func (p *webGLRenderingContextBaseImpl) GetBufferParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.Call("getBufferParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetParameter(pname GLenum) interface{} {
	return Wrap(p.Call("getParameter", uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetError() GLenum {
	return GLenum(p.Call("getError").Int())
}

func (p *webGLRenderingContextBaseImpl) GetFramebufferAttachmentParameter(target, attachment, pname GLenum) {
	p.Call("getFramebufferAttachmentParameter", uint(target), uint(attachment), uint(pname))
}

func (p *webGLRenderingContextBaseImpl) GetProgramParameter(program WebGLProgram, pname GLenum) interface{} {
	return Wrap(p.Call("getProgramParameter", JSValue(program), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetProgramInfoLog(program WebGLProgram) string {
	return p.Call("getProgramInfoLog", JSValue(program)).String()
}

func (p *webGLRenderingContextBaseImpl) GetRenderbufferParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.Call("getRenderbufferParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderParameter(shader WebGLShader, pname GLenum) interface{} {
	return Wrap(p.Call("getShaderParameter", JSValue(shader), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderPrecisionFormat(shaderType GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat {
	return wrapWebGLShaderPrecisionFormat(p.Call("getShaderPrecisionFormat", uint(shaderType), uint(precisiontype)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderInfoLog(shader WebGLShader) string {
	return p.Call("getShaderInfoLog", JSValue(shader)).String()
}

func (p *webGLRenderingContextBaseImpl) GetShaderSource(shader WebGLShader) string {
	return p.Call("getShaderSource", JSValue(shader)).String()
}

func (p *webGLRenderingContextBaseImpl) GetTexParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.Call("getTexParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetUniform(program WebGLProgram, location WebGLUniformLocation) interface{} {
	return Wrap(p.Call("getUniform", JSValue(program), JSValue(location)))
}

func (p *webGLRenderingContextBaseImpl) GetUniformLocation(program WebGLProgram, name string) WebGLUniformLocation {
	return wrapWebGLUniformLocation(p.Call("getUniformLocation", JSValue(program), name))
}

func (p *webGLRenderingContextBaseImpl) GetVertexAttrib(index uint, pname GLenum) interface{} {
	return Wrap(p.Call("getVertexAttrib", index, uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetVertexAttribOffset(index uint, pname GLenum) int {
	return p.Call("getVertexAttribOffset", index, uint(pname)).Int()
}

func (p *webGLRenderingContextBaseImpl) Hint(target GLenum, mode GLenum) {
	p.Call("hint", uint(target), uint(mode))
}

func (p *webGLRenderingContextBaseImpl) IsBuffer(buffer WebGLBuffer) bool {
	return p.Call("isBuffer", JSValue(buffer)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsEnabled(c GLenum) bool {
	return p.Call("isEnabled", uint(c)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsFramebuffer(buffer WebGLFramebuffer) bool {
	return p.Call("isFramebuffer", JSValue(buffer)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsProgram(program WebGLProgram) bool {
	return p.Call("isProgram", JSValue(program)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsRenderbuffer(buffer WebGLRenderbuffer) bool {
	return p.Call("isRenderbuffer", JSValue(buffer)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsShader(shader WebGLShader) bool {
	return p.Call("isShader", JSValue(shader)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsTexture(texture WebGLTexture) bool {
	return p.Call("isTexture", JSValue(texture)).Bool()
}

func (p *webGLRenderingContextBaseImpl) LineWidth(width float32) {
	p.Call("lineWidth", width)
}

func (p *webGLRenderingContextBaseImpl) LinkProgram(program WebGLProgram) {
	p.Call("linkProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) PixelStorei(pname GLenum, param int) {
	p.Call("pixelStorei", uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) PolygonOffset(factor float32, units float32) {
	p.Call("polygonOffset", factor, units)
}

func (p *webGLRenderingContextBaseImpl) ReadPixels(x int, y int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.Call("readPixels", x, y, width, height, uint(format), uint(typ), JSValue(pixels))
}

func (p *webGLRenderingContextBaseImpl) RenderbufferStorage(target GLenum, format GLenum, width int, height int) {
	p.Call("renderbufferStorage", uint(target), uint(format), width, height)
}

func (p *webGLRenderingContextBaseImpl) SampleCoverage(value float32, invert bool) {
	p.Call("sampleCoverage", value, invert)
}

func (p *webGLRenderingContextBaseImpl) Scissor(x int, y int, width int, height int) {
	p.Call("scissor", x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) ShaderSource(shader WebGLShader, source string) {
	p.Call("shaderSource", JSValue(shader), source)
}

func (p *webGLRenderingContextBaseImpl) StencilFunc(fn GLenum, ref int, mask uint) {
	p.Call("stencilFunc", uint(fn), ref, mask)
}

func (p *webGLRenderingContextBaseImpl) StencilFuncSeparate(face GLenum, fn GLenum, ref int, mask uint) {
	p.Call("stencilFuncSeparate", uint(face), uint(fn), ref, mask)
}

func (p *webGLRenderingContextBaseImpl) StencilMask(mask uint) {
	p.Call("stencilMask", mask)
}

func (p *webGLRenderingContextBaseImpl) StencilMaskSeparate(face GLenum, mask uint) {
	p.Call("stencilMaskSeparate", uint(face), mask)
}

func (p *webGLRenderingContextBaseImpl) StencilOp(fail, zfail, zpass GLenum) {
	p.Call("stencilOp", uint(fail), uint(zfail), uint(zpass))
}

func (p *webGLRenderingContextBaseImpl) StencilOpSeparate(face, fail, zfail, zpass GLenum) {
	p.Call("stencilOpSeparate", face, fail, zfail, zpass)
}

func (p *webGLRenderingContextBaseImpl) TexImage2DBuffer(target GLenum, level int, internalFormat int, width int, height int, border int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	if pixels != nil {
		p.Call("texImage2D", uint(target), level, internalFormat, width, height, border, uint(format), uint(typ), JSValue(pixels))
	} else {
		p.Call("texImage2D", uint(target), level, internalFormat, width, height, border, uint(format), uint(typ), nil)
	}
}

func (p *webGLRenderingContextBaseImpl) TexImage2DSource(target GLenum, level int, internalFormat int, format GLenum, typ GLenum, source TexImageSource) {
	p.Call("texImage2D", uint(target), level, internalFormat, uint(format), uint(typ), JSValue(source))
}

func (p *webGLRenderingContextBaseImpl) TexParameterf(target GLenum, pname GLenum, param float32) {
	p.Call("texParameterf", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexParameteri(target GLenum, pname GLenum, param int) {
	p.Call("texParameteri", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DBuffer(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.Call("texSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), uint(typ), JSValue(pixels))
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DSource(target GLenum, level int, xoffset int, yoffset int, format GLenum, typ GLenum, source TexImageSource) {
	p.Call("texSubImage2D", uint(target), level, xoffset, yoffset, uint(format), uint(typ), JSValue(source))
}

func (p *webGLRenderingContextBaseImpl) Uniform1f(location WebGLUniformLocation, x float32) {
	p.Call("uniform1f", JSValue(location), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2f(location WebGLUniformLocation, x float32, y float32) {
	p.Call("uniform2f", JSValue(location), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3f(location WebGLUniformLocation, x float32, y float32, z float32) {
	p.Call("uniform3f", JSValue(location), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4f(location WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	p.Call("uniform4f", JSValue(location), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1i(location WebGLUniformLocation, x int) {
	p.Call("uniform1i", JSValue(location), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2i(location WebGLUniformLocation, x int, y int) {
	p.Call("uniform2i", JSValue(location), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3i(location WebGLUniformLocation, x int, y int, z int) {
	p.Call("uniform3i", JSValue(location), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4i(location WebGLUniformLocation, x int, y int, z int, w int) {
	p.Call("uniform4i", JSValue(location), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform1fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform2fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform3fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform4fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform1iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform1iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform2iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform3iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform4iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix2fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix2fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix3fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix3fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix4fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix4fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UseProgram(program WebGLProgram) {
	p.Call("useProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) ValidateProgram(program WebGLProgram) {
	p.Call("validateProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1f(index uint, x float32) {
	p.Call("vertexAttrib1f", index, x)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2f(index uint, x, y float32) {
	p.Call("vertexAttrib2f", index, x, y)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3f(index uint, x, y, z float32) {
	p.Call("vertexAttrib3f", index, x, y, z)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4f(index uint, x, y, z, w float32) {
	p.Call("vertexAttrib4f", index, x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib1fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib2fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib3fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib4fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttribPointer(index uint, size int, typ GLenum, normalized bool, stride int, offset int) {
	p.Call("vertexAttribPointer", index, size, uint(typ), normalized, stride, offset)
}

func (p *webGLRenderingContextBaseImpl) Viewport(x, y, width, height int) {
	p.Call("viewport", x, y, width, height)
}

// -------------8<---------------------------------------

type webGLRenderingContextImpl struct {
	*webGLRenderingContextBaseImpl
}

func wrapWebGLRenderingContext(v Value) WebGLRenderingContext {
	if v.Valid() {
		return &webGLRenderingContextImpl{
			webGLRenderingContextBaseImpl: newWebGLRenderingContextBaseImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------
