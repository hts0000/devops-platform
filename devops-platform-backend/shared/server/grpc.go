package server

import (
	"context"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(*grpc.Server)
	Logger       *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameFiled := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameFiled, zap.Error(err))
	}

	var opts []grpc.ServerOption
	in := grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// TODO: jwt token parse
		return handler(ctx, req)
	})
	opts = append(opts, in)
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)

	c.Logger.Info("server started", nameFiled, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
