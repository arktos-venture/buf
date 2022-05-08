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

type DividendsHTTPServer interface {
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Last(context.Context, *DividendsLastRequest) (*DividendsLastReply, error)
	Search(context.Context, *DividendsRequest) (*DividendsReply, error)
}

func RegisterDividendsHTTPServer(s *http.Server, srv DividendsHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/dividends/{exchange}/{ticker}", _Dividends_Last1_HTTP_Handler(srv))
	r.POST("/v1/dividends", _Dividends_Search6_HTTP_Handler(srv))
	r.GET("/healthz", _Dividends_Health13_HTTP_Handler(srv))
}

func _Dividends_Last1_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DividendsLastRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Last")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Last(ctx, req.(*DividendsLastRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DividendsLastReply)
		return ctx.Result(200, reply)
	}
}

func _Dividends_Search6_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DividendsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*DividendsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DividendsReply)
		return ctx.Result(200, reply)
	}
}

func _Dividends_Health13_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Health")
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

type DividendsHTTPClient interface {
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Last(ctx context.Context, req *DividendsLastRequest, opts ...http.CallOption) (rsp *DividendsLastReply, err error)
	Search(ctx context.Context, req *DividendsRequest, opts ...http.CallOption) (rsp *DividendsReply, err error)
}

type DividendsHTTPClientImpl struct {
	cc *http.Client
}

func NewDividendsHTTPClient(client *http.Client) DividendsHTTPClient {
	return &DividendsHTTPClientImpl{client}
}

func (c *DividendsHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DividendsHTTPClientImpl) Last(ctx context.Context, in *DividendsLastRequest, opts ...http.CallOption) (*DividendsLastReply, error) {
	var out DividendsLastReply
	pattern := "/v1/dividends/{exchange}/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Last"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DividendsHTTPClientImpl) Search(ctx context.Context, in *DividendsRequest, opts ...http.CallOption) (*DividendsReply, error) {
	var out DividendsReply
	pattern := "/v1/dividends"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
