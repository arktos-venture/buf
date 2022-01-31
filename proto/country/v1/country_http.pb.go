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

type CountryHTTPServer interface {
	Get(context.Context, *CountryRequest) (*CountryReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Indicator(context.Context, *CountryIndicatorRequest) (*CountryIndicatorReply, error)
	Search(context.Context, *CountrySearchRequest) (*CountryReplies, error)
}

func RegisterCountryHTTPServer(s *http.Server, srv CountryHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/country/{country}", _Country_Get6_HTTP_Handler(srv))
	r.POST("/v1/country", _Country_Search6_HTTP_Handler(srv))
	r.GET("/v1/country/{country}/{indicator}", _Country_Indicator0_HTTP_Handler(srv))
	r.GET("/healthz", _Country_Health14_HTTP_Handler(srv))
}

func _Country_Get6_HTTP_Handler(srv CountryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/country.info.v1.get.Country/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*CountryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CountryReply)
		return ctx.Result(200, reply)
	}
}

func _Country_Search6_HTTP_Handler(srv CountryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountrySearchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/country.info.v1.get.Country/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*CountrySearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CountryReplies)
		return ctx.Result(200, reply)
	}
}

func _Country_Indicator0_HTTP_Handler(srv CountryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountryIndicatorRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/country.info.v1.get.Country/Indicator")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Indicator(ctx, req.(*CountryIndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CountryIndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Country_Health14_HTTP_Handler(srv CountryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/country.info.v1.get.Country/Health")
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

type CountryHTTPClient interface {
	Get(ctx context.Context, req *CountryRequest, opts ...http.CallOption) (rsp *CountryReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Indicator(ctx context.Context, req *CountryIndicatorRequest, opts ...http.CallOption) (rsp *CountryIndicatorReply, err error)
	Search(ctx context.Context, req *CountrySearchRequest, opts ...http.CallOption) (rsp *CountryReplies, err error)
}

type CountryHTTPClientImpl struct {
	cc *http.Client
}

func NewCountryHTTPClient(client *http.Client) CountryHTTPClient {
	return &CountryHTTPClientImpl{client}
}

func (c *CountryHTTPClientImpl) Get(ctx context.Context, in *CountryRequest, opts ...http.CallOption) (*CountryReply, error) {
	var out CountryReply
	pattern := "/v1/country/{country}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/country.info.v1.get.Country/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CountryHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/country.info.v1.get.Country/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CountryHTTPClientImpl) Indicator(ctx context.Context, in *CountryIndicatorRequest, opts ...http.CallOption) (*CountryIndicatorReply, error) {
	var out CountryIndicatorReply
	pattern := "/v1/country/{country}/{indicator}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/country.info.v1.get.Country/Indicator"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CountryHTTPClientImpl) Search(ctx context.Context, in *CountrySearchRequest, opts ...http.CallOption) (*CountryReplies, error) {
	var out CountryReplies
	pattern := "/v1/country"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/country.info.v1.get.Country/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
