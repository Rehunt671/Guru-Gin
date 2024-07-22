package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
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
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to get form data")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("No files uploaded")
	}

	classifications, err := h.DetectObjectsWithGRPC(files)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to classify images: %v", err))
	}

	return ctx.JSON(fiber.Map{
		"classifications": classifications,
	})
}

func (h *mlHandler) DetectObjectsWithGRPC(filesHeader []*multipart.FileHeader) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var images []*services.ImageRequest
	for _, fileHeader := range filesHeader {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}

		req := &services.ImageRequest{
			Image: buf.Bytes(),
		}

		images = append(images, req)
	}

	req := &services.ImagesRequest{
		Images: images,
	}

	res, err := h.mlClient.DetectObjects(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to classify images: %v", err)
	}

	return res.Classifications, nil
}
