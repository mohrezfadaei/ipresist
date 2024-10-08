package config

import (
	"os"
	"strconv"
)

var (
	DEBUG                 bool
	ADDRESS               string
	PORT                  string
	DB_HOST               string
	DB_PORT               string
	DB_USER               string
	DB_NAME               string
	DB_PASSWORD           string
	TIMEZONE              string
	INFLUXDB_HOST         string
	INFLUXDB_TOKEN        string
	INFLUXDB_ORG          string
	INFLUXDB_BUCKET       string
	HEALTH_CHECK_INTERVAL int
)

func LoadConfig() {
	if debug, exists := os.LookupEnv("DEBUG"); exists {
		DEBUG, _ = strconv.ParseBool(debug)
	} else {
		DEBUG = false
	}

	if address, exists := os.LookupEnv("ADDRESS"); exists {
		ADDRESS = address
	} else {
		ADDRESS = "127.0.0.1" // default address
	}

	if port, exists := os.LookupEnv("PORT"); exists {
		PORT = port
	} else {
		PORT = "3000" // default port
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	TIMEZONE = os.Getenv("TIMEZONE")

	INFLUXDB_HOST = os.Getenv("INFLUXDB_HOST")
	INFLUXDB_TOKEN = os.Getenv("INFLUXDB_TOKEN")
	INFLUXDB_ORG = os.Getenv("INFLUXDB_ORG")
	INFLUXDB_BUCKET = os.Getenv("INFLUXDB_BUCKET")

	interval, err := strconv.Atoi(os.Getenv("HEALTH_CHECK_INTERVAL"))
	if err != nil || interval <= 0 {
		HEALTH_CHECK_INTERVAL = 300 // Default to 300 seconds
	} else {
		HEALTH_CHECK_INTERVAL = interval
	}
}
