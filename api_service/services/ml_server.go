package services

import context "context"

type mlServer struct {
}

func NewMLServiceServer() MLServiceServer {
	return &mlServer{}
}

func (s *mlServer) mustEmbedUnimplementedMLServiceServer() {}
func (s *mlServer) DetectObjects(ctx context.Context, req *ImagesRequest) (*ImageResponse, error) {
	return nil, nil
}
