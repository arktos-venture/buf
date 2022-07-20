// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Exchanges

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

// ExchangesClient is the client API for Exchanges service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangesClient interface {
	// Public API: Get Exchange properties
	Get(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeReply, error)
	// Private API: Get Stats Exchange
	Stats(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeStatsReply, error)
	// Public API: Search Exchanges available
	Search(ctx context.Context, in *ExchangeSearchRequest, opts ...grpc.CallOption) (*ExchangeReplies, error)
}

type exchangesClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangesClient(cc grpc.ClientConnInterface) ExchangesClient {
	return &exchangesClient{cc}
}

func (c *exchangesClient) Get(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeReply, error) {
	out := new(ExchangeReply)
	err := c.cc.Invoke(ctx, "/exchanges.v1.Exchanges/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangesClient) Stats(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeStatsReply, error) {
	out := new(ExchangeStatsReply)
	err := c.cc.Invoke(ctx, "/exchanges.v1.Exchanges/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangesClient) Search(ctx context.Context, in *ExchangeSearchRequest, opts ...grpc.CallOption) (*ExchangeReplies, error) {
	out := new(ExchangeReplies)
	err := c.cc.Invoke(ctx, "/exchanges.v1.Exchanges/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangesServer is the server API for Exchanges service.
// All implementations must embed UnimplementedExchangesServer
// for forward compatibility
type ExchangesServer interface {
	// Public API: Get Exchange properties
	Get(context.Context, *ExchangeRequest) (*ExchangeReply, error)
	// Private API: Get Stats Exchange
	Stats(context.Context, *ExchangeRequest) (*ExchangeStatsReply, error)
	// Public API: Search Exchanges available
	Search(context.Context, *ExchangeSearchRequest) (*ExchangeReplies, error)
	mustEmbedUnimplementedExchangesServer()
}

// UnimplementedExchangesServer must be embedded to have forward compatible implementations.
type UnimplementedExchangesServer struct {
}

func (UnimplementedExchangesServer) Get(context.Context, *ExchangeRequest) (*ExchangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExchangesServer) Stats(context.Context, *ExchangeRequest) (*ExchangeStatsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedExchangesServer) Search(context.Context, *ExchangeSearchRequest) (*ExchangeReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedExchangesServer) mustEmbedUnimplementedExchangesServer() {}

// UnsafeExchangesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangesServer will
// result in compilation errors.
type UnsafeExchangesServer interface {
	mustEmbedUnimplementedExchangesServer()
}

func RegisterExchangesServer(s grpc.ServiceRegistrar, srv ExchangesServer) {
	s.RegisterService(&Exchanges_ServiceDesc, srv)
}

func _Exchanges_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchanges.v1.Exchanges/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangesServer).Get(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchanges_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangesServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchanges.v1.Exchanges/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangesServer).Stats(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchanges_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchanges.v1.Exchanges/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangesServer).Search(ctx, req.(*ExchangeSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Exchanges_ServiceDesc is the grpc.ServiceDesc for Exchanges service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exchanges_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchanges.v1.Exchanges",
	HandlerType: (*ExchangesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Exchanges_Get_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Exchanges_Stats_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Exchanges_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exchanges/v1/exchanges.proto",
}
