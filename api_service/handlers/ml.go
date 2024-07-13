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
	ClassifyImages(*fiber.Ctx) error
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

func (h *mlHandler) ClassifyImages(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to get multipart form")
	}

	files := form.File["files"]
	var images [][]byte

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Failed to open file")
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Failed to read file")
		}
		images = append(images, buf.Bytes())
	}

	classifications, err := h.classifyImagesWithGRPC(images)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to classify images: %v", err))
	}

	return ctx.JSON(fiber.Map{
		"classifications": classifications,
	})
}

func (h *mlHandler) classifyImagesWithGRPC(images [][]byte) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &services.ImageRequest{
		Images: images,
	}

	res, err := h.mlClient.ClassifyImages(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to classify images: %v", err)
	}

	return res.GetClassifications(), nil
}
