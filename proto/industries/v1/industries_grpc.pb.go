// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Industries

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

// IndustriesClient is the client API for Industries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndustriesClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IndustryReply, error)
	Activities(ctx context.Context, in *IndustryActivitiesRequest, opts ...grpc.CallOption) (*IndustrySearchReply, error)
	Search(ctx context.Context, in *IndustrySearchRequest, opts ...grpc.CallOption) (*IndustrySearchReply, error)
}

type industriesClient struct {
	cc grpc.ClientConnInterface
}

func NewIndustriesClient(cc grpc.ClientConnInterface) IndustriesClient {
	return &industriesClient{cc}
}

func (c *industriesClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IndustryReply, error) {
	out := new(IndustryReply)
	err := c.cc.Invoke(ctx, "/industries.v1.Industries/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *industriesClient) Activities(ctx context.Context, in *IndustryActivitiesRequest, opts ...grpc.CallOption) (*IndustrySearchReply, error) {
	out := new(IndustrySearchReply)
	err := c.cc.Invoke(ctx, "/industries.v1.Industries/Activities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *industriesClient) Search(ctx context.Context, in *IndustrySearchRequest, opts ...grpc.CallOption) (*IndustrySearchReply, error) {
	out := new(IndustrySearchReply)
	err := c.cc.Invoke(ctx, "/industries.v1.Industries/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndustriesServer is the server API for Industries service.
// All implementations must embed UnimplementedIndustriesServer
// for forward compatibility
type IndustriesServer interface {
	List(context.Context, *emptypb.Empty) (*IndustryReply, error)
	Activities(context.Context, *IndustryActivitiesRequest) (*IndustrySearchReply, error)
	Search(context.Context, *IndustrySearchRequest) (*IndustrySearchReply, error)
	mustEmbedUnimplementedIndustriesServer()
}

// UnimplementedIndustriesServer must be embedded to have forward compatible implementations.
type UnimplementedIndustriesServer struct {
}

func (UnimplementedIndustriesServer) List(context.Context, *emptypb.Empty) (*IndustryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIndustriesServer) Activities(context.Context, *IndustryActivitiesRequest) (*IndustrySearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activities not implemented")
}
func (UnimplementedIndustriesServer) Search(context.Context, *IndustrySearchRequest) (*IndustrySearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedIndustriesServer) mustEmbedUnimplementedIndustriesServer() {}

// UnsafeIndustriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndustriesServer will
// result in compilation errors.
type UnsafeIndustriesServer interface {
	mustEmbedUnimplementedIndustriesServer()
}

func RegisterIndustriesServer(s grpc.ServiceRegistrar, srv IndustriesServer) {
	s.RegisterService(&Industries_ServiceDesc, srv)
}

func _Industries_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndustriesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/industries.v1.Industries/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndustriesServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Industries_Activities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndustryActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndustriesServer).Activities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/industries.v1.Industries/Activities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndustriesServer).Activities(ctx, req.(*IndustryActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Industries_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndustrySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndustriesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/industries.v1.Industries/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndustriesServer).Search(ctx, req.(*IndustrySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Industries_ServiceDesc is the grpc.ServiceDesc for Industries service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Industries_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "industries.v1.Industries",
	HandlerType: (*IndustriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Industries_List_Handler,
		},
		{
			MethodName: "Activities",
			Handler:    _Industries_Activities_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Industries_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/industries/v1/industries.proto",
}
