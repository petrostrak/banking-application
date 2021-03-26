package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviroment variable not defined")
	}
}

func urlMapping() {
	sanityCheck()
	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting the server SERVER_ADDRESS=localhost SERVER_PORT=8282 ./banking-application
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router); err != nil {
		panic(err)
	}
}
