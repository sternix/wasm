// +build js,wasm

package wasm

// -------------8<---------------------------------------

type dataTransferImpl struct {
	Value
}

func wrapDataTransfer(v Value) DataTransfer {
	if v.valid() {
		return &dataTransferImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataTransferImpl) DropEffect() string {
	return p.get("dropEffect").toString()
}

func (p *dataTransferImpl) SetDropEffect(effect string) {
	p.set("dropEffect", effect)
}

func (p *dataTransferImpl) EffectAllowed() string {
	return p.get("effectAllowed").toString()
}

func (p *dataTransferImpl) SetEffectAllowed(param string) {
	p.set("effectAllowed", param)
}

func (p *dataTransferImpl) Items() DataTransferItemList {
	return wrapDataTransferItemList(p.get("items"))
}

func (p *dataTransferImpl) SetDragImage(image Element, x int, y int) {
	p.call("setDragImage", JSValueOf(image), x, y)
}

func (p *dataTransferImpl) Types() []string {
	return stringSequenceToSlice(p.get("types"))
}

func (p *dataTransferImpl) Data(format string) string {
	return p.call("getData", format).toString()
}

func (p *dataTransferImpl) SetData(format string, data string) {
	p.call("setData", format, data)
}

func (p *dataTransferImpl) ClearData(format ...string) {
	switch len(format) {
	case 0:
		p.call("clearData")
	default:
		p.call("clearData", format[0])
	}
}

func (p *dataTransferImpl) Files() []File {
	return fileListToSlice(p.get("files"))
}

// -------------8<---------------------------------------

type dataTransferItemListImpl struct {
	Value
}

func wrapDataTransferItemList(v Value) DataTransferItemList {
	if v.valid() {
		return &dataTransferItemListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataTransferItemListImpl) Length() uint {
	return p.get("length").toUint()
}

func (p *dataTransferItemListImpl) Index(index uint) DataTransferItem {
	return wrapDataTransferItem(p.call("DataTransferItem", index))
}

func (p *dataTransferItemListImpl) Add(data string, typ string) DataTransferItem {
	return wrapDataTransferItem(p.call("add", data, typ))
}

func (p *dataTransferItemListImpl) AddFile(data File) DataTransferItem {
	return wrapDataTransferItem(p.call("add", JSValueOf(data)))
}

func (p *dataTransferItemListImpl) Remove(index uint) {
	p.call("remove", index)
}

func (p *dataTransferItemListImpl) Clear() {
	p.call("clear")
}

// -------------8<---------------------------------------

type dataTransferItemImpl struct {
	Value
}

func wrapDataTransferItem(v Value) DataTransferItem {
	if v.valid() {
		return &dataTransferItemImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataTransferItemImpl) Kind() string {
	return p.get("kind").toString()
}

func (p *dataTransferItemImpl) Type() string {
	return p.get("type").toString()
}

func (p *dataTransferItemImpl) AsString(cb FunctionStringCallback) {
	p.call("getAsString", cb.jsCallback())
}

func (p *dataTransferItemImpl) AsFile() File {
	return wrapFile(p.call("getAsFile"))
}

// -------------8<---------------------------------------

type dragEventImpl struct {
	*mouseEventImpl
}

func wrapDragEvent(v Value) DragEvent {
	if v.valid() {
		return &dragEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *dragEventImpl) DataTransfer() DataTransfer {
	return wrapDataTransfer(p.get("dataTransfer"))
}

// -------------8<---------------------------------------

func NewDragEvent(typ string, dei ...DragEventInit) DragEvent {
	if jsDragEvent := jsGlobal.get("DragEvent"); jsDragEvent.valid() {
		switch len(dei) {
		case 0:
			return wrapDragEvent(jsDragEvent.jsNew(typ))
		default:
			return wrapDragEvent(jsDragEvent.jsNew(typ, dei[0].JSValue()))
		}
	}
	return nil
}
