package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils"
	"strconv"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func GetUserById(c *fiber.Ctx) error {
	param := c.Params("id")

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

func GetUserByName(c *fiber.Ctx) error {
	param := c.Params("name")

	user, err := services.NewUserService().GetByUsername(param)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Get by name failure", "errors": err.Error()})
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

func UpdateToken(c *fiber.Ctx) error {
	var payload *dto.FcmToken

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Parse token failure", "errors": err.Error()})
	}

	errors := utils.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	id, e := strconv.ParseUint(c.Params("id"), 10, 64)
	if e != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Parse id failure", "errors": e.Error()})
	}

	err := services.NewUserService().UpdateFcmToken(uint(id), payload.V)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Update token failure", "errors": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Token updated"})
}
