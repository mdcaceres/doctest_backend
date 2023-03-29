package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
	"strconv"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func GetUser(c *fiber.Ctx) error {
	param := c.Query("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Parse int failure", "errors": err.Error()})
	}

	user, err := services.NewUserService().GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Get by id failure", "errors": err.Error()})
	}

	return c.JSON(user)
}

func GetAll(c *fiber.Ctx) error {
	user, err := services.NewUserService().GetAll()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Get All failure", "errors": err.Error()})
	}
	return c.JSON(user)
}
