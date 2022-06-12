// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package currencies_v1

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

type CurrenciesHTTPServer interface {
	Delete(context.Context, *CurrencyDeleteRequest) (*CurrencyDeleteReply, error)
	Get(context.Context, *CurrencyRequest) (*CurrencyReply, error)
	List(context.Context, *emptypb.Empty) (*CurrencyReplies, error)
}

func RegisterCurrenciesHTTPServer(s *http.Server, srv CurrenciesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/currency/{ticker}", _Currencies_Get2_HTTP_Handler(srv))
	r.GET("/v1/currencies", _Currencies_List1_HTTP_Handler(srv))
	r.DELETE("/v1/currencies", _Currencies_Delete3_HTTP_Handler(srv))
}

func _Currencies_Get2_HTTP_Handler(srv CurrenciesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrencyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/currencies.v1.Currencies/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*CurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrencyReply)
		return ctx.Result(200, reply)
	}
}

func _Currencies_List1_HTTP_Handler(srv CurrenciesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/currencies.v1.Currencies/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrencyReplies)
		return ctx.Result(200, reply)
	}
}

func _Currencies_Delete3_HTTP_Handler(srv CurrenciesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrencyDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/currencies.v1.Currencies/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*CurrencyDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrencyDeleteReply)
		return ctx.Result(200, reply)
	}
}

type CurrenciesHTTPClient interface {
	Delete(ctx context.Context, req *CurrencyDeleteRequest, opts ...http.CallOption) (rsp *CurrencyDeleteReply, err error)
	Get(ctx context.Context, req *CurrencyRequest, opts ...http.CallOption) (rsp *CurrencyReply, err error)
	List(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CurrencyReplies, err error)
}

type CurrenciesHTTPClientImpl struct {
	cc *http.Client
}

func NewCurrenciesHTTPClient(client *http.Client) CurrenciesHTTPClient {
	return &CurrenciesHTTPClientImpl{client}
}

func (c *CurrenciesHTTPClientImpl) Delete(ctx context.Context, in *CurrencyDeleteRequest, opts ...http.CallOption) (*CurrencyDeleteReply, error) {
	var out CurrencyDeleteReply
	pattern := "/v1/currencies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/currencies.v1.Currencies/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CurrenciesHTTPClientImpl) Get(ctx context.Context, in *CurrencyRequest, opts ...http.CallOption) (*CurrencyReply, error) {
	var out CurrencyReply
	pattern := "/v1/currency/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/currencies.v1.Currencies/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CurrenciesHTTPClientImpl) List(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CurrencyReplies, error) {
	var out CurrencyReplies
	pattern := "/v1/currencies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/currencies.v1.Currencies/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
