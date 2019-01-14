// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/FileAPI/#dfn-file
	File interface {
		Name() string
		LastModified() int
	}

	// https://www.w3.org/TR/FileAPI/#dfn-Blob
	Blob interface {
		Size() int
		Type() string
		Slice(...interface{}) Blob
	}

	// https://www.w3.org/TR/FileAPI/#typedefdef-blobpart
	// typedef (BufferSource or Blob or USVString) BlobPart;
	BlobPart interface{}

	// https://www.w3.org/TR/FileAPI/#dfn-filereader
	FileReader interface {
		EventTarget

		ReadAsArrayBuffer(Blob)
		ReadAsBinaryString(Blob)
		ReadAsText(Blob, ...string)
		ReadAsDataURL(Blob)
		Abort()
		ReadyState() FileReaderState

		// File or Blob data
		Result() []byte // (DOMString or ArrayBuffer)? , if result is string convert to []byte
		Error() DOMException

		OnLoadStart(func(Event)) EventHandler
		OnProgress(func(Event)) EventHandler
		OnLoad(func(Event)) EventHandler
		OnAbort(func(Event)) EventHandler
		OnError(func(Event)) EventHandler
		OnLoadEnd(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/FileAPI/#dfn-FileReaderSync
	FileReaderSync interface {
		ReadAsArrayBuffer(Blob) ArrayBuffer
		ReadAsBinaryString(Blob) string
		ReadAsText(Blob, ...string) string
		ReadAsDataURL(Blob) string
	}

	// https://xhr.spec.whatwg.org/#interface-progressevent
	ProgressEvent interface {
		Event

		LengthComputable() bool
		Loaded() int
		Total() int
	}
)

type EndingType string

const (
	EndingTypeTransparent EndingType = "transparent"
	EndingTypeNative      EndingType = "native"
)

type FileReaderState uint

const (
	FileReaderStateEmpty   FileReaderState = 0
	FileReaderStateLoading FileReaderState = 1
	FileReaderStateDone    FileReaderState = 2
)

// -------------8<---------------------------------------

// https://w3c.github.io/FileAPI/#dfn-BlobPropertyBag
type BlobPropertyBag struct {
	Type string
	// working draft
	// Endings EndingType `json:"endings"` // default transparent
}

func (p BlobPropertyBag) toJSObject() Value {
	o := jsObject.jsNew()
	o.set("type", p.Type)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/FileAPI/#dfn-FilePropertyBag
type FilePropertyBag struct {
	BlobPropertyBag

	LastModified int
}

func (p FilePropertyBag) toJSObject() Value {
	o := p.BlobPropertyBag.toJSObject()
	o.set("lastModified", p.LastModified)
	return o
}

// -------------8<---------------------------------------

// https://xhr.spec.whatwg.org/#progresseventinit
type ProgressEventInit struct {
	EventInit

	LengthComputable bool
	Loaded           int
	Total            int
}

func (p ProgressEventInit) toJSObject() Value {
	o := p.EventInit.toJSObject()
	o.set("lengthComputable", p.LengthComputable)
	o.set("loaded", p.Loaded)
	o.set("total", p.Total)
	return o
}
