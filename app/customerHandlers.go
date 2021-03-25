package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"petrostrak/banking-application/service"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Petros", City: "Athens", Zipcode: "17456"},
	// 	{Name: "Maggie", City: "Athens", Zipcode: "17456"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		// xml encoder
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// json encoder
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	// pass the Id to the service
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
