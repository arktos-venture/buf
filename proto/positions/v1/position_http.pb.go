// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package positions_v1

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

type PositionHTTPServer interface {
	Companies(context.Context, *PositionRequest) (*PositionCompanyReplies, error)
	Currencies(context.Context, *PositionRequest) (*PositionCurrencyReplies, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

func RegisterPositionHTTPServer(s *http.Server, srv PositionHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/positions/{account}/companies", _Position_Companies0_HTTP_Handler(srv))
	r.GET("/v1/positions/{account}/currencies", _Position_Currencies0_HTTP_Handler(srv))
	r.GET("/healthz", _Position_Health13_HTTP_Handler(srv))
}

func _Position_Companies0_HTTP_Handler(srv PositionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PositionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/positions.v1.Position/Companies")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Companies(ctx, req.(*PositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PositionCompanyReplies)
		return ctx.Result(200, reply)
	}
}

func _Position_Currencies0_HTTP_Handler(srv PositionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PositionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/positions.v1.Position/Currencies")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Currencies(ctx, req.(*PositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PositionCurrencyReplies)
		return ctx.Result(200, reply)
	}
}

func _Position_Health13_HTTP_Handler(srv PositionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/positions.v1.Position/Health")
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

type PositionHTTPClient interface {
	Companies(ctx context.Context, req *PositionRequest, opts ...http.CallOption) (rsp *PositionCompanyReplies, err error)
	Currencies(ctx context.Context, req *PositionRequest, opts ...http.CallOption) (rsp *PositionCurrencyReplies, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type PositionHTTPClientImpl struct {
	cc *http.Client
}

func NewPositionHTTPClient(client *http.Client) PositionHTTPClient {
	return &PositionHTTPClientImpl{client}
}

func (c *PositionHTTPClientImpl) Companies(ctx context.Context, in *PositionRequest, opts ...http.CallOption) (*PositionCompanyReplies, error) {
	var out PositionCompanyReplies
	pattern := "/v1/positions/{account}/companies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/positions.v1.Position/Companies"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PositionHTTPClientImpl) Currencies(ctx context.Context, in *PositionRequest, opts ...http.CallOption) (*PositionCurrencyReplies, error) {
	var out PositionCurrencyReplies
	pattern := "/v1/positions/{account}/currencies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/positions.v1.Position/Currencies"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PositionHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/positions.v1.Position/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
