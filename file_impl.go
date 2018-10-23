// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type fileImpl struct {
	*objectImpl
}

func newFile(v js.Value) File {
	if isNil(v) {
		return nil
	}

	return &fileImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *fileImpl) Name() string {
	return p.Get("name").String()
}

func (p *fileImpl) LastModified() int {
	return p.Get("lastModified").Int()
}

// -------------8<---------------------------------------

type blobImpl struct {
	*objectImpl
}

func newBlob(v js.Value) Blob {
	if isNil(v) {
		return nil
	}
	return &blobImpl{
		objectImpl: newObjectImpl(v),
	}
}

func (p *blobImpl) Size() int {
	return p.Get("size").Int()
}

func (p *blobImpl) Type() string {
	return p.Get("type").String()
}

func (p *blobImpl) Slice(args ...interface{}) Blob {
	switch len(args) {
	case 1:
		if start, ok := args[0].(int); ok {
			return newBlob(p.Call("slice", start))
		}
	case 2:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				return newBlob(p.Call("slice", start, end))
			}
		}
	case 3:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				if contentType, ok := args[2].(string); ok {
					return newBlob(p.Call("slice", start, end, contentType))
				}
			}
		}
	}
	// wrong parameter count or parameter not given
	return newBlob(p.Call("slice"))
}

// -------------8<---------------------------------------

type fileReaderImpl struct {
	*eventTargetImpl
}

func newFileReader(v js.Value) FileReader {
	if isNil(v) {
		return nil
	}

	return &fileReaderImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *fileReaderImpl) ReadAsArrayBuffer(blob Blob) {
	p.Call("readAsArrayBuffer", blob.JSValue())
}

func (p *fileReaderImpl) ReadAsBinaryString(blob Blob) {
	p.Call("readAsBinaryString", blob.JSValue())
}

func (p *fileReaderImpl) ReadAsText(blob Blob, label ...string) {
	switch len(label) {
	case 0:
		p.Call("readAsText", blob.JSValue())
	default: // 1 or more
		p.Call("readAsText", blob.JSValue(), label[0])
	}
}

func (p *fileReaderImpl) ReadAsDataURL(blob Blob) {
	p.Call("readAsDataURL", blob.JSValue())
}

func (p *fileReaderImpl) Abort() {
	p.Call("abort")
}

func (p *fileReaderImpl) ReadyState() FileReaderState {
	return FileReaderState(p.Get("readyState").Int())
}

// TODO
func (p *fileReaderImpl) Result() []byte {
	v := p.Get("result")

	switch v.Type() {
	case js.TypeString:
		return []byte(v.String())
	case js.TypeObject: // ArrayBuffer
		return newArrayBuffer(v).ToByteSlice()
	default:
		return nil
	}
}

func (p *fileReaderImpl) Error() DOMException {
	return newDOMException(p.Get("error"))
}

func (p *fileReaderImpl) OnLoadStart(fn func(Event)) EventHandler {
	return p.On("loadstart", fn)
}

func (p *fileReaderImpl) OnProgress(fn func(Event)) EventHandler {
	return p.On("progress", fn)
}

func (p *fileReaderImpl) OnLoad(fn func(Event)) EventHandler {
	return p.On("load", fn)
}

func (p *fileReaderImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

func (p *fileReaderImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

func (p *fileReaderImpl) OnLoadEnd(fn func(Event)) EventHandler {
	return p.On("loadend", fn)
}

// -------------8<---------------------------------------

type fileReaderSyncImpl struct {
	js.Value
}

func newFileReaderSync(v js.Value) FileReaderSync {
	if isNil(v) {
		return nil
	}

	return &fileReaderSyncImpl{
		Value: v,
	}
}

func (p *fileReaderSyncImpl) ReadAsArrayBuffer(blob Blob) ArrayBuffer {
	return newArrayBuffer(p.Call("readAsArrayBuffer", blob.JSValue()))
}

func (p *fileReaderSyncImpl) ReadAsBinaryString(blob Blob) string {
	return p.Call("readAsBinaryString", blob.JSValue()).String()
}

func (p *fileReaderSyncImpl) ReadAsText(blob Blob, label ...string) string {
	switch len(label) {
	case 0:
		return p.Call("readAsText", blob.JSValue()).String()
	default:
		return p.Call("readAsText", blob.JSValue(), label[0]).String()
	}
}

func (p *fileReaderSyncImpl) ReadAsDataURL(blob Blob) string {
	return p.Call("readAsDataURL", blob.JSValue()).String()
}

// -------------8<---------------------------------------

type progressEventImpl struct {
	*eventImpl
}

func newProgressEvent(v js.Value) ProgressEvent {
	if isNil(v) {
		return nil
	}

	return &progressEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *progressEventImpl) LengthComputable() bool {
	return p.Get("lengthComputable").Bool()
}

func (p *progressEventImpl) Loaded() int {
	return p.Get("loaded").Int()
}

func (p *progressEventImpl) Total() int {
	return p.Get("total").Int()
}

// -------------8<---------------------------------------

func NewBlob(args ...interface{}) Blob {
	jsBlob := js.Global().Get("Blob")
	if isNil(jsBlob) {
		return nil
	}

	switch len(args) {
	case 1:
		if ar, ok := args[0].([]byte); ok {
			ta := js.TypedArrayOf(ar)
			defer ta.Release()
			return newBlob(jsBlob.New(ta))
		}
	case 2:
		if ar, ok := args[0].([]byte); ok {
			if options, ok := args[1].(BlobPropertyBag); ok {
				ta := js.TypedArrayOf(ar)
				defer ta.Release()
				return newBlob(jsBlob.New(ta, toJSONObject(options)))
			}
		}
	}

	return newBlob(jsBlob.New())
}

func NewFile(fileBits []byte, fileName string, options ...FilePropertyBag) File {
	jsFile := js.Global().Get("File")
	if isNil(jsFile) {
		return nil
	}

	ta := js.TypedArrayOf(fileBits)
	defer ta.Release()

	switch len(options) {
	case 0:
		return newFile(jsFile.New(ta, fileName))
	default:
		return newFile(jsFile.New(ta, fileName, toJSONObject(options[0])))
	}
}

func NewFileReader() FileReader {
	jsFileReader := js.Global().Get("FileReader")
	if isNil(jsFileReader) {
		return nil
	}

	return newFileReader(jsFileReader.New())
}

func NewFileReaderSync() FileReaderSync {
	jsFileReaderSync := js.Global().Get("FileReaderSync")
	if isNil(jsFileReaderSync) {
		return nil
	}

	return newFileReaderSync(jsFileReaderSync.New())
}

func NewProgressEvent(typ string, pei ...ProgressEventInit) ProgressEvent {
	jsProgressEvent := js.Global().Get("ProgressEvent")
	if isNil(jsProgressEvent) {
		return nil
	}

	if len(pei) > 0 {
		return newProgressEvent(jsProgressEvent.New(typ, toJSONObject(pei[0])))
	}
	return newProgressEvent(jsProgressEvent.New(typ))
}
