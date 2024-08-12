package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohrezfadaei/ipresist/config"
)

var STATUS_MESSAGES = map[int]string{
	100: "OK",
	101: "Resource is not implemented",
	102: "Invalid UUID format",
	103: "Resource not found",
	104: "Bad request",
	105: "Failed to create resource",
	106: "Failed to update resource",
	107: "Failed to delete resource",
	108: "Invalid data provided",
}

func Jsonify(state interface{}, metadata interface{}, status int, code int, headers map[string]string) fiber.Map {
	resource := fiber.Map{
		"result": state,
		"status": fiber.Map{
			"code": code,
		},
		"_metadata": metadata,
	}
	if config.DEBUG {
		resource["status"].(fiber.Map)["message"] = STATUS_MESSAGES[code]
	}
	return resource
}
