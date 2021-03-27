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
