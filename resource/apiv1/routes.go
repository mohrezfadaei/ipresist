package apiv1

import (
	"github.com/mohrezfadaei/ipresist/resource/apiv1/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/ips", handlers.GetIPAddrs)
	api.Get("/ip/:id", handlers.GetIPByID)
	api.Post("/ip", handlers.CreateIP)
	api.Put("/ip/:id", handlers.UpdateIP)
	api.Delete("/ip/:id", handlers.DeleteIP)
}
