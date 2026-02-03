package car

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type addCarRequest struct {
	Name  string `json:"name"`
	Brand string `json:"eamil"`
	Price int    `json:"price"`
}

func (h *Handler) AddCar(c *fiber.Ctx) error {
	var req addCarRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	ownerIDStr := c.Locals("user_id").(string)
	ownerID, _ := uuid.Parse(ownerIDStr)

	if err := h.service.AddCar(ownerID, req.Name, req.Brand, req.Price); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) ListCars(c *fiber.Ctx) error {
	cars, err := h.service.ListCars()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(cars)
}
