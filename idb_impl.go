// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type idbRequestImpl struct {
	*eventTargetImpl
}

func wrapIDBRequest(v js.Value) IDBRequest {
	if p := newIDBRequestImpl(v); p != nil {
		return p
	}
	return nil
}

func newIDBRequestImpl(v js.Value) *idbRequestImpl {
	if isNil(v) {
		return nil
	}

	return &idbRequestImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *idbRequestImpl) Result() interface{} {
	return Wrap(p.Get("result"))
}

func (p *idbRequestImpl) Error() DOMException {
	return wrapDOMException(p.Get("error"))
}

func (p *idbRequestImpl) Source() IDBRequestSource {
	return wrapIDBRequestSource(p.Get("source"))
}

func (p *idbRequestImpl) Transaction() IDBTransaction {
	return wrapIDBTransaction(p.Get("transaction"))
}

func (p *idbRequestImpl) ReadyState() IDBRequestReadyState {
	return IDBRequestReadyState(p.Get("readyState").String())
}

func (p *idbRequestImpl) OnSuccess(fn func(Event)) EventHandler {
	return p.On("success", fn)
}

func (p *idbRequestImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

// -------------8<---------------------------------------
// (IDBObjectStore or IDBIndex or IDBCursor)?
// TODO

type idbRequestSourceImpl struct {
	js.Value
}

func wrapIDBRequestSource(v js.Value) IDBRequestSource {
	if isNil(v) {
		return nil
	}

	return &idbRequestSourceImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type idbOpenDBRequestImpl struct {
	*idbRequestImpl
}

func wrapIDBOpenDBRequest(v js.Value) IDBOpenDBRequest {
	if isNil(v) {
		return nil
	}

	return &idbOpenDBRequestImpl{
		idbRequestImpl: newIDBRequestImpl(v),
	}
}

func (p *idbOpenDBRequestImpl) OnBlocked(fn func(Event)) EventHandler {
	return p.On("blocked", fn)
}

func (p *idbOpenDBRequestImpl) OnUpgradeNeeded(fn func(IDBVersionChangeEvent)) EventHandler {
	return p.On("upgradeneeded", func(e Event) {
		if ve, ok := e.(IDBVersionChangeEvent); ok {
			fn(ve)
		}
	})
}

// -------------8<---------------------------------------

type idbVersionChangeEventImpl struct {
	*eventImpl
}

func wrapIDBVersionChangeEvent(v js.Value) IDBVersionChangeEvent {
	if isNil(v) {
		return nil
	}

	return &idbVersionChangeEventImpl{
		eventImpl: newEventImpl(v),
	}
}

func (p *idbVersionChangeEventImpl) OldVersion() int {
	return p.Get("oldVersion").Int()
}

func (p *idbVersionChangeEventImpl) NewVersion() int {
	return p.Get("newVersion").Int()
}

// -------------8<---------------------------------------

type idbFactoryImpl struct {
	js.Value
}

func wrapIDBFactory(v js.Value) IDBFactory {
	if isNil(v) {
		return nil
	}

	return &idbFactoryImpl{
		Value: v,
	}
}

func (p *idbFactoryImpl) Open(name string, version ...int) IDBOpenDBRequest {
	if len(version) > 0 {
		return wrapIDBOpenDBRequest(p.Call("open", name, version[0]))
	}

	return wrapIDBOpenDBRequest(p.Call("open", name))
}

func (p *idbFactoryImpl) DeleteDatabase(name string) IDBOpenDBRequest {
	return wrapIDBOpenDBRequest(p.Call("deleteDatabase", name))
}

func (p *idbFactoryImpl) Cmp(first interface{}, second interface{}) int {
	return p.Call("cmp", first, second).Int()
}

// -------------8<---------------------------------------

type idbDatabaseImpl struct {
	*eventTargetImpl
}

func wrapIDBDatabase(v js.Value) IDBDatabase {
	if isNil(v) {
		return nil
	}

	return &idbDatabaseImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *idbDatabaseImpl) Name() string {
	return p.Get("name").String()
}

func (p *idbDatabaseImpl) Version() int {
	return p.Get("version").Int()
}

func (p *idbDatabaseImpl) ObjectStoreNames() []string {
	return domStringListToSlice(p.Get("objectStoreNames"))
}

func (p *idbDatabaseImpl) Transaction(storeNames []string, mode ...IDBTransactionMode) IDBTransaction {
	arr := sliceToJsArray(storeNames)
	switch len(mode) {
	case 0:
		return wrapIDBTransaction(p.Call("transaction", arr))
	default:
		return wrapIDBTransaction(p.Call("transaction", arr, mode[0]))
	}
}

func (p *idbDatabaseImpl) Close() {
	p.Call("close")
}

func (p *idbDatabaseImpl) CreateObjectStore(name string, options ...IDBObjectStoreParameters) IDBObjectStore {
	switch len(options) {
	case 0:
		return wrapIDBObjectStore(p.Call("createObjectStore", name))
	default:
		return wrapIDBObjectStore(p.Call("createObjectStore", name, options[0].toDict()))
	}
}

func (p *idbDatabaseImpl) DeleteObjectStore(name string) {
	p.Call("deleteObjectStore", name)
}

func (p *idbDatabaseImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

func (p *idbDatabaseImpl) OnClose(fn func(Event)) EventHandler {
	return p.On("close", fn)
}

func (p *idbDatabaseImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

func (p *idbDatabaseImpl) OnVersionchange(fn func(Event)) EventHandler {
	return p.On("versionchange", fn)
}

// -------------8<---------------------------------------

type idbObjectStoreImpl struct {
	js.Value
}

func wrapIDBObjectStore(v js.Value) IDBObjectStore {
	if isNil(v) {
		return nil
	}

	return &idbObjectStoreImpl{
		Value: v,
	}
}

func (p *idbObjectStoreImpl) Name() string {
	return p.Value.Get("name").String()
}

func (p *idbObjectStoreImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *idbObjectStoreImpl) KeyPath() string {
	return p.Value.Get("keyPath").String()
}

func (p *idbObjectStoreImpl) IndexNames() []string {
	return domStringListToSlice(p.Value.Get("indexNames"))
}

func (p *idbObjectStoreImpl) Transaction() IDBTransaction {
	return wrapIDBTransaction(p.Value.Get("transaction"))
}

func (p *idbObjectStoreImpl) AutoIncrement() bool {
	return p.Value.Get("autoIncrement").Bool()
}

func (p *idbObjectStoreImpl) Put(value interface{}, key ...interface{}) IDBRequest {
	switch len(key) {
	case 0:
		return wrapIDBRequest(p.Call("put", value))
	default:
		return wrapIDBRequest(p.Call("put", value, key[0]))
	}
}

func (p *idbObjectStoreImpl) Add(value interface{}, key ...interface{}) IDBRequest {
	switch len(key) {
	case 0:
		return wrapIDBRequest(p.Call("add", value))
	default:
		return wrapIDBRequest(p.Call("add", value, key[0]))
	}
}

func (p *idbObjectStoreImpl) Delete(query interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("delete", query))
}

func (p *idbObjectStoreImpl) Clear() IDBRequest {
	return wrapIDBRequest(p.Call("clear"))
}

func (p *idbObjectStoreImpl) Get(query interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("get", query))
}

func (p *idbObjectStoreImpl) Key(query interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("getKey", query))
}

func (p *idbObjectStoreImpl) All(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("getAll", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.Call("getAll", args[0], count))
		}
	}

	return wrapIDBRequest(p.Call("getAll"))
}

func (p *idbObjectStoreImpl) AllKeys(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("getAllKeys", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.Call("getAllKeys", args[0], count))
		}
	}

	return wrapIDBRequest(p.Call("getAllKeys"))
}

func (p *idbObjectStoreImpl) Count(query ...interface{}) IDBRequest {
	switch len(query) {
	case 0:
		return wrapIDBRequest(p.Call("count"))
	default:
		return wrapIDBRequest(p.Call("count", query[0]))
	}
}

func (p *idbObjectStoreImpl) OpenCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("openCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.Call("openCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.Call("openCursor"))
}

func (p *idbObjectStoreImpl) OpenKeyCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("openKeyCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.Call("openKeyCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.Call("openKeyCursor"))
}

func (p *idbObjectStoreImpl) Index(name string) IDBIndex {
	return wrapIDBIndex(p.Call("index", name))
}

func (p *idbObjectStoreImpl) CreateIndex(name string, keyPath string, options ...IDBIndexParameters) IDBIndex {
	switch len(options) {
	case 0:
		return wrapIDBIndex(p.Call("createIndex", name, keyPath))
	default:
		return wrapIDBIndex(p.Call("createIndex", name, keyPath, options[0].toDict()))
	}
}

func (p *idbObjectStoreImpl) DeleteIndex(name string) {
	p.Call("deleteIndex", name)
}

// -------------8<---------------------------------------

type idbIndexImpl struct {
	js.Value
}

func wrapIDBIndex(v js.Value) IDBIndex {
	if isNil(v) {
		return nil
	}

	return &idbIndexImpl{
		Value: v,
	}
}

func (p *idbIndexImpl) Name() string {
	return p.Value.Get("name").String()
}

func (p *idbIndexImpl) SetName(name string) {
	p.Set("name", name)
}

func (p *idbIndexImpl) ObjectStore() IDBObjectStore {
	return wrapIDBObjectStore(p.Value.Get("objectStore"))
}

func (p *idbIndexImpl) KeyPath() string {
	return p.Value.Get("keyPath").String()
}

func (p *idbIndexImpl) MultiEntry() bool {
	return p.Value.Get("multiEntry").Bool()
}

func (p *idbIndexImpl) Unique() bool {
	return p.Value.Get("unique").Bool()
}

func (p *idbIndexImpl) Get(query interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("get", query))
}

func (p *idbIndexImpl) Key(query interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("getKey", query))
}

func (p *idbIndexImpl) All(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("getAll", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.Call("getAll", args[0], count))
		}
	}

	return wrapIDBRequest(p.Call("getAll"))
}

func (p *idbIndexImpl) AllKeys(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("getAllKeys", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.Call("getAllKeys", args[0], count))
		}
	}

	return wrapIDBRequest(p.Call("getAllKeys"))
}

func (p *idbIndexImpl) Count(query ...interface{}) IDBRequest {
	switch len(query) {
	case 0:
		return wrapIDBRequest(p.Call("count"))
	default:
		return wrapIDBRequest(p.Call("count", query[0]))
	}
}

func (p *idbIndexImpl) OpenCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("openCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.Call("openCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.Call("openCursor"))
}

func (p *idbIndexImpl) OpenKeyCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.Call("openKeyCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.Call("openKeyCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.Call("openKeyCursor"))
}

// -------------8<---------------------------------------

type idbKeyRangeImpl struct {
	js.Value
}

func wrapIDBKeyRange(v js.Value) IDBKeyRange {
	if isNil(v) {
		return nil
	}

	return &idbKeyRangeImpl{
		Value: v,
	}
}

func (p *idbKeyRangeImpl) Lower() interface{} {
	return Wrap(p.Get("lower"))
}

func (p *idbKeyRangeImpl) Upper() interface{} {
	return Wrap(p.Get("upper"))
}

func (p *idbKeyRangeImpl) LowerOpen() bool {
	return p.Get("lowerOpen").Bool()
}

func (p *idbKeyRangeImpl) UpperOpen() bool {
	return p.Get("upperOpen").Bool()
}

//static
func (p *idbKeyRangeImpl) Only(value interface{}) IDBKeyRange {
	return wrapIDBKeyRange(p.Call("only", value))
}

//static
func (p *idbKeyRangeImpl) LowerBound(lower interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 0:
		return wrapIDBKeyRange(p.Call("lowerBound", lower))
	default:
		return wrapIDBKeyRange(p.Call("lowerBound", lower, open[0]))
	}
}

//static
func (p *idbKeyRangeImpl) UpperBound(upper interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 0:
		return wrapIDBKeyRange(p.Call("upperBound", upper))
	default:
		return wrapIDBKeyRange(p.Call("upperBound", upper, open[0]))
	}
}

//static
func (p *idbKeyRangeImpl) Bound(lower interface{}, upper interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 1:
		return wrapIDBKeyRange(p.Call("bound", lower, upper, open[0]))
	case 2:
		return wrapIDBKeyRange(p.Call("bound", lower, upper, open[0], open[1]))
	default:
		return wrapIDBKeyRange(p.Call("bound", lower, upper))
	}
}

func (p *idbKeyRangeImpl) Includes(key interface{}) bool {
	return p.Call("_includes", key).Bool()
}

// -------------8<---------------------------------------

type idbCursorImpl struct {
	js.Value
}

func wrapIDBCursor(v js.Value) IDBCursor {
	if p := newIDBCursorImpl(v); p != nil {
		return p
	}
	return nil
}

func newIDBCursorImpl(v js.Value) *idbCursorImpl {
	if isNil(v) {
		return nil
	}

	return &idbCursorImpl{
		Value: v,
	}
}

func (p *idbCursorImpl) Source() IDBCursorSource {
	return wrapIDBCursorSource(p.Get("source"))
}

func (p *idbCursorImpl) Direction() IDBCursorDirection {
	return IDBCursorDirection(p.Get("direction").String())
}

func (p *idbCursorImpl) Key() interface{} {
	return Wrap(p.Get("key"))
}

func (p *idbCursorImpl) PrimaryKey() interface{} {
	return Wrap(p.Get("primaryKey"))
}

func (p *idbCursorImpl) Advance(count int) {
	p.Call("advance", count)
}

func (p *idbCursorImpl) Continue(key ...interface{}) {
	switch len(key) {
	case 0:
		p.Call("continue")
	default:
		p.Call("continue", key[0])
	}
}

func (p *idbCursorImpl) ContinuePrimaryKey(key interface{}, primaryKey interface{}) {
	p.Call("continuePrimaryKey", key, primaryKey)
}

func (p *idbCursorImpl) Update(value interface{}) IDBRequest {
	return wrapIDBRequest(p.Call("update", value))
}

func (p *idbCursorImpl) Delete() IDBRequest {
	return wrapIDBRequest(p.Call("delete"))
}

// -------------8<---------------------------------------
// (IDBObjectStore or IDBIndex)
type idbCursorSourceImpl struct {
	js.Value
}

func wrapIDBCursorSource(v js.Value) IDBCursorSource {
	if isNil(v) {
		return nil
	}

	return &idbCursorSourceImpl{
		Value: v,
	}
}

// -------------8<---------------------------------------

type idbCursorWithValueImpl struct {
	*idbCursorImpl
}

func wrapIDBCursorWithValue(v js.Value) IDBCursorWithValue {
	if isNil(v) {
		return nil
	}

	return &idbCursorWithValueImpl{
		idbCursorImpl: newIDBCursorImpl(v),
	}
}

func (p *idbCursorWithValueImpl) Value() interface{} {
	return Wrap(p.Get("value"))
}

// -------------8<---------------------------------------

type idbTransactionImpl struct {
	*eventTargetImpl
}

func wrapIDBTransaction(v js.Value) IDBTransaction {
	if isNil(v) {
		return nil
	}

	return &idbTransactionImpl{
		eventTargetImpl: newEventTargetImpl(v),
	}
}

func (p *idbTransactionImpl) ObjectStoreNames() []string {
	return domStringListToSlice(p.Get("objectStoreNames"))
}

func (p *idbTransactionImpl) Mode() IDBTransactionMode {
	return IDBTransactionMode(p.Get("mode").String())
}

func (p *idbTransactionImpl) DB() IDBDatabase {
	return wrapIDBDatabase(p.Get("db"))
}

func (p *idbTransactionImpl) Error() DOMException {
	return wrapDOMException(p.Get("error"))
}

func (p *idbTransactionImpl) ObjectStore(name string) IDBObjectStore {
	return wrapIDBObjectStore(p.Call("objectStore", name))
}

func (p *idbTransactionImpl) Abort() {
	p.Call("abort")
}

func (p *idbTransactionImpl) OnAbort(fn func(Event)) EventHandler {
	return p.On("abort", fn)
}

func (p *idbTransactionImpl) OnComplete(fn func(Event)) EventHandler {
	return p.On("complete", fn)
}

func (p *idbTransactionImpl) OnError(fn func(Event)) EventHandler {
	return p.On("error", fn)
}

// -------------8<---------------------------------------

func NewIDBVersionChangeEvent(typ string, vce ...IDBVersionChangeEventInit) IDBVersionChangeEvent {
	jsIDBVersionChangeEvent := js.Global().Get("IDBVersionChangeEvent")
	if isNil(jsIDBVersionChangeEvent) {
		return nil
	}

	switch len(vce) {
	case 0:
		return wrapIDBVersionChangeEvent(jsIDBVersionChangeEvent.New(typ))
	default:
		return wrapIDBVersionChangeEvent(jsIDBVersionChangeEvent.New(typ, vce[0].toDict()))
	}
}
