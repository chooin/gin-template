package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var AppDebug bool

func init() {
	var err error
	mode := os.Getenv("MODE")
	if mode == "" {
		panic("failed to load dotenv")
	} else {
		if err = godotenv.Load(mode); err != nil {
			panic("failed to load dotenv")
		}
	}
	AppDebug, err = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		panic("missing APP_DEBUG in .env file")
	}
}
