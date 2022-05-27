// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package indices_v1

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

type IndicesHTTPServer interface {
	Get(context.Context, *IndicesRequest) (*IndicesReply, error)
	List(context.Context, *IndicesExchangeRequest) (*IndicesShortReplies, error)
}

func RegisterIndicesHTTPServer(s *http.Server, srv IndicesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/indice/{ticker}", _Indices_Get1_HTTP_Handler(srv))
	r.GET("/v1/indices/exchange/{exchange}", _Indices_List0_HTTP_Handler(srv))
}

func _Indices_Get1_HTTP_Handler(srv IndicesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indices.v1.Indices/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*IndicesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicesReply)
		return ctx.Result(200, reply)
	}
}

func _Indices_List0_HTTP_Handler(srv IndicesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicesExchangeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indices.v1.Indices/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*IndicesExchangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicesShortReplies)
		return ctx.Result(200, reply)
	}
}

type IndicesHTTPClient interface {
	Get(ctx context.Context, req *IndicesRequest, opts ...http.CallOption) (rsp *IndicesReply, err error)
	List(ctx context.Context, req *IndicesExchangeRequest, opts ...http.CallOption) (rsp *IndicesShortReplies, err error)
}

type IndicesHTTPClientImpl struct {
	cc *http.Client
}

func NewIndicesHTTPClient(client *http.Client) IndicesHTTPClient {
	return &IndicesHTTPClientImpl{client}
}

func (c *IndicesHTTPClientImpl) Get(ctx context.Context, in *IndicesRequest, opts ...http.CallOption) (*IndicesReply, error) {
	var out IndicesReply
	pattern := "/v1/indice/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indices.v1.Indices/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicesHTTPClientImpl) List(ctx context.Context, in *IndicesExchangeRequest, opts ...http.CallOption) (*IndicesShortReplies, error) {
	var out IndicesShortReplies
	pattern := "/v1/indices/exchange/{exchange}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indices.v1.Indices/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
