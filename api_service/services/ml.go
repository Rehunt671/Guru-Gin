package services

import (
	"bytes"
	context "context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"gitlab.com/gurugin/repositories"
)

type MLService interface {
	DetectObjectsWithGRPC(filesHeader []*multipart.FileHeader) ([]string, error)
}

type mlService struct {
	mlClient         MLServiceClient
	recipeRepository repositories.RecipeRepository
}

func NewMLService(mlClient MLServiceClient, recipeRepository repositories.RecipeRepository) MLService {
	return &mlService{
		mlClient:         mlClient,
		recipeRepository: recipeRepository,
	}
}

func (s *mlService) DetectObjectsWithGRPC(filesHeader []*multipart.FileHeader) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Create the gRPC stream
	stream, err := s.mlClient.DetectObjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create stream: %v", err)
	}

	// Send all image files
	for _, fileHeader := range filesHeader {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		defer file.Close() // Ensure file is closed after processing

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}

		req := &ImageRequest{
			Info: &ImageInfo{
				ImageType: filepath.Ext(fileHeader.Filename),
			},
			Data: buf.Bytes(),
		}

		if err := stream.Send(req); err != nil {
			return nil, fmt.Errorf("failed to send request: %v", err)
		}
	}

	if err := stream.CloseSend(); err != nil {
		return nil, fmt.Errorf("failed to close stream: %v", err)
	}

	var classifications []string
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// End of stream, exit loop
			break
		} else if err != nil {
			return nil, fmt.Errorf("error receiving stream: %v", err)
		}
		classifications = append(classifications, resp.Classification)
	}

	return classifications, nil
}
