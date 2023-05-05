package configs

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func LoadEnv() {
	dir := os.Getenv("GOPATH") + "/src/expense-logger/.env"
	err := env.Load(dir)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
