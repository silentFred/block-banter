package config

import (
	"fmt"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// TODO spin up separate docker image when this works~
func LoadConfig() Config {
	return Config{
		Host:     "localhost", //os.Getenv("DB_HOST"),
		User:     "test",      //os.Getenv("DB_USER"),
		Password: "test",      //os.Getenv("DB_PASSWORD"),
		DBName:   "wallet",    //os.Getenv("DB_NAME"),
		Port:     "5432",      //os.Getenv("DB_PORT"),
		SSLMode:  "disable",   //os.Getenv("DB_SSLMODE"),
	}
}

//migrate -database "postgres://test:test@localhost:5432/wallet?sslmode=disable" -path ./migrations up

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}
