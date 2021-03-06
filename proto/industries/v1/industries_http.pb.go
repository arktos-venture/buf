// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Industries

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type IndustriesHTTPServer interface {
	Activities(context.Context, *IndustryActivitiesRequest) (*IndustrySearchReply, error)
	List(context.Context, *emptypb.Empty) (*IndustryReply, error)
	Search(context.Context, *IndustrySearchRequest) (*IndustrySearchReply, error)
}

func RegisterIndustriesHTTPServer(s *http.Server, srv IndustriesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/industries", _Industries_List2_HTTP_Handler(srv))
	r.POST("/v1/industries/activities", _Industries_Activities0_HTTP_Handler(srv))
	r.GET("/v1/industry/{ticker}", _Industries_Search6_HTTP_Handler(srv))
}

func _Industries_List2_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/industries.v1.Industries/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndustryReply)
		return ctx.Result(200, reply)
	}
}

func _Industries_Activities0_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndustryActivitiesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/industries.v1.Industries/Activities")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Activities(ctx, req.(*IndustryActivitiesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndustrySearchReply)
		return ctx.Result(200, reply)
	}
}

func _Industries_Search6_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndustrySearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/industries.v1.Industries/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*IndustrySearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndustrySearchReply)
		return ctx.Result(200, reply)
	}
}

type IndustriesHTTPClient interface {
	Activities(ctx context.Context, req *IndustryActivitiesRequest, opts ...http.CallOption) (rsp *IndustrySearchReply, err error)
	List(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *IndustryReply, err error)
	Search(ctx context.Context, req *IndustrySearchRequest, opts ...http.CallOption) (rsp *IndustrySearchReply, err error)
}

type IndustriesHTTPClientImpl struct {
	cc *http.Client
}

func NewIndustriesHTTPClient(client *http.Client) IndustriesHTTPClient {
	return &IndustriesHTTPClientImpl{client}
}

func (c *IndustriesHTTPClientImpl) Activities(ctx context.Context, in *IndustryActivitiesRequest, opts ...http.CallOption) (*IndustrySearchReply, error) {
	var out IndustrySearchReply
	pattern := "/v1/industries/activities"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/industries.v1.Industries/Activities"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndustriesHTTPClientImpl) List(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*IndustryReply, error) {
	var out IndustryReply
	pattern := "/v1/industries"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/industries.v1.Industries/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndustriesHTTPClientImpl) Search(ctx context.Context, in *IndustrySearchRequest, opts ...http.CallOption) (*IndustrySearchReply, error) {
	var out IndustrySearchReply
	pattern := "/v1/industry/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/industries.v1.Industries/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
