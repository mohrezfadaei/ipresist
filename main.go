package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mohrezfadaei/ipresist/internal/db"
	"github.com/mohrezfadaei/ipresist/internal/scheduler"
	"github.com/mohrezfadaei/ipresist/resource/apiv1"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mohrezfadaei/ipresist/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.LoadConfig()
	db.ConnectDB()
	defer db.DB.Close()
	db.ConnectInfluxDB()
	defer func() {
		if err := db.InfluxClient.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	scheduler.StartHealthCheckScheduler(time.Duration(config.HEALTH_CHECK_INTERVAL) * time.Second)

	db.RunMigrations()

	app := fiber.New()

	apiv1.SetupRoutes(app)

	address := fmt.Sprintf("%s:%s", config.ADDRESS, config.PORT)
	log.Printf("Starting server on %s", address)
	log.Fatal(app.Listen(address))
}
