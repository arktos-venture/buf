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

// IndexesClient is the client API for Indexes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexesClient interface {
	Get(ctx context.Context, in *IndexesRequest, opts ...grpc.CallOption) (*IndexesReply, error)
	Search(ctx context.Context, in *IndexesSearchRequest, opts ...grpc.CallOption) (*IndexesSearchReplies, error)
	Create(ctx context.Context, in *IndexesCreateRequest, opts ...grpc.CallOption) (*IndexesReply, error)
	Update(ctx context.Context, in *IndexesCreateRequest, opts ...grpc.CallOption) (*IndexesReply, error)
	Delete(ctx context.Context, in *IndexesRequest, opts ...grpc.CallOption) (*IndexesReply, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type indexesClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexesClient(cc grpc.ClientConnInterface) IndexesClient {
	return &indexesClient{cc}
}

func (c *indexesClient) Get(ctx context.Context, in *IndexesRequest, opts ...grpc.CallOption) (*IndexesReply, error) {
	out := new(IndexesReply)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexesClient) Search(ctx context.Context, in *IndexesSearchRequest, opts ...grpc.CallOption) (*IndexesSearchReplies, error) {
	out := new(IndexesSearchReplies)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexesClient) Create(ctx context.Context, in *IndexesCreateRequest, opts ...grpc.CallOption) (*IndexesReply, error) {
	out := new(IndexesReply)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexesClient) Update(ctx context.Context, in *IndexesCreateRequest, opts ...grpc.CallOption) (*IndexesReply, error) {
	out := new(IndexesReply)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexesClient) Delete(ctx context.Context, in *IndexesRequest, opts ...grpc.CallOption) (*IndexesReply, error) {
	out := new(IndexesReply)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexesClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/indexes.v1.Indexes/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexesServer is the server API for Indexes service.
// All implementations must embed UnimplementedIndexesServer
// for forward compatibility
type IndexesServer interface {
	Get(context.Context, *IndexesRequest) (*IndexesReply, error)
	Search(context.Context, *IndexesSearchRequest) (*IndexesSearchReplies, error)
	Create(context.Context, *IndexesCreateRequest) (*IndexesReply, error)
	Update(context.Context, *IndexesCreateRequest) (*IndexesReply, error)
	Delete(context.Context, *IndexesRequest) (*IndexesReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedIndexesServer()
}

// UnimplementedIndexesServer must be embedded to have forward compatible implementations.
type UnimplementedIndexesServer struct {
}

func (UnimplementedIndexesServer) Get(context.Context, *IndexesRequest) (*IndexesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedIndexesServer) Search(context.Context, *IndexesSearchRequest) (*IndexesSearchReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedIndexesServer) Create(context.Context, *IndexesCreateRequest) (*IndexesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIndexesServer) Update(context.Context, *IndexesCreateRequest) (*IndexesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIndexesServer) Delete(context.Context, *IndexesRequest) (*IndexesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIndexesServer) Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedIndexesServer) mustEmbedUnimplementedIndexesServer() {}

// UnsafeIndexesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexesServer will
// result in compilation errors.
type UnsafeIndexesServer interface {
	mustEmbedUnimplementedIndexesServer()
}

func RegisterIndexesServer(s grpc.ServiceRegistrar, srv IndexesServer) {
	s.RegisterService(&Indexes_ServiceDesc, srv)
}

func _Indexes_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Get(ctx, req.(*IndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexes_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexesSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Search(ctx, req.(*IndexesSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexesCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Create(ctx, req.(*IndexesCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexes_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexesCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Update(ctx, req.(*IndexesCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Delete(ctx, req.(*IndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexes_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexesServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexes.v1.Indexes/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexesServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Indexes_ServiceDesc is the grpc.ServiceDesc for Indexes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Indexes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indexes.v1.Indexes",
	HandlerType: (*IndexesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Indexes_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Indexes_Search_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Indexes_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Indexes_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Indexes_Delete_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Indexes_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indexes/v1/indexes.proto",
}