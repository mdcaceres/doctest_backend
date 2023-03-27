package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
