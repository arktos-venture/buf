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

// QuotesClient is the client API for Quotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotesClient interface {
	Company(ctx context.Context, in *QuotesCompanyRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Currency(ctx context.Context, in *QuotesCurrencyRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Industry(ctx context.Context, in *QuotesIndustryRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Exchange(ctx context.Context, in *QuotesExchangeRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Country(ctx context.Context, in *QuotesCountryRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Index(ctx context.Context, in *QuotesIndexRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Account(ctx context.Context, in *QuotesAccountRequest, opts ...grpc.CallOption) (*QuotesReply, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type quotesClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotesClient(cc grpc.ClientConnInterface) QuotesClient {
	return &quotesClient{cc}
}

func (c *quotesClient) Company(ctx context.Context, in *QuotesCompanyRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Company", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Currency(ctx context.Context, in *QuotesCurrencyRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Currency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Industry(ctx context.Context, in *QuotesIndustryRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Industry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Exchange(ctx context.Context, in *QuotesExchangeRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Exchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Country(ctx context.Context, in *QuotesCountryRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Country", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Index(ctx context.Context, in *QuotesIndexRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Account(ctx context.Context, in *QuotesAccountRequest, opts ...grpc.CallOption) (*QuotesReply, error) {
	out := new(QuotesReply)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotesClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/quotes.v1.Quotes/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotesServer is the server API for Quotes service.
// All implementations must embed UnimplementedQuotesServer
// for forward compatibility
type QuotesServer interface {
	Company(context.Context, *QuotesCompanyRequest) (*QuotesReply, error)
	Currency(context.Context, *QuotesCurrencyRequest) (*QuotesReply, error)
	Industry(context.Context, *QuotesIndustryRequest) (*QuotesReply, error)
	Exchange(context.Context, *QuotesExchangeRequest) (*QuotesReply, error)
	Country(context.Context, *QuotesCountryRequest) (*QuotesReply, error)
	Index(context.Context, *QuotesIndexRequest) (*QuotesReply, error)
	Account(context.Context, *QuotesAccountRequest) (*QuotesReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedQuotesServer()
}

// UnimplementedQuotesServer must be embedded to have forward compatible implementations.
type UnimplementedQuotesServer struct {
}

func (UnimplementedQuotesServer) Company(context.Context, *QuotesCompanyRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Company not implemented")
}
func (UnimplementedQuotesServer) Currency(context.Context, *QuotesCurrencyRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currency not implemented")
}
func (UnimplementedQuotesServer) Industry(context.Context, *QuotesIndustryRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Industry not implemented")
}
func (UnimplementedQuotesServer) Exchange(context.Context, *QuotesExchangeRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedQuotesServer) Country(context.Context, *QuotesCountryRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Country not implemented")
}
func (UnimplementedQuotesServer) Index(context.Context, *QuotesIndexRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedQuotesServer) Account(context.Context, *QuotesAccountRequest) (*QuotesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (UnimplementedQuotesServer) Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
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

func _Quotes_Company_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Company(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Company",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Company(ctx, req.(*QuotesCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Currency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Currency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Currency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Currency(ctx, req.(*QuotesCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Industry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesIndustryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Industry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Industry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Industry(ctx, req.(*QuotesIndustryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Exchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Exchange(ctx, req.(*QuotesExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Country_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Country(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Country",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Country(ctx, req.(*QuotesCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Index(ctx, req.(*QuotesIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Account(ctx, req.(*QuotesAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quotes_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.v1.Quotes/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Health(ctx, req.(*emptypb.Empty))
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
			MethodName: "Company",
			Handler:    _Quotes_Company_Handler,
		},
		{
			MethodName: "Currency",
			Handler:    _Quotes_Currency_Handler,
		},
		{
			MethodName: "Industry",
			Handler:    _Quotes_Industry_Handler,
		},
		{
			MethodName: "Exchange",
			Handler:    _Quotes_Exchange_Handler,
		},
		{
			MethodName: "Country",
			Handler:    _Quotes_Country_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Quotes_Index_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _Quotes_Account_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Quotes_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/quotes/v1/quotes.proto",
}
