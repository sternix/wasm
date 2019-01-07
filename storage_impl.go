// +build js,wasm

package wasm

// -------------8<---------------------------------------

func NewStorageEvent(typ string, sei ...StorageEventInit) StorageEvent {
	if jsStorageEvent := jsGlobal.Get("StorageEvent"); jsStorageEvent.Valid() {
		switch len(sei) {
		case 0:
			return wrapStorageEvent(jsStorageEvent.New(typ))
		default:
			return wrapStorageEvent(jsStorageEvent.New(typ, sei[0].toJSObject()))
		}
	}
	return nil
}

// -------------8<---------------------------------------

type storageImpl struct {
	Value
}

func wrapStorage(v Value) Storage {
	if v.Valid() {
		return &storageImpl{
			Value: v,
		}
	}
	return nil
}

func (p *storageImpl) SetItem(key, value string) {
	p.Call("setItem", key, value)
}

func (p *storageImpl) Item(key string) string {
	return p.Call("getItem", key).String()
}

func (p *storageImpl) Clear() {
	p.Call("clear")
}

func (p *storageImpl) Length() int {
	return p.Call("length").Int()
}

func (p *storageImpl) Key(i int) string {
	return p.Call("key", i).String()
}

func (p *storageImpl) RemoveItem(key string) {
	p.Call("removeItem", key)
}

// -------------8<---------------------------------------

type storageEventImpl struct {
	*eventImpl
}

func wrapStorageEvent(v Value) StorageEvent {
	if v.Valid() {
		return &storageEventImpl{
			eventImpl: newEventImpl(v),
		}
	}
	return nil
}

func (p *storageEventImpl) Key() string {
	return p.Get("key").String()
}

func (p *storageEventImpl) OldValue() string {
	return p.Get("oldValue").String()
}

func (p *storageEventImpl) NewValue() string {
	return p.Get("newValue").String()
}

func (p *storageEventImpl) Url() string {
	return p.Get("url").String()
}

func (p *storageEventImpl) StorageArea() Storage {
	return wrapStorage(p.Get("storageArea"))
}

// -------------8<---------------------------------------
