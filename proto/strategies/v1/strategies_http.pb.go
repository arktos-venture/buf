// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Strategies

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

type StrategiesHTTPServer interface {
	Create(context.Context, *StrategyModifyRequest) (*StrategyReply, error)
	Delete(context.Context, *StrategyDeleteRequest) (*StrategyDelete, error)
	Get(context.Context, *StrategyRequest) (*StrategyReply, error)
	List(context.Context, *StrategyListRequest) (*StrategyReplies, error)
	Update(context.Context, *StrategyModifyRequest) (*StrategyReply, error)
}

func RegisterStrategiesHTTPServer(s *http.Server, srv StrategiesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/strategy/{ticker}", _Strategies_Get1_HTTP_Handler(srv))
	r.GET("/v1/strategies", _Strategies_List0_HTTP_Handler(srv))
	r.POST("/v1/strategies", _Strategies_Create0_HTTP_Handler(srv))
	r.PUT("/v1/strategy/{ticker}", _Strategies_Update0_HTTP_Handler(srv))
	r.DELETE("/v1/strategies", _Strategies_Delete0_HTTP_Handler(srv))
}

func _Strategies_Get1_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*StrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategies_List0_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*StrategyListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReplies)
		return ctx.Result(200, reply)
	}
}

func _Strategies_Create0_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyModifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*StrategyModifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategies_Update0_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyModifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*StrategyModifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

func _Strategies_Delete0_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*StrategyDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyDelete)
		return ctx.Result(200, reply)
	}
}

type StrategiesHTTPClient interface {
	Create(ctx context.Context, req *StrategyModifyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	Delete(ctx context.Context, req *StrategyDeleteRequest, opts ...http.CallOption) (rsp *StrategyDelete, err error)
	Get(ctx context.Context, req *StrategyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
	List(ctx context.Context, req *StrategyListRequest, opts ...http.CallOption) (rsp *StrategyReplies, err error)
	Update(ctx context.Context, req *StrategyModifyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
}

type StrategiesHTTPClientImpl struct {
	cc *http.Client
}

func NewStrategiesHTTPClient(client *http.Client) StrategiesHTTPClient {
	return &StrategiesHTTPClientImpl{client}
}

func (c *StrategiesHTTPClientImpl) Create(ctx context.Context, in *StrategyModifyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategies"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategiesHTTPClientImpl) Delete(ctx context.Context, in *StrategyDeleteRequest, opts ...http.CallOption) (*StrategyDelete, error) {
	var out StrategyDelete
	pattern := "/v1/strategies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategiesHTTPClientImpl) Get(ctx context.Context, in *StrategyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategiesHTTPClientImpl) List(ctx context.Context, in *StrategyListRequest, opts ...http.CallOption) (*StrategyReplies, error) {
	var out StrategyReplies
	pattern := "/v1/strategies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StrategiesHTTPClientImpl) Update(ctx context.Context, in *StrategyModifyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/{ticker}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
