package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"petrostrak/banking-application/domain"
	"petrostrak/banking-application/resources"
	"petrostrak/banking-application/service"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviroment variable not defined")
	}
}

func urlMapping() {
	sanityCheck()
	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)
	// accountRepositoryDb := domain.NewAccountRepository(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// starting the server SERVER_ADDRESS=localhost SERVER_PORT=8282 ./banking-application
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router); err != nil {
		panic(err)
	}
}

func getDbClient() *sqlx.DB {
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
	return client
}
