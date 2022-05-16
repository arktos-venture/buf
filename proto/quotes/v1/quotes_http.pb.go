// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package quotes_v1

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

type QuotesHTTPServer interface {
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Last(context.Context, *QuotesLastRequest) (*QuotesLastReply, error)
	Search(context.Context, *QuotesRequest) (*QuotesReply, error)
}

func RegisterQuotesHTTPServer(s *http.Server, srv QuotesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/quotes/{exchange}/{ticker}", _Quotes_Last0_HTTP_Handler(srv))
	r.POST("/v1/quotes", _Quotes_Search2_HTTP_Handler(srv))
	r.GET("/healthz", _Quotes_Health4_HTTP_Handler(srv))
}

func _Quotes_Last0_HTTP_Handler(srv QuotesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QuotesLastRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/quotes.v1.Quotes/Last")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Last(ctx, req.(*QuotesLastRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QuotesLastReply)
		return ctx.Result(200, reply)
	}
}

func _Quotes_Search2_HTTP_Handler(srv QuotesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QuotesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/quotes.v1.Quotes/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*QuotesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QuotesReply)
		return ctx.Result(200, reply)
	}
}

func _Quotes_Health4_HTTP_Handler(srv QuotesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/quotes.v1.Quotes/Health")
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

type QuotesHTTPClient interface {
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Last(ctx context.Context, req *QuotesLastRequest, opts ...http.CallOption) (rsp *QuotesLastReply, err error)
	Search(ctx context.Context, req *QuotesRequest, opts ...http.CallOption) (rsp *QuotesReply, err error)
}

type QuotesHTTPClientImpl struct {
	cc *http.Client
}

func NewQuotesHTTPClient(client *http.Client) QuotesHTTPClient {
	return &QuotesHTTPClientImpl{client}
}

func (c *QuotesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/quotes.v1.Quotes/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuotesHTTPClientImpl) Last(ctx context.Context, in *QuotesLastRequest, opts ...http.CallOption) (*QuotesLastReply, error) {
	var out QuotesLastReply
	pattern := "/v1/quotes/{exchange}/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/quotes.v1.Quotes/Last"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuotesHTTPClientImpl) Search(ctx context.Context, in *QuotesRequest, opts ...http.CallOption) (*QuotesReply, error) {
	var out QuotesReply
	pattern := "/v1/quotes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/quotes.v1.Quotes/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
