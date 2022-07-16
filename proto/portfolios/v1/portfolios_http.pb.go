// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Portfolios

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

type PortfoliosHTTPServer interface {
	Create(context.Context, *PortfolioCreateRequest) (*PortfolioReply, error)
	Delete(context.Context, *PortfolioDeleteRequest) (*PortfolioDelete, error)
	Search(context.Context, *PortfolioSearchRequest) (*PortfolioReplies, error)
	Status(context.Context, *PortfolioStatusRequest) (*PortfolioReply, error)
	Update(context.Context, *PortfolioUpdateRequest) (*PortfolioReply, error)
}

func RegisterPortfoliosHTTPServer(s *http.Server, srv PortfoliosHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/portfolios/{account}/{currency}", _Portfolios_Status1_HTTP_Handler(srv))
	r.GET("/v1/portfolios/{account}", _Portfolios_Search6_HTTP_Handler(srv))
	r.POST("/v1/portfolios/{account}", _Portfolios_Create6_HTTP_Handler(srv))
	r.PUT("/v1/portfolios/{account}/{currency}", _Portfolios_Update6_HTTP_Handler(srv))
	r.DELETE("/v1/portfolios/{account}/{currency}", _Portfolios_Delete7_HTTP_Handler(srv))
}

func _Portfolios_Status1_HTTP_Handler(srv PortfoliosHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PortfolioStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/portfolios.v1.Portfolios/Status")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Status(ctx, req.(*PortfolioStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PortfolioReply)
		return ctx.Result(200, reply)
	}
}

func _Portfolios_Search6_HTTP_Handler(srv PortfoliosHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PortfolioSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/portfolios.v1.Portfolios/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*PortfolioSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PortfolioReplies)
		return ctx.Result(200, reply)
	}
}

func _Portfolios_Create6_HTTP_Handler(srv PortfoliosHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PortfolioCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/portfolios.v1.Portfolios/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*PortfolioCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PortfolioReply)
		return ctx.Result(200, reply)
	}
}

func _Portfolios_Update6_HTTP_Handler(srv PortfoliosHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PortfolioUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/portfolios.v1.Portfolios/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*PortfolioUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PortfolioReply)
		return ctx.Result(200, reply)
	}
}

func _Portfolios_Delete7_HTTP_Handler(srv PortfoliosHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PortfolioDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/portfolios.v1.Portfolios/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*PortfolioDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PortfolioDelete)
		return ctx.Result(200, reply)
	}
}

type PortfoliosHTTPClient interface {
	Create(ctx context.Context, req *PortfolioCreateRequest, opts ...http.CallOption) (rsp *PortfolioReply, err error)
	Delete(ctx context.Context, req *PortfolioDeleteRequest, opts ...http.CallOption) (rsp *PortfolioDelete, err error)
	Search(ctx context.Context, req *PortfolioSearchRequest, opts ...http.CallOption) (rsp *PortfolioReplies, err error)
	Status(ctx context.Context, req *PortfolioStatusRequest, opts ...http.CallOption) (rsp *PortfolioReply, err error)
	Update(ctx context.Context, req *PortfolioUpdateRequest, opts ...http.CallOption) (rsp *PortfolioReply, err error)
}

type PortfoliosHTTPClientImpl struct {
	cc *http.Client
}

func NewPortfoliosHTTPClient(client *http.Client) PortfoliosHTTPClient {
	return &PortfoliosHTTPClientImpl{client}
}

func (c *PortfoliosHTTPClientImpl) Create(ctx context.Context, in *PortfolioCreateRequest, opts ...http.CallOption) (*PortfolioReply, error) {
	var out PortfolioReply
	pattern := "/v1/portfolios/{account}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/portfolios.v1.Portfolios/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortfoliosHTTPClientImpl) Delete(ctx context.Context, in *PortfolioDeleteRequest, opts ...http.CallOption) (*PortfolioDelete, error) {
	var out PortfolioDelete
	pattern := "/v1/portfolios/{account}/{currency}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/portfolios.v1.Portfolios/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortfoliosHTTPClientImpl) Search(ctx context.Context, in *PortfolioSearchRequest, opts ...http.CallOption) (*PortfolioReplies, error) {
	var out PortfolioReplies
	pattern := "/v1/portfolios/{account}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/portfolios.v1.Portfolios/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortfoliosHTTPClientImpl) Status(ctx context.Context, in *PortfolioStatusRequest, opts ...http.CallOption) (*PortfolioReply, error) {
	var out PortfolioReply
	pattern := "/v1/portfolios/{account}/{currency}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/portfolios.v1.Portfolios/Status"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortfoliosHTTPClientImpl) Update(ctx context.Context, in *PortfolioUpdateRequest, opts ...http.CallOption) (*PortfolioReply, error) {
	var out PortfolioReply
	pattern := "/v1/portfolios/{account}/{currency}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/portfolios.v1.Portfolios/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
