package handlers

import (
	"github.com/mohrezfadaei/ipresist/utils"

	"github.com/gofiber/fiber/v2"
)

// GetIPAddrs handles GET requests to fetch all IP addresses.
//
//	GET /api/v1/ips
func GetIPAddrs(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
}

// GetIPByID handles GET requests to fetch a specific IP address by its UUID.
//
//	GET /api/v1/ip/<id: uuid>
func GetIPByID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
}

// CreateIP handles POST requests to create a new IP address.
//
//	POST /api/v1/ip
func CreateIP(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
}

// UpdateIP handles PUT requests to update an existing IP address by its UUID.
//
//	PUT /api/v1/ip/<id: uuid>
func UpdateIP(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
}

// DeleteIP handles DELETE requests to remove an IP address by its UUID.
//
//	DELETE /api/v1/ip/<id: uuid>
func DeleteIP(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
}
