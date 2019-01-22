// +build js,wasm

package wasm

/*
TODO: most arguments like ...interface{}
are panic'able restrict parameter to ValueOf
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
	Value
}

func wrapConsole(v Value) Console {
	if v.valid() {
		return &consoleImpl{
			Value: v,
		}
	}
	return nil
}

func (p *consoleImpl) Assert(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("assert")
	case 1:
		if condition, ok := args[0].(bool); ok {
			p.call("assert", condition)
		}
	default: // more than 1
		// TODO: WRONG
		if condition, ok := args[0].(bool); ok {
			var params []interface{}
			for i := 1; i < len(args); i++ {
				params = append(params, args[i])
			}

			p.call("assert", condition, params)
		}
	}
}

func (p *consoleImpl) Clear() {
	p.call("clear")
}

func (p *consoleImpl) Debug(data ...interface{}) {
	p.call("debug", data...)
}

func (p *consoleImpl) Error(data ...interface{}) {
	p.call("error", data...)
}

func (p *consoleImpl) Info(data ...interface{}) {
	p.call("info", data...)
}

func (p *consoleImpl) Log(data ...interface{}) {
	p.call("log", data...)
}

func (p *consoleImpl) Table(tabularData interface{}, properties ...[]string) {
	switch len(properties) {
	case 0:
		p.call("table", tabularData)
	default:
		p.call("table", tabularData, ToJSArray(properties[0]))
	}
}

func (p *consoleImpl) Trace(data ...interface{}) {
	p.call("trace", data...)
}

func (p *consoleImpl) Warn(data ...interface{}) {
	p.call("warn", data...)
}

func (p *consoleImpl) Dir(item interface{}, options ...interface{}) {
	switch len(options) {
	case 0:
		p.call("dir", item)
	default:
		p.call("dir", item, options[0])
	}
}

func (p *consoleImpl) Dirxml(data ...interface{}) {
	p.call("dirxml", data...)
}

func (p *consoleImpl) Count(label ...string) {
	switch len(label) {
	case 0:
		p.call("count")
	default:
		p.call("count", label[0])
	}
}

func (p *consoleImpl) CountReset(label ...string) {
	switch len(label) {
	case 0:
		p.call("countReset")
	default:
		p.call("countReset", label[0])
	}
}

func (p *consoleImpl) Group(data ...interface{}) {
	p.call("group", data...)
}

func (p *consoleImpl) GroupCollapsed(data ...interface{}) {
	p.call("groupCollapsed", data...)
}

func (p *consoleImpl) GroupEnd() {
	p.call("groupEnd")
}

func (p *consoleImpl) Time(label ...string) {
	switch len(label) {
	case 0:
		p.call("time")
	case 1:
		p.call("time", label[0])
	}
}

func (p *consoleImpl) TimeLog(args ...interface{}) {
	switch len(args) {
	case 0:
		p.call("timeLog")
	case 1:
		if label, ok := args[0].(string); ok {
			p.call("timeLog", label)
		}
	default:
		//TODO WRONG
		if label, ok := args[0].(string); ok {
			var params []interface{}
			for i := 1; i < len(args); i++ {
				params = append(params, args[i])
			}
			p.call("timeLog", label, params)
		}
	}
}

func (p *consoleImpl) TimeEnd(label ...string) {
	switch len(label) {
	case 0:
		p.call("timeEnd")
	default:
		p.call("timeEnd", label[0])
	}
}
