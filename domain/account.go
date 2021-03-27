package domain

import (
	"petrostrak/banking-application/dto"
	"petrostrak/banking-application/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		a.AccountID,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
