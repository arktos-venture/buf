// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Forexes

import (
	context "context"
	v1 "github.com/arktos-venture/buf/proto/quotes/v1"
	v11 "github.com/arktos-venture/buf/proto/strategies/v1"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ForexesHTTPServer interface {
	Create(context.Context, *ForexCreateRequest) (*ForexReply, error)
	Delete(context.Context, *ForexDeleteRequest) (*ForexDelete, error)
	Get(context.Context, *ForexRequest) (*ForexReply, error)
	LastQuote(context.Context, *ForexRequest) (*v1.QuoteReply, error)
	List(context.Context, *ForexListRequest) (*ForexList, error)
	Quotes(context.Context, *ForexQuotesRequest) (*v1.QuoteReplies, error)
	Stats(context.Context, *ForexRequest) (*ForexStatsReply, error)
	Strategies(context.Context, *ForexStrategiesRequest) (*v11.StrategiesReplies, error)
}

func RegisterForexesHTTPServer(s *http.Server, srv ForexesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/forexes/{ticker}", _Forexes_Get7_HTTP_Handler(srv))
	r.GET("/v1/forexes/{ticker}/stats", _Forexes_Stats4_HTTP_Handler(srv))
	r.GET("/v1/forexes/{ticker}/quotes/last", _Forexes_LastQuote2_HTTP_Handler(srv))
	r.POST("/v1/forexes/quotes", _Forexes_Quotes2_HTTP_Handler(srv))
	r.GET("/v1/forexes/{ticker}/strategies", _Forexes_Strategies3_HTTP_Handler(srv))
	r.GET("/v1/forexes/{currency}/pairs", _Forexes_List3_HTTP_Handler(srv))
	r.POST("/v1/forexes", _Forexes_Create7_HTTP_Handler(srv))
	r.DELETE("/v1/forexes", _Forexes_Delete6_HTTP_Handler(srv))
}

func _Forexes_Get7_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
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

func _Forexes_Stats4_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
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

func _Forexes_LastQuote2_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/LastQuote")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LastQuote(ctx, req.(*ForexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.QuoteReply)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Quotes2_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexQuotesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Quotes")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Quotes(ctx, req.(*ForexQuotesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.QuoteReplies)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Strategies3_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
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
		reply := out.(*v11.StrategiesReplies)
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
		reply := out.(*ForexList)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Create7_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*ForexCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForexReply)
		return ctx.Result(200, reply)
	}
}

func _Forexes_Delete6_HTTP_Handler(srv ForexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForexDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/forexes.v1.Forexes/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*ForexDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForexDelete)
		return ctx.Result(200, reply)
	}
}

type ForexesHTTPClient interface {
	Create(ctx context.Context, req *ForexCreateRequest, opts ...http.CallOption) (rsp *ForexReply, err error)
	Delete(ctx context.Context, req *ForexDeleteRequest, opts ...http.CallOption) (rsp *ForexDelete, err error)
	Get(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *ForexReply, err error)
	LastQuote(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *v1.QuoteReply, err error)
	List(ctx context.Context, req *ForexListRequest, opts ...http.CallOption) (rsp *ForexList, err error)
	Quotes(ctx context.Context, req *ForexQuotesRequest, opts ...http.CallOption) (rsp *v1.QuoteReplies, err error)
	Stats(ctx context.Context, req *ForexRequest, opts ...http.CallOption) (rsp *ForexStatsReply, err error)
	Strategies(ctx context.Context, req *ForexStrategiesRequest, opts ...http.CallOption) (rsp *v11.StrategiesReplies, err error)
}

type ForexesHTTPClientImpl struct {
	cc *http.Client
}

func NewForexesHTTPClient(client *http.Client) ForexesHTTPClient {
	return &ForexesHTTPClientImpl{client}
}

func (c *ForexesHTTPClientImpl) Create(ctx context.Context, in *ForexCreateRequest, opts ...http.CallOption) (*ForexReply, error) {
	var out ForexReply
	pattern := "/v1/forexes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) Delete(ctx context.Context, in *ForexDeleteRequest, opts ...http.CallOption) (*ForexDelete, error) {
	var out ForexDelete
	pattern := "/v1/forexes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
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

func (c *ForexesHTTPClientImpl) LastQuote(ctx context.Context, in *ForexRequest, opts ...http.CallOption) (*v1.QuoteReply, error) {
	var out v1.QuoteReply
	pattern := "/v1/forexes/{ticker}/quotes/last"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/LastQuote"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ForexesHTTPClientImpl) List(ctx context.Context, in *ForexListRequest, opts ...http.CallOption) (*ForexList, error) {
	var out ForexList
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

func (c *ForexesHTTPClientImpl) Quotes(ctx context.Context, in *ForexQuotesRequest, opts ...http.CallOption) (*v1.QuoteReplies, error) {
	var out v1.QuoteReplies
	pattern := "/v1/forexes/quotes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/forexes.v1.Forexes/Quotes"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
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

func (c *ForexesHTTPClientImpl) Strategies(ctx context.Context, in *ForexStrategiesRequest, opts ...http.CallOption) (*v11.StrategiesReplies, error) {
	var out v11.StrategiesReplies
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
