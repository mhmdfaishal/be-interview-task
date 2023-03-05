package delivery

import (
	"be-interview-task/internal/contracts"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type YaraHandler struct {
	yaraUsecase contracts.YaraUsecase
}

func NewYaraHandler(f *fiber.App, vu contracts.YaraUsecase) {
	handler := &YaraHandler{
		yaraUsecase: vu,
	}

	f.Post("/analyze/file", handler.AnalyzeFile)
	f.Get("/get/result/:id", handler.GetResult)
}

func (v *YaraHandler) AnalyzeFile(c *fiber.Ctx) error {
	var ctx = c.Context()

	file, err := c.FormFile("file")
	if err != nil {
		if err.Error() == "request has no multipart/form-data Content-Type" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "No file found in request",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error uploading file",
			"error":   err.Error(),
		})
	}
	xApiKey := c.Get("x-apikey")
	if xApiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "API Key is required",
		})
	}

	result, err := v.yaraUsecase.AnalyzeFile(ctx, xApiKey, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error uploading file",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "File uploaded successfully",
		"data":    result,
	})
}

func (v *YaraHandler) GetResult(c *fiber.Ctx) error {
	var ctx = c.Context()

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID is required",
		})
	}
	//convert percent on path to true string
	id, err := url.QueryUnescape(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting result",
			"error":   err.Error(),
		})
	}
	xApiKey := c.Get("x-apikey")
	if xApiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "API Key is required",
		})
	}

	result, err := v.yaraUsecase.GetResult(ctx, xApiKey, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting result",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Result fetched successfully",
		"data":    result,
	})
}
