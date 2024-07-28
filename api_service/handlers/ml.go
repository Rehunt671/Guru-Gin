package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
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

	stream, err := h.mlClient.DetectObjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create stream: %v", err)
	}

	for _, fileHeader := range filesHeader {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			file.Close()
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
		file.Close()

		req := &services.ImageRequest{
			Info: &services.ImageInfo{
				ImageType: filepath.Ext(fileHeader.Filename),
			},
			Data: buf.Bytes(),
		}

		if err := stream.Send(req); err != nil {
			return nil, fmt.Errorf("failed to send request: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive response: %v", err)
	}

	return res.Classifications, nil
}
