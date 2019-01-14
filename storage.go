// +build js,wasm

package wasm

type (

	// https://www.w3.org/TR/webstorage/#storage-0
	Storage interface {
		Length() int
		Key(int) string
		Item(string) string
		SetItem(string, string)
		RemoveItem(string)
		Clear()
	}

	// https://www.w3.org/TR/webstorage/#dom-sessionstorage
	WindowSessionStorage interface {
		SessionStorage() Storage
	}

	// https://www.w3.org/TR/webstorage/#dom-localstorage
	WindowLocalStorage interface {
		LocalStorage() Storage
	}

	// https://www.w3.org/TR/webstorage/#storageevent
	StorageEvent interface {
		Event

		Key() string
		OldValue() string
		NewValue() string
		Url() string
		StorageArea() Storage
	}
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/webstorage/#storageeventinit
type StorageEventInit struct {
	EventInit

	Key         string
	OldValue    string
	NewValue    string
	Url         string
	StorageArea Storage
}

func (p StorageEventInit) toJSObject() Value {
	o := p.EventInit.toJSObject()
	o.set("key", p.Key)
	o.set("oldValue", p.OldValue)
	o.set("newValue", p.NewValue)
	o.set("url", p.Url)
	o.set("storageArea", JSValue(p.StorageArea))
	return o
}
