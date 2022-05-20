// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package forexes_v1

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

type ForexesHTTPServer interface {
	Get(context.Context, *ForexRequest) (*ForexReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	List(context.Context, *ForexListRequest) (*ForexListReply, error)
}

func RegisterForexesHTTPServer(s *http.Server, srv ForexesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/forexes/{forex}", _Forexes_Get4_HTTP_Handler(srv))
	r.GET("/v1/forexes/{currency}/pairs", _Forexes_List4_HTTP_Handler(srv))
	r.GET("/healthz", _Forexes_Health11_HTTP_Handler(srv))
}

func _Forexes_Get4_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*ForexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForexReply)
		return ctx.Result(200, reply)
	}
}

func _Forexes_List4_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ForexListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForexListReply)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Health11_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Health")
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

type ForexesHTTPClient interface {
	Get(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *ForexReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	List(ctx context.Context, req *ForexListRequest, opts ...http.CallOption) (rsp *ForexListReply, err error)
}

type ForexesHTTPClientImpl struct {
	cc *http.Client
}

func NewForexesHTTPClient(client *http.Client) ForexesHTTPClient {
	return &ForexesHTTPClientImpl{client}
}

func (c *ForexesHTTPClientImpl) Get(ctx context.Context, in *ForexRequest, opts ...http.CallOption) (*ForexReply, error) {
	var out ForexReply
	pattern := "/v1/forexes/{forex}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) List(ctx context.Context, in *ForexListRequest, opts ...http.CallOption) (*ForexListReply, error) {
	var out ForexListReply
	pattern := "/v1/forexes/{currency}/pairs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
