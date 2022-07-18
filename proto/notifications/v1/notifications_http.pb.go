// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Notifications

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

type NotificationsHTTPServer interface {
	Create(context.Context, *NotificationCreateRequest) (*NotificationReply, error)
	Delete(context.Context, *NotificationDeleteRequest) (*NotificationDelete, error)
	Search(context.Context, *NotificationSearchRequest) (*NotificationReplies, error)
}

func RegisterNotificationsHTTPServer(s *http.Server, srv NotificationsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/notifications/{account}", _Notifications_Create9_HTTP_Handler(srv))
	r.GET("/v1/notifications/{account}", _Notifications_Search11_HTTP_Handler(srv))
	r.DELETE("/v1/notifications", _Notifications_Delete9_HTTP_Handler(srv))
}

func _Notifications_Create9_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
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

func _Notifications_Search11_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
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

func _Notifications_Delete9_HTTP_Handler(srv NotificationsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotificationDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/notifications.v1.Notifications/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*NotificationDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NotificationDelete)
		return ctx.Result(200, reply)
	}
}

type NotificationsHTTPClient interface {
	Create(ctx context.Context, req *NotificationCreateRequest, opts ...http.CallOption) (rsp *NotificationReply, err error)
	Delete(ctx context.Context, req *NotificationDeleteRequest, opts ...http.CallOption) (rsp *NotificationDelete, err error)
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
	pattern := "/v1/notifications/{account}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationsHTTPClientImpl) Delete(ctx context.Context, in *NotificationDeleteRequest, opts ...http.CallOption) (*NotificationDelete, error) {
	var out NotificationDelete
	pattern := "/v1/notifications"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationsHTTPClientImpl) Search(ctx context.Context, in *NotificationSearchRequest, opts ...http.CallOption) (*NotificationReplies, error) {
	var out NotificationReplies
	pattern := "/v1/notifications/{account}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/notifications.v1.Notifications/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
