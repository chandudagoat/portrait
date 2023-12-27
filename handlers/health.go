package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHealth(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
