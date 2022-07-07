// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Orders

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

type OrdersHTTPServer interface {
	Cancel(context.Context, *OrderCancelRequest) (*OrderCancel, error)
	Create(context.Context, *OrderCreateRequest) (*OrderModifyReply, error)
	Positions(context.Context, *PositionRequest) (*PositionReplies, error)
	Search(context.Context, *OrderSearchRequest) (*OrderReplies, error)
	Status(context.Context, *OrderStatusRequest) (*OrderReply, error)
	Update(context.Context, *OrderUpdateRequest) (*OrderModifyReply, error)
}

func RegisterOrdersHTTPServer(s *http.Server, srv OrdersHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/positions/{account}/{portfolioUUID}", _Orders_Positions0_HTTP_Handler(srv))
	r.GET("/v1/orders/{account}/{portfolioUUID}/{orderUUID}", _Orders_Status0_HTTP_Handler(srv))
	r.POST("/v1/orders/{account}/{portfolioUUID}/search", _Orders_Search4_HTTP_Handler(srv))
	r.POST("/v1/orders/{account}/{portfolioUUID}", _Orders_Create3_HTTP_Handler(srv))
	r.PATCH("/v1/orders/{account}/{portfolioUUID}/{orderUUID}", _Orders_Update2_HTTP_Handler(srv))
	r.DELETE("/v1/orders/{account}/{portfolioUUID}/{orderUUID}/cancel", _Orders_Cancel0_HTTP_Handler(srv))
}

func _Orders_Positions0_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PositionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Positions")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Positions(ctx, req.(*PositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PositionReplies)
		return ctx.Result(200, reply)
	}
}

func _Orders_Status0_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Status")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Status(ctx, req.(*OrderStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderReply)
		return ctx.Result(200, reply)
	}
}

func _Orders_Search4_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderSearchRequest
		if err := ctx.Bind(&in); err != nil {
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

func _Orders_Create3_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
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
		reply := out.(*OrderModifyReply)
		return ctx.Result(200, reply)
	}
}

func _Orders_Update2_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
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
		reply := out.(*OrderModifyReply)
		return ctx.Result(200, reply)
	}
}

func _Orders_Cancel0_HTTP_Handler(srv OrdersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderCancelRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/orders.v1.Orders/Cancel")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Cancel(ctx, req.(*OrderCancelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderCancel)
		return ctx.Result(200, reply)
	}
}

type OrdersHTTPClient interface {
	Cancel(ctx context.Context, req *OrderCancelRequest, opts ...http.CallOption) (rsp *OrderCancel, err error)
	Create(ctx context.Context, req *OrderCreateRequest, opts ...http.CallOption) (rsp *OrderModifyReply, err error)
	Positions(ctx context.Context, req *PositionRequest, opts ...http.CallOption) (rsp *PositionReplies, err error)
	Search(ctx context.Context, req *OrderSearchRequest, opts ...http.CallOption) (rsp *OrderReplies, err error)
	Status(ctx context.Context, req *OrderStatusRequest, opts ...http.CallOption) (rsp *OrderReply, err error)
	Update(ctx context.Context, req *OrderUpdateRequest, opts ...http.CallOption) (rsp *OrderModifyReply, err error)
}

type OrdersHTTPClientImpl struct {
	cc *http.Client
}

func NewOrdersHTTPClient(client *http.Client) OrdersHTTPClient {
	return &OrdersHTTPClientImpl{client}
}

func (c *OrdersHTTPClientImpl) Cancel(ctx context.Context, in *OrderCancelRequest, opts ...http.CallOption) (*OrderCancel, error) {
	var out OrderCancel
	pattern := "/v1/orders/{account}/{portfolioUUID}/{orderUUID}/cancel"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/orders.v1.Orders/Cancel"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Create(ctx context.Context, in *OrderCreateRequest, opts ...http.CallOption) (*OrderModifyReply, error) {
	var out OrderModifyReply
	pattern := "/v1/orders/{account}/{portfolioUUID}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Positions(ctx context.Context, in *PositionRequest, opts ...http.CallOption) (*PositionReplies, error) {
	var out PositionReplies
	pattern := "/v1/positions/{account}/{portfolioUUID}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Positions"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Search(ctx context.Context, in *OrderSearchRequest, opts ...http.CallOption) (*OrderReplies, error) {
	var out OrderReplies
	pattern := "/v1/orders/{account}/{portfolioUUID}/search"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Status(ctx context.Context, in *OrderStatusRequest, opts ...http.CallOption) (*OrderReply, error) {
	var out OrderReply
	pattern := "/v1/orders/{account}/{portfolioUUID}/{orderUUID}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/orders.v1.Orders/Status"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrdersHTTPClientImpl) Update(ctx context.Context, in *OrderUpdateRequest, opts ...http.CallOption) (*OrderModifyReply, error) {
	var out OrderModifyReply
	pattern := "/v1/orders/{account}/{portfolioUUID}/{orderUUID}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/orders.v1.Orders/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
