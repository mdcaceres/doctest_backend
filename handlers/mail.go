package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services/mail"
)

func SendSimple(c *fiber.Ctx) error {
	var payload *dto.MailData

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}
	err := mail.NewEmailService().SendSimple(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
