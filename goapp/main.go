package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi from aws codepipeline  %s\n", r.URL.Path)
	})

	http.ListenAndServe(":81", nil)
}
