package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
)

func CreateBug(c *fiber.Ctx) error {
	var payload *dto.BugRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	bug, err := services.NewBugService().Create(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bug": bug}})
}

func GetAllBugsByProjectId(c *fiber.Ctx) error {
	param := c.Params("id")
	bugs, err := services.NewBugService().GetAllByProjectId(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bugs": bugs}})
}

func GetBugByUserId(c *fiber.Ctx) error {
	param := c.Params("id")
	bugs, err := services.NewBugService().GetAllByUserId(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bugs": bugs}})
}
