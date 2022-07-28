// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc2
// source: gofound.proto

package gofoundpd

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

// GofoundServiceClient is the client API for GofoundService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GofoundServiceClient interface {
	Welcome(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*WelcomeResponse, error)
	GC(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Status(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type gofoundServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGofoundServiceClient(cc grpc.ClientConnInterface) GofoundServiceClient {
	return &gofoundServiceClient{cc}
}

func (c *gofoundServiceClient) Welcome(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*WelcomeResponse, error) {
	out := new(WelcomeResponse)
	err := c.cc.Invoke(ctx, "/gofound.v1.GofoundService/Welcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gofoundServiceClient) GC(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/gofound.v1.GofoundService/GC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gofoundServiceClient) Status(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/gofound.v1.GofoundService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gofoundServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/gofound.v1.GofoundService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GofoundServiceServer is the server API for GofoundService service.
// All implementations must embed UnimplementedGofoundServiceServer
// for forward compatibility
type GofoundServiceServer interface {
	Welcome(context.Context, *EmptyRequest) (*WelcomeResponse, error)
	GC(context.Context, *EmptyRequest) (*EmptyResponse, error)
	Status(context.Context, *EmptyRequest) (*StatusResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedGofoundServiceServer()
}

// UnimplementedGofoundServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGofoundServiceServer struct {
}

func (UnimplementedGofoundServiceServer) Welcome(context.Context, *EmptyRequest) (*WelcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}
func (UnimplementedGofoundServiceServer) GC(context.Context, *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GC not implemented")
}
func (UnimplementedGofoundServiceServer) Status(context.Context, *EmptyRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedGofoundServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedGofoundServiceServer) mustEmbedUnimplementedGofoundServiceServer() {}

// UnsafeGofoundServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GofoundServiceServer will
// result in compilation errors.
type UnsafeGofoundServiceServer interface {
	mustEmbedUnimplementedGofoundServiceServer()
}

func RegisterGofoundServiceServer(s grpc.ServiceRegistrar, srv GofoundServiceServer) {
	s.RegisterService(&GofoundService_ServiceDesc, srv)
}

func _GofoundService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GofoundServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofound.v1.GofoundService/Welcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GofoundServiceServer).Welcome(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GofoundService_GC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GofoundServiceServer).GC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofound.v1.GofoundService/GC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GofoundServiceServer).GC(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GofoundService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GofoundServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofound.v1.GofoundService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GofoundServiceServer).Status(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GofoundService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GofoundServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofound.v1.GofoundService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GofoundServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GofoundService_ServiceDesc is the grpc.ServiceDesc for GofoundService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GofoundService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gofound.v1.GofoundService",
	HandlerType: (*GofoundServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _GofoundService_Welcome_Handler,
		},
		{
			MethodName: "GC",
			Handler:    _GofoundService_GC_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _GofoundService_Status_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _GofoundService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gofound.proto",
}
