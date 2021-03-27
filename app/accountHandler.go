package app

import (
	"encoding/json"
	"net/http"
	"petrostrak/banking-application/dto"
	"petrostrak/banking-application/service"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	var req dto.NewAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerID = customerID
		acc, appErr := ah.service.NewAccount(req)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, acc)
		}
	}
}

func (ah AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := ah.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
