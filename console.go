// +build js,wasm

package wasm

import (
	"syscall/js"
)

/*
TODO: most arguments like ...interface{}
are panic'able restrict parameter to js.ValueOf
*/

// https://console.spec.whatwg.org/

type (
	// https://console.spec.whatwg.org/#namespacedef-console
	Console interface {
		Assert(...interface{})
		Clear()
		Debug(...interface{})
		Error(...interface{})
		Info(...interface{})
		Log(...interface{})
		Table(interface{}, ...[]string)
		Trace(...interface{})
		Warn(...interface{})
		Dir(interface{}, ...interface{})
		Dirxml(...interface{})
		Count(...string)
		CountReset(...string)
		Group(...interface{})
		GroupCollapsed(...interface{})
		GroupEnd()
		Time(...string)
		TimeLog(...interface{})
		TimeEnd(...string)
	}
)

// -------------8<---------------------------------------

type consoleImpl struct {
	js.Value
}

func wrapConsole(v js.Value) Console {
	if isNil(v) {
		return nil
	}

	return &consoleImpl{
		Value: v,
	}
}

func (p *consoleImpl) Assert(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("assert")
	case 1:
		if condition, ok := args[0].(bool); ok {
			p.Call("assert", condition)
		}
	default: // more than 1
		// TODO: WRONG
		if condition, ok := args[0].(bool); ok {
			var params []interface{}
			for i := 1; i < len(args); i++ {
				params = append(params, args[i])
			}

			p.Call("assert", condition, params)
		}
	}
}

func (p *consoleImpl) Clear() {
	p.Call("clear")
}

func (p *consoleImpl) Debug(data ...interface{}) {
	p.Call("debug", data...)
}

func (p *consoleImpl) Error(data ...interface{}) {
	p.Call("error", data...)
}

func (p *consoleImpl) Info(data ...interface{}) {
	p.Call("info", data...)
}

func (p *consoleImpl) Log(data ...interface{}) {
	p.Call("log", data...)
}

func (p *consoleImpl) Table(tabularData interface{}, properties ...[]string) {
	switch len(properties) {
	case 0:
		p.Call("table", tabularData)
	default:
		p.Call("table", tabularData, sliceToJsArray(properties[0]))
	}
}

func (p *consoleImpl) Trace(data ...interface{}) {
	p.Call("trace", data...)
}

func (p *consoleImpl) Warn(data ...interface{}) {
	p.Call("warn", data...)
}

func (p *consoleImpl) Dir(item interface{}, options ...interface{}) {
	switch len(options) {
	case 0:
		p.Call("dir", item)
	default:
		p.Call("dir", item, options[0])
	}
}

func (p *consoleImpl) Dirxml(data ...interface{}) {
	p.Call("dirxml", data...)
}

func (p *consoleImpl) Count(label ...string) {
	switch len(label) {
	case 0:
		p.Call("count")
	default:
		p.Call("count", label[0])
	}
}

func (p *consoleImpl) CountReset(label ...string) {
	switch len(label) {
	case 0:
		p.Call("countReset")
	default:
		p.Call("countReset", label[0])
	}
}

func (p *consoleImpl) Group(data ...interface{}) {
	p.Call("group", data...)
}

func (p *consoleImpl) GroupCollapsed(data ...interface{}) {
	p.Call("groupCollapsed", data...)
}

func (p *consoleImpl) GroupEnd() {
	p.Call("groupEnd")
}

func (p *consoleImpl) Time(label ...string) {
	switch len(label) {
	case 0:
		p.Call("time")
	case 1:
		p.Call("time", label[0])
	}
}

func (p *consoleImpl) TimeLog(args ...interface{}) {
	switch len(args) {
	case 0:
		p.Call("timeLog")
	case 1:
		if label, ok := args[0].(string); ok {
			p.Call("timeLog", label)
		}
	default:
		//TODO WRONG
		if label, ok := args[0].(string); ok {
			var params []interface{}
			for i := 1; i < len(args); i++ {
				params = append(params, args[i])
			}
			p.Call("timeLog", label, params)
		}
	}
}

func (p *consoleImpl) TimeEnd(label ...string) {
	switch len(label) {
	case 0:
		p.Call("timeEnd")
	default:
		p.Call("timeEnd", label[0])
	}
}
