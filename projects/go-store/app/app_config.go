package app

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/raflynagachi/go-store/helpers"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (appConfig *AppConfig) setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}
	appConfig.AppName = helpers.GetEnv("APP_NAME", "Go-Store")
	appConfig.AppEnv = helpers.GetEnv("APP_ENV", "development")
	appConfig.AppPort = helpers.GetEnv("APP_PORT", "8080")

}
