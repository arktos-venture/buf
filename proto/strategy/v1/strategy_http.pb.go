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

type StrategyHTTPServer interface {
	Account(context.Context, *StrategyAccountRequest) (*StrategyReply, error)
	Company(context.Context, *StrategyCompanyRequest) (*StrategyReply, error)
	Country(context.Context, *StrategyCountryRequest) (*StrategyReply, error)
	Currency(context.Context, *StrategyCurrencyRequest) (*StrategyReply, error)
	Exchange(context.Context, *StrategyExchangeRequest) (*StrategyReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Index(context.Context, *StrategyIndexRequest) (*StrategyReply, error)
	Industry(context.Context, *StrategyIndustryRequest) (*StrategyReply, error)
}

func RegisterStrategyHTTPServer(s *http.Server, srv StrategyHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/strategy/company", _Strategy_Company2_HTTP_Handler(srv))
	r.POST("/v1/strategy/currency", _Strategy_Currency2_HTTP_Handler(srv))
	r.POST("/v1/strategy/{exchange}/industry", _Strategy_Industry2_HTTP_Handler(srv))
	r.POST("/v1/strategy/exchange", _Strategy_Exchange2_HTTP_Handler(srv))
	r.POST("/v1/strategy/country", _Strategy_Country2_HTTP_Handler(srv))
	r.POST("/v1/strategy/index", _Strategy_Index2_HTTP_Handler(srv))
	r.POST("/v1/strategy/account", _Strategy_Account2_HTTP_Handler(srv))
	r.GET("/healthz", _Strategy_Health16_HTTP_Handler(srv))
}

func _Strategy_Company2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyCompanyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Company")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Company(ctx, req.(*StrategyCompanyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Currency2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyCurrencyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Currency")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Currency(ctx, req.(*StrategyCurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Industry2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyIndustryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Industry")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Industry(ctx, req.(*StrategyIndustryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Exchange2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyExchangeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Exchange")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Exchange(ctx, req.(*StrategyExchangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Country2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyCountryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Country")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Country(ctx, req.(*StrategyCountryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Index2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyIndexRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Index")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Index(ctx, req.(*StrategyIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Account2_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Account")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Account(ctx, req.(*StrategyAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategy_Health16_HTTP_Handler(srv StrategyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategy.v1.Strategy/Health")
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

type StrategyHTTPClient interface {
	Account(ctx context.Context, req *StrategyAccountRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Company(ctx context.Context, req *StrategyCompanyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Country(ctx context.Context, req *StrategyCountryRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Currency(ctx context.Context, req *StrategyCurrencyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Exchange(ctx context.Context, req *StrategyExchangeRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Index(ctx context.Context, req *StrategyIndexRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Industry(ctx context.Context, req *StrategyIndustryRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
}

type StrategyHTTPClientImpl struct {
	cc *http.Client
}

func NewStrategyHTTPClient(client *http.Client) StrategyHTTPClient {
	return &StrategyHTTPClientImpl{client}
}

func (c *StrategyHTTPClientImpl) Account(ctx context.Context, in *StrategyAccountRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/account"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Account"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Company(ctx context.Context, in *StrategyCompanyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/company"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Company"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Country(ctx context.Context, in *StrategyCountryRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/country"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Country"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Currency(ctx context.Context, in *StrategyCurrencyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/currency"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Currency"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Exchange(ctx context.Context, in *StrategyExchangeRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/exchange"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Exchange"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Index(ctx context.Context, in *StrategyIndexRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/index"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Index"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategyHTTPClientImpl) Industry(ctx context.Context, in *StrategyIndustryRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/{exchange}/industry"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategy.v1.Strategy/Industry"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
