package main

import (
	"petrostrak/banking-application/app"
	"petrostrak/banking-application/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
