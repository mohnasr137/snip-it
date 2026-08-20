package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() Config {
	godotenv.Load(".env")

	return Config{
		Port: os.Getenv("PORT"),
	}
}
