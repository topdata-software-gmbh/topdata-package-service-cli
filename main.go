package main

import (
	"github.com/joho/godotenv"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/commands"
	"github.com/topdata-software-gmbh/topdata-package-service-cli/pkg"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	baseURL := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	topdataPackageServiceClient := pkg.NewTopdataPackageServiceClient(baseURL, username, password)

	commands.Execute(topdataPackageServiceClient)
}
