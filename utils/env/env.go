package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	OPENAI_API_KEY string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	if OPENAI_API_KEY == "" {
		log.Fatal("ENV ERROR: OPENAI_API_KEY is not set")
		os.Exit(1)
	}
}
