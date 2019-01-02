// +build js,wasm
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func TestConsole(console wasm.Console) {
	fmt.Println("----------------")
	fmt.Println("Console Tests")
	console.Log("Log1")
	console.Log("Log1", "Log2")
	console.Log("Log1", 1)
	console.Log("Log1", 1, 12.5, true)
	console.Clear()
	//console.Assert(2 != 3, "2 ")
	console.Log(1)
	console.Log("test")
	console.Log(true)
	console.Log(124.54)
	console.Log()
	fmt.Println("----------------")
}
