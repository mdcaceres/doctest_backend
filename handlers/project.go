package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils"
	"strconv"
)

func CreateProject(c *fiber.Ctx) error {
	var payload *dto.ProjectRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	projectResponse, serviceError := services.NewProjectService().Create(payload)

	if serviceError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": serviceError.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": projectResponse}})
}

func GetProjects(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	ps, e := services.NewProjectService().GetAll(fmt.Sprintf("%v", userId))
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": e.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"projects": ps}})
}

func GetProject(c *fiber.Ctx) error {
	param := c.Params("id")
	projectId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return err
	}

	p, e := services.NewProjectService().Get(uint(projectId))
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": e.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": p}})
}

func JoinProject(c *fiber.Ctx) error {
	/*var payload *dto.JoinProject
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}*/

	//userId := c.Query("userId")
	//projectId := c.Query("projectId")

	payload := &dto.JoinProject{
		ProjectId: c.Query("projectId"),
		UserId:    c.Query("userId"),
	}

	//payload.ProjectId = projectId
	//payload.UserId = userId

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": errors})
	}
	r, e := services.NewProjectService().Join(c, payload)
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": e.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": r}})
}
