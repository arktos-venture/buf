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

type ExchangesHTTPServer interface {
	Get(context.Context, *ExchangeRequest) (*ExchangeReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	IsOpen(context.Context, *ExchangeIsOpenRequest) (*ExchangeIsOpenReply, error)
	List(context.Context, *ExchangeListRequest) (*ExchangeReplies, error)
}

func RegisterExchangesHTTPServer(s *http.Server, srv ExchangesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/exchange/{exchange}/isopen", _Exchanges_IsOpen0_HTTP_Handler(srv))
	r.GET("/v1/exchange/{exchange}", _Exchanges_Get1_HTTP_Handler(srv))
	r.GET("/v1/exchanges", _Exchanges_List0_HTTP_Handler(srv))
	r.GET("/healthz", _Exchanges_Health4_HTTP_Handler(srv))
}

func _Exchanges_IsOpen0_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeIsOpenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/IsOpen")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.IsOpen(ctx, req.(*ExchangeIsOpenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeIsOpenReply)
		return ctx.Result(200, reply)
	}
}

func _Exchanges_Get1_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*ExchangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeReply)
		return ctx.Result(200, reply)
	}
}

func _Exchanges_List0_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ExchangeListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeReplies)
		return ctx.Result(200, reply)
	}
}

func _Exchanges_Health4_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/Health")
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

type ExchangesHTTPClient interface {
	Get(ctx context.Context, req *ExchangeRequest, opts ...http.CallOption) (rsp *ExchangeReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	IsOpen(ctx context.Context, req *ExchangeIsOpenRequest, opts ...http.CallOption) (rsp *ExchangeIsOpenReply, err error)
	List(ctx context.Context, req *ExchangeListRequest, opts ...http.CallOption) (rsp *ExchangeReplies, err error)
}

type ExchangesHTTPClientImpl struct {
	cc *http.Client
}

func NewExchangesHTTPClient(client *http.Client) ExchangesHTTPClient {
	return &ExchangesHTTPClientImpl{client}
}

func (c *ExchangesHTTPClientImpl) Get(ctx context.Context, in *ExchangeRequest, opts ...http.CallOption) (*ExchangeReply, error) {
	var out ExchangeReply
	pattern := "/v1/exchange/{exchange}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) IsOpen(ctx context.Context, in *ExchangeIsOpenRequest, opts ...http.CallOption) (*ExchangeIsOpenReply, error) {
	var out ExchangeIsOpenReply
	pattern := "/v1/exchange/{exchange}/isopen"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/IsOpen"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) List(ctx context.Context, in *ExchangeListRequest, opts ...http.CallOption) (*ExchangeReplies, error) {
	var out ExchangeReplies
	pattern := "/v1/exchanges"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
