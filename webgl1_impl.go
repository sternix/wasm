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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
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
	if v.valid() {
		return &webGLActiveInfoImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLActiveInfoImpl) Size() int {
	return p.get("size").toInt()
}

func (p *webGLActiveInfoImpl) Type() GLenum {
	return GLenum(p.get("type").toInt())
}

func (p *webGLActiveInfoImpl) Name() string {
	return p.get("name").toString()
}

// -------------8<---------------------------------------

type webGLShaderPrecisionFormatImpl struct {
	Value
}

func wrapWebGLShaderPrecisionFormat(v Value) WebGLShaderPrecisionFormat {
	if v.valid() {
		return &webGLShaderPrecisionFormatImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLShaderPrecisionFormatImpl) RangeMin() int {
	return p.get("rangeMin").toInt()
}

func (p *webGLShaderPrecisionFormatImpl) RangeMax() int {
	return p.get("rangeMax").toInt()
}

func (p *webGLShaderPrecisionFormatImpl) Precision() int {
	return p.get("precision").toInt()
}

// -------------8<---------------------------------------

func wrapWebGLContextAttributes(v Value) WebGLContextAttributes {
	if v.valid() {
		return WebGLContextAttributes{
			Alpha:                        v.get("alpha").toBool(),
			Depth:                        v.get("depth").toBool(),
			Stencil:                      v.get("stencil").toBool(),
			Antialias:                    v.get("antialias").toBool(),
			PremultipliedAlpha:           v.get("premultipliedAlpha").toBool(),
			PreserveDrawingBuffer:        v.get("preserveDrawingBuffer").toBool(),
			PowerPreference:              WebGLPowerPreference(v.get("powerPreference").toString()),
			FailIfMajorPerformanceCaveat: v.get("failIfMajorPerformanceCaveat").toBool(),
		}
	}
	return WebGLContextAttributes{}
}

// -------------8<---------------------------------------

type texImageSourceImpl struct {
	Value
}

func wrapTexImageSource(v Value) TexImageSource {
	if v.valid() {
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
	if v.valid() {
		return &webGLRenderingContextBaseImpl{
			Value: v,
		}
	}
	return nil
}

func (p *webGLRenderingContextBaseImpl) Canvas() HTMLCanvasElement {
	return wrapHTMLCanvasElement(p.get("canvas"))
}

func (p *webGLRenderingContextBaseImpl) DrawingBufferWidth() int {
	return p.get("drawingBufferWidth").toInt()
}

func (p *webGLRenderingContextBaseImpl) drawingBufferHeight() int {
	return p.get("drawingBufferHeight").toInt()
}

func (p *webGLRenderingContextBaseImpl) GetContextAttributes() WebGLContextAttributes {
	return wrapWebGLContextAttributes(p.call("getContextAttributes"))
}

func (p *webGLRenderingContextBaseImpl) IsContextLost() bool {
	return p.call("isContextLost").toBool()
}

func (p *webGLRenderingContextBaseImpl) GetSupportedExtensions() []string {
	return stringSequenceToSlice(p.call("getSupportedExtensions"))
}

func (p *webGLRenderingContextBaseImpl) GetExtension(name string) interface{} {
	return Wrap(p.call("getExtension", name))
}

func (p *webGLRenderingContextBaseImpl) ActiveTexture(texture GLenum) {
	p.call("activeTexture", uint(texture))
}

func (p *webGLRenderingContextBaseImpl) AttachShader(program WebGLProgram, shader WebGLShader) {
	p.call("attachShader", JSValue(program), JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) BindAttribLocation(program WebGLProgram, index int, name string) {
	p.call("bindAttribLocation", JSValue(program), index, name)
}

func (p *webGLRenderingContextBaseImpl) BindBuffer(target GLenum, buffer WebGLBuffer) {
	p.call("bindBuffer", uint(target), JSValue(buffer))
}

func (p *webGLRenderingContextBaseImpl) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) {
	p.call("bindFramebuffer", uint(target), JSValue(framebuffer))
}

func (p *webGLRenderingContextBaseImpl) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) {
	p.call("bindRenderbuffer", uint(target), JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) BindTexture(target GLenum, texture WebGLTexture) {
	p.call("bindTexture", uint(target), JSValue(texture))
}

func (p *webGLRenderingContextBaseImpl) BlendColor(red, green, blue, alpha float32) {
	p.call("blendColor", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) BlendEquation(mode GLenum) {
	p.call("blendEquation", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) {
	p.call("blendEquationSeparate", uint(modeRGB), uint(modeAlpha))
}

func (p *webGLRenderingContextBaseImpl) BlendFunc(sfactor GLenum, dfactor GLenum) {
	p.call("blendFunc", uint(sfactor), uint(dfactor))
}

func (p *webGLRenderingContextBaseImpl) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLenum) {
	p.call("blendFuncSeparate", uint(srcRGB), uint(dstRGB), uint(srcAlpha), uint(dstAlpha))
}

func (p *webGLRenderingContextBaseImpl) BufferData(target GLenum, size int, usage GLenum) {
	p.call("bufferData", uint(target), size, uint(usage))
}

func (p *webGLRenderingContextBaseImpl) BufferDataSource(target GLenum, data BufferSource, usage GLenum) {
	p.call("bufferData", uint(target), JSValue(data), uint(usage))
}

func (p *webGLRenderingContextBaseImpl) BufferSubData(target GLenum, offset int, data BufferSource) {
	p.call("bufferSubData", uint(target), offset, JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(p.call("checkFramebufferStatus", uint(target)).toInt())
}

func (p *webGLRenderingContextBaseImpl) Clear(mask GLenum) {
	p.call("clear", uint(mask))
}

func (p *webGLRenderingContextBaseImpl) ClearColor(red, green, blue, alpha float32) {
	p.call("clearColor", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) ClearDepth(depth float32) {
	p.call("clearDepth", depth)
}

func (p *webGLRenderingContextBaseImpl) ClearStencil(s int) {
	p.call("clearStencil", s)
}

func (p *webGLRenderingContextBaseImpl) ColorMask(red, green, blue, alpha bool) {
	p.call("colorMask", red, green, blue, alpha)
}

func (p *webGLRenderingContextBaseImpl) CompileShader(shader WebGLShader) {
	p.call("compileShader", JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) CompressedTexImage2D(target GLenum, level int, internalFormat GLenum, width int, height int, border int, data ArrayBufferView) {
	p.call("compressedTexImage2D", uint(target), level, uint(internalFormat), width, height, border, JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CompressedTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, data ArrayBufferView) {
	p.call("compressedTexSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), JSValue(data))
}

func (p *webGLRenderingContextBaseImpl) CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	p.call("copyTexImage2D", uint(target), level, uint(internalFormat), x, y, width, height, border)
}

func (p *webGLRenderingContextBaseImpl) CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	p.call("copyTexSubImage2D", uint(target), level, xoffset, yoffset, x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) CreateBuffer() WebGLBuffer {
	return wrapWebGLBuffer(p.call("createBuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateFramebuffer() WebGLFramebuffer {
	return wrapWebGLFramebuffer(p.call("createFramebuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateProgram() WebGLProgram {
	return wrapWebGLProgram(p.call("createProgram"))
}

func (p *webGLRenderingContextBaseImpl) CreateRenderbuffer() WebGLRenderbuffer {
	return wrapWebGLRenderbuffer(p.call("createRenderbuffer"))
}

func (p *webGLRenderingContextBaseImpl) CreateShader(typ GLenum) WebGLShader {
	return wrapWebGLShader(p.call("createShader", uint(typ)))
}

func (p *webGLRenderingContextBaseImpl) CreateTexture() WebGLTexture {
	return wrapWebGLTexture(p.call("createTexture"))
}

func (p *webGLRenderingContextBaseImpl) CullFace(mode GLenum) {
	p.call("cullFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) DeleteBuffer(buffer WebGLBuffer) {
	p.call("deleteBuffer", JSValue(buffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteFramebuffer(frameBuffer WebGLFramebuffer) {
	p.call("deleteFramebuffer", JSValue(frameBuffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteProgram(program WebGLProgram) {
	p.call("deleteProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) {
	p.call("deleteRenderbuffer", JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) DeleteShader(shader WebGLShader) {
	p.call("deleteShader", JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) DeleteTexture(texture WebGLTexture) {
	p.call("deleteTexture", JSValue(texture))
}

func (p *webGLRenderingContextBaseImpl) DepthFunc(fn GLenum) {
	p.call("depthFunc", uint(fn))
}

func (p *webGLRenderingContextBaseImpl) DepthMask(flag bool) {
	p.call("depthMask", flag)
}

func (p *webGLRenderingContextBaseImpl) DepthRange(zNear, zFar float32) {
	p.call("depthRange", zNear, zFar)
}

func (p *webGLRenderingContextBaseImpl) DetachShader(program WebGLProgram, shader WebGLShader) {
	p.call("detachShader", JSValue(program), JSValue(shader))
}

func (p *webGLRenderingContextBaseImpl) Disable(c GLenum) {
	p.call("disable", uint(c))
}

func (p *webGLRenderingContextBaseImpl) DisableVertexAttribArray(index uint) {
	p.call("disableVertexAttribArray", index)
}

func (p *webGLRenderingContextBaseImpl) DrawArrays(mode GLenum, first int, count int) {
	p.call("drawArrays", uint(mode), first, count)
}

func (p *webGLRenderingContextBaseImpl) DrawElements(mode GLenum, count int, typ GLenum, offset int) {
	p.call("drawElements", uint(mode), count, uint(typ), offset)
}

func (p *webGLRenderingContextBaseImpl) Enable(c GLenum) {
	p.call("enable", uint(c))
}

func (p *webGLRenderingContextBaseImpl) EnableVertexAttribArray(index uint) {
	p.call("enableVertexAttribArray", index)
}

func (p *webGLRenderingContextBaseImpl) Finish() {
	p.call("finish")
}

func (p *webGLRenderingContextBaseImpl) Flush() {
	p.call("flush")
}

func (p *webGLRenderingContextBaseImpl) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) {
	p.call("framebufferRenderbuffer", uint(target), uint(attachment), uint(renderbuffertarget), JSValue(renderbuffer))
}

func (p *webGLRenderingContextBaseImpl) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level int) {
	p.call("framebufferTexture2D", uint(target), uint(attachment), uint(textarget), JSValue(texture), level)
}

func (p *webGLRenderingContextBaseImpl) FrontFace(mode GLenum) {
	p.call("frontFace", uint(mode))
}

func (p *webGLRenderingContextBaseImpl) GenerateMipmap(target GLenum) {
	p.call("generateMipmap", uint(target))
}

func (p *webGLRenderingContextBaseImpl) GetActiveAttrib(program WebGLProgram, index uint) WebGLActiveInfo {
	return wrapWebGLActiveInfo(p.call("getActiveAttrib", JSValue(program), index))
}

func (p *webGLRenderingContextBaseImpl) GetActiveUniform(program WebGLProgram, index uint) WebGLActiveInfo {
	return wrapWebGLActiveInfo(p.call("getActiveUniform", JSValue(program), index))
}

func (p *webGLRenderingContextBaseImpl) GetAttachedShaders(program WebGLProgram) []WebGLShader {
	if s := p.call("getAttachedShaders", JSValue(program)).toSlice(); s != nil {
		ret := make([]WebGLShader, len(s))
		for i, v := range s {
			ret[i] = wrapWebGLShader(v)
		}
		return ret
	}
	return nil
}

func (p *webGLRenderingContextBaseImpl) GetAttribLocation(program WebGLProgram, name string) int {
	return p.call("getAttribLocation", JSValue(program), name).toInt()
}

func (p *webGLRenderingContextBaseImpl) GetBufferParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.call("getBufferParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetParameter(pname GLenum) interface{} {
	return Wrap(p.call("getParameter", uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetError() GLenum {
	return GLenum(p.call("getError").toInt())
}

func (p *webGLRenderingContextBaseImpl) GetFramebufferAttachmentParameter(target, attachment, pname GLenum) {
	p.call("getFramebufferAttachmentParameter", uint(target), uint(attachment), uint(pname))
}

func (p *webGLRenderingContextBaseImpl) GetProgramParameter(program WebGLProgram, pname GLenum) interface{} {
	return Wrap(p.call("getProgramParameter", JSValue(program), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetProgramInfoLog(program WebGLProgram) string {
	return p.call("getProgramInfoLog", JSValue(program)).toString()
}

func (p *webGLRenderingContextBaseImpl) GetRenderbufferParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.call("getRenderbufferParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderParameter(shader WebGLShader, pname GLenum) interface{} {
	return Wrap(p.call("getShaderParameter", JSValue(shader), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderPrecisionFormat(shaderType GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat {
	return wrapWebGLShaderPrecisionFormat(p.call("getShaderPrecisionFormat", uint(shaderType), uint(precisiontype)))
}

func (p *webGLRenderingContextBaseImpl) GetShaderInfoLog(shader WebGLShader) string {
	return p.call("getShaderInfoLog", JSValue(shader)).toString()
}

func (p *webGLRenderingContextBaseImpl) GetShaderSource(shader WebGLShader) string {
	return p.call("getShaderSource", JSValue(shader)).toString()
}

func (p *webGLRenderingContextBaseImpl) GetTexParameter(target GLenum, pname GLenum) interface{} {
	return Wrap(p.call("getTexParameter", uint(target), uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetUniform(program WebGLProgram, location WebGLUniformLocation) interface{} {
	return Wrap(p.call("getUniform", JSValue(program), JSValue(location)))
}

func (p *webGLRenderingContextBaseImpl) GetUniformLocation(program WebGLProgram, name string) WebGLUniformLocation {
	return wrapWebGLUniformLocation(p.call("getUniformLocation", JSValue(program), name))
}

func (p *webGLRenderingContextBaseImpl) GetVertexAttrib(index uint, pname GLenum) interface{} {
	return Wrap(p.call("getVertexAttrib", index, uint(pname)))
}

func (p *webGLRenderingContextBaseImpl) GetVertexAttribOffset(index uint, pname GLenum) int {
	return p.call("getVertexAttribOffset", index, uint(pname)).toInt()
}

func (p *webGLRenderingContextBaseImpl) Hint(target GLenum, mode GLenum) {
	p.call("hint", uint(target), uint(mode))
}

func (p *webGLRenderingContextBaseImpl) IsBuffer(buffer WebGLBuffer) bool {
	return p.call("isBuffer", JSValue(buffer)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsEnabled(c GLenum) bool {
	return p.call("isEnabled", uint(c)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsFramebuffer(buffer WebGLFramebuffer) bool {
	return p.call("isFramebuffer", JSValue(buffer)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsProgram(program WebGLProgram) bool {
	return p.call("isProgram", JSValue(program)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsRenderbuffer(buffer WebGLRenderbuffer) bool {
	return p.call("isRenderbuffer", JSValue(buffer)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsShader(shader WebGLShader) bool {
	return p.call("isShader", JSValue(shader)).toBool()
}

func (p *webGLRenderingContextBaseImpl) IsTexture(texture WebGLTexture) bool {
	return p.call("isTexture", JSValue(texture)).toBool()
}

func (p *webGLRenderingContextBaseImpl) LineWidth(width float32) {
	p.call("lineWidth", width)
}

func (p *webGLRenderingContextBaseImpl) LinkProgram(program WebGLProgram) {
	p.call("linkProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) PixelStorei(pname GLenum, param int) {
	p.call("pixelStorei", uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) PolygonOffset(factor float32, units float32) {
	p.call("polygonOffset", factor, units)
}

func (p *webGLRenderingContextBaseImpl) ReadPixels(x int, y int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.call("readPixels", x, y, width, height, uint(format), uint(typ), JSValue(pixels))
}

func (p *webGLRenderingContextBaseImpl) RenderbufferStorage(target GLenum, format GLenum, width int, height int) {
	p.call("renderbufferStorage", uint(target), uint(format), width, height)
}

func (p *webGLRenderingContextBaseImpl) SampleCoverage(value float32, invert bool) {
	p.call("sampleCoverage", value, invert)
}

func (p *webGLRenderingContextBaseImpl) Scissor(x int, y int, width int, height int) {
	p.call("scissor", x, y, width, height)
}

func (p *webGLRenderingContextBaseImpl) ShaderSource(shader WebGLShader, source string) {
	p.call("shaderSource", JSValue(shader), source)
}

func (p *webGLRenderingContextBaseImpl) StencilFunc(fn GLenum, ref int, mask uint) {
	p.call("stencilFunc", uint(fn), ref, mask)
}

func (p *webGLRenderingContextBaseImpl) StencilFuncSeparate(face GLenum, fn GLenum, ref int, mask uint) {
	p.call("stencilFuncSeparate", uint(face), uint(fn), ref, mask)
}

func (p *webGLRenderingContextBaseImpl) StencilMask(mask uint) {
	p.call("stencilMask", mask)
}

func (p *webGLRenderingContextBaseImpl) StencilMaskSeparate(face GLenum, mask uint) {
	p.call("stencilMaskSeparate", uint(face), mask)
}

func (p *webGLRenderingContextBaseImpl) StencilOp(fail, zfail, zpass GLenum) {
	p.call("stencilOp", uint(fail), uint(zfail), uint(zpass))
}

func (p *webGLRenderingContextBaseImpl) StencilOpSeparate(face, fail, zfail, zpass GLenum) {
	p.call("stencilOpSeparate", face, fail, zfail, zpass)
}

func (p *webGLRenderingContextBaseImpl) TexImage2DBuffer(target GLenum, level int, internalFormat int, width int, height int, border int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.call("texImage2D", uint(target), level, internalFormat, width, height, border, uint(format), uint(typ), JSValue(pixels))
}

func (p *webGLRenderingContextBaseImpl) TexImage2DSource(target GLenum, level int, internalFormat int, format GLenum, typ GLenum, source TexImageSource) {
	p.call("texImage2D", uint(target), level, internalFormat, uint(format), uint(typ), JSValue(source))
}

func (p *webGLRenderingContextBaseImpl) TexParameterf(target GLenum, pname GLenum, param float32) {
	p.call("texParameterf", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexParameteri(target GLenum, pname GLenum, param int) {
	p.call("texParameteri", uint(target), uint(pname), param)
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DBuffer(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels ArrayBufferView) {
	p.call("texSubImage2D", uint(target), level, xoffset, yoffset, width, height, uint(format), uint(typ), JSValue(pixels))
}

func (p *webGLRenderingContextBaseImpl) TexSubImage2DSource(target GLenum, level int, xoffset int, yoffset int, format GLenum, typ GLenum, source TexImageSource) {
	p.call("texSubImage2D", uint(target), level, xoffset, yoffset, uint(format), uint(typ), JSValue(source))
}

func (p *webGLRenderingContextBaseImpl) Uniform1f(location WebGLUniformLocation, x float32) {
	p.call("uniform1f", JSValue(location), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2f(location WebGLUniformLocation, x float32, y float32) {
	p.call("uniform2f", JSValue(location), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3f(location WebGLUniformLocation, x float32, y float32, z float32) {
	p.call("uniform3f", JSValue(location), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4f(location WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	p.call("uniform4f", JSValue(location), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1i(location WebGLUniformLocation, x int) {
	p.call("uniform1i", JSValue(location), x)
}

func (p *webGLRenderingContextBaseImpl) Uniform2i(location WebGLUniformLocation, x int, y int) {
	p.call("uniform2i", JSValue(location), x, y)
}

func (p *webGLRenderingContextBaseImpl) Uniform3i(location WebGLUniformLocation, x int, y int, z int) {
	p.call("uniform3i", JSValue(location), x, y, z)
}

func (p *webGLRenderingContextBaseImpl) Uniform4i(location WebGLUniformLocation, x int, y int, z int, w int) {
	p.call("uniform4i", JSValue(location), x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) Uniform1fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniform1fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniform2fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniform3fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4fv(location WebGLUniformLocation, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniform4fv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform1iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.call("uniform1iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform2iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.call("uniform2iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform3iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.call("uniform3iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) Uniform4iv(location WebGLUniformLocation, v []int) {
	ta := js.TypedArrayOf(v)
	p.call("uniform4iv", JSValue(location), ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix2fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniformMatrix2fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix3fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniformMatrix3fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UniformMatrix4fv(location WebGLUniformLocation, transpose bool, v []float32) {
	ta := js.TypedArrayOf(v)
	p.call("uniformMatrix4fv", JSValue(location), transpose, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) UseProgram(program WebGLProgram) {
	p.call("useProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) ValidateProgram(program WebGLProgram) {
	p.call("validateProgram", JSValue(program))
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1f(index uint, x float32) {
	p.call("vertexAttrib1f", index, x)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2f(index uint, x, y float32) {
	p.call("vertexAttrib2f", index, x, y)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3f(index uint, x, y, z float32) {
	p.call("vertexAttrib3f", index, x, y, z)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4f(index uint, x, y, z, w float32) {
	p.call("vertexAttrib4f", index, x, y, z, w)
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib1fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.call("vertexAttrib1fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib2fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.call("vertexAttrib2fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib3fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.call("vertexAttrib3fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttrib4fv(index uint, values []float32) {
	ta := js.TypedArrayOf(values)
	p.call("vertexAttrib4fv", index, ta)
	ta.Release()
}

func (p *webGLRenderingContextBaseImpl) VertexAttribPointer(index uint, size int, typ GLenum, normalized bool, stride int, offset int) {
	p.call("vertexAttribPointer", index, size, uint(typ), normalized, stride, offset)
}

func (p *webGLRenderingContextBaseImpl) Viewport(x, y, width, height int) {
	p.call("viewport", x, y, width, height)
}

// -------------8<---------------------------------------

type webGLRenderingContextImpl struct {
	*webGLRenderingContextBaseImpl
}

func wrapWebGLRenderingContext(v Value) WebGLRenderingContext {
	if v.valid() {
		return &webGLRenderingContextImpl{
			webGLRenderingContextBaseImpl: newWebGLRenderingContextBaseImpl(v),
		}
	}
	return nil
}

// -------------8<---------------------------------------
