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

type FundamentalHTTPServer interface {
	Get(context.Context, *FundamentalRequest) (*FundamentalReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

func RegisterFundamentalHTTPServer(s *http.Server, srv FundamentalHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/company/fundamentals", _Fundamental_Get2_HTTP_Handler(srv))
	r.GET("/healthz", _Fundamental_Health6_HTTP_Handler(srv))
}

func _Fundamental_Get2_HTTP_Handler(srv FundamentalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FundamentalRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fundamental.v1.Fundamental/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*FundamentalRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FundamentalReply)
		return ctx.Result(200, reply)
	}
}

func _Fundamental_Health6_HTTP_Handler(srv FundamentalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fundamental.v1.Fundamental/Health")
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

type FundamentalHTTPClient interface {
	Get(ctx context.Context, req *FundamentalRequest, opts ...http.CallOption) (rsp *FundamentalReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type FundamentalHTTPClientImpl struct {
	cc *http.Client
}

func NewFundamentalHTTPClient(client *http.Client) FundamentalHTTPClient {
	return &FundamentalHTTPClientImpl{client}
}

func (c *FundamentalHTTPClientImpl) Get(ctx context.Context, in *FundamentalRequest, opts ...http.CallOption) (*FundamentalReply, error) {
	var out FundamentalReply
	pattern := "/v1/company/fundamentals"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/fundamental.v1.Fundamental/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FundamentalHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/fundamental.v1.Fundamental/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
