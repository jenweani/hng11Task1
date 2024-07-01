package configs

import (
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
