package repository

import (
	"AstrologService/pkg/model"
	"context"
	"time"
)

type ImageImplementation interface {
	InsertImage(ctx context.Context, picture *model.ImageModel) (int64, error)
	GetImageByDate(ctx context.Context, date time.Time) (*model.ImageModel, error)
	DeleteImageByDate(ctx context.Context, date time.Time) (int64, error)
	GetAllImages(ctx context.Context) ([]*model.ImageModel, error)
}

type Implementation interface {
	ImageImplementation
}
