package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils"
)

// post
func CreateInvitation(c *fiber.Ctx) error {
	var payload *dto.InvitationRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err})
	}

	err := utils.ValidateStruct(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err})
	}

	resp, errors := services.NewInvitationService().Create(c, payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": errors})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"invitation": resp}})
}

//get
