package domain

import (
	"context"

	v1 "github.com/huseinnashr/open-scanner-be/api/v1"
)

type OpenCVUsecase interface {
	WarpPerspective(ctx context.Context, param *v1.WarpPerspectiveRequest) (v1.Image, error)
}
