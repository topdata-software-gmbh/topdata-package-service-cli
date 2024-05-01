package main

import (
	"github.com/joho/godotenv"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/cmd"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/pkg"
	"log"
	"os"
)

var client *pkg.TopdataPackageServiceClient

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	baseURL := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	client = pkg.NewClient(baseURL, username, password)

	cmd.Execute()
}
