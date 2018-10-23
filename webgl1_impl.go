// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type webGLObjectImpl struct {
	*objectImpl
}

func newWebGLObject(v js.Value) WebGLObject {
	if p := newWebGLObjectImpl(v); p != nil {
		return p
	}
	return nil
}

// KEEP
func newWebGLObjectImpl(v js.Value) *webGLObjectImpl {
	if isNil(v) {
		return nil
	}

	return &webGLObjectImpl{
		objectImpl: newObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLBufferImpl struct {
	*webGLObjectImpl
}

func newWebGLBuffer(v js.Value) WebGLBuffer {
	if isNil(v) {
		return nil
	}

	return &webGLBufferImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLFramebufferImpl struct {
	*webGLObjectImpl
}

func newWebGLFramebuffer(v js.Value) WebGLFramebuffer {
	if isNil(v) {
		return nil
	}

	return &webGLFramebufferImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLProgramImpl struct {
	*webGLObjectImpl
}

func newWebGLProgram(v js.Value) WebGLProgram {
	if isNil(v) {
		return nil
	}

	return &webGLProgramImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLRenderbufferImpl struct {
	*webGLObjectImpl
}

func newWebGLRenderbuffer(v js.Value) WebGLRenderbuffer {
	if isNil(v) {
		return nil
	}

	return &webGLRenderbufferImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLShaderImpl struct {
	*webGLObjectImpl
}

func newWebGLShader(v js.Value) WebGLShader {
	if isNil(v) {
		return nil
	}

	return &webGLShaderImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLTextureImpl struct {
	*webGLObjectImpl
}

func newWebGLTexture(v js.Value) WebGLTexture {
	if isNil(v) {
		return nil
	}

	return &webGLTextureImpl{
		webGLObjectImpl: newWebGLObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLUniformLocationImpl struct {
	*objectImpl
}

func newWebGLUniformLocation(v js.Value) WebGLUniformLocation {
	if isNil(v) {
		return nil
	}

	return &webGLUniformLocationImpl{
		objectImpl: newObjectImpl(v),
	}
}

// -------------8<---------------------------------------

type webGLActiveInfoImpl struct {
	*objectImpl
}

func newWebGLActiveInfo(v js.Value) WebGLActiveInfo {
	if isNil(v) {
		return nil
	}

	return &webGLActiveInfoImpl{
		objectImpl: newObjectImpl(v),
	}
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
	*objectImpl
}

func newWebGLShaderPrecisionFormat(v js.Value) WebGLShaderPrecisionFormat {
	if isNil(v) {
		return nil
	}

	return &webGLShaderPrecisionFormatImpl{
		objectImpl: newObjectImpl(v),
	}
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

func newWebGLContextAttributes(v js.Value) WebGLContextAttributes {
	if isNil(v) {
		return WebGLContextAttributes{}
	}

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

// -------------8<---------------------------------------

type texImageSourceImpl struct {
	*objectImpl
}

func newTexImageSource(v js.Value) TexImageSource {
	if isNil(v) {
		return nil
	}

	return &texImageSourceImpl{
		objectImpl: newObjectImpl(v),
	}
}

// -------------8<---------------------------------------

var _ WebGLRenderingContextBase = &webGLRenderingContextBaseImpl{}

type webGLRenderingContextBaseImpl struct {
	*objectImpl
}

// KEEP
func newWebGLRenderingContextBaseImpl(v js.Value) *webGLRenderingContextBaseImpl {
	if isNil(v) {
		return nil
	}

	return &webGLRenderingContextBaseImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *webGLRenderingContextBaseImpl) Canvas() HTMLCanvasElement {
	return newHTMLCanvasElement(p.Get("canvas"))
}

func (p *webGLRenderingContextBaseImpl) DrawingBufferWidth() int {
	return p.Get("drawingBufferWidth").Int()
}

func (p *webGLRenderingContextBaseImpl) drawingBufferHeight() int {
	return p.Get("drawingBufferHeight").Int()
}

func (p *webGLRenderingContextBaseImpl) GetContextAttributes() WebGLContextAttributes {
	return newWebGLContextAttributes(p.Call("getContextAttributes"))
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
	p.Call("attachShader", program.JSValue(), shader.JSValue())
}

func (p *webGLRenderingContextBaseImpl) BindAttribLocation(program WebGLProgram, index int, name string) {
	p.Call("bindAttribLocation", program.JSValue(), index, name)
}

func (p *webGLRenderingContextBaseImpl) BindBuffer(target GLenum, buffer WebGLBuffer) {
	p.Call("bindBuffer", uint(target), buffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) {
	p.Call("bindFramebuffer", uint(target), framebuffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) {
	p.Call("bindRenderbuffer", uint(target), renderbuffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) BindTexture(target GLenum, texture WebGLTexture) {
	p.Call("bindTexture", uint(target), texture.JSValue())
}

func (p *webGLRenderingContextBaseImpl) BlendColor(red, green, blue, alpha float64) {
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

func (p *webGLRenderingContextBaseImpl) BufferDataSource(target GLenum, data ArrayBufferView, usage GLenum) {
	p.Call("bufferData", uint(target), data.JSValue(), uint(usage))
}

func (p *webGLRenderingContextBaseImpl) BufferSubData(target GLenum, offset int, data ArrayBufferView) {
	p.Call("bufferSubData", uint(target), offset, data)
}

func (p *webGLRenderingContextBaseImpl) CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(p.Call("checkFramebufferStatus", uint(target)).Int())
}

func (p *webGLRenderingContextBaseImpl) Clear(mask uint) {
	p.Call("clear", mask)
}

func (p *webGLRenderingContextBaseImpl) ClearColor(red, green, blue, alpha float64) {
	p.Call("clearColor", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) ClearDepth(depth float64) {
	p.Call("clearDepth", depth)
}

func (p *webGLRenderingContextBaseImpl) ClearStencil(s int) {
	p.Call("clearStencil", s)
}

func (p *webGLRenderingContextBaseImpl) ColorMask(red, green, blue, alpha bool) {
	p.Call("colorMask", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) CompileShader(shader WebGLShader) {
	p.Call("compileShader", shader.JSValue())
}

func (p *webGLRenderingContextBaseImpl) CompressedTexImage2D(target GLenum, level int, internalFormat GLenum, width int, height int, border int, data ArrayBufferView) {
	p.Call("compressedTexImage2D", uint(target), level, uint(internalFormat), width, height, border, data.JSValue())
}

func (p *webGLRenderingContextBaseImpl) CompressedTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, data ArrayBufferView) {
	p.Call("compressedTexSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), data.JSValue())
}

func (p *webGLRenderingContextBaseImpl) CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	p.Call("copyTexImage2D", uint(target), level, uint(internalFormat), x, y, width, height, border)
}

func (p *webGLRenderingContextBaseImpl) CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	p.Call("copyTexSubImage2D", uint(target), level, xoffset, yoffset, x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) CreateBuffer() WebGLBuffer {
	return newWebGLBuffer(p.Call("createBuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateFramebuffer() WebGLFramebuffer {
	return newWebGLFramebuffer(p.Call("createFramebuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateProgram() WebGLProgram {
	return newWebGLProgram(p.Call("createProgram"))
}

func (p *webGLRenderingContextBaseImpl) CreateRenderbuffer() WebGLRenderbuffer {
	return newWebGLRenderbuffer(p.Call("createRenderbuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateShader(typ GLenum) WebGLShader {
	return newWebGLShader(p.Call("createShader", uint(typ)))
}

func (p *webGLRenderingContextBaseImpl) CreateTexture() WebGLTexture {
	return newWebGLTexture(p.Call("createTexture"))
}

func (p *webGLRenderingContextBaseImpl) CullFace(mode GLenum) {
	p.Call("cullFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) DeleteBuffer(buffer WebGLBuffer) {
	p.Call("deleteBuffer", buffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DeleteFramebuffer(frameBuffer WebGLFramebuffer) {
	p.Call("deleteFramebuffer", frameBuffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DeleteProgram(program WebGLProgram) {
	p.Call("deleteProgram", program.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) {
	p.Call("deleteRenderbuffer", renderbuffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DeleteShader(shader WebGLShader) {
	p.Call("deleteShader", shader.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DeleteTexture(texture WebGLTexture) {
	p.Call("deleteTexture", texture.JSValue())
}

func (p *webGLRenderingContextBaseImpl) DepthFunc(fn GLenum) {
	p.Call("depthFunc", uint(fn))
}

func (p *webGLRenderingContextBaseImpl) DepthMask(flag bool) {
	p.Call("depthMask", flag)
}

func (p *webGLRenderingContextBaseImpl) DepthRange(zNear, zFar float64) {
	p.Call("depthRange", zNear, zFar)
}

func (p *webGLRenderingContextBaseImpl) DetachShader(program WebGLProgram, shader WebGLShader) {
	p.Call("detachShader", program.JSValue(), shader.JSValue())
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
	p.Call("framebufferRenderbuffer", uint(target), uint(attachment), uint(renderbuffertarget), renderbuffer.JSValue())
}

func (p *webGLRenderingContextBaseImpl) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level int) {
	p.Call("framebufferTexture2D", uint(target), uint(attachment), uint(textarget), texture.JSValue(), level)
}

func (p *webGLRenderingContextBaseImpl) FrontFace(mode GLenum) {
	p.Call("frontFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) GenerateMipmap(target GLenum) {
	p.Call("generateMipmap", uint(target))
}

func (p *webGLRenderingContextBaseImpl) GetActiveAttrib(program WebGLProgram, index uint) WebGLActiveInfo {
	return newWebGLActiveInfo(p.Call("getActiveAttrib", program.JSValue(), index))
}

func (p *webGLRenderingContextBaseImpl) GetActiveUniform(program WebGLProgram, index uint) WebGLActiveInfo {
	return newWebGLActiveInfo(p.Call("getActiveUniform", program.JSValue(), index))
}

func (p *webGLRenderingContextBaseImpl) GetAttachedShaders(program WebGLProgram) []WebGLShader {
	s := arrayToSlice(p.Call("getAttachedShaders", program.JSValue()))
	if s == nil {
		return nil
	}

	ret := make([]WebGLShader, len(s))
	for i, v := range s {
		ret[i] = newWebGLShader(v)
	}

	return ret
}

func (p *webGLRenderingContextBaseImpl) GetAttribLocation(program WebGLProgram, name string) int {
	return p.Call("getAttribLocation", program.JSValue(), name).Int()
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
	return Wrap(p.Call("getProgramParameter", program.JSValue(), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetProgramInfoLog(program WebGLProgram) string {
	return p.Call("getProgramInfoLog", program.JSValue()).String()
}

func (p *webGLRenderingContextBaseImpl) GetRenderbufferParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.Call("getRenderbufferParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderParameter(shader WebGLShader, pname GLenum) interface{} {
	return Wrap(p.Call("getShaderParameter", shader.JSValue(), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderPrecisionFormat(shaderType GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat {
	return newWebGLShaderPrecisionFormat(p.Call("getShaderPrecisionFormat", uint(shaderType), uint(precisiontype)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderInfoLog(shader WebGLShader) string {
	return p.Call("getShaderInfoLog", shader.JSValue()).String()
}

func (p *webGLRenderingContextBaseImpl) GetShaderSource(shader WebGLShader) string {
	return p.Call("getShaderSource", shader.JSValue()).String()
}

func (p *webGLRenderingContextBaseImpl) GetTexParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.Call("getTexParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetUniform(program WebGLProgram, location WebGLUniformLocation) interface{} {
	return Wrap(p.Call("getUniform", program.JSValue(), location.JSValue()))
}

func (p *webGLRenderingContextBaseImpl) GetUniformLocation(program WebGLProgram, name string) WebGLUniformLocation {
	return newWebGLUniformLocation(p.Call("getUniformLocation", program.JSValue(), name))
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
	return p.Call("isBuffer", buffer.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsEnabled(c GLenum) bool {
	return p.Call("isEnabled", uint(c)).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsFramebuffer(buffer WebGLFramebuffer) bool {
	return p.Call("isFramebuffer", buffer.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsProgram(program WebGLProgram) bool {
	return p.Call("isProgram", program.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsRenderbuffer(buffer WebGLRenderbuffer) bool {
	return p.Call("isRenderbuffer", buffer.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsShader(shader WebGLShader) bool {
	return p.Call("isShader", shader.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) IsTexture(texture WebGLTexture) bool {
	return p.Call("isTexture", texture.JSValue()).Bool()
}

func (p *webGLRenderingContextBaseImpl) LineWidth(width float64) {
	p.Call("lineWidth", width)
}

func (p *webGLRenderingContextBaseImpl) LinkProgram(program WebGLProgram) {
	p.Call("linkProgram", program.JSValue())
}

func (p *webGLRenderingContextBaseImpl) PixelStorei(pname GLenum, param int) {
	p.Call("pixelStorei", uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) PolygonOffset(factor float64, units float64) {
	p.Call("polygonOffset", factor, units)
}

func (p *webGLRenderingContextBaseImpl) ReadPixels(x int, y int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.Call("readPixels", x, y, width, height, uint(format), uint(typ), pixels.JSValue())
}

func (p *webGLRenderingContextBaseImpl) RenderbufferStorage(target GLenum, format GLenum, width int, height int) {
	p.Call("renderbufferStorage", uint(target), uint(format), width, height)
}

func (p *webGLRenderingContextBaseImpl) SampleCoverage(value float64, invert bool) {
	p.Call("sampleCoverage", value, invert)
}

func (p *webGLRenderingContextBaseImpl) Scissor(x int, y int, width int, height int) {
	p.Call("scissor", x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) ShaderSource(shader WebGLShader, source string) {
	p.Call("shaderSource", shader.JSValue(), source)
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
	p.Call("texImage2D", uint(target), level, internalFormat, width, height, border, uint(format), uint(typ), pixels.JSValue())
}

func (p *webGLRenderingContextBaseImpl) TexImage2DSource(target GLenum, level int, internalFormat int, format GLenum, typ GLenum, source TexImageSource) {
	p.Call("texImage2D", uint(target), level, internalFormat, uint(format), uint(typ), source.JSValue())
}

func (p *webGLRenderingContextBaseImpl) TexParameterf(target GLenum, pname GLenum, param float64) {
	p.Call("texParameterf", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexParameteri(target GLenum, pname GLenum, param int) {
	p.Call("texParameteri", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DBuffer(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.Call("texSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), uint(typ), pixels.JSValue())
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DSource(target GLenum, level int, xoffset int, yoffset int, format GLenum, typ GLenum, source TexImageSource) {
	p.Call("texSubImage2D", uint(target), level, xoffset, yoffset, uint(format), uint(typ), source.JSValue())
}

func (p *webGLRenderingContextBaseImpl) Uniform1f(location WebGLUniformLocation, x float64) {
	p.Call("uniform1f", location.JSValue(), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2f(location WebGLUniformLocation, x float64, y float64) {
	p.Call("uniform2f", location.JSValue(), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3f(location WebGLUniformLocation, x float64, y float64, z float64) {
	p.Call("uniform3f", location.JSValue(), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4f(location WebGLUniformLocation, x float64, y float64, z float64, w float64) {
	p.Call("uniform4f", location.JSValue(), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1i(location WebGLUniformLocation, x int) {
	p.Call("uniform1i", location.JSValue(), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2i(location WebGLUniformLocation, x int, y int) {
	p.Call("uniform2i", location.JSValue(), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3i(location WebGLUniformLocation, x int, y int, z int) {
	p.Call("uniform3i", location.JSValue(), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4i(location WebGLUniformLocation, x int, y int, z int, w int) {
	p.Call("uniform4i", location.JSValue(), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1fv(location WebGLUniformLocation, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform1fv", location.JSValue(), v)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2fv(location WebGLUniformLocation, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform2fv", location.JSValue(), v)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3fv(location WebGLUniformLocation, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform3fv", location.JSValue(), v)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4fv(location WebGLUniformLocation, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform4fv", location.JSValue(), v)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform1iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform1iv", location.JSValue(), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform2iv", location.JSValue(), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform3iv", location.JSValue(), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.Call("uniform4iv", location.JSValue(), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix2fv(location WebGLUniformLocation, transpose bool, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix2fv", location.JSValue(), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix3fv(location WebGLUniformLocation, transpose bool, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix3fv", location.JSValue(), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix4fv(location WebGLUniformLocation, transpose bool, v []float64) {
	ta := js.TypedArrayOf(v)
	p.Call("uniformMatrix4fv", location.JSValue(), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UseProgram(program WebGLProgram) {
	p.Call("useProgram", program.JSValue())
}

func (p *webGLRenderingContextBaseImpl) ValidateProgram(program WebGLProgram) {
	p.Call("validateProgram", program.JSValue())
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1f(index uint, x float64) {
	p.Call("vertexAttrib1f", index, x)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2f(index uint, x, y float64) {
	p.Call("vertexAttrib2f", index, x, y)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3f(index uint, x, y, z float64) {
	p.Call("vertexAttrib3f", index, x, y, z)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4f(index uint, x, y, z, w float64) {
	p.Call("vertexAttrib4f", index, x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1fv(index uint, values []float64) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib1fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2fv(index uint, values []float64) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib2fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3fv(index uint, values []float64) {
	ta := js.TypedArrayOf(values)
	p.Call("vertexAttrib3fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4fv(index uint, values []float64) {
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

func newWebGLRenderingContext(v js.Value) WebGLRenderingContext {
	if isNil(v) {
		return nil
	}

	return &webGLRenderingContextImpl{
		webGLRenderingContextBaseImpl: newWebGLRenderingContextBaseImpl(v),
	}
}

// -------------8<---------------------------------------
