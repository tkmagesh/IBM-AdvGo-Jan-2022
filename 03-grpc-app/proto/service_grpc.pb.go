// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	// Request & Response
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Server Streaming
	GeneratePrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (AppService_GeneratePrimesClient, error)
	// Client Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AppService_ComputeAverageClient, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.AppService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GeneratePrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (AppService_GeneratePrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppService_ServiceDesc.Streams[0], "/proto.AppService/GeneratePrimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &appServiceGeneratePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppService_GeneratePrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type appServiceGeneratePrimesClient struct {
	grpc.ClientStream
}

func (x *appServiceGeneratePrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AppService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppService_ServiceDesc.Streams[1], "/proto.AppService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &appServiceComputeAverageClient{stream}
	return x, nil
}

type AppService_ComputeAverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type appServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *appServiceComputeAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appServiceComputeAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	// Request & Response
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Server Streaming
	GeneratePrimes(*PrimeRequest, AppService_GeneratePrimesServer) error
	// Client Streaming
	ComputeAverage(AppService_ComputeAverageServer) error
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAppServiceServer) GeneratePrimes(*PrimeRequest, AppService_GeneratePrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GeneratePrimes not implemented")
}
func (UnimplementedAppServiceServer) ComputeAverage(AppService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AppService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GeneratePrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServiceServer).GeneratePrimes(m, &appServiceGeneratePrimesServer{stream})
}

type AppService_GeneratePrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type appServiceGeneratePrimesServer struct {
	grpc.ServerStream
}

func (x *appServiceGeneratePrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AppService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServiceServer).ComputeAverage(&appServiceComputeAverageServer{stream})
}

type AppService_ComputeAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type appServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *appServiceComputeAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appServiceComputeAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AppService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GeneratePrimes",
			Handler:       _AppService_GeneratePrimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _AppService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
