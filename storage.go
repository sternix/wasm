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

	// https://www.w3.org/TR/webstorage/#storageeventinit
	StorageEventInit struct {
		EventInit

		Key         string  `json:"key"`
		OldValue    string  `json:"oldValue"`
		NewValue    string  `json:"newValue"`
		Url         string  `json:"url"`
		StorageArea Storage `json:"storageArea"`
	}
)
