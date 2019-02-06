// +build js,wasm

package wasm

type (
	// https://www.w3.org/TR/FileAPI/#dfn-Blob
	Blob interface {
		Size() uint
		Type() string
		Slice(...interface{}) Blob

		JSValue() jsValue
	}

	// https://www.w3.org/TR/FileAPI/#dfn-file
	File interface {
		Blob

		Name() string
		LastModified() int
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
		ReadyState() FileReaderReadyState

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
		Loaded() uint
		Total() uint
	}
)

type EndingType string

const (
	EndingTypeTransparent EndingType = "transparent"
	EndingTypeNative      EndingType = "native"
)

type FileReaderReadyState uint16

const (
	FileReaderReadyStateEmpty   FileReaderReadyState = 0
	FileReaderReadyStateLoading FileReaderReadyState = 1
	FileReaderReadyStateDone    FileReaderReadyState = 2
)

// -------------8<---------------------------------------

// https://w3c.github.io/FileAPI/#dfn-BlobPropertyBag
type BlobPropertyBag struct {
	Type string
	// working draft
	// Endings EndingType `json:"endings"` // default transparent
}

func (p BlobPropertyBag) JSValue() jsValue {
	o := jsObject.New()
	o.Set("type", p.Type)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/FileAPI/#dfn-FilePropertyBag
type FilePropertyBag struct {
	BlobPropertyBag

	LastModified int
}

func (p FilePropertyBag) JSValue() jsValue {
	o := p.BlobPropertyBag.JSValue()
	o.Set("lastModified", p.LastModified)
	return o
}

// -------------8<---------------------------------------

// https://xhr.spec.whatwg.org/#progresseventinit
type ProgressEventInit struct {
	EventInit

	LengthComputable bool
	Loaded           uint
	Total            uint
}

func (p ProgressEventInit) JSValue() jsValue {
	o := p.EventInit.JSValue()
	o.Set("lengthComputable", p.LengthComputable)
	o.Set("loaded", p.Loaded)
	o.Set("total", p.Total)
	return o
}
