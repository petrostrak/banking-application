package domain

import "petrostrak/banking-application/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DataOfBirth string
	Status      string
}

type CustomerReposiroty interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
