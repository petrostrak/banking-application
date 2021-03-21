package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)

	// starting the server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
