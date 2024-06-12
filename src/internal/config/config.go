package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	API_NAME    string `env:"API_NAME,required"`
	API_VERSION string `env:"API_VERSION,required"`
}

var Cfg = Config{}

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	err = env.Parse(&Cfg) // 👈 Parse environment variables into `Config`

	if err != nil {

		log.Fatalf("unable to parse ennvironment variables: %e", err)

	}

	fmt.Println("Config:")

	fmt.Printf("API_NAME: %s\n", Cfg.API_NAME)

	fmt.Printf("API_VERSION: %d\n", Cfg.API_NAME)

}
