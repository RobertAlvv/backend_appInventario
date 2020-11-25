package handlers

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to the API AppInventariado!", "data": nil, "status": "success"})
}
