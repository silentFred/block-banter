package database

import (
	"block-banter/config"
	"block-banter/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	cfg := config.LoadConfig()
	dsn := cfg.DSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&repository.TransferEvent{})
}
