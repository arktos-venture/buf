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

type OrdersHTTPServer interface {
	Create(context.Context, *OrderCreateRequest) (*OrderReply, error)
	Delete(context.Context, *OrderDeleteRequest) (*OrderDelete, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Search(context.Context, *OrderSearchRequest) (*OrderReplies, error)
	Update(context.Context, *OrderUpdateRequest) (*OrderReply, error)
}

func RegisterOrdersHTTPServer(s *http.Server, srv OrdersHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/{account}/orders", _Orders_Search4_HTTP_Handler(srv))
	r.POST("/v1/{account}/order", _Orders_Create1_HTTP_Handler(srv))
	r.PUT("/v1/{account}/order/{orderId}", _Orders_Update0_HTTP_Handler(srv))
	r.DELETE("/v1/{account}/order/{orderId}", _Orders_Delete0_HTTP_Handler(srv))
	r.GET("/healthz", _Orders_Health10_HTTP_Handler(srv))
}

func _Orders_Search4_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*OrderSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderReplies)
		return ctx.Result(200, reply)
	}
}

func _Orders_Create1_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*OrderCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderReply)
		return ctx.Result(200, reply)
	}
}

func _Orders_Update0_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*OrderUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderReply)
		return ctx.Result(200, reply)
	}
}

func _Orders_Delete0_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*OrderDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderDelete)
		return ctx.Result(200, reply)
	}
}

func _Orders_Health10_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Health")
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

type OrdersHTTPClient interface {
	Create(ctx context.Context, req *OrderCreateRequest, opts ...http.CallOption) (rsp *OrderReply, err error)
	Delete(ctx context.Context, req *OrderDeleteRequest, opts ...http.CallOption) (rsp *OrderDelete, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Search(ctx context.Context, req *OrderSearchRequest, opts ...http.CallOption) (rsp *OrderReplies, err error)
	Update(ctx context.Context, req *OrderUpdateRequest, opts ...http.CallOption) (rsp *OrderReply, err error)
}

type OrdersHTTPClientImpl struct {
	cc *http.Client
}

func NewOrdersHTTPClient(client *http.Client) OrdersHTTPClient {
	return &OrdersHTTPClientImpl{client}
}

func (c *OrdersHTTPClientImpl) Create(ctx context.Context, in *OrderCreateRequest, opts ...http.CallOption) (*OrderReply, error) {
	var out OrderReply
	pattern := "/v1/{account}/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Delete(ctx context.Context, in *OrderDeleteRequest, opts ...http.CallOption) (*OrderDelete, error) {
	var out OrderDelete
	pattern := "/v1/{account}/order/{orderId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/orders.v1.Orders/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/orders.v1.Orders/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Search(ctx context.Context, in *OrderSearchRequest, opts ...http.CallOption) (*OrderReplies, error) {
	var out OrderReplies
	pattern := "/v1/{account}/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/orders.v1.Orders/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Update(ctx context.Context, in *OrderUpdateRequest, opts ...http.CallOption) (*OrderReply, error) {
	var out OrderReply
	pattern := "/v1/{account}/order/{orderId}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
