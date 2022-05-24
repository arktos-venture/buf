// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package strategies_v1

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
	Execute(context.Context, *StrategyRequest) (*StrategyReply, error)
}

func RegisterStrategiesHTTPServer(s *http.Server, srv StrategiesHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/strategy/execute", _Strategies_Execute1_HTTP_Handler(srv))
}

func _Strategies_Execute1_HTTP_Handler(srv StrategiesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StrategyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/strategies.v1.Strategies/Execute")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Execute(ctx, req.(*StrategyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StrategyReply)
		return ctx.Result(200, reply)
	}
}

type StrategiesHTTPClient interface {
	Execute(ctx context.Context, req *StrategyRequest, opts ...http.CallOption) (rsp *StrategyReply, err error)
}

type StrategiesHTTPClientImpl struct {
	cc *http.Client
}

func NewStrategiesHTTPClient(client *http.Client) StrategiesHTTPClient {
	return &StrategiesHTTPClientImpl{client}
}

func (c *StrategiesHTTPClientImpl) Execute(ctx context.Context, in *StrategyRequest, opts ...http.CallOption) (*StrategyReply, error) {
	var out StrategyReply
	pattern := "/v1/strategy/execute"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/strategies.v1.Strategies/Execute"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
