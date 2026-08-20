package dotenv

import (
	"github.com/joho/godotenv"
)

type Config map[string]string

func Load() (Config, error) {
	return godotenv.Read(".env")
}
