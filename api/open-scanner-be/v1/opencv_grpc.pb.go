// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: open-scanner-be/v1/opencv.proto

package v1

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

const (
	OpenCVService_WarpPerspective_FullMethodName = "/v1.OpenCVService/WarpPerspective"
)

// OpenCVServiceClient is the client API for OpenCVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenCVServiceClient interface {
	WarpPerspective(ctx context.Context, in *WarpPerspectiveRequest, opts ...grpc.CallOption) (*Image, error)
}

type openCVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenCVServiceClient(cc grpc.ClientConnInterface) OpenCVServiceClient {
	return &openCVServiceClient{cc}
}

func (c *openCVServiceClient) WarpPerspective(ctx context.Context, in *WarpPerspectiveRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, OpenCVService_WarpPerspective_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenCVServiceServer is the server API for OpenCVService service.
// All implementations must embed UnimplementedOpenCVServiceServer
// for forward compatibility
type OpenCVServiceServer interface {
	WarpPerspective(context.Context, *WarpPerspectiveRequest) (*Image, error)
	mustEmbedUnimplementedOpenCVServiceServer()
}

// UnimplementedOpenCVServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpenCVServiceServer struct {
}

func (UnimplementedOpenCVServiceServer) WarpPerspective(context.Context, *WarpPerspectiveRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WarpPerspective not implemented")
}
func (UnimplementedOpenCVServiceServer) mustEmbedUnimplementedOpenCVServiceServer() {}

// UnsafeOpenCVServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenCVServiceServer will
// result in compilation errors.
type UnsafeOpenCVServiceServer interface {
	mustEmbedUnimplementedOpenCVServiceServer()
}

func RegisterOpenCVServiceServer(s grpc.ServiceRegistrar, srv OpenCVServiceServer) {
	s.RegisterService(&OpenCVService_ServiceDesc, srv)
}

func _OpenCVService_WarpPerspective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarpPerspectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenCVServiceServer).WarpPerspective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenCVService_WarpPerspective_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenCVServiceServer).WarpPerspective(ctx, req.(*WarpPerspectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenCVService_ServiceDesc is the grpc.ServiceDesc for OpenCVService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenCVService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OpenCVService",
	HandlerType: (*OpenCVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WarpPerspective",
			Handler:    _OpenCVService_WarpPerspective_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "open-scanner-be/v1/opencv.proto",
}
