package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type Client struct {
	Context context.Context
	GRPC    *grpc.ClientConn
}
