package service

import (
	"AstrologService/pkg/model"
	"AstrologService/pkg/repository"
	"context"
	"time"
)

type ImageImplementation interface {
	InsertImage(ctx context.Context, picture *model.ImageModel) (int64, error)
	GetImageByDate(ctx context.Context, date time.Time) (*model.ImageModel, error)
	DeleteImageByDate(ctx context.Context, date time.Time) (int64, error)
	GetAllImages(ctx context.Context) ([]*model.ImageModel, error)
}

type Service struct {
	repository repository.Implementation
}

type Implementation interface {
	ImageImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
