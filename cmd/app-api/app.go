package main

import (
	"context"

	"github.com/huseinnashr/open-scanner-be/internal/config"
	"github.com/huseinnashr/open-scanner-be/internal/handler/api"
	"github.com/huseinnashr/open-scanner-be/internal/usecase"
)

func startApp(ctx context.Context, config *config.Config) error {
	openCVUsecase := usecase.NewOpenCVUsecase()
	openCVApiHandler := api.NewOpenCVHandler(openCVUsecase)

	if err := startServer(ctx, config, openCVApiHandler); err != nil {
		return err
	}

	return nil
}
