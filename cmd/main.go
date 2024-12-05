package cmd

import (
	"Marketplace/internal/config"
	"Marketplace/internal/di"
	"log"
	"os"

	_ "Marketplace/docs"
)

// @title Marketplace
// @version 1.0
// @description API Server for Marketplace for influences and companies
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal(diErr)
	} else {
		server.Run(infoLog, errorLog)
	}
}
