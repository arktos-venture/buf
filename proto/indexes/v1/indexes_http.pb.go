// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package indexes_v1

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

type IndexesHTTPServer interface {
	Create(context.Context, *IndexCreateRequest) (*IndexReply, error)
	Delete(context.Context, *IndexDeleteRequest) (*IndexDelete, error)
	Get(context.Context, *IndexRequest) (*IndexReply, error)
	Search(context.Context, *IndexSearchRequest) (*IndexSearchReplies, error)
	Update(context.Context, *IndexCreateRequest) (*IndexReply, error)
}

func RegisterIndexesHTTPServer(s *http.Server, srv IndexesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/indexes/{ticker}", _Indexes_Get9_HTTP_Handler(srv))
	r.GET("/v1/indexes", _Indexes_Search10_HTTP_Handler(srv))
	r.POST("/v1/indexes", _Indexes_Create8_HTTP_Handler(srv))
	r.PUT("/v1/indexes", _Indexes_Update5_HTTP_Handler(srv))
	r.DELETE("/v1/indexes", _Indexes_Delete13_HTTP_Handler(srv))
}

func _Indexes_Get9_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*IndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Search10_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*IndexSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexSearchReplies)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Create8_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*IndexCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Update5_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*IndexCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Delete13_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*IndexDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexDelete)
		return ctx.Result(200, reply)
	}
}

type IndexesHTTPClient interface {
	Create(ctx context.Context, req *IndexCreateRequest, opts ...http.CallOption) (rsp *IndexReply, err error)
	Delete(ctx context.Context, req *IndexDeleteRequest, opts ...http.CallOption) (rsp *IndexDelete, err error)
	Get(ctx context.Context, req *IndexRequest, opts ...http.CallOption) (rsp *IndexReply, err error)
	Search(ctx context.Context, req *IndexSearchRequest, opts ...http.CallOption) (rsp *IndexSearchReplies, err error)
	Update(ctx context.Context, req *IndexCreateRequest, opts ...http.CallOption) (rsp *IndexReply, err error)
}

type IndexesHTTPClientImpl struct {
	cc *http.Client
}

func NewIndexesHTTPClient(client *http.Client) IndexesHTTPClient {
	return &IndexesHTTPClientImpl{client}
}

func (c *IndexesHTTPClientImpl) Create(ctx context.Context, in *IndexCreateRequest, opts ...http.CallOption) (*IndexReply, error) {
	var out IndexReply
	pattern := "/v1/indexes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Delete(ctx context.Context, in *IndexDeleteRequest, opts ...http.CallOption) (*IndexDelete, error) {
	var out IndexDelete
	pattern := "/v1/indexes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Get(ctx context.Context, in *IndexRequest, opts ...http.CallOption) (*IndexReply, error) {
	var out IndexReply
	pattern := "/v1/indexes/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Search(ctx context.Context, in *IndexSearchRequest, opts ...http.CallOption) (*IndexSearchReplies, error) {
	var out IndexSearchReplies
	pattern := "/v1/indexes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Update(ctx context.Context, in *IndexCreateRequest, opts ...http.CallOption) (*IndexReply, error) {
	var out IndexReply
	pattern := "/v1/indexes"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
