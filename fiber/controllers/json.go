package controllers

import "github.com/gofiber/fiber/v2"

func ReturnJSON(c *fiber.Ctx) error {
	data := new(interface{})

	if err := c.BodyParser(data); err != nil {
		return err
	}

	return c.JSON(data)
}
