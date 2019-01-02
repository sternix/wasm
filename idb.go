// +build js,wasm

package wasm

import (
	"syscall/js"
)

type (
	// https://www.w3.org/TR/IndexedDB/#idbrequest
	IDBRequest interface {
		EventTarget

		Result() interface{}
		Error() DOMException
		Source() IDBRequestSource
		Transaction() IDBTransaction
		ReadyState() IDBRequestReadyState
		OnSuccess(func(Event)) EventHandler
		OnError(func(Event)) EventHandler
	}

	// (IDBObjectStore or IDBIndex or IDBCursor)?
	IDBRequestSource interface{}

	// https://www.w3.org/TR/IndexedDB/#idbopendbrequest
	IDBOpenDBRequest interface {
		IDBRequest

		OnBlocked(func(Event)) EventHandler
		OnUpgradeNeeded(func(IDBVersionChangeEvent)) EventHandler
	}

	// https://www.w3.org/TR/IndexedDB/#idbversionchangeevent
	IDBVersionChangeEvent interface {
		Event

		OldVersion() int
		NewVersion() int
	}

	// https://www.w3.org/TR/IndexedDB/#idbfactory
	IDBFactory interface {
		Open(string, ...int) IDBOpenDBRequest
		DeleteDatabase(string) IDBOpenDBRequest
		Cmp(interface{}, interface{}) int
	}

	// https://www.w3.org/TR/IndexedDB/#idbdatabase
	IDBDatabase interface {
		EventTarget

		Name() string
		Version() int
		ObjectStoreNames() []string
		Transaction([]string, ...IDBTransactionMode) IDBTransaction
		Close()
		CreateObjectStore(string, ...IDBObjectStoreParameters) IDBObjectStore
		DeleteObjectStore(string)
		OnAbort(func(Event)) EventHandler
		OnClose(func(Event)) EventHandler
		OnError(func(Event)) EventHandler
		OnVersionchange(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/IndexedDB/#idbobjectstore
	IDBObjectStore interface {
		Name() string
		SetName(string)
		KeyPath() string // any
		IndexNames() []string
		Transaction() IDBTransaction
		AutoIncrement() bool
		Put(interface{}, ...interface{}) IDBRequest
		Add(interface{}, ...interface{}) IDBRequest
		Delete(interface{}) IDBRequest
		Clear() IDBRequest
		Get(interface{}) IDBRequest
		Key(interface{}) IDBRequest
		All(...interface{}) IDBRequest
		AllKeys(...interface{}) IDBRequest
		Count(...interface{}) IDBRequest
		OpenCursor(...interface{}) IDBRequest
		OpenKeyCursor(...interface{}) IDBRequest
		Index(string) IDBIndex
		CreateIndex(string, string, ...IDBIndexParameters) IDBIndex
		DeleteIndex(string)
	}

	// https://www.w3.org/TR/IndexedDB/#idbindex
	IDBIndex interface {
		Name() string
		SetName(string)
		ObjectStore() IDBObjectStore
		KeyPath() string
		MultiEntry() bool
		Unique() bool
		Get(interface{}) IDBRequest
		Key(interface{}) IDBRequest
		All(...interface{}) IDBRequest
		AllKeys(...interface{}) IDBRequest
		Count(...interface{}) IDBRequest
		OpenCursor(...interface{}) IDBRequest
		OpenKeyCursor(...interface{}) IDBRequest
	}

	// https://www.w3.org/TR/IndexedDB/#idbkeyrange
	IDBKeyRange interface {
		Lower() interface{}
		Upper() interface{}
		LowerOpen() bool
		UpperOpen() bool
		Only(interface{}) IDBKeyRange                        //static
		LowerBound(interface{}, ...bool) IDBKeyRange         //static
		UpperBound(interface{}, ...bool) IDBKeyRange         //static
		Bound(interface{}, interface{}, ...bool) IDBKeyRange //static
		Includes(interface{}) bool
	}

	// https://www.w3.org/TR/IndexedDB/#idbcursor
	IDBCursor interface {
		Source() IDBCursorSource
		Direction() IDBCursorDirection
		Key() interface{}
		PrimaryKey() interface{}
		Advance(int)
		Continue(...interface{})
		ContinuePrimaryKey(interface{}, interface{})
		Update(interface{}) IDBRequest
		Delete() IDBRequest
	}

	// (IDBObjectStore or IDBIndex)
	IDBCursorSource interface{}

	// https://www.w3.org/TR/IndexedDB/#idbcursorwithvalue
	IDBCursorWithValue interface {
		IDBCursor

		Value() interface{}
	}

	// https://www.w3.org/TR/IndexedDB/#idbtransaction
	IDBTransaction interface {
		EventTarget

		ObjectStoreNames() []string // DOMStringList
		Mode() IDBTransactionMode
		DB() IDBDatabase
		Error() DOMException
		ObjectStore(string) IDBObjectStore
		Abort()
		OnAbort(func(Event)) EventHandler
		OnComplete(func(Event)) EventHandler
		OnError(func(Event)) EventHandler
	}
)

// https://www.w3.org/TR/IndexedDB/#enumdef-idbrequestreadystate
type IDBRequestReadyState string

const (
	IDBRequestReadyStatePending IDBRequestReadyState = "pending"
	IDBRequestReadyStateDone    IDBRequestReadyState = "done"
)

// https://www.w3.org/TR/IndexedDB/#enumdef-idbcursordirection
type IDBCursorDirection string

const (
	IDBCursorDirectionNext       IDBCursorDirection = "next"
	IDBCursorDirectionNextUnique IDBCursorDirection = "nextunique"
	IDBCursorDirectionPrev       IDBCursorDirection = "prev"
	IDBCursorDirectionPrevUnique IDBCursorDirection = "prevunique"
)

// https://www.w3.org/TR/IndexedDB/#enumdef-idbtransactionmode
type IDBTransactionMode string

const (
	IDBTransactionModeReadOnly      IDBTransactionMode = "readonly"
	IDBTransactionModeReadWrite     IDBTransactionMode = "readwrite"
	IDBTransactionModeVersionChange IDBTransactionMode = "versionchange"
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/IndexedDB/#dictdef-idbversionchangeeventinit
type IDBVersionChangeEventInit struct {
	OldVersion int
	NewVersion int // default null
}

func (p IDBVersionChangeEventInit) toJSObject() js.Value {
	o := jsObject.New()
	o.Set("oldVersion", p.OldVersion)
	o.Set("newVersion", p.NewVersion)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/IndexedDB/#dictdef-idbobjectstoreparameters
type IDBObjectStoreParameters struct {
	KeyPath       []string
	AutoIncrement bool
}

func (p IDBObjectStoreParameters) toJSObject() js.Value {
	o := jsObject.New()
	o.Set("keyPath", sliceToJsArray(p.KeyPath))
	o.Set("autoIncrement", p.AutoIncrement)
	return o
}

// -------------8<---------------------------------------

// https://www.w3.org/TR/IndexedDB/#dictdef-idbindexparameters
type IDBIndexParameters struct {
	Unique     bool
	MultiEntry bool
}

func (p IDBIndexParameters) toJSObject() js.Value {
	o := jsObject.New()
	o.Set("unique", p.Unique)
	o.Set("multiEntry", p.MultiEntry)
	return o
}
