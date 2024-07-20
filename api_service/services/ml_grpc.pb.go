// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: ml.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MLServiceClient is the client API for MLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MLServiceClient interface {
	// Bidirectional streaming RPC method for image classification.
	DetectObjects(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
}

type mLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMLServiceClient(cc grpc.ClientConnInterface) MLServiceClient {
	return &mLServiceClient{cc}
}

func (c *mLServiceClient) DetectObjects(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ml.MLService/DetectObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MLServiceServer is the server API for MLService service.
// All implementations must embed UnimplementedMLServiceServer
// for forward compatibility
type MLServiceServer interface {
	// Bidirectional streaming RPC method for image classification.
	DetectObjects(context.Context, *ImageRequest) (*ImageResponse, error)
	mustEmbedUnimplementedMLServiceServer()
}

// UnimplementedMLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMLServiceServer struct {
}

func (UnimplementedMLServiceServer) DetectObjects(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectObjects not implemented")
}
func (UnimplementedMLServiceServer) mustEmbedUnimplementedMLServiceServer() {}

// UnsafeMLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MLServiceServer will
// result in compilation errors.
type UnsafeMLServiceServer interface {
	mustEmbedUnimplementedMLServiceServer()
}

func RegisterMLServiceServer(s grpc.ServiceRegistrar, srv MLServiceServer) {
	s.RegisterService(&MLService_ServiceDesc, srv)
}

func _MLService_DetectObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MLServiceServer).DetectObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ml.MLService/DetectObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MLServiceServer).DetectObjects(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MLService_ServiceDesc is the grpc.ServiceDesc for MLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ml.MLService",
	HandlerType: (*MLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectObjects",
			Handler:    _MLService_DetectObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ml.proto",
}
