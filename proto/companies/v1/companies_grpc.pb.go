// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1Companies

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

// CompaniesClient is the client API for Companies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompaniesClient interface {
	// Public API
	// Get Company properties
	Get(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyReply, error)
	// Private API
	// List Companies, only for start imports
	List(ctx context.Context, in *CompanyListRequest, opts ...grpc.CallOption) (*CompanyReplies, error)
	// Private API
	// Create a new Company
	Create(ctx context.Context, in *CompanyCreateRequest, opts ...grpc.CallOption) (*CompanyReply, error)
	// Private API
	// Update properties of Company
	Update(ctx context.Context, in *CompanyUpdateRequest, opts ...grpc.CallOption) (*CompanyReply, error)
	// Private API
	// Delete Company
	Delete(ctx context.Context, in *CompanyDeleteRequest, opts ...grpc.CallOption) (*CompanyDelete, error)
}

type companiesClient struct {
	cc grpc.ClientConnInterface
}

func NewCompaniesClient(cc grpc.ClientConnInterface) CompaniesClient {
	return &companiesClient{cc}
}

func (c *companiesClient) Get(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyReply, error) {
	out := new(CompanyReply)
	err := c.cc.Invoke(ctx, "/companies.v1.Companies/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) List(ctx context.Context, in *CompanyListRequest, opts ...grpc.CallOption) (*CompanyReplies, error) {
	out := new(CompanyReplies)
	err := c.cc.Invoke(ctx, "/companies.v1.Companies/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) Create(ctx context.Context, in *CompanyCreateRequest, opts ...grpc.CallOption) (*CompanyReply, error) {
	out := new(CompanyReply)
	err := c.cc.Invoke(ctx, "/companies.v1.Companies/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) Update(ctx context.Context, in *CompanyUpdateRequest, opts ...grpc.CallOption) (*CompanyReply, error) {
	out := new(CompanyReply)
	err := c.cc.Invoke(ctx, "/companies.v1.Companies/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) Delete(ctx context.Context, in *CompanyDeleteRequest, opts ...grpc.CallOption) (*CompanyDelete, error) {
	out := new(CompanyDelete)
	err := c.cc.Invoke(ctx, "/companies.v1.Companies/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompaniesServer is the server API for Companies service.
// All implementations must embed UnimplementedCompaniesServer
// for forward compatibility
type CompaniesServer interface {
	// Public API
	// Get Company properties
	Get(context.Context, *CompanyRequest) (*CompanyReply, error)
	// Private API
	// List Companies, only for start imports
	List(context.Context, *CompanyListRequest) (*CompanyReplies, error)
	// Private API
	// Create a new Company
	Create(context.Context, *CompanyCreateRequest) (*CompanyReply, error)
	// Private API
	// Update properties of Company
	Update(context.Context, *CompanyUpdateRequest) (*CompanyReply, error)
	// Private API
	// Delete Company
	Delete(context.Context, *CompanyDeleteRequest) (*CompanyDelete, error)
	mustEmbedUnimplementedCompaniesServer()
}

// UnimplementedCompaniesServer must be embedded to have forward compatible implementations.
type UnimplementedCompaniesServer struct {
}

func (UnimplementedCompaniesServer) Get(context.Context, *CompanyRequest) (*CompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCompaniesServer) List(context.Context, *CompanyListRequest) (*CompanyReplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCompaniesServer) Create(context.Context, *CompanyCreateRequest) (*CompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCompaniesServer) Update(context.Context, *CompanyUpdateRequest) (*CompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCompaniesServer) Delete(context.Context, *CompanyDeleteRequest) (*CompanyDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCompaniesServer) mustEmbedUnimplementedCompaniesServer() {}

// UnsafeCompaniesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompaniesServer will
// result in compilation errors.
type UnsafeCompaniesServer interface {
	mustEmbedUnimplementedCompaniesServer()
}

func RegisterCompaniesServer(s grpc.ServiceRegistrar, srv CompaniesServer) {
	s.RegisterService(&Companies_ServiceDesc, srv)
}

func _Companies_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.v1.Companies/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).Get(ctx, req.(*CompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.v1.Companies/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).List(ctx, req.(*CompanyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.v1.Companies/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).Create(ctx, req.(*CompanyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.v1.Companies/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).Update(ctx, req.(*CompanyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.v1.Companies/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).Delete(ctx, req.(*CompanyDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Companies_ServiceDesc is the grpc.ServiceDesc for Companies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Companies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "companies.v1.Companies",
	HandlerType: (*CompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Companies_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Companies_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Companies_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Companies_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Companies_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/companies/v1/companies.proto",
}
