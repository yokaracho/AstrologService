package db

const (
	insertImageQuery = `INSERT INTO images
				   (title, "date", url, hd_url, media_type, explanation, thumbnail_url,  copyright) 
				   VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	getImageByDate = `SELECT id, title, "date", url, hd_url, media_type, explanation, thumbnail_url,  copyright 
					  FROM images WHERE "date" = $1`

	deleteImageByDateQuery = `DELETE FROM images WHERE "date" = $1`

	getAllImage = `SELECT id, title, "date", url, hd_url, media_type, explanation, thumbnail_url,  copyright 
					  FROM images`
)
