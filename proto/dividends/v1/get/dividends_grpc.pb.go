// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DividendsClient is the client API for Dividends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DividendsClient interface {
	Search(ctx context.Context, in *DividendsRequest, opts ...grpc.CallOption) (*DividendsReply, error)
	Analytics(ctx context.Context, in *DividendsRequest, opts ...grpc.CallOption) (*DividendsReply, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dividendsClient struct {
	cc grpc.ClientConnInterface
}

func NewDividendsClient(cc grpc.ClientConnInterface) DividendsClient {
	return &dividendsClient{cc}
}

func (c *dividendsClient) Search(ctx context.Context, in *DividendsRequest, opts ...grpc.CallOption) (*DividendsReply, error) {
	out := new(DividendsReply)
	err := c.cc.Invoke(ctx, "/dividends.v1.Dividends/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dividendsClient) Analytics(ctx context.Context, in *DividendsRequest, opts ...grpc.CallOption) (*DividendsReply, error) {
	out := new(DividendsReply)
	err := c.cc.Invoke(ctx, "/dividends.v1.Dividends/Analytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dividendsClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dividends.v1.Dividends/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DividendsServer is the server API for Dividends service.
// All implementations must embed UnimplementedDividendsServer
// for forward compatibility
type DividendsServer interface {
	Search(context.Context, *DividendsRequest) (*DividendsReply, error)
	Analytics(context.Context, *DividendsRequest) (*DividendsReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedDividendsServer()
}

// UnimplementedDividendsServer must be embedded to have forward compatible implementations.
type UnimplementedDividendsServer struct {
}

func (UnimplementedDividendsServer) Search(context.Context, *DividendsRequest) (*DividendsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDividendsServer) Analytics(context.Context, *DividendsRequest) (*DividendsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analytics not implemented")
}
func (UnimplementedDividendsServer) Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedDividendsServer) mustEmbedUnimplementedDividendsServer() {}

// UnsafeDividendsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DividendsServer will
// result in compilation errors.
type UnsafeDividendsServer interface {
	mustEmbedUnimplementedDividendsServer()
}

func RegisterDividendsServer(s grpc.ServiceRegistrar, srv DividendsServer) {
	s.RegisterService(&Dividends_ServiceDesc, srv)
}

func _Dividends_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DividendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DividendsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dividends.v1.Dividends/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DividendsServer).Search(ctx, req.(*DividendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dividends_Analytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DividendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DividendsServer).Analytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dividends.v1.Dividends/Analytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DividendsServer).Analytics(ctx, req.(*DividendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dividends_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DividendsServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dividends.v1.Dividends/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DividendsServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Dividends_ServiceDesc is the grpc.ServiceDesc for Dividends service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dividends_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dividends.v1.Dividends",
	HandlerType: (*DividendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Dividends_Search_Handler,
		},
		{
			MethodName: "Analytics",
			Handler:    _Dividends_Analytics_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Dividends_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dividends/v1/get/dividends.proto",
}
