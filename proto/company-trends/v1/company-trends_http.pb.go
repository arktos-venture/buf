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

type CompanyTrendsHTTPServer interface {
	Get(context.Context, *TrendRequest) (*TrendIdReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Search(context.Context, *TrendRequest) (*TrendReply, error)
}

func RegisterCompanyTrendsHTTPServer(s *http.Server, srv CompanyTrendsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/company/trendId", _CompanyTrends_Get8_HTTP_Handler(srv))
	r.POST("/v1/company/trends", _CompanyTrends_Search10_HTTP_Handler(srv))
	r.GET("/healthz", _CompanyTrends_Health19_HTTP_Handler(srv))
}

func _CompanyTrends_Get8_HTTP_Handler(srv CompanyTrendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TrendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/company.trends.v1.CompanyTrends/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*TrendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TrendIdReply)
		return ctx.Result(200, reply)
	}
}

func _CompanyTrends_Search10_HTTP_Handler(srv CompanyTrendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TrendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/company.trends.v1.CompanyTrends/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*TrendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TrendReply)
		return ctx.Result(200, reply)
	}
}

func _CompanyTrends_Health19_HTTP_Handler(srv CompanyTrendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/company.trends.v1.CompanyTrends/Health")
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

type CompanyTrendsHTTPClient interface {
	Get(ctx context.Context, req *TrendRequest, opts ...http.CallOption) (rsp *TrendIdReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Search(ctx context.Context, req *TrendRequest, opts ...http.CallOption) (rsp *TrendReply, err error)
}

type CompanyTrendsHTTPClientImpl struct {
	cc *http.Client
}

func NewCompanyTrendsHTTPClient(client *http.Client) CompanyTrendsHTTPClient {
	return &CompanyTrendsHTTPClientImpl{client}
}

func (c *CompanyTrendsHTTPClientImpl) Get(ctx context.Context, in *TrendRequest, opts ...http.CallOption) (*TrendIdReply, error) {
	var out TrendIdReply
	pattern := "/v1/company/trendId"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/company.trends.v1.CompanyTrends/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyTrendsHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/company.trends.v1.CompanyTrends/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyTrendsHTTPClientImpl) Search(ctx context.Context, in *TrendRequest, opts ...http.CallOption) (*TrendReply, error) {
	var out TrendReply
	pattern := "/v1/company/trends"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/company.trends.v1.CompanyTrends/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
