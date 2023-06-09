package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/services"
	"io"
	"strconv"
)

func UploadFileToCase(c *fiber.Ctx) error {
	param := c.Params("id")

	caseId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return err
	}

	b, err := GetByteSlice(c)
	if err != nil {
		return err
	}

	err = services.NewCaseService().SaveFiles(uint(caseId), b)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func UploadFileToBug(c *fiber.Ctx) error {
	param := c.Params("id")

	bugId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return err
	}

	b, err := GetByteSlice(c)
	if err != nil {
		return err
	}

	err = services.NewBugService().SaveFiles(uint(bugId), b)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func GetByteSlice(c *fiber.Ctx) ([][]byte, error) {
	multiForm, err := c.MultipartForm()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
	}
	b := make([][]byte, 0)

	files := multiForm.File["files"]
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
		}
		defer src.Close()

		fileBytes, err := io.ReadAll(src)
		if err != nil {
			return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "errors": err.Error()})
		}
		b = append(b, fileBytes)
	}

	return b, nil
}
