package services

import context "context"

type mlServer struct {
}

func NewMLServiceServer() MLServiceServer {
	return &mlServer{}
}

func (s *mlServer) mustEmbedUnimplementedMLServiceServer() {}
func (s *mlServer) ClassifyImages(ctx context.Context, req *ImageRequest) (*ImageResponse, error) {
	return nil, nil
}
