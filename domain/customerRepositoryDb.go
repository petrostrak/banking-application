package domain

import (
	"database/sql"
	"log"
	"petrostrak/banking-application/errs"
	"petrostrak/banking-application/resources"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers`
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?`
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		log.Println("Error while quering customer table", err.Error())
		return nil, errs.NewUnexpectedError("Error while quering customer table")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.City,
			&c.DataOfBirth,
			&c.Zipcode,
			&c.Status,
		); err != nil {
			log.Println("Error while reading rows", err.Error())
			return nil, errs.NewUnexpectedError("Error while reading rows")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?`

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	if err := row.Scan(
		&c.ID,
		&c.Name,
		&c.City,
		&c.DataOfBirth,
		&c.Zipcode,
		&c.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scaning customer", err.Error())
		return nil, errs.NewUnexpectedError("Unexpexted database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	client, err := sql.Open("mysql", resources.MySQLCredentials)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
