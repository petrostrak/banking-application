package main

import (
	"petrostrak/banking-application/app"
	"petrostrak/banking-application/logger"
)

// SERVER_ADDRESS=localhost SERVER_PORT=8000 ./banking-application
func main() {
	logger.Info("Starting the application...")
	app.Start()
}
