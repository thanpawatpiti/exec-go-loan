package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *InitHandler) Login(c *fiber.Ctx) error {
	return c.SendString("login")
}
