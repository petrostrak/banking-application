package domain

import (
	"database/sql"
	"log"
	"petrostrak/banking-application/resources"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers`

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while quering customer table", err.Error())
		return nil, err
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
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
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
