package db

import (
	"AstrologService/pkg/model"
	"context"
	"fmt"
	"time"
)

func (r *Repository) InsertImage(ctx context.Context, image *model.ImageModel) (int64, error) {
	result, err := r.pool.Exec(ctx, insertImageQuery, image.Title, time.Now().Format(time.DateOnly), image.URL, image.HDURL, image.MediaType, image.Explanation, image.ThumbURL, image.Copyright)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}

func (r *Repository) GetImageByDate(ctx context.Context, date time.Time) (*model.ImageModel, error) {
	var image model.ImageModel
	err := r.pool.QueryRow(ctx, getImageByDate, date.Format(time.DateOnly)).Scan(&image.ID, &image.Title, &image.Date, &image.URL, &image.HDURL, &image.MediaType, &image.Explanation, &image.ThumbURL, &image.Copyright)
	fmt.Println(date)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *Repository) GetAllImages(ctx context.Context) ([]*model.ImageModel, error) {
	var images []*model.ImageModel
	rows, err := r.pool.Query(ctx, getAllImage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var image model.ImageModel

		if err := rows.Scan(&image.ID, &image.Title, &image.Date, &image.URL, &image.HDURL, &image.MediaType, &image.Explanation, &image.ThumbURL, &image.Copyright); err != nil {
			return nil, err
		}
		images = append(images, &image)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

func (r *Repository) DeleteImageByDate(ctx context.Context, date time.Time) (int64, error) {
	result, err := r.pool.Exec(ctx, deleteImageByDateQuery, date)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}
