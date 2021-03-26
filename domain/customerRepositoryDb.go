package domain

import (
	"database/sql"
	"petrostrak/banking-application/errs"
	"petrostrak/banking-application/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers`
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?`
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while quering customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Error while quering customer table")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer
	customerSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?`

	if err := d.client.Get(&c, customerSql, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scaning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpexted database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
