package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"api-server/internal/pkg/config"
)

func ConnectDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		cfg.DBUser(),
		cfg.DBPassword(),
		cfg.DBHost(),
		cfg.DBPort(),
		cfg.DBName(),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB connection failed. %w", err))
	}
	return db
}
