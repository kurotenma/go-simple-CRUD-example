package configTools

import (
	"github.com/joho/godotenv"
	"os"
)

func Init() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}
	env := os.Getenv("ENV_MODE")
	if env != "DEVELOPMENT" {
		err = godotenv.Load(".env.production")
	} else {
		err = godotenv.Load(".env")
	}
	return
}
