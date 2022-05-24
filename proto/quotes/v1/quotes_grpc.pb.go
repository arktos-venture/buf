// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package quotes_v1

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

// QuotesClient is the client API for Quotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotesClient interface {
	Last(ctx context.Context, in *QuotesLastRequest, opts ...grpc.CallOption) (*QuotesLastReply, error)
	Search(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesReply, error)
}

type quotesClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotesClient(cc grpc.ClientConnInterface) QuotesClient {
	return &quotesClient{cc}
}

func (c *quotesClient) Last(ctx context.Context, in *QuotesLastRequest, opts ...grpc.CallOption) (*QuotesLastReply, error) {
	out := new(QuotesLastReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Last", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Search(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotesServer is the server API for Quotes service.
// All implementations must embed UnimplementedQuotesServer
// for forward compatibility
type QuotesServer interface {
	Last(context.Context, *QuotesLastRequest) (*QuotesLastReply, error)
	Search(context.Context, *QuotesRequest) (*QuotesReply, error)
	mustEmbedUnimplementedQuotesServer()
}

// UnimplementedQuotesServer must be embedded to have forward compatible implementations.
type UnimplementedQuotesServer struct {
}

func (UnimplementedQuotesServer) Last(context.Context, *QuotesLastRequest) (*QuotesLastReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Last not implemented")
}
func (UnimplementedQuotesServer) Search(context.Context, *QuotesRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedQuotesServer) mustEmbedUnimplementedQuotesServer() {}

// UnsafeQuotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotesServer will
// result in compilation errors.
type UnsafeQuotesServer interface {
	mustEmbedUnimplementedQuotesServer()
}

func RegisterQuotesServer(s grpc.ServiceRegistrar, srv QuotesServer) {
	s.RegisterService(&Quotes_ServiceDesc, srv)
}

func _Quotes_Last_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesLastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Last(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Last",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Last(ctx, req.(*QuotesLastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Search(ctx, req.(*QuotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Quotes_ServiceDesc is the grpc.ServiceDesc for Quotes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Quotes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quotes.v1.Quotes",
	HandlerType: (*QuotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Last",
			Handler:    _Quotes_Last_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Quotes_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/quotes/v1/quotes.proto",
}
