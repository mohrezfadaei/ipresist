package main

import (
	"fmt"
	"log"

	"github.com/mohrezfadaei/ipresist/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.LoadConfig()

	app := fiber.New()

	address := fmt.Sprintf("%s:%s", config.ADDRESS, config.PORT)
	log.Printf("Starting server on %s", address)
	log.Fatal(app.Listen(address))
}
