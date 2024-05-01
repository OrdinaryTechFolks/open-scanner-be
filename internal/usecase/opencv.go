package usecase

import (
	"context"
	"image"

	v1 "github.com/huseinnashr/open-scanner-be/api/open-scanner-be/v1"
	"github.com/huseinnashr/open-scanner-be/internal/domain"
	"gocv.io/x/gocv"
)

type OpenCVUsecase struct{}

func NewOpenCVUsecase() domain.OpenCVUsecase {
	return &OpenCVUsecase{}
}

func (u *OpenCVUsecase) WarpPerspective(ctx context.Context, param *v1.WarpPerspectiveRequest) (v1.Image, error) {
	srcImg := param.GetSrcImage()
	srcCorners := param.GetSrcCorners()
	destSize := param.GetDestSize()

	srcMat, err := gocv.NewMatFromBytes(int(srcImg.GetHeight()), int(srcImg.GetWidth()), gocv.MatTypeCV8UC4, srcImg.GetData())
	if err != nil {
		return v1.Image{}, nil
	}
	destMat := gocv.NewMat()

	destCorners := []gocv.Point2f{{X: 0, Y: 0}, {X: destSize.GetWidth(), Y: 0},
		{X: destSize.Width, Y: destSize.Height}, {X: 0, Y: destSize.Width}}
	perspectiveMat := gocv.GetPerspectiveTransform2f(
		gocv.NewPoint2fVectorFromPoints(Point2FToGoCVs(srcCorners)),
		gocv.NewPoint2fVectorFromPoints(destCorners),
	)

	gocv.WarpPerspective(srcMat, &destMat, perspectiveMat, Size2FToImagePoint(destSize))

	return v1.Image{
		Width:    int32(destMat.Cols()),
		Height:   int32(destMat.Rows()),
		Channels: int32(destMat.Channels()),
		Data:     destMat.ToBytes(),
	}, nil
}

func Size2FToImagePoint(src *v1.Size2F) image.Point {
	return image.Point{X: int(src.GetWidth()), Y: int(src.GetHeight())}
}

func Point2FToGoCVs(src []*v1.Point2F) []gocv.Point2f {
	dest := make([]gocv.Point2f, len(src))
	for i, point := range src {
		dest[i] = Point2FToGoCV(point)
	}
	return dest
}

func Point2FToGoCV(src *v1.Point2F) gocv.Point2f {
	return gocv.Point2f{X: src.X, Y: src.Y}
}
