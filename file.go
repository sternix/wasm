// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://www.w3.org/TR/FileAPI/#dfn-file
	File interface {
		js.Wrapper

		Name() string
		LastModified() int
	}

	// https://www.w3.org/TR/FileAPI/#dfn-Blob
	Blob interface {
		js.Wrapper

		Size() int
		Type() string
		Slice(...interface{}) Blob
	}

	// https://www.w3.org/TR/FileAPI/#typedefdef-blobpart
	BlobPart interface {
		js.Wrapper
	} // typedef (BufferSource or Blob or USVString) BlobPart;

	// https://w3c.github.io/FileAPI/#dfn-BlobPropertyBag
	BlobPropertyBag struct {
		Type string `json:"type"`
		// working draft
		// Endings EndingType `json:"endings"` // default transparent
	}

	// https://www.w3.org/TR/FileAPI/#dfn-FilePropertyBag
	FilePropertyBag struct {
		BlobPropertyBag

		LastModified int `json:"lastModified"`
	}

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
		js.Wrapper

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

	// https://xhr.spec.whatwg.org/#progresseventinit
	ProgressEventInit struct {
		EventInit

		LengthComputable bool `json:"lengthComputable"`
		Loaded           int  `json:"loaded"`
		Total            int  `json:"total"`
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
