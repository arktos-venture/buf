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

type SplitsHTTPServer interface {
	Bulk(context.Context, *SplitBulkRequest) (*SplitBulkReplies, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	List(context.Context, *SplitPeriodRequest) (*SplitReplies, error)
	ListByTimestamp(context.Context, *SplitTimestampRequest) (*SplitReplies, error)
}

func RegisterSplitsHTTPServer(s *http.Server, srv SplitsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/company/{exchange}/{ticker}/splits", _Splits_List4_HTTP_Handler(srv))
	r.POST("/v1/company/{exchange}/{ticker}/splits/ts", _Splits_ListByTimestamp0_HTTP_Handler(srv))
	r.POST("/v1/exchange/{exchange}/splits/bulk", _Splits_Bulk0_HTTP_Handler(srv))
	r.GET("/healthz", _Splits_Health16_HTTP_Handler(srv))
}

func _Splits_List4_HTTP_Handler(srv SplitsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SplitPeriodRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/splits.v1.Splits/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*SplitPeriodRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SplitReplies)
		return ctx.Result(200, reply)
	}
}

func _Splits_ListByTimestamp0_HTTP_Handler(srv SplitsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SplitTimestampRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/splits.v1.Splits/ListByTimestamp")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListByTimestamp(ctx, req.(*SplitTimestampRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SplitReplies)
		return ctx.Result(200, reply)
	}
}

func _Splits_Bulk0_HTTP_Handler(srv SplitsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SplitBulkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/splits.v1.Splits/Bulk")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Bulk(ctx, req.(*SplitBulkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SplitBulkReplies)
		return ctx.Result(200, reply)
	}
}

func _Splits_Health16_HTTP_Handler(srv SplitsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/splits.v1.Splits/Health")
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

type SplitsHTTPClient interface {
	Bulk(ctx context.Context, req *SplitBulkRequest, opts ...http.CallOption) (rsp *SplitBulkReplies, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	List(ctx context.Context, req *SplitPeriodRequest, opts ...http.CallOption) (rsp *SplitReplies, err error)
	ListByTimestamp(ctx context.Context, req *SplitTimestampRequest, opts ...http.CallOption) (rsp *SplitReplies, err error)
}

type SplitsHTTPClientImpl struct {
	cc *http.Client
}

func NewSplitsHTTPClient(client *http.Client) SplitsHTTPClient {
	return &SplitsHTTPClientImpl{client}
}

func (c *SplitsHTTPClientImpl) Bulk(ctx context.Context, in *SplitBulkRequest, opts ...http.CallOption) (*SplitBulkReplies, error) {
	var out SplitBulkReplies
	pattern := "/v1/exchange/{exchange}/splits/bulk"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/splits.v1.Splits/Bulk"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SplitsHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/splits.v1.Splits/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SplitsHTTPClientImpl) List(ctx context.Context, in *SplitPeriodRequest, opts ...http.CallOption) (*SplitReplies, error) {
	var out SplitReplies
	pattern := "/v1/company/{exchange}/{ticker}/splits"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/splits.v1.Splits/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SplitsHTTPClientImpl) ListByTimestamp(ctx context.Context, in *SplitTimestampRequest, opts ...http.CallOption) (*SplitReplies, error) {
	var out SplitReplies
	pattern := "/v1/company/{exchange}/{ticker}/splits/ts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/splits.v1.Splits/ListByTimestamp"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
