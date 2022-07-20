// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Instruments

import (
	context "context"
	v1 "github.com/arktos-venture/buf/proto/strategies/v1"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type InstrumentsHTTPServer interface {
	Get(context.Context, *InstrumentRequest) (*InstrumentReply, error)
	Search(context.Context, *InstrumentSearchRequest) (*InstrumentReplies, error)
	Stats(context.Context, *InstrumentRequest) (*InstrumentStatsReply, error)
	Strategies(context.Context, *InstrumentStrategiesRequest) (*v1.StrategiesReplies, error)
}

func RegisterInstrumentsHTTPServer(s *http.Server, srv InstrumentsHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/instrument/{id}", _Instruments_Get2_HTTP_Handler(srv))
	r.GET("/v1/instrument/{id}/stats", _Instruments_Stats0_HTTP_Handler(srv))
	r.GET("/v1/instrument/{id}/strategies", _Instruments_Strategies0_HTTP_Handler(srv))
	r.POST("/v1/instruments/search", _Instruments_Search2_HTTP_Handler(srv))
}

func _Instruments_Get2_HTTP_Handler(srv InstrumentsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InstrumentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments.v1.instruments/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*InstrumentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InstrumentReply)
		return ctx.Result(200, reply)
	}
}

func _Instruments_Stats0_HTTP_Handler(srv InstrumentsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InstrumentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments.v1.instruments/Stats")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Stats(ctx, req.(*InstrumentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InstrumentStatsReply)
		return ctx.Result(200, reply)
	}
}

func _Instruments_Strategies0_HTTP_Handler(srv InstrumentsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InstrumentStrategiesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments.v1.instruments/Strategies")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Strategies(ctx, req.(*InstrumentStrategiesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.StrategiesReplies)
		return ctx.Result(200, reply)
	}
}

func _Instruments_Search2_HTTP_Handler(srv InstrumentsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InstrumentSearchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/instruments.v1.instruments/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*InstrumentSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InstrumentReplies)
		return ctx.Result(200, reply)
	}
}

type InstrumentsHTTPClient interface {
	Get(ctx context.Context, req *InstrumentRequest, opts ...http.CallOption) (rsp *InstrumentReply, err error)
	Search(ctx context.Context, req *InstrumentSearchRequest, opts ...http.CallOption) (rsp *InstrumentReplies, err error)
	Stats(ctx context.Context, req *InstrumentRequest, opts ...http.CallOption) (rsp *InstrumentStatsReply, err error)
	Strategies(ctx context.Context, req *InstrumentStrategiesRequest, opts ...http.CallOption) (rsp *v1.StrategiesReplies, err error)
}

type InstrumentsHTTPClientImpl struct {
	cc *http.Client
}

func NewInstrumentsHTTPClient(client *http.Client) InstrumentsHTTPClient {
	return &InstrumentsHTTPClientImpl{client}
}

func (c *InstrumentsHTTPClientImpl) Get(ctx context.Context, in *InstrumentRequest, opts ...http.CallOption) (*InstrumentReply, error) {
	var out InstrumentReply
	pattern := "/v1/instrument/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/instruments.v1.instruments/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *InstrumentsHTTPClientImpl) Search(ctx context.Context, in *InstrumentSearchRequest, opts ...http.CallOption) (*InstrumentReplies, error) {
	var out InstrumentReplies
	pattern := "/v1/instruments/search"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/instruments.v1.instruments/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *InstrumentsHTTPClientImpl) Stats(ctx context.Context, in *InstrumentRequest, opts ...http.CallOption) (*InstrumentStatsReply, error) {
	var out InstrumentStatsReply
	pattern := "/v1/instrument/{id}/stats"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/instruments.v1.instruments/Stats"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *InstrumentsHTTPClientImpl) Strategies(ctx context.Context, in *InstrumentStrategiesRequest, opts ...http.CallOption) (*v1.StrategiesReplies, error) {
	var out v1.StrategiesReplies
	pattern := "/v1/instrument/{id}/strategies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/instruments.v1.instruments/Strategies"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
