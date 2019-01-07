// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type fileImpl struct {
	Value
}

func wrapFile(v Value) File {
	if v.Valid() {
		return &fileImpl{
			Value: v,
		}
	}
	return nil
}

func (p *fileImpl) Name() string {
	return p.Get("name").String()
}

func (p *fileImpl) LastModified() int {
	return p.Get("lastModified").Int()
}

// -------------8<---------------------------------------

type blobImpl struct {
	Value
}

func wrapBlob(v Value) Blob {
	if v.Valid() {
		return &blobImpl{
			Value: v,
		}
	}
	return nil
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
			return wrapBlob(p.Call("slice", start))
		}
	case 2:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				return wrapBlob(p.Call("slice", start, end))
			}
		}
	case 3:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				if contentType, ok := args[2].(string); ok {
					return wrapBlob(p.Call("slice", start, end, contentType))
				}
			}
		}
	}
	// wrong parameter count or parameter not given
	return wrapBlob(p.Call("slice"))
}

// -------------8<---------------------------------------

type fileReaderImpl struct {
	*eventTargetImpl
}

func wrapFileReader(v Value) FileReader {
	if v.Valid() {
		return &fileReaderImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *fileReaderImpl) ReadAsArrayBuffer(blob Blob) {
	p.Call("readAsArrayBuffer", JSValue(blob))
}

func (p *fileReaderImpl) ReadAsBinaryString(blob Blob) {
	p.Call("readAsBinaryString", JSValue(blob))
}

func (p *fileReaderImpl) ReadAsText(blob Blob, label ...string) {
	switch len(label) {
	case 0:
		p.Call("readAsText", JSValue(blob))
	default: // 1 or more
		p.Call("readAsText", JSValue(blob), label[0])
	}
}

func (p *fileReaderImpl) ReadAsDataURL(blob Blob) {
	p.Call("readAsDataURL", JSValue(blob))
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
	case TypeString:
		return []byte(v.String())
	case TypeObject: // ArrayBuffer
		return wrapArrayBuffer(v).ToByteSlice()
	default:
		return nil
	}
}

func (p *fileReaderImpl) Error() DOMException {
	return wrapDOMException(p.Get("error"))
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
	Value
}

func wrapFileReaderSync(v Value) FileReaderSync {
	if v.Valid() {
		return &fileReaderSyncImpl{
			Value: v,
		}
	}
	return nil
}

func (p *fileReaderSyncImpl) ReadAsArrayBuffer(blob Blob) ArrayBuffer {
	return wrapArrayBuffer(p.Call("readAsArrayBuffer", JSValue(blob)))
}

func (p *fileReaderSyncImpl) ReadAsBinaryString(blob Blob) string {
	return p.Call("readAsBinaryString", JSValue(blob)).String()
}

func (p *fileReaderSyncImpl) ReadAsText(blob Blob, label ...string) string {
	switch len(label) {
	case 0:
		return p.Call("readAsText", JSValue(blob)).String()
	default:
		return p.Call("readAsText", JSValue(blob), label[0]).String()
	}
}

func (p *fileReaderSyncImpl) ReadAsDataURL(blob Blob) string {
	return p.Call("readAsDataURL", JSValue(blob)).String()
}

// -------------8<---------------------------------------

type progressEventImpl struct {
	*eventImpl
}

func wrapProgressEvent(v Value) ProgressEvent {
	if v.Valid() {
		return &progressEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
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
	if jsBlob := jsGlobal.Get("Blob"); jsBlob.Valid() {
		switch len(args) {
		case 1:
			if ar, ok := args[0].([]byte); ok {
				ta := js.TypedArrayOf(ar)
				defer ta.Release()
				return wrapBlob(jsBlob.New(ta))
			}
		case 2:
			if ar, ok := args[0].([]byte); ok {
				if options, ok := args[1].(BlobPropertyBag); ok {
					ta := js.TypedArrayOf(ar)
					defer ta.Release()
					return wrapBlob(jsBlob.New(ta, options.toJSObject()))
				}
			}
		}

		return wrapBlob(jsBlob.New())
	}
	return nil
}

func NewFile(fileBits []byte, fileName string, options ...FilePropertyBag) File {
	if jsFile := jsGlobal.Get("File"); jsFile.Valid() {
		ta := js.TypedArrayOf(fileBits)
		defer ta.Release()

		switch len(options) {
		case 0:
			return wrapFile(jsFile.New(ta, fileName))
		default:
			return wrapFile(jsFile.New(ta, fileName, options[0].toJSObject()))
		}
	}
	return nil
}

func NewFileReader() FileReader {
	if jsFileReader := jsGlobal.Get("FileReader"); jsFileReader.Valid() {
		return wrapFileReader(jsFileReader.New())
	}
	return nil
}

func NewFileReaderSync() FileReaderSync {
	if jsFileReaderSync := jsGlobal.Get("FileReaderSync"); jsFileReaderSync.Valid() {
		return wrapFileReaderSync(jsFileReaderSync.New())
	}
	return nil
}

func NewProgressEvent(typ string, pei ...ProgressEventInit) ProgressEvent {
	if jsProgressEvent := jsGlobal.Get("ProgressEvent"); jsProgressEvent.Valid() {
		switch len(pei) {
		case 0:
			return wrapProgressEvent(jsProgressEvent.New(typ))
		default:
			return wrapProgressEvent(jsProgressEvent.New(typ, pei[0].toJSObject()))
		}
	}
	return nil
}
