package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/services"
	"io"
	"strconv"
)

func UploadProjectImage(c *fiber.Ctx) error {
	param := c.Params("id")

	projectId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return err
	}

	multiForm, err := c.MultipartForm()
	fmt.Sprintf("%v\n", multiForm)
	src, err := multiForm.File["file"][0].Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	/*src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}*/
	defer src.Close()

	fileBytes, err := io.ReadAll(src)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	err = services.NewProjectService().SaveProjectImage(uint(projectId), fileBytes)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func Serve(c *fiber.Ctx) error {
	return c.SendFile("./uploads/" + c.Params("id"))
}
