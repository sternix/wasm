// +build js,wasm

package wasm

// -------------8<---------------------------------------

type idbRequestImpl struct {
	*eventTargetImpl
}

func wrapIDBRequest(v Value) IDBRequest {
	if p := newIDBRequestImpl(v); p != nil {
		return p
	}
	return nil
}

func newIDBRequestImpl(v Value) *idbRequestImpl {
	if v.valid() {
		return &idbRequestImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *idbRequestImpl) Result() interface{} {
	return Wrap(p.get("result"))
}

func (p *idbRequestImpl) Error() DOMException {
	return wrapDOMException(p.get("error"))
}

func (p *idbRequestImpl) Source() IDBRequestSource {
	return wrapIDBRequestSource(p.get("source"))
}

func (p *idbRequestImpl) Transaction() IDBTransaction {
	return wrapIDBTransaction(p.get("transaction"))
}

func (p *idbRequestImpl) ReadyState() IDBRequestReadyState {
	return IDBRequestReadyState(p.get("readyState").toString())
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
	Value
}

func wrapIDBRequestSource(v Value) IDBRequestSource {
	if v.valid() {
		return &idbRequestSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type idbOpenDBRequestImpl struct {
	*idbRequestImpl
}

func wrapIDBOpenDBRequest(v Value) IDBOpenDBRequest {
	if v.valid() {
		return &idbOpenDBRequestImpl{
			idbRequestImpl: newIDBRequestImpl(v),
		}
	}
	return nil
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

func wrapIDBVersionChangeEvent(v Value) IDBVersionChangeEvent {
	if v.valid() {
		return &idbVersionChangeEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *idbVersionChangeEventImpl) OldVersion() uint {
	return p.get("oldVersion").toUint()
}

func (p *idbVersionChangeEventImpl) NewVersion() uint {
	return p.get("newVersion").toUint()
}

// -------------8<---------------------------------------

type idbFactoryImpl struct {
	Value
}

func wrapIDBFactory(v Value) IDBFactory {
	if v.valid() {
		return &idbFactoryImpl{
			Value: v,
		}
	}
	return nil
}

func (p *idbFactoryImpl) Open(name string, version ...int) IDBOpenDBRequest {
	if len(version) > 0 {
		return wrapIDBOpenDBRequest(p.call("open", name, version[0]))
	}

	return wrapIDBOpenDBRequest(p.call("open", name))
}

func (p *idbFactoryImpl) DeleteDatabase(name string) IDBOpenDBRequest {
	return wrapIDBOpenDBRequest(p.call("deleteDatabase", name))
}

func (p *idbFactoryImpl) Cmp(first interface{}, second interface{}) int {
	return p.call("cmp", first, second).toInt()
}

// -------------8<---------------------------------------

type idbDatabaseImpl struct {
	*eventTargetImpl
}

func wrapIDBDatabase(v Value) IDBDatabase {
	if v.valid() {
		return &idbDatabaseImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *idbDatabaseImpl) Name() string {
	return p.get("name").toString()
}

func (p *idbDatabaseImpl) Version() uint {
	return p.get("version").toUint()
}

func (p *idbDatabaseImpl) ObjectStoreNames() []string {
	return domStringListToSlice(p.get("objectStoreNames"))
}

func (p *idbDatabaseImpl) Transaction(storeNames []string, mode ...IDBTransactionMode) IDBTransaction {
	arr := sliceToJsArray(storeNames)
	switch len(mode) {
	case 0:
		return wrapIDBTransaction(p.call("transaction", arr))
	default:
		return wrapIDBTransaction(p.call("transaction", arr, mode[0]))
	}
}

func (p *idbDatabaseImpl) Close() {
	p.call("close")
}

func (p *idbDatabaseImpl) CreateObjectStore(name string, options ...IDBObjectStoreParameters) IDBObjectStore {
	switch len(options) {
	case 0:
		return wrapIDBObjectStore(p.call("createObjectStore", name))
	default:
		return wrapIDBObjectStore(p.call("createObjectStore", name, options[0].toJSObject()))
	}
}

func (p *idbDatabaseImpl) DeleteObjectStore(name string) {
	p.call("deleteObjectStore", name)
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
	Value
}

func wrapIDBObjectStore(v Value) IDBObjectStore {
	if v.valid() {
		return &idbObjectStoreImpl{
			Value: v,
		}
	}
	return nil
}

func (p *idbObjectStoreImpl) Name() string {
	return p.get("name").toString()
}

func (p *idbObjectStoreImpl) SetName(name string) {
	p.set("name", name)
}

func (p *idbObjectStoreImpl) KeyPath() string {
	return p.get("keyPath").toString()
}

func (p *idbObjectStoreImpl) IndexNames() []string {
	return domStringListToSlice(p.get("indexNames"))
}

func (p *idbObjectStoreImpl) Transaction() IDBTransaction {
	return wrapIDBTransaction(p.get("transaction"))
}

func (p *idbObjectStoreImpl) AutoIncrement() bool {
	return p.get("autoIncrement").toBool()
}

func (p *idbObjectStoreImpl) Put(value interface{}, key ...interface{}) IDBRequest {
	switch len(key) {
	case 0:
		return wrapIDBRequest(p.call("put", value))
	default:
		return wrapIDBRequest(p.call("put", value, key[0]))
	}
}

func (p *idbObjectStoreImpl) Add(value interface{}, key ...interface{}) IDBRequest {
	switch len(key) {
	case 0:
		return wrapIDBRequest(p.call("add", value))
	default:
		return wrapIDBRequest(p.call("add", value, key[0]))
	}
}

func (p *idbObjectStoreImpl) Delete(query interface{}) IDBRequest {
	return wrapIDBRequest(p.call("delete", query))
}

func (p *idbObjectStoreImpl) Clear() IDBRequest {
	return wrapIDBRequest(p.call("clear"))
}

func (p *idbObjectStoreImpl) Get(query interface{}) IDBRequest {
	return wrapIDBRequest(p.call("get", query))
}

func (p *idbObjectStoreImpl) Key(query interface{}) IDBRequest {
	return wrapIDBRequest(p.call("getKey", query))
}

func (p *idbObjectStoreImpl) All(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("getAll", JSValue(args[0])))
	case 2:
		if count, ok := args[1].(uint); ok {
			return wrapIDBRequest(p.call("getAll", JSValue(args[0]), count))
		}
	}

	return wrapIDBRequest(p.call("getAll"))
}

func (p *idbObjectStoreImpl) AllKeys(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("getAllKeys", JSValue(args[0])))
	case 2:
		if count, ok := args[1].(uint); ok {
			return wrapIDBRequest(p.call("getAllKeys", JSValue(args[0]), count))
		}
	}

	return wrapIDBRequest(p.call("getAllKeys"))
}

func (p *idbObjectStoreImpl) Count(query ...interface{}) IDBRequest {
	switch len(query) {
	case 0:
		return wrapIDBRequest(p.call("count"))
	default:
		return wrapIDBRequest(p.call("count", JSValue(query[0])))
	}
}

func (p *idbObjectStoreImpl) OpenCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("openCursor", JSValue(args[0])))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.call("openCursor", JSValue(args[0]), string(direction)))
		}
	}

	return wrapIDBRequest(p.call("openCursor"))
}

func (p *idbObjectStoreImpl) OpenKeyCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("openKeyCursor", JSValue(args[0])))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.call("openKeyCursor", JSValue(args[0]), string(direction)))
		}
	}

	return wrapIDBRequest(p.call("openKeyCursor"))
}

func (p *idbObjectStoreImpl) Index(name string) IDBIndex {
	return wrapIDBIndex(p.call("index", name))
}

func (p *idbObjectStoreImpl) CreateIndex(name string, keyPath string, options ...IDBIndexParameters) IDBIndex {
	switch len(options) {
	case 0:
		return wrapIDBIndex(p.call("createIndex", name, keyPath))
	default:
		return wrapIDBIndex(p.call("createIndex", name, keyPath, options[0].toJSObject()))
	}
}

func (p *idbObjectStoreImpl) DeleteIndex(name string) {
	p.call("deleteIndex", name)
}

// -------------8<---------------------------------------

type idbIndexImpl struct {
	Value
}

func wrapIDBIndex(v Value) IDBIndex {
	if v.valid() {
		return &idbIndexImpl{
			Value: v,
		}
	}
	return nil
}

func (p *idbIndexImpl) Name() string {
	return p.get("name").toString()
}

func (p *idbIndexImpl) SetName(name string) {
	p.set("name", name)
}

func (p *idbIndexImpl) ObjectStore() IDBObjectStore {
	return wrapIDBObjectStore(p.get("objectStore"))
}

func (p *idbIndexImpl) KeyPath() string {
	return p.get("keyPath").toString()
}

func (p *idbIndexImpl) MultiEntry() bool {
	return p.get("multiEntry").toBool()
}

func (p *idbIndexImpl) Unique() bool {
	return p.get("unique").toBool()
}

func (p *idbIndexImpl) Get(query interface{}) IDBRequest {
	return wrapIDBRequest(p.call("get", query))
}

func (p *idbIndexImpl) Key(query interface{}) IDBRequest {
	return wrapIDBRequest(p.call("getKey", query))
}

func (p *idbIndexImpl) All(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("getAll", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.call("getAll", args[0], count))
		}
	}

	return wrapIDBRequest(p.call("getAll"))
}

func (p *idbIndexImpl) AllKeys(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("getAllKeys", args[0]))
	case 2:
		if count, ok := args[1].(int); ok {
			return wrapIDBRequest(p.call("getAllKeys", args[0], count))
		}
	}

	return wrapIDBRequest(p.call("getAllKeys"))
}

func (p *idbIndexImpl) Count(query ...interface{}) IDBRequest {
	switch len(query) {
	case 0:
		return wrapIDBRequest(p.call("count"))
	default:
		return wrapIDBRequest(p.call("count", query[0]))
	}
}

func (p *idbIndexImpl) OpenCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("openCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.call("openCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.call("openCursor"))
}

func (p *idbIndexImpl) OpenKeyCursor(args ...interface{}) IDBRequest {
	switch len(args) {
	case 1:
		return wrapIDBRequest(p.call("openKeyCursor", args[0]))
	case 2:
		if direction, ok := args[1].(IDBCursorDirection); ok {
			return wrapIDBRequest(p.call("openKeyCursor", args[0], string(direction)))
		}
	}

	return wrapIDBRequest(p.call("openKeyCursor"))
}

// -------------8<---------------------------------------

type idbKeyRangeImpl struct {
	Value
}

func wrapIDBKeyRange(v Value) IDBKeyRange {
	if v.valid() {
		return &idbKeyRangeImpl{
			Value: v,
		}
	}
	return nil
}

func (p *idbKeyRangeImpl) Lower() interface{} {
	return Wrap(p.get("lower"))
}

func (p *idbKeyRangeImpl) Upper() interface{} {
	return Wrap(p.get("upper"))
}

func (p *idbKeyRangeImpl) LowerOpen() bool {
	return p.get("lowerOpen").toBool()
}

func (p *idbKeyRangeImpl) UpperOpen() bool {
	return p.get("upperOpen").toBool()
}

//static
func (p *idbKeyRangeImpl) Only(value interface{}) IDBKeyRange {
	return wrapIDBKeyRange(p.call("only", value))
}

//static
func (p *idbKeyRangeImpl) LowerBound(lower interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 0:
		return wrapIDBKeyRange(p.call("lowerBound", lower))
	default:
		return wrapIDBKeyRange(p.call("lowerBound", lower, open[0]))
	}
}

//static
func (p *idbKeyRangeImpl) UpperBound(upper interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 0:
		return wrapIDBKeyRange(p.call("upperBound", upper))
	default:
		return wrapIDBKeyRange(p.call("upperBound", upper, open[0]))
	}
}

//static
func (p *idbKeyRangeImpl) Bound(lower interface{}, upper interface{}, open ...bool) IDBKeyRange {
	switch len(open) {
	case 1:
		return wrapIDBKeyRange(p.call("bound", lower, upper, open[0]))
	case 2:
		return wrapIDBKeyRange(p.call("bound", lower, upper, open[0], open[1]))
	default:
		return wrapIDBKeyRange(p.call("bound", lower, upper))
	}
}

func (p *idbKeyRangeImpl) Includes(key interface{}) bool {
	return p.call("_includes", key).toBool()
}

// -------------8<---------------------------------------

type idbCursorImpl struct {
	Value
}

func wrapIDBCursor(v Value) IDBCursor {
	if p := newIDBCursorImpl(v); p != nil {
		return p
	}
	return nil
}

func newIDBCursorImpl(v Value) *idbCursorImpl {
	if v.valid() {
		return &idbCursorImpl{
			Value: v,
		}
	}
	return nil
}

func (p *idbCursorImpl) Source() IDBCursorSource {
	return wrapIDBCursorSource(p.get("source"))
}

func (p *idbCursorImpl) Direction() IDBCursorDirection {
	return IDBCursorDirection(p.get("direction").toString())
}

func (p *idbCursorImpl) Key() interface{} {
	return Wrap(p.get("key"))
}

func (p *idbCursorImpl) PrimaryKey() interface{} {
	return Wrap(p.get("primaryKey"))
}

func (p *idbCursorImpl) Advance(count uint) {
	p.call("advance", count)
}

func (p *idbCursorImpl) Continue(key ...interface{}) {
	switch len(key) {
	case 0:
		p.call("continue")
	default:
		p.call("continue", key[0])
	}
}

func (p *idbCursorImpl) ContinuePrimaryKey(key interface{}, primaryKey interface{}) {
	p.call("continuePrimaryKey", key, primaryKey)
}

func (p *idbCursorImpl) Update(value interface{}) IDBRequest {
	return wrapIDBRequest(p.call("update", value))
}

func (p *idbCursorImpl) Delete() IDBRequest {
	return wrapIDBRequest(p.call("delete"))
}

// -------------8<---------------------------------------
// (IDBObjectStore or IDBIndex)
type idbCursorSourceImpl struct {
	Value
}

func wrapIDBCursorSource(v Value) IDBCursorSource {
	if v.valid() {
		return &idbCursorSourceImpl{
			Value: v,
		}
	}
	return nil
}

// -------------8<---------------------------------------

type idbCursorWithValueImpl struct {
	*idbCursorImpl
}

func wrapIDBCursorWithValue(v Value) IDBCursorWithValue {
	if v.valid() {
		return &idbCursorWithValueImpl{
			idbCursorImpl: newIDBCursorImpl(v),
		}
	}
	return nil
}

func (p *idbCursorWithValueImpl) Value() interface{} {
	return Wrap(p.get("value"))
}

// -------------8<---------------------------------------

type idbTransactionImpl struct {
	*eventTargetImpl
}

func wrapIDBTransaction(v Value) IDBTransaction {
	if v.valid() {
		return &idbTransactionImpl{
			eventTargetImpl: newEventTargetImpl(v),
		}
	}
	return nil
}

func (p *idbTransactionImpl) ObjectStoreNames() []string {
	return domStringListToSlice(p.get("objectStoreNames"))
}

func (p *idbTransactionImpl) Mode() IDBTransactionMode {
	return IDBTransactionMode(p.get("mode").toString())
}

func (p *idbTransactionImpl) DB() IDBDatabase {
	return wrapIDBDatabase(p.get("db"))
}

func (p *idbTransactionImpl) Error() DOMException {
	return wrapDOMException(p.get("error"))
}

func (p *idbTransactionImpl) ObjectStore(name string) IDBObjectStore {
	return wrapIDBObjectStore(p.call("objectStore", name))
}

func (p *idbTransactionImpl) Abort() {
	p.call("abort")
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
	if jsIDBVersionChangeEvent := jsGlobal.get("IDBVersionChangeEvent"); jsIDBVersionChangeEvent.valid() {
		switch len(vce) {
		case 0:
			return wrapIDBVersionChangeEvent(jsIDBVersionChangeEvent.jsNew(typ))
		default:
			return wrapIDBVersionChangeEvent(jsIDBVersionChangeEvent.jsNew(typ, vce[0].toJSObject()))
		}
	}
	return nil
}
