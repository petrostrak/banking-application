package app

import (
	"net/http"
	"petrostrak/banking-application/domain"
	"petrostrak/banking-application/service"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

func urlMapping() {
	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting the server
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
