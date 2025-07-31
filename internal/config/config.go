package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig(dbPath string) (*Config, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &Config{DB: db}, nil
}
