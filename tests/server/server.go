package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Path[1:]) > 0 {
		http.ServeFile(w, r, r.URL.Path[1:])
	} else {
		http.ServeFile(w, r, "wasm_exec.html")
	}
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
