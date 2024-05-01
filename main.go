package main

import (
	"github.com/topdata-software-gmbh/topdata-package-service-cli/cmd"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/pkg"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var client *pkg.topdataPackageServiceClient

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	client = pkg.NewClient(baseURL, username, password)

	cmd.Execute()
}

func main() {
    cmd.Execute()
}
