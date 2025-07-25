package constants

import (
	"github.com/SwipEats/SwipEats/server/internal/config"
	"strconv"
)

var PORT int
var DB_HOST string
var DB_PORT int
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var JWT_SECRET string
var UNSPLASHED_KEY string

func InitEnv() {
	config.LoadEnv()

	portStr := config.GetEnv("PORT", "8080")
	dbPortStr := config.GetEnv("DB_PORT", "3306")
	DB_HOST = config.GetEnv("DB_HOST", "localhost")
	DB_USER = config.GetEnv("DB_USER", "root")
	DB_PASSWORD = config.GetEnv("DB_PASSWORD", "")
	DB_NAME = config.GetEnv("DB_NAME", "swipeats_db")
	UNSPLASHED_KEY = config.GetEnv("UNSPLASHED_KEY", "your_unsplash_key_here")

	JWT_SECRET = config.GetEnv("JWT_SECRET", "defaultsecretkey")
	
	port, err := strconv.Atoi(portStr)

	if err != nil {
		panic("Invalid PORT value in .env file: " + portStr)
	}

	dbPort, err := strconv.Atoi(dbPortStr)

	if err != nil {
		panic("Invalid DB_PORT value in .env file: " + dbPortStr)
	}

	PORT = port
	DB_PORT = dbPort
}