package apiv1

import (
	"github.com/mohrezfadaei/ipresist/resource/apiv1/controllers"
	"github.com/mohrezfadaei/ipresist/resource/apiv1/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	ipHandler := handlers.IPHandler{
		Controller: controllers.IPController{},
	}

	api := app.Group("/api/v1")
	api.Get("/ips", ipHandler.GetIPAddrs)
	api.Get("/ip/:id", ipHandler.GetIPByID)
	api.Post("/ip", ipHandler.CreateIP)
	api.Put("/ip/:id", ipHandler.UpdateIP)
	api.Delete("/ip/:id", ipHandler.DeleteIP)
}
