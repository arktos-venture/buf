// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package news_v1

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

type NewsHTTPServer interface {
	Delete(context.Context, *NewsDeleteRequest) (*NewsDelete, error)
	Search(context.Context, *NewsRequest) (*NewsReplies, error)
}

func RegisterNewsHTTPServer(s *http.Server, srv NewsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/news", _News_Search3_HTTP_Handler(srv))
	r.DELETE("/v1/news", _News_Delete6_HTTP_Handler(srv))
}

func _News_Search3_HTTP_Handler(srv NewsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NewsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/news.v1.News/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*NewsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewsReplies)
		return ctx.Result(200, reply)
	}
}

func _News_Delete6_HTTP_Handler(srv NewsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NewsDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/news.v1.News/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*NewsDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewsDelete)
		return ctx.Result(200, reply)
	}
}

type NewsHTTPClient interface {
	Delete(ctx context.Context, req *NewsDeleteRequest, opts ...http.CallOption) (rsp *NewsDelete, err error)
	Search(ctx context.Context, req *NewsRequest, opts ...http.CallOption) (rsp *NewsReplies, err error)
}

type NewsHTTPClientImpl struct {
	cc *http.Client
}

func NewNewsHTTPClient(client *http.Client) NewsHTTPClient {
	return &NewsHTTPClientImpl{client}
}

func (c *NewsHTTPClientImpl) Delete(ctx context.Context, in *NewsDeleteRequest, opts ...http.CallOption) (*NewsDelete, error) {
	var out NewsDelete
	pattern := "/v1/news"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/news.v1.News/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NewsHTTPClientImpl) Search(ctx context.Context, in *NewsRequest, opts ...http.CallOption) (*NewsReplies, error) {
	var out NewsReplies
	pattern := "/v1/news"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/news.v1.News/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
