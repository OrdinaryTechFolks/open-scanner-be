package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/huseinnashr/open-scanner-be/api/open-scanner-be/v1"
	"github.com/huseinnashr/open-scanner-be/cmd/app-api/middleware"
	"github.com/huseinnashr/open-scanner-be/internal/config"
	"golang.org/x/sync/errgroup"
)

type IServer interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func startServer(ctx context.Context, config *config.Config, openCVAPIHandler v1.OpenCVServiceServer) error {
	var servers []IServer

	var opts []grpc.ServerOption
	if address := config.App.GRPC.Address; address != "" {
		opts = append(opts, grpc.Address(address))
	}
	if timeout := config.App.GRPC.Timeout; timeout != 0 {
		opts = append(opts, grpc.Timeout(timeout))
	}
	opts = append(opts, grpc.Middleware(
		middleware.ServerMetadata(config.App.Version),
	))
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterOpenCVServiceServer(grpcServer, openCVAPIHandler)
	servers = append(servers, grpcServer)

	group, ctx := errgroup.WithContext(ctx)
	for _, server := range servers {
		server := server
		group.Go(func() error {
			if err := server.Start(ctx); err != nil {
				return err
			}
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
