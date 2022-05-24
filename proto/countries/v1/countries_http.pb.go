// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package countries_v1

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

type CountriesHTTPServer interface {
	Get(context.Context, *CountryRequest) (*CountryReply, error)
	Indicator(context.Context, *CountryIndicatorRequest) (*CountryIndicatorReply, error)
	Search(context.Context, *CountrySearchRequest) (*CountryReplies, error)
}

func RegisterCountriesHTTPServer(s *http.Server, srv CountriesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/country/{country}", _Countries_Get5_HTTP_Handler(srv))
	r.POST("/v1/countries", _Countries_Search5_HTTP_Handler(srv))
	r.GET("/v1/country/{country}/{indicator}", _Countries_Indicator0_HTTP_Handler(srv))
}

func _Countries_Get5_HTTP_Handler(srv CountriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/countries.v1.Countries/Get")
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

func _Countries_Search5_HTTP_Handler(srv CountriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountrySearchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/countries.v1.Countries/Search")
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

func _Countries_Indicator0_HTTP_Handler(srv CountriesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CountryIndicatorRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/countries.v1.Countries/Indicator")
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

type CountriesHTTPClient interface {
	Get(ctx context.Context, req *CountryRequest, opts ...http.CallOption) (rsp *CountryReply, err error)
	Indicator(ctx context.Context, req *CountryIndicatorRequest, opts ...http.CallOption) (rsp *CountryIndicatorReply, err error)
	Search(ctx context.Context, req *CountrySearchRequest, opts ...http.CallOption) (rsp *CountryReplies, err error)
}

type CountriesHTTPClientImpl struct {
	cc *http.Client
}

func NewCountriesHTTPClient(client *http.Client) CountriesHTTPClient {
	return &CountriesHTTPClientImpl{client}
}

func (c *CountriesHTTPClientImpl) Get(ctx context.Context, in *CountryRequest, opts ...http.CallOption) (*CountryReply, error) {
	var out CountryReply
	pattern := "/v1/country/{country}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/countries.v1.Countries/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CountriesHTTPClientImpl) Indicator(ctx context.Context, in *CountryIndicatorRequest, opts ...http.CallOption) (*CountryIndicatorReply, error) {
	var out CountryIndicatorReply
	pattern := "/v1/country/{country}/{indicator}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/countries.v1.Countries/Indicator"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CountriesHTTPClientImpl) Search(ctx context.Context, in *CountrySearchRequest, opts ...http.CallOption) (*CountryReplies, error) {
	var out CountryReplies
	pattern := "/v1/countries"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/countries.v1.Countries/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
