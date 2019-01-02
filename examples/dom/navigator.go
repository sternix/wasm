// +build js,wasm

package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func TestNavigator(nav wasm.Navigator) {
	fmt.Println("---------------------")
	fmt.Println("Navigator Tests")

	fmt.Printf("Navigator.AppCodeName: %s\n", nav.AppCodeName())
	fmt.Printf("Navigator.AppName: %s\n", nav.AppName())
	fmt.Printf("Navigator.AppVersion: %s\n", nav.AppVersion())
	fmt.Printf("Navigator.Platform: %s\n", nav.Platform())
	fmt.Printf("Navigator.Product: %s\n", nav.Product())
	fmt.Printf("Navigator.UserAgent: %s\n", nav.UserAgent())

	fmt.Printf("Navigator.Language: %s\n", nav.Language())

	fmt.Print("Navigator.Languages:")
	for _, l := range nav.Languages() {
		fmt.Printf(" %s ", l)
	}
	fmt.Println()
	fmt.Printf("Navigator.CookieEnabled: %t\n", nav.CookieEnabled())

	fmt.Println("---------------------")
}
