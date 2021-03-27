package domain

import (
	"petrostrak/banking-application/dto"
	"petrostrak/banking-application/errs"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		a.AccountID,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	// if a.Amount < amount {
	// 	return false
	// }
	return a.Amount < amount
}
