package handlers

import (
	"strconv"

	"github.com/mohrezfadaei/ipresist/resource/apiv1/controllers"
	"github.com/mohrezfadaei/ipresist/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mohrezfadaei/ipresist/internal/db"
)

type IPHandler struct {
	Controller controllers.IPController
}

// GetIPAddrs handles GET requests to fetch all IP addresses.
//
//	GET /api/v1/ips
func (h IPHandler) GetIPAddrs(c *fiber.Ctx) error {
	status := c.Query("status")
	sort := c.Query("sort")
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	ips, err := h.Controller.GetAll(status, sort, offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 101, nil))
	}
	return c.JSON(utils.Jsonify(ips, nil, 200, 100, nil))
}

// GetIPByID handles GET requests to fetch a specific IP address by its UUID.
//
//	GET /api/v1/ip/<id: uuid>
func (h IPHandler) GetIPByID(c *fiber.Ctx) error {
	ipID := c.Params("id")

	id, err := uuid.Parse(ipID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 102, nil))
	}

	ip, err := h.Controller.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Jsonify(nil, nil, 404, 103, nil))
	}
	return c.JSON(utils.Jsonify(ip, nil, 200, 100, nil))
}

// CreateIP handles POST requests to create a new IP address.
//
//	POST /api/v1/ip
func (h IPHandler) CreateIP(c *fiber.Ctx) error {
	var input struct {
		IPAddress string `json:"ipaddress"`
		Note      string `json:"note"`
		Status    string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 104, nil))
	}

	if _, ok := controllers.ValidStatuses[input.Status]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 108, nil))
	}

	ip := db.IP{
		IPAddress: input.IPAddress,
		Note:      input.Note,
		Status:    db.IPStatus(input.Status),
	}

	if err := h.Controller.Create(&ip); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 105, nil))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Jsonify(ip, nil, 201, 100, nil))
}

// UpdateIP handles PUT requests to update an existing IP address by its UUID.
//
//	PUT /api/v1/ip/<id: uuid>
func (h IPHandler) UpdateIP(c *fiber.Ctx) error {
	ipID := c.Params("id")

	id, err := uuid.Parse(ipID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 102, nil))
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 104, nil))
	}

	if status, ok := data["status"].(string); ok {
		if _, valid := controllers.ValidStatuses[status]; !valid {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 108, nil))
		}
	}

	ip, err := h.Controller.Update(id, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 106, nil))
	}
	return c.JSON(utils.Jsonify(ip, nil, 200, 100, nil))
}

// DeleteIP handles DELETE requests to remove an IP address by its UUID.
//
//	DELETE /api/v1/ip/<id: uuid>
func (h IPHandler) DeleteIP(c *fiber.Ctx) error {
	ipID := c.Params("id")

	id, err := uuid.Parse(ipID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Jsonify(nil, nil, 400, 102, nil))
	}

	if err := h.Controller.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Jsonify(nil, nil, 500, 107, nil))
	}
	return c.Status(fiber.StatusAccepted).JSON(utils.Jsonify(nil, nil, 200, 100, nil))
}
