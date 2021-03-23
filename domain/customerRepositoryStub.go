package domain

type CustomerReposirotyStub struct {
	customers []Customer
}

func (s CustomerReposirotyStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerReposirotyStub {
	customers := []Customer{
		{"1001", "Petros", "Athens", "17456", "1986-05-10", "1"},
		{"1001", "Maggie", "Athens", "17456", "1984-03-8", "1"},
	}
	return CustomerReposirotyStub{customers: customers}
}
