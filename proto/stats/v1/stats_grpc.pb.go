// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stats_v1

import (
	context "context"
	v1 "github.com/arktos-venture/buf/proto/quotes/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	// Public API
	// Get Stat properties
	Last(ctx context.Context, in *v1.QuoteLastRequest, opts ...grpc.CallOption) (*StatReply, error)
	// Public API
	// Get Stat properties
	Search(ctx context.Context, in *v1.QuoteRequest, opts ...grpc.CallOption) (*StatReplies, error)
	// Private API
	// Create a new Stat
	Create(ctx context.Context, in *StatCreateRequest, opts ...grpc.CallOption) (*StatReply, error)
	// Private API
	// Update properties of Stat
	Update(ctx context.Context, in *StatUpdateRequest, opts ...grpc.CallOption) (*StatReply, error)
	// Private API
	// Delete Stat
	Delete(ctx context.Context, in *v1.QuoteDeleteRequest, opts ...grpc.CallOption) (*StatDelete, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) Last(ctx context.Context, in *v1.QuoteLastRequest, opts ...grpc.CallOption) (*StatReply, error) {
	out := new(StatReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Last", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Search(ctx context.Context, in *v1.QuoteRequest, opts ...grpc.CallOption) (*StatReplies, error) {
	out := new(StatReplies)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Create(ctx context.Context, in *StatCreateRequest, opts ...grpc.CallOption) (*StatReply, error) {
	out := new(StatReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Update(ctx context.Context, in *StatUpdateRequest, opts ...grpc.CallOption) (*StatReply, error) {
	out := new(StatReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Delete(ctx context.Context, in *v1.QuoteDeleteRequest, opts ...grpc.CallOption) (*StatDelete, error) {
	out := new(StatDelete)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility
type StatsServer interface {
	// Public API
	// Get Stat properties
	Last(context.Context, *v1.QuoteLastRequest) (*StatReply, error)
	// Public API
	// Get Stat properties
	Search(context.Context, *v1.QuoteRequest) (*StatReplies, error)
	// Private API
	// Create a new Stat
	Create(context.Context, *StatCreateRequest) (*StatReply, error)
	// Private API
	// Update properties of Stat
	Update(context.Context, *StatUpdateRequest) (*StatReply, error)
	// Private API
	// Delete Stat
	Delete(context.Context, *v1.QuoteDeleteRequest) (*StatDelete, error)
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (UnimplementedStatsServer) Last(context.Context, *v1.QuoteLastRequest) (*StatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Last not implemented")
}
func (UnimplementedStatsServer) Search(context.Context, *v1.QuoteRequest) (*StatReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedStatsServer) Create(context.Context, *StatCreateRequest) (*StatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStatsServer) Update(context.Context, *StatUpdateRequest) (*StatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStatsServer) Delete(context.Context, *v1.QuoteDeleteRequest) (*StatDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_Last_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QuoteLastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Last(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Last",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Last(ctx, req.(*v1.QuoteLastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Search(ctx, req.(*v1.QuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Create(ctx, req.(*StatCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Update(ctx, req.(*StatUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QuoteDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Delete(ctx, req.(*v1.QuoteDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats.v1.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Last",
			Handler:    _Stats_Last_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Stats_Search_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Stats_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Stats_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Stats_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/stats/v1/stats.proto",
}
