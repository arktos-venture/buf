// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Dividends

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

type DividendsHTTPServer interface {
	Delete(context.Context, *DividendDeleteRequest) (*DividendDelete, error)
	Last(context.Context, *DividendLastRequest) (*DividendLastReply, error)
	Search(context.Context, *DividendRequest) (*DividendReply, error)
}

func RegisterDividendsHTTPServer(s *http.Server, srv DividendsHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/dividend/{exchange}/{ticker}", _Dividends_Last2_HTTP_Handler(srv))
	r.POST("/v1/dividends", _Dividends_Search9_HTTP_Handler(srv))
	r.DELETE("/v1/dividends", _Dividends_Delete10_HTTP_Handler(srv))
}

func _Dividends_Last2_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DividendLastRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Last")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Last(ctx, req.(*DividendLastRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DividendLastReply)
		return ctx.Result(200, reply)
	}
}

func _Dividends_Search9_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DividendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*DividendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DividendReply)
		return ctx.Result(200, reply)
	}
}

func _Dividends_Delete10_HTTP_Handler(srv DividendsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DividendDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/dividends.v1.Dividends/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DividendDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DividendDelete)
		return ctx.Result(200, reply)
	}
}

type DividendsHTTPClient interface {
	Delete(ctx context.Context, req *DividendDeleteRequest, opts ...http.CallOption) (rsp *DividendDelete, err error)
	Last(ctx context.Context, req *DividendLastRequest, opts ...http.CallOption) (rsp *DividendLastReply, err error)
	Search(ctx context.Context, req *DividendRequest, opts ...http.CallOption) (rsp *DividendReply, err error)
}

type DividendsHTTPClientImpl struct {
	cc *http.Client
}

func NewDividendsHTTPClient(client *http.Client) DividendsHTTPClient {
	return &DividendsHTTPClientImpl{client}
}

func (c *DividendsHTTPClientImpl) Delete(ctx context.Context, in *DividendDeleteRequest, opts ...http.CallOption) (*DividendDelete, error) {
	var out DividendDelete
	pattern := "/v1/dividends"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DividendsHTTPClientImpl) Last(ctx context.Context, in *DividendLastRequest, opts ...http.CallOption) (*DividendLastReply, error) {
	var out DividendLastReply
	pattern := "/v1/dividend/{exchange}/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Last"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DividendsHTTPClientImpl) Search(ctx context.Context, in *DividendRequest, opts ...http.CallOption) (*DividendReply, error) {
	var out DividendReply
	pattern := "/v1/dividends"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/dividends.v1.Dividends/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
