package api

import (
	"context"

	v1 "github.com/huseinnashr/open-scanner-be/api/v1"
	"github.com/huseinnashr/open-scanner-be/internal/domain"
	"github.com/huseinnashr/open-scanner-be/internal/pkg/tracer"
)

type OpenCVHandler struct {
	v1.UnimplementedOpenCVServiceServer
	accountUsecase domain.OpenCVUsecase
}

func NewOpenCVHandler(accountUsecase domain.OpenCVUsecase) v1.OpenCVServiceServer {
	return &OpenCVHandler{
		accountUsecase: accountUsecase,
	}
}

func (h *OpenCVHandler) WarpPerspective(ctx context.Context, req *v1.WarpPerspectiveRequest) (*v1.Image, error) {
	ctx, span := tracer.Start(ctx, "handler.Signup")
	defer span.End()

	image, err := h.accountUsecase.WarpPerspective(ctx, req)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
