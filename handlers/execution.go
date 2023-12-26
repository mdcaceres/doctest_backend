package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils"
	"time"
)

func ExecuteTest(c *fiber.Ctx) error {

	var payload *dto.TestExecutionRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	executedCase, err := services.NewExecutionService().SaveTestExecution(payload)
	if err != nil {
		return err
	}

	t := dto.GetTestExecutionResponse(executedCase)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"test_execution": t}})
}

func GetAverage(c *fiber.Ctx) error {

	projectId := c.Query("project_id")
	status := c.Query("status")
	start, startError := time.Parse(time.RFC3339, c.Query("start"))
	end, endError := time.Parse(time.RFC3339, c.Query("end"))

	if startError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": startError})
	}
	if endError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": endError})
	}

	p, err := services.NewExecutionService().GetPercentage(projectId, status, start, end)
	if err != nil {
		return err
	}

	r := dto.TestExecutionAverageResponse{
		Percentage: p,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"test_execution": r}})
}
