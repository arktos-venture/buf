// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package indexes_v1

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

type IndexesHTTPServer interface {
	Create(context.Context, *IndexesCreateRequest) (*IndexesReply, error)
	Delete(context.Context, *IndexesRequest) (*IndexesReply, error)
	Get(context.Context, *IndexesRequest) (*IndexesReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Search(context.Context, *IndexesSearchRequest) (*IndexesSearchReplies, error)
	Update(context.Context, *IndexesCreateRequest) (*IndexesReply, error)
}

func RegisterIndexesHTTPServer(s *http.Server, srv IndexesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/indexes/{ref}", _Indexes_Get7_HTTP_Handler(srv))
	r.GET("/v1/indexes", _Indexes_Search9_HTTP_Handler(srv))
	r.POST("/v1/indexes", _Indexes_Create3_HTTP_Handler(srv))
	r.PUT("/v1/indexes", _Indexes_Update2_HTTP_Handler(srv))
	r.DELETE("/v1/indexes/{ref}", _Indexes_Delete1_HTTP_Handler(srv))
	r.GET("/healthz", _Indexes_Health20_HTTP_Handler(srv))
}

func _Indexes_Get7_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*IndexesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexesReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Search9_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexesSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*IndexesSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexesSearchReplies)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Create3_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexesCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*IndexesCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexesReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Update2_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexesCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*IndexesCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexesReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Delete1_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndexesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*IndexesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndexesReply)
		return ctx.Result(200, reply)
	}
}

func _Indexes_Health20_HTTP_Handler(srv IndexesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indexes.v1.Indexes/Health")
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

type IndexesHTTPClient interface {
	Create(ctx context.Context, req *IndexesCreateRequest, opts ...http.CallOption) (rsp *IndexesReply, err error)
	Delete(ctx context.Context, req *IndexesRequest, opts ...http.CallOption) (rsp *IndexesReply, err error)
	Get(ctx context.Context, req *IndexesRequest, opts ...http.CallOption) (rsp *IndexesReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Search(ctx context.Context, req *IndexesSearchRequest, opts ...http.CallOption) (rsp *IndexesSearchReplies, err error)
	Update(ctx context.Context, req *IndexesCreateRequest, opts ...http.CallOption) (rsp *IndexesReply, err error)
}

type IndexesHTTPClientImpl struct {
	cc *http.Client
}

func NewIndexesHTTPClient(client *http.Client) IndexesHTTPClient {
	return &IndexesHTTPClientImpl{client}
}

func (c *IndexesHTTPClientImpl) Create(ctx context.Context, in *IndexesCreateRequest, opts ...http.CallOption) (*IndexesReply, error) {
	var out IndexesReply
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

func (c *IndexesHTTPClientImpl) Delete(ctx context.Context, in *IndexesRequest, opts ...http.CallOption) (*IndexesReply, error) {
	var out IndexesReply
	pattern := "/v1/indexes/{ref}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Get(ctx context.Context, in *IndexesRequest, opts ...http.CallOption) (*IndexesReply, error) {
	var out IndexesReply
	pattern := "/v1/indexes/{ref}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/indexes.v1.Indexes/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndexesHTTPClientImpl) Search(ctx context.Context, in *IndexesSearchRequest, opts ...http.CallOption) (*IndexesSearchReplies, error) {
	var out IndexesSearchReplies
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

func (c *IndexesHTTPClientImpl) Update(ctx context.Context, in *IndexesCreateRequest, opts ...http.CallOption) (*IndexesReply, error) {
	var out IndexesReply
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
