// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package screener_v1

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

// ScreenerClient is the client API for Screener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScreenerClient interface {
	// Public API
	Search(ctx context.Context, in *ScreenerRequest, opts ...grpc.CallOption) (*ScreenerReplies, error)
}

type screenerClient struct {
	cc grpc.ClientConnInterface
}

func NewScreenerClient(cc grpc.ClientConnInterface) ScreenerClient {
	return &screenerClient{cc}
}

func (c *screenerClient) Search(ctx context.Context, in *ScreenerRequest, opts ...grpc.CallOption) (*ScreenerReplies, error) {
	out := new(ScreenerReplies)
	err := c.cc.Invoke(ctx, "/screener.v1.Screener/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScreenerServer is the server API for Screener service.
// All implementations must embed UnimplementedScreenerServer
// for forward compatibility
type ScreenerServer interface {
	// Public API
	Search(context.Context, *ScreenerRequest) (*ScreenerReplies, error)
	mustEmbedUnimplementedScreenerServer()
}

// UnimplementedScreenerServer must be embedded to have forward compatible implementations.
type UnimplementedScreenerServer struct {
}

func (UnimplementedScreenerServer) Search(context.Context, *ScreenerRequest) (*ScreenerReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedScreenerServer) mustEmbedUnimplementedScreenerServer() {}

// UnsafeScreenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScreenerServer will
// result in compilation errors.
type UnsafeScreenerServer interface {
	mustEmbedUnimplementedScreenerServer()
}

func RegisterScreenerServer(s grpc.ServiceRegistrar, srv ScreenerServer) {
	s.RegisterService(&Screener_ServiceDesc, srv)
}

func _Screener_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScreenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenerServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/screener.v1.Screener/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenerServer).Search(ctx, req.(*ScreenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Screener_ServiceDesc is the grpc.ServiceDesc for Screener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Screener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "screener.v1.Screener",
	HandlerType: (*ScreenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Screener_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/screener/v1/screener.proto",
}
