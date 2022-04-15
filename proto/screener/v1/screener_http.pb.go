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

type ScreenerHTTPServer interface {
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Search(context.Context, *ScreenerRequest) (*ScreenerReplies, error)
}

func RegisterScreenerHTTPServer(s *http.Server, srv ScreenerHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/screener", _Screener_Search0_HTTP_Handler(srv))
	r.GET("/healthz", _Screener_Health2_HTTP_Handler(srv))
}

func _Screener_Search0_HTTP_Handler(srv ScreenerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ScreenerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/screener.v1.Screener/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*ScreenerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ScreenerReplies)
		return ctx.Result(200, reply)
	}
}

func _Screener_Health2_HTTP_Handler(srv ScreenerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/screener.v1.Screener/Health")
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

type ScreenerHTTPClient interface {
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Search(ctx context.Context, req *ScreenerRequest, opts ...http.CallOption) (rsp *ScreenerReplies, err error)
}

type ScreenerHTTPClientImpl struct {
	cc *http.Client
}

func NewScreenerHTTPClient(client *http.Client) ScreenerHTTPClient {
	return &ScreenerHTTPClientImpl{client}
}

func (c *ScreenerHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/screener.v1.Screener/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ScreenerHTTPClientImpl) Search(ctx context.Context, in *ScreenerRequest, opts ...http.CallOption) (*ScreenerReplies, error) {
	var out ScreenerReplies
	pattern := "/v1/screener"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/screener.v1.Screener/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
