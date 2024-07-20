package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitialEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Panic("No .env file found")
	}
}

func GetAPIPort() string {
	port, ok := os.LookupEnv("API_PORT")

	if !ok {
		return "8000"
	}

	return port
}

func GetMLHOST() string {
	host, ok := os.LookupEnv("ML_HOST")

	if !ok {
		return "localhost"
	}

	return host

}

func GetMLPort() string {
	port, ok := os.LookupEnv("ML_PORT")

	if !ok {
		return "8000"
	}

	return port
}
