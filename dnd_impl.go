// +build js,wasm

package wasm

// -------------8<---------------------------------------

type dataTransferImpl struct {
	Value
}

func wrapDataTransfer(v Value) DataTransfer {
	if v.Valid() {
		return &dataTransferImpl{
			Value: v,
		}
	}
	return nil
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
	return wrapDataTransferItemList(p.Get("items"))
}

func (p *dataTransferImpl) SetDragImage(image Element, x int, y int) {
	p.Call("setDragImage", JSValue(image), x, y)
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
	Value
}

func wrapDataTransferItemList(v Value) DataTransferItemList {
	if v.Valid() {
		return &dataTransferItemListImpl{
			Value: v,
		}
	}
	return nil
}

func (p *dataTransferItemListImpl) Length() int {
	return p.Get("length").Int()
}

func (p *dataTransferItemListImpl) Index(index int) DataTransferItem {
	return wrapDataTransferItem(p.Call("DataTransferItem", index))
}

func (p *dataTransferItemListImpl) Add(data string, typ string) DataTransferItem {
	return wrapDataTransferItem(p.Call("add", data, typ))
}

func (p *dataTransferItemListImpl) AddFile(data File) DataTransferItem {
	return wrapDataTransferItem(p.Call("add", JSValue(data)))
}

func (p *dataTransferItemListImpl) Remove(index int) {
	p.Call("remove", index)
}

func (p *dataTransferItemListImpl) Clear() {
	p.Call("clear")
}

// -------------8<---------------------------------------

type dataTransferItemImpl struct {
	Value
}

func wrapDataTransferItem(v Value) DataTransferItem {
	if v.Valid() {
		return &dataTransferItemImpl{
			Value: v,
		}
	}
	return nil
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
	return wrapFile(p.Call("getAsFile"))
}

// -------------8<---------------------------------------

type dragEventImpl struct {
	*mouseEventImpl
}

func wrapDragEvent(v Value) DragEvent {
	if v.Valid() {
		return &dragEventImpl{
			mouseEventImpl: newMouseEventImpl(v),
		}
	}
	return nil
}

func (p *dragEventImpl) DataTransfer() DataTransfer {
	return wrapDataTransfer(p.Get("dataTransfer"))
}

// -------------8<---------------------------------------

func NewDragEvent(typ string, dei ...DragEventInit) DragEvent {
	if jsDragEvent := jsGlobal.Get("DragEvent"); jsDragEvent.Valid() {
		switch len(dei) {
		case 0:
			return wrapDragEvent(jsDragEvent.New(typ))
		default:
			return wrapDragEvent(jsDragEvent.New(typ, dei[0].toJSObject()))
		}
	}
	return nil
}
