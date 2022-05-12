// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1

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
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	List(context.Context, *emptypb.Empty) (*IndustryReply, error)
	Search(context.Context, *IndustrySearchRequest) (*IndustrySearchReply, error)
}

func RegisterIndustriesHTTPServer(s *http.Server, srv IndustriesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/industries", _Industries_List3_HTTP_Handler(srv))
	r.POST("/v1/industries", _Industries_Search2_HTTP_Handler(srv))
	r.GET("/healthz", _Industries_Health8_HTTP_Handler(srv))
}

func _Industries_List3_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
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

func _Industries_Search2_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndustrySearchRequest
		if err := ctx.Bind(&in); err != nil {
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

func _Industries_Health8_HTTP_Handler(srv IndustriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/industries.v1.Industries/Health")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Health(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type IndustriesHTTPClient interface {
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	List(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *IndustryReply, err error)
	Search(ctx context.Context, req *IndustrySearchRequest, opts ...http.CallOption) (rsp *IndustrySearchReply, err error)
}

type IndustriesHTTPClientImpl struct {
	cc *http.Client
}

func NewIndustriesHTTPClient(client *http.Client) IndustriesHTTPClient {
	return &IndustriesHTTPClientImpl{client}
}

func (c *IndustriesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/industries.v1.Industries/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
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
	pattern := "/v1/industries"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/industries.v1.Industries/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
