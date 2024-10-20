package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *InitHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	res, err := h.Cmd.GetProfile(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
