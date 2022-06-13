// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package positions_v1

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

// PositionsClient is the client API for Positions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionsClient interface {
	// Public API
	// Search Positions opened
	Search(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReplies, error)
	// Private API
	// Delete Positions open or closed
	Delete(ctx context.Context, in *PositionDeleteRequest, opts ...grpc.CallOption) (*PositionDelete, error)
}

type positionsClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionsClient(cc grpc.ClientConnInterface) PositionsClient {
	return &positionsClient{cc}
}

func (c *positionsClient) Search(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReplies, error) {
	out := new(PositionReplies)
	err := c.cc.Invoke(ctx, "/positions.v1.Positions/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionsClient) Delete(ctx context.Context, in *PositionDeleteRequest, opts ...grpc.CallOption) (*PositionDelete, error) {
	out := new(PositionDelete)
	err := c.cc.Invoke(ctx, "/positions.v1.Positions/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionsServer is the server API for Positions service.
// All implementations must embed UnimplementedPositionsServer
// for forward compatibility
type PositionsServer interface {
	// Public API
	// Search Positions opened
	Search(context.Context, *PositionRequest) (*PositionReplies, error)
	// Private API
	// Delete Positions open or closed
	Delete(context.Context, *PositionDeleteRequest) (*PositionDelete, error)
	mustEmbedUnimplementedPositionsServer()
}

// UnimplementedPositionsServer must be embedded to have forward compatible implementations.
type UnimplementedPositionsServer struct {
}

func (UnimplementedPositionsServer) Search(context.Context, *PositionRequest) (*PositionReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedPositionsServer) Delete(context.Context, *PositionDeleteRequest) (*PositionDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPositionsServer) mustEmbedUnimplementedPositionsServer() {}

// UnsafePositionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionsServer will
// result in compilation errors.
type UnsafePositionsServer interface {
	mustEmbedUnimplementedPositionsServer()
}

func RegisterPositionsServer(s grpc.ServiceRegistrar, srv PositionsServer) {
	s.RegisterService(&Positions_ServiceDesc, srv)
}

func _Positions_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/positions.v1.Positions/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionsServer).Search(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Positions_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/positions.v1.Positions/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionsServer).Delete(ctx, req.(*PositionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Positions_ServiceDesc is the grpc.ServiceDesc for Positions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Positions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "positions.v1.Positions",
	HandlerType: (*PositionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Positions_Search_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Positions_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/positions/v1/positions.proto",
}
