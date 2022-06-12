// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package companies_v1

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

type CompaniesHTTPServer interface {
	Create(context.Context, *CompanyCreateRequest) (*CompanyReply, error)
	Delete(context.Context, *CompanyDeleteRequest) (*CompanyDeleteReply, error)
	Get(context.Context, *CompanyRequest) (*CompanyReply, error)
	Update(context.Context, *CompanyUpdateRequest) (*CompanyReply, error)
}

func RegisterCompaniesHTTPServer(s *http.Server, srv CompaniesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/company/{exchange}/{ticker}", _Companies_Get8_HTTP_Handler(srv))
	r.POST("/v1/companies", _Companies_Create4_HTTP_Handler(srv))
	r.PUT("/v1/company/{exchange}/{ticker}", _Companies_Update3_HTTP_Handler(srv))
	r.DELETE("/v1/companies", _Companies_Delete12_HTTP_Handler(srv))
}

func _Companies_Get8_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*CompanyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Companies_Create4_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CompanyCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Companies_Update3_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*CompanyUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Companies_Delete12_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*CompanyDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyDeleteReply)
		return ctx.Result(200, reply)
	}
}

type CompaniesHTTPClient interface {
	Create(ctx context.Context, req *CompanyCreateRequest, opts ...http.CallOption) (rsp *CompanyReply, err error)
	Delete(ctx context.Context, req *CompanyDeleteRequest, opts ...http.CallOption) (rsp *CompanyDeleteReply, err error)
	Get(ctx context.Context, req *CompanyRequest, opts ...http.CallOption) (rsp *CompanyReply, err error)
	Update(ctx context.Context, req *CompanyUpdateRequest, opts ...http.CallOption) (rsp *CompanyReply, err error)
}

type CompaniesHTTPClientImpl struct {
	cc *http.Client
}

func NewCompaniesHTTPClient(client *http.Client) CompaniesHTTPClient {
	return &CompaniesHTTPClientImpl{client}
}

func (c *CompaniesHTTPClientImpl) Create(ctx context.Context, in *CompanyCreateRequest, opts ...http.CallOption) (*CompanyReply, error) {
	var out CompanyReply
	pattern := "/v1/companies"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/companies.v1.Companies/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompaniesHTTPClientImpl) Delete(ctx context.Context, in *CompanyDeleteRequest, opts ...http.CallOption) (*CompanyDeleteReply, error) {
	var out CompanyDeleteReply
	pattern := "/v1/companies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/companies.v1.Companies/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompaniesHTTPClientImpl) Get(ctx context.Context, in *CompanyRequest, opts ...http.CallOption) (*CompanyReply, error) {
	var out CompanyReply
	pattern := "/v1/company/{exchange}/{ticker}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/companies.v1.Companies/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompaniesHTTPClientImpl) Update(ctx context.Context, in *CompanyUpdateRequest, opts ...http.CallOption) (*CompanyReply, error) {
	var out CompanyReply
	pattern := "/v1/company/{exchange}/{ticker}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/companies.v1.Companies/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
