package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/services"
)

type MLHandler interface {
	TrainModel(*fiber.Ctx) error
	DetectObjects(*fiber.Ctx) error
}

type mlHandler struct {
	mlService services.MLService
}

func NewMLHandler(mlService services.MLService) MLHandler {
	return &mlHandler{
		mlService: mlService,
	}
}

func (h *mlHandler) TrainModel(ctx *fiber.Ctx) error {
	// Implementation for training the model
	return nil
}

func (h *mlHandler) DetectObjects(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to get form data")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("No files uploaded")
	}

	classifications, err := h.mlService.DetectObjectsWithGRPC(files)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to classify images: %v", err))
	}

	return ctx.JSON(fiber.Map{
		"classifications": classifications,
	})
}
