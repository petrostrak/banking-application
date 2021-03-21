package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Petros", City: "Athens", Zipcode: "17456"},
		{Name: "Maggie", City: "Athens", Zipcode: "17456"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting the server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
