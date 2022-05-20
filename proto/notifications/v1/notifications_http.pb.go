// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package notifications_v1

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

type NotificationsHTTPServer interface {
	Create(context.Context, *NotificationCreateRequest) (*NotificationReply, error)
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Search(context.Context, *NotificationSearchRequest) (*NotificationReplies, error)
}

func RegisterNotificationsHTTPServer(s *http.Server, srv NotificationsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/{account}/notifications", _Notifications_Create2_HTTP_Handler(srv))
	r.GET("/v1/{account}/notifications", _Notifications_Search7_HTTP_Handler(srv))
	r.GET("/healthz", _Notifications_Health17_HTTP_Handler(srv))
}

func _Notifications_Create2_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotificationCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/notifications.v1.Notifications/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*NotificationCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotificationReply)
		return ctx.Result(200, reply)
	}
}

func _Notifications_Search7_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotificationSearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/notifications.v1.Notifications/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*NotificationSearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotificationReplies)
		return ctx.Result(200, reply)
	}
}

func _Notifications_Health17_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/notifications.v1.Notifications/Health")
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

type NotificationsHTTPClient interface {
	Create(ctx context.Context, req *NotificationCreateRequest, opts ...http.CallOption) (rsp *NotificationReply, err error)
	Health(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Search(ctx context.Context, req *NotificationSearchRequest, opts ...http.CallOption) (rsp *NotificationReplies, err error)
}

type NotificationsHTTPClientImpl struct {
	cc *http.Client
}

func NewNotificationsHTTPClient(client *http.Client) NotificationsHTTPClient {
	return &NotificationsHTTPClientImpl{client}
}

func (c *NotificationsHTTPClientImpl) Create(ctx context.Context, in *NotificationCreateRequest, opts ...http.CallOption) (*NotificationReply, error) {
	var out NotificationReply
	pattern := "/v1/{account}/notifications"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationsHTTPClientImpl) Health(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/healthz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Health"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationsHTTPClientImpl) Search(ctx context.Context, in *NotificationSearchRequest, opts ...http.CallOption) (*NotificationReplies, error) {
	var out NotificationReplies
	pattern := "/v1/{account}/notifications"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
