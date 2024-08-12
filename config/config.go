package config

import (
	"os"
	"strconv"
)

var (
	DEBUG   bool
	ADDRESS string
	PORT    string
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
}
