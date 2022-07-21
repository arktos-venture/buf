// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Forexes

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

type ForexesHTTPServer interface {
	Get(context.Context, *ForexRequest) (*ForexReply, error)
	List(context.Context, *ForexListRequest) (*ForexReplies, error)
	Stats(context.Context, *ForexRequest) (*ForexStatsReply, error)
	Strategies(context.Context, *ForexStrategiesRequest) (*v1.StrategiesReplies, error)
}

func RegisterForexesHTTPServer(s *http.Server, srv ForexesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/forexes/{ticker}", _Forexes_Get8_HTTP_Handler(srv))
	r.GET("/v1/forexes/{ticker}/stats", _Forexes_Stats3_HTTP_Handler(srv))
	r.GET("/v1/forexes/{ticker}/strategies", _Forexes_Strategies2_HTTP_Handler(srv))
	r.GET("/v1/forexes/{currency}/pairs", _Forexes_List3_HTTP_Handler(srv))
}

func _Forexes_Get8_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
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

func _Forexes_Stats3_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Stats")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Stats(ctx, req.(*ForexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForexStatsReply)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Strategies2_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexStrategiesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Strategies")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Strategies(ctx, req.(*ForexStrategiesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.StrategiesReplies)
		return ctx.Result(200, reply)
	}
}

func _Forexes_List3_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
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
		reply := out.(*ForexReplies)
		return ctx.Result(200, reply)
	}
}

type ForexesHTTPClient interface {
	Get(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *ForexReply, err error)
	List(ctx context.Context, req *ForexListRequest, opts ...http.CallOption) (rsp *ForexReplies, err error)
	Stats(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *ForexStatsReply, err error)
	Strategies(ctx context.Context, req *ForexStrategiesRequest, opts ...http.CallOption) (rsp *v1.StrategiesReplies, err error)
}

type ForexesHTTPClientImpl struct {
	cc *http.Client
}

func NewForexesHTTPClient(client *http.Client) ForexesHTTPClient {
	return &ForexesHTTPClientImpl{client}
}

func (c *ForexesHTTPClientImpl) Get(ctx context.Context, in *ForexRequest, opts ...http.CallOption) (*ForexReply, error) {
	var out ForexReply
	pattern := "/v1/forexes/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) List(ctx context.Context, in *ForexListRequest, opts ...http.CallOption) (*ForexReplies, error) {
	var out ForexReplies
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

func (c *ForexesHTTPClientImpl) Stats(ctx context.Context, in *ForexRequest, opts ...http.CallOption) (*ForexStatsReply, error) {
	var out ForexStatsReply
	pattern := "/v1/forexes/{ticker}/stats"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Stats"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) Strategies(ctx context.Context, in *ForexStrategiesRequest, opts ...http.CallOption) (*v1.StrategiesReplies, error) {
	var out v1.StrategiesReplies
	pattern := "/v1/forexes/{ticker}/strategies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Strategies"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
