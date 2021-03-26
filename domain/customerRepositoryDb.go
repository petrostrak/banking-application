package domain

import (
	"database/sql"
	"petrostrak/banking-application/errs"
	"petrostrak/banking-application/logger"
	"petrostrak/banking-application/resources"
	"time"

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

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	// dbUser := os.Getenv("DB_USER")
	// dbPasswd := os.Getenv("DB_PASSWD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_Name")

	// 	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", resources.MySQLCredentials)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
