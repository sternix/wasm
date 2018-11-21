// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

func NewStorageEvent(typ string, sei ...StorageEventInit) StorageEvent {
	jsStorageEvent := js.Global().Get("StorageEvent")
	if isNil(jsStorageEvent) {
		return nil
	}

	switch len(sei) {
	case 0:
		return newStorageEvent(jsStorageEvent.New(typ))
	default:
		return newStorageEvent(jsStorageEvent.New(typ, toJSONObject(sei[0])))
	}
}

// -------------8<---------------------------------------

type storageImpl struct {
	js.Value
}

func newStorage(v js.Value) Storage {
	if isNil(v) {
		return nil
	}

	return &storageImpl{
		Value: v,
	}
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

func newStorageEvent(v js.Value) StorageEvent {
	if isNil(v) {
		return nil
	}

	return &storageEventImpl{
		eventImpl: newEventImpl(v),
	}
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
	return newStorage(p.Get("storageArea"))
}

// -------------8<---------------------------------------
