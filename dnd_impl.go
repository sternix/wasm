// +build js,wasm

package wasm

import (
	"syscall/js"
)

// -------------8<---------------------------------------

type dataTransferImpl struct {
	js.Value
}

func newDataTransfer(v js.Value) DataTransfer {
	if isNil(v) {
		return nil
	}

	return &dataTransferImpl{
		Value: v,
	}
}

func (p *dataTransferImpl) DropEffect() string {
	return p.Get("dropEffect").String()
}

func (p *dataTransferImpl) SetDropEffect(effect string) {
	p.Set("dropEffect", effect)
}

func (p *dataTransferImpl) EffectAllowed() string {
	return p.Get("effectAllowed").String()
}

func (p *dataTransferImpl) SetEffectAllowed(param string) {
	p.Set("effectAllowed", param)
}

func (p *dataTransferImpl) Items() DataTransferItemList {
	return newDataTransferItemList(p.Get("items"))
}

func (p *dataTransferImpl) SetDragImage(image Element, x int, y int) {
	p.Call("setDragImage", image.JSValue(), x, y)
}

func (p *dataTransferImpl) Types() []string {
	return stringSequenceToSlice(p.Get("types"))
}

func (p *dataTransferImpl) Data(format string) string {
	return p.Call("getData", format).String()
}

func (p *dataTransferImpl) SetData(format string, data string) {
	p.Call("setData", format, data)
}

func (p *dataTransferImpl) ClearData(format ...string) {
	switch len(format) {
	case 0:
		p.Call("clearData")
	default:
		p.Call("clearData", format[0])
	}
}

func (p *dataTransferImpl) Files() []File {
	return fileListToSlice(p.Get("files"))
}

// -------------8<---------------------------------------

type dataTransferItemListImpl struct {
	js.Value
}

func newDataTransferItemList(v js.Value) DataTransferItemList {
	if isNil(v) {
		return nil
	}

	return &dataTransferItemListImpl{
		Value: v,
	}
}

func (p *dataTransferItemListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *dataTransferItemListImpl) Index(index int) DataTransferItem {
	return newDataTransferItem(p.Call("DataTransferItem", index))
}

func (p *dataTransferItemListImpl) Add(data string, typ string) DataTransferItem {
	return newDataTransferItem(p.Call("add", data, typ))
}

func (p *dataTransferItemListImpl) AddFile(data File) DataTransferItem {
	return newDataTransferItem(p.Call("add", data.JSValue()))
}

func (p *dataTransferItemListImpl) Remove(index int) {
	p.Call("remove", index)
}

func (p *dataTransferItemListImpl) Clear() {
	p.Call("clear")
}

// -------------8<---------------------------------------

type dataTransferItemImpl struct {
	js.Value
}

func newDataTransferItem(v js.Value) DataTransferItem {
	if isNil(v) {
		return nil
	}

	return &dataTransferItemImpl{
		Value: v,
	}
}

func (p *dataTransferItemImpl) Kind() string {
	return p.Get("kind").String()
}

func (p *dataTransferItemImpl) Type() string {
	return p.Get("type").String()
}

func (p *dataTransferItemImpl) AsString(cb FunctionStringCallback) {
	p.Call("getAsString", cb.jsCallback())
}

func (p *dataTransferItemImpl) AsFile() File {
	return newFile(p.Call("getAsFile"))
}

// -------------8<---------------------------------------

type dragEventImpl struct {
	*mouseEventImpl
}

func newDragEvent(v js.Value) DragEvent {
	if isNil(v) {
		return nil
	}

	return &dragEventImpl{
		mouseEventImpl: newMouseEventImpl(v),
	}
}

func (p *dragEventImpl) DataTransfer() DataTransfer {
	return newDataTransfer(p.Get("dataTransfer"))
}

// -------------8<---------------------------------------

func NewDragEvent(typ string, dei ...DragEventInit) DragEvent {
	jsDragEvent := js.Global().Get("DragEvent")
	if isNil(jsDragEvent) {
		return nil
	}

	if len(dei) > 0 {
		return newDragEvent(jsDragEvent.New(typ, dei[0].toDict()))
	}

	return newDragEvent(jsDragEvent.New(typ))
}
