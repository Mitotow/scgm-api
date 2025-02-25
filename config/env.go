package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type EnvironmentVariables struct {
	AppPort    string `envconfig:"APP_PORT" default:"8080"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USERNAME" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
}

var variables EnvironmentVariables

func init() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find .env file, error : " + err.Error())
	}

	err = envconfig.Process("", &variables)
	if err != nil {
		log.Fatal("Cannot config environment variables, error : " + err.Error())
	}
}

func GetEnv() *EnvironmentVariables {
	return &variables
}
