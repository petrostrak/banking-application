package service

import (
	"petrostrak/banking-application/domain"
	"petrostrak/banking-application/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerReposiroty
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerReposiroty) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
