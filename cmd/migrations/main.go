package main

import (
	"api-server/internal/app/article/infrastructure/persistence"
	"api-server/internal/pkg/config"
	"api-server/internal/pkg/db"
	"fmt"
)

func main() {
	config.InitConfig()
	db := db.ConnectDB()
	err := db.AutoMigrate(&persistence.Article{})
	if err != nil {
		fmt.Println(fmt.Errorf("migration error. %w", err))
	}
}
