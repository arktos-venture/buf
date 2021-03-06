// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Indices

import (
	context "context"
	v1 "github.com/arktos-venture/buf/proto/strategies/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IndicesClient is the client API for Indices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndicesClient interface {
	// Public API: Get Indice properties
	Get(ctx context.Context, in *IndiceRequest, opts ...grpc.CallOption) (*IndiceReply, error)
	// Private API: Get Stats Indice
	Stats(ctx context.Context, in *IndiceRequest, opts ...grpc.CallOption) (*IndiceStatsReply, error)
	// Public API: Get Strategies Results Indice
	Strategies(ctx context.Context, in *IndiceStrategiesRequest, opts ...grpc.CallOption) (*v1.StrategiesReplies, error)
	// Public API: Search Indices available
	Search(ctx context.Context, in *IndiceSearchRequest, opts ...grpc.CallOption) (*IndiceReplies, error)
}

type indicesClient struct {
	cc grpc.ClientConnInterface
}

func NewIndicesClient(cc grpc.ClientConnInterface) IndicesClient {
	return &indicesClient{cc}
}

func (c *indicesClient) Get(ctx context.Context, in *IndiceRequest, opts ...grpc.CallOption) (*IndiceReply, error) {
	out := new(IndiceReply)
	err := c.cc.Invoke(ctx, "/indices.v1.Indices/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicesClient) Stats(ctx context.Context, in *IndiceRequest, opts ...grpc.CallOption) (*IndiceStatsReply, error) {
	out := new(IndiceStatsReply)
	err := c.cc.Invoke(ctx, "/indices.v1.Indices/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicesClient) Strategies(ctx context.Context, in *IndiceStrategiesRequest, opts ...grpc.CallOption) (*v1.StrategiesReplies, error) {
	out := new(v1.StrategiesReplies)
	err := c.cc.Invoke(ctx, "/indices.v1.Indices/Strategies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicesClient) Search(ctx context.Context, in *IndiceSearchRequest, opts ...grpc.CallOption) (*IndiceReplies, error) {
	out := new(IndiceReplies)
	err := c.cc.Invoke(ctx, "/indices.v1.Indices/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndicesServer is the server API for Indices service.
// All implementations must embed UnimplementedIndicesServer
// for forward compatibility
type IndicesServer interface {
	// Public API: Get Indice properties
	Get(context.Context, *IndiceRequest) (*IndiceReply, error)
	// Private API: Get Stats Indice
	Stats(context.Context, *IndiceRequest) (*IndiceStatsReply, error)
	// Public API: Get Strategies Results Indice
	Strategies(context.Context, *IndiceStrategiesRequest) (*v1.StrategiesReplies, error)
	// Public API: Search Indices available
	Search(context.Context, *IndiceSearchRequest) (*IndiceReplies, error)
	mustEmbedUnimplementedIndicesServer()
}

// UnimplementedIndicesServer must be embedded to have forward compatible implementations.
type UnimplementedIndicesServer struct {
}

func (UnimplementedIndicesServer) Get(context.Context, *IndiceRequest) (*IndiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedIndicesServer) Stats(context.Context, *IndiceRequest) (*IndiceStatsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedIndicesServer) Strategies(context.Context, *IndiceStrategiesRequest) (*v1.StrategiesReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Strategies not implemented")
}
func (UnimplementedIndicesServer) Search(context.Context, *IndiceSearchRequest) (*IndiceReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedIndicesServer) mustEmbedUnimplementedIndicesServer() {}

// UnsafeIndicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndicesServer will
// result in compilation errors.
type UnsafeIndicesServer interface {
	mustEmbedUnimplementedIndicesServer()
}

func RegisterIndicesServer(s grpc.ServiceRegistrar, srv IndicesServer) {
	s.RegisterService(&Indices_ServiceDesc, srv)
}

func _Indices_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indices.v1.Indices/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicesServer).Get(ctx, req.(*IndiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indices_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicesServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indices.v1.Indices/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicesServer).Stats(ctx, req.(*IndiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indices_Strategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicesServer).Strategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indices.v1.Indices/Strategies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicesServer).Strategies(ctx, req.(*IndiceStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indices_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indices.v1.Indices/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicesServer).Search(ctx, req.(*IndiceSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Indices_ServiceDesc is the grpc.ServiceDesc for Indices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Indices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indices.v1.Indices",
	HandlerType: (*IndicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Indices_Get_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Indices_Stats_Handler,
		},
		{
			MethodName: "Strategies",
			Handler:    _Indices_Strategies_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Indices_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indices/v1/indices.proto",
}
