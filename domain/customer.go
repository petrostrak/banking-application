package domain

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DataOfBirth string
	Status      string
}

type CustomerReposiroty interface {
	FindAll() ([]Customer, error)
}
