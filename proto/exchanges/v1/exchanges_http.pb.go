// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package exchanges_v1

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

type ExchangesHTTPServer interface {
	Delete(context.Context, *ExchangeDeleteRequest) (*ExchangeDelete, error)
	Get(context.Context, *ExchangeRequest) (*ExchangeReply, error)
	IsOpen(context.Context, *ExchangeIsOpenRequest) (*ExchangeIsOpenReply, error)
	Search(context.Context, *ExchangeSearchRequest) (*ExchangeReplies, error)
}

func RegisterExchangesHTTPServer(s *http.Server, srv ExchangesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/exchange/{ticker}/isopen", _Exchanges_IsOpen0_HTTP_Handler(srv))
	r.GET("/v1/exchange/{ticker}", _Exchanges_Get3_HTTP_Handler(srv))
	r.GET("/v1/exchanges", _Exchanges_Search3_HTTP_Handler(srv))
	r.DELETE("/v1/exchanges", _Exchanges_Delete3_HTTP_Handler(srv))
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

func _Exchanges_Get3_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
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

func _Exchanges_Search3_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*ExchangeSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeReplies)
		return ctx.Result(200, reply)
	}
}

func _Exchanges_Delete3_HTTP_Handler(srv ExchangesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/exchanges.v1.Exchanges/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*ExchangeDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeDelete)
		return ctx.Result(200, reply)
	}
}

type ExchangesHTTPClient interface {
	Delete(ctx context.Context, req *ExchangeDeleteRequest, opts ...http.CallOption) (rsp *ExchangeDelete, err error)
	Get(ctx context.Context, req *ExchangeRequest, opts ...http.CallOption) (rsp *ExchangeReply, err error)
	IsOpen(ctx context.Context, req *ExchangeIsOpenRequest, opts ...http.CallOption) (rsp *ExchangeIsOpenReply, err error)
	Search(ctx context.Context, req *ExchangeSearchRequest, opts ...http.CallOption) (rsp *ExchangeReplies, err error)
}

type ExchangesHTTPClientImpl struct {
	cc *http.Client
}

func NewExchangesHTTPClient(client *http.Client) ExchangesHTTPClient {
	return &ExchangesHTTPClientImpl{client}
}

func (c *ExchangesHTTPClientImpl) Delete(ctx context.Context, in *ExchangeDeleteRequest, opts ...http.CallOption) (*ExchangeDelete, error) {
	var out ExchangeDelete
	pattern := "/v1/exchanges"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) Get(ctx context.Context, in *ExchangeRequest, opts ...http.CallOption) (*ExchangeReply, error) {
	var out ExchangeReply
	pattern := "/v1/exchange/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) IsOpen(ctx context.Context, in *ExchangeIsOpenRequest, opts ...http.CallOption) (*ExchangeIsOpenReply, error) {
	var out ExchangeIsOpenReply
	pattern := "/v1/exchange/{ticker}/isopen"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/IsOpen"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExchangesHTTPClientImpl) Search(ctx context.Context, in *ExchangeSearchRequest, opts ...http.CallOption) (*ExchangeReplies, error) {
	var out ExchangeReplies
	pattern := "/v1/exchanges"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/exchanges.v1.Exchanges/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
