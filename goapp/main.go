package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		fmt.Fprintf(w, "hi %s\n", r.URL.Path)
=======
		fmt.Fprintf(w, "amazon  %s\n", r.URL.Path)
>>>>>>> 6e6ebf6a822b03d14a292d7350f8c4abc5487387
	})

	http.ListenAndServe(":81", nil)
}
