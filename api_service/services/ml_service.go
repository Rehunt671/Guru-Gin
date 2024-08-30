package services

import (
	"bytes"
	context "context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"
)

type MLService interface {
	DetectObjectsWithGRPC(filesHeader []*multipart.FileHeader) ([]string, error)
}

type mlService struct {
	mlClient MLServiceClient
}

func NewMLService(mlClient MLServiceClient) MLService {
	return &mlService{
		mlClient: mlClient,
	}
}

func (s *mlService) DetectObjectsWithGRPC(filesHeader []*multipart.FileHeader) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := s.mlClient.DetectObjects(ctx)
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

	res, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive response: %v", err)
	}

	return res.Classifications, nil
}
