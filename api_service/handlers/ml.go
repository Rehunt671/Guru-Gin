package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/services"
)

type MLHandler interface {
	TrainModel(*fiber.Ctx) error
	DetectObjects(*fiber.Ctx) error
}

type mlHandler struct {
	mlClient services.MLServiceClient
}

func NewMLHandler(mlClient services.MLServiceClient) MLHandler {
	return &mlHandler{
		mlClient: mlClient,
	}
}

func (h *mlHandler) TrainModel(ctx *fiber.Ctx) error {
	// Implementation for training the model
	return nil
}

func (h *mlHandler) DetectObjects(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to get file from form")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to open file")
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to read file")
	}

	classifications, err := h.DetectObjectWithGRPC(buf.Bytes())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to classify image: %v", err))
	}

	return ctx.JSON(fiber.Map{
		"classifications": classifications,
	})
}

func (h *mlHandler) DetectObjectWithGRPC(image []byte) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &services.ImageRequest{
		Image: image,
	}

	res, err := h.mlClient.DetectObjects(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to classify image: %v", err)
	}

	return res.GetClassifications(), nil
}
