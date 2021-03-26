package domain

import "petrostrak/banking-application/errs"

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AcountRepository interface {
	Save(Account) (*Account, errs.AppError)
}
