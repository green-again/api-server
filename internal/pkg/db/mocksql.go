package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func NewMockSQL() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	mockEngine, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return mockEngine, mock
}
