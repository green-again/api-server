package config

import "os"

type Config struct {
	dbConfig
}

type dbConfig struct {
	dbHost string
	dbUser string
	dbPassword string
	dbName string
	dbPort string
}

func DBHost() string {
	return config.dbHost
}

func DBPort() string {
	return config.dbPort
}

func DBName() string {
	return config.dbName
}

func DBUser() string {
	return config.dbUser
}

func DBPassword() string {
	return config.dbPassword
}

var config Config

func InitConfig() {
	config = Config{
		dbConfig{
			dbHost:     os.Getenv("DB_HOST"),
			dbUser:     os.Getenv("DB_USER"),
			dbPassword: os.Getenv("DB_PASSWORD"),
			dbName:     os.Getenv("DB_NAME"),
			dbPort:     os.Getenv("DB_PORT"),
		},
	}
}
