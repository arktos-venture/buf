// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package strategies_v1

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

// StrategiesClient is the client API for Strategies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategiesClient interface {
	Execute(ctx context.Context, in *StrategyRequest, opts ...grpc.CallOption) (*StrategyReply, error)
}

type strategiesClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategiesClient(cc grpc.ClientConnInterface) StrategiesClient {
	return &strategiesClient{cc}
}

func (c *strategiesClient) Execute(ctx context.Context, in *StrategyRequest, opts ...grpc.CallOption) (*StrategyReply, error) {
	out := new(StrategyReply)
	err := c.cc.Invoke(ctx, "/strategies.v1.Strategies/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategiesServer is the server API for Strategies service.
// All implementations must embed UnimplementedStrategiesServer
// for forward compatibility
type StrategiesServer interface {
	Execute(context.Context, *StrategyRequest) (*StrategyReply, error)
	mustEmbedUnimplementedStrategiesServer()
}

// UnimplementedStrategiesServer must be embedded to have forward compatible implementations.
type UnimplementedStrategiesServer struct {
}

func (UnimplementedStrategiesServer) Execute(context.Context, *StrategyRequest) (*StrategyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedStrategiesServer) mustEmbedUnimplementedStrategiesServer() {}

// UnsafeStrategiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategiesServer will
// result in compilation errors.
type UnsafeStrategiesServer interface {
	mustEmbedUnimplementedStrategiesServer()
}

func RegisterStrategiesServer(s grpc.ServiceRegistrar, srv StrategiesServer) {
	s.RegisterService(&Strategies_ServiceDesc, srv)
}

func _Strategies_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategiesServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strategies.v1.Strategies/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategiesServer).Execute(ctx, req.(*StrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Strategies_ServiceDesc is the grpc.ServiceDesc for Strategies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Strategies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strategies.v1.Strategies",
	HandlerType: (*StrategiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Strategies_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/strategies/v1/strategies.proto",
}
