package handlers

import (
	"github.com/gofiber/fiber/v2"
	usecase "github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"
)

func (h *InitHandler) Login(c *fiber.Ctx) error {
	// Get the request body
	var body usecase.RequestLogin
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Call cmd interface
	res, err := h.Cmd.Login(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *InitHandler) Refresh(c *fiber.Ctx) error {
	// Get the request body
	var body usecase.RequestRefresh
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Call cmd interface
	res, err := h.Cmd.Refresh(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
