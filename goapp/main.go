package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":81", nil)
}
