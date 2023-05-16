package configs

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func loadEnv() {
	dir := os.Getenv("GOPATH") + "/src/stocker-quant/.env"
	err := env.Load(dir)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
