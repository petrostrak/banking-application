package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

func urlMapping() {
	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting the server
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
