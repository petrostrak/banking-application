package domain

import "petrostrak/banking-application/errs"

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DataOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerReposiroty interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
