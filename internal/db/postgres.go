package db

import (
	"fmt"

	"github.com/priyanshu334/go_car_book/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBUser,
		config.Cfg.DBPass,
		config.Cfg.DBName,
		config.Cfg.DBSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
