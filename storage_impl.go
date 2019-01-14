// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewStorageEvent(typ string, sei ...StorageEventInit) StorageEvent {
	if jsStorageEvent := jsGlobal.get("StorageEvent"); jsStorageEvent.valid() {
		switch len(sei) {
		case 0:
			return wrapStorageEvent(jsStorageEvent.jsNew(typ))
		default:
			return wrapStorageEvent(jsStorageEvent.jsNew(typ, sei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type storageImpl struct {
	Value
}

func wrapStorage(v Value) Storage {
	if v.valid() {
		return &storageImpl{
			Value: v,
		}
	}
	return nil
}

func (p *storageImpl) SetItem(key, value string) {
	p.call("setItem", key, value)
}

func (p *storageImpl) Item(key string) string {
	return p.call("getItem", key).toString()
}

func (p *storageImpl) Clear() {
	p.call("clear")
}

func (p *storageImpl) Length() int {
	return p.call("length").toInt()
}

func (p *storageImpl) Key(i int) string {
	return p.call("key", i).toString()
}

func (p *storageImpl) RemoveItem(key string) {
	p.call("removeItem", key)
}

// -------------8<---------------------------------------

type storageEventImpl struct {
	*eventImpl
}

func wrapStorageEvent(v Value) StorageEvent {
	if v.valid() {
		return &storageEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *storageEventImpl) Key() string {
	return p.get("key").toString()
}

func (p *storageEventImpl) OldValue() string {
	return p.get("oldValue").toString()
}

func (p *storageEventImpl) NewValue() string {
	return p.get("newValue").toString()
}

func (p *storageEventImpl) Url() string {
	return p.get("url").toString()
}

func (p *storageEventImpl) StorageArea() Storage {
	return wrapStorage(p.get("storageArea"))
}

// -------------8<---------------------------------------
