// +build js,wasm

package wasm

// -------------8<---------------------------------------

type fileImpl struct {
	Value
}

func wrapFile(v Value) File {
	if v.valid() {
		return &fileImpl{
			Value: v,
		}
	}
	return nil
}

func (p *fileImpl) Name() string {
	return p.get("name").toString()
}

func (p *fileImpl) LastModified() int {
	return p.get("lastModified").toInt()
}

// -------------8<---------------------------------------

type blobImpl struct {
	Value
}

func wrapBlob(v Value) Blob {
	if v.valid() {
		return &blobImpl{
			Value: v,
		}
	}
	return nil
}

func (p *blobImpl) Size() uint {
	return p.get("size").toUint()
}

func (p *blobImpl) Type() string {
	return p.get("type").toString()
}

func (p *blobImpl) Slice(args ...interface{}) Blob {
	switch len(args) {
	case 1:
		if start, ok := args[0].(int); ok {
			return wrapBlob(p.call("slice", start))
		}
	case 2:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				return wrapBlob(p.call("slice", start, end))
			}
		}
	case 3:
		if start, ok := args[0].(int); ok {
			if end, ok := args[1].(int); ok {
				if contentType, ok := args[2].(string); ok {
					return wrapBlob(p.call("slice", start, end, contentType))
				}
			}
		}
	}
	// wrong parameter count or parameter not given
	return wrapBlob(p.call("slice"))
}

// -------------8<---------------------------------------

type fileReaderImpl struct {
	*eventTargetImpl
}

func wrapFileReader(v Value) FileReader {
	if v.valid() {
		return &fileReaderImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *fileReaderImpl) ReadAsArrayBuffer(blob Blob) {
	p.call("readAsArrayBuffer", JSValueOf(blob))
}

func (p *fileReaderImpl) ReadAsBinaryString(blob Blob) {
	p.call("readAsBinaryString", JSValueOf(blob))
}

func (p *fileReaderImpl) ReadAsText(blob Blob, label ...string) {
	switch len(label) {
	case 0:
		p.call("readAsText", JSValueOf(blob))
	default: // 1 or more
		p.call("readAsText", JSValueOf(blob), label[0])
	}
}

func (p *fileReaderImpl) ReadAsDataURL(blob Blob) {
	p.call("readAsDataURL", JSValueOf(blob))
}

func (p *fileReaderImpl) Abort() {
	p.call("abort")
}

func (p *fileReaderImpl) ReadyState() FileReaderState {
	return FileReaderState(p.get("readyState").toInt())
}

// TODO
func (p *fileReaderImpl) Result() []byte {
	v := p.get("result")
	switch v.jsValue.Type() {
	case TypeString:
		return []byte(v.toString())
	case TypeObject: // ArrayBuffer
		return wrapArrayBuffer(v).ToByteSlice()
	default:
		return nil
	}
}

func (p *fileReaderImpl) Error() DOMException {
	return wrapDOMException(p.get("error"))
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
	if v.valid() {
		return &fileReaderSyncImpl{
			Value: v,
		}
	}
	return nil
}

func (p *fileReaderSyncImpl) ReadAsArrayBuffer(blob Blob) ArrayBuffer {
	return wrapArrayBuffer(p.call("readAsArrayBuffer", JSValueOf(blob)))
}

func (p *fileReaderSyncImpl) ReadAsBinaryString(blob Blob) string {
	return p.call("readAsBinaryString", JSValueOf(blob)).toString()
}

func (p *fileReaderSyncImpl) ReadAsText(blob Blob, label ...string) string {
	switch len(label) {
	case 0:
		return p.call("readAsText", JSValueOf(blob)).toString()
	default:
		return p.call("readAsText", JSValueOf(blob), label[0]).toString()
	}
}

func (p *fileReaderSyncImpl) ReadAsDataURL(blob Blob) string {
	return p.call("readAsDataURL", JSValueOf(blob)).toString()
}

// -------------8<---------------------------------------

type progressEventImpl struct {
	*eventImpl
}

func wrapProgressEvent(v Value) ProgressEvent {
	if v.valid() {
		return &progressEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *progressEventImpl) LengthComputable() bool {
	return p.get("lengthComputable").toBool()
}

func (p *progressEventImpl) Loaded() uint {
	return p.get("loaded").toUint()
}

func (p *progressEventImpl) Total() uint {
	return p.get("total").toUint()
}

// -------------8<---------------------------------------

func NewBlob(args ...interface{}) Blob {
	if jsBlob := jsGlobal.get("Blob"); jsBlob.valid() {
		switch len(args) {
		case 1:
			if ar, ok := args[0].([]byte); ok {
				ta := jsTypedArrayOf(ar)
				defer ta.Release()
				return wrapBlob(jsBlob.jsNew(ta))
			}
		case 2:
			if ar, ok := args[0].([]byte); ok {
				if options, ok := args[1].(BlobPropertyBag); ok {
					ta := jsTypedArrayOf(ar)
					defer ta.Release()
					return wrapBlob(jsBlob.jsNew(ta, options.JSValue()))
				}
			}
		}

		return wrapBlob(jsBlob.jsNew())
	}
	return nil
}

func NewFile(fileBits []byte, fileName string, options ...FilePropertyBag) File {
	if jsFile := jsGlobal.get("File"); jsFile.valid() {
		ta := jsTypedArrayOf(fileBits)
		defer ta.Release()

		switch len(options) {
		case 0:
			return wrapFile(jsFile.jsNew(ta, fileName))
		default:
			return wrapFile(jsFile.jsNew(ta, fileName, options[0].JSValue()))
		}
	}
	return nil
}

func NewFileReader() FileReader {
	if jsFileReader := jsGlobal.get("FileReader"); jsFileReader.valid() {
		return wrapFileReader(jsFileReader.jsNew())
	}
	return nil
}

func NewFileReaderSync() FileReaderSync {
	if jsFileReaderSync := jsGlobal.get("FileReaderSync"); jsFileReaderSync.valid() {
		return wrapFileReaderSync(jsFileReaderSync.jsNew())
	}
	return nil
}

func NewProgressEvent(typ string, pei ...ProgressEventInit) ProgressEvent {
	if jsProgressEvent := jsGlobal.get("ProgressEvent"); jsProgressEvent.valid() {
		switch len(pei) {
		case 0:
			return wrapProgressEvent(jsProgressEvent.jsNew(typ))
		default:
			return wrapProgressEvent(jsProgressEvent.jsNew(typ, pei[0].JSValue()))
		}
	}
	return nil
}
