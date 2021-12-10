package config

import "os"

type Config struct {
	dbConfig
}

type dbConfig struct {
	dbHost     string
	dbUser     string
	dbPassword string
	dbName     string
	dbPort     string
}

func (c Config) DBHost() string {
	return c.dbHost
}

func (c Config) DBPort() string {
	return c.dbPort
}

func (c Config) DBName() string {
	return c.dbName
}

func (c Config) DBUser() string {
	return c.dbUser
}

func (c Config) DBPassword() string {
	return c.dbPassword
}

func Make() Config {
	return Config{
		dbConfig{
			dbHost:     os.Getenv("DB_HOST"),
			dbUser:     os.Getenv("DB_USER"),
			dbPassword: os.Getenv("DB_PASSWORD"),
			dbName:     os.Getenv("DB_NAME"),
			dbPort:     os.Getenv("DB_PORT"),
		},
	}
}
