package app

import "net/http"

func Start() {
	urlMapping()
	// starting the server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
