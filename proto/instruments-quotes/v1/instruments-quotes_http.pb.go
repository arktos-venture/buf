// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1InstrumentsQuotes

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type QuotesHTTPServer interface {
	Last(context.Context, *QuoteRequest) (*QuoteReply, error)
	Search(context.Context, *QuoteSearchRequest) (*QuoteReplies, error)
}

func RegisterQuotesHTTPServer(s *http.Server, srv QuotesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/quote/{id}/last", _Quotes_Last1_HTTP_Handler(srv))
	r.POST("/v1/quotes", _Quotes_Search1_HTTP_Handler(srv))
}

func _Quotes_Last1_HTTP_Handler(srv QuotesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QuoteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments_quotes.v1.quotes/Last")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Last(ctx, req.(*QuoteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QuoteReply)
		return ctx.Result(200, reply)
	}
}

func _Quotes_Search1_HTTP_Handler(srv QuotesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QuoteSearchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments_quotes.v1.quotes/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*QuoteSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QuoteReplies)
		return ctx.Result(200, reply)
	}
}

type QuotesHTTPClient interface {
	Last(ctx context.Context, req *QuoteRequest, opts ...http.CallOption) (rsp *QuoteReply, err error)
	Search(ctx context.Context, req *QuoteSearchRequest, opts ...http.CallOption) (rsp *QuoteReplies, err error)
}

type QuotesHTTPClientImpl struct {
	cc *http.Client
}

func NewQuotesHTTPClient(client *http.Client) QuotesHTTPClient {
	return &QuotesHTTPClientImpl{client}
}

func (c *QuotesHTTPClientImpl) Last(ctx context.Context, in *QuoteRequest, opts ...http.CallOption) (*QuoteReply, error) {
	var out QuoteReply
	pattern := "/v1/quote/{id}/last"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/instruments_quotes.v1.quotes/Last"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QuotesHTTPClientImpl) Search(ctx context.Context, in *QuoteSearchRequest, opts ...http.CallOption) (*QuoteReplies, error) {
	var out QuoteReplies
	pattern := "/v1/quotes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/instruments_quotes.v1.quotes/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
