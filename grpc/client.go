package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func NewClient(hostname string, timeout time.Duration) *Client {
	var (
		ctx  context.Context
		conn *grpc.ClientConn
		err  error
	)

	if conn, err = grpc.Dial(hostname, grpc.WithInsecure()); err != nil {
		return nil
	}

	ctx, _ = context.WithTimeout(context.Background(), timeout)

	return &Client{
		Context: ctx,
		GRPC:    conn,
	}
}
