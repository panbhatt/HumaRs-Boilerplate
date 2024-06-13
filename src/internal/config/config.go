package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
THis is the Config Struct that will be used across the Application
*/

type Config struct {
	API_NAME    string `env:"API_NAME,required"`
	API_VERSION string `env:"API_VERSION,required"`
	API_PORT    string `env:"API_PORT,required"`
	ROUTER      string `env:"ROUTER,required"`

	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_HOST     string `env:"DB_HOST,required"`
	DB_PORT     string `env:"DB_PORT,required"`
	DB_NAME     string `env:"DB_NAME,required"`

	DB *gorm.DB
}

var Cfg = Config{}

func Init() {
	initConfig()
	initDB()
}

func initConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	err = env.Parse(&Cfg) // 👈 Parse environment variables into `Config`

	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	slog.Info("Config:", "API_NAME", Cfg.API_NAME, "API_VERSION", Cfg.API_VERSION, "API_PORT", Cfg.API_PORT)

}

/*
*
This function will initial the SLOGGER
*/
func initLogger() {

}

func initDB() {
	// Create an slog-gorm instance
	gormLogger := slogGorm.New() // use slog.Default() by default
	var dsn string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", Cfg.DB_USER, Cfg.DB_PASSWORD, Cfg.DB_HOST, Cfg.DB_PORT, Cfg.DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		slog.Error("An Error occured, while connectiong to the Database. ")
		os.Exit(1)
	}

	// Setting Various DB Properties.
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	Cfg.DB = db

	slog.Info("Database has been successfully connected ", "URL", dsn)

}
