package services

type mlServer struct {
}

func NewMLServiceServer() MLServiceServer {
	return &mlServer{}
}

func (s *mlServer) mustEmbedUnimplementedMLServiceServer() {}
func (s *mlServer) DetectObjects(stream MLService_DetectObjectsServer) error {
	return nil
}
