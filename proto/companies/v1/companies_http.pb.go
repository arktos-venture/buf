// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package companies_v1

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

type CompaniesHTTPServer interface {
	Get(context.Context, *CompanyRequest) (*CompanyReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Similars(context.Context, *CompanyRequest) (*CompanySimilarsReply, error)
	Stats(context.Context, *CompanyRequest) (*CompanyStatsReply, error)
}

func RegisterCompaniesHTTPServer(s *http.Server, srv CompaniesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/company/{exchange}/{ticker}", _Companies_Get6_HTTP_Handler(srv))
	r.GET("/v1/company/{exchange}/{ticker}/stats", _Companies_Stats0_HTTP_Handler(srv))
	r.GET("/v1/company/{exchange}/{ticker}/similars", _Companies_Similars0_HTTP_Handler(srv))
	r.GET("/healthz", _Companies_Health17_HTTP_Handler(srv))
}

func _Companies_Get6_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
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

func _Companies_Stats0_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Stats")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Stats(ctx, req.(*CompanyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyStatsReply)
		return ctx.Result(200, reply)
	}
}

func _Companies_Similars0_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Similars")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Similars(ctx, req.(*CompanyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanySimilarsReply)
		return ctx.Result(200, reply)
	}
}

func _Companies_Health17_HTTP_Handler(srv CompaniesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/companies.v1.Companies/Health")
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

type CompaniesHTTPClient interface {
	Get(ctx context.Context, req *CompanyRequest, opts ...http.CallOption) (rsp *CompanyReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Similars(ctx context.Context, req *CompanyRequest, opts ...http.CallOption) (rsp *CompanySimilarsReply, err error)
	Stats(ctx context.Context, req *CompanyRequest, opts ...http.CallOption) (rsp *CompanyStatsReply, err error)
}

type CompaniesHTTPClientImpl struct {
	cc *http.Client
}

func NewCompaniesHTTPClient(client *http.Client) CompaniesHTTPClient {
	return &CompaniesHTTPClientImpl{client}
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

func (c *CompaniesHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/companies.v1.Companies/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompaniesHTTPClientImpl) Similars(ctx context.Context, in *CompanyRequest, opts ...http.CallOption) (*CompanySimilarsReply, error) {
	var out CompanySimilarsReply
	pattern := "/v1/company/{exchange}/{ticker}/similars"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/companies.v1.Companies/Similars"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompaniesHTTPClientImpl) Stats(ctx context.Context, in *CompanyRequest, opts ...http.CallOption) (*CompanyStatsReply, error) {
	var out CompanyStatsReply
	pattern := "/v1/company/{exchange}/{ticker}/stats"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/companies.v1.Companies/Stats"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
