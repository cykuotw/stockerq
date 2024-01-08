package configs

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func loadEnv() {
	exeMode := os.Getenv("executeMode")
	dir := ""
	if exeMode == "Debug" || exeMode == "" {
		dir = os.Getenv("GOPATH") + "/src/stockerq/.env"
	} else if exeMode == "Release" {
		dir = "./.env"
	}
	err := env.Load(dir)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
