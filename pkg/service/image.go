package service

import (
	"AstrologService/pkg/model"
	"context"
	"time"
)

func (s *Service) InsertImage(ctx context.Context, image *model.ImageModel) (int64, error) {
	return s.repository.InsertImage(ctx, image)
}

func (s *Service) GetImageByDate(ctx context.Context, date time.Time) (*model.ImageModel, error) {
	return s.repository.GetImageByDate(ctx, date)
}

func (s *Service) DeleteImageByDate(ctx context.Context, date time.Time) (int64, error) {
	return s.repository.DeleteImageByDate(ctx, date)
}

func (s *Service) GetAllImages(ctx context.Context) ([]*model.ImageModel, error) {
	return s.repository.GetAllImages(ctx)
}
