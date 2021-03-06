// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Fundamentals

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

// FundamentalsClient is the client API for Fundamentals service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FundamentalsClient interface {
	Get(ctx context.Context, in *FundamentalRequest, opts ...grpc.CallOption) (*FundamentalReply, error)
}

type fundamentalsClient struct {
	cc grpc.ClientConnInterface
}

func NewFundamentalsClient(cc grpc.ClientConnInterface) FundamentalsClient {
	return &fundamentalsClient{cc}
}

func (c *fundamentalsClient) Get(ctx context.Context, in *FundamentalRequest, opts ...grpc.CallOption) (*FundamentalReply, error) {
	out := new(FundamentalReply)
	err := c.cc.Invoke(ctx, "/fundamentals.v1.Fundamentals/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FundamentalsServer is the server API for Fundamentals service.
// All implementations must embed UnimplementedFundamentalsServer
// for forward compatibility
type FundamentalsServer interface {
	Get(context.Context, *FundamentalRequest) (*FundamentalReply, error)
	mustEmbedUnimplementedFundamentalsServer()
}

// UnimplementedFundamentalsServer must be embedded to have forward compatible implementations.
type UnimplementedFundamentalsServer struct {
}

func (UnimplementedFundamentalsServer) Get(context.Context, *FundamentalRequest) (*FundamentalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFundamentalsServer) mustEmbedUnimplementedFundamentalsServer() {}

// UnsafeFundamentalsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FundamentalsServer will
// result in compilation errors.
type UnsafeFundamentalsServer interface {
	mustEmbedUnimplementedFundamentalsServer()
}

func RegisterFundamentalsServer(s grpc.ServiceRegistrar, srv FundamentalsServer) {
	s.RegisterService(&Fundamentals_ServiceDesc, srv)
}

func _Fundamentals_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundamentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundamentalsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fundamentals.v1.Fundamentals/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundamentalsServer).Get(ctx, req.(*FundamentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fundamentals_ServiceDesc is the grpc.ServiceDesc for Fundamentals service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fundamentals_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fundamentals.v1.Fundamentals",
	HandlerType: (*FundamentalsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Fundamentals_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fundamentals/v1/fundamentals.proto",
}
