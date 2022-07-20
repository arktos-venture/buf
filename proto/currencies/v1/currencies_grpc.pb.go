// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Currencies

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

// CurrenciesClient is the client API for Currencies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrenciesClient interface {
	// Public API: Get Currency properties
	Get(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyReply, error)
	// Public API: List Currencies available
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CurrencyReplies, error)
}

type currenciesClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrenciesClient(cc grpc.ClientConnInterface) CurrenciesClient {
	return &currenciesClient{cc}
}

func (c *currenciesClient) Get(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyReply, error) {
	out := new(CurrencyReply)
	err := c.cc.Invoke(ctx, "/currencies.v1.Currencies/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currenciesClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CurrencyReplies, error) {
	out := new(CurrencyReplies)
	err := c.cc.Invoke(ctx, "/currencies.v1.Currencies/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrenciesServer is the server API for Currencies service.
// All implementations must embed UnimplementedCurrenciesServer
// for forward compatibility
type CurrenciesServer interface {
	// Public API: Get Currency properties
	Get(context.Context, *CurrencyRequest) (*CurrencyReply, error)
	// Public API: List Currencies available
	List(context.Context, *emptypb.Empty) (*CurrencyReplies, error)
	mustEmbedUnimplementedCurrenciesServer()
}

// UnimplementedCurrenciesServer must be embedded to have forward compatible implementations.
type UnimplementedCurrenciesServer struct {
}

func (UnimplementedCurrenciesServer) Get(context.Context, *CurrencyRequest) (*CurrencyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCurrenciesServer) List(context.Context, *emptypb.Empty) (*CurrencyReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCurrenciesServer) mustEmbedUnimplementedCurrenciesServer() {}

// UnsafeCurrenciesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrenciesServer will
// result in compilation errors.
type UnsafeCurrenciesServer interface {
	mustEmbedUnimplementedCurrenciesServer()
}

func RegisterCurrenciesServer(s grpc.ServiceRegistrar, srv CurrenciesServer) {
	s.RegisterService(&Currencies_ServiceDesc, srv)
}

func _Currencies_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrenciesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currencies.v1.Currencies/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrenciesServer).Get(ctx, req.(*CurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currencies_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrenciesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currencies.v1.Currencies/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrenciesServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Currencies_ServiceDesc is the grpc.ServiceDesc for Currencies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Currencies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currencies.v1.Currencies",
	HandlerType: (*CurrenciesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Currencies_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Currencies_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/currencies/v1/currencies.proto",
}
