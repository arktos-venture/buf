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

// CountriesClient is the client API for Countries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountriesClient interface {
	Get(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryReply, error)
	Search(ctx context.Context, in *CountrySearchRequest, opts ...grpc.CallOption) (*CountryReplies, error)
	Indicator(ctx context.Context, in *CountryIndicatorRequest, opts ...grpc.CallOption) (*CountryIndicatorReply, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type countriesClient struct {
	cc grpc.ClientConnInterface
}

func NewCountriesClient(cc grpc.ClientConnInterface) CountriesClient {
	return &countriesClient{cc}
}

func (c *countriesClient) Get(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryReply, error) {
	out := new(CountryReply)
	err := c.cc.Invoke(ctx, "/countries.v1.Countries/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesClient) Search(ctx context.Context, in *CountrySearchRequest, opts ...grpc.CallOption) (*CountryReplies, error) {
	out := new(CountryReplies)
	err := c.cc.Invoke(ctx, "/countries.v1.Countries/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesClient) Indicator(ctx context.Context, in *CountryIndicatorRequest, opts ...grpc.CallOption) (*CountryIndicatorReply, error) {
	out := new(CountryIndicatorReply)
	err := c.cc.Invoke(ctx, "/countries.v1.Countries/Indicator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/countries.v1.Countries/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountriesServer is the server API for Countries service.
// All implementations must embed UnimplementedCountriesServer
// for forward compatibility
type CountriesServer interface {
	Get(context.Context, *CountryRequest) (*CountryReply, error)
	Search(context.Context, *CountrySearchRequest) (*CountryReplies, error)
	Indicator(context.Context, *CountryIndicatorRequest) (*CountryIndicatorReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedCountriesServer()
}

// UnimplementedCountriesServer must be embedded to have forward compatible implementations.
type UnimplementedCountriesServer struct {
}

func (UnimplementedCountriesServer) Get(context.Context, *CountryRequest) (*CountryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCountriesServer) Search(context.Context, *CountrySearchRequest) (*CountryReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedCountriesServer) Indicator(context.Context, *CountryIndicatorRequest) (*CountryIndicatorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Indicator not implemented")
}
func (UnimplementedCountriesServer) Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCountriesServer) mustEmbedUnimplementedCountriesServer() {}

// UnsafeCountriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountriesServer will
// result in compilation errors.
type UnsafeCountriesServer interface {
	mustEmbedUnimplementedCountriesServer()
}

func RegisterCountriesServer(s grpc.ServiceRegistrar, srv CountriesServer) {
	s.RegisterService(&Countries_ServiceDesc, srv)
}

func _Countries_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.v1.Countries/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Get(ctx, req.(*CountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Countries_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountrySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.v1.Countries/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Search(ctx, req.(*CountrySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Countries_Indicator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryIndicatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Indicator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.v1.Countries/Indicator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Indicator(ctx, req.(*CountryIndicatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Countries_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.v1.Countries/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Countries_ServiceDesc is the grpc.ServiceDesc for Countries service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Countries_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "countries.v1.Countries",
	HandlerType: (*CountriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Countries_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Countries_Search_Handler,
		},
		{
			MethodName: "Indicator",
			Handler:    _Countries_Indicator_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Countries_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/countries/v1/countries.proto",
}
